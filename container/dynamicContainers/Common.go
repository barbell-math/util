// This package serves to define the set of dynamic containers and expose them
// as interfaces.
package dynamicContainers

import "github.com/barbell-math/util/container/containerTypes"

// An interface that only allows read operations on a vector.
type ReadVector[K ~int | ~int8 | ~int16 | ~int32 | ~int64, V any] interface {
	containerTypes.RWSyncable
	containerTypes.Length
	containerTypes.Capacity
	containerTypes.ReadOps[K,V]
	containerTypes.ReadKeyedOps[K,V]
	containerTypes.LastElemRead[V]
	containerTypes.FirstElemRead[V]
}
// An interface that only allows write operations on a vector.
type WriteVector[K ~int | ~int8 | ~int16 | ~int32 | ~int64, V any] interface {
	containerTypes.RWSyncable
	containerTypes.Length
	containerTypes.Capacity
	containerTypes.WriteOps[K,V]
	containerTypes.WriteKeyedOps[K,V]
	containerTypes.LastElemWrite[V]
	containerTypes.FirstElemWrite[V]
	containerTypes.DeleteOps[K,V]
	containerTypes.DeleteKeyedOps[K,V]
	containerTypes.LastElemDelete[V]
	containerTypes.FirstElemDelete[V]
}
// An interface that represents a vector with no restrictions on reading or
// writing.
type Vector[K ~int | ~int8 | ~int16 | ~int32 | ~int64, V any] interface {
	ReadVector[K,V]
	WriteVector[K,V]
}
