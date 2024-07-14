package containerTypes

// An interface that enforces implementation of delete-only, value-only, operations.
type DeleteOps[K any, V any] interface {
	Pop(v V) int
}

// An interface that enforces implementation of delete-only, value-only,
// sequential operations.
type DeleteSequentialOps[K any, V any] interface {
	PopSequential(v V, num int) int
}

// An interface that enforces implementation of delete-only, key/value, operations.
type DeleteKeyedOps[K any, V any] interface {
	Delete(idx K) error
}

// An interface that enforces implementation of delete-only, key/value, operations.
type DeleteKeyedSequentialOps[K any, V any] interface {
	DeleteSequential(start int, end int) error
}

// An interface that enforces the implementation of delete-only first element access.
type FirstElemDelete[V any] interface {
	PopFront() (V, error)
}

// An interface that enforces the implementation of delete-only last element access.
type LastElemDelete[V any] interface {
	PopBack() (V, error)
}

// An interface that enforces implementaiton of delete-only, graph structure,
// operations.
type DeleteGraphOps[V any, E any] interface {
	DeleteVertex(v V) error
	DeleteVertexPntr(v *V) error
	DeleteEdge(e E) error
	DeleteEdgePntr(e *E) error
}

// An interface that enforces implementaiton of delete-only, directed, graph
// structure, operations.
type DeleteDirectedGraphOps[V any, E any] interface {
	DeleteLink(from V, to V, e E) error
	DeleteLinkPntr(from *V, to *V, e *E) error
	DeleteLinks(from V, to V) error
	DeleteLinksPntr(from *V, to *V) error
}

// An interface that enforces implementaiton of delete-only, undirected, graph
// structure, operations.
type DeleteUndirectedGraphOps[V any, E any] interface {
	DeleteUndirectedLink(from V, to V, e E) error
	DeleteUndirectedLinkPntr(from *V, to *V, e *E) error
	DeleteUndirectedLinks(from V, to V) error
	DeleteUndirectedLinksPntr(from *V, to *V) error
}
