package containers

import (
	"sync"

	"github.com/barbell-math/util/algo/hash"
	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/algo/widgets"
	"github.com/barbell-math/util/container/containerTypes"
)


type (
    internalHashSetImpl[T any] map[uint64]T

    // A type to represent a set that dynamically grows as elements are added.
    // The set will maintain uniqueness and is internally implemented with a 
    // hashing method. The type constraints on the generics define the logic 
    // for how value specific operations, such as equality comparisons, will be 
    // handled.
    HashSet[T any, U widgets.WidgetInterface[T]] struct {
        internalHashSetImpl[T]
    }

    // A synchronized version of HashSet. All operations will be wrapped in the
    // appropriate calls the embedded RWMutex. A pointer to a RWMutex is embedded
    // rather than a value to avoid copying the lock value.
    SyncedHashSet[T any, U widgets.WidgetInterface[T]] struct {
    	*sync.RWMutex
    	HashSet[T,U]
    }
)

// Creates a new hash set initialized with enough memory to hold size elements. 
// Size must be >= 0, an error will be returned if it is not. If size is 0 the 
// hash set will be initialized with 0 elements.
func NewHashSet[
    T any,
    U widgets.WidgetInterface[T],
](size int) (HashSet[T,U],error) {
    if size<0 {
        return HashSet[T, U]{}, getSizeError(size)
    }
    return HashSet[T, U]{
        internalHashSetImpl: make(internalHashSetImpl[T], size),
    }, nil
}

// Creates a new synced hash set initialized with enough memory to hold size 
// elements. Size must be >= 0, an error will be returned if it is not. If size 
// is 0 the hash set will be initialized with 0 elements. The underlying RWMutex 
// value will be fully unlocked upon initialization.
func NewSyncedHashSet[T any, U widgets.WidgetInterface[T]](
    size int,
) (SyncedHashSet[T,U],error) {
    rv,err:=NewHashSet[T,U](size)
    return SyncedHashSet[T,U]{
	RWMutex: &sync.RWMutex{},
        HashSet: rv,
    },err
}

// A empty pass through function that performs no action. HashSet will call all 
// the appropriate locking methods despite not being synced, just nothing will
// happen. This is done so that SyncedHashSet can simply embed a HashSet and
// override the appropriate locking methods to implement the correct behavior
// without needing to make any additional changes such as wrapping every single
// method from HashSet.
func (h *HashSet[T,U])Lock() { }

// A empty pass through function that performs no action. HashSet will call all 
// the appropriate locking methods despite not being synced, just nothing will
// happen. This is done so that SyncedHashSet can simply embed a HashSet and
// override the appropriate locking methods to implement the correct behavior
// without needing to make any additional changes such as wrapping every single
// method from HashSet.
func (h *HashSet[T,U])Unlock() { }

// A empty pass through function that performs no action. HashSet will call all 
// the appropriate locking methods despite not being synced, just nothing will
// happen. This is done so that SyncedHashSet can simply embed a HashSet and
// override the appropriate locking methods to implement the correct behavior
// without needing to make any additional changes such as wrapping every single
// method from HashSet.
func (h *HashSet[T,U])RLock() { }

// A empty pass through function that performs no action. HashSet will call all 
// the appropriate locking methods despite not being synced, just nothing will
// happen. This is done so that SyncedHashSet can simply embed a HashSet and
// override the appropriate locking methods to implement the correct behavior
// without needing to make any additional changes such as wrapping every single
// method from HashSet.
func (h *HashSet[T,U])RUnlock() { }

// The SyncedHashSet method to override the HashSet pass through function and 
// actually apply the mutex operation.
func (h *SyncedHashSet[T,U])Lock() { h.RWMutex.Lock() }

// The SyncedHashSet method to override the HashSet pass through function and 
// actually apply the mutex operation.
func (h *SyncedHashSet[T,U])Unlock() { h.RWMutex.Unlock() }

// The SyncedHashSet method to override the HashSet pass through function and 
// actually apply the mutex operation.
func (h *SyncedHashSet[T,U])RLock() { h.RWMutex.RLock() }

// The SyncedHashSet method to override the HashSet pass through function and 
// actually apply the mutex operation.
func (h *SyncedHashSet[T,U])RUnlock() { h.RWMutex.RUnlock() }

// Returns the length of the vector.
func (h *HashSet[T, U])Length() int {
    h.RLock()
    defer h.RUnlock()
    return len(h.internalHashSetImpl)
}

func (h *HashSet[T,U])Vals() iter.Iter[T] {
    return iter.NoElem[T]()
}

func (h *HashSet[T,U])ValPntrs() iter.Iter[*T] {
    return iter.NoElem[*T]()
}

func (h *HashSet[T,U])Contains(v T) bool {
    return false
}

func (h *HashSet[T,U])ContainsPntr(v *T) bool {
    return false
}

func (h *HashSet[T, U])AppendUnique(vals ...T) error {
    return nil
}

func (h *HashSet[T, U])Pop(v T, num int) int {
    return 0
}

func (h *HashSet[T, U])Clear() {

}

func (h *HashSet[T,U])UnorderedEq(
    other containerTypes.ComparisonsOtherConstraint[T],
) bool {
    return false
}

func (h *HashSet[T,U])Intersection(
    l containerTypes.ComparisonsOtherConstraint[T],
    r containerTypes.ComparisonsOtherConstraint[T],
) {

}

func (h *HashSet[T,U])Union(
    l containerTypes.ComparisonsOtherConstraint[T],
    r containerTypes.ComparisonsOtherConstraint[T],
) {

}

func (h *HashSet[T,U])Difference(
    l containerTypes.ComparisonsOtherConstraint[T],
    r containerTypes.ComparisonsOtherConstraint[T],
) {

}

func (h *HashSet[T,U])IsSuperset(
    other containerTypes.ComparisonsOtherConstraint[T],
) bool {
    return false
}

func (h *HashSet[T,U])IsSubset(
    other containerTypes.ComparisonsOtherConstraint[T],
) bool {
    return false
}

func (v *Vector[T, U])Eq(l *Vector[T,U], r *Vector[T,U]) bool {
    return false
}

func (v *Vector[T, U])Lt(l *Vector[T,U], r *Vector[T,U]) bool {
    return false
}

func (c *Vector[T, U])Hash(other *Vector[T,U]) hash.Hash {
    return hash.Hash(0)
}

func (v *Vector[T, U])Zero(other *Vector[T,U]) {

}
