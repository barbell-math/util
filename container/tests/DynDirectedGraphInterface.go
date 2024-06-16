package tests

import (
	"testing"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/container/dynamicContainers"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/test"
)


func graphReadInterface[T any, U any](c dynamicContainers.ReadGraph[T, U])   {}
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

func testHashDirectedGraphEdgesHelper(
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
// Tests the Edges method functionality of a dynamic graph.
func DynDirectedGraphInterfaceEdges(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
){
	testHashDirectedGraphEdgesHelper(factory, 0, t)
	testHashDirectedGraphEdgesHelper(factory, 1, t)
	testHashDirectedGraphEdgesHelper(factory, 2, t)
	testHashDirectedGraphEdgesHelper(factory, 5, t)
}

func testHashDirectedGraphEdgePntrsHelper(
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

// Tests the EdgePntrs method functionality of a dynamic graph.
func DynMapInterfaceEdgePntrs(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	container := factory(0)
	if container.IsAddressable() {
		testHashDirectedGraphEdgePntrsHelper(factory, 0, t)
		testHashDirectedGraphEdgePntrsHelper(factory, 1, t)
		testHashDirectedGraphEdgePntrsHelper(factory, 2, t)
		testHashDirectedGraphEdgePntrsHelper(factory, 5, t)
	} else {
		test.Panics(func() { container.EdgePntrs() }, t)
	}
}

func testHashDirectedGraphVerticesHelper(
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
// Tests the Vertices method functionality of a dynamic graph.
func DynDirectedGraphInterfaceVertices(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
){
	testHashDirectedGraphVerticesHelper(factory, 0, t)
	testHashDirectedGraphVerticesHelper(factory, 1, t)
	testHashDirectedGraphVerticesHelper(factory, 2, t)
	testHashDirectedGraphVerticesHelper(factory, 5, t)
}

func testHashDirectedGraphVerticePntrsHelper(
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

// Tests the VerticePntrs method functionality of a dynamic graph.
func DynHashDirectedGraphInterfaceVerticePntrs(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	container := factory(0)
	if container.IsAddressable() {
		testHashDirectedGraphVerticePntrsHelper(factory, 0, t)
		testHashDirectedGraphVerticePntrsHelper(factory, 1, t)
		testHashDirectedGraphVerticePntrsHelper(factory, 2, t)
		testHashDirectedGraphVerticePntrsHelper(factory, 5, t)
	} else {
		test.Panics(func() { container.VerticePntrs() }, t)
	}
}

func hashDirectedGraphLinkHelper(
	container dynamicContainers.DirectedGraph[int, int],
	numVertices int,
	numEdges int,
	links [][3]int,
	t *testing.T,
) {
	for i:=0; i<numVertices; i++ {
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
// Tests the Link method functionality of a dynamic graph.
func DynHashDirectedGraphLink(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
){
	hashDirectedGraphLinkHelper(
		factory(0), 5, 4,
		[][3]int{
			[3]int{0,1,0},
			[3]int{1,2,1},
			[3]int{2,3,2},
			[3]int{3,4,3},
		},
		t,
	)
	hashDirectedGraphLinkHelper(
		factory(0), 5, 5,
		[][3]int{
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

func hashDirectedGraphLinkPntrHelper(
	container dynamicContainers.DirectedGraph[int, int],
	numVertices int,
	numEdges int,
	links [][3]int,
	t *testing.T,
) {
	for i:=0; i<numVertices; i++ {
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
// Tests the LinkPntr method functionality of a dynamic graph.
func DynHashDirectedGraphLinkPntr(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
){
	hashDirectedGraphLinkPntrHelper(
		factory(0), 5, 4,
		[][3]int{
			[3]int{0,1,0},
			[3]int{1,2,1},
			[3]int{2,3,2},
			[3]int{3,4,3},
		},
		t,
	)
	hashDirectedGraphLinkPntrHelper(
		factory(0), 5, 5,
		[][3]int{
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

func hashDirectedGraphOutEdgesHelper(
	container dynamicContainers.DirectedGraph[int, int],
	numVertices int,
	numEdges int,
	links [][3]int,
	outEdges map[int][]int,
	t *testing.T,
) {
	for i:=0; i<numVertices; i++ {
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

	for i:=0; i<numVertices; i++ {
		iterOutEdges,err:=container.OutEdges(i).Collect()
		test.Nil(err,t)
		test.Eq(len(outEdges[i]),len(iterOutEdges),t)

		for j:=0; j<len(outEdges[i]); j++ {
			found:=false
			for k:=0; k<len(iterOutEdges) && !found; k++ {
				found=(iterOutEdges[k]==outEdges[i][j])
			}
			test.True(found,t)
		}
	}

	err:=container.OutEdges(-1).Consume()
	test.ContainsError(customerr.InvalidValue,err,t)
	test.ContainsError(containerTypes.KeyError,err,t)
}
// Tests the OutEdges method functionality of a dynamic graph.
func DynHashDirectedGraphOutEdges(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
){
	hashDirectedGraphOutEdgesHelper(
		factory(0), 6, 4,
		[][3]int{
			[3]int{0,1,0},
			[3]int{1,2,1},
			[3]int{2,3,2},
			[3]int{3,4,3},
		},
		map[int][]int{
			0: []int{0},
			1: []int{1},
			2: []int{2},
			3: []int{3},
			4: []int{},
			5: []int{},
		},
		t,
	)
	hashDirectedGraphOutEdgesHelper(
		factory(0), 5, 5,
		[][3]int{
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

func hashDirectedGraphOutEdgePntrsHelper(
	container dynamicContainers.DirectedGraph[int, int],
	numVertices int,
	numEdges int,
	links [][3]int,
	outEdges map[int][]int,
	t *testing.T,
) {
	for i:=0; i<numVertices; i++ {
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

	for i:=0; i<numVertices; i++ {
		iterOutEdges,err:=container.OutEdgePntrs(&i).Collect()
		test.Nil(err,t)
		test.Eq(len(outEdges[i]),len(iterOutEdges),t)

		for j:=0; j<len(outEdges[i]); j++ {
			found:=false
			for k:=0; k<len(iterOutEdges) && !found; k++ {
				found=(*iterOutEdges[k]==outEdges[i][j])
			}
			test.True(found,t)
		}
	}

	err:=container.OutEdges(-1).Consume()
	test.ContainsError(customerr.InvalidValue,err,t)
	test.ContainsError(containerTypes.KeyError,err,t)
}
// Tests the OutEdgePntrs method functionality of a dynamic graph.
func DynHashDirectedGraphOutEdgePntrs(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
){
	container:=factory(0)
	if container.IsAddressable() {
		hashDirectedGraphOutEdgesHelper(
			factory(0), 6, 4,
			[][3]int{
				[3]int{0,1,0},
				[3]int{1,2,1},
				[3]int{2,3,2},
				[3]int{3,4,3},
			},
			map[int][]int{
				0: []int{0},
				1: []int{1},
				2: []int{2},
				3: []int{3},
				4: []int{},
				5: []int{},
			},
			t,
		)
		hashDirectedGraphOutEdgesHelper(
			factory(0), 5, 5,
			[][3]int{
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

func hashDirectedGraphOutVerticesHelper(
	container dynamicContainers.DirectedGraph[int, int],
	numVertices int,
	numEdges int,
	links [][3]int,
	outVertices map[int][]int,
	t *testing.T,
) {
	for i:=0; i<numVertices; i++ {
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

	for i:=0; i<numVertices; i++ {
		iterOutVertices,err:=container.OutVertices(i).Collect()
		test.Nil(err,t)
		test.Eq(len(outVertices[i]),len(iterOutVertices),t)

		for j:=0; j<len(outVertices[i]); j++ {
			found:=false
			for k:=0; k<len(iterOutVertices) && !found; k++ {
				found=(iterOutVertices[k]==outVertices[i][j])
			}
			test.True(found,t)
		}
	}

	err:=container.OutVertices(-1).Consume()
	test.ContainsError(customerr.InvalidValue,err,t)
	test.ContainsError(containerTypes.KeyError,err,t)
}
// Tests the OutVertices method functionality of a dynamic graph.
func DynHashDirectedGraphOutVertices(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
){
	hashDirectedGraphOutVerticesHelper(
		factory(0), 6, 4,
		[][3]int{
			[3]int{0,1,0},
			[3]int{1,2,1},
			[3]int{2,3,2},
			[3]int{3,4,3},
		},
		map[int][]int{
			0: []int{1},
			1: []int{2},
			2: []int{3},
			3: []int{4},
			4: []int{},
			5: []int{},
		},
		t,
	)
	hashDirectedGraphOutVerticesHelper(
		factory(0), 5, 5,
		[][3]int{
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

func hashDirectedGraphOutVerticePntrsHelper(
	container dynamicContainers.DirectedGraph[int, int],
	numVertices int,
	numEdges int,
	links [][3]int,
	outVertices map[int][]int,
	t *testing.T,
) {
	for i:=0; i<numVertices; i++ {
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

	for i:=0; i<numVertices; i++ {
		iterOutVertices,err:=container.OutVerticePntrs(&i).Collect()
		test.Nil(err,t)
		test.Eq(len(outVertices[i]),len(iterOutVertices),t)

		for j:=0; j<len(outVertices[i]); j++ {
			found:=false
			for k:=0; k<len(iterOutVertices) && !found; k++ {
				found=(*iterOutVertices[k]==outVertices[i][j])
			}
			test.True(found,t)
		}
	}

	tmp:=-1
	err:=container.OutVerticePntrs(&tmp).Consume()
	test.ContainsError(customerr.InvalidValue,err,t)
	test.ContainsError(containerTypes.KeyError,err,t)
}
// Tests the OutVerticePntrs method functionality of a dynamic graph.
func DynHashDirectedGraphOutVerticePntrs(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
){
	container:=factory(0)
	if container.IsAddressable() {
		hashDirectedGraphOutVerticePntrsHelper(
			factory(0), 6, 4,
			[][3]int{
				[3]int{0,1,0},
				[3]int{1,2,1},
				[3]int{2,3,2},
				[3]int{3,4,3},
			},
			map[int][]int{
				0: []int{1},
				1: []int{2},
				2: []int{3},
				3: []int{4},
				4: []int{},
				5: []int{},
			},
			t,
		)
		hashDirectedGraphOutVerticePntrsHelper(
			factory(0), 5, 5,
			[][3]int{
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

func hashDirectedGraphOutEdgesAndVerticesHelper(
	container dynamicContainers.DirectedGraph[int, int],
	numVertices int,
	numEdges int,
	links [][3]int,
	outEdges map[int][]int,
	outVertices map[int][]int,
	t *testing.T,
) {
	for i:=0; i<numVertices; i++ {
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

	for i:=0; i<numVertices; i++ {
		iterOutEdgesAndVertices,err:=container.OutEdgesAndVertices(i).Collect()
		test.Nil(err,t)
		test.Eq(len(outEdges[i]),len(iterOutEdgesAndVertices),t)

		for j:=0; j<len(outEdges[i]); j++ {
			found:=false
			for k:=0; k<len(iterOutEdgesAndVertices) && !found; k++ {
				found=(iterOutEdgesAndVertices[k].A==outEdges[i][j])
			}
			test.True(found,t)
		}
		for j:=0; j<len(outVertices[i]); j++ {
			found:=false
			for k:=0; k<len(iterOutEdgesAndVertices) && !found; k++ {
				found=(iterOutEdgesAndVertices[k].B==outVertices[i][j])
			}
			test.True(found,t)
		}
	}

	err:=container.OutVertices(-1).Consume()
	test.ContainsError(customerr.InvalidValue,err,t)
	test.ContainsError(containerTypes.KeyError,err,t)
}
// Tests the OutEdgesAndVertices method functionality of a dynamic graph.
func DynHashDirectedGraphOutEdgesAndVertices(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
){
	hashDirectedGraphOutEdgesAndVerticesHelper(
		factory(0), 6, 4,
		[][3]int{
			[3]int{0,1,0},
			[3]int{1,2,1},
			[3]int{2,3,2},
			[3]int{3,4,3},
		},
		map[int][]int{
			0: []int{0},
			1: []int{1},
			2: []int{2},
			3: []int{3},
			4: []int{},
			5: []int{},
		},
		map[int][]int{
			0: []int{1},
			1: []int{2},
			2: []int{3},
			3: []int{4},
			4: []int{},
			5: []int{},
		},
		t,
	)
	hashDirectedGraphOutEdgesAndVerticesHelper(
		factory(0), 5, 5,
		[][3]int{
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
			0: []int{0,1,2,3},
			1: []int{1,2,3},
			2: []int{2,3},
			3: []int{4},
			4: []int{},
			5: []int{},
		},
		map[int][]int{
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

func hashDirectedGraphOutEdgesAndVerticePntrsHelper(
	container dynamicContainers.DirectedGraph[int, int],
	numVertices int,
	numEdges int,
	links [][3]int,
	outEdges map[int][]int,
	outVertices map[int][]int,
	t *testing.T,
) {
	for i:=0; i<numVertices; i++ {
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

	for i:=0; i<numVertices; i++ {
		iterOutEdgesAndVertices,err:=container.
			OutEdgesAndVerticePntrs(&i).Collect()
		test.Nil(err,t)
		test.Eq(len(outEdges[i]),len(iterOutEdgesAndVertices),t)

		for j:=0; j<len(outEdges[i]); j++ {
			found:=false
			for k:=0; k<len(iterOutEdgesAndVertices) && !found; k++ {
				found=(*iterOutEdgesAndVertices[k].A==outEdges[i][j])
			}
			test.True(found,t)
		}
		for j:=0; j<len(outVertices[i]); j++ {
			found:=false
			for k:=0; k<len(iterOutEdgesAndVertices) && !found; k++ {
				found=(*iterOutEdgesAndVertices[k].B==outVertices[i][j])
			}
			test.True(found,t)
		}
	}

	err:=container.OutVertices(-1).Consume()
	test.ContainsError(customerr.InvalidValue,err,t)
	test.ContainsError(containerTypes.KeyError,err,t)
}
// Tests the OutEdgesAndVerticePntrs method functionality of a dynamic graph.
func DynHashDirectedGraphOutEdgesAndVerticePntrs(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
){
	container:=factory(0)
	if container.IsAddressable() {
		hashDirectedGraphOutEdgesAndVerticePntrsHelper(
			factory(0), 6, 4,
			[][3]int{
				[3]int{0,1,0},
				[3]int{1,2,1},
				[3]int{2,3,2},
				[3]int{3,4,3},
			},
			map[int][]int{
				0: []int{0},
				1: []int{1},
				2: []int{2},
				3: []int{3},
				4: []int{},
				5: []int{},
			},
			map[int][]int{
				0: []int{1},
				1: []int{2},
				2: []int{3},
				3: []int{4},
				4: []int{},
				5: []int{},
			},
			t,
		)
		hashDirectedGraphOutEdgesAndVerticePntrsHelper(
			factory(0), 5, 5,
			[][3]int{
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
				0: []int{0,1,2,3},
				1: []int{1,2,3},
				2: []int{2,3},
				3: []int{4},
				4: []int{},
				5: []int{},
			},
			map[int][]int{
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
