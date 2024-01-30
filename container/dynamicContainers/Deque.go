package dynamicContainers

import "github.com/barbell-math/util/container/containerTypes"

// An interface that only allows read operations on a vector.
type ReadDeque[V any] interface {
	containerTypes.RWSyncable
	containerTypes.Length
	containerTypes.Capacity
	containerTypes.LastElemRead[V]
	containerTypes.FirstElemRead[V]
}
// An interface that only allows write operations on a vector.
type WriteDeque[V any] interface {
	containerTypes.RWSyncable
	containerTypes.Clear
	containerTypes.Length
	containerTypes.Capacity
	containerTypes.LastElemWrite[V]
	containerTypes.FirstElemWrite[V]
	containerTypes.LastElemDelete[V]
	containerTypes.FirstElemDelete[V]
}
// An interface that represents a vector with no restrictions on reading or
// writing.
type Deque[V any] interface {
	ReadDeque[V]
	WriteDeque[V]
}