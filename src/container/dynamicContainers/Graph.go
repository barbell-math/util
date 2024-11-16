package dynamicContainers

import (
	"github.com/barbell-math/util/src/container/containerTypes"
)

// An interface that only allows read operations in a directed graph.
type ReadDirectedGraph[V any, E any] interface {
	containerTypes.RWSyncable
	containerTypes.Addressable
	containerTypes.ReadGraphOps[V, E]
	containerTypes.Comparisons[
		containerTypes.GraphComparisonsConstraint[V, E],
		E,
	]
	containerTypes.KeyedComparisons[
		containerTypes.GraphComparisonsConstraint[V, E],
		V,
		E,
	]
}

// An interface that only allows write operations on a directed graph.
type WriteDirectedGraph[V any, E any] interface {
	containerTypes.RWSyncable
	containerTypes.Clear
	containerTypes.WriteGraphOps[V, E]
	containerTypes.DeleteGraphOps[V, E]
	containerTypes.DeleteDirectedGraphOps[V, E]
	containerTypes.SetOperations[
		containerTypes.GraphComparisonsConstraint[V, E],
		V,
	]
}

// An interface that represents a directed graph with no restrictions on reading
// or writing.
type DirectedGraph[V any, E any] interface {
	ReadDirectedGraph[V, E]
	WriteDirectedGraph[V, E]
}
