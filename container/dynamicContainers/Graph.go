package dynamicContainers

import (
	"github.com/barbell-math/util/container/containerTypes"
)


// An interface that only allows read operations on a graph. All graphs share
// the same read interface, regardless of subtype.
type ReadGraph[V any, E any] interface {
	containerTypes.RWSyncable
	containerTypes.Addressable
	containerTypes.ReadGraphOps[V,E]
	containerTypes.Comparisons[ReadGraph[V,E], V, E]
}

// An interface that only allows write operations on a directed graph.
type WriteDirectedGraph[V any, E any] interface {
	containerTypes.RWSyncable
	containerTypes.Clear
	containerTypes.WriteGraphOps[V,E]
	containerTypes.WriteDirectedGraphOps[V,E]
	containerTypes.DeleteGraphOps[V,E]
	containerTypes.DeleteDirectedGraphOps[V,E]
}
// An interface that only allows write operations on an undirected graph.
type WriteUndirectedGraph[V any, E any] interface {
	containerTypes.RWSyncable
	containerTypes.Clear
	containerTypes.WriteGraphOps[V,E]
	containerTypes.WriteUndirectedGraphOps[V,E]
	containerTypes.DeleteGraphOps[V,E]
	containerTypes.DeleteUndirectedGraphOps[V,E]
}

// An interface that represents a directed graph with no restrictions on reading
// or writing.
type DirectedGraph[V any, E any] interface {
	ReadGraph[V,E]
	WriteDirectedGraph[V,E]
}

// An interface that represents an undirected graph with no restrictions on
// reading or writing.
type UndirectedGraph[V any, E any] interface {
	ReadGraph[V,E]
	WriteDirectedGraph[V,E]
}
