package dynamicContainers

import (
	"github.com/barbell-math/util/container/containerTypes"
)

type ReadGraph[V any, E any] interface {
	containerTypes.RWSyncable
	containerTypes.Addressable
	containerTypes.ReadGraphOps[V,E]
	// containerTypes.Comparisons[] - TODO - have to wait for compiler bug to be fixed :(
}
type WriteGraph[V any, E any] interface {
	containerTypes.RWSyncable
	containerTypes.Clear
	containerTypes.WriteGraphOps[V,E]
	containerTypes.DeleteGraphOps[V,E]
}
type Graph[V any, E any] interface {
	ReadGraph[V,E]
	WriteGraph[V,E]
}
