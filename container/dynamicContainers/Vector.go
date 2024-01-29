package dynamicContainers

import "github.com/barbell-math/util/container/containerTypes"

// An interface that only allows read operations on a vector.
type ReadVector[K ~int, V any] interface {
	containerTypes.RWSyncable
	containerTypes.Length
	containerTypes.Capacity
	containerTypes.ReadOps[K,V]
	containerTypes.ReadKeyedOps[K,V]
}
// An interface that only allows write operations on a vector.
type WriteVector[K ~int, V any] interface {
	containerTypes.RWSyncable
	containerTypes.Clear
	containerTypes.Length
	containerTypes.Capacity
	containerTypes.WriteOps[K,V]
	containerTypes.WriteKeyedOps[K,V]
	containerTypes.DeleteOps[K,V]
	containerTypes.DeleteKeyedOps[K,V]
}
// An interface that represents a vector with no restrictions on reading or
// writing.
type Vector[K ~int, V any] interface {
	ReadVector[K,V]
	WriteVector[K,V]
}
