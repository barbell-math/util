// This package serves to define the set of types that the containers will
// implement. These interfaces are meant to be composed together if necessary
// to allow for very specific sub-types of containers to be specified.
package containerTypes

import (
	"math"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/container/basic"
)


const (
    // A constant that may be passed to the pop function of the [DeleteOps] 
    // interface to specify popping all values.
    PopAll int=math.MaxInt
)

// An interface that exposes the RWMutex's interface as well as a convience
// function that will be used to determine if something is synced or if it
// simply implements this interface as a pass through.
type RWSyncable interface {
    Lock()
    RLock()
    Unlock()
    RUnlock()
    IsSynced() bool
}

// An interface that determines if a container is addressable or not.
type Addressable interface { IsAddressable() bool }

// An interface that allows access to the containers length.
type Length interface { Length() int }

// An interface that allows access to the containers capacity.
type Capacity interface { Capacity() int }

// An iterface that allows a container to be cleared, deleting all values.
type Clear interface { Clear() }

// An interface that defines what it means to have static capacity. Note that
// there is no DynamicCapacity interface, meaning that all containers are
// considered to be dynamic by default and are only static if explicity stated
// so by the implementation of this interface.
type StaticCapacity interface {
    Capacity
    Full() bool
}

// An interface that defines what value-only comparisons can be performed on a
// container.
type Comparisons[K any, V any] interface {
    UnorderedEq(other ComparisonsOtherConstraint[V]) bool
    Intersection(l ComparisonsOtherConstraint[V], r ComparisonsOtherConstraint[V])
    Union(l ComparisonsOtherConstraint[V], r ComparisonsOtherConstraint[V])
    Difference(l ComparisonsOtherConstraint[V], r ComparisonsOtherConstraint[V])
    IsSuperset(other ComparisonsOtherConstraint[V]) bool
    IsSubset(other ComparisonsOtherConstraint[V]) bool
}

// An interface that defines what key/value comparisons can be performed on a
// container that has keyed values.
type KeyedComparisons[K any, V any] interface {
    KeyedEq(other KeyedComparisonsOtherConstraint[K,V]) bool
}

// An interface that defines what kinds values can be passed to the methods in
// the [Comparisons] interface.
type ComparisonsOtherConstraint[V any] interface {
    RWSyncable
    Addressable
    Length
    ReadOps[V]
}

// An interface that defines what kinds values can be passed to the methods in
// the [KeyedComparisons] interface.
type KeyedComparisonsOtherConstraint[K any, V any] interface {
    RWSyncable
    Addressable
    Length
    ReadOps[V] // TODO - needed??
    ReadKeyedOps[K,V]
}

// An interface that enforces implementation of read-only, value-only, operations.
type ReadOps[V any] interface {
    Vals() iter.Iter[V]
    ValPntrs() iter.Iter[*V]
    Contains(v V) bool
    ContainsPntr(v *V) bool
}

// An interface that enforces implementation of read-only, key/value, operations.
type ReadKeyedOps[K any, V any] interface {
    Get(k K) (V,error)
    GetPntr(k K) (*V,error)
    KeyOf(v V) (K,bool)
    Keys() iter.Iter[K]
}

// An interface that enforces implementation of write-only, value-only, unique 
// valued, operations.
type WriteUniqueOps[K any, V any] interface {
    AppendUnique(vals ...V) error
}

// An interface that enforces implementation of write-only, key/value, unique 
// valued, operations.
type WriteUniqueKeyedOps[K any, V any] interface {
    EmplaceUnique(idx K, v V) error
}

// TODO
// make set
// Add arith to widgets - map will need it for sequential ops, separate widget type?? Probably should be
// ZeroVal() T
// UnitVal() T
// Map interface tests
// make map
// reimpl circular buffer
// make static container tests
// add window producer it iterface file

// An interface that enforces implementation of write-only, value-only, operations.
type WriteOps[V any] interface {
    Append(vals ...V) error
}
// An interface that enforces implementation of write-only, key/value, operations.
type WriteKeyedOps[K any, V any] interface {
    Set(kvPairs ...basic.Pair[K,V]) error;
    SetSequential(k K, v ...V) error;
}
// An interface that enforces implementation of write-only, key/value, 
// dynamic key, operations. A dynamic key operation is an operation that allows
// changing the value of a key but also allows changing of the keys as a
// result of that operation.
type WriteDynKeyedOps[K any, V any] interface {
    Insert(kvPairs ...basic.Pair[K,V]) error
    InsertSequential(idx K, v ...V) error
}
// An interface that enforces implementation of write-only, key/value, 
// static key, operations. A static key operation is an operation that allows
// changing the value of a key but does not allow changing of the keys as a
// result of that operation.
type WriteStaticKeyedOps[K any, V any] interface {
    Emplace(kvPairs ...basic.Pair[K,V]) error;
    EmplaceSequential(idk K, v ...V) error;
}

// An interface that enforces implementation of delete-only, value-only, operations.
type DeleteOps[K any, V any] interface {
    Pop(v V, num int) int
}
// An interface that enforces implementation of delete-only, key/value, operations.
type DeleteKeyedOps[K any, V any] interface {
    DeleteOps[K,V]
    Delete(idx K) error
}

// An interface that enforces the implementation of read-only first element access.
type FirstElemRead[V any] interface {
    PeekFront() (V,error);
    PeekPntrFront() (*V,error);
}
// An interface that enforces the implementation of write-only first element access.
type FirstElemWrite[V any] interface {
    PushFront(v ...V) error;
    ForcePushFront(v ...V)
}
// An interface that enforces the implementation of delete-only first element access.
type FirstElemDelete[V any] interface {
    PopFront() (V,error);
}

// An interface that enforces the implementation of read-only last element access.
type LastElemRead[V any] interface {
    PeekBack() (V,error);
    PeekPntrBack() (*V,error)
}
type LastElemWrite[V any] interface {
    PushBack(v ...V) (error);
    ForcePushBack(v ...V)
}
// An interface that enforces the implementation of delete-only last element access.
type LastElemDelete[V any] interface {
    PopBack() (V,error);
}
