package tests

import (
	"testing"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/container/dynamicContainers"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/test"
)


func graphReadInterface[T any, U any](c dynamicContainers.ReadDirectedGraph[T, U])   {}
func graphWriteInterface[T any, U any](c dynamicContainers.WriteDirectedGraph[T, U]) {}
func graphInterface[T any, U any](c dynamicContainers.DirectedGraph[T, U])           {}

// Tests that the value supplied by the factory implements the
// [containerTypes.RWSyncable] interface.
func DynDirectedGraphInterfaceSyncableInterface[V any, E any](
	factory func(capacity int) dynamicContainers.DirectedGraph[V, E],
	t *testing.T,
) {
	var container containerTypes.RWSyncable = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.RWSyncable] interface.
func DynDirectedGraphInterfaceAddressableInterface[V any, E any](
	factory func(capacity int) dynamicContainers.DirectedGraph[V, E],
	t *testing.T,
) {
	var container containerTypes.Addressable = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Clear] interface.
func DynDirectedGraphInterfaceClearInterface[V any, E any](
	factory func(capacity int) dynamicContainers.DirectedGraph[V, E],
	t *testing.T,
) {
	var container containerTypes.Clear = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.DirectedGraphRead] interface.
func ReadDynDirectedGraphInterface[V any, E any](
	factory func(capacity int) dynamicContainers.DirectedGraph[V, E],
	t *testing.T,
) {
	graphReadInterface[V, E](factory(0))
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.WriteDirectedGraph] interface.
func WriteDynDirectedGraphInterface[V any, E any](
	factory func(capacity int) dynamicContainers.DirectedGraph[V, E],
	t *testing.T,
) {
	graphWriteInterface[V, E](factory(0))
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.DirectedGraph] interface.
func DynDirectedGraphInterfaceInterface[V any, E any](
	factory func(capacity int) dynamicContainers.DirectedGraph[V, E],
	t *testing.T,
) {
	graphInterface[V, E](factory(0))
}

// Tests that the value supplied by the factory does not implement the
// [staticContainers.Map] interface.
func DynDirectedGraphInterfaceStaticCapacityInterface[V any, E any](
	factory func(capacity int) dynamicContainers.DirectedGraph[V, E],
	t *testing.T,
) {
	test.Panics(
		func() {
			var c any
			c = factory(0)
			c2 := c.(containerTypes.StaticCapacity)
			_ = c2
		},
		t,
	)
}

func graphContainsEdgeHelper(
	g dynamicContainers.DirectedGraph[int,int],
	l int,
	t *testing.T,
){
	for i:=0; i<l && g.AddEdges(i)==nil; i++ {}
	for i:=0; i<l; i++ {
		test.True(g.ContainsEdge(i),t)
	}
	test.False(g.ContainsEdge(-1),t)
	test.False(g.ContainsEdge(l),t)
}
// Tests the ContainsEdge method functionality of a dynamic graph
func DynDirectedGraphInterfaceContainsEdge(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	graphContainsEdgeHelper(factory(0), 0, t)
	graphContainsEdgeHelper(factory(0), 1, t)
	graphContainsEdgeHelper(factory(0), 2, t)
	graphContainsEdgeHelper(factory(0), 5, t)
}

func graphContainsEdgePntrHelper(
	g dynamicContainers.DirectedGraph[int,int],
	l int,
	t *testing.T,
){
	for i:=0; i<l && g.AddEdges(i)==nil; i++ {}
	for i:=0; i<l; i++ {
		test.True(g.ContainsEdgePntr(&i),t)
	}
	tmp:=-1
	test.False(g.ContainsEdgePntr(&tmp),t)
	test.False(g.ContainsEdgePntr(&l),t)
}
// Tests the ContainsEdgePntr method functionality of a dynamic graph
func DynDirectedGraphInterfaceContainsEdgePntr(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	graphContainsEdgePntrHelper(factory(0), 0, t)
	graphContainsEdgePntrHelper(factory(0), 1, t)
	graphContainsEdgePntrHelper(factory(0), 2, t)
	graphContainsEdgePntrHelper(factory(0), 5, t)
}

func graphContainsVertexHelper(
	g dynamicContainers.DirectedGraph[int,int],
	l int,
	t *testing.T,
){
	for i:=0; i<l && g.AddVertices(i)==nil; i++ {}
	for i:=0; i<l; i++ {
		test.True(g.ContainsVertex(i),t)
	}
	test.False(g.ContainsVertex(-1),t)
	test.False(g.ContainsVertex(l),t)
}
// Tests the ContainsVertex method functionality of a dynamic graph
func DynDirectedGraphInterfaceContainsVertex(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	graphContainsVertexHelper(factory(0), 0, t)
	graphContainsVertexHelper(factory(0), 1, t)
	graphContainsVertexHelper(factory(0), 2, t)
	graphContainsVertexHelper(factory(0), 5, t)
}

func graphContainsVertexPntrHelper(
	g dynamicContainers.DirectedGraph[int,int],
	l int,
	t *testing.T,
){
	for i:=0; i<l && g.AddVertices(i)==nil; i++ {}
	for i:=0; i<l; i++ {
		test.True(g.ContainsVertexPntr(&i),t)
	}
	tmp:=-1
	test.False(g.ContainsVertexPntr(&tmp),t)
	test.False(g.ContainsVertexPntr(&l),t)
}
// Tests the ContainsVertexPntr method functionality of a dynamic graph
func DynDirectedGraphInterfaceContainsVertexPntr(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	graphContainsVertexPntrHelper(factory(0), 0, t)
	graphContainsVertexPntrHelper(factory(0), 1, t)
	graphContainsVertexPntrHelper(factory(0), 2, t)
	graphContainsVertexPntrHelper(factory(0), 5, t)
}

func testDirectedGraphEdgesHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	l int,
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < l; i++ {
		container.AddEdges(i)
	}
	cnt := 0
	container.Edges().ForEach(func(index, val int) (iter.IteratorFeedback, error) {
		cnt++
		test.True(container.ContainsEdge(val),t)
		return iter.Continue, nil
	})
	test.Eq(l, cnt, t)
}
// Tests the Edges method functionality of a dynamic directed graph.
func DynDirectedGraphInterfaceEdges(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
){
	testDirectedGraphEdgesHelper(factory, 0, t)
	testDirectedGraphEdgesHelper(factory, 1, t)
	testDirectedGraphEdgesHelper(factory, 2, t)
	testDirectedGraphEdgesHelper(factory, 5, t)
}

func testDirectedGraphEdgePntrsHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	l int,
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < l; i++ {
		container.AddEdges(i)
	}
	cnt := 0
	container.EdgePntrs().ForEach(
		func(index int, val *int) (iter.IteratorFeedback, error) {
			cnt++
			test.True(container.ContainsEdge(*val), t)
			return iter.Continue, nil
		},
	)
	test.Eq(l, cnt, t)
}

// Tests the EdgePntrs method functionality of a dynamic directed graph.
func DynMapInterfaceEdgePntrs(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	container := factory(0)
	if container.IsAddressable() {
		testDirectedGraphEdgePntrsHelper(factory, 0, t)
		testDirectedGraphEdgePntrsHelper(factory, 1, t)
		testDirectedGraphEdgePntrsHelper(factory, 2, t)
		testDirectedGraphEdgePntrsHelper(factory, 5, t)
	} else {
		test.Panics(func() { container.EdgePntrs() }, t)
	}
}

func testDirectedGraphVerticesHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	l int,
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < l; i++ {
		container.AddVertices(i)
	}
	cnt := 0
	container.Vertices().ForEach(
		func(index, val int) (iter.IteratorFeedback, error) {
			cnt++
			test.True(container.ContainsVertex(val),t)
			return iter.Continue, nil
		},
	)
	test.Eq(l, cnt, t)
}
// Tests the Vertices method functionality of a dynamic directed graph.
func DynDirectedGraphInterfaceVertices(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
){
	testDirectedGraphVerticesHelper(factory, 0, t)
	testDirectedGraphVerticesHelper(factory, 1, t)
	testDirectedGraphVerticesHelper(factory, 2, t)
	testDirectedGraphVerticesHelper(factory, 5, t)
}

func testDirectedGraphVerticePntrsHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	l int,
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < l; i++ {
		container.AddEdges(i)
	}
	cnt := 0
	container.VerticePntrs().ForEach(
		func(index int, val *int) (iter.IteratorFeedback, error) {
			cnt++
			test.True(container.ContainsEdge(*val), t)
			return iter.Continue, nil
		},
	)
	test.Eq(l, cnt, t)
}

// Tests the VerticePntrs method functionality of a dynamic directed graph.
func DynDirectedGraphInterfaceVerticePntrs(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	container := factory(0)
	if container.IsAddressable() {
		testDirectedGraphVerticePntrsHelper(factory, 0, t)
		testDirectedGraphVerticePntrsHelper(factory, 1, t)
		testDirectedGraphVerticePntrsHelper(factory, 2, t)
		testDirectedGraphVerticePntrsHelper(factory, 5, t)
	} else {
		test.Panics(func() { container.VerticePntrs() }, t)
	}
}

func directedGraphLinkHelper(
	container dynamicContainers.DirectedGraph[int, int],
	numVertices int,
	numEdges int,
	links [][3]int,
	t *testing.T,
) {
	for i:=0; i<numEdges; i++ {
		test.Nil(container.AddEdges(i),t)
	}
	for i:=0; i<numVertices; i++ {
		test.Nil(container.AddVertices(i),t)
	}
	for i:=0; i<len(links); i++ {
		test.Nil(container.Link(links[i][0],links[i][1],links[i][2]),t)
	}
	for i:=0; i<len(links); i++ {
		test.True(container.ContainsLink(links[i][0],links[i][1],links[i][2]),t)
	}
	test.Eq(len(links),container.NumLinks(),t)

	for i:=0; i<len(links); i++ {
		test.Nil(container.Link(links[i][0],links[i][1],links[i][2]),t)
	}
	test.Eq(len(links),container.NumLinks(),t)

	err:=container.Link(-1,0,0)
	test.ContainsError(customerr.InvalidValue,err,t)
	test.ContainsError(containerTypes.KeyError,err,t)
	err=container.Link(0,-1,0)
	test.ContainsError(customerr.InvalidValue,err,t)
	test.ContainsError(containerTypes.KeyError,err,t)
	err=container.Link(0,0,-1)
	test.ContainsError(customerr.InvalidValue,err,t)
	test.ContainsError(containerTypes.KeyError,err,t)
}
// Tests the Link method functionality of a dynamic directed graph.
func DynDirectedGraphLink(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
){
	directedGraphLinkHelper(
		factory(0), 5, 4,
		[][3]int{
			// from, to, e
			[3]int{0,1,0},
			[3]int{1,2,1},
			[3]int{2,3,2},
			[3]int{3,4,3},
		},
		t,
	)
	directedGraphLinkHelper(
		factory(0), 5, 5,
		[][3]int{
			// from, to, e
			[3]int{0,1,0},
			[3]int{0,2,1},
			[3]int{0,3,2},
			[3]int{0,4,3},
			[3]int{1,2,1},
			[3]int{1,3,2},
			[3]int{1,4,3},
			[3]int{2,3,2},
			[3]int{2,4,3},
			[3]int{3,4,4},
		},
		t,
	)
}

func directedGraphLinkPntrHelper(
	container dynamicContainers.DirectedGraph[int, int],
	numVertices int,
	numEdges int,
	links [][3]int,
	t *testing.T,
) {
	for i:=0; i<numEdges; i++ {
		test.Nil(container.AddEdges(i),t)
	}
	for i:=0; i<numVertices; i++ {
		test.Nil(container.AddVertices(i),t)
	}
	for i:=0; i<len(links); i++ {
		test.Nil(container.LinkPntr(&links[i][0],&links[i][1],&links[i][2]),t)
	}
	for i:=0; i<len(links); i++ {
		test.True(container.ContainsLink(links[i][0],links[i][1],links[i][2]),t)
	}
	test.Eq(len(links),container.NumLinks(),t)

	for i:=0; i<len(links); i++ {
		test.Nil(container.LinkPntr(&links[i][0],&links[i][1],&links[i][2]),t)
	}
	test.Eq(len(links),container.NumLinks(),t)

	tmp1:=-1
	tmp2:=0
	err:=container.LinkPntr(&tmp1,&tmp2,&tmp2)
	test.ContainsError(customerr.InvalidValue,err,t)
	test.ContainsError(containerTypes.KeyError,err,t)
	err=container.LinkPntr(&tmp2,&tmp1,&tmp2)
	test.ContainsError(customerr.InvalidValue,err,t)
	test.ContainsError(containerTypes.KeyError,err,t)
	err=container.LinkPntr(&tmp2,&tmp2,&tmp1)
	test.ContainsError(customerr.InvalidValue,err,t)
	test.ContainsError(containerTypes.KeyError,err,t)
}
// Tests the LinkPntr method functionality of a dynamic directed graph.
func DynDirectedGraphLinkPntr(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
){
	directedGraphLinkPntrHelper(
		factory(0), 5, 4,
		[][3]int{
			// from, to, e
			[3]int{0,1,0},
			[3]int{1,2,1},
			[3]int{2,3,2},
			[3]int{3,4,3},
		},
		t,
	)
	directedGraphLinkPntrHelper(
		factory(0), 5, 5,
		[][3]int{
			// from, to, e
			[3]int{0,1,0},
			[3]int{0,2,1},
			[3]int{0,3,2},
			[3]int{0,4,3},
			[3]int{1,2,1},
			[3]int{1,3,2},
			[3]int{1,4,3},
			[3]int{2,3,2},
			[3]int{2,4,3},
			[3]int{3,4,4},
		},
		t,
	)
}

func directedGraphOutEdgesHelper(
	container dynamicContainers.DirectedGraph[int, int],
	numVertices int,
	numEdges int,
	links [][3]int,
	outEdges map[int][]int,
	t *testing.T,
) {
	for i:=0; i<numEdges; i++ {
		test.Nil(container.AddEdges(i),t)
	}
	for i:=0; i<numVertices; i++ {
		test.Nil(container.AddVertices(i),t)
	}
	for i:=0; i<len(links); i++ {
		test.Nil(container.Link(links[i][0],links[i][1],links[i][2]),t)
	}
	for i:=0; i<len(links); i++ {
		test.True(container.ContainsLink(links[i][0],links[i][1],links[i][2]),t)
	}
	test.Eq(len(links),container.NumLinks(),t)

	for i:=0; i<numVertices; i++ {
		mappedValues,err:=container.OutEdges(i).Collect()
		test.Eq(len(outEdges[i]),len(mappedValues),t)
		test.Nil(err,t)
		test.SlicesMatchUnordered[int](outEdges[i], mappedValues, t)
	}

	err:=container.OutEdges(-1).Consume()
	test.ContainsError(customerr.InvalidValue,err,t)
	test.ContainsError(containerTypes.KeyError,err,t)
}
// Tests the OutEdges method functionality of a dynamic directed graph.
func DynDirectedGraphOutEdges(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
){
	directedGraphOutEdgesHelper(
		factory(0), 6, 4,
		[][3]int{
			// from, to, e
			[3]int{0,1,0},
			[3]int{1,2,1},
			[3]int{2,3,2},
			[3]int{3,4,3},
		},
		map[int][]int{
			// from, e
			0: []int{0},
			1: []int{1},
			2: []int{2},
			3: []int{3},
			4: []int{},
			5: []int{},
		},
		t,
	)
	directedGraphOutEdgesHelper(
		factory(0), 5, 5,
		[][3]int{
			// from, to, e
			[3]int{0,1,0},
			[3]int{0,2,1},
			[3]int{0,3,2},
			[3]int{0,4,3},
			[3]int{1,2,1},
			[3]int{1,3,2},
			[3]int{1,4,3},
			[3]int{2,3,2},
			[3]int{2,4,3},
			[3]int{3,4,4},
		},
		map[int][]int{
			// from, e
			0: []int{0,1,2,3},
			1: []int{1,2,3},
			2: []int{2,3},
			3: []int{4},
			4: []int{},
			5: []int{},
		},
		t,
	)
}

func directedGraphOutEdgePntrsHelper(
	container dynamicContainers.DirectedGraph[int, int],
	numVertices int,
	numEdges int,
	links [][3]int,
	outEdges map[int][]int,
	t *testing.T,
) {
	for i:=0; i<numEdges; i++ {
		test.Nil(container.AddEdges(i),t)
	}
	for i:=0; i<numVertices; i++ {
		test.Nil(container.AddVertices(i),t)
	}
	for i:=0; i<len(links); i++ {
		test.Nil(container.Link(links[i][0],links[i][1],links[i][2]),t)
	}
	for i:=0; i<len(links); i++ {
		test.True(container.ContainsLink(links[i][0],links[i][1],links[i][2]),t)
	}
	test.Eq(len(links),container.NumLinks(),t)

	for i:=0; i<numVertices; i++ {
		mappedValues,err:=iter.PntrToVal[int](
			container.OutEdgePntrs(&i),
		).Collect()
		test.Eq(len(outEdges[i]),len(mappedValues),t)
		test.Nil(err,t)
		test.SlicesMatchUnordered[int](outEdges[i], mappedValues, t)
	}

	err:=container.OutEdges(-1).Consume()
	test.ContainsError(customerr.InvalidValue,err,t)
	test.ContainsError(containerTypes.KeyError,err,t)
}
// Tests the OutEdgePntrs method functionality of a dynamic directed graph.
func DynDirectedGraphOutEdgePntrs(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
){
	container:=factory(0)
	if container.IsAddressable() {
		directedGraphOutEdgesHelper(
			factory(0), 6, 4,
			[][3]int{
				// from, to, e
				[3]int{0,1,0},
				[3]int{1,2,1},
				[3]int{2,3,2},
				[3]int{3,4,3},
			},
			map[int][]int{
				// from, e
				0: []int{0},
				1: []int{1},
				2: []int{2},
				3: []int{3},
				4: []int{},
				5: []int{},
			},
			t,
		)
		directedGraphOutEdgesHelper(
			factory(0), 5, 5,
			[][3]int{
				// from, to, e
				[3]int{0,1,0},
				[3]int{0,2,1},
				[3]int{0,3,2},
				[3]int{0,4,3},
				[3]int{1,2,1},
				[3]int{1,3,2},
				[3]int{1,4,3},
				[3]int{2,3,2},
				[3]int{2,4,3},
				[3]int{3,4,4},
			},
			map[int][]int{
				// from, e
				0: []int{0,1,2,3},
				1: []int{1,2,3},
				2: []int{2,3},
				3: []int{4},
				4: []int{},
				5: []int{},
			},
			t,
		)
	} else {
		tmp:=0
		test.Panics(func() { container.OutEdgePntrs(&tmp) }, t)
	}
}

func directedGraphOutVerticesHelper(
	container dynamicContainers.DirectedGraph[int, int],
	numVertices int,
	numEdges int,
	links [][3]int,
	outVertices map[int][]int,
	t *testing.T,
) {
	for i:=0; i<numEdges; i++ {
		test.Nil(container.AddEdges(i),t)
	}
	for i:=0; i<numVertices; i++ {
		test.Nil(container.AddVertices(i),t)
	}
	for i:=0; i<len(links); i++ {
		test.Nil(container.Link(links[i][0],links[i][1],links[i][2]),t)
	}
	for i:=0; i<len(links); i++ {
		test.True(container.ContainsLink(links[i][0],links[i][1],links[i][2]),t)
	}
	test.Eq(len(links),container.NumLinks(),t)

	for i:=0; i<numVertices; i++ {
		mappedValues,err:=container.OutVertices(i).Collect()
		test.Eq(len(outVertices[i]),len(mappedValues),t)
		test.Nil(err,t)
		test.SlicesMatchUnordered[int](outVertices[i], mappedValues, t)
	}

	err:=container.OutVertices(-1).Consume()
	test.ContainsError(customerr.InvalidValue,err,t)
	test.ContainsError(containerTypes.KeyError,err,t)
}
// Tests the OutVertices method functionality of a dynamic directed graph.
func DynDirectedGraphOutVertices(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
){
	directedGraphOutVerticesHelper(
		factory(0), 6, 4,
		[][3]int{
			// from, to, e
			[3]int{0,1,0},
			[3]int{1,2,1},
			[3]int{2,3,2},
			[3]int{3,4,3},
		},
		map[int][]int{
			// from, to
			0: []int{1},
			1: []int{2},
			2: []int{3},
			3: []int{4},
			4: []int{},
			5: []int{},
		},
		t,
	)
	directedGraphOutVerticesHelper(
		factory(0), 5, 5,
		[][3]int{
			// from, to, e
			[3]int{0,1,0},
			[3]int{0,2,1},
			[3]int{0,3,2},
			[3]int{0,4,3},
			[3]int{1,2,1},
			[3]int{1,3,2},
			[3]int{1,4,3},
			[3]int{2,3,2},
			[3]int{2,4,3},
			[3]int{3,4,4},
		},
		map[int][]int{
			// from, to
			0: []int{1,2,3,4},
			1: []int{2,3,4},
			2: []int{3,4},
			3: []int{4},
			4: []int{},
			5: []int{},
		},
		t,
	)
}

func directedGraphOutVerticePntrsHelper(
	container dynamicContainers.DirectedGraph[int, int],
	numVertices int,
	numEdges int,
	links [][3]int,
	outVertices map[int][]int,
	t *testing.T,
) {
	for i:=0; i<numEdges; i++ {
		test.Nil(container.AddEdges(i),t)
	}
	for i:=0; i<numVertices; i++ {
		test.Nil(container.AddVertices(i),t)
	}
	for i:=0; i<len(links); i++ {
		test.Nil(container.Link(links[i][0],links[i][1],links[i][2]),t)
	}
	for i:=0; i<len(links); i++ {
		test.True(container.ContainsLink(links[i][0],links[i][1],links[i][2]),t)
	}
	test.Eq(len(links),container.NumLinks(),t)

	for i:=0; i<numVertices; i++ {
		mappedValues,err:=iter.PntrToVal[int](
			container.OutVerticePntrs(&i),
		).Collect()
		test.Eq(len(outVertices[i]),len(mappedValues),t)
		test.Nil(err,t)
		test.SlicesMatchUnordered[int](outVertices[i], mappedValues, t)
	}

	tmp:=-1
	err:=container.OutVerticePntrs(&tmp).Consume()
	test.ContainsError(customerr.InvalidValue,err,t)
	test.ContainsError(containerTypes.KeyError,err,t)
}
// Tests the OutVerticePntrs method functionality of a dynamic directed graph.
func DynDirectedGraphOutVerticePntrs(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
){
	container:=factory(0)
	if container.IsAddressable() {
		directedGraphOutVerticePntrsHelper(
			factory(0), 6, 4,
			[][3]int{
				// from, to, e
				[3]int{0,1,0},
				[3]int{1,2,1},
				[3]int{2,3,2},
				[3]int{3,4,3},
			},
			map[int][]int{
				// from, to
				0: []int{1},
				1: []int{2},
				2: []int{3},
				3: []int{4},
				4: []int{},
				5: []int{},
			},
			t,
		)
		directedGraphOutVerticePntrsHelper(
			factory(0), 5, 5,
			[][3]int{
				// from, to, e
				[3]int{0,1,0},
				[3]int{0,2,1},
				[3]int{0,3,2},
				[3]int{0,4,3},
				[3]int{1,2,1},
				[3]int{1,3,2},
				[3]int{1,4,3},
				[3]int{2,3,2},
				[3]int{2,4,3},
				[3]int{3,4,4},
			},
			map[int][]int{
				// from, to
				0: []int{1,2,3,4},
				1: []int{2,3,4},
				2: []int{3,4},
				3: []int{4},
				4: []int{},
				5: []int{},
			},
			t,
		)
	} else {
		tmp:=0
		test.Panics(func() { container.OutEdgePntrs(&tmp) }, t)
	}
}

func directedGraphOutEdgesAndVerticesHelper(
	container dynamicContainers.DirectedGraph[int, int],
	numVertices int,
	numEdges int,
	links [][3]int,
	outEdges map[int][]int,
	outVertices map[int][]int,
	t *testing.T,
) {
	for i:=0; i<numEdges; i++ {
		test.Nil(container.AddEdges(i),t)
	}
	for i:=0; i<numVertices; i++ {
		test.Nil(container.AddVertices(i),t)
	}
	for i:=0; i<len(links); i++ {
		test.Nil(container.Link(links[i][0],links[i][1],links[i][2]),t)
	}
	for i:=0; i<len(links); i++ {
		test.True(container.ContainsLink(links[i][0],links[i][1],links[i][2]),t)
	}
	test.Eq(len(links),container.NumLinks(),t)

	for i:=0; i<numVertices; i++ {
		iterOutEdgesAndVertices,err:=container.OutEdgesAndVertices(i).Collect()
		test.Nil(err,t)
		test.Eq(len(outEdges[i]),len(iterOutEdgesAndVertices),t)

		iterOutEdges,err:=iter.Map[basic.Pair[int,int],int](
			iter.SliceElems[basic.Pair[int,int]](iterOutEdgesAndVertices),
			func(index int, val basic.Pair[int, int]) (int, error) {
				return val.A, nil
			},
		).Collect()
		test.Nil(err,t)
		test.SlicesMatchUnordered[int](outEdges[i], iterOutEdges, t)

		iterOutVertices,err:=iter.Map[basic.Pair[int,int],int](
			iter.SliceElems[basic.Pair[int,int]](iterOutEdgesAndVertices),
			func(index int, val basic.Pair[int, int]) (int, error) {
				return val.B, nil
			},
		).Collect()
		test.Nil(err,t)
		test.SlicesMatchUnordered[int](outVertices[i], iterOutVertices, t)
	}

	err:=container.OutVertices(-1).Consume()
	test.ContainsError(customerr.InvalidValue,err,t)
	test.ContainsError(containerTypes.KeyError,err,t)
}
// Tests the OutEdgesAndVertices method functionality of a dynamic directed graph.
func DynDirectedGraphOutEdgesAndVertices(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
){
	directedGraphOutEdgesAndVerticesHelper(
		factory(0), 6, 4,
		[][3]int{
			// from, to, e
			[3]int{0,1,0},
			[3]int{1,2,1},
			[3]int{2,3,2},
			[3]int{3,4,3},
		},
		map[int][]int{
			// from, e
			0: []int{0},
			1: []int{1},
			2: []int{2},
			3: []int{3},
			4: []int{},
			5: []int{},
		},
		map[int][]int{
			// from, to
			0: []int{1},
			1: []int{2},
			2: []int{3},
			3: []int{4},
			4: []int{},
			5: []int{},
		},
		t,
	)
	directedGraphOutEdgesAndVerticesHelper(
		factory(0), 5, 5,
		[][3]int{
			// from, to, e
			[3]int{0,1,0},
			[3]int{0,2,1},
			[3]int{0,3,2},
			[3]int{0,4,3},
			[3]int{1,2,1},
			[3]int{1,3,2},
			[3]int{1,4,3},
			[3]int{2,3,2},
			[3]int{2,4,3},
			[3]int{3,4,4},
		},
		map[int][]int{
			// from, e
			0: []int{0,1,2,3},
			1: []int{1,2,3},
			2: []int{2,3},
			3: []int{4},
			4: []int{},
			5: []int{},
		},
		map[int][]int{
			// from, to
			0: []int{1,2,3,4},
			1: []int{2,3,4},
			2: []int{3,4},
			3: []int{4},
			4: []int{},
			5: []int{},
		},
		t,
	)
}

func directedGraphOutEdgesAndVerticePntrsHelper(
	container dynamicContainers.DirectedGraph[int, int],
	numVertices int,
	numEdges int,
	links [][3]int,
	outEdges map[int][]int,
	outVertices map[int][]int,
	t *testing.T,
) {
	for i:=0; i<numEdges; i++ {
		test.Nil(container.AddEdges(i),t)
	}
	for i:=0; i<numVertices; i++ {
		test.Nil(container.AddVertices(i),t)
	}
	for i:=0; i<len(links); i++ {
		test.Nil(container.Link(links[i][0],links[i][1],links[i][2]),t)
	}
	for i:=0; i<len(links); i++ {
		test.True(container.ContainsLink(links[i][0],links[i][1],links[i][2]),t)
	}
	test.Eq(len(links),container.NumLinks(),t)

	for i:=0; i<numVertices; i++ {
		iterOutEdgesAndVertices,err:=container.
			OutEdgesAndVerticePntrs(&i).Collect()
		test.Nil(err,t)
		test.Eq(len(outEdges[i]),len(iterOutEdgesAndVertices),t)

		iterOutEdges,err:=iter.Map[basic.Pair[*int,*int],int](
			iter.SliceElems[basic.Pair[*int,*int]](iterOutEdgesAndVertices),
			func(index int, val basic.Pair[*int, *int]) (int, error) {
				return *val.A, nil
			},
		).Collect()
		test.Nil(err,t)
		test.SlicesMatchUnordered[int](outEdges[i], iterOutEdges, t)

		iterOutVertices,err:=iter.Map[basic.Pair[*int,*int],int](
			iter.SliceElems[basic.Pair[*int,*int]](iterOutEdgesAndVertices),
			func(index int, val basic.Pair[*int, *int]) (int, error) {
				return *val.B, nil
			},
		).Collect()
		test.Nil(err,t)
		test.SlicesMatchUnordered[int](outVertices[i], iterOutVertices, t)
	}

	err:=container.OutVertices(-1).Consume()
	test.ContainsError(customerr.InvalidValue,err,t)
	test.ContainsError(containerTypes.KeyError,err,t)
}
// Tests the OutEdgesAndVerticePntrs method functionality of a dynamic directed graph.
func DynDirectedGraphOutEdgesAndVerticePntrs(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
){
	container:=factory(0)
	if container.IsAddressable() {
		directedGraphOutEdgesAndVerticePntrsHelper(
			factory(0), 6, 4,
			[][3]int{
				// from, to, e
				[3]int{0,1,0},
				[3]int{1,2,1},
				[3]int{2,3,2},
				[3]int{3,4,3},
			},
			map[int][]int{
				// from, e
				0: []int{0},
				1: []int{1},
				2: []int{2},
				3: []int{3},
				4: []int{},
				5: []int{},
			},
			map[int][]int{
				// from, to
				0: []int{1},
				1: []int{2},
				2: []int{3},
				3: []int{4},
				4: []int{},
				5: []int{},
			},
			t,
		)
		directedGraphOutEdgesAndVerticePntrsHelper(
			factory(0), 5, 5,
			[][3]int{
				// from, to, e
				[3]int{0,1,0},
				[3]int{0,2,1},
				[3]int{0,3,2},
				[3]int{0,4,3},
				[3]int{1,2,1},
				[3]int{1,3,2},
				[3]int{1,4,3},
				[3]int{2,3,2},
				[3]int{2,4,3},
				[3]int{3,4,4},
			},
			map[int][]int{
				// from, e
				0: []int{0,1,2,3},
				1: []int{1,2,3},
				2: []int{2,3},
				3: []int{4},
				4: []int{},
				5: []int{},
			},
			map[int][]int{
				// from, to
				0: []int{1,2,3,4},
				1: []int{2,3,4},
				2: []int{3,4},
				3: []int{4},
				4: []int{},
				5: []int{},
			},
			t,
		)
	} else {
		tmp:=-1
		test.Panics(func() { container.OutEdgesAndVerticePntrs(&tmp) }, t)
	}
}

func directedGraphEdgesBetweenHelper(
	container dynamicContainers.DirectedGraph[int, int],
	numVertices int,
	numEdges int,
	links [][3]int,
	structure map[int]map[int][]int,
	t *testing.T,
){
	for i:=0; i<numEdges; i++ {
		test.Nil(container.AddEdges(i),t)
	}
	for i:=0; i<numVertices; i++ {
		test.Nil(container.AddVertices(i),t)
	}
	for i:=0; i<len(links); i++ {
		test.Nil(container.Link(links[i][0],links[i][1],links[i][2]),t)
	}
	for i:=0; i<len(links); i++ {
		test.True(container.ContainsLink(links[i][0],links[i][1],links[i][2]),t)
	}
	test.Eq(len(links),container.NumLinks(),t)

	for from,allTo:=range(structure) {
		for to,edges:=range(allTo) {
			res,err:=container.EdgesBetween(from, to).Collect()
			test.Nil(err,t)
			test.SlicesMatchUnordered[int](edges, res, t)
		}
	}
}
// Tests the EdgesBetween method functionality of a dynamic directed graph.
func DynDirectedGraphEdgesBetween(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphEdgesBetweenHelper(
		factory(0),
		3,
		5,
		[][3]int{
			// from, to, e
			[3]int{0,1,0},
			[3]int{0,1,1},
			[3]int{0,1,2},
			[3]int{0,2,3},
			[3]int{0,2,4},
			[3]int{1,0,0},
			[3]int{1,2,1},
		},
		map[int]map[int][]int{
			// from: to: e
			0: map[int][]int{
				// to: e
				0: []int{},
				1: []int{0,1,2},
				2: []int{3,4},
			},
			1: map[int][]int{
				// to: e
				0: []int{0},
				1: []int{},
				2: []int{1},
			},
		},
		t,
	)
}

func directedGraphEdgesBetweenPntrHelper(
	container dynamicContainers.DirectedGraph[int, int],
	numVertices int,
	numEdges int,
	links [][3]int,
	structure map[int]map[int][]int,
	t *testing.T,
){
	for i:=0; i<numEdges; i++ {
		test.Nil(container.AddEdges(i),t)
	}
	for i:=0; i<numVertices; i++ {
		test.Nil(container.AddVertices(i),t)
	}
	for i:=0; i<len(links); i++ {
		test.Nil(container.Link(links[i][0],links[i][1],links[i][2]),t)
	}
	for i:=0; i<len(links); i++ {
		test.True(container.ContainsLink(links[i][0],links[i][1],links[i][2]),t)
	}
	test.Eq(len(links),container.NumLinks(),t)

	for from,allTo:=range(structure) {
		for to,edges:=range(allTo) {
			res,err:=iter.PntrToVal[int](
				container.EdgesBetweenPntr(&from, &to),
			).Collect()
			test.Nil(err,t)
			test.SlicesMatchUnordered[int](edges, res, t)
		}
	}
}
// Tests the EdgesBetweenPntr method functionality of a dynamic directed graph.
func DynDirectedGraphEdgesBetweenPntr(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	container:=factory(0)
	if container.IsAddressable() {
		directedGraphEdgesBetweenPntrHelper(
			factory(0),
			3,
			5,
			[][3]int{
				// from, to, e
				[3]int{0,1,0},
				[3]int{0,1,1},
				[3]int{0,1,2},
				[3]int{0,2,3},
				[3]int{0,2,4},
				[3]int{1,0,0},
				[3]int{1,2,1},
			},
			map[int]map[int][]int{
				// from: to: e
				0: map[int][]int{
					// to: e
					0: []int{},
					1: []int{0,1,2},
					2: []int{3,4},
				},
				1: map[int][]int{
					// to: e
					0: []int{0},
					1: []int{},
					2: []int{1},
				},
			},
			t,
		)
	} else {
		tmp:=-1
		test.Panics(func() { container.EdgesBetweenPntr(&tmp, &tmp) }, t)
	}
}

func directedGraphDeleteLinkHelper(
	container dynamicContainers.DirectedGraph[int, int],
	numVertices int,
	numEdges int,
	links [][3]int,
	deleteLinks [][3]int,
	t *testing.T,
) {
	for i:=0; i<numEdges; i++ {
		test.Nil(container.AddEdges(i),t)
	}
	for i:=0; i<numVertices; i++ {
		test.Nil(container.AddVertices(i),t)
	}
	for i:=0; i<len(links); i++ {
		test.Nil(container.Link(links[i][0],links[i][1],links[i][2]),t)
	}
	for i:=0; i<len(links); i++ {
		test.True(container.ContainsLink(links[i][0],links[i][1],links[i][2]),t)
	}
	test.Eq(len(links),container.NumLinks(),t)

	for i:=0; i<len(deleteLinks); i++ {
		test.Nil(
			container.DeleteLink(
				deleteLinks[i][0],
				deleteLinks[i][1],
				deleteLinks[i][2],
			),
			t,
		)
	}
	for i:=0; i<len(deleteLinks); i++ {
		test.False(container.ContainsLink(links[i][0],links[i][1],links[i][2]),t)
	}
	test.Eq(len(links)-len(deleteLinks),container.NumLinks(),t)
}
// Tests the DeleteLink method functionality of a dynamic directed graph.
func DynDirectedGraphDeleteLink(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphDeleteLinkPntrHelper(
		factory(0),
		3,
		5,
		[][3]int{
			// from, to, e
			[3]int{0,1,0},
			[3]int{0,1,1},
			[3]int{0,1,2},
			[3]int{0,2,3},
			[3]int{0,2,4},
			[3]int{1,0,0},
			[3]int{1,2,1},
		},
		[][3]int{
			[3]int{0,1,0},
			[3]int{0,1,1},
			[3]int{1,0,0},
		},
		t,
	)
}

// TODO - need to test the error portion of all methods (invalid vertex and edges)
func directedGraphDeleteLinkPntrHelper(
	container dynamicContainers.DirectedGraph[int, int],
	numVertices int,
	numEdges int,
	links [][3]int,
	deleteLinks [][3]int,
	t *testing.T,
) {
	for i:=0; i<numEdges; i++ {
		test.Nil(container.AddEdges(i),t)
	}
	for i:=0; i<numVertices; i++ {
		test.Nil(container.AddVertices(i),t)
	}
	for i:=0; i<len(links); i++ {
		test.Nil(container.Link(links[i][0],links[i][1],links[i][2]),t)
	}
	for i:=0; i<len(links); i++ {
		test.True(container.ContainsLink(links[i][0],links[i][1],links[i][2]),t)
	}
	test.Eq(len(links),container.NumLinks(),t)

	for i:=0; i<len(deleteLinks); i++ {
		test.Nil(
			container.DeleteLinkPntr(
				&deleteLinks[i][0],
				&deleteLinks[i][1],
				&deleteLinks[i][2],
			),
			t,
		)
	}
	for i:=0; i<len(deleteLinks); i++ {
		test.False(
			container.ContainsLink(
				deleteLinks[i][0],
				deleteLinks[i][1],
				deleteLinks[i][2],
			),
			t,
		)
	}
	test.Eq(len(links)-len(deleteLinks),container.NumLinks(),t)
}
// Tests the DeleteLinkPntr method functionality of a dynamic directed graph.
func DynDirectedGraphDeleteLinkPntr(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphDeleteLinkPntrHelper(
		factory(0),
		3,
		5,
		[][3]int{
			// from, to, e
			[3]int{0,1,0},
			[3]int{0,1,1},
			[3]int{0,1,2},
			[3]int{0,2,3},
			[3]int{0,2,4},
			[3]int{1,0,0},
			[3]int{1,2,1},
		},
		[][3]int{
			[3]int{0,1,0},
			[3]int{0,1,1},
			[3]int{1,0,0},
		},
		t,
	)
}
