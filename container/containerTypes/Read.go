package containerTypes

import (
	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/container/basic"
)

// An interface that enforces implementation of read-only, value-only, operations.
type ReadOps[V any] interface {
	Vals() iter.Iter[V]
	ValPntrs() iter.Iter[*V]
	Contains(v V) bool
	ContainsPntr(v *V) bool
}

// An interface that enforces implementation of read-only, key/value, operations.
type ReadKeyedOps[K any, V any] interface {
	Get(k K) (V, error)
	GetPntr(k K) (*V, error)
	KeyOf(v V) (K, bool)
	Keys() iter.Iter[K]
}

// An interface that enforces the implementation of read-only first element access.
type FirstElemRead[V any] interface {
	PeekFront() (V, error)
	PeekPntrFront() (*V, error)
}

// An interface that enforces the implementation of read-only last element access.
type LastElemRead[V any] interface {
	PeekBack() (V, error)
	PeekPntrBack() (*V, error)
}

// An interface that enforces implementation of read-only, graph structure,
// opertions.
type ReadGraphOps[V any, E any] interface {
	NumLinks() int
	NumEdges() int
	NumVertices() int
	Edges() iter.Iter[E]
	EdgePntrs() iter.Iter[*E]
	Vertices() iter.Iter[V]
	VerticePntrs() iter.Iter[*V]
	ContainsEdge(e E) bool
	ContainsEdgePntr(e *E) bool
	ContainsVertex(v V) bool
	ContainsVertexPntr(v *V) bool
	OutEdges(v V) iter.Iter[E]
	OutEdgePntrs(v *V) iter.Iter[*E]
	NumOutLinks(v V) int
	NumOutLinksPntr(v *V) int
	OutVertices(v V) iter.Iter[V]
	OutVerticePntrs(v *V) iter.Iter[*V]
	OutEdgesAndVertices(v V) iter.Iter[basic.Pair[E, V]]
	OutEdgesAndVerticePntrs(v *V) iter.Iter[basic.Pair[*E, *V]]
	EdgesBetween(from V, to V) iter.Iter[E]
	EdgesBetweenPntr(from *V, to *V) iter.Iter[*E]
	ContainsLink(from V, to V, e E) bool
	ContainsLinkPntr(from *V, to *V, e *E) bool
}

// An interface that enforces implementation of read-only, directed, graph
// structure, operations.
type ReadDirectedGraphOps[V any, E any] interface {
	NumLinks() int
}

// An interface that enforces implementation of read-only, undirected, graph
// structure, operations.
type ReadUndirectedGraphOps[V any, E any] interface {
	NumUndirectedLinks() int
}
