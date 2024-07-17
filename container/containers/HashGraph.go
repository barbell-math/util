package containers

import (
	"fmt"
	"sync"

	"github.com/barbell-math/util/algo/hash"
	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/algo/widgets"
	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/container/containerTypes"
)

type (
	edgeHash              hash.Hash
	vertexHash            hash.Hash
	graphLink             basic.Pair[edgeHash, vertexHash]
	internalHashGraphImpl map[vertexHash]Vector[graphLink, *graphLink]

	// This is used when only the vertex part of a graph edge is needed
	vertexOnlyGraphLinkWidget graphLink
	// This is used when only the edge part of a graph edge is needed
	edgeOnlyGraphLinkWidget graphLink

	// A type to represent an arbitrary graph with the specified vertex and edge
	// types. The graph will maintain a set of vertices that are connected by a
	// set of edges. The type constraints on the generics define the logic for
	// for how specific operations, such as equality comparisons, will be
	// handled. The hash and equals methods defined in the widget types *must*
	// be congruent as they are both used when creating the graph internally.
	// The graph will grow as edges and vertices are added.
	HashGraph[
		V any,
		E any,
		VI widgets.WidgetInterface[V],
		EI widgets.WidgetInterface[E],
	] struct {
		numLinks int
		edges    map[edgeHash]E
		vertices map[vertexHash]V
		graph    internalHashGraphImpl
	}

	// A synchronized version of HashGraph. All operations will be wrapped in
	// the appropriate calls to the embedded RWMutex. A pointer to a RWMutex is
	// embedded rather than a value to avoid copying the lock value.
	SyncedHashGraph[
		V any,
		E any,
		VI widgets.WidgetInterface[V],
		EI widgets.WidgetInterface[E],
	] struct {
		*sync.RWMutex
		HashGraph[V, E, VI, EI]
	}
)

func (_ *graphLink) Eq(l *graphLink, r *graphLink) bool {
	hw := widgets.BuiltinHash{}
	return (hw.Eq((*hash.Hash)(&l.A), (*hash.Hash)(&r.A)) &&
		hw.Eq((*hash.Hash)(&l.B), (*hash.Hash)(&r.B)))
}
func (_ *graphLink) Lt(l *graphLink, r *graphLink) bool {
	hw := widgets.BuiltinHash{}
	return (hw.Lt((*hash.Hash)(&l.A), (*hash.Hash)(&r.A)) &&
		hw.Lt((*hash.Hash)(&l.B), (*hash.Hash)(&r.B)))
}
func (_ *graphLink) Hash(other *graphLink) hash.Hash {
	return ((hash.Hash)(other.A)).Combine((hash.Hash)(other.B))
}
func (_ *graphLink) Zero(other *graphLink) {
	*other = graphLink{A: edgeHash(0), B: vertexHash(0)}
}

func (_ *vertexOnlyGraphLinkWidget) Eq(l *graphLink, r *graphLink) bool {
	hw := widgets.BuiltinHash{}
	return hw.Eq((*hash.Hash)(&l.B), (*hash.Hash)(&r.B))
}
func (_ *vertexOnlyGraphLinkWidget) Lt(l *graphLink, r *graphLink) bool {
	hw := widgets.BuiltinHash{}
	return hw.Lt((*hash.Hash)(&l.B), (*hash.Hash)(&r.B))
}
func (_ *vertexOnlyGraphLinkWidget) Hash(other *graphLink) hash.Hash {
	return (hash.Hash)(other.B)
}
func (_ *vertexOnlyGraphLinkWidget) Zero(other *graphLink) {
	*other = graphLink{A: edgeHash(0), B: vertexHash(0)}
}

func (_ *edgeOnlyGraphLinkWidget) Eq(l *graphLink, r *graphLink) bool {
	hw := widgets.BuiltinHash{}
	return hw.Eq((*hash.Hash)(&l.A), (*hash.Hash)(&r.A))
}
func (_ *edgeOnlyGraphLinkWidget) Lt(l *graphLink, r *graphLink) bool {
	hw := widgets.BuiltinHash{}
	return hw.Lt((*hash.Hash)(&l.A), (*hash.Hash)(&r.A))
}
func (_ *edgeOnlyGraphLinkWidget) Hash(other *graphLink) hash.Hash {
	return (hash.Hash)(other.A)
}
func (_ *edgeOnlyGraphLinkWidget) Zero(other *graphLink) {
	*other = graphLink{A: edgeHash(0), B: vertexHash(0)}
}

// Creates a new hash graph initialized with enough memory to hold the specified
// amount of vertices and edges. Both numVertices and numEdges must be >=0, an
// error will be returned if either one violates that rule. If either size is 0
// then the associated map will be initialized with 0 elements.
func NewHashGraph[
	V any,
	E any,
	VI widgets.WidgetInterface[V],
	EI widgets.WidgetInterface[E],
](numVertices int, numEdges int) (HashGraph[V, E, VI, EI], error) {
	if numVertices < 0 {
		return HashGraph[V, E, VI, EI]{}, getSizeError(numVertices)
	}
	if numEdges < 0 {
		return HashGraph[V, E, VI, EI]{}, getSizeError(numEdges)
	}
	em := make(map[edgeHash]E, numEdges)
	vm := make(map[vertexHash]V, numVertices)
	g := make(internalHashGraphImpl, numVertices)
	return HashGraph[V, E, VI, EI]{
		numLinks: 0,
		edges:    em,
		vertices: vm,
		graph:    g,
	}, nil
}

// Creates a new hash graph initialized with enough memory to hold the specified
// amount of vertices and edges. Both numVertices and numEdges must be >=0, an
// error will be returned if either one violates that rule. If either size is 0
// then the associated map will be initialized with 0 elements. The underlying
// RWMutex value will be fully unlocked upon initialization.
func NewSyncedHashGraph[
	V any,
	E any,
	VI widgets.WidgetInterface[V],
	EI widgets.WidgetInterface[E],
](numVertices int, numEdges int) (SyncedHashGraph[V, E, VI, EI], error) {
	if numVertices < 0 {
		return SyncedHashGraph[V, E, VI, EI]{}, getSizeError(numVertices)
	}
	if numEdges < 0 {
		return SyncedHashGraph[V, E, VI, EI]{}, getSizeError(numEdges)
	}
	rv, err := NewHashGraph[V, E, VI, EI](numVertices, numEdges)
	return SyncedHashGraph[V, E, VI, EI]{
		RWMutex:   &sync.RWMutex{},
		HashGraph: rv,
	}, err
}

// A empty pass through function that performs no action. Needed for the
// [dynamicContainers.Comparisons] interface.
func (g *HashGraph[V, E, VI, EI]) Lock() {}

// A empty pass through function that performs no action. Needed for the
// [dynamicContainers.Comparisons] interface.
func (g *HashGraph[V, E, VI, EI]) Unlock() {}

// A empty pass through function that performs no action. Needed for the
// [dynamicContainers.Comparisons] interface.
func (g *HashGraph[V, E, VI, EI]) RLock() {}

// A empty pass through function that performs no action. Needed for the
// [dynamicContainers.Comparisons] interface.
func (g *HashGraph[V, E, VI, EI]) RUnlock() {}

// The SyncedHashGraph method to override the HashGraph pass through function
// and actually apply the mutex operation.
func (g *SyncedHashGraph[V, E, VI, EI]) Lock() { g.RWMutex.Lock() }

// The SyncedHashGraph method to override the HashGraph pass through function
// and actually apply the mutex operation.
func (g *SyncedHashGraph[V, E, VI, EI]) Unlock() { g.RWMutex.Unlock() }

// The SyncedHashGraph method to override the HashGraph pass through function
// and actually apply the mutex operation.
func (g *SyncedHashGraph[V, E, VI, EI]) RLock() { g.RWMutex.RLock() }

// The SyncedHashGraph method to override the HashGraph pass through function
// and actually apply the mutex operation.
func (g *SyncedHashGraph[V, E, VI, EI]) RUnlock() { g.RWMutex.RUnlock() }

// Returns false, hash graphs are not addressable. (Due to being built out of
// maps.)
func (g *HashGraph[V, E, VI, EI]) IsAddressable() bool { return false }

// Returns false, a hash graph is not synced.
func (g *HashGraph[V, E, VI, EI]) IsSynced() bool { return false }

// Returns true, a synced hash graph is synced.
func (g *SyncedHashGraph[V, E, VI, EI]) IsSynced() bool { return true }

// Description: NumEdges will return the number of edges in the graph. This will
// include any unconnected edges.
//
// Time Complexity: O(1)
func (g *HashGraph[V, E, VI, EI]) NumEdges() int {
	return len(g.edges)
}

// Description: NumEdges will return the number of edges in the graph. This will
// include any unconnected edges.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (g *SyncedHashGraph[V, E, VI, EI]) NumEdges() int {
	g.RLock()
	defer g.RUnlock()
	return len(g.edges)
}

// Description: NumVertices will return the number of edges in the graph. This
// will include any unconnected vertices.
//
// Time Complexity: O(1)
func (g *HashGraph[V, E, VI, EI]) NumVertices() int {
	return len(g.vertices)
}

// Description: NumVertices will return the number of edges in the graph. This
// will include any unconnected vertices.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (g *SyncedHashGraph[V, E, VI, EI]) NumVertices() int {
	g.RLock()
	defer g.RUnlock()
	return len(g.vertices)
}

// Description: NumLinks will return the number of links in the graph. This is
// different from the number of edges, as the number of links will not include
// any orphaned edges.
//
// Time Complexity: O(1)
func (g *HashGraph[V, E, VI, EI]) NumLinks() int {
	return g.numLinks
}

// Description: NumLinks will return the number of links in the graph. This is
// different from the number of edges, as the number of links will not include
// any orphaned edges.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (g *SyncedHashGraph[V, E, VI, EI]) NumLinks() int {
	g.RLock()
	defer g.RUnlock()
	return g.numLinks
}

// Description: Returns an iterator that iterates over the edges in the graph.
//
// Time Complexity: O(n), where n=num edges
func (g *HashGraph[V, E, VI, EI]) Edges() iter.Iter[E] {
	return iter.MapVals[edgeHash, E](g.edges)
}

// Description: Modifies the iterator chain returned by the underlying
// [HashGraph.Edges] method such that a read lock will be placed on the
// underlying hash graph when the iterator is consumed. The hash graph will have
// a read lock the entire time the iteration is being performed. The lock will
// not be applied until the iterator starts to be consumed.
//
// Lock Type: Read
//
// Time Complexity: O(n), where n=num edges
func (g *SyncedHashGraph[V, E, VI, EI]) Edges() iter.Iter[E] {
	return iter.MapVals[edgeHash, E](g.edges).SetupTeardown(
		func() error { g.RLock(); return nil },
		func() error { g.RUnlock(); return nil },
	)
}

// Panics, hash graphs are not addressable.
func (g *HashGraph[V, E, VI, EI]) EdgePntrs() iter.Iter[*E] {
	panic(getNonAddressablePanicText("hash graph"))
}

// Description: Returns an iterator that iterates over the vertices in the
// graph.
//
// Time Complexity: O(n), where n=num edges
func (g *HashGraph[V, E, VI, EI]) Vertices() iter.Iter[V] {
	return iter.MapVals[vertexHash, V](g.vertices)
}

// Description: Modifies the iterator chain returned by the underlying
// [HashGraph.Vertices] method such that a read lock will be placed on the
// underlying hash graph when the iterator is consumed. The hash graph will have
// a read lock the entire time the iteration is being performed. The lock will
// not be applied until the iterator starts to be consumed.
//
// Lock Type: Read
//
// Time Complexity: O(n), where n=num vertices
func (g *SyncedHashGraph[V, E, VI, EI]) Vertices() iter.Iter[V] {
	return iter.MapVals[vertexHash, V](g.vertices).SetupTeardown(
		func() error { g.RLock(); return nil },
		func() error { g.RUnlock(); return nil },
	)
}

// Panics, hash graphs are not addressable.
func (g *HashGraph[V, E, VI, EI]) VerticePntrs() iter.Iter[*V] {
	panic(getNonAddressablePanicText("hash graph"))
}

// Description: Returns true if the supplied vertex is contained within the
// graph. All equality comparisons are performed by the generic VI widget type
// that the hash graph was initialized with.
//
// Time Complexity: O(1)
func (g *HashGraph[V, E, VI, EI]) ContainsVertex(v V) bool {
	vw := widgets.Widget[V, VI]{}
	_, ok := g.vertices[vertexHash(vw.Hash(&v))]
	return ok
}

// Description: Places a read lock on the underlying graph before checking if
// the supplied vertex is contained in the graph, returning true if it is. All
// equality comparisons are performed by the generic VI widget type that the
// hash graph was initialized with.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (g *SyncedHashGraph[V, E, VI, EI]) ContainsVertex(v V) bool {
	g.RLock()
	defer g.RUnlock()
	vw := widgets.Widget[V, VI]{}
	_, ok := g.vertices[vertexHash(vw.Hash(&v))]
	return ok
}

// Description: ContainsVertexPntr will return true if the supplied vertex is in
// the hash graph, false otherwise. All equality comparisons are performed by
// the generic VI widget type that the hash graph was initialized with.
//
// Time Complexity: O(1)
func (g *HashGraph[V, E, VI, EI]) ContainsVertexPntr(v *V) bool {
	vw := widgets.Widget[V, VI]{}
	_, ok := g.vertices[vertexHash(vw.Hash(v))]
	return ok
}

// Description: Places a read lock on the underlying hash graph and then calls
// the underlying hash graph [HashGraph.ContainsVertexPntr] method.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (g *SyncedHashGraph[V, E, VI, EI]) ContainsVertexPntr(v *V) bool {
	g.RLock()
	defer g.RUnlock()
	vw := widgets.Widget[V, VI]{}
	_, ok := g.vertices[vertexHash(vw.Hash(v))]
	return ok
}

// Description: Returns true if the supplied edge is contained within the
// graph. All equality comparisons are performed by the generic EI widget type
// that the hash graph was initialized with.
//
// Time Complexity: O(1)
func (g *HashGraph[V, E, VI, EI]) ContainsEdge(e E) bool {
	ew := widgets.Widget[E, EI]{}
	_, ok := g.edges[edgeHash(ew.Hash(&e))]
	return ok
}

// Description: Places a read lock on the underlying graph before checking if
// the supplied edge is contained in the graph, returning true if it is. All
// equality comparisons are performed by the generic EI widget type that the
// hash graph was initialized with.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (g *SyncedHashGraph[V, E, VI, EI]) ContainsEdge(e E) bool {
	g.RLock()
	defer g.RUnlock()
	ew := widgets.Widget[E, EI]{}
	_, ok := g.edges[edgeHash(ew.Hash(&e))]
	return ok
}

// Description: ContainsEdgePntr will return true if the supplied edge is in the
// hash graph, false otherwise. All equality comparisons are performed by the
// generic EI widget type that the hash graph was initialized with.
//
// Time Complexity: O(1)
func (g *HashGraph[V, E, VI, EI]) ContainsEdgePntr(e *E) bool {
	ew := widgets.Widget[E, EI]{}
	_, ok := g.edges[edgeHash(ew.Hash(e))]
	return ok
}

// Description: Places a read lock on the underlying hash graph and then calls
// the underlying hash graph [HashGraph.ContainsEdgePntr] method.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (g *SyncedHashGraph[V, E, VI, EI]) ContainsEdgePntr(e *E) bool {
	g.RLock()
	defer g.RUnlock()
	ew := widgets.Widget[E, EI]{}
	_, ok := g.edges[edgeHash(ew.Hash(e))]
	return ok
}

// Description: Returns true if the supplied edge links the supplied vertices.
//
// Time Complexity: O(n), where n=num outgoing edges from the starting vertex.
func (g *HashGraph[V, E, VI, EI]) ContainsLink(from V, to V, e E) bool {
	return g.ContainsLinkPntr(&from, &to, &e)
}

// Description: Places a read lock on the underlying hash graph and then calls
// the underlying hash graph [HashGraph.ContainsLinkPntr] method. The pntr
// variant is called to avoid copying the V and E generic arguments twice, which
// could be expensive with large generic types.
//
// Lock Type: Read
//
// Time Complexity: O(n), where n=num outgoing edges from the starting vertex.
func (g *SyncedHashGraph[V, E, VI, EI]) ContainsLink(from V, to V, e E) bool {
	g.RLock()
	defer g.RUnlock()
	return g.HashGraph.ContainsLinkPntr(&from, &to, &e)
}

// Description: Returns true if the supplied edge links the supplied vertices.
//
// Time Complexity: O(n), where n=num outgoing edges from the starting vertex.
func (g *HashGraph[V, E, VI, EI]) ContainsLinkPntr(from *V, to *V, e *E) bool {
	vw := widgets.Widget[V, VI]{}
	ew := widgets.Widget[E, EI]{}
	fromHash := vertexHash(vw.Hash(from))
	toHash := vertexHash(vw.Hash(to))
	eHash := edgeHash(ew.Hash(e))

	if _, ok := g.vertices[fromHash]; !ok {
		return false
	}
	if _, ok := g.vertices[toHash]; !ok {
		return false
	}
	if _, ok := g.edges[eHash]; !ok {
		return false
	}

	gNode, _ := g.graph[fromHash]
	linkExists := false
	for i := 0; i < len(gNode) && !linkExists; i++ {
		linkExists = (gNode[i].A == eHash && gNode[i].B == toHash)
	}
	if !linkExists {
		return false
	}
	return true
}

// Description: Places a read lock on the underlying hash graph and then calls
// the underlying hash graph [HashGraph.ContainsLinkPntr] method.
//
// Lock Type: Read
//
// Time Complexity: O(n), where n=num outgoing edges from the starting vertex.
func (g *SyncedHashGraph[V, E, VI, EI]) ContainsLinkPntr(from *V, to *V, e *E) bool {
	g.RLock()
	defer g.RUnlock()
	return g.HashGraph.ContainsLinkPntr(from, to, e)
}

// Description: Returns the number of outgoing edges from the supplied vertex
//
// Time Complexity: O(1)
func (g *HashGraph[V, E, VI, EI]) NumOutLinks(v V) int {
	return g.NumOutLinksPntr(&v)
}

// Description: Places a read lock on the underlying hash graph and then calls
// the underlying hash graph [HashGraph.NumOutLinksPntr] method.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (g *SyncedHashGraph[V, E, VI, EI]) NumOutLinks(v V) int {
	g.RLock()
	defer g.RUnlock()
	return g.HashGraph.NumOutLinksPntr(&v)
}

// Description: Returns the number of outgoing edges from the supplied vertex
//
// Time Complexity: O(1)
func (g *HashGraph[V, E, VI, EI]) NumOutLinksPntr(v *V) int {
	vw := widgets.Widget[V, VI]{}
	vHash := vertexHash(vw.Hash(v))

	if _, ok := g.vertices[vHash]; !ok {
		return 0
	}
	if _, ok := g.graph[vHash]; !ok {
		return 0
	}
	return len(g.graph[vHash])
}

// Description: Places a read lock on the underlying hash graph and then calls
// the underlying hash graph [HashGraph.NumOutLinksPntr] method.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (g *SyncedHashGraph[V, E, VI, EI]) NumOutLinksPntr(v *V) int {
	g.RLock()
	defer g.RUnlock()
	return g.HashGraph.NumOutLinksPntr(v)
}

// Description: Returns an iterator that supplies all of the outgoing edges from
// the supplied vertex. Duplicate edges will not be filtered out, meaning a
// single edge may be returned multiple times by the iterator.
//
// Time Complexity: O(n), where n=num of outgoing edges
func (g *HashGraph[V, E, VI, EI]) OutEdges(v V) iter.Iter[E] {
	return g.outEdgesImpl(&v)
}

// Description: Modifies the iterator chain returned by the underlying
// [HashGraph.OutEdges] method such that a read lock will be placed on the
// underlying hash graph when the iterator is consumed. The hash graph will have
// a read lock the entire time the iteration is being performed. The lock will
// not be applied until the iterator chain starts to be consumed.
//
// Lock Type: Read
//
// Time Complexity: O(n), where n=num of outgoing edges
func (g *SyncedHashGraph[V, E, VI, EI]) OutEdges(v V) iter.Iter[E] {
	return g.HashGraph.outEdgesImpl(&v).SetupTeardown(
		func() error { g.RLock(); return nil },
		func() error { g.RUnlock(); return nil },
	)
}

// Panics, hash graphs are non-addressable.
func (g *HashGraph[V, E, VI, EI]) OutEdgePntrs(v *V) iter.Iter[*E] {
	panic(getNonAddressablePanicText("hash graph"))
}

func (g *HashGraph[V, E, VI, EI]) outEdgesImpl(v *V) iter.Iter[E] {
	vw := widgets.Widget[V, VI]{}
	vHash := vertexHash(vw.Hash(v))
	if _, ok := g.vertices[vHash]; !ok {
		var tmp E
		return iter.ValElem[E](tmp, getVertexError[V](v), 1)
	}
	if _, ok := g.graph[vHash]; !ok {
		// It is a valid vertex, just has no out going edges
		return iter.NoElem[E]()
	}
	return iter.Map[graphLink, E](
		iter.SliceElems[graphLink](g.graph[vHash]),
		func(index int, val graphLink) (E, error) {
			return g.edges[val.A], nil
		},
	)
}

// Description: Returns an iterator that supplies all of the outgoing vertices
// from the supplied vertex. Duplicate vertices will not be filtered out,
// meaning a single vertex may be returned multiple times by the iterator.
//
// Time Complexity: O(n), where n=num of outgoing edges
func (g *HashGraph[V, E, VI, EI]) OutVertices(v V) iter.Iter[V] {
	return g.outVerticesImpl(&v)
}

// Description: Modifies the iterator chain returned by the underlying
// [HashGraph.OutVertices] method such that a read lock will be placed on the
// underlying hash graph when the iterator is consumed. The hash graph will have
// a read lock the entire time the iteration is being performed. The lock will
// not be applied until the iterator chain starts to be consumed.
//
// Lock Type: Read
//
// Time Complexity: O(n), where n=num of outgoing edges
func (g *SyncedHashGraph[V, E, VI, EI]) OutVertices(v V) iter.Iter[V] {
	return g.HashGraph.outVerticesImpl(&v).SetupTeardown(
		func() error { g.RLock(); return nil },
		func() error { g.RUnlock(); return nil },
	)
}

// Panics, hash graphs are non-addressable
func (g *HashGraph[V, E, VI, EI]) OutVerticePntrs(v *V) iter.Iter[*V] {
	panic(getNonAddressablePanicText("hash graph"))
}

func (g *HashGraph[V, E, VI, EI]) outVerticesImpl(v *V) iter.Iter[V] {
	vw := widgets.Widget[V, VI]{}
	vHash := vertexHash(vw.Hash(v))
	if _, ok := g.vertices[vHash]; !ok {
		var tmp V
		return iter.ValElem[V](tmp, getVertexError[V](v), 1)
	}
	if _, ok := g.graph[vHash]; !ok {
		// It is a valid vertex, just has no out going edges
		return iter.NoElem[V]()
	}
	return iter.Map[graphLink, V](
		iter.SliceElems[graphLink](g.graph[vHash]),
		func(index int, val graphLink) (V, error) {
			return g.vertices[val.B], nil
		},
	)
}

// Description: Returns an iterator that supplies all of the outgoing edges
// paired with there associated vertices.
//
// Time Complexity: O(n), where n=num of outgoing edges
func (g *HashGraph[V, E, VI, EI]) OutEdgesAndVertices(
	v V,
) iter.Iter[basic.Pair[E, V]] {
	return g.outEdgesAndVerticesImpl(&v)
}

// Description: Modifies the iterator chain returned by the underlying
// [HashGraph.OutEdgesAndVertices] method such that a read lock will be placed
// on the underlying hash graph when the iterator is consumed. The hash graph
// will have a read lock the entire time the iteration is being performed. The
// lock will not be applied until the iterator chain starts to be consumed.
//
// Lock Type: Read
//
// Time Complexity: O(n), where n=num of outgoing edges
func (g *SyncedHashGraph[V, E, VI, EI]) OutEdgeAndVertices(
	v V,
) iter.Iter[basic.Pair[E, V]] {
	return g.HashGraph.outEdgesAndVerticesImpl(&v).SetupTeardown(
		func() error { g.RLock(); return nil },
		func() error { g.RUnlock(); return nil },
	)
}

func (g *HashGraph[V, E, VI, EI]) outEdgesAndVerticesImpl(
	v *V,
) iter.Iter[basic.Pair[E, V]] {
	vw := widgets.Widget[V, VI]{}
	vHash := vertexHash(vw.Hash(v))
	if _, ok := g.vertices[vHash]; !ok {
		var tmp basic.Pair[E, V]
		return iter.ValElem[basic.Pair[E, V]](
			tmp,
			getVertexError[V](v),
			1,
		)
	}
	if _, ok := g.graph[vHash]; !ok {
		// It is a valid vertex, just has no out going edges
		return iter.NoElem[basic.Pair[E, V]]()
	}
	return iter.Map[graphLink, basic.Pair[E, V]](
		iter.SliceElems[graphLink](g.graph[vHash]),
		func(index int, val graphLink) (basic.Pair[E, V], error) {
			return basic.Pair[E, V]{g.edges[val.A], g.vertices[val.B]}, nil
		},
	)
}

// Panics, hash graphs are non-addressable.
func (g *HashGraph[V, E, VI, EI]) OutEdgesAndVerticePntrs(
	v *V,
) iter.Iter[basic.Pair[*E, *V]] {
	panic(getNonAddressablePanicText("hash graph"))
}

// Description: Returns the list of edges that exist between the supplied
// vertices. Any returned edges will follow the direction specified by the
// arguments.
//
// Time Complexity: O(n), where n=the number of outgoing edges on the from
// vertex
func (g *HashGraph[V, E, VI, EI]) EdgesBetween(from V, to V) iter.Iter[E] {
	return g.edgesBetweenImpl(&from, &to)
}

// Description: Modifies the iterator chain returned by the underlying
// [HashGraph.EdgesBetween] method such that a read lock will be placed
// on the underlying hash graph when the iterator is consumed. The hash graph
// will have a read lock the entire time the iteration is being performed. The
// lock will not be applied until the iterator chain starts to be consumed.
//
// Lock Type: Read
//
// Time Complexity: O(n), where n=the number of outgoing edges on the from
// vertex
func (g *SyncedHashGraph[V, E, VI, EI]) EdgesBetween(from V, to V) iter.Iter[E] {
	return g.HashGraph.edgesBetweenImpl(&from, &to).SetupTeardown(
		func() error { g.RLock(); return nil },
		func() error { g.RUnlock(); return nil },
	)
}

func (g *HashGraph[V, E, VI, EI]) edgesBetweenImpl(from *V, to *V) iter.Iter[E] {
	vw := widgets.Widget[V, VI]{}
	fromHash := vertexHash(vw.Hash(from))
	toHash := vertexHash(vw.Hash(to))

	if _, ok := g.vertices[fromHash]; !ok {
		var tmp E
		return iter.ValElem[E](tmp, getVertexError[V](from), 1)
	}
	if _, ok := g.vertices[toHash]; !ok {
		var tmp E
		return iter.ValElem[E](tmp, getVertexError[V](to), 1)
	}
	if _, ok := g.graph[fromHash]; !ok {
		// The from vertex is a valid vertex, just has no outgoing edges.
		return iter.NoElem[E]()
	}

	return iter.Map[graphLink, E](
		iter.SliceElems[graphLink](g.graph[fromHash]).Filter(
			func(index int, val graphLink) bool {
				return val.B == toHash
			},
		),
		func(index int, val graphLink) (E, error) {
			return g.edges[val.A], nil
		},
	)
}

// Panics, hash graphs are non-addressable.
func (g *HashGraph[V, E, VI, EI]) EdgesBetweenPntr(from *V, to *V) iter.Iter[*E] {
	panic(getNonAddressablePanicText("hash graph"))
}

// TODO - probably delete
// func (g *HashGraph[V, E, VI, EI])Links() iter.Iter[struct{from V; to V; e E}] {
// 	type tmpHash struct {
// 		from vertexHash
// 		to vertexHash
// 		e edgeHash
// 	}
// 	prevIndex:=-1
// 	return iter.Map[tmpHash, struct {from, to V; e E}](
// 		iter.Map[vertexHash, tmpHash](
// 			iter.MapKeys[vertexHash](g.graph),
// 			func(index int, val vertexHash,) (tmpHash, error) {
// 				return struct{from vertexHash; to vertexHash; e edgeHash}{
// 					from: val,
// 				}, nil
// 			},
// 		).Next(
// 			func(
// 				index int,
// 				val tmpHash,
// 				status iter.IteratorFeedback,
// 			) (
// 				iter.IteratorFeedback,
// 				tmpHash,
// 				error,
// 			) {
// 				if status==iter.Break {
// 					return iter.Break, tmpHash{}, nil
// 				}
// 				if prevIndex<len(g.graph[val.from]) {
// 					prevIndex++
// 					return iter.Continue, tmpHash{
// 						from: val.from,
// 						to: g.graph[val.from][prevIndex].B,
// 						e: g.graph[val.from][prevIndex].A,
// 					}, nil
// 				}
// 				prevIndex=-1
// 				return iter.Iterate, tmpHash{}, nil
// 			},
// 		),
// 		func(index int, val tmpHash) (struct{from V; to V; e E}, error) {
// 			return struct{from V; to V; e E}{
// 				from: g.vertices[val.from],
// 				to: g.vertices[val.to],
// 				e: g.edges[val.e],
// 			}, nil
// 		},
// 	)
// }
//
// func (g *HashGraph[V, E, VI, EI])LinkPntrs() iter.Iter[struct{from *V; to *V; e *E}] {
// 	return iter.NoElem[struct{from *V; to *V; e *E}]()
// }

// Description: Adds edges to the graph without connecting them to any vertices.
// Duplicate edges will be ignored. This method will never return an error.
//
// Time Complexity: O(n), where n=len(e)
func (g *HashGraph[V, E, VI, EI]) AddEdges(e ...E) error {
	ew := widgets.Widget[E, EI]{}
	for _, iterE := range e {
		iterEHash := edgeHash(ew.Hash(&iterE))
		if _, ok := g.edges[iterEHash]; !ok {
			g.edges[iterEHash] = iterE
		}
	}
	return nil
}

// Description: Places a write lock on the underlying hash graph before calling
// the underlying hash graphs [HashGraph.AddEdges] method.
//
// Lock Type: Write
//
// Time Complexity: O(n), where n=len(e)
func (g *SyncedHashGraph[V, E, VI, EI]) AddEdges(e ...E) error {
	g.Lock()
	defer g.Unlock()
	return g.HashGraph.AddEdges(e...)
}

// Description: Adds a vertex to the graph, if it does not already exist
// according to the hash and equals method on the vertex widget interface. Non-
// unique vertices will not be added. This function will never return an error.
//
// Time Complexity: O(n), where n=len(v)
func (g *HashGraph[V, E, VI, EI]) AddVertices(v ...V) error {
	ew := widgets.Widget[V, VI]{}
	for _, iterV := range v {
		iterVHash := vertexHash(ew.Hash(&iterV))
		if _, ok := g.vertices[iterVHash]; !ok {
			g.vertices[iterVHash] = iterV
		}
	}
	return nil
}

// Description: Places a write lock on the underlying hash graph and then adds
// the vertices to the underlying graph. Exhibits the same behavior as the
// non-synced [HashGraph.AddVertices] method.
//
// Lock Type: Write
//
// Time Complexity: O(n), where n=len(v)
func (g *SyncedHashGraph[V, E, VI, EI]) AddVertices(v ...V) error {
	g.Lock()
	defer g.Unlock()
	return g.HashGraph.AddVertices(v...)
}

// Description: Adds a link between an existing edge and vertices in the graph.
// The edge and vertices must have been added to the graph prior to calling
// this function or an error will be returned. If a link already exists between
// the provided vertices with the provided edge then no action will be taken and
// no error will be returned.
//
// Time Complexity: O(n), where n=num of outgoing edges from the start vertex
func (g *HashGraph[V, E, VI, EI]) Link(from V, to V, e E) error {
	return g.LinkPntr(&from, &to, &e)
}

// Description: Places a write lock on the underlying hash graph and then adds
// the link to the graph. Exhibits the same behavior as the non-synced
// [HashGraph.Link] method.
//
// Lock Type: Write
//
// Time Complexity: O(n), where n=num of outgoing edges from the start vertex
func (g *SyncedHashGraph[V, E, VI, EI]) Link(from V, to V, e E) error {
	g.Lock()
	defer g.Unlock()
	return g.HashGraph.LinkPntr(&from, &to, &e)
}

// Description: Adds a link between an existing edge and vertices in the graph.
// The edge and vertices must have been added to the graph prior to calling
// this function or an error will be returned. If a link already exists between
// the provided vertices with the provided edge then no action will be taken and
// no error will be returned.
//
// Time Complexity: O(n), where n=num of outgoing edges from the start vertex
func (g *HashGraph[V, E, VI, EI]) LinkPntr(from *V, to *V, e *E) error {
	vw := widgets.Widget[V, VI]{}
	ew := widgets.Widget[E, EI]{}
	fromHash := vertexHash(vw.Hash(from))
	toHash := vertexHash(vw.Hash(to))
	eHash := edgeHash(ew.Hash(e))

	if _, ok := g.vertices[fromHash]; !ok {
		return getVertexError[V](from)
	}
	if _, ok := g.vertices[toHash]; !ok {
		return getVertexError[V](to)
	}
	if _, ok := g.edges[eHash]; !ok {
		return getEdgeError[E](e)
	}

	gl := graphLink{eHash, toHash}
	gNode, _ := g.graph[fromHash]
	if gNode.Contains(gl) {
		return nil
	}

	g.numLinks++
	gNode.Append(graphLink{eHash, toHash})
	g.graph[fromHash] = gNode
	if gNode, ok := g.graph[toHash]; !ok {
		g.graph[toHash] = gNode
	}
	return nil
}

// Description: Places a write lock on the underlying hash graph and then adds
// the link to the graph. Exhibits the same behavior as the non-synced
// [HashGraph.LinkPntr] method.
//
// Lock Type: Write
//
// Time Complexity: O(n), where n=num of outgoing edges on the from vertex
func (g *SyncedHashGraph[V, E, VI, EI]) LinkPntr(from *V, to *V, e *E) error {
	g.Lock()
	defer g.Unlock()
	return g.HashGraph.LinkPntr(from, to, e)
}

// Description: Deletes a vertex from the graph, removing any links that
// previously used the vertex. No edges will be deleted, meaning this operation
// may result in orphaned edges.
//
// Time Complexity: O(n), where n=num of links in the graph
func (g *HashGraph[V, E, VI, EI]) DeleteVertex(v V) error {
	return g.DeleteVertexPntr(&v)
}

// Description: Places a write lock on the underlying hash graph and then
// removes the vertex from the graph. Exhibits the same behavior as the
// non-synced [HashGraph.DeleteVertexPntr] method.
//
// Lock Type: Write
//
// Time Complexity: O(n), where n=num of links in the graph
func (g *SyncedHashGraph[V, E, VI, EI]) DeleteVertex(v V) error {
	g.Lock()
	defer g.Unlock()
	return g.HashGraph.DeleteVertexPntr(&v)
}

// Description: Deletes a vertex from the graph, removing any links that
// previously used the vertex. No edges will be deleted, meaning this operation
// may result in orphaned edges.
//
// Time Complexity: O(n), where n=num of links in the graph
func (g *HashGraph[V, E, VI, EI]) DeleteVertexPntr(v *V) error {
	vw := widgets.Widget[V, VI]{}
	vHash := vertexHash(vw.Hash(v))

	if _, ok := g.vertices[vHash]; !ok {
		return getVertexError[V](v)
	}
	tmpV := g.vertices[vHash]
	vw.Zero(&tmpV)
	delete(g.vertices, vHash)

	if gNode, ok := g.graph[vHash]; ok {
		g.numLinks -= gNode.Length()
		delete(g.graph, vHash)
	}

	for iterHash, _ := range g.graph {
		gNode := (Vector[graphLink, *vertexOnlyGraphLinkWidget])(
			([]graphLink)(g.graph[iterHash]),
		)
		g.numLinks -= gNode.Pop(graphLink{B: vHash})
		if len(gNode) == 0 {
			delete(g.graph, iterHash)
		} else {
			g.graph[iterHash] = (Vector[graphLink, *graphLink])(
				([]graphLink)(gNode),
			)
		}
	}
	return nil
}

// Description: Places a write lock on the underlying hash graph and then
// removes the vertex from the graph. Exhibits the same behavior as the
// non-synced [HashGraph.DeleteVertexPntr] method.
//
// Lock Type: Write
//
// Time Complexity: O(n), where n=num of links in the graph
func (g *SyncedHashGraph[V, E, VI, EI]) DeleteVertexPntr(v *V) error {
	g.Lock()
	defer g.Unlock()
	return g.HashGraph.DeleteVertexPntr(v)
}

// Description: Deletes an edge from the graph, removing any links that
// previously used the vertex. No vertices will be deleted, meaning this
// operation may result in orphaned vertices.
//
// Time Complexity: O(n), where n=num of links in the graph
func (g *HashGraph[V, E, VI, EI]) DeleteEdge(e E) error {
	return g.DeleteEdgePntr(&e)
}

// Description: Places a write lock on the underlying hash graph and then
// removes the edge from the graph. Exhibits the same behavior as the
// non-synced [HashGraph.DeleteEdgePntr] method.
//
// Lock Type: Write
//
// Time Complexity: O(n), where n=num of links in the graph
func (g *SyncedHashGraph[V, E, VI, EI]) DeleteEdge(e E) error {
	g.Lock()
	defer g.Unlock()
	return g.HashGraph.DeleteEdgePntr(&e)
}

// Description: Deletes an edge from the graph, removing any links that
// previously used the vertex. No vertices will be deleted, meaning this
// operation may result in orphaned vertices.
//
// Time Complexity: O(n), where n=num of links in the graph
func (g *HashGraph[V, E, VI, EI]) DeleteEdgePntr(e *E) error {
	ew := widgets.Widget[E, EI]{}
	eHash := edgeHash(ew.Hash(e))

	if _, ok := g.edges[eHash]; !ok {
		return getEdgeError[E](e)
	}
	tmpE := g.edges[eHash]
	ew.Zero(&tmpE)
	delete(g.edges, eHash)

	for iterHash, _ := range g.graph {
		gNode := (Vector[graphLink, *edgeOnlyGraphLinkWidget])(
			([]graphLink)(g.graph[iterHash]),
		)
		g.numLinks -= gNode.Pop(graphLink{A: eHash})
		if len(gNode) == 0 {
			delete(g.graph, iterHash)
		} else {
			g.graph[iterHash] = (Vector[graphLink, *graphLink])(
				([]graphLink)(gNode),
			)
		}
	}
	return nil
}

// Description: Places a write lock on the underlying hash graph and then
// removes the edge from the graph. Exhibits the same behavior as the
// non-synced [HashGraph.DeleteEdgePntr] method.
//
// Lock Type: Write
//
// Time Complexity: O(n), where n=num of links in the graph
func (g *SyncedHashGraph[V, E, VI, EI]) DeleteEdgePntr(e *E) error {
	g.Lock()
	defer g.Unlock()
	return g.HashGraph.DeleteEdgePntr(e)
}

// Description: Removes a link within the graph without removing the underlying
// vertices or edge. This may leave vertices with no links, and edges that don't
// correspond to any links. An error will be returned if either vertice does not
// exist in the graph or if the supplied edge does not exist in the graph.
//
// Time Complexity: O(n), where n=num of outgoing edges on the from vertex
func (g *HashGraph[V, E, VI, EI]) DeleteLink(from V, to V, e E) error {
	return g.DeleteLinkPntr(&from, &to, &e)
}

// Description: Places a write lock on the underlying hash graph before calling
// the underlying hash graphs [HashGraph.DeleteLinkPntr] method.
//
// Lock Type: Write
//
// Time Complexity: O(n), where n=num of outgoing edges on the from vertex
func (g *SyncedHashGraph[V, E, VI, EI]) DeleteLink(from V, to V, e E) error {
	g.Lock()
	defer g.Unlock()
	return g.HashGraph.DeleteLinkPntr(&from, &to, &e)
}

// Description: Removes a link within the graph without removing the underlying
// vertices or edge. This may leave vertices with no links, and edges that don't
// correspond to any links. An error will be returned if either vertice does not
// exist in the graph or if the supplied edge does not exist in the graph.
//
// Time Complexity: O(n), where n=num of outgoing edges on the from vertex
func (g *HashGraph[V, E, VI, EI]) DeleteLinkPntr(from *V, to *V, e *E) error {
	vw := widgets.Widget[V, VI]{}
	ew := widgets.Widget[E, EI]{}
	fromHash := vertexHash(vw.Hash(from))
	toHash := vertexHash(vw.Hash(to))
	eHash := edgeHash(ew.Hash(e))

	if _, ok := g.vertices[fromHash]; !ok {
		return getVertexError[V](from)
	}
	if _, ok := g.vertices[toHash]; !ok {
		return getVertexError[V](to)
	}
	if _, ok := g.edges[eHash]; !ok {
		return getEdgeError[E](e)
	}

	gNode, _ := g.graph[fromHash]
	if idx, found := gNode.KeyOf(graphLink{eHash, toHash}); found {
		if len(gNode) > 1 {
			gNode.Delete(idx)
			g.numLinks--
			g.graph[fromHash] = gNode
		} else if len(gNode) == 1 {
			delete(g.graph, fromHash)
		}
	}
	return nil
}

// Description: Places a write lock on the underlying hash graph before calling
// the underlying hash graphs [HashGraph.DeleteLinkPntr] method.
//
// Lock Type: Write
//
// Time Complexity: O(n), where n=num of outgoing edges on the from vertex
func (g *SyncedHashGraph[V, E, VI, EI]) DeleteLinkPntr(
	from *V,
	to *V,
	e *E,
) error {
	g.Lock()
	defer g.Unlock()
	return g.HashGraph.DeleteLinkPntr(from, to, e)
}

// Description: Removes all links starting at from and ending at to without
// removing the underlying vertices or edges. This may leave vertices with no
// links, and edges that don't correspond to any links. An error will be
// returned if either vertex does not exist in the graph.
//
// Time Complexity: O(n), where n=num of outgoing edges on the from vertex
func (g *HashGraph[V, E, VI, EI]) DeleteLinks(from V, to V) error {
	return g.DeleteLinksPntr(&from, &to)
}

// Description: Places a write lock on the underlying hash graph before calling
// the underlying hash graphs [HashGraph.DeleteLinksPntr] method.
//
// Lock Type: Write
//
// Time Complexity: O(n), where n=num of outgoing edges on the from vertex
func (g *SyncedHashGraph[V, E, VI, EI]) DeleteLinks(from V, to V) error {
	g.Lock()
	defer g.Unlock()
	return g.HashGraph.DeleteLinksPntr(&from, &to)
}

// Description: Removes all links starting at from and ending at to without
// removing the underlying vertices or edges. This may leave vertices with no
// links, and edges that don't correspond to any links. An error will be
// returned if either vertex does not exist in the graph.
//
// Time Complexity: O(n), where n=num of outgoing edges on the from vertex
func (g *HashGraph[V, E, VI, EI]) DeleteLinksPntr(from *V, to *V) error {
	vw := widgets.Widget[V, VI]{}
	fromHash := vertexHash(vw.Hash(from))
	toHash := vertexHash(vw.Hash(to))

	if _, ok := g.vertices[fromHash]; !ok {
		return getVertexError[V](from)
	}
	if _, ok := g.vertices[toHash]; !ok {
		return getVertexError[V](to)
	}

	gNode := (Vector[graphLink, *vertexOnlyGraphLinkWidget])(
		([]graphLink)(g.graph[fromHash]),
	)
	g.numLinks -= gNode.Pop(graphLink{B: toHash})

	if len(gNode) == 0 {
		delete(g.graph, fromHash)
	} else {
		g.graph[fromHash] = (Vector[graphLink, *graphLink])(([]graphLink)(gNode))
	}
	return nil
}

// Description: Places a write lock on the underlying hash graph before calling
// the underlying hash graphs [HashGraph.DeleteLinksPntr] method.
//
// Lock Type: Write
//
// Time Complexity: O(n), where n=num of outgoing edges on the from vertex
func (g *SyncedHashGraph[V, E, VI, EI]) DeleteLinksPntr(from *V, to *V) error {
	g.Lock()
	g.Unlock()
	return g.HashGraph.DeleteLinksPntr(from, to)
}

// Description: Removes all edges, vertices, and links.
//
// Time Complexity: O(n+m), where n=num vertices and m=num edges
func (g *HashGraph[V, E, VI, EI]) Clear() {
	ew := widgets.Widget[E, EI]{}
	vw := widgets.Widget[V, VI]{}
	for _, iterE := range g.edges {
		ew.Zero(&iterE)
	}
	for _, iterV := range g.vertices {
		vw.Zero(&iterV)
	}

	g.edges = make(map[edgeHash]E)
	g.vertices = make(map[vertexHash]V)
	g.graph = make(internalHashGraphImpl)
	g.numLinks = 0
}

// Description: Places a write lock on the underlying hash graph before calling
// the underlying [HashGraph.Clear] method.
//
// Lock Type: Write
//
// Time Complexity: O(n+m), where n=num vertices and m=num edges
func (g *SyncedHashGraph[V, E, VI, EI]) Clear() {
	g.Lock()
	defer g.Unlock()
	g.HashGraph.Clear()
}

// Description: Returns true if the two supplied graphs are considered equal.
// In order for two graphs to be equal they must have the same structure and all
// of the corresponding vertices and edges must be equal as defined by the Eq
// method of the supplied VI and EI widgets.
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the methods on other. In big-O it might look something like this,
// O(n*O(other.ContainsLink) + m*O(other.ContainsVertexPntr) + p*O(other.ContainsEdgePntr))
// where:
//  - n is the number of links in this graph
//  - m is the number of vertices in this graph
//  - p is the number of edges in this graph
//  - O(other.ContainsLink) represents the time complexity of the ContainsLink method on other
//  - O(other.ContainsVertexPntr) represents the time complexity of the ContainsVertexPntr method on other
//  - O(other.ContainsEdgePntr) represents the time complexity of the ContainsEdgePntr method on other
func (g *HashGraph[V, E, VI, EI]) KeyedEq(
	other containerTypes.GraphComparisonsConstraint[V, E],
) bool {
	if !(g.NumEdges() == other.NumEdges() &&
		g.NumVertices() == other.NumVertices() &&
		g.NumLinks() == other.NumLinks()) {
		return false
	}

	for _,v:=range(g.vertices) {
		if !other.ContainsVertexPntr(&v) {
			return false
		}
	}
	for _,e:=range(g.edges) {
		if !other.ContainsEdgePntr(&e) {
			return false
		}
	}

	for from, gNode := range g.graph {
		if len(gNode) != other.NumOutLinks(g.vertices[from]) {
			return false
		}
		for _, gLink := range gNode {
			if !other.ContainsLink(
				g.vertices[from],
				g.vertices[gLink.B],
				g.edges[gLink.A],
			) {
				return false
			}
		}
	}
	return true
}

// Description: Places a read lock on the underlying hash graph and then calls
// the underlying hash graph [HashGraph.KeyedEq] method. Attempts to place a
// read lock on other but whether or not that happens is implementation
// dependent.
//
// Lock Type: Read on this hash graph, read on other
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the methods on other. In big-O it might look something like this,
// O(n*O(other.ContainsLink) + m*O(other.ContainsVertexPntr) + p*O(other.ContainsEdgePntr))
// where:
//  - n is the number of links in this graph
//  - m is the number of vertices in this graph
//  - p is the number of edges in this graph
//  - O(other.ContainsLink) represents the time complexity of the ContainsLink method on other
//  - O(other.ContainsVertexPntr) represents the time complexity of the ContainsVertexPntr method on other
//  - O(other.ContainsEdgePntr) represents the time complexity of the ContainsEdgePntr method on other
func (g *SyncedHashGraph[V, E, VI, EI]) KeyedEq(
	other containerTypes.GraphComparisonsConstraint[V, E],
) bool {
	g.RLock()
	other.RLock()
	defer g.RUnlock()
	defer other.RUnlock()
	return g.HashGraph.KeyedEq(other)
}

// Isomorphic equality
func (g *HashGraph[V, E, VI, EI]) UnorderedEq(
	other containerTypes.GraphComparisonsConstraint[V, E],
) bool {
	return false
}

// Description: Takes the intersection of l and r and puts the result in this
// graph. All values from this graph will be cleared before storing the new
// result. Vertices and edges are compared using the supplied VI and EI widgets.
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the ContainsLinkPntr method on other. In big-O it might look something like
// this, O(n*O(other.ContainsLinkPntr))), where n is the number of links in r
// and O(other.ContainsLinkPntr) represents the time complexity of the
// ContainsLinkPntr method on other.
func (g *HashGraph[V, E, VI, EI]) Intersection(
	l containerTypes.GraphComparisonsConstraint[V, E],
	r containerTypes.GraphComparisonsConstraint[V, E],
) {
	newG,err:=NewHashGraph[V,E,VI,EI](r.NumVertices()/2, r.NumEdges()/2)
	if err!=nil {
		panic(fmt.Sprintf("An error occurred making a new hash graph: %s",err))
	}
	// This implementation chooses to optimize the case where a link is not
	// created in the graph. It does this by using pointers to values whenever
	// possible. Note that in the case when a link must be made that values
	// will be copied from out of scope, which might entail the GC.
	addressableSafeVerticesIter[V,E](r).ForEach(
		func(index int, fromV *V) (iter.IteratorFeedback, error) {
			if !l.ContainsVertexPntr(fromV) {
				return iter.Continue, nil
			}
			addressableSafeOutVerticesAndEdgesIter[V,E](r,*fromV).ForEach(
				func(
					index int,
					toVAndE basic.Pair[*E, *V],
				) (iter.IteratorFeedback, error) {
					if l.ContainsLinkPntr(fromV, toVAndE.B, toVAndE.A) {
						newG.AddEdges(*toVAndE.A)
						newG.AddVertices(*fromV, *toVAndE.B)
						newG.LinkPntr(fromV, toVAndE.B, toVAndE.A)
					}
					return iter.Continue, nil
				},
			)
			return iter.Continue, nil
		},
	)
	g.Clear()
	*g=newG
}

// Description: Places a write lock on the underlying hash graph and then calls
// the underlying hash graph [HashGraph.Intersection] method. Attempts to place
// a read lock on l and r but whether or not that happens is implementation
// dependent.
//
// Lock Type: Write on this vector, read on l and r
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the ContainsLinkPntr method on other. In big-O it might look something like
// this, O(n*O(other.ContainsLinkPntr))), where n is the number of links in r
// and O(other.ContainsLinkPntr) represents the time complexity of the
// ContainsLinkPntr method on other.
func (g *SyncedHashGraph[V, E, VI, EI]) Intersection(
	l containerTypes.GraphComparisonsConstraint[V,E],
	r containerTypes.GraphComparisonsConstraint[V,E],
) {
	g.Lock()
	l.RLock()
	r.RLock()
	defer g.Unlock()
	defer l.RUnlock()
	defer r.RUnlock()
	g.HashGraph.Intersection(l,r)
}

// Description: Takes the union of l and r and puts the result in this
// graph. All values from this graph will be cleared before storing the new
// result. Vertices and edges are compared using the supplied VI and EI widgets.
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the ContainsLinkPntr method on other. In big-O it might look something like
// this, O((n+m)*O(other.ContainsLinkPntr))), where n is the number of links in
// r and m is the number of links in l. O(other.ContainsLinkPntr) represents the
// time complexity of the ContainsLinkPntr method on other.
func (g *HashGraph[V, E, VI, EI]) Union(
	l containerTypes.GraphComparisonsConstraint[V, E],
	r containerTypes.GraphComparisonsConstraint[V, E],
) {
	newG,err:=NewHashGraph[V,E,VI,EI](r.NumVertices()/2, r.NumEdges()/2)
	if err!=nil {
		panic(fmt.Sprintf("An error occurred making a new hash graph: %s",err))
	}
	opTemplate:= func(fromV *V) func(
		index int,
		toVAndE basic.Pair[*E, *V],
	) (iter.IteratorFeedback, error) {
		return func(
			index int,
			toVAndE basic.Pair[*E, *V],
		) (iter.IteratorFeedback, error) {
			newG.AddEdges(*toVAndE.A)
			newG.AddVertices(*fromV, *toVAndE.B)
			newG.LinkPntr(fromV, toVAndE.B, toVAndE.A)
			return iter.Continue, nil
		}
	}
	// This implementation chooses to optimize the case where a link is not
	// created in the graph. It does this by using pointers to values whenever
	// possible. Note that in the case when a link must be made that values
	// will be copied from out of scope, which might entail the GC.
	addressableSafeVerticesIter[V,E](r).ForEach(
		func(index int, fromV *V) (iter.IteratorFeedback, error) {
			op:=opTemplate(fromV)
			addressableSafeOutVerticesAndEdgesIter[V,E](r,*fromV).ForEach(op)
			return iter.Continue, nil
		},
	)
	addressableSafeVerticesIter[V,E](l).ForEach(
		func(index int, fromV *V) (iter.IteratorFeedback, error) {
			op:=opTemplate(fromV)
			addressableSafeOutVerticesAndEdgesIter[V,E](l,*fromV).ForEach(op)
			return iter.Continue, nil
		},
	)
	g.Clear()
	*g=newG
}

// Description: Places a write lock on the underlying hash graph and then calls
// the underlying hash graph [HashGraph.Union] method. Attempts to place
// a read lock on l and r but whether or not that happens is implementation
// dependent.
//
// Lock Type: Write on this vector, read on l and r
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the ContainsLinkPntr method on other. In big-O it might look something like
// this, O((n+m)*O(other.ContainsLinkPntr))), where n is the number of links in
// r and m is the number of links in l. O(other.ContainsLinkPntr) represents the
// time complexity of the ContainsLinkPntr method on other.
func (g *SyncedHashGraph[V, E, VI, EI])Union(
	l containerTypes.GraphComparisonsConstraint[V, E],
	r containerTypes.GraphComparisonsConstraint[V, E],
) {
	g.Lock()
	l.RLock()
	r.RLock()
	defer g.Unlock()
	defer l.RUnlock()
	defer r.RUnlock()
	g.HashGraph.Union(l,r)
}

// Description: Takes the difference  of l and r and puts the result in this
// graph. All values from this graph will be cleared before storing the new
// result. Vertices and edges are compared using the supplied VI and EI widgets.
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the ContainsLinkPntr method on other. In big-O it might look something like
// this, O((n+m)*O(other.ContainsLinkPntr))), where n is the number of links in
// r and m is the number of links in l. O(other.ContainsLinkPntr) represents the
// time complexity of the ContainsLinkPntr method on other.
func (g *HashGraph[V, E, VI, EI]) Difference(
	l containerTypes.GraphComparisonsConstraint[V, E],
	r containerTypes.GraphComparisonsConstraint[V, E],
) {
	newG,err:=NewHashGraph[V,E,VI,EI](r.NumVertices()/2, r.NumEdges()/2)
	if err!=nil {
		panic(fmt.Sprintf("An error occurred making a new hash graph: %s",err))
	}
	opTemplate:= func(fromV *V) func(
		index int,
		toVAndE basic.Pair[*E, *V],
	) (iter.IteratorFeedback, error) {
		return func(
			index int,
			toVAndE basic.Pair[*E, *V],
		) (iter.IteratorFeedback, error) {
			if r.ContainsLinkPntr(fromV,toVAndE.B,toVAndE.A) {
				return iter.Continue, nil
			}
			newG.AddEdges(*toVAndE.A)
			newG.AddVertices(*fromV, *toVAndE.B)
			newG.LinkPntr(fromV, toVAndE.B, toVAndE.A)
			return iter.Continue, nil
		}
	}
	// This implementation chooses to optimize the case where a link is not
	// created in the graph. It does this by using pointers to values whenever
	// possible. Note that in the case when a link must be made that values
	// will be copied from out of scope, which might entail the GC.
	addressableSafeVerticesIter[V,E](l).ForEach(
		func(index int, fromV *V) (iter.IteratorFeedback, error) {
			op:=opTemplate(fromV)
			addressableSafeOutVerticesAndEdgesIter[V,E](l,*fromV).ForEach(op)
			return iter.Continue, nil
		},
	)
	g.Clear()
	*g=newG
}

// Description: Places a write lock on the underlying hash graph and then calls
// the underlying hash graph [HashGraph.Difference] method. Attempts to place
// a read lock on l and r but whether or not that happens is implementation
// dependent.
//
// Lock Type: Write on this vector, read on l and r
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the ContainsLinkPntr method on other. In big-O it might look something like
// this, O((n+m)*O(other.ContainsLinkPntr))), where n is the number of links in
// r and m is the number of links in l. O(other.ContainsLinkPntr) represents the
// time complexity of the ContainsLinkPntr method on other.
func (g *SyncedHashGraph[V, E, VI, EI])Difference(
	l containerTypes.GraphComparisonsConstraint[V, E],
	r containerTypes.GraphComparisonsConstraint[V, E],
) {
	g.Lock()
	r.RLock()
	l.RLock()
	defer g.Unlock()
	defer r.RUnlock()
	defer l.RUnlock()
	g.HashGraph.Difference(l,r)
}

// Description: Returns true if this graph is a superset of other. In order for 
// this graph to be a superset of other, it must have all of others vertices,
// edges, and links. It may have other vertices, edges, or links that are not in
// other. All of the corresponding vertices and edges must be equal as defined
// by the Eq method of the supplied VI and EI widgets.
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the methods on other. In big-O it might look something like this,
// O(n*O(other.ContainsLink) + m*O(other.ContainsVertexPntr) + p*O(other.ContainsEdgePntr))
// where:
//  - n is the number of links in this graph
//  - m is the number of vertices in this graph
//  - p is the number of edges in this graph
//  - O(other.ContainsLink) represents the time complexity of the ContainsLink method on other
//  - O(other.ContainsVertexPntr) represents the time complexity of the ContainsVertexPntr method on other
//  - O(other.ContainsEdgePntr) represents the time complexity of the ContainsEdgePntr method on other
func (g *HashGraph[V, E, VI, EI]) IsSuperset(
	other containerTypes.GraphComparisonsConstraint[V, E],
) bool {
	if (g.NumVertices()<other.NumVertices() ||
		g.NumEdges()<other.NumEdges() ||
		g.NumLinks()<other.NumLinks()) {
		return false
	}

	rv:=true
	addressableSafeVerticesIter[V,E](other).ForEach(
		func(index int, fromV *V) (iter.IteratorFeedback, error) {
			if rv=g.ContainsVertexPntr(fromV); !rv {
				return iter.Break, nil
			}
			return iter.Continue, nil
		},
	)
	if !rv {
		return false
	}

	addressableSafeEdgesIter[V,E](other).ForEach(
		func(index int, fromV *E) (iter.IteratorFeedback, error) {
			if rv=g.ContainsEdgePntr(fromV); !rv {
				return iter.Break, nil
			}
			return iter.Continue, nil
		},
	)
	if !rv {
		return false
	}

	addressableSafeVerticesIter[V,E](other).ForEach(
		func(index int, fromV *V) (iter.IteratorFeedback, error) {
			addressableSafeOutVerticesAndEdgesIter[V,E](other,*fromV).ForEach(
				func(index int, val basic.Pair[*E, *V]) (iter.IteratorFeedback, error) {
					if rv=g.ContainsLinkPntr(fromV, val.B, val.A); !rv {
						return iter.Break, nil
					}
					return iter.Continue, nil
				},
			)
			return iter.Continue, nil
		},
	)
	return rv
}

// Description: Places a read lock on the underlying hash graph and then calls
// the underlying hash graph [HashGraph.IsSuperSet] method. Attempts to place a
// read lock on other but whether or not that happens is implementation
// dependent.
//
// Lock Type: Read on this hash graph, read on other
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the methods in other. In big-O it might look something like this,
// O(n*O(other.ContainsLink) + m*O(other.ContainsVertexPntr) + p*O(other.ContainsEdgePntr))
// where:
//  - n is the number of links in this graph
//  - m is the number of vertices in this graph
//  - p is the number of edges in this graph
//  - O(other.ContainsLink) represents the time complexity of the ContainsLink method on other
//  - O(other.ContainsVertexPntr) represents the time complexity of the ContainsVertexPntr method on other
//  - O(other.ContainsEdgePntr) represents the time complexity of the ContainsEdgePntr method on other
func (g *SyncedHashGraph[V, E, VI, EI])IsSuperSet(
	other containerTypes.GraphComparisonsConstraint[V, E],
) bool {
	g.RLock()
	other.RLock()
	defer g.RUnlock()
	defer other.RUnlock()
	return g.HashGraph.IsSuperset(other)
}

// Description: Returns true if this graph is a subset of other. In order for 
// this graph to be a subset of other, other must have all of this graphs
// vertices, edges, and links. Other may have other vertices, edges, or links
// that are not in this graph. All of the corresponding vertices and edges must
// be equal as defined by the Eq method of the supplied VI and EI widgets.
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the methods on other. In big-O it might look something like this,
// O(n*O(other.ContainsLink) + m*O(other.ContainsVertexPntr) + p*O(other.ContainsEdgePntr))
// where:
//  - n is the number of links in this graph
//  - m is the number of vertices in this graph
//  - p is the number of edges in this graph
//  - O(other.ContainsLink) represents the time complexity of the ContainsLink method on other
//  - O(other.ContainsVertexPntr) represents the time complexity of the ContainsVertexPntr method on other
//  - O(other.ContainsEdgePntr) represents the time complexity of the ContainsEdgePntr method on other
func (g *HashGraph[V, E, VI, EI]) IsSubset(
	other containerTypes.GraphComparisonsConstraint[V, E],
) bool {
	if (g.NumVertices()>other.NumVertices() ||
		g.NumEdges()>other.NumEdges() ||
		g.NumLinks()>other.NumLinks()) {
		return false
	}

	for _, v:=range(g.vertices) {
		if !other.ContainsVertexPntr(&v) {
			return false
		}
	}
	for _, e:=range(g.edges) {
		if !other.ContainsEdgePntr(&e) {
			return false
		}
	}

	for vHash, gNode:=range(g.graph) {
		for _, gLink:=range(gNode) {
			if !other.ContainsLink(
				g.vertices[vHash],
				g.vertices[gLink.B],
				g.edges[gLink.A],
			) {
				return false
			}
		}
	}
	return true
}

// Description: Places a read lock on the underlying hash graph and then calls
// the underlying hash graph [HashGraph.IsSubset] method. Attempts to place a
// read lock on other but whether or not that happens is implementation
// dependent.
//
// Lock Type: Read on this hash graph, read on other
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the methods in other. In big-O it might look something like this,
// O(n*O(other.ContainsLink) + m*O(other.ContainsVertexPntr) + p*O(other.ContainsEdgePntr))
// where:
//  - n is the number of links in this graph
//  - m is the number of vertices in this graph
//  - p is the number of edges in this graph
//  - O(other.ContainsLink) represents the time complexity of the ContainsLink method on other
//  - O(other.ContainsVertexPntr) represents the time complexity of the ContainsVertexPntr method on other
//  - O(other.ContainsEdgePntr) represents the time complexity of the ContainsEdgePntr method on other
func (g *SyncedHashGraph[V, E, VI, EI])IsSubset(
	other containerTypes.GraphComparisonsConstraint[V, E],
) bool {
	g.RLock()
	other.RLock()
	defer g.RUnlock()
	defer other.RUnlock()
	return g.HashGraph.IsSubset(other)
}

func (_ *HashGraph[V, E, VI, EI]) Eq(
	l *HashGraph[V, E, VI, EI],
	r *HashGraph[V, E, VI, EI],
) bool {
	return l.KeyedEq(r)
}

func (_ *SyncedHashGraph[V, E, VI, EI])Eq(
	l *SyncedHashGraph[V, E, VI, EI],
	r *SyncedHashGraph[V, E, VI, EI],
) bool {
	l.RLock()
	r.RLock()
	defer l.RUnlock()
	defer r.RUnlock()
	return l.HashGraph.KeyedEq(r)
}

// Panics, hash graphs cannot be compared for order.
func (_ *HashGraph[V, E, VI, EI]) Lt(
	l *HashGraph[V, E, VI, EI],
	r *HashGraph[V, E, VI, EI],
) bool {
	panic("Hash graphs maps cannot be compared relative to each other.")
}

// Panics, hash graphs cannot be compared for order.
func (_ *SyncedHashGraph[V, E, VI, EI]) Lt(
	l *SyncedHashGraph[V, E, VI, EI],
	r *SyncedHashGraph[V, E, VI, EI],
) bool {
	panic("Hash graphs maps cannot be compared relative to each other.")
}

func (_ *HashGraph[V, E, VI, EI]) Hash(other *HashGraph[V, E, VI, EI]) hash.Hash {
	return hash.Hash(0)
}

// An zero function that implements the [algo.widget.WidgetInterface] interface.
// Internally this is equivalent to [HashGraph.Clear].
func (_ *HashGraph[V, E, VI, EI]) Zero(other *HashGraph[V, E, VI, EI]) {
	other.Clear()
}

// An zero function that implements the [algo.widget.WidgetInterface] interface.
// Internally this is equivalent to [SyncedHashGraph.Clear].
func (_ *SyncedHashGraph[V, E, VI, EI]) Zero(other *SyncedHashGraph[V, E, VI, EI]) {
	other.Clear()
}
