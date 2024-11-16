package dynamicContainers

import "github.com/barbell-math/util/src/container/containerTypes"

// An interface that only allows read operations on a stack.
type ReadStack[V any] interface {
	containerTypes.RWSyncable
	containerTypes.Addressable
	containerTypes.Length
	containerTypes.Capacity
	containerTypes.LastElemRead[V]
}

// An interface that only allows write operations on a stack.
type WriteStack[V any] interface {
	containerTypes.RWSyncable
	containerTypes.Clear
	containerTypes.Length
	containerTypes.Capacity
	containerTypes.LastElemWrite[V]
	containerTypes.LastElemDelete[V]
}

// An interface that represents a stack with no restrictions on reading or
// writing.
type Stack[V any] interface {
	ReadStack[V]
	WriteStack[V]
}
