package containers

import (
	"testing"

	"github.com/barbell-math/util/algo/widgets"
	"github.com/barbell-math/util/test"
)

//go:generate go run interfaceTest.go -type=HashGraph -category=dynamic -interface=DirectedGraph -genericDecl=[int,int] -factory=generateHashGraph
//go:generate go run interfaceTest.go -type=SyncedHashGraph -category=dynamic -interface=DirectedGraph -genericDecl=[int,int] -factory=generateSyncedHashGraph

func generateHashGraph(
	capacity int,
) HashGraph[int, int, widgets.BuiltinInt, widgets.BuiltinInt] {
	v, _ := NewHashGraph[
		int, int,
		widgets.BuiltinInt, widgets.BuiltinInt,
	](capacity, capacity)
	return v
}

func generateSyncedHashGraph(
	capacity int,
) SyncedHashGraph[int, int, widgets.BuiltinInt, widgets.BuiltinInt] {
	v, _ := NewSyncedHashGraph[
		int, int,
		widgets.BuiltinInt, widgets.BuiltinInt,
	]( capacity, capacity)
	return v
}

func TestHashGraphEq(t *testing.T) {
	g1,err:=NewHashGraph[
		string, int,
		widgets.BuiltinString, widgets.BuiltinInt,
	](0,0)
	test.Nil(err,t)
	g2,err:=NewHashGraph[
		string, int,
		widgets.BuiltinString, widgets.BuiltinInt,
	](0,0)
	test.Nil(err,t)

	test.Nil(g1.AddVertices("one", "two", "three", "four", "five"),t)
	test.Nil(g1.AddEdges(1,2,3,4),t)
	test.Nil(g1.Link("one", "two", 1),t)
	test.Nil(g1.Link("two", "three", 2),t)
	test.Nil(g1.Link("three", "four", 3),t)
	test.Nil(g1.Link("four", "five", 4),t)

	test.Nil(g2.AddVertices("one", "two", "three", "four", "five"),t)
	test.Nil(g2.AddEdges(1,2,3,4),t)
	test.Nil(g2.Link("one", "two", 1),t)
	test.Nil(g2.Link("two", "three", 2),t)
	test.Nil(g2.Link("three", "four", 3),t)
	test.Nil(g2.Link("four", "five", 4),t)

	test.True(g1.Eq(&g1, &g2),t)
	test.True(g1.Eq(&g2, &g1),t)

	g2.Link("five", "one", 1)

	test.False(g1.Eq(&g1, &g2),t)
	test.False(g1.Eq(&g2, &g1),t)
}

func TestHashGraphLt(t *testing.T) {
	g,err:=NewHashGraph[
		string, int,
		widgets.BuiltinString, widgets.BuiltinInt,
	](0,0)
	test.Nil(err,t)
	test.Panics(
		func() {
			g.Lt(&g, &g)
		},
		t,
	)
}

func TestHashGraphHash(t *testing.T){
	g1,err:=NewHashGraph[
		string, int,
		widgets.BuiltinString, widgets.BuiltinInt,
	](0,0)
	test.Nil(err,t)
	g2,err:=NewHashGraph[
		string, int,
		widgets.BuiltinString, widgets.BuiltinInt,
	](0,0)
	test.Nil(err,t)

	test.Nil(g1.AddVertices("one", "two", "three", "four", "five"),t)
	test.Nil(g1.AddEdges(1,2,3,4),t)
	test.Nil(g1.Link("one", "two", 1),t)
	test.Nil(g1.Link("two", "three", 2),t)
	test.Nil(g1.Link("three", "four", 3),t)
	test.Nil(g1.Link("four", "five", 4),t)

	test.Nil(g2.AddVertices("one", "two", "three", "four", "five"),t)
	test.Nil(g2.AddEdges(1,2,3,4),t)
	test.Nil(g2.Link("one", "two", 1),t)
	test.Nil(g2.Link("two", "three", 2),t)
	test.Nil(g2.Link("three", "four", 3),t)
	test.Nil(g2.Link("four", "five", 4),t)

	test.Eq(g1.Hash(&g1),g1.Hash(&g2),t)
	test.Eq(g2.Hash(&g1),g2.Hash(&g2),t)

	g2.Link("five", "one", 1)
	test.Neq(g1.Hash(&g1),g1.Hash(&g2),t)
	test.Neq(g2.Hash(&g1),g2.Hash(&g2),t)

	h:=g1.Hash(&g1)
	for i:=0; i<100; i++ {
		test.Eq(h, g1.Hash(&g1), t)
	}
}

func TestHashGraphZero(t *testing.T) {
	g1,err:=NewHashGraph[
		string, int,
		widgets.BuiltinString, widgets.BuiltinInt,
	](0,0)
	test.Nil(err,t)

	test.Nil(g1.AddVertices("one", "two", "three", "four", "five"),t)
	test.Nil(g1.AddEdges(1,2,3,4),t)
	test.Nil(g1.Link("one", "two", 1),t)
	test.Nil(g1.Link("two", "three", 2),t)
	test.Nil(g1.Link("three", "four", 3),t)
	test.Nil(g1.Link("four", "five", 4),t)

	g1.Zero(&g1)
	test.Eq(0, g1.NumVertices(), t)
	test.Eq(0, g1.NumEdges(), t)
	test.Eq(0, g1.NumLinks(), t)
	test.Eq(0, len(g1.vertices), t)
	test.Eq(0, len(g1.edges), t)
	test.Eq(0, len(g1.graph), t)
}
