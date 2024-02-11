package containers

import (
	"sync"

	"github.com/barbell-math/util/algo/hash"
	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/algo/widgets"
	"github.com/barbell-math/util/container/containerTypes"
)


type (
    internalHashSetImpl[T any] map[hash.Hash]T

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

// Returns the number of values in the hash set.
func (h *HashSet[T, U])Length() int {
    h.RLock()
    defer h.RUnlock()
    return len(h.internalHashSetImpl)
}

// Returns an iterator that iterates over the values in the hash set. The hash 
// set will have a read lock the entire time the iteration is being performed. 
// The lock will not be applied until the iterator is consumed.
func (h *HashSet[T,U])Vals() iter.Iter[T] {
    return iter.MapVals[hash.Hash,T](h.internalHashSetImpl).SetupTeardown(
        func() error { h.RLock(); return nil },
        func() error { h.RUnlock(); return nil },
    )
}

// Returns an iterator that iterates over the pointers to the values in the 
// hash set. The hash set will have a read lock the entire time the iteration is 
// being performed. The lock will not be applied until the iterator is consumed.
func (h *HashSet[T,U])ValPntrs() iter.Iter[*T] {
    return iter.Map[T,*T](
        iter.MapVals[hash.Hash,T](h.internalHashSetImpl).SetupTeardown(
            func() error { h.RLock(); return nil },
            func() error { h.RUnlock(); return nil },
        ),
        func(index int, val T) (*T, error) { return &val,nil },
    )
}

// Contains will return true if the supplied value is in the hash set, false
// otherwise. All equality comparisons are performed by the generic U widget
// type that the hash set was initialized with. The time complexity of Contains
// on a hash set is O(1).
func (h *HashSet[T,U])Contains(v T) bool {
    return h.ContainsPntr(&v)
}

// ContainsPntr will return true if the supplied value is in the hash set, false
// otherwise. All equality comparisons are performed by the generic U widget
// type that the hash set was initialized with. The time complexity of 
// ContainsPntr on a hash set is O(1).
func (h *HashSet[T,U])ContainsPntr(v *T) bool {
    h.RLock()
    defer h.RUnlock()
    rv:=false
    w:=widgets.NewWidget[T,U]()
    vHash:=w.Hash(v)
    for i:=0; !rv; i++ {
        hashPlacement:=vHash+hash.Hash(i)
        if iterV,foundPlace:=h.internalHashSetImpl[hashPlacement]; foundPlace {
            rv=w.Eq(v,&iterV)
        } else {
            break
        }
    }
    return rv
}

// AppendUnique will append the supplied values to the hash set if they are not
// already present in the hash set (unique). Non-unique values will not be 
// appended. This function will never return an error. The time complexity of 
// AppendUnique is O(m) where m is the number of values to append.
func (h *HashSet[T, U])AppendUnique(vals ...T) error {
    h.Lock()
    defer h.Unlock()
    for i:=0; i<len(vals); i++ {
        h.appendOp(&vals[i])
    }
    return nil
}

// Note - this function assumes the appropriate locks have already been placed
// on the set.
func (h *HashSet[T, U])appendOp(v *T) {
    w:=widgets.NewWidget[T,U]()
    valHash:=w.Hash(v)
    for j:=0; ; j++ {
        hashPlacement:=valHash+hash.Hash(j)
        if iterV,found:=h.internalHashSetImpl[hashPlacement]; !found {
            h.internalHashSetImpl[hashPlacement]=*v
            break
        } else if w.Eq(v,&iterV) {
            break
        }
    }
}

// Pop will remove the val from the hash aset. All equality comparisons are 
// performed by the generic U widget type that the hash set was initialized 
// with. If num is <=0 then no values will be poped and the hash set
// will not change. Any values for num >=1 will result in the same behavior of
// the specified value being deleted from the set (because a set will only ever
// contain a single value).
func (h *HashSet[T, U])Pop(v T, num int) int {
    if num<=0 {
        return 0
    }
    h.RLock()
    defer h.RUnlock()
    w:=widgets.NewWidget[T,U]()
    vHash:=w.Hash(&v)
    for i:=0; ; i++ {
        hashPlacement:=vHash+hash.Hash(i)
        if iterV,foundPlace:=h.internalHashSetImpl[hashPlacement]; foundPlace {
            if found:=w.Eq(&v,&iterV); found {
                delete(h.internalHashSetImpl,hashPlacement)    
                return 1
            }
        } else {
            return 0
        }
    }
}

// Clears all values from the hash set.
func (h *HashSet[T, U])Clear() {
    h.Lock()
    defer h.RUnlock()
    h.internalHashSetImpl=make(internalHashSetImpl[T])
}

// Returns true if the elements in h are all contained in other and the elements
// of other are all contained in h, regardless of position. Returns false 
// otherwise. This implementation of UnorderedEq is dependent on the time 
// complexity of the implementation of the ContainsPntr method on other. In 
// big-O it might look something like this, O(n*O(other.ContainsPntr))), where n 
// is the number of elements in h and O(other.ContainsPntr) represents the 
// time complexity of the containsPntr method on other with m values. Read locks 
// will be placed on both this hash set and the other hash set.
func (h *HashSet[T,U])UnorderedEq(
    other containerTypes.ComparisonsOtherConstraint[T],
) bool {
    h.RLock()
    other.RLock()
    defer h.RUnlock()
    defer other.RUnlock()
    if len(h.internalHashSetImpl)!=other.Length() {
        return false
    }
    for _,iterV:=range(h.internalHashSetImpl) {
        if !other.ContainsPntr(&iterV) {
            return false
        }
    }
    return true
}

func (h *HashSet[T,U])Intersection(
    l containerTypes.ComparisonsOtherConstraint[T],
    r containerTypes.ComparisonsOtherConstraint[T],
) {
    h.Lock()
    l.RLock()
    r.RLock()
    defer h.Unlock()
    defer l.RUnlock()
    defer r.RUnlock()
    h.internalHashSetImpl=make(internalHashSetImpl[T],(l.Length()+r.Length())/2)
    r.Vals().ForEach(func(index int, val T) (iter.IteratorFeedback, error) {
        if l.ContainsPntr(&val) {
            h.appendOp(&val)
        }
        return iter.Continue,nil
    })
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

func (h *HashSet[T, U])Eq(l *HashSet[T,U], r *HashSet[T,U]) bool {
    return false
}

func (h *HashSet[T, U])Lt(l *HashSet[T,U], r *HashSet[T,U]) bool {
    return false
}

func (h *HashSet[T, U])Hash(other *HashSet[T,U]) hash.Hash {
    return hash.Hash(0)
}

func (h *HashSet[T, U])Zero(other *HashSet[T,U]) {

}
