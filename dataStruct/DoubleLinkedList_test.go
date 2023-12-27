package dataStruct

import (
	"testing"

	"github.com/barbell-math/util/test"
)

func pointerTestHelper[T any](n *node[T], pntrs Pair[*node[T],*node[T]], t *testing.T){
    test.BasicTest(pntrs.A,n.Prev,
        "A node did not have the correct previous pointer.",t,
    )
    test.BasicTest(pntrs.B,n.Next,
        "A node did not have the correct next pointer.",t,
    )
}
func TestDoubleLinkedListPushBack(t *testing.T) {
    d:=DoubleLinkedList[int]{}
    d.PushBack(1)
    test.BasicTest(d.Length(),1,
        "The length of the double linked list was not incremented.",t,
    )
    test.BasicTest(1,d.endNode.Val,
        "The element was not added to the linked list correctly.",t,
    )
    pointerTestHelper[int](d.endNode,Pair[*node[int], *node[int]]{nil,nil},t)
    test.BasicTest(1,d.startNode.Val,
        "The element was not added to the linked list correctly.",t,
    )
    pointerTestHelper[int](d.endNode,Pair[*node[int], *node[int]]{nil,nil},t)
    d.PushBack(2)
    test.BasicTest(d.Length(),2,
        "The length of the double linked list was not incremented.",t,
    )
    test.BasicTest(1,d.startNode.Val,
        "The element was not added to the linked list correctly.",t,
    )
    test.BasicTest(2,d.endNode.Val,
        "The element was not added to the linked list correctly.",t,
    )
    pointerTestHelper[int](
        d.endNode,Pair[*node[int], *node[int]]{d.startNode,nil},t,
    )
    pointerTestHelper[int](
        d.startNode,Pair[*node[int], *node[int]]{nil,d.endNode},t,
    )
    d.PushBack(3)
    test.BasicTest(d.Length(),3,
        "The length of the double linked list was not incremented.",t,
    )
    test.BasicTest(1,d.startNode.Val,
        "The element was not added to the linked list correctly.",t,
    )
    test.BasicTest(3,d.endNode.Val,
        "The element was not added to the linked list correctly.",t,
    )
    test.BasicTest(2,d.startNode.Next.Val,
        "The element was not added to the linked list correctly.",t,
    )
    pointerTestHelper[int](
        d.endNode,Pair[*node[int], *node[int]]{d.startNode.Next,nil},t,
    )
    pointerTestHelper[int](
        d.startNode,Pair[*node[int], *node[int]]{nil,d.endNode.Prev},t,
    )
    pointerTestHelper[int](
        d.startNode.Next,Pair[*node[int], *node[int]]{d.startNode,d.endNode},t,
    )
}

func TestDoubleLinkedListPushFront(t *testing.T){
    d:=DoubleLinkedList[int]{}
    d.PushFront(1)
    test.BasicTest(d.Length(),1,
        "The length of the double linked list was not incremented.",t,
    )
    test.BasicTest(1,d.endNode.Val,
        "The element was not added to the linked list correctly.",t,
    )
    pointerTestHelper[int](d.endNode,Pair[*node[int], *node[int]]{nil,nil},t)
    test.BasicTest(1,d.startNode.Val,
        "The element was not added to the linked list correctly.",t,
    )
    pointerTestHelper[int](d.endNode,Pair[*node[int], *node[int]]{nil,nil},t)
    d.PushFront(2)
    test.BasicTest(d.Length(),2,
        "The length of the double linked list was not incremented.",t,
    )
    test.BasicTest(2,d.startNode.Val,
        "The element was not added to the linked list correctly.",t,
    )
    test.BasicTest(1,d.endNode.Val,
        "The element was not added to the linked list correctly.",t,
    )
    pointerTestHelper[int](
        d.endNode,Pair[*node[int], *node[int]]{d.startNode,nil},t,
    )
    pointerTestHelper[int](
        d.startNode,Pair[*node[int], *node[int]]{nil,d.endNode},t,
    )
    d.PushFront(3)
    test.BasicTest(d.Length(),3,
        "The length of the double linked list was not incremented.",t,
    )
    test.BasicTest(3,d.startNode.Val,
        "The element was not added to the linked list correctly.",t,
    )
    test.BasicTest(1,d.endNode.Val,
        "The element was not added to the linked list correctly.",t,
    )
    test.BasicTest(2,d.startNode.Next.Val,
        "The element was not added to the linked list correctly.",t,
    )
    pointerTestHelper[int](
        d.endNode,Pair[*node[int], *node[int]]{d.startNode.Next,nil},t,
    )
    pointerTestHelper[int](
        d.startNode,Pair[*node[int], *node[int]]{nil,d.endNode.Prev},t,
    )
    pointerTestHelper[int](
        d.startNode.Next,Pair[*node[int], *node[int]]{d.startNode,d.endNode},t,
    )
}

func TestPopBack(t *testing.T){
    d:=DoubleLinkedList[int]{}
    d.PushBack(1)
    d.PushBack(2)
    d.PushBack(3)
    v,err:=d.PopBack();
    test.BasicTest(3,v,
        "Pop back did not return the correct value.",t,
    )
    test.BasicTest(nil,err,
        "Pop back returned an error when it shouldn't have.",t,
    )
    test.BasicTest(2,d.Length(),
        "The length of the double linked list was not decremented.",t,
    )
    test.BasicTest(2,d.endNode.Val,
        "The last element was not removed correctly.",t,
    )
    pointerTestHelper[int](
        d.endNode,Pair[*node[int], *node[int]]{d.startNode,nil},t,
    )
    v,err=d.PopBack();
    test.BasicTest(2,v,
        "Pop back did not return the correct value.",t,
    )
    test.BasicTest(nil,err,
        "Pop back returned an error when it shouldn't have.",t,
    )
    test.BasicTest(1,d.Length(),
        "The length of the double linked list was not decremented.",t,
    )
    test.BasicTest(1,d.endNode.Val,
        "The last element was not removed correctly.",t,
    )
    pointerTestHelper[int](
        d.endNode,Pair[*node[int], *node[int]]{nil,nil},t,
    )
    v,err=d.PopBack();
    test.BasicTest(1,v,
        "Pop back did not return the correct value.",t,
    )
    test.BasicTest(nil,err,
        "Pop back returned an error when it shouldn't have.",t,
    )
    test.BasicTest(0,d.Length(),
        "The length of the double linked list was not decremented.",t,
    )
    test.BasicTest((*node[int])(nil),d.startNode,
        "The start node was not reset to nil with an empty deque.",t,
    )
    test.BasicTest((*node[int])(nil),d.endNode,
        "The end node was not reset to nil with an empty deque.",t,
    )
    v,err=d.PopBack()
    if !IsQueueEmpty(err) {
        test.FormatError(QueueEmpty(""),err,
            "The correct error was not raised with an empty queue.",t,
        )
    }
}

func TestPopFront(t *testing.T){
    d:=DoubleLinkedList[int]{}
    d.PushBack(1)
    d.PushBack(2)
    d.PushBack(3)
    v,err:=d.PopFront();
    test.BasicTest(1,v,
        "Pop back did not return the correct value.",t,
    )
    test.BasicTest(nil,err,
        "Pop back returned an error when it shouldn't have.",t,
    )
    test.BasicTest(2,d.Length(),
        "The length of the double linked list was not decremented.",t,
    )
    test.BasicTest(2,d.startNode.Val,
        "The first element was not removed correctly.",t,
    )
    pointerTestHelper[int](
        d.startNode,Pair[*node[int], *node[int]]{nil,d.endNode},t,
    )
    v,err=d.PopFront();
    test.BasicTest(2,v,
        "Pop back did not return the correct value.",t,
    )
    test.BasicTest(nil,err,
        "Pop back returned an error when it shouldn't have.",t,
    )
    test.BasicTest(1,d.Length(),
        "The length of the double linked list was not decremented.",t,
    )
    test.BasicTest(3,d.startNode.Val,
        "The first element was not removed correctly.",t,
    )
    pointerTestHelper[int](
        d.startNode,Pair[*node[int], *node[int]]{nil,nil},t,
    )
    v,err=d.PopFront();
    test.BasicTest(3,v,
        "Pop back did not return the correct value.",t,
    )
    test.BasicTest(nil,err,
        "Pop back returned an error when it shouldn't have.",t,
    )
    test.BasicTest(0,d.Length(),
        "The length of the double linked list was not decremented.",t,
    )
    test.BasicTest((*node[int])(nil),d.startNode,
        "The start node was not reset to nil with an empty deque.",t,
    )
    test.BasicTest((*node[int])(nil),d.endNode,
        "The end node was not reset to nil with an empty deque.",t,
    )
    v,err=d.PopFront()
    if !IsQueueEmpty(err) {
        test.FormatError(QueueEmpty(""),err,
            "The correct error was not raised with an empty queue.",t,
        )
    }
}

func TestDoubleLinkedListPeekPntrFront(t *testing.T) {
    // d:=DoubleLinkedList[int]{}
    // _,err:=d.PeekPntrFront() 
    // if !customerr.IsValOutsideRange(err) {
    //     test.FormatError(customerr.IsValOutsideRange(""),err,
    //         "PeekPntrFront returned the incorrect error or no error at all.",t,
    //     )
    // }
}
