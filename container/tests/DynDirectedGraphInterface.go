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

type graphConstruction struct {
	vertices iter.Iter[int]
	edges iter.Iter[int]
	links [][3]int
	numVertices int
	numEdges int
}

func (g *graphConstruction)makeGraph(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) dynamicContainers.DirectedGraph[int, int] {
	g.numEdges=0
	g.numVertices=0
	rv:=factory(0)
	g.edges.ForEach(func(index, val int) (iter.IteratorFeedback, error) {
		g.numEdges++
		test.Nil(rv.AddEdges(val), t)
		return iter.Continue, nil
	})
	g.vertices.ForEach(func(index, val int) (iter.IteratorFeedback, error) {
		g.numVertices++
		test.Nil(rv.AddVertices(val), t)
		return iter.Continue, nil
	})
	for i := 0; i < len(g.links); i++ {
		test.Nil(rv.Link( g.links[i][0], g.links[i][1], g.links[i][2],), t)
	}
	for i := 0; i < len(g.links); i++ {
		test.True(
			rv.ContainsLink(
				g.links[i][0],
				g.links[i][1],
				g.links[i][2],
			),
			t,
		)
	}
	test.Eq(len(g.links), rv.NumLinks(), t)
	test.Eq(g.numVertices, rv.NumVertices(), t)
	test.Eq(g.numEdges, rv.NumEdges(), t)
	return rv
}

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
	g dynamicContainers.DirectedGraph[int, int],
	l int,
	t *testing.T,
) {
	for i := 0; i < l && g.AddEdges(i) == nil; i++ {
	}
	for i := 0; i < l; i++ {
		test.True(g.ContainsEdge(i), t)
	}
	test.False(g.ContainsEdge(-1), t)
	test.False(g.ContainsEdge(l), t)
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
	g dynamicContainers.DirectedGraph[int, int],
	l int,
	t *testing.T,
) {
	for i := 0; i < l && g.AddEdges(i) == nil; i++ {
	}
	for i := 0; i < l; i++ {
		test.True(g.ContainsEdgePntr(&i), t)
	}
	tmp := -1
	test.False(g.ContainsEdgePntr(&tmp), t)
	test.False(g.ContainsEdgePntr(&l), t)
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
	g dynamicContainers.DirectedGraph[int, int],
	l int,
	t *testing.T,
) {
	for i := 0; i < l && g.AddVertices(i) == nil; i++ {
	}
	for i := 0; i < l; i++ {
		test.True(g.ContainsVertex(i), t)
	}
	test.False(g.ContainsVertex(-1), t)
	test.False(g.ContainsVertex(l), t)
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
	g dynamicContainers.DirectedGraph[int, int],
	l int,
	t *testing.T,
) {
	for i := 0; i < l && g.AddVertices(i) == nil; i++ {
	}
	for i := 0; i < l; i++ {
		test.True(g.ContainsVertexPntr(&i), t)
	}
	tmp := -1
	test.False(g.ContainsVertexPntr(&tmp), t)
	test.False(g.ContainsVertexPntr(&l), t)
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

func directedGraphEdgesHelper(
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
		test.True(container.ContainsEdge(val), t)
		return iter.Continue, nil
	})
	test.Eq(l, cnt, t)
}

// Tests the Edges method functionality of a dynamic directed graph.
func DynDirectedGraphInterfaceEdges(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphEdgesHelper(factory, 0, t)
	directedGraphEdgesHelper(factory, 1, t)
	directedGraphEdgesHelper(factory, 2, t)
	directedGraphEdgesHelper(factory, 5, t)
}

func directedGraphEdgePntrsHelper(
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
		directedGraphEdgePntrsHelper(factory, 0, t)
		directedGraphEdgePntrsHelper(factory, 1, t)
		directedGraphEdgePntrsHelper(factory, 2, t)
		directedGraphEdgePntrsHelper(factory, 5, t)
	} else {
		test.Panics(func() { container.EdgePntrs() }, t)
	}
}

func directedGraphVerticesHelper(
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
			test.True(container.ContainsVertex(val), t)
			return iter.Continue, nil
		},
	)
	test.Eq(l, cnt, t)
}

// Tests the Vertices method functionality of a dynamic directed graph.
func DynDirectedGraphInterfaceVertices(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphVerticesHelper(factory, 0, t)
	directedGraphVerticesHelper(factory, 1, t)
	directedGraphVerticesHelper(factory, 2, t)
	directedGraphVerticesHelper(factory, 5, t)
}

func directedGraphVerticePntrsHelper(
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
		directedGraphVerticePntrsHelper(factory, 0, t)
		directedGraphVerticePntrsHelper(factory, 1, t)
		directedGraphVerticePntrsHelper(factory, 2, t)
		directedGraphVerticePntrsHelper(factory, 5, t)
	} else {
		test.Panics(func() { container.VerticePntrs() }, t)
	}
}

func directedGraphLinkHelper(
	factory func (capacity int) dynamicContainers.DirectedGraph[int, int],
	construction graphConstruction,
	t *testing.T,
) {
	container:=construction.makeGraph(factory, t)

	for i := 0; i < len(construction.links); i++ {
		test.Nil(
			container.Link(
				construction.links[i][0],
				construction.links[i][1],
				construction.links[i][2],
			),
			t,
		)
	}
	test.Eq(len(construction.links), container.NumLinks(), t)
	test.Eq(construction.numVertices, container.NumVertices(), t)
	test.Eq(construction.numEdges, container.NumEdges(), t)

	err := container.Link(-1, 0, 0)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)
	err = container.Link(0, -1, 0)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)
	err = container.Link(0, 0, -1)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)
}

// Tests the Link method functionality of a dynamic directed graph.
func DynDirectedGraphLink(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphLinkHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,5,1),
			edges: iter.Range[int](0,4,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 1},
				[3]int{2, 3, 2},
				[3]int{3, 4, 3},
			},
		},
		t,
	)
	directedGraphLinkHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,5,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 2, 1},
				[3]int{0, 3, 2},
				[3]int{0, 4, 3},
				[3]int{1, 2, 1},
				[3]int{1, 3, 2},
				[3]int{1, 4, 3},
				[3]int{2, 3, 2},
				[3]int{2, 4, 3},
				[3]int{3, 4, 4},
			},
		},
		t,
	)
}

func directedGraphLinkPntrHelper(
	factory func (capacity int) dynamicContainers.DirectedGraph[int, int],
	construction graphConstruction,
	t *testing.T,
) {
	container:=construction.makeGraph(factory,t)

	for i := 0; i < len(construction.links); i++ {
		test.Nil(
			container.LinkPntr(
				&construction.links[i][0],
				&construction.links[i][1],
				&construction.links[i][2],
			),
			t,
		)
	}
	test.Eq(len(construction.links), container.NumLinks(), t)
	test.Eq(construction.numVertices, container.NumVertices(), t)
	test.Eq(construction.numEdges, container.NumEdges(), t)

	tmp1 := -1
	tmp2 := 0
	err := container.LinkPntr(&tmp1, &tmp2, &tmp2)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)
	err = container.LinkPntr(&tmp2, &tmp1, &tmp2)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)
	err = container.LinkPntr(&tmp2, &tmp2, &tmp1)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)
}

// Tests the LinkPntr method functionality of a dynamic directed graph.
func DynDirectedGraphLinkPntr(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphLinkPntrHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,5,1),
			edges: iter.Range[int](0,4,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 1},
				[3]int{2, 3, 2},
				[3]int{3, 4, 3},
			},
		},
		t,
	)
	directedGraphLinkPntrHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,5,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 2, 1},
				[3]int{0, 3, 2},
				[3]int{0, 4, 3},
				[3]int{1, 2, 1},
				[3]int{1, 3, 2},
				[3]int{1, 4, 3},
				[3]int{2, 3, 2},
				[3]int{2, 4, 3},
				[3]int{3, 4, 4},
			},
		},
		t,
	)
}

func directedGraphNumOutLinksHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	construction graphConstruction,
	outEdges map[int][]int,
	t *testing.T,
) {
	container:=construction.makeGraph(factory,t)
	for i := 0; i < construction.numVertices; i++ {
		test.Eq(len(outEdges[i]), container.NumOutLinks(i), t)
	}

	test.Eq(0, container.NumOutLinks(-1), t)
}
// Tests the NumOutLinks method functionality of a dynamic directed graph.
func DynDirectedGraphNumOutLinks(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphNumOutLinksHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,6,1),
			edges: iter.Range[int](0,4,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 1},
				[3]int{2, 3, 2},
				[3]int{3, 4, 3},
			},
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
	directedGraphNumOutLinksHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,5,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 2, 1},
				[3]int{0, 3, 2},
				[3]int{0, 4, 3},
				[3]int{1, 2, 1},
				[3]int{1, 3, 2},
				[3]int{1, 4, 3},
				[3]int{2, 3, 2},
				[3]int{2, 4, 3},
				[3]int{3, 4, 4},
			},
		},
		map[int][]int{
			// from, e
			0: []int{0, 1, 2, 3},
			1: []int{1, 2, 3},
			2: []int{2, 3},
			3: []int{4},
			4: []int{},
			5: []int{},
		},
		t,
	)
}

func directedGraphNumOutLinksPntrHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	construction graphConstruction,
	outEdges map[int][]int,
	t *testing.T,
) {
	container:=construction.makeGraph(factory, t)
	for i := 0; i < construction.numVertices; i++ {
		test.Eq(len(outEdges[i]), container.NumOutLinksPntr(&i), t)
	}

	tmp := -1
	test.Eq(0, container.NumOutLinksPntr(&tmp), t)
}
// Tests the NumOutLinksPntr method functionality of a dynamic directed graph.
func DynDirectedGraphNumOutLinksPntr(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphNumOutLinksPntrHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,6,1),
			edges: iter.Range[int](0,4,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 1},
				[3]int{2, 3, 2},
				[3]int{3, 4, 3},
			},
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
	directedGraphNumOutLinksPntrHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,5,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 2, 1},
				[3]int{0, 3, 2},
				[3]int{0, 4, 3},
				[3]int{1, 2, 1},
				[3]int{1, 3, 2},
				[3]int{1, 4, 3},
				[3]int{2, 3, 2},
				[3]int{2, 4, 3},
				[3]int{3, 4, 4},
			},
		},
		map[int][]int{
			// from, e
			0: []int{0, 1, 2, 3},
			1: []int{1, 2, 3},
			2: []int{2, 3},
			3: []int{4},
			4: []int{},
			5: []int{},
		},
		t,
	)
}

func directedGraphOutEdgesHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	construction graphConstruction,
	outEdges map[int][]int,
	t *testing.T,
) {
	container:=construction.makeGraph(factory, t)
	for i := 0; i < construction.numVertices; i++ {
		mappedValues, err := container.OutEdges(i).Collect()
		test.Eq(len(outEdges[i]), len(mappedValues), t)
		test.Nil(err, t)
		test.SlicesMatchUnordered[int](outEdges[i], mappedValues, t)
	}

	err := container.OutEdges(-1).Consume()
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)
}
// Tests the OutEdges method functionality of a dynamic directed graph.
func DynDirectedGraphOutEdges(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphOutEdgesHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,6,1),
			edges: iter.Range[int](0,4,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 1},
				[3]int{2, 3, 2},
				[3]int{3, 4, 3},
			},
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
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,5,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 2, 1},
				[3]int{0, 3, 2},
				[3]int{0, 4, 3},
				[3]int{1, 2, 1},
				[3]int{1, 3, 2},
				[3]int{1, 4, 3},
				[3]int{2, 3, 2},
				[3]int{2, 4, 3},
				[3]int{3, 4, 4},
			},
		},
		map[int][]int{
			// from, e
			0: []int{0, 1, 2, 3},
			1: []int{1, 2, 3},
			2: []int{2, 3},
			3: []int{4},
			4: []int{},
			5: []int{},
		},
		t,
	)
}

func directedGraphOutEdgePntrsHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	construction graphConstruction,
	outEdges map[int][]int,
	t *testing.T,
) {
	container:=construction.makeGraph(factory, t)
	for i := 0; i < construction.numVertices; i++ {
		mappedValues, err := iter.PntrToVal[int](
			container.OutEdgePntrs(&i),
		).Collect()
		test.Eq(len(outEdges[i]), len(mappedValues), t)
		test.Nil(err, t)
		test.SlicesMatchUnordered[int](outEdges[i], mappedValues, t)
	}

	err := container.OutEdges(-1).Consume()
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)
}
// Tests the OutEdgePntrs method functionality of a dynamic directed graph.
func DynDirectedGraphOutEdgePntrs(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	container := factory(0)
	if container.IsAddressable() {
		directedGraphOutEdgesHelper(
			factory,
			graphConstruction{
				vertices: iter.Range[int](0,6,1),
				edges: iter.Range[int](0,4,1),
				links: [][3]int{
					// from, to, e
					[3]int{0, 1, 0},
					[3]int{1, 2, 1},
					[3]int{2, 3, 2},
					[3]int{3, 4, 3},
				},
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
			factory,
			graphConstruction{
				vertices: iter.Range[int](0,5,1),
				edges: iter.Range[int](0,5,1),
				links: [][3]int{
					// from, to, e
					[3]int{0, 1, 0},
					[3]int{0, 2, 1},
					[3]int{0, 3, 2},
					[3]int{0, 4, 3},
					[3]int{1, 2, 1},
					[3]int{1, 3, 2},
					[3]int{1, 4, 3},
					[3]int{2, 3, 2},
					[3]int{2, 4, 3},
					[3]int{3, 4, 4},
				},
			},
			map[int][]int{
				// from, e
				0: []int{0, 1, 2, 3},
				1: []int{1, 2, 3},
				2: []int{2, 3},
				3: []int{4},
				4: []int{},
				5: []int{},
			},
			t,
		)
	} else {
		tmp := 0
		test.Panics(func() { container.OutEdgePntrs(&tmp) }, t)
	}
}

func directedGraphOutVerticesHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	construction graphConstruction,
	outVertices map[int][]int,
	t *testing.T,
) {
	container:=construction.makeGraph(factory, t)
	for i := 0; i < construction.numVertices; i++ {
		mappedValues, err := container.OutVertices(i).Collect()
		test.Eq(len(outVertices[i]), len(mappedValues), t)
		test.Nil(err, t)
		test.SlicesMatchUnordered[int](outVertices[i], mappedValues, t)
	}

	err := container.OutVertices(-1).Consume()
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)
}
// Tests the OutVertices method functionality of a dynamic directed graph.
func DynDirectedGraphOutVertices(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphOutVerticesHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,6,1),
			edges: iter.Range[int](0,4,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 1},
				[3]int{2, 3, 2},
				[3]int{3, 4, 3},
			},
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
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,5,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 2, 1},
				[3]int{0, 3, 2},
				[3]int{0, 4, 3},
				[3]int{1, 2, 1},
				[3]int{1, 3, 2},
				[3]int{1, 4, 3},
				[3]int{2, 3, 2},
				[3]int{2, 4, 3},
				[3]int{3, 4, 4},
			},
		},
		map[int][]int{
			// from, to
			0: []int{1, 2, 3, 4},
			1: []int{2, 3, 4},
			2: []int{3, 4},
			3: []int{4},
			4: []int{},
			5: []int{},
		},
		t,
	)
}

func directedGraphOutVerticePntrsHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	construction graphConstruction,
	outVertices map[int][]int,
	t *testing.T,
) {
	container:=construction.makeGraph(factory, t)
	for i := 0; i < construction.numVertices; i++ {
		mappedValues, err := iter.PntrToVal[int](
			container.OutVerticePntrs(&i),
		).Collect()
		test.Eq(len(outVertices[i]), len(mappedValues), t)
		test.Nil(err, t)
		test.SlicesMatchUnordered[int](outVertices[i], mappedValues, t)
	}

	tmp := -1
	err := container.OutVerticePntrs(&tmp).Consume()
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)
}
// Tests the OutVerticePntrs method functionality of a dynamic directed graph.
func DynDirectedGraphOutVerticePntrs(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	container := factory(0)
	if container.IsAddressable() {
		directedGraphOutVerticePntrsHelper(
			factory,
			graphConstruction{
				vertices: iter.Range[int](0,6,1),
				edges: iter.Range[int](0,4,1),
				links: [][3]int{
					// from, to, e
					[3]int{0, 1, 0},
					[3]int{1, 2, 1},
					[3]int{2, 3, 2},
					[3]int{3, 4, 3},
				},
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
			factory,
			graphConstruction{
				vertices: iter.Range[int](0,5,1),
				edges: iter.Range[int](0,5,1),
				links: [][3]int{
					// from, to, e
					[3]int{0, 1, 0},
					[3]int{0, 2, 1},
					[3]int{0, 3, 2},
					[3]int{0, 4, 3},
					[3]int{1, 2, 1},
					[3]int{1, 3, 2},
					[3]int{1, 4, 3},
					[3]int{2, 3, 2},
					[3]int{2, 4, 3},
					[3]int{3, 4, 4},
				},
			},
			map[int][]int{
				// from, to
				0: []int{1, 2, 3, 4},
				1: []int{2, 3, 4},
				2: []int{3, 4},
				3: []int{4},
				4: []int{},
				5: []int{},
			},
			t,
		)
	} else {
		tmp := 0
		test.Panics(func() { container.OutEdgePntrs(&tmp) }, t)
	}
}

func directedGraphOutEdgesAndVerticesHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	construction graphConstruction,
	outEdges map[int][]int,
	outVertices map[int][]int,
	t *testing.T,
) {
	container:=construction.makeGraph(factory, t)
	for i := 0; i < construction.numVertices; i++ {
		iterOutEdgesAndVertices, err := container.OutEdgesAndVertices(i).Collect()
		test.Nil(err, t)
		test.Eq(len(outEdges[i]), len(iterOutEdgesAndVertices), t)

		iterOutEdges, err := iter.Map[basic.Pair[int, int], int](
			iter.SliceElems[basic.Pair[int, int]](iterOutEdgesAndVertices),
			func(index int, val basic.Pair[int, int]) (int, error) {
				return val.A, nil
			},
		).Collect()
		test.Nil(err, t)
		test.SlicesMatchUnordered[int](outEdges[i], iterOutEdges, t)

		iterOutVertices, err := iter.Map[basic.Pair[int, int], int](
			iter.SliceElems[basic.Pair[int, int]](iterOutEdgesAndVertices),
			func(index int, val basic.Pair[int, int]) (int, error) {
				return val.B, nil
			},
		).Collect()
		test.Nil(err, t)
		test.SlicesMatchUnordered[int](outVertices[i], iterOutVertices, t)
	}
	err := container.OutVertices(-1).Consume()
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)
}
// Tests the OutEdgesAndVertices method functionality of a dynamic directed graph.
func DynDirectedGraphOutEdgesAndVertices(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphOutEdgesAndVerticesHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,6,1),
			edges: iter.Range[int](0,4,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 1},
				[3]int{2, 3, 2},
				[3]int{3, 4, 3},
			},
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
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,5,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 2, 1},
				[3]int{0, 3, 2},
				[3]int{0, 4, 3},
				[3]int{1, 2, 1},
				[3]int{1, 3, 2},
				[3]int{1, 4, 3},
				[3]int{2, 3, 2},
				[3]int{2, 4, 3},
				[3]int{3, 4, 4},
			},
		},
		map[int][]int{
			// from, e
			0: []int{0, 1, 2, 3},
			1: []int{1, 2, 3},
			2: []int{2, 3},
			3: []int{4},
			4: []int{},
			5: []int{},
		},
		map[int][]int{
			// from, to
			0: []int{1, 2, 3, 4},
			1: []int{2, 3, 4},
			2: []int{3, 4},
			3: []int{4},
			4: []int{},
			5: []int{},
		},
		t,
	)
}

func directedGraphOutEdgesAndVerticePntrsHelper(
	factory func(capacuty int) dynamicContainers.DirectedGraph[int, int],
	construction graphConstruction,
	outEdges map[int][]int,
	outVertices map[int][]int,
	t *testing.T,
) {
	container:=construction.makeGraph(factory, t)
	for i := 0; i < construction.numVertices; i++ {
		iterOutEdgesAndVertices, err := container.
			OutEdgesAndVerticePntrs(&i).Collect()
		test.Nil(err, t)
		test.Eq(len(outEdges[i]), len(iterOutEdgesAndVertices), t)

		iterOutEdges, err := iter.Map[basic.Pair[*int, *int], int](
			iter.SliceElems[basic.Pair[*int, *int]](iterOutEdgesAndVertices),
			func(index int, val basic.Pair[*int, *int]) (int, error) {
				return *val.A, nil
			},
		).Collect()
		test.Nil(err, t)
		test.SlicesMatchUnordered[int](outEdges[i], iterOutEdges, t)

		iterOutVertices, err := iter.Map[basic.Pair[*int, *int], int](
			iter.SliceElems[basic.Pair[*int, *int]](iterOutEdgesAndVertices),
			func(index int, val basic.Pair[*int, *int]) (int, error) {
				return *val.B, nil
			},
		).Collect()
		test.Nil(err, t)
		test.SlicesMatchUnordered[int](outVertices[i], iterOutVertices, t)
	}

	err := container.OutVertices(-1).Consume()
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)
}
// Tests the OutEdgesAndVerticePntrs method functionality of a dynamic directed graph.
func DynDirectedGraphOutEdgesAndVerticePntrs(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	container := factory(0)
	if container.IsAddressable() {
		directedGraphOutEdgesAndVerticePntrsHelper(
			factory,
			graphConstruction{
				vertices: iter.Range[int](0,6,1),
				edges: iter.Range[int](0,4,1),
				links: [][3]int{
					// from, to, e
					[3]int{0, 1, 0},
					[3]int{1, 2, 1},
					[3]int{2, 3, 2},
					[3]int{3, 4, 3},
				},
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
			factory,
			graphConstruction{
				vertices: iter.Range[int](0,5,1),
				edges: iter.Range[int](0,5,1),
				links: [][3]int{
					// from, to, e
					[3]int{0, 1, 0},
					[3]int{0, 2, 1},
					[3]int{0, 3, 2},
					[3]int{0, 4, 3},
					[3]int{1, 2, 1},
					[3]int{1, 3, 2},
					[3]int{1, 4, 3},
					[3]int{2, 3, 2},
					[3]int{2, 4, 3},
					[3]int{3, 4, 4},
				},
			},
			map[int][]int{
				// from, e
				0: []int{0, 1, 2, 3},
				1: []int{1, 2, 3},
				2: []int{2, 3},
				3: []int{4},
				4: []int{},
				5: []int{},
			},
			map[int][]int{
				// from, to
				0: []int{1, 2, 3, 4},
				1: []int{2, 3, 4},
				2: []int{3, 4},
				3: []int{4},
				4: []int{},
				5: []int{},
			},
			t,
		)
	} else {
		tmp := -1
		test.Panics(func() { container.OutEdgesAndVerticePntrs(&tmp) }, t)
	}
}

func directedGraphEdgesBetweenHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	construction graphConstruction,
	structure map[int]map[int][]int,
	t *testing.T,
) {
	container:=construction.makeGraph(factory,t)
	for from, allTo := range structure {
		for to, edges := range allTo {
			res, err := container.EdgesBetween(from, to).Collect()
			test.Nil(err, t)
			test.SlicesMatchUnordered[int](edges, res, t)
		}
	}

	res, err := container.EdgesBetween(0, -1).Collect()
	test.SlicesMatchUnordered[int]([]int{}, res, t)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)

	res, err = container.EdgesBetween(-1, 0).Collect()
	test.SlicesMatchUnordered[int]([]int{}, res, t)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)

	res, err = container.EdgesBetween(-1, -1).Collect()
	test.SlicesMatchUnordered[int]([]int{}, res, t)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)
}
// Tests the EdgesBetween method functionality of a dynamic directed graph.
func DynDirectedGraphEdgesBetween(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphEdgesBetweenHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 1, 1},
				[3]int{0, 1, 2},
				[3]int{0, 2, 3},
				[3]int{0, 2, 4},
				[3]int{1, 0, 0},
				[3]int{1, 2, 1},
			},
		},
		map[int]map[int][]int{
			// from: to: e
			0: map[int][]int{
				// to: e
				0: []int{},
				1: []int{0, 1, 2},
				2: []int{3, 4},
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
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	construction graphConstruction,
	structure map[int]map[int][]int,
	t *testing.T,
) {
	container:=construction.makeGraph(factory,t)
	for from, allTo := range structure {
		for to, edges := range allTo {
			res, err := iter.PntrToVal[int](
				container.EdgesBetweenPntr(&from, &to),
			).Collect()
			test.Nil(err, t)
			test.SlicesMatchUnordered[int](edges, res, t)
		}
	}

	tmp1 := 0
	tmp2 := -1
	res, err := iter.PntrToVal[int](
		container.EdgesBetweenPntr(&tmp1, &tmp2),
	).Collect()
	test.SlicesMatchUnordered[int]([]int{}, res, t)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)

	res, err = iter.PntrToVal[int](
		container.EdgesBetweenPntr(&tmp2, &tmp1),
	).Collect()
	test.SlicesMatchUnordered[int]([]int{}, res, t)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)

	res, err = iter.PntrToVal[int](
		container.EdgesBetweenPntr(&tmp2, &tmp2),
	).Collect()
	test.SlicesMatchUnordered[int]([]int{}, res, t)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)
}
// Tests the EdgesBetweenPntr method functionality of a dynamic directed graph.
func DynDirectedGraphEdgesBetweenPntr(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	container := factory(0)
	if container.IsAddressable() {
		directedGraphEdgesBetweenPntrHelper(
			factory,
			graphConstruction{
				vertices: iter.Range[int](0,3,1),
				edges: iter.Range[int](0,5,1),
				links: [][3]int{
					// from, to, e
					[3]int{0, 1, 0},
					[3]int{0, 1, 1},
					[3]int{0, 1, 2},
					[3]int{0, 2, 3},
					[3]int{0, 2, 4},
					[3]int{1, 0, 0},
					[3]int{1, 2, 1},
				},
			},
			map[int]map[int][]int{
				// from: to: e
				0: map[int][]int{
					// to: e
					0: []int{},
					1: []int{0, 1, 2},
					2: []int{3, 4},
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
		tmp := -1
		test.Panics(func() { container.EdgesBetweenPntr(&tmp, &tmp) }, t)
	}
}

func directedGraphDeleteLinkHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	construction graphConstruction,
	deleteLinks [][3]int,
	t *testing.T,
) {
	container:=construction.makeGraph(factory, t)
	for i := 0; i < len(deleteLinks); i++ {
		test.Nil(
			container.DeleteLink(
				deleteLinks[i][0],
				deleteLinks[i][1],
				deleteLinks[i][2],
			),
			t,
		)
	}
	for i := 0; i < len(deleteLinks); i++ {
		test.False(
			container.ContainsLink(
				deleteLinks[i][0],
				deleteLinks[i][1],
				deleteLinks[i][2],
			),
			t,
		)
	}
	test.Eq(construction.numEdges, container.NumEdges(), t)
	test.Eq(construction.numVertices, container.NumVertices(), t)
	test.Eq(len(construction.links)-len(deleteLinks), container.NumLinks(), t)

	err := container.DeleteLink(-1, 0, 0)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)

	err = container.DeleteLink(0, -1, 0)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)

	err = container.DeleteLink(0, 0, -1)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)
}
// Tests the DeleteLink method functionality of a dynamic directed graph.
func DynDirectedGraphDeleteLink(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphDeleteLinkHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 1, 1},
				[3]int{0, 1, 2},
				[3]int{0, 2, 3},
				[3]int{0, 2, 4},
				[3]int{1, 0, 0},
				[3]int{1, 2, 1},
			},
		},
		[][3]int{
			[3]int{0, 1, 0},
			[3]int{0, 1, 1},
			[3]int{1, 0, 0},
		},
		t,
	)
}

func directedGraphDeleteLinkPntrHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	construction graphConstruction,
	deleteLinks [][3]int,
	t *testing.T,
) {
	container:=construction.makeGraph(factory, t)
	for i := 0; i < len(deleteLinks); i++ {
		test.Nil(
			container.DeleteLinkPntr(
				&deleteLinks[i][0],
				&deleteLinks[i][1],
				&deleteLinks[i][2],
			),
			t,
		)
	}
	for i := 0; i < len(deleteLinks); i++ {
		test.False(
			container.ContainsLink(
				deleteLinks[i][0],
				deleteLinks[i][1],
				deleteLinks[i][2],
			),
			t,
		)
	}
	test.Eq(construction.numEdges, container.NumEdges(), t)
	test.Eq(construction.numVertices, container.NumVertices(), t)
	test.Eq(len(construction.links)-len(deleteLinks), container.NumLinks(), t)

	tmp1 := 0
	tmp2 := -1
	err := container.DeleteLinkPntr(&tmp2, &tmp1, &tmp1)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)

	err = container.DeleteLinkPntr(&tmp1, &tmp2, &tmp1)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)

	err = container.DeleteLinkPntr(&tmp1, &tmp1, &tmp2)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)
}
// Tests the DeleteLinkPntr method functionality of a dynamic directed graph.
func DynDirectedGraphDeleteLinkPntr(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphDeleteLinkPntrHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 1, 1},
				[3]int{0, 1, 2},
				[3]int{0, 2, 3},
				[3]int{0, 2, 4},
				[3]int{1, 0, 0},
				[3]int{1, 2, 1},
			},
		},
		[][3]int{
			[3]int{0, 1, 0},
			[3]int{0, 1, 1},
			[3]int{1, 0, 0},
		},
		t,
	)
}

func directedGraphDeleteLinksHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	construction graphConstruction,
	deleteLinks [2]int,
	numDeletedLinks int,
	t *testing.T,
) {
	container:=construction.makeGraph(factory, t)
	test.Nil(container.DeleteLinks(deleteLinks[0], deleteLinks[1]), t)
	for i := 0; i < construction.numEdges; i++ {
		test.False(
			container.ContainsLink(deleteLinks[0], deleteLinks[1], i),
			t,
		)
	}
	test.Eq(construction.numEdges, container.NumEdges(), t)
	test.Eq(construction.numVertices, container.NumVertices(), t)
	test.Eq(len(construction.links)-numDeletedLinks, container.NumLinks(), t)

	err := container.DeleteLinks(-1, 0)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)

	err = container.DeleteLinks(0, -1)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)
}
// Tests the DeleteLinks method functionality of a dynamic directed graph.
func DynDirectedGraphDeleteLinks(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphDeleteLinksHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 1, 1},
				[3]int{0, 1, 2},
				[3]int{0, 2, 3},
				[3]int{0, 2, 4},
				[3]int{1, 0, 0},
				[3]int{1, 2, 1},
			},
		},
		[2]int{0, 1},
		3,
		t,
	)
	directedGraphDeleteLinksHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 1, 1},
				[3]int{0, 1, 2},
				[3]int{0, 2, 3},
				[3]int{0, 2, 4},
				[3]int{1, 0, 0},
				[3]int{1, 2, 1},
			},
		},
		[2]int{1, 0},
		1,
		t,
	)
}

func directedGraphDeleteLinksPntrHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	construction graphConstruction,
	deleteLinks [2]int,
	numDeletedLinks int,
	t *testing.T,
) {
	container:=construction.makeGraph(factory, t)
	test.Nil(container.DeleteLinksPntr(&deleteLinks[0], &deleteLinks[1]), t)
	for i := 0; i < construction.numEdges; i++ {
		test.False(
			container.ContainsLink(deleteLinks[0], deleteLinks[1], i),
			t,
		)
	}
	test.Eq(construction.numEdges, container.NumEdges(), t)
	test.Eq(construction.numVertices, container.NumVertices(), t)
	test.Eq(len(construction.links)-numDeletedLinks, container.NumLinks(), t)

	err := container.DeleteLinks(-1, 0)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)

	err = container.DeleteLinks(0, -1)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)
}
// Tests the DeleteLinks method functionality of a dynamic directed graph.
func DynDirectedGraphDeleteLinksPntr(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphDeleteLinksPntrHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 1, 1},
				[3]int{0, 1, 2},
				[3]int{0, 2, 3},
				[3]int{0, 2, 4},
				[3]int{1, 0, 0},
				[3]int{1, 2, 1},
			},
		},
		[2]int{0, 1},
		3,
		t,
	)
	directedGraphDeleteLinksPntrHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 1, 1},
				[3]int{0, 1, 2},
				[3]int{0, 2, 3},
				[3]int{0, 2, 4},
				[3]int{1, 0, 0},
				[3]int{1, 2, 1},
			},
		},
		[2]int{1, 0},
		1,
		t,
	)
}

func directedGraphDeleteVertexHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	construction graphConstruction,
	deleteVertices []int,
	numDeletedLinks int,
	t *testing.T,
) {
	container:=construction.makeGraph(factory,t)
	for i := 0; i < len(deleteVertices); i++ {
		test.Nil(container.DeleteVertex(deleteVertices[i]), t)
	}
	for i := 0; i < len(deleteVertices); i++ {
		test.False(container.ContainsVertex(deleteVertices[i]), t)
	}
	test.Eq(construction.numEdges, container.NumEdges(), t)
	test.Eq(
		construction.numVertices-len(deleteVertices),
		container.NumVertices(),
		t,
	)
	test.Eq(len(construction.links)-numDeletedLinks, container.NumLinks(), t)

	err := container.DeleteVertex(-1)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)
}
// Tests the DeleteVertex method functionality of a dynamic directed graph.
func DynDirectedGraphDeleteVertex(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphDeleteVertexHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 1, 1},
				[3]int{0, 1, 2},
				[3]int{0, 2, 3},
				[3]int{0, 2, 4},
				[3]int{1, 0, 0},
				[3]int{1, 2, 1},
			},
		},
		[]int{0},
		6,
		t,
	)
	directedGraphDeleteVertexHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 1, 1},
				[3]int{0, 1, 2},
				[3]int{0, 2, 3},
				[3]int{0, 2, 4},
				[3]int{1, 0, 0},
				[3]int{1, 2, 1},
			},
		},
		[]int{1},
		5,
		t,
	)
	directedGraphDeleteVertexHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 1, 1},
				[3]int{0, 1, 2},
				[3]int{0, 2, 3},
				[3]int{0, 2, 4},
				[3]int{1, 0, 0},
				[3]int{1, 2, 1},
			},
		},
		[]int{1, 2},
		7,
		t,
	)
}

func directedGraphDeleteVertexPntrHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	construction graphConstruction,
	deleteVertices []int,
	numDeletedLinks int,
	t *testing.T,
) {
	container:=construction.makeGraph(factory, t)
	for i := 0; i < len(deleteVertices); i++ {
		test.Nil(container.DeleteVertexPntr(&deleteVertices[i]), t)
	}
	for i := 0; i < len(deleteVertices); i++ {
		test.False(container.ContainsVertex(deleteVertices[i]), t)
	}
	test.Eq(construction.numEdges, container.NumEdges(), t)
	test.Eq(
		construction.numVertices-len(deleteVertices),
		container.NumVertices(),
		t,
	)
	test.Eq(len(construction.links)-numDeletedLinks, container.NumLinks(), t)

	tmp := -1
	err := container.DeleteVertexPntr(&tmp)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)
}
// Tests the DeleteVertexPntr method functionality of a dynamic directed graph.
func DynDirectedGraphDeleteVertexPntr(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphDeleteVertexPntrHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 1, 1},
				[3]int{0, 1, 2},
				[3]int{0, 2, 3},
				[3]int{0, 2, 4},
				[3]int{1, 0, 0},
				[3]int{1, 2, 1},
			},
		},
		[]int{0},
		6,
		t,
	)
	directedGraphDeleteVertexPntrHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 1, 1},
				[3]int{0, 1, 2},
				[3]int{0, 2, 3},
				[3]int{0, 2, 4},
				[3]int{1, 0, 0},
				[3]int{1, 2, 1},
			},
		},
		[]int{1},
		5,
		t,
	)
	directedGraphDeleteVertexPntrHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 1, 1},
				[3]int{0, 1, 2},
				[3]int{0, 2, 3},
				[3]int{0, 2, 4},
				[3]int{1, 0, 0},
				[3]int{1, 2, 1},
			},
		},
		[]int{1, 2},
		7,
		t,
	)
}

func directedGraphDeleteEdgeHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	construction graphConstruction,
	deleteEdges []int,
	numDeletedLinks int,
	t *testing.T,
) {
	container:=construction.makeGraph(factory, t)
	for i := 0; i < len(deleteEdges); i++ {
		test.Nil(container.DeleteEdge(deleteEdges[i]), t)
	}
	for i := 0; i < len(deleteEdges); i++ {
		test.False(container.ContainsEdge(deleteEdges[i]), t)
	}
	test.Eq(construction.numEdges-len(deleteEdges), container.NumEdges(), t)
	test.Eq(construction.numVertices, container.NumVertices(), t)
	test.Eq(len(construction.links)-numDeletedLinks, container.NumLinks(), t)

	err := container.DeleteEdge(-1)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)
}
// Tests the DeleteEdge method functionality of a dynamic directed graph.
func DynDirectedGraphDeleteEdge(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphDeleteEdgeHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 1, 1},
				[3]int{0, 1, 2},
				[3]int{0, 2, 3},
				[3]int{0, 2, 4},
				[3]int{1, 0, 0},
				[3]int{1, 2, 1},
			},
		},
		[]int{0},
		2,
		t,
	)
	directedGraphDeleteEdgeHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 1, 1},
				[3]int{0, 1, 2},
				[3]int{0, 2, 3},
				[3]int{0, 2, 4},
				[3]int{1, 0, 0},
				[3]int{1, 2, 1},
			},
		},
		[]int{2},
		1,
		t,
	)
	directedGraphDeleteEdgeHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 1, 1},
				[3]int{0, 1, 2},
				[3]int{0, 2, 3},
				[3]int{0, 2, 4},
				[3]int{1, 0, 0},
				[3]int{1, 2, 1},
			},
		},
		[]int{1, 2},
		3,
		t,
	)
}

func directedGraphDeleteEdgePntrHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	construction graphConstruction,
	deleteEdges []int,
	numDeletedLinks int,
	t *testing.T,
) {
	container:=construction.makeGraph(factory, t)
	for i := 0; i < len(deleteEdges); i++ {
		test.Nil(container.DeleteEdgePntr(&deleteEdges[i]), t)
	}
	for i := 0; i < len(deleteEdges); i++ {
		test.False(container.ContainsEdge(deleteEdges[i]), t)
	}
	test.Eq(construction.numEdges-len(deleteEdges), container.NumEdges(), t)
	test.Eq(construction.numVertices, container.NumVertices(), t)
	test.Eq(len(construction.links)-numDeletedLinks, container.NumLinks(), t)

	tmp := -1
	err := container.DeleteEdgePntr(&tmp)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(containerTypes.KeyError, err, t)
}
// Tests the DeleteEdgePntr method functionality of a dynamic directed graph.
func DynDirectedGraphDeleteEdgePntr(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphDeleteEdgePntrHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 1, 1},
				[3]int{0, 1, 2},
				[3]int{0, 2, 3},
				[3]int{0, 2, 4},
				[3]int{1, 0, 0},
				[3]int{1, 2, 1},
			},
		},
		[]int{0},
		2,
		t,
	)
	directedGraphDeleteEdgePntrHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 1, 1},
				[3]int{0, 1, 2},
				[3]int{0, 2, 3},
				[3]int{0, 2, 4},
				[3]int{1, 0, 0},
				[3]int{1, 2, 1},
			},
		},
		[]int{2},
		1,
		t,
	)
	directedGraphDeleteEdgePntrHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 1, 1},
				[3]int{0, 1, 2},
				[3]int{0, 2, 3},
				[3]int{0, 2, 4},
				[3]int{1, 0, 0},
				[3]int{1, 2, 1},
			},
		},
		[]int{1, 2},
		3,
		t,
	)
}

// Tests the Clear method functionality of a dynamic directed graph.
func DynDirectedGraphClear(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	container:=(&graphConstruction{
		vertices: iter.Range[int](0,5,1),
		edges: iter.Range[int](0,5,1),
		links: [][3]int{
			// from, to, e
			[3]int{0, 1, 0},
			[3]int{0, 1, 1},
			[3]int{0, 1, 2},
			[3]int{0, 2, 3},
			[3]int{0, 2, 4},
			[3]int{1, 0, 0},
			[3]int{1, 2, 1},
		},
	}).makeGraph(factory, t)
	container.Clear()
	test.Eq(0, container.NumLinks(), t)
	test.Eq(0, container.NumVertices(), t)
	test.Eq(0, container.NumEdges(), t)

	container.Clear()
	test.Eq(0, container.NumLinks(), t)
	test.Eq(0, container.NumVertices(), t)
	test.Eq(0, container.NumEdges(), t)
}

func directedGraphKeyedEqHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	g1 graphConstruction,
	g2 graphConstruction,
	expResult bool,
	t *testing.T,
) {
	container1:=g1.makeGraph(factory, t)
	container2:=g2.makeGraph(factory, t)
	test.Eq(expResult, container1.KeyedEq(container2), t)
	test.Eq(expResult, container2.KeyedEq(container1), t)
}
// Tests the KeyedEq method functionality of a dynamic directed graph.
func DynDirectedGraphKeyedEq(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphKeyedEqHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,2,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](0,2,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 0, 0},
			},
		},
		true,
		t,
	)
	directedGraphKeyedEqHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 0, 0},
				[3]int{2, 1, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](0,2,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 0, 0},
			},
		},
		false,
		t,
	)
	directedGraphKeyedEqHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,2,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 0, 1},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](0,2,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 0, 0},
			},
		},
		false,
		t,
	)
	directedGraphKeyedEqHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,2,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 1, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](0,2,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 0, 0},
			},
		},
		false,
		t,
	)
	directedGraphKeyedEqHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,2,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 0, 0},
				[3]int{1, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](0,2,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 0, 0},
			},
		},
		false,
		t,
	)
	directedGraphKeyedEqHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,4,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 2, 1},
				[3]int{0, 3, 2},
				[3]int{1, 2, 0},
				[3]int{1, 3, 2},
				[3]int{1, 0, 4},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](0,4,1),
			edges: iter.Range[int](0,5,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{0, 2, 1},
				[3]int{0, 3, 2},
				[3]int{1, 2, 0},
				[3]int{1, 3, 2},
				[3]int{1, 0, 4},
			},
		},
		true,
		t,
	)
	directedGraphKeyedEqHelper(
		factory,
		graphConstruction{
			vertices: iter.SliceElems[int]([]int{0,1,5}),
			edges: iter.SliceElems[int]([]int{0,5}),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,2,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 0, 0},
			},
		},
		false,
		t,
	)
}

func directedGraphIntersectionHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	g1 graphConstruction,
	g2 graphConstruction,
	expG graphConstruction,
	t *testing.T,
) {
	container1:=g1.makeGraph(factory, t)
	container2:=g2.makeGraph(factory, t)
	container3:=expG.makeGraph(factory, t)
	container4:=factory(0)
	container4.Intersection(container1, container2)
	test.True(container3.KeyedEq(container4),t)

	container5:=factory(0)
	container5.Intersection(container2, container1)
	test.True(container3.KeyedEq(container5),t)
}
// Tests the Intersection method functionality of a dynamic directed graph.
func DynDirectedGraphIntersection(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphIntersectionHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,2,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](0,2,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.SliceElems[int]([]int{0,1}),
			edges: iter.SliceElems[int]([]int{0}),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 0, 0},
			},
		},
		t,
	)
	directedGraphIntersectionHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 0},
				[3]int{2, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](1,4,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{1, 2, 0},
				[3]int{2, 3, 0},
				[3]int{3, 1, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](1,3,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{1, 2, 0},
			},
		},
		t,
	)
	directedGraphIntersectionHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 0},
				[3]int{2, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](1,4,1),
			edges: iter.Range[int](0,2,1),
			links: [][3]int{
				// from, to, e
				[3]int{1, 2, 1},
				[3]int{2, 3, 0},
				[3]int{3, 1, 0},
			},
		},
		graphConstruction{
			vertices: iter.NoElem[int](),
			edges: iter.NoElem[int](),
			links: [][3]int{},
		},
		t,
	)
	directedGraphIntersectionHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 0},
				[3]int{2, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](1,5,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{2, 3, 0},
				[3]int{3, 4, 0},
				[3]int{4, 2, 0},
			},
		},
		graphConstruction{
			vertices: iter.NoElem[int](),
			edges: iter.NoElem[int](),
			links: [][3]int{},
		},
		t,
	)
}

func directedGraphUnionHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	g1 graphConstruction,
	g2 graphConstruction,
	expG graphConstruction,
	t *testing.T,
) {
	container1:=g1.makeGraph(factory, t)
	container2:=g2.makeGraph(factory, t)
	container3:=expG.makeGraph(factory, t)
	container4:=factory(0)
	container4.Union(container1, container2)
	test.True(container3.KeyedEq(container4),t)

	container5:=factory(0)
	container5.Union(container2, container1)
	test.True(container3.KeyedEq(container5),t)
}
// Tests the Union method functionality of a dynamic directed graph.
func DynDirectedGraphUnion(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphUnionHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,2,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](0,2,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.SliceElems[int]([]int{0,1}),
			edges: iter.SliceElems[int]([]int{0}),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 0, 0},
			},
		},
		t,
	)
	directedGraphUnionHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 0},
				[3]int{2, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](1,4,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{1, 2, 0},
				[3]int{2, 3, 0},
				[3]int{3, 1, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](0,4,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 0},
				[3]int{2, 0, 0},
				[3]int{2, 3, 0},
				[3]int{3, 1, 0},
			},
		},
		t,
	)
	directedGraphUnionHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 0},
				[3]int{2, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](1,4,1),
			edges: iter.Range[int](0,2,1),
			links: [][3]int{
				// from, to, e
				[3]int{1, 2, 1},
				[3]int{2, 3, 0},
				[3]int{3, 1, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](0,4,1),
			edges: iter.Range[int](0,2,1),
			links: [][3]int{
				[3]int{0, 1, 0},
				[3]int{1, 2, 0},
				[3]int{2, 0, 0},
				[3]int{1, 2, 1},
				[3]int{2, 3, 0},
				[3]int{3, 1, 0},
			},
		},
		t,
	)
	directedGraphUnionHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 0},
				[3]int{2, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](1,5,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{2, 3, 0},
				[3]int{3, 4, 0},
				[3]int{4, 2, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](0,5,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				[3]int{0, 1, 0},
				[3]int{1, 2, 0},
				[3]int{2, 0, 0},
				[3]int{2, 3, 0},
				[3]int{3, 4, 0},
				[3]int{4, 2, 0},
			},
		},
		t,
	)
}

func directedGraphDifferenceHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	g1 graphConstruction,
	g2 graphConstruction,
	expG graphConstruction,
	t *testing.T,
) {
	container1:=g1.makeGraph(factory, t)
	container2:=g2.makeGraph(factory, t)
	container3:=expG.makeGraph(factory, t)
	container4:=factory(0)
	container4.Difference(container1, container2)
	test.True(container3.KeyedEq(container4),t)
}
// Tests the Difference method functionality of a dynamic directed graph.
func DynDirectedGraphDifference(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphDifferenceHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,2,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](0,2,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.NoElem[int](),
			edges: iter.NoElem[int](),
			links: [][3]int{},
		},
		t,
	)
	directedGraphDifferenceHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 0},
				[3]int{2, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](1,4,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{1, 2, 0},
				[3]int{2, 3, 0},
				[3]int{3, 1, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{2, 0, 0},
			},
		},
		t,
	)
	directedGraphDifferenceHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 0},
				[3]int{2, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](1,4,1),
			edges: iter.Range[int](0,2,1),
			links: [][3]int{
				// from, to, e
				[3]int{1, 2, 1},
				[3]int{2, 3, 0},
				[3]int{3, 1, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				[3]int{0, 1, 0},
				[3]int{1, 2, 0},
				[3]int{2, 0, 0},
			},
		},
		t,
	)
	directedGraphDifferenceHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 0},
				[3]int{2, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](1,5,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{2, 3, 0},
				[3]int{3, 4, 0},
				[3]int{4, 2, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				[3]int{0, 1, 0},
				[3]int{1, 2, 0},
				[3]int{2, 0, 0},
			},
		},
		t,
	)
	directedGraphDifferenceHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,5,1),
			edges: iter.Range[int](0,2,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 0},
				[3]int{2, 3, 0},
				[3]int{3, 4, 0},
				[3]int{4, 0, 0},
				
				[3]int{0,2,1},
				[3]int{0,3,1},
				[3]int{1,3,1},
				[3]int{1,4,1},
				[3]int{2,4,1},
				[3]int{2,0,1},
				[3]int{3,0,1},
				[3]int{3,1,1},
				[3]int{4,1,1},
				[3]int{4,2,1},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](0,5,1),
			edges: iter.Range[int](1,2,1),
			links: [][3]int{
				// from, to, e
				[3]int{0,2,1},
				[3]int{0,3,1},
				[3]int{1,3,1},
				[3]int{1,4,1},
				[3]int{2,4,1},
				[3]int{2,0,1},
				[3]int{3,0,1},
				[3]int{3,1,1},
				[3]int{4,1,1},
				[3]int{4,2,1},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](0,5,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				[3]int{0, 1, 0},
				[3]int{1, 2, 0},
				[3]int{2, 3, 0},
				[3]int{3, 4, 0},
				[3]int{4, 0, 0},
			},
		},
		t,
	)
}

func directedGraphIsSupersetHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	g1 graphConstruction,
	g2 graphConstruction,
	expRes bool,
	t *testing.T,
) {
	container1:=g1.makeGraph(factory, t)
	container2:=g2.makeGraph(factory, t)
	test.Eq(expRes, container1.IsSuperset(container2),t)
	if container1.KeyedEq(container2) {
		test.Eq(expRes, container2.IsSuperset(container1),t)
	}
}
// Tests the IsSuperset method functionality of a dynamic directed graph.
func DynDirectedGraphIsSuperset(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphIsSupersetHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,2,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](0,2,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 0, 0},
			},
		},
		true,
		t,
	)
	directedGraphIsSupersetHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 0},
				[3]int{2, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](1,4,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{1, 2, 0},
				[3]int{2, 3, 0},
				[3]int{3, 1, 0},
			},
		},
		false,
		t,
	)
	directedGraphIsSupersetHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 0},
				[3]int{2, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](1,4,1),
			edges: iter.Range[int](0,2,1),
			links: [][3]int{
				// from, to, e
				[3]int{1, 2, 1},
				[3]int{2, 3, 0},
				[3]int{3, 1, 0},
			},
		},
		false,
		t,
	)
	directedGraphIsSupersetHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 0},
				[3]int{2, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](1,5,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{2, 3, 0},
				[3]int{3, 4, 0},
				[3]int{4, 2, 0},
			},
		},
		false,
		t,
	)
	directedGraphIsSupersetHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,5,1),
			edges: iter.Range[int](0,2,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 0},
				[3]int{2, 3, 0},
				[3]int{3, 4, 0},
				[3]int{4, 0, 0},
				
				[3]int{0,2,1},
				[3]int{0,3,1},
				[3]int{1,3,1},
				[3]int{1,4,1},
				[3]int{2,4,1},
				[3]int{2,0,1},
				[3]int{3,0,1},
				[3]int{3,1,1},
				[3]int{4,1,1},
				[3]int{4,2,1},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](0,5,1),
			edges: iter.Range[int](1,2,1),
			links: [][3]int{
				// from, to, e
				[3]int{0,2,1},
				[3]int{0,3,1},
				[3]int{1,3,1},
				[3]int{1,4,1},
				[3]int{2,4,1},
				[3]int{2,0,1},
				[3]int{3,0,1},
				[3]int{3,1,1},
				[3]int{4,1,1},
				[3]int{4,2,1},
			},
		},
		true,
		t,
	)
}

func directedGraphIsSubsetHelper(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	g1 graphConstruction,
	g2 graphConstruction,
	expRes bool,
	t *testing.T,
) {
	container1:=g1.makeGraph(factory, t)
	container2:=g2.makeGraph(factory, t)
	test.Eq(expRes, container1.IsSubset(container2),t)
	if container1.KeyedEq(container2) {
		test.Eq(expRes, container2.IsSubset(container1),t)
	}
}
// Tests the IsSuperset method functionality of a dynamic directed graph.
func DynDirectedGraphIsSubset(
	factory func(capacity int) dynamicContainers.DirectedGraph[int, int],
	t *testing.T,
) {
	directedGraphIsSubsetHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,2,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](0,2,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 0, 0},
			},
		},
		true,
		t,
	)
	directedGraphIsSubsetHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 0},
				[3]int{2, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](1,4,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{1, 2, 0},
				[3]int{2, 3, 0},
				[3]int{3, 1, 0},
			},
		},
		false,
		t,
	)
	directedGraphIsSubsetHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 0},
				[3]int{2, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](1,4,1),
			edges: iter.Range[int](0,2,1),
			links: [][3]int{
				// from, to, e
				[3]int{1, 2, 1},
				[3]int{2, 3, 0},
				[3]int{3, 1, 0},
			},
		},
		false,
		t,
	)
	directedGraphIsSubsetHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,3,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 0},
				[3]int{2, 0, 0},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](1,5,1),
			edges: iter.Range[int](0,1,1),
			links: [][3]int{
				// from, to, e
				[3]int{2, 3, 0},
				[3]int{3, 4, 0},
				[3]int{4, 2, 0},
			},
		},
		false,
		t,
	)
	directedGraphIsSubsetHelper(
		factory,
		graphConstruction{
			vertices: iter.Range[int](0,5,1),
			edges: iter.Range[int](1,2,1),
			links: [][3]int{
				// from, to, e
				[3]int{0,2,1},
				[3]int{0,3,1},
				[3]int{1,3,1},
				[3]int{1,4,1},
				[3]int{2,4,1},
				[3]int{2,0,1},
				[3]int{3,0,1},
				[3]int{3,1,1},
				[3]int{4,1,1},
				[3]int{4,2,1},
			},
		},
		graphConstruction{
			vertices: iter.Range[int](0,5,1),
			edges: iter.Range[int](0,2,1),
			links: [][3]int{
				// from, to, e
				[3]int{0, 1, 0},
				[3]int{1, 2, 0},
				[3]int{2, 3, 0},
				[3]int{3, 4, 0},
				[3]int{4, 0, 0},
				
				[3]int{0,2,1},
				[3]int{0,3,1},
				[3]int{1,3,1},
				[3]int{1,4,1},
				[3]int{2,4,1},
				[3]int{2,0,1},
				[3]int{3,0,1},
				[3]int{3,1,1},
				[3]int{4,1,1},
				[3]int{4,2,1},
			},
		},
		true,
		t,
	)
}
