package lexer

import (
	"github.com/barbell-math/util/algo/hash"
	"github.com/barbell-math/util/algo/widgets"

	// "github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/container/containers"
)

type (
	nfaID   int
	nfaFlag int
	// nfaNode struct {
	// 	flags       nfaFlag
	// 	transitions []basic.Pair[byte, nfaID]
	// }

	// NFA map[nfaID]nfaNode

	nfaNode struct {
		id    nfaID
		flags nfaFlag
	}
	alphabetRange[A any, AI widgets.WidgetInterface[A]] containers.HashSet[A, AI]
	NFA[A any, AI widgets.WidgetInterface[A]]           containers.HashGraph[
		nfaNode, containers.HashSet[A, AI],
		*nfaNode, *containers.HashSet[A, AI],
	]
)

const (
	nfaStart  nfaID   = 0
	nfaSource nfaFlag = 1 << iota
	nfaSink
)

func (_ *nfaNode) Eq(l *nfaNode, r *nfaNode) bool {
	return l.id == r.id
}
func (_ *nfaNode) Lt(l *nfaNode, r *nfaNode) bool {
	return l.id < r.id
}
func (_ *nfaNode) Hash(other *nfaNode) hash.Hash {
	return hash.Hash(other.id)
}
func (_ *nfaNode) Zero(other *nfaNode) {
	*other = nfaNode{}
}

//	func (n *nfaNode) addTransition(c byte, id nfaID) {
//		n.transitions = append(n.transitions, basic.Pair[byte, nfaID]{c, id})
//	}
func NewNFA[A any, AI widgets.WidgetInterface[A]]() NFA[A, AI] {
	rv, _ := containers.NewHashGraph[
		nfaNode, containers.HashSet[A, AI],
		*nfaNode, *containers.HashSet[A, AI],
	](0, 0)
	return NFA[A, AI](rv)
}

func (n NFA[A, AI]) initNFA() {
	lambdaEdge, _ := containers.NewHashSet[A, AI](0)
	n.Clear()
	n.AddVertices(nfaNode{
		id:    nfaID(0),
		flags: nfaSource | nfaSink,
	})
	n.AddEdges(lambdaEdge)
}

func (n NFA[A, AI]) AppendTransition(c A) {
	if n.NumVertices() == 0 {
		n.initNFA()
	}
	newNode := nfaNode{
		id:    nfaID(n.NumVertices()),
		flags: nfaSink,
	}
	edge := containers.HashSetValInit[A, AI](c)
	sinkIdentifier := nfaNode{id: nfaID(n.NumVertices()) - 1}

	n.AddVertices(newNode)
	n.AddEdges(edge)
	n.UpdateVertex(sinkIdentifier, func(orig *nfaNode) {
		orig.flags &= ^nfaSink
	})
	n.LinkPntr(&sinkIdentifier, &newNode, &edge)
}

// func (n NFA) AppendNFA(other NFA) {
// 	// Other is nothing more than a just inited NFA, nothing to add
// 	if len(other) <= 1 || other.nfaSink() <= 0 {
// 		return
// 	}
// 	if len(n) == 0 {
// 		n.initNFA()
// 	}
// 	n.changeAndSave(n.nfaSink(), func(node *nfaNode) error {
// 		for _, v := range other[nfaStart].transitions {
// 			node.addTransition(v.A, v.B+n.nfaSink())
// 		}
// 		node.flags &= ^nfaSink
// 		return nil
// 	})
//
// 	oldNFASize := n.nfaSink()
// 	for i := nfaStart + 1; i <= other.nfaSink(); i++ {
// 		iterNode := other[i]
// 		iterNode.flags &= ^nfaSource
// 		for j := 0; j < len(iterNode.transitions); j++ {
// 			iterNode.transitions[j].B += oldNFASize
// 		}
// 		n.addNode(iterNode)
// 	}
// }
//
// func (n NFA) ApplyKleene() {
// 	if len(n) == 0 {
// 		n.initNFA()
// 	}
// 	alreadyKleened := false
// 	for i := 0; i < len(n[n.nfaSink()].transitions) && !alreadyKleened; i++ {
// 		v := n[n.nfaSink()].transitions[i]
// 		alreadyKleened = (v.A == lambdaChar && v.B == n.nfaSink()-1)
// 	}
// 	alreadyKleened = (alreadyKleened && len(n[nfaStart].transitions) == 1)
// 	alreadyKleened = (alreadyKleened && n[nfaStart].transitions[0].A == lambdaChar)
// 	alreadyKleened = (alreadyKleened && n[nfaStart].transitions[0].B == n.nfaSink()-1)
// 	if !alreadyKleened {
// 		n.changeAndSave(n.nfaSink(), func(node *nfaNode) error {
// 			node.flags &= ^nfaSink
// 			node.addTransition(lambdaChar, n.nfaSink()+2)
// 			return nil
// 		})
// 		n.addNode(nfaNode{
// 			flags:       0,
// 			transitions: n[nfaStart].transitions,
// 		})
// 		n.changeAndSave(nfaStart, func(node *nfaNode) error {
// 			node.transitions = []basic.Pair[byte, nfaID]{{lambdaChar, n.nfaSink()}}
// 			return nil
// 		})
// 		n.addNode(nfaNode{
// 			flags:       nfaSink,
// 			transitions: []basic.Pair[byte, nfaID]{{lambdaChar, n.nfaSink()}},
// 		})
// 	}
// }
//
// func (n NFA) AddBranch(other NFA) {
// 	// Other is nothing more than a just inited NFA, nothing to add
// 	if len(other) <= 1 || other.nfaSink() <= 0 {
// 		return
// 	}
// 	// This NFA is nothing more than a just inited NFA, simply copy other
// 	if len(n) <= 1 || n.nfaSink() <= 0 {
// 		n.clearAndCopy(other)
// 		return
// 	}
// 	n.changeAndSave(nfaStart, func(node *nfaNode) error {
// 		for _, v := range other[nfaStart].transitions {
// 			node.addTransition(v.A, v.B+n.nfaSink())
// 		}
// 		return nil
// 	})
// 	n.changeAndSave(n.nfaSink(), func(node *nfaNode) error {
// 		node.flags &= ^nfaSink
// 		node.addTransition(lambdaChar, n.nfaSink()+other.nfaSink()+1)
// 		return nil
// 	})
//
// 	oldNFASize := n.nfaSink()
// 	for i := nfaStart + 1; i < other.nfaSink(); i++ {
// 		iterNode := other[i]
// 		iterNode.flags &= ^(nfaSource | nfaSink)
// 		for j := 0; j < len(iterNode.transitions); j++ {
// 			iterNode.transitions[j].B += oldNFASize
// 		}
// 		n.addNode(iterNode)
// 	}
//
// 	otherEndNode := other[other.nfaSink()]
// 	for j := 0; j < len(otherEndNode.transitions); j++ {
// 		otherEndNode.transitions[j].B += oldNFASize
// 	}
// 	otherEndNode.flags &= ^(nfaSource | nfaSink)
// 	otherEndNode.addTransition(lambdaChar, n.nfaSink()+2)
// 	n.addNode(otherEndNode)
//
// 	n.addNode(nfaNode{
// 		flags:       nfaSink,
// 		transitions: []basic.Pair[byte, nfaID]{},
// 	})
// }
//
// func (n NFA) addNode(node nfaNode) {
// 	n[n.nfaSink()+1] = node
// }
//
// func (n NFA) changeAndSave(id nfaID, op func(node *nfaNode) error) error {
// 	node := n[id]
// 	if err := op(&node); err == nil {
// 		n[id] = node
// 	} else {
// 		return err
// 	}
// 	return nil
// }
//
// func (n NFA) nfaSink() nfaID {
// 	return nfaID(len(n)) - 1
// }
//
// func (n NFA) clearAndCopy(other NFA) {
// 	for k, _ := range n {
// 		delete(n, k)
// 	}
// 	for k, v := range other {
// 		n[k] = v
// 	}
// }
