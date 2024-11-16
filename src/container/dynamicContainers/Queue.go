package dynamicContainers

import "github.com/barbell-math/util/src/container/containerTypes"

// An interface that only allows read operations on a queue.
type ReadQueue[V any] interface {
	containerTypes.RWSyncable
	containerTypes.Addressable
	containerTypes.Length
	containerTypes.Capacity
	containerTypes.FirstElemRead[V]
}

// An interface that only allows write operations on a queue.
type WriteQueue[V any] interface {
	containerTypes.RWSyncable
	containerTypes.Clear
	containerTypes.Length
	containerTypes.Capacity
	containerTypes.LastElemWrite[V]
	containerTypes.FirstElemDelete[V]
}

// An interface that represents a queue with no restrictions on reading or
// writing.
type Queue[V any] interface {
	ReadQueue[V]
	WriteQueue[V]
}
