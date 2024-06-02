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
func graphWriteInterface[T any, U any](c dynamicContainers.WriteGraph[T, U]) {}
func graphInterface[T any, U any](c dynamicContainers.Graph[T, U])           {}

// Tests that the value supplied by the factory implements the
// [containerTypes.RWSyncable] interface.
func DynGraphInterfaceSyncableInterface[V any, E any](
	factory func(capacity int) dynamicContainers.Graph[V, E],
	t *testing.T,
) {
	var container containerTypes.RWSyncable = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.RWSyncable] interface.
func DynGraphInterfaceAddressableInterface[V any, E any](
	factory func(capacity int) dynamicContainers.Graph[V, E],
	t *testing.T,
) {
	var container containerTypes.Addressable = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Clear] interface.
func DynGraphInterfaceClearInterface[V any, E any](
	factory func(capacity int) dynamicContainers.Graph[V, E],
	t *testing.T,
) {
	var container containerTypes.Clear = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.GraphRead] interface.
func ReadDynGraphInterface[V any, E any](
	factory func(capacity int) dynamicContainers.Graph[V, E],
	t *testing.T,
) {
	graphReadInterface[V, E](factory(0))
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.WriteGraph] interface.
func WriteDynGraphInterface[V any, E any](
	factory func(capacity int) dynamicContainers.Graph[V, E],
	t *testing.T,
) {
	graphWriteInterface[V, E](factory(0))
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.Graph] interface.
func DynGraphInterfaceInterface[V any, E any](
	factory func(capacity int) dynamicContainers.Graph[V, E],
	t *testing.T,
) {
	graphInterface[V, E](factory(0))
}

// Tests that the value supplied by the factory does not implement the
// [staticContainers.Map] interface.
func DynGraphInterfaceStaticCapacityInterface[V any, E any](
	factory func(capacity int) dynamicContainers.Graph[V, E],
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
	g dynamicContainers.Graph[int,int],
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
func DynGraphInterfaceContainsEdge(
	factory func(capacity int) dynamicContainers.Graph[int, int],
	t *testing.T,
) {
	graphContainsEdgeHelper(factory(0), 0, t)
	graphContainsEdgeHelper(factory(0), 1, t)
	graphContainsEdgeHelper(factory(0), 2, t)
	graphContainsEdgeHelper(factory(0), 5, t)
}

func graphContainsEdgePntrHelper(
	g dynamicContainers.Graph[int,int],
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
func DynGraphInterfaceContainsEdgePntr(
	factory func(capacity int) dynamicContainers.Graph[int, int],
	t *testing.T,
) {
	graphContainsEdgePntrHelper(factory(0), 0, t)
	graphContainsEdgePntrHelper(factory(0), 1, t)
	graphContainsEdgePntrHelper(factory(0), 2, t)
	graphContainsEdgePntrHelper(factory(0), 5, t)
}

func graphContainsVertexHelper(
	g dynamicContainers.Graph[int,int],
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
func DynGraphInterfaceContainsVertex(
	factory func(capacity int) dynamicContainers.Graph[int, int],
	t *testing.T,
) {
	graphContainsVertexHelper(factory(0), 0, t)
	graphContainsVertexHelper(factory(0), 1, t)
	graphContainsVertexHelper(factory(0), 2, t)
	graphContainsVertexHelper(factory(0), 5, t)
}

func graphContainsVertexPntrHelper(
	g dynamicContainers.Graph[int,int],
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
func DynGraphInterfaceContainsVertexPntr(
	factory func(capacity int) dynamicContainers.Graph[int, int],
	t *testing.T,
) {
	graphContainsVertexPntrHelper(factory(0), 0, t)
	graphContainsVertexPntrHelper(factory(0), 1, t)
	graphContainsVertexPntrHelper(factory(0), 2, t)
	graphContainsVertexPntrHelper(factory(0), 5, t)
}

func testHashGraphEdgesHelper(
	factory func(capacity int) dynamicContainers.Graph[int, int],
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
func DynGraphInterfaceEdges(
	factory func(capacity int) dynamicContainers.Graph[int, int],
	t *testing.T,
){
	testHashGraphEdgesHelper(factory, 0, t)
	testHashGraphEdgesHelper(factory, 1, t)
	testHashGraphEdgesHelper(factory, 2, t)
	testHashGraphEdgesHelper(factory, 5, t)
}

func testHashGraphEdgePntrsHelper(
	factory func(capacity int) dynamicContainers.Graph[int, int],
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
	factory func(capacity int) dynamicContainers.Graph[int, int],
	t *testing.T,
) {
	container := factory(0)
	if container.IsAddressable() {
		testHashGraphEdgePntrsHelper(factory, 0, t)
		testHashGraphEdgePntrsHelper(factory, 1, t)
		testHashGraphEdgePntrsHelper(factory, 2, t)
		testHashGraphEdgePntrsHelper(factory, 5, t)
	} else {
		test.Panics(func() { container.EdgePntrs() }, t)
	}
}

func testHashGraphVerticesHelper(
	factory func(capacity int) dynamicContainers.Graph[int, int],
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
func DynGraphInterfaceVertices(
	factory func(capacity int) dynamicContainers.Graph[int, int],
	t *testing.T,
){
	testHashGraphVerticesHelper(factory, 0, t)
	testHashGraphVerticesHelper(factory, 1, t)
	testHashGraphVerticesHelper(factory, 2, t)
	testHashGraphVerticesHelper(factory, 5, t)
}

func testHashGraphVerticePntrsHelper(
	factory func(capacity int) dynamicContainers.Graph[int, int],
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
func DynHashGraphInterfaceVerticePntrs(
	factory func(capacity int) dynamicContainers.Graph[int, int],
	t *testing.T,
) {
	container := factory(0)
	if container.IsAddressable() {
		testHashGraphVerticePntrsHelper(factory, 0, t)
		testHashGraphVerticePntrsHelper(factory, 1, t)
		testHashGraphVerticePntrsHelper(factory, 2, t)
		testHashGraphVerticePntrsHelper(factory, 5, t)
	} else {
		test.Panics(func() { container.VerticePntrs() }, t)
	}
}

func hashGraphLinkHelper(
	container dynamicContainers.Graph[int, int],
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
func DynHashGraphLink(
	factory func(capacity int) dynamicContainers.Graph[int, int],
	t *testing.T,
){
	hashGraphLinkHelper(
		factory(0), 5, 4,
		[][3]int{
			[3]int{0,1,0},
			[3]int{1,2,1},
			[3]int{2,3,2},
			[3]int{3,4,3},
		},
		t,
	)
	hashGraphLinkHelper(
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

func hashGraphLinkPntrHelper(
	container dynamicContainers.Graph[int, int],
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
func DynHashGraphLinkPntr(
	factory func(capacity int) dynamicContainers.Graph[int, int],
	t *testing.T,
){
	hashGraphLinkPntrHelper(
		factory(0), 5, 4,
		[][3]int{
			[3]int{0,1,0},
			[3]int{1,2,1},
			[3]int{2,3,2},
			[3]int{3,4,3},
		},
		t,
	)
	hashGraphLinkPntrHelper(
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

func hashGraphOutEdgesHelper(
	container dynamicContainers.Graph[int, int],
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
func DynHashGraphOutEdges(
	factory func(capacity int) dynamicContainers.Graph[int, int],
	t *testing.T,
){
	hashGraphOutEdgesHelper(
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
	hashGraphOutEdgesHelper(
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

func hashGraphOutEdgePntrsHelper(
	container dynamicContainers.Graph[int, int],
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
func DynHashGraphOutEdgePntrs(
	factory func(capacity int) dynamicContainers.Graph[int, int],
	t *testing.T,
){
	container:=factory(0)
	if container.IsAddressable() {
		hashGraphOutEdgesHelper(
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
		hashGraphOutEdgesHelper(
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

func hashGraphOutVerticesHelper(
	container dynamicContainers.Graph[int, int],
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
func DynHashGraphOutVertices(
	factory func(capacity int) dynamicContainers.Graph[int, int],
	t *testing.T,
){
	hashGraphOutVerticesHelper(
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
	hashGraphOutVerticesHelper(
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

func hashGraphOutVerticePntrsHelper(
	container dynamicContainers.Graph[int, int],
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
func DynHashGraphOutVerticePntrs(
	factory func(capacity int) dynamicContainers.Graph[int, int],
	t *testing.T,
){
	container:=factory(0)
	if container.IsAddressable() {
		hashGraphOutVerticePntrsHelper(
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
		hashGraphOutVerticePntrsHelper(
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

func hashGraphOutEdgesAndVerticesHelper(
	container dynamicContainers.Graph[int, int],
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
func DynHashGraphOutEdgesAndVertices(
	factory func(capacity int) dynamicContainers.Graph[int, int],
	t *testing.T,
){
	hashGraphOutEdgesAndVerticesHelper(
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
	hashGraphOutEdgesAndVerticesHelper(
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

func hashGraphOutEdgesAndVerticePntrsHelper(
	container dynamicContainers.Graph[int, int],
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
func DynHashGraphOutEdgesAndVerticePntrs(
	factory func(capacity int) dynamicContainers.Graph[int, int],
	t *testing.T,
){
	container:=factory(0)
	if container.IsAddressable() {
		hashGraphOutEdgesAndVerticePntrsHelper(
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
		hashGraphOutEdgesAndVerticePntrsHelper(
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
