package dynamicContainers

import "github.com/barbell-math/util/container/containerTypes"

// An interface that only allows read operations on a vector.
type ReadStack[K ~int, V any] interface {
	containerTypes.RWSyncable
	containerTypes.Length
	containerTypes.Capacity
	containerTypes.LastElemRead[V]
}
// An interface that only allows write operations on a vector.
type WriteStack[K ~int, V any] interface {
	containerTypes.RWSyncable
	containerTypes.Clear
	containerTypes.Length
	containerTypes.Capacity
	containerTypes.LastElemWrite[V]
	containerTypes.LastElemDelete[V]
}
// An interface that represents a vector with no restrictions on reading or
// writing.
type Stack[K ~int, V any] interface {
	ReadStack[K,V]
	WriteStack[K,V]
}
