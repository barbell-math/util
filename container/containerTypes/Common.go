// This package serves to define the set of types that the containers will
// implement. These interfaces are meant to be composed together if necessary
// to allow for very specific sub-types of containers to be specified.
package containerTypes

import "math"


const (
    // A constant that may be passed to the pop function of the [DeleteOps] 
    // interface to specify popping all values.
    PopAll int=math.MaxInt
)

// The interface that the RWMutex exposes.
type RWSyncable interface {
    Lock()
    Unlock()
    RLock()
    RUnlock()
}

// An interface that allows access to the containers length.
type Length interface { Length() int }

// An interface that allows access to the containers capacity.
type Capacity interface { Capacity() int }

// An interface that defines what it means to have static capacity. Note that
// there is no DynamicCapacity interface, meaning that all containers are
// considered to be dynamic by default and are only static if explicity stated
// so by the implementation of this interface.
type StaticCapacity interface {
    Capacity
    Full() bool
}

// An interface the enforces implementation of read-only, value-only, operations.
type ReadOps[K any, V any] interface {
    Contains(v V) bool
}
// An interface the enforces implementation of read-only, key/value, operations.
type ReadKeyedOps[K any, V any] interface {
    ReadOps[K,V]
    Get(k K) (V,error)
    GetPntr(k K) (*V,error)
    KeyOf(v V) (K,bool)
}

// An interface the enforces implementation of write-only, value-only, operations.
type WriteOps[K any, V any] interface {
    Append(vals ...V) error
}
// An interface the enforces implementation of write-only, key/value, operations.
type WriteKeyedOps[K any, V any] interface {
    WriteOps[K,V]
    Emplace(idx K, v V) error;
    Push(idx K, v ...V) error;
}

// An interface the enforces implementation of delete-only, value-only, operations.
type DeleteOps[K any, V any] interface {
    Pop(v V, num int) int
}
// An interface the enforces implementation of delete-only, key/value, operations.
type DeleteKeyedOps[K any, V any] interface {
    DeleteOps[K,V]
    Delete(idx K) error
    Clear()
}

// An interface the enforces the implementation of read-only first element access.
type FirstElemRead[V any] interface {
    PeekFront() (V,error);
    PeekPntrFront() (*V,error);
}
// An interface the enforces the implementation of write-only first element access.
type FirstElemWrite[V any] interface {
    PushFront(v V) error;
    ForcePushFront(v V)
}
// An interface the enforces the implementation of delete-only first element access.
type FirstElemDelete[V any] interface {
    PopFront() (V,error);
}

// An interface the enforces the implementation of read-only last element access.
type LastElemRead[V any] interface {
    PeekBack() (V,error);
    PeekPntrBack() (*V,error);
}
// An interface the enforces the implementation of write-only last element access.
type LastElemWrite[V any] interface {
    PushBack(v V) (error);
    ForcePushBack(v V)
}
// An interface the enforces the implementation of delete-only last element access.
type LastElemDelete[V any] interface {
    PopBack() (V,error);
}
