package dynamicContainers

import "github.com/barbell-math/util/container/containerTypes"

// An interface that only allows read operations on a vector.
type ReadVector[V any] interface {
	containerTypes.Capacity
	ReadMap[int,V]
}
// An interface that only allows write operations on a vector.
type WriteVector[V any] interface {
	containerTypes.Capacity
	containerTypes.WriteOps[V]
	WriteMap[int,V]
}
// An interface that represents a vector with no restrictions on reading or
// writing.
type Vector[V any] interface {
	ReadVector[V]
	WriteVector[V]
}
