// This package serves to define the set of types that the containers will
// implement. These interfaces are meant to be composed together if necessary
// to allow for very specific sub-types of containers to be specified.
package containerTypes

import (
	"math"

	"github.com/barbell-math/util/algo/iter"
)


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
    ReadOps[V]
    RWSyncable
    Length
}

// An interface that defines what kinds values can be passed to the methods in
// the [KeyedComparisons] interface.
type KeyedComparisonsOtherConstraint[K any, V any] interface {
    ReadKeyedOps[K,V]
    ReadOps[V] // TODO - needed??
    RWSyncable
    Length
}

// TODO
// Add copy?? - would need to add widget support
//  Copy(v V) OI

// An interface the enforces implementation of read-only, value-only, operations.
type ReadOps[V any] interface {
    Vals() iter.Iter[V]
    ValPntrs() iter.Iter[*V]
    Contains(v V) bool
    ContainsPntr(v *V) bool
}
// An interface the enforces implementation of read-only, key/value, operations.
type ReadKeyedOps[K any, V any] interface {
    Get(k K) (V,error)
    GetPntr(k K) (*V,error)
    KeyOf(v V) (K,bool)
    Keys() iter.Iter[K]
    KeyPntrs() iter.Iter[*K]
    // KeyRange() func(k K) bool
    // KeyPntrRange() func(k *K) bool
    // KeyValRange() func(k K, v V) bool
    // KeyValPntrRange() func(k *K, v *V) bool
}

// An interface the enforces implementation of write-only, value-only, unique 
// valued, operations.
type WriteUniqueOps[K any, V any] interface {
    AppendUnique(vals ...V) error
}

// An interface the enforces implementation of write-only, key/value, unique 
// valued, operations.
type WriteUniqueKeyedOps[K any, V any] interface {
    EmplaceUnique(idx K, v V) error
}

// An interface the enforces implementation of write-only, value-only, operations.
type WriteOps[V any] interface {
    Append(vals ...V) error
}
// An interface the enforces implementation of write-only, key/value, operations.
type WriteKeyedOps[K any, V any] interface {
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
}

// An interface the enforces the implementation of read-only first element access.
type FirstElemRead[V any] interface {
    PeekFront() (V,error);
    PeekPntrFront() (*V,error);
}
// An interface the enforces the implementation of write-only first element access.
type FirstElemWrite[V any] interface {
    PushFront(v ...V) error;
    ForcePushFront(v ...V)
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
    PushBack(v ...V) (error);
    ForcePushBack(v ...V)
}
// An interface the enforces the implementation of delete-only last element access.
type LastElemDelete[V any] interface {
    PopBack() (V,error);
}
