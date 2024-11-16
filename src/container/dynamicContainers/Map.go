package dynamicContainers

import "github.com/barbell-math/util/src/container/containerTypes"

// An interface that only allows read operations on a map.
type ReadMap[K any, V any] interface {
	containerTypes.RWSyncable
	containerTypes.Addressable
	containerTypes.Length
	containerTypes.ReadOps[V]
	containerTypes.ReadKeyedOps[K, V]
	containerTypes.KeyedComparisons[
		containerTypes.KeyedComparisonsOtherConstraint[K, V],
		K,
		V,
	]
}

// An interface that only allows write operations on a map.
type WriteMap[K any, V any] interface {
	containerTypes.RWSyncable
	containerTypes.Clear
	containerTypes.Length
	containerTypes.WriteKeyedOps[K, V]
	containerTypes.WriteStaticKeyedOps[K, V]
	containerTypes.DeleteOps[K, V]
	containerTypes.DeleteKeyedOps[K, V]
}

// An interface that represents a map with no restrictions on reading or
// writing.
type Map[K any, V any] interface {
	ReadMap[K, V]
	WriteMap[K, V]
}
