package dataStruct

import (
    "fmt"

    customerr "github.com/barbell-math/util/err"
)

type node[T any] struct {
    Val T
    Next *node[T]
    Prev *node[T]
}

type Movement[T any] func(n *node[T]) *node[T]
func Forward[T any](n *node[T]) *node[T] { 
    if n!=nil {
        return n.Next 
    }
    return nil
}
func Backward[T any](n *node[T]) *node[T] {
    if n!=nil {
        return n.Prev 
    }
    return nil
}

type DoubleLinkedList[T any] struct {
    startNode *node[T]
    endNode *node[T]
    numNodes int
}

func (d *DoubleLinkedList[T])Length() int {
    return d.numNodes
}

func (d *DoubleLinkedList[T])Capacity() int {
    return d.numNodes
}

func (d *DoubleLinkedList[T])PeekPntrFront(idx int) (*T,error) {
    return d.peekPntr(idx)
}

func (d *DoubleLinkedList[T])PeekFront(idx int) (T,error) {
    if p,err:=d.PeekPntrFront(idx); err==nil {
        return *p,nil
    } else {
        var tmp T
        return tmp,err
    }
}

func (d *DoubleLinkedList[T])PeekPntrBack(idx int) (*T,error) {
    return d.peekPntr(idx)
}

func (d *DoubleLinkedList[T])peekPntr(idx int) (*T,error) {
    if idx<0 || idx>=d.numNodes || d.numNodes<=0 {
        return nil,customerr.ValOutsideRange(fmt.Sprintf(
            "Index out of bounds. | Len: %d Idx: %d",d.numNodes,idx,
        ))
    }
    var n *node[T]
    if d.distanceFromFront(idx)<d.distanceFromBack(idx) {
        n=d.getNode(d.startNode,Forward[T],idx)
    } else {
        n=d.getNode(d.endNode,Backward[T],idx)
    }
    return &n.Val,nil
}

func (d *DoubleLinkedList[T])Peek(idx int) (T,error) {
    var tmp T
    if v,err:=d.peekPntr(idx); err==nil {
        return *v,nil
    } else {
        return tmp,err
    }
}

func (d *DoubleLinkedList[T])distanceFromFront(idx int) int {
    return idx+1
}
func (d *DoubleLinkedList[T])distanceFromBack(idx int) int {
    return d.numNodes-idx
}

func (d *DoubleLinkedList[T])getNode(n *node[T], m Movement[T], idx int) *node[T] {
    i:=0
    for n=d.startNode; n!=nil && i<idx; n=m(n) {}
    return n
}

func (d *DoubleLinkedList[T])PushBack(v T) error {
    n:=node[T]{ Val: v }
    if d.numNodes==0 {
        d.startNode=&n
    } else {
        d.endNode.Next=&n
    }
    n.Prev=d.endNode
    d.endNode=&n
    d.numNodes+=1
    return nil
}

func (d *DoubleLinkedList[T])PushFront(v T) error {
    n:=node[T]{ Val: v }
    if d.numNodes==0 {
        d.endNode=&n
    } else {
        d.startNode.Prev=&n
    }
    n.Next=d.startNode
    d.startNode=&n
    d.numNodes+=1
    return nil
}

func (d *DoubleLinkedList[T])PopBack() (T,error) {
    if d.numNodes<=0 {
        var tmp T
        return tmp,QueueEmpty("Nothing to pop!")
    } else if d.numNodes==1 {
        rv:=d.endNode.Val
        d.endNode=nil
        d.startNode=nil
        d.numNodes-=1
        return rv,nil
    } else if d.numNodes==2 {
        rv:=d.endNode.Val
        d.endNode.Prev=nil
        d.endNode=d.startNode
        d.startNode.Next=nil
        d.numNodes-=1
        return rv,nil
    } else {
        rv:=d.endNode.Val
        d.endNode=d.endNode.Prev
        d.endNode.Next.Prev=nil
        d.endNode.Next=nil
        d.numNodes-=1
        return rv,nil
    }
}

func (d *DoubleLinkedList[T])PopFront() (T,error) {
    if d.numNodes<=0 {
        var tmp T
        return tmp,QueueEmpty("Nothing to pop!")
    } else if d.numNodes==1 {
        rv:=d.startNode.Val
        d.endNode=nil
        d.startNode=nil
        d.numNodes-=1
        return rv,nil
    } else if d.numNodes==2 {
        rv:=d.startNode.Val
        d.startNode.Next=nil
        d.startNode=d.endNode
        d.startNode.Prev=nil
        d.numNodes-=1
        return rv,nil
    } else {
        rv:=d.startNode.Val
        d.startNode=d.startNode.Next
        d.startNode.Prev.Next=nil
        d.startNode.Prev=nil
        d.numNodes-=1
        return rv,nil
    }
}
