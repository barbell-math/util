// This package serves to define the set of types that the containers will
// implement. These interfaces are meant to be composed together if necessary
// to allow for very specific sub-types of containers to be specified.
package containerTypes

import "math"

const (
	// A constant that may be passed to the pop function of the [DeleteOps]
	// interface to specify popping all values.
	PopAll int = math.MaxInt
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
type Addressable interface{ IsAddressable() bool }

// An interface that allows access to the containers length.
type Length interface{ Length() int }

// An interface that allows access to the containers capacity.
type Capacity interface{ Capacity() int }

// An iterface that allows a container to be cleared, deleting all values.
type Clear interface{ Clear() }

// An interface that defines what it means to have static capacity. Note that
// there is no DynamicCapacity interface, meaning that all containers are
// considered to be dynamic by default and are only static if explicity stated
// so by the implementation of this interface.
type StaticCapacity interface {
	Capacity
	Full() bool
}

// An interface that defines what set-wise operations can be performed on a 
// container.
type SetOperations[RI any, V any] interface {
	Intersection(l RI, r RI)
	Union(l RI, r RI)
	Difference(l RI, r RI)
}

// An interface that defines what value-only comparisons can be performed on a
// container.
type Comparisons[RI any, V any] interface {
	UnorderedEq(other RI) bool
	IsSuperset(other RI) bool
	IsSubset(other RI) bool
}

// An interface that defines what key/value comparisons can be performed on a
// container that has keyed values.
type KeyedComparisons[RI any, K any, V any] interface {
	KeyedEq(other RI) bool
}

// An interface that defines what kinds of values can be passed to the methods
// in the [Comparisons] interface.
type ComparisonsOtherConstraint[V any] interface {
	RWSyncable
	Addressable
	Length
	ReadOps[V]
}

// An interface that defines what kinds of values can be passed to the methods
// in the [KeyedComparisons] interface.
type KeyedComparisonsOtherConstraint[K any, V any] interface {
	RWSyncable
	Addressable
	Length
	ReadKeyedOps[K, V]
}

// An interface that defines what kinds of values can be passed to the methods
// in the [Comparisons] and [KeyedComparisons] interfaces.
type GraphComparisonsConstraint[V any, E any] interface {
	RWSyncable
	Addressable
	ReadGraphOps[V, E]
}
