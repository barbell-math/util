package containers

import (
	"sync"

	"github.com/barbell-math/util/algo/hash"
	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/algo/widgets"
	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/container/dynamicContainers"
)

type (
	edgeHash hash.Hash
	vertexHash hash.Hash
	graphEdge basic.Pair[edgeHash, vertexHash]

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
		edges map[edgeHash]E
		vertices map[vertexHash]V
		graph map[vertexHash][]graphEdge
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
		HashGraph[V,E,VI,EI]
	}
)

// Creates a new hash graph initialized with enough memory to hold the specified
// amount of vertices and edges. Both numVertices and numEdges must be >=0, an
// error will be returned if either one violates that rule. If either size is 0
// then the associated map will be initialized with 0 elements.
func NewHashGraph[
	V any,
	E any,
	VI widgets.WidgetInterface[V],
	EI widgets.WidgetInterface[E],
](numVertices int, numEdges int) (HashGraph[V,E,VI,EI],error) {
	if numVertices<0 {
		return HashGraph[V, E, VI, EI]{}, getSizeError(numVertices)
	}
	if numEdges<0 {
		return HashGraph[V, E, VI, EI]{}, getSizeError(numEdges)
	}
	em:=make(map[edgeHash]E,numEdges)
	vm:=make(map[vertexHash]V,numVertices)
	g:=make(map[vertexHash][]graphEdge,numVertices)
	return HashGraph[V, E, VI, EI]{
		numLinks: 0,
		edges: em,
		vertices: vm,
		graph: g,
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
](numVertices int, numEdges int) (SyncedHashGraph[V,E,VI,EI],error) {
	if numVertices<0 {
		return SyncedHashGraph[V, E, VI, EI]{}, getSizeError(numVertices)
	}
	if numEdges<0 {
		return SyncedHashGraph[V, E, VI, EI]{}, getSizeError(numEdges)
	}
	rv,err:=NewHashGraph[V,E,VI,EI](numVertices, numEdges)
	return SyncedHashGraph[V, E, VI, EI]{
		RWMutex: &sync.RWMutex{},
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
func (g *HashGraph[V,E,VI,EI]) IsSynced() bool { return false }

// Returns true, a synced hash graph is synced.
func (g *SyncedHashGraph[V,E,VI,EI]) IsSynced() bool { return true }

// Description: NumEdges will return the number of edges in the graph. This will
// include any unconnected edges.
//
// Time Complexity: O(1)
func (g *HashGraph[V,E,VI,EI])NumEdges() int {
	return len(g.edges)
}
// Description: NumEdges will return the number of edges in the graph. This will
// include any unconnected edges.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (g *SyncedHashGraph[V,E,VI,EI])NumEdges() int {
	g.RLock()
	defer g.RUnlock()
	return len(g.edges)
}

// Description: NumVertices will return the number of edges in the graph. This
// will include any unconnected vertices.
//
// Time Complexity: O(1)
func (g *HashGraph[V,E,VI,EI])NumVertices() int {
	return len(g.vertices)
}

// Description: NumVertices will return the number of edges in the graph. This
// will include any unconnected vertices.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (g *SyncedHashGraph[V,E,VI,EI])NumVertices() int {
	g.RLock()
	defer g.RUnlock()
	return len(g.vertices)
}

// Description: NumLinks will return the number of links in the graph. This is
// different from the number of edges, as the number of links will not include
// any orphaned edges.
//
// Time Complexity: O(1)
func (g *HashGraph[V,E,VI,EI])NumLinks() int {
	return g.numLinks
}
// Description: NumLinks will return the number of links in the graph. This is
// different from the number of edges, as the number of links will not include
// any orphaned edges.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (g *SyncedHashGraph[V,E,VI,EI])NumLinks() int {
	g.RLock()
	defer g.RUnlock()
	return g.numLinks
}

// Description: Returns an iterator that iterates over the edges in the graph.
//
// Time Complexity: O(n), where n=num edges
func (g *HashGraph[V,E,VI,EI])Edges() iter.Iter[E] {
	return iter.MapVals[edgeHash,E](g.edges)
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
func (g *SyncedHashGraph[V,E,VI,EI])Edges() iter.Iter[E] {
	return iter.MapVals[edgeHash,E](g.edges).SetupTeardown(
		func() error { g.RLock(); return nil },
		func() error { g.RUnlock(); return nil },
	)
}

// Panics, hash graphs are not addressable.
func (g *HashGraph[V,E,VI,EI]) EdgePntrs() iter.Iter[*E] {
	panic(getNonAddressablePanicText("hash graph"))
}

// Description: Returns an iterator that iterates over the vertices in the
// graph.
//
// Time Complexity: O(n), where n=num edges
func (g *HashGraph[V,E,VI,EI])Vertices() iter.Iter[V] {
	return iter.MapVals[vertexHash,V](g.vertices)
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
func (g *SyncedHashGraph[V,E,VI,EI])Vertices() iter.Iter[V] {
	return iter.MapVals[vertexHash,V](g.vertices).SetupTeardown(
		func() error { g.RLock(); return nil },
		func() error { g.RUnlock(); return nil },
	)
}

// Panics, hash graphs are not addressable.
func (g *HashGraph[V,E,VI,EI]) VerticePntrs() iter.Iter[*V] {
	panic(getNonAddressablePanicText("hash graph"))
}

// Description: Returns true if the supplied vertex is contained within the
// graph. All equality comparisons are performed by the generic VI widget type
// that the hash graph was initialized with.
//
// Time Complexity: O(1)
func (g *HashGraph[V,E,VI,EI])ContainsVertex(v V) bool {
	vw:=widgets.Widget[V,VI]{}
	_,ok:=g.vertices[vertexHash(vw.Hash(&v))]
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
func (g *SyncedHashGraph[V,E,VI,EI])ContainsVertex(v V) bool {
	g.RLock()
	defer g.RUnlock()
	vw:=widgets.Widget[V,VI]{}
	_,ok:=g.vertices[vertexHash(vw.Hash(&v))]
	return ok
}

// Description: ContainsVertexPntr will return true if the supplied vertex is in
// the hash graph, false otherwise. All equality comparisons are performed by
// the generic VI widget type that the hash graph was initialized with.
//
// Time Complexity: O(1)
func (g *HashGraph[V,E,VI,EI])ContainsVertexPntr(v *V) bool {
	vw:=widgets.Widget[V,VI]{}
	_,ok:=g.vertices[vertexHash(vw.Hash(v))]
	return ok
}

// Description: Places a read lock on the underlying hash graph and then calls
// the underlying hash graph [HashGraph.ContainsVertexPntr] method.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (g *SyncedHashGraph[V,E,VI,EI])ContainsVertexPntr(v *V) bool {
	g.RLock()
	defer g.RUnlock()
	vw:=widgets.Widget[V,VI]{}
	_,ok:=g.vertices[vertexHash(vw.Hash(v))]
	return ok
}

// Description: Returns true if the supplied edge is contained within the
// graph. All equality comparisons are performed by the generic EI widget type
// that the hash graph was initialized with.
//
// Time Complexity: O(1)
func (g *HashGraph[V,E,VI,EI])ContainsEdge(e E) bool {
	ew:=widgets.Widget[E,EI]{}
	_,ok:=g.edges[edgeHash(ew.Hash(&e))]
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
func (g *SyncedHashGraph[V,E,VI,EI])ContainsEdge(e E) bool {
	g.RLock()
	defer g.RUnlock()
	ew:=widgets.Widget[E,EI]{}
	_,ok:=g.edges[edgeHash(ew.Hash(&e))]
	return ok
}

// Description: ContainsEdgePntr will return true if the supplied edge is in the
// hash graph, false otherwise. All equality comparisons are performed by the
// generic EI widget type that the hash graph was initialized with.
//
// Time Complexity: O(1)
func (g *HashGraph[V,E,VI,EI])ContainsEdgePntr(e *E) bool {
	ew:=widgets.Widget[E,EI]{}
	_,ok:=g.edges[edgeHash(ew.Hash(e))]
	return ok
}

// Description: Places a read lock on the underlying hash graph and then calls
// the underlying hash graph [HashGraph.ContainsEdgePntr] method.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (g *SyncedHashGraph[V,E,VI,EI])ContainsEdgePntr(e *E) bool {
	g.RLock()
	defer g.RUnlock()
	ew:=widgets.Widget[E,EI]{}
	_,ok:=g.edges[edgeHash(ew.Hash(e))]
	return ok
}

// Description: Returns true if the supplied edge links the supplied vertices.
//
// Time Complexity: O(n), where n=num outgoing edges from the starting vertex.
func (g *HashGraph[V,E,VI,EI])ContainsLink(from V, to V, e E) bool {
	return g.ContainsLinkPntr(&from,&to,&e)
}

// Description: Places a read lock on the underlying hash graph and then calls
// the underlying hash graph [HashGraph.ContainsLinkPntr] method. The pntr
// variant is called to avoid copying the V and E generic arguments twice, which
// could be expensive with large generic types.
//
// Lock Type: Read
//
// Time Complexity: O(n), where n=num outgoing edges from the starting vertex.
func (g *SyncedHashGraph[V,E,VI,EI])ContainsLink(from V, to V, e E) bool {
	g.RLock()
	defer g.RUnlock()
	return g.HashGraph.ContainsLinkPntr(&from,&to,&e)
}

// Description: Returns true if the supplied edge links the supplied vertices.
//
// Time Complexity: O(n), where n=num outgoing edges from the starting vertex.
func (g *HashGraph[V,E,VI,EI])ContainsLinkPntr(from *V, to *V, e *E) bool {
	vw:=widgets.Widget[V,VI]{}
	ew:=widgets.Widget[E,EI]{}
	fromHash:=vertexHash(vw.Hash(from))
	toHash:=vertexHash(vw.Hash(to))
	eHash:=edgeHash(ew.Hash(e))

	if _,ok:=g.vertices[fromHash]; !ok {
		return false
	}
	if _,ok:=g.vertices[toHash]; !ok {
		return false
	}
	if _,ok:=g.edges[eHash]; !ok {
		return false
	}

	gNode,_:=g.graph[fromHash]
	linkExists:=false
	for i:=0; i<len(gNode) && !linkExists; i++ {
		linkExists=(gNode[i].A==eHash && gNode[i].B==toHash)
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
func (g *SyncedHashGraph[V,E,VI,EI])ContainsLinkPntr(from *V, to *V, e *E) bool {
	g.RLock()
	defer g.RUnlock()
	return g.HashGraph.ContainsLinkPntr(from,to,e)
}

// Description: Returns an iterator that supplies all of the outgoing edges from
// the supplied vertex. Duplicate edges will not be filtered out, meaning a
// single edge may be returned multiple times by the iterator.
//
// Time Complexity: O(n), where n=num of outgoing edges
func (g *HashGraph[V,E,VI,EI])OutEdges(v V) iter.Iter[E] {
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
func (g *SyncedHashGraph[V, E, VI, EI])OutEdges(v V) iter.Iter[E] {
	return g.HashGraph.outEdgesImpl(&v).SetupTeardown(
		func() error { g.RLock(); return nil },
		func() error { g.RUnlock(); return nil },
	)
}

// Panics, hash graphs are non-addressable.
func (g *HashGraph[V, E, VI, EI])OutEdgePntrs(v *V) iter.Iter[*E] {
	panic(getNonAddressablePanicText("hash graph"))
}

func (g *HashGraph[V, E, VI, EI])outEdgesImpl(v *V) iter.Iter[E] {
	vw:=widgets.Widget[V,VI]{}
	vHash:=vertexHash(vw.Hash(v))
	if _,ok:=g.vertices[vHash]; !ok {
		var tmp E
		return iter.ValElem[E](tmp,getVertexError[V](v),1)
	}
	if _,ok:=g.graph[vHash]; !ok {
		// It is a valid vertex, just has no out going edges
		return iter.NoElem[E]()
	}
	return iter.Map[graphEdge,E](
		iter.SliceElems[graphEdge](g.graph[vHash]),
		func(index int, val graphEdge) (E, error) {
			return g.edges[val.A], nil
		},
	)
}

// Description: Returns an iterator that supplies all of the outgoing vertices
// from the supplied vertex. Duplicate vertices will not be filtered out, 
// meaning a single vertex may be returned multiple times by the iterator.
//
// Time Complexity: O(n), where n=num of outgoing edges
func (g *HashGraph[V,E,VI,EI])OutVertices(v V) iter.Iter[V] {
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
func (g *SyncedHashGraph[V,E,VI,EI])OutVertices(v V) iter.Iter[V] {
	return g.HashGraph.outVerticesImpl(&v).SetupTeardown(
		func() error { g.RLock(); return nil },
		func() error { g.RUnlock(); return nil },
	)
}

// Panics, hash graphs are non-addressable
func (g *HashGraph[V, E, VI, EI])OutVerticePntrs(v *V) iter.Iter[*V] {
	panic(getNonAddressablePanicText("hash graph"))
}

func (g *HashGraph[V,E,VI,EI])outVerticesImpl(v *V) iter.Iter[V] {
	vw:=widgets.Widget[V,VI]{}
	vHash:=vertexHash(vw.Hash(v))
	if _,ok:=g.vertices[vHash]; !ok {
		var tmp V
		return iter.ValElem[V](tmp,getVertexError[V](v),1)
	}
	if _,ok:=g.graph[vHash]; !ok {
		// It is a valid vertex, just has no out going edges
		return iter.NoElem[V]()
	}
	return iter.Map[graphEdge,V](
		iter.SliceElems[graphEdge](g.graph[vHash]),
		func(index int, val graphEdge) (V, error) {
			return g.vertices[val.B], nil
		},
	)
}

// Description: Returns an iterator that supplies all of the outgoing edges
// paired with there associated vertices.
//
// Time Complexity: O(n), where n=num of outgoing edges
func (g *HashGraph[V,E,VI,EI])OutEdgesAndVertices(
	v V,
) iter.Iter[basic.Pair[E,V]] {
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
func (g *SyncedHashGraph[V,E,VI,EI])OutEdgeAndVertices(
	v V,
) iter.Iter[basic.Pair[E,V]] {
	return g.HashGraph.outEdgesAndVerticesImpl(&v).SetupTeardown(
		func() error { g.RLock(); return nil },
		func() error { g.RUnlock(); return nil },
	)
}

func (g *HashGraph[V,E,VI,EI])outEdgesAndVerticesImpl(
	v *V,
) iter.Iter[basic.Pair[E,V]] {
	vw:=widgets.Widget[V,VI]{}
	vHash:=vertexHash(vw.Hash(v))
	if _,ok:=g.vertices[vHash]; !ok {
		var tmp basic.Pair[E,V]
		return iter.ValElem[basic.Pair[E,V]](
			tmp,
			getVertexError[V](v),
			1,
		)
	}
	if _,ok:=g.graph[vHash]; !ok {
		// It is a valid vertex, just has no out going edges
		return iter.NoElem[basic.Pair[E,V]]()
	}
	return iter.Map[graphEdge,basic.Pair[E,V]](
		iter.SliceElems[graphEdge](g.graph[vHash]),
		func(index int, val graphEdge) (basic.Pair[E,V], error) {
			return basic.Pair[E,V]{g.edges[val.A], g.vertices[val.B]}, nil
		},
	)
}

// Panics, hash graphs are non-addressable.
func (g *HashGraph[V,E,VI,EI])OutEdgesAndVerticePntrs(
	v *V,
) iter.Iter[basic.Pair[*E,*V]] {
	panic(getNonAddressablePanicText("hash graph"))
}

// Description: Returns the list of edges that exist between the supplied
// vertices. Any returned edges will follow the direction specified by the
// arguments.
//
// Time Complexity: O(n), where n=the number of edges that are returned
func (g *HashGraph[V,E,VI,EI])EdgesBetween(from V, to V) iter.Iter[E] {
	vw:=widgets.Widget[V,VI]{}
	fromHash:=vertexHash(vw.Hash(&from))
	toHash:=vertexHash(vw.Hash(&to))

	if _,ok:=g.vertices[fromHash]; !ok {
		var tmp E
		return iter.ValElem[E](tmp,getVertexError[V](&from),1)
	}
	if _,ok:=g.vertices[toHash]; !ok {
		var tmp E
		return iter.ValElem[E](tmp,getVertexError[V](&to),1)
	}
	if _,ok:=g.graph[fromHash]; !ok {
		// The from vertex is a valid vertex, just has no outgoing edges.
		return iter.NoElem[E]()
	}

	return iter.Map[graphEdge,E](
		iter.SliceElems[graphEdge](g.graph[fromHash]).Filter(
			func(index int, val graphEdge) bool {
				return val.B==toHash
			},
		),
		func(index int, val graphEdge) (E, error) {
			return g.edges[val.A], nil
		},
	)
}

// Panics, hash graphs are non-addressable.
func (g *HashGraph[V,E,VI,EI])EdgesBetweenPntr(from *V, to *V) iter.Iter[*E] {
	panic(getNonAddressablePanicText("hash graph"))
}

// Description: Adds edges to the graph without connecting them to any vertices.
// Duplicate edges will be ignored. This method will never return an error.
//
// Time Complexity: O(n), where n=len(e)
func (g *HashGraph[V,E,VI,EI])AddEdges(e ...E) error {
	ew:=widgets.Widget[E,EI]{}
	for _,iterE:=range(e) {
		iterEHash:=edgeHash(ew.Hash(&iterE))
		if _,ok:=g.edges[iterEHash]; !ok {
			g.edges[iterEHash]=iterE
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
func (g *SyncedHashGraph[V,E,VI,EI])AddEdges(e ...E) error {
	g.Lock()
	defer g.Unlock()
	return g.HashGraph.AddEdges(e...)
}

// Description: Adds a vertex to the graph, if it does not already exist
// according to the hash and equals method on the vertex widget interface. Non-
// unique vertices will not be added. This function will never return an error.
//
// Time Complexity: O(n), where n=len(v)
func (g *HashGraph[V,E,VI,EI])AddVertices(v ...V) error {
	ew:=widgets.Widget[V,VI]{}
	for _,iterV:=range(v) {
		iterVHash:=vertexHash(ew.Hash(&iterV))
		if _,ok:=g.vertices[iterVHash]; !ok {
			g.vertices[iterVHash]=iterV
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
func (g *SyncedHashGraph[V,E,VI,EI])AddVertices(v ...V) error {
	g.Lock()
	defer g.Unlock()
	return g.HashGraph.AddVertices(v...)
}

// Description: Adds a link between an existing edge and vertices in the graph.
// The edge and vertices must have been added to the graph prior to calling
// this funciton or an error will be returned. If a link already exists between
// the provided vertices with the provided edge then no action will be taken and
// no error will be returned.
//
// Time Complexity: O(n), where n=num of outgoing edges from the start vertex
func (g *HashGraph[V,E,VI,EI])Link(from V, to V, e E) error {
	return g.LinkPntr(&from,&to,&e)
}

// Description: Places a write lock on the underlying hash graph and then adds
// the link to the graph. Exhibits the same behavior as the non-synced
// [HashGraph.Link] method.
//
// Lock Type: Write
//
// Time Complexity: O(n), where n=num of outgoing edges from the start vertex
func (g *SyncedHashGraph[V, E, VI, EI])Link(from V, to V, e E) error {
	g.Lock()
	defer g.Unlock()
	return g.HashGraph.LinkPntr(&from,&to,&e)
}

// Description: Adds a link between an existing edge and vertices in the graph.
// The edge and vertices must have been added to the graph prior to calling
// this funciton or an error will be returned. If a link already exists between
// the provided vertices with the provided edge then no action will be taken and
// no error will be returned.
//
// Time Complexity: O(n), where n=num of outgoing edges from the start vertex
func (g *HashGraph[V,E,VI,EI])LinkPntr(from *V, to *V, e *E) error {
	vw:=widgets.Widget[V,VI]{}
	ew:=widgets.Widget[E,EI]{}
	fromHash:=vertexHash(vw.Hash(from))
	toHash:=vertexHash(vw.Hash(to))
	eHash:=edgeHash(ew.Hash(e))

	if _,ok:=g.vertices[fromHash]; !ok {
		return getVertexError[V](from)
	}
	if _,ok:=g.vertices[toHash]; !ok {
		return getVertexError[V](to)
	}
	if _,ok:=g.edges[eHash]; !ok {
		return getEdgeError[E](e)
	}

	gNode,_:=g.graph[fromHash]
	linkExists:=false
	for i:=0; i<len(gNode) && !linkExists; i++ {
		linkExists=(gNode[i].A==eHash && gNode[i].B==toHash)
	}
	if linkExists {
		return nil
	}

	g.numLinks++
	gNode = append(gNode, graphEdge{eHash, toHash})
	g.graph[fromHash]=gNode
	if gNode,ok:=g.graph[toHash]; !ok {
		g.graph[toHash]=gNode
	}
	return nil
}

// Description: Places a write lock on the underlying hash graph and then adds
// the link to the graph. Exhibits the same behavior as the non-synced
// [HashGraph.LinkPntr] method.
//
// Lock Type: Write
//
// Time Complexity: O(n), where n=num of outgoing edges from the start vertex
func (g *SyncedHashGraph[V, E, VI, EI])LinkPntr(from *V, to *V, e *E) error {
	g.Lock()
	defer g.Unlock()
	return g.HashGraph.LinkPntr(from,to,e)
}

func (g *HashGraph[V,E,VI,EI])DeleteVertex(v V) error {
	return nil
}

func (g *HashGraph[V,E,VI,EI])DeleteVertexPntr(v *V) error {
	return nil
}

func (g *HashGraph[V,E,VI,EI])DeleteEdge(e E) error {
	return nil
}

func (g *HashGraph[V, E, VI, EI])DeleteEdgePntr(e *E) error {
	return nil
}

func (g *HashGraph[V,E,VI,EI])DeleteLink(from V, to V, e E) error {
	return nil
}

func (g *HashGraph[V,E,VI,EI])DeleteLinkPntr(from *V, to *V, e *E) error {
	return nil
}

func (g *HashGraph[V,E,VI,EI])DeleteLinks(from V, to V) error {
	return nil
}

func (g *HashGraph[V,E,VI,EI])DeleteLinksPntr(from *V, to *V) error {
	return nil
}

func (g *HashGraph[V,E,VI,EI])Clear() {}

func (g *HashGraph[V,E,VI,EI])KeyedEq(
	other dynamicContainers.ReadGraph[V,E],
) bool {
	return false
}
func (g *HashGraph[V,E,VI,EI])UnorderedEq(
	other dynamicContainers.ReadGraph[V,E],
) bool {
	return false
}
func (g *HashGraph[V,E,VI,EI])Intersection(
	l dynamicContainers.ReadGraph[V,E],
	r dynamicContainers.ReadGraph[V,E],
) {}
func (g *HashGraph[V,E,VI,EI])Union(
	l dynamicContainers.ReadGraph[V,E],
	r dynamicContainers.ReadGraph[V,E],
) {}
func (g *HashGraph[V,E,VI,EI])Difference(
	l dynamicContainers.ReadGraph[V,E],
	r dynamicContainers.ReadGraph[V,E],
) {}
func (g *HashGraph[V,E,VI,EI])IsSuperset(
	other dynamicContainers.ReadGraph[V,E],
) bool {
	return false
}
func (g *HashGraph[V,E,VI,EI])IsSubset(
	other dynamicContainers.ReadGraph[V,E],
) bool {
	return false
}

func (_ *HashGraph[V,E,VI,EI])Eq(
	l *HashGraph[V,E,VI,EI],
	r *HashGraph[V,E,VI,EI],
) bool {
	return false
}
func (_ *HashGraph[V,E,VI,EI])Lt(
	l *HashGraph[V,E,VI,EI],
	r *HashGraph[V,E,VI,EI],
) bool {
	return false
}
func (_ *HashGraph[V,E,VI,EI])Hash(other *HashGraph[V,E,VI,EI]) hash.Hash {
	return hash.Hash(0)
}
func (_ *HashGraph[V,E,VI,EI])Zero(other *HashGraph[V,E,VI,EI]) {}
