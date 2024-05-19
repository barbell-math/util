package lexer

import (
	"github.com/barbell-math/util/container/basic"
)

type (
	nfaID int
	nfaFlag int
	nfaNode struct {
		flags nfaFlag
		transitions []basic.Pair[byte,nfaID]
	}

	NFA struct {
		// curId nfaID
		graph map[nfaID]nfaNode
	}
)

const (
	start nfaID=0
	source nfaFlag=1<<iota
	sink
)

func (n *nfaNode)addTransition(c byte, id nfaID) {
	n.transitions=append(n.transitions, basic.Pair[byte,nfaID]{c,id})
}

func NewNFA() NFA {
	rv:=NFA{}
	rv.initNFA()
	return rv
}

func (n *NFA)initNFA() {
	n.graph=map[nfaID]nfaNode{
		start: {
			flags: source | sink,
			transitions: []basic.Pair[byte,nfaID]{},
		},
	}
	// n.curId=0
}

func (n *NFA)AppendTransition(c byte) {
	if len(n.graph)==0 {
		n.initNFA()
	}
	n.changeAndSave(n.sink(), func(node *nfaNode) error {
		node.flags&=^sink
		node.addTransition(c,n.sink()+1)
		return nil
	})
	n.addNode(nfaNode {
		flags: sink,
		transitions: []basic.Pair[byte,nfaID]{},
	})
}

func (n *NFA)AppendNFA(other *NFA) {
	// Other is nothing more than a just inited NFA, nothing to add
	if len(other.graph)<=1 || other.sink()<=0 {
		return
	}
	if len(n.graph)==0 {
		n.initNFA()
	}
	n.changeAndSave(n.sink(), func(node *nfaNode) error {
		for _,v:=range(other.graph[start].transitions) {
			node.addTransition(v.A,v.B+n.sink())
		}
		node.flags&=^sink
		return nil
	})

	oldNFASize:=n.sink()
	for i:=start+1; i<=other.sink(); i++ {
		iterNode:=other.graph[i]
		iterNode.flags&=^source
		for j:=0; j<len(iterNode.transitions); j++ {
			iterNode.transitions[j].B+=oldNFASize
		}
		n.addNode(iterNode)
	}
}

func (n *NFA)ApplyKleene() {
	if len(n.graph)==0 {
		n.initNFA()
	}
	alreadyKleened:=false
	for i:=0; i<len(n.graph[n.sink()].transitions) && !alreadyKleened; i++ {
		v:=n.graph[n.sink()].transitions[i]
		alreadyKleened=(v.A==lambdaChar && v.B==n.sink()-1)
	}
	alreadyKleened=(alreadyKleened && len(n.graph[start].transitions)==1)
	alreadyKleened=(alreadyKleened && n.graph[start].transitions[0].A==lambdaChar)
	alreadyKleened=(alreadyKleened && n.graph[start].transitions[0].B==n.sink()-1)
	if !alreadyKleened {
		n.changeAndSave(n.sink(), func(node *nfaNode) error {
			node.flags&=^sink
			node.addTransition(lambdaChar, n.sink()+2)
			return nil
		})
		n.addNode(nfaNode{
			flags: 0,
			transitions: n.graph[start].transitions,
		})
		n.changeAndSave(start, func(node *nfaNode) error {
			node.transitions=[]basic.Pair[byte, nfaID]{{lambdaChar, n.sink()}}
			return nil
		})
		n.addNode(nfaNode{
			flags: sink,
			transitions: []basic.Pair[byte, nfaID]{{lambdaChar, n.sink()}},
		})
	}
}

func (n *NFA)AddBranch(other *NFA) {
	// Other is nothing more than a just inited NFA, nothing to add
	if len(other.graph)<=1 || other.sink()<=0 {
		return
	}
	// This NFA is nothing more than a just inited NFA, simply copy other
	if len(n.graph)<=1 || n.sink()<=0 {
		*n=*other
		return
	}
	n.changeAndSave(start, func(node *nfaNode) error {
		for _,v:=range(other.graph[start].transitions) {
			node.addTransition(v.A,v.B+n.sink())
		}
		return nil
	})
	n.changeAndSave(n.sink(), func(node *nfaNode) error {
		node.flags&=^sink
		node.addTransition(lambdaChar, n.sink()+other.sink()+1)
		return nil
	})

	oldNFASize:=n.sink()
	for i:=start+1; i<other.sink(); i++ {
		iterNode:=other.graph[i]
		iterNode.flags&=^(source|sink)
		for j:=0; j<len(iterNode.transitions); j++ {
			iterNode.transitions[j].B+=oldNFASize
		}
		n.addNode(iterNode)
	}

	otherEndNode:=other.graph[other.sink()]
	for j:=0; j<len(otherEndNode.transitions); j++ {
		otherEndNode.transitions[j].B+=oldNFASize
	}
	otherEndNode.flags&=^(source|sink)
	otherEndNode.addTransition(lambdaChar, n.sink()+2)
	n.addNode(otherEndNode)

	n.addNode(nfaNode{
		flags: sink,
		transitions: []basic.Pair[byte,nfaID]{},
	})
}

func (n *NFA)addNode(node nfaNode) {
	// n.curId++
	// n.graph[n.curId]=node
	n.graph[n.sink()+1]=node
}

func (n *NFA)changeAndSave(id nfaID, op func(node *nfaNode) error ) error {
	node:=n.graph[id]
	if err:=op(&node); err==nil {
		n.graph[id]=node
	} else {
		return err
	}
	return nil
}

func (n *NFA)sink() nfaID {
	return nfaID(len(n.graph))-1
}
