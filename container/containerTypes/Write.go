package containerTypes

import "github.com/barbell-math/util/container/basic"

// An interface that enforces implementation of write-only, value-only, unique
// valued, operations.
type WriteUniqueOps[K any, V any] interface {
    AppendUnique(vals ...V) error
}

// An interface that enforces implementation of write-only, key/value, unique 
// valued, operations.
type WriteUniqueKeyedOps[K any, V any] interface {
    EmplaceUnique(idx K, v V) error
}

// An interface that enforces implementation of write-only, value-only, operations.
type WriteOps[V any] interface {
    Append(vals ...V) error
}
// An interface that enforces implementation of write-only, key/value, operations.
type WriteKeyedOps[K any, V any] interface {
    Set(kvPairs ...basic.Pair[K,V]) error;
}
// An interface that enforces implementation of write-only, key/value, 
// sequential operations.
type WriteKeyedSequentialOps[K any, V any] interface {
    SetSequential(k K, v ...V) error;
}

// An interface that enforces implementation of write-only, key/value, 
// dynamic key, operations. A dynamic key operation is an operation that allows
// changing the value of a key but also allows changing of multiple keys as a
// result of that operation.
type WriteDynKeyedOps[K any, V any] interface {
    Insert(kvPairs ...basic.Pair[K,V]) error
    InsertSequential(idx K, v ...V) error
}
// An interface that enforces implementation of write-only, key/value, 
// static key, sequential, operations. A static key operation is an operation 
// that allows changing the value of a key but does not allow changing of
// multiple keys as a result of that operation.
type WriteStaticKeyedOps[K any, V any] interface {
    Emplace(kvPairs ...basic.Pair[K,V]) error;
}
// An interface that enforces implementation of write-only, key/value, 
// static key, operations. A static key operation is an operation that allows
// changing the value of a key but does not allow changing of multiple keys as a
// result of that operation.
type WriteStaticKeyedSequentialOps[K any, V any] interface {
    EmplaceSequential(idk K, v ...V) error;
}

// An interface that enforces the implementation of write-only first element access.
type FirstElemWrite[V any] interface {
    PushFront(v ...V) error;
    ForcePushFront(v ...V)
}

// An interface that enforces the implemntation of write-only last element access.
type LastElemWrite[V any] interface {
    PushBack(v ...V) (error);
    ForcePushBack(v ...V)
}

// An interface that enforces implementation of write-only, graph structure,
// operations.
type WriteGraphOps[V any, E any] interface {
    AddEdges(e ...E) error
    AddVertices(v ...V) error
}

// An interface that enforces implementaiton of write-only, undirected, graph
// structure operations.
type WriteUndirectedGraphOps[V any, E any] interface {
    LinkUndirected(from V, to V, e E) error
    LinkUndirectedPntr(from *V, to *V, e *E) error
}

// An interface that enforces implementaiton of write-only, directed, graph
// structure operations.
type WriteDirectedGraphOps[V any, E any] interface {
    Link(from V, to V, e E) error
    LinkPntr(from *V, to *V, e *E) error
}
