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
		HashSet[T, U]
	}
)

// Creates a new hash set initialized with enough memory to hold size elements.
// Size must be >= 0, an error will be returned if it is not. If size is 0 the
// hash set will be initialized with 0 elements.
func NewHashSet[
	T any,
	U widgets.WidgetInterface[T],
](size int) (HashSet[T, U], error) {
	if size < 0 {
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
) (SyncedHashSet[T, U], error) {
	rv, err := NewHashSet[T, U](size)
	return SyncedHashSet[T, U]{
		RWMutex: &sync.RWMutex{},
		HashSet: rv,
	}, err
}

// Converts the supplied hash set to a syncronized hash set. Beware: The original
// non-synced hash set will remain useable.
func (v *HashSet[T, U]) ToSynced() SyncedHashSet[T, U] {
	return SyncedHashSet[T, U]{
		RWMutex: &sync.RWMutex{},
		HashSet: *v,
	}
}

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (h *HashSet[T, U]) Lock() {}

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (h *HashSet[T, U]) Unlock() {}

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (h *HashSet[T, U]) RLock() {}

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (h *HashSet[T, U]) RUnlock() {}

// The SyncedHashSet method to override the HashSet pass through function and
// actually apply the mutex operation.
func (h *SyncedHashSet[T, U]) Lock() { h.RWMutex.Lock() }

// The SyncedHashSet method to override the HashSet pass through function and
// actually apply the mutex operation.
func (h *SyncedHashSet[T, U]) Unlock() { h.RWMutex.Unlock() }

// The SyncedHashSet method to override the HashSet pass through function and
// actually apply the mutex operation.
func (h *SyncedHashSet[T, U]) RLock() { h.RWMutex.RLock() }

// The SyncedHashSet method to override the HashSet pass through function and
// actually apply the mutex operation.
func (h *SyncedHashSet[T, U]) RUnlock() { h.RWMutex.RUnlock() }

// Returns false, hash sets are not addressable.
func (h *HashSet[T, U]) IsAddressable() bool { return false }

// Returns false, a hash set is not synced.
func (h *HashSet[T, U]) IsSynced() bool { return false }

// Returns true, a synced hash set is synced. :O
func (h *SyncedHashSet[T, U]) IsSynced() bool { return true }

// Description: Returns the number of values in the hash set.
//
// Time Complexity: O(1)
func (h *HashSet[T, U]) Length() int {
	return len(h.internalHashSetImpl)
}

// Description: Places a read lock on the underlying hash set and then calls the
// underlying hash set's [HashSet.Length] method.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (h *SyncedHashSet[T, U]) Length() int {
	h.RLock()
	defer h.RUnlock()
	return h.HashSet.Length()
}

// Description: Returns an iterator that iterates over the values in the hash
// set.
//
// Time Complexity: O(n)
func (h *HashSet[T, U]) Vals() iter.Iter[T] {
	return iter.MapVals[hash.Hash, T](h.internalHashSetImpl)
}

// Description: Modifies the iterator chain returned by the unerlying
// [hash set.Vals] method such that a read lock will be placed on the underlying
// hash set when iterator is consumer. The hash set will have a read lock the
// entire time the iteration is being performed. The lock will not be applied
// until the iterator starts to be consumed.
//
// Lock Type: Read
//
// Time Complexity: O(n)
func (h *SyncedHashSet[T, U]) Vals() iter.Iter[T] {
	return h.HashSet.Vals().SetupTeardown(
		func() error { h.RLock(); return nil },
		func() error { h.RUnlock(); return nil },
	)
}

// Panics, hash sets are not addressable.
func (h *HashSet[T, U]) ValPntrs() iter.Iter[*T] {
	panic(getNonAddressablePanicText("hash set"))
}

// Description: Contains will return true if the supplied value is in the hash
// set, false otherwise. All equality comparisons are performed by the generic U
// widget type that the hash set was initialized with.
//
// Time Complexity: The time complexity of Contains on a hash set is O(1).
func (h *HashSet[T, U]) Contains(v T) bool {
	return h.ContainsPntr(&v)
}

// Description: Places a read lock on the underlying hash set and then calls the
// underlying hash sets [HashSet.ContainsPntr] method.
//
// Lock Type: Read
//
// Time Complexity: O(n) (linear search)
func (h *SyncedHashSet[T, U]) Contains(v T) bool {
	h.RLock()
	defer h.RUnlock()
	return h.HashSet.ContainsPntr(&v)
}

// Description: ContainsPntr will return true if the supplied value is in the
// hash set, false otherwise. All equality comparisons are performed by the
// generic U widget type that the hash set was initialized with.
//
// Time Complexity: O(n) (linear search)
func (h *HashSet[T, U]) ContainsPntr(v *T) bool {
	_, rv := h.getHashPosition(v)
	return rv
}

// Description: Places a read lock on the underlying hash set and then calls the
// underlying hash sets [HashSet.ContainsPntr] method.
//
// Lock Type: Read
//
// Time Complexity: O(n) (linear search)
func (h *SyncedHashSet[T, U]) ContainsPntr(v *T) bool {
	h.RLock()
	defer h.RUnlock()
	return h.HashSet.ContainsPntr(v)
}

func (h *HashSet[T, U]) getHashPosition(v *T) (hash.Hash, bool) {
	w := widgets.NewWidget[T, U]()
	for i := w.Hash(v); ; i++ {
		if iterV, found := h.internalHashSetImpl[i]; found && w.Eq(v, &iterV) {
			return i, true
		} else if !found {
			return hash.Hash(0), false
		}
	}
}

// Description: AppendUnique will append the supplied values to the hash set if
// they are not already present in the hash set (unique). Non-unique values will
// not be appended. This function will never return an error.
//
// Time Complexity: Best case O(m) (no reallocation), worst case O(n+m)
// (reallocation), where m=len(vals).
func (h *HashSet[T, U]) AppendUnique(vals ...T) error {
	for i := 0; i < len(vals); i++ {
		h.appendOp(&vals[i])
	}
	return nil
}

// Description: Places a write lock on the underlying hash set and then calls
// the underlying hash sets [HashSet.AppendUnique] implementaion method. The
// [HashSet.AppendUnique] method is not called directly to avoid copying the
// vals varargs twice, which could be expensive with a large type for the T
// generic or a large number of values.
//
// Lock Type: Write
//
// Time Complexity: Best case O(m) (no reallocation), worst case O(n+m)
// (reallocation), where m=len(vals).
func (h *SyncedHashSet[T, U]) AppendUnique(vals ...T) error {
	h.Lock()
	defer h.Unlock()
	for i := 0; i < len(vals); i++ {
		h.appendOp(&vals[i])
	}
	return nil
}

func (h *HashSet[T, U]) appendOp(v *T) {
	w := widgets.NewWidget[T, U]()
	for i := w.Hash(v); ; i++ {
		if iterV, found := h.internalHashSetImpl[i]; !found {
			h.internalHashSetImpl[i] = *v
			break
		} else if w.Eq(v, &iterV) {
			break
		}
	}
}

// Description: Pop will remove the first num occurrences of val in the hash set.
// All equality comparisons are performed by the generic U widget type that the
// hash set was initialized with. If num is <=0 then no values will be poped and
// the hash set will not change.
//
// Time Complexity: O(m), where m=num
func (h *HashSet[T, U]) Pop(v T) int {
	return h.popImpl(&v)
}

// Description: Places a write lock on the underlying hash set and then calls
// the underlying hash sets [hash set.Pop] implementation method. The
// [HashSet.Pop] method is not called directly to avoid copying the vals varargs
// twice, which could be expensive with a large type for the T generic or a
// large number of values.
//
// Lock Type: Write
//
// Time Complexity: O(m), where m=num
func (h *SyncedHashSet[T, U]) Pop(v T) int {
	h.Lock()
	defer h.Unlock()
	return h.HashSet.popImpl(&v)
}

func (h *HashSet[T, U]) popImpl(v *T) int {
	w := widgets.NewWidget[T, U]()
	if i, cont := h.getHashPosition(v); cont {
		delete(h.internalHashSetImpl, i)
		curPos := i
		for j := i + 1; ; j++ {
			if iterV, found := h.internalHashSetImpl[j]; found {
				// If the hash of iterV is > the curPos then do not move it, it
				// is already in the correct position, but continue on this
				// collision chain to check for other values that could move.
				if w.Hash(&iterV) > curPos {
					continue
				}
				h.internalHashSetImpl[curPos] = h.internalHashSetImpl[j]
				delete(h.internalHashSetImpl, j)
				curPos = j
			} else {
				break
			}
		}
		return 1
	}
	return 0
}

// Description: Clears all values from the hash set. Equivalent to making a new
// hash set and setting it equal to the current one.
//
// Time Complexity: O(1)
func (h *HashSet[T, U]) Clear() {
	w := widgets.NewWidget[T, U]()
	for _, v := range h.internalHashSetImpl {
		w.Zero(&v)
	}
	h.internalHashSetImpl = make(internalHashSetImpl[T])
}

// Description: Places a write lock on the underlying hash set and then calls
// the underlying hash sets [HashSet.Clear] method.
//
// Lock Type: Write
//
// Time Complexity: O(1)
func (h *SyncedHashSet[T, U]) Clear() {
	h.Lock()
	defer h.Unlock()
	h.HashSet.Clear()
}

// Description: Returns true if the elements in h are all contained in other and
// the elements of other are all contained in h, regardless of position. Returns
// false otherwise.
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the ContainsPntr method on other. In big-O it might look something like this,
// O(n*O(other.ContainsPntr))), where O(other.ContainsPntr) represents the time
// complexity of the ContainsPntr method on other with m values.
func (h *HashSet[T, U]) UnorderedEq(
	other containerTypes.ComparisonsOtherConstraint[T],
) bool {
	if len(h.internalHashSetImpl) != other.Length() {
		return false
	}
	for _, iterV := range h.internalHashSetImpl {
		if !other.ContainsPntr(&iterV) {
			return false
		}
	}
	return true
}

// Description: Places a read lock on the underlying hash set and then calls the
// underlying hash sets [HashSet.UnorderedEq] method. Attempts to place a read
// lock on other but whether or not that happens is implementation dependent.
//
// Lock Type: Read on this hash set, read on other
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the ContainsPntr method on other. In big-O it might look something like this,
// O(n*O(other.ContainsPntr))), where O(other.ContainsPntr) represents the time
// complexity of the ContainsPntr method on other with m values.
func (h *SyncedHashSet[T, U]) UnorderedEq(
	other containerTypes.ComparisonsOtherConstraint[T],
) bool {
	h.RLock()
	other.RLock()
	defer h.RUnlock()
	defer other.RUnlock()
	return h.HashSet.UnorderedEq(other)
}

// Description: Populates the hash set with the intersection of values from the
// l and r containers. This hash set will be cleared before storing the result.
// When clearing, the new resulting hash set will be initialized with zero
// length and enough backing capacity to store (l.Length()+r.Length())/2
// elements before reallocating. This means that there should be at most 1
// reallocation beyond this initial allocation, and that additional allocation
// should only occur when the length of the intersection is greater than the
// average length of the l and r hash sets. This logic is predicated on the fact
// that intersections will likely be much smaller than the original hash sets.
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the ContainsPntr method on l and r. In big-O it might look something like
// this, O(O(r.ContainsPntr)*O(l.ContainsPntr)), where O(r.ContainsPntr)
// represents the time complexity of the containsPntr method on r and
// O(l.ContainsPntr) represents the time complexity of the containsPntr method
// on l.
func (h *HashSet[T, U]) Intersection(
	l containerTypes.ComparisonsOtherConstraint[T],
	r containerTypes.ComparisonsOtherConstraint[T],
) {
	newH := HashSet[T, U]{
		internalHashSetImpl: make(
			internalHashSetImpl[T], (l.Length()+r.Length())/2,
		),
	}
	addressableSafeValIter[T](
		r,
		func(index int, val *T) (iter.IteratorFeedback, error) {
			if l.ContainsPntr(val) {
				newH.appendOp(val)
			}
			return iter.Continue, nil
		},
	)
	h.Clear()
	*h = newH
}

// Description: Places a write lock on the underlying hash set and then calls
// the underlying hash sets [HashSet.Intersection] method. Attempts to place a
// read lock on l and r but whether or not that happens is implementation
// dependent.
//
// Lock Type: Write on this hash set, read on l and r
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the ContainsPntr method on l and r. In big-O it might look something like
// this, O(O(r.ContainsPntr)*O(l.ContainsPntr)), where O(r.ContainsPntr)
// represents the time complexity of the containsPntr method on r and
// O(l.ContainsPntr) represents the time complexity of the containsPntr method
// on l.
func (h *SyncedHashSet[T, U]) Intersection(
	l containerTypes.ComparisonsOtherConstraint[T],
	r containerTypes.ComparisonsOtherConstraint[T],
) {
	h.Lock()
	l.RLock()
	r.RLock()
	defer h.Unlock()
	defer l.RUnlock()
	defer r.RUnlock()
	h.HashSet.Intersection(l, r)
}

// Description: Populates the hash set with the union of values from the l and r
// containers. This hash set will be cleared before storing the result. When
// clearing, the new resulting hash set will be initialized with zero capacity
// and enough backing memory to store the average of the maximum and minimum
// possible union sizes before reallocating. This means that there should be at
// most 1 reallocation beyond this initial allocation, and that additional
// allocation should only occur when the length of the union is greater than the
// average length of the minumum and maximum possible union sizes. This logic is
// predicated on the fact that unions will likely be much smaller than the
// original hash sets.
//
// Time Complexity: O((n+m)*(n+m)), where n is the number of values in l and m
// is the number of values in r.
func (h *HashSet[T, U]) Union(
	l containerTypes.ComparisonsOtherConstraint[T],
	r containerTypes.ComparisonsOtherConstraint[T],
) {
	minLen := max(l.Length(), r.Length())
	maxLen := l.Length() + r.Length()
	newH := HashSet[T, U]{
		internalHashSetImpl: make(internalHashSetImpl[T], (minLen+maxLen)/2),
	}
	op := func(index int, val *T) (iter.IteratorFeedback, error) {
		newH.appendOp(val) // This also works with AppendUnique?? Shouldn't, check sync
		return iter.Continue, nil
	}
	addressableSafeValIter[T](l, op)
	addressableSafeValIter[T](r, op)
	h.Clear()
	*h = newH
}

// Description: Places a write lock on the underlying hash set and then calls
// the underlying hash sets [hash set.Union] method. Attempts to place a read
// lock on l and r but whether or not that happens is implementation dependent.
//
// Lock Type: Write on this hash set, read on l and r
//
// Time Complexity: O((n+m)*(n+m)), where n is the number of values in l and m
// is the number of values in r.
func (h *SyncedHashSet[T, U]) Union(
	l containerTypes.ComparisonsOtherConstraint[T],
	r containerTypes.ComparisonsOtherConstraint[T],
) {
	h.Lock()
	l.RLock()
	r.RLock()
	defer h.Unlock()
	defer l.RUnlock()
	defer r.RUnlock()
	h.HashSet.Union(l, r)
}

// Description: Populates the hash set with the result of taking the difference
// of r from l. This hash set will be cleared before storing the result. When
// clearing, the new resulting hash set will be initialized with zero capacity
// and enough backing memory to store half the length of l. This means that
// there should be at most 1 reallocation beyond this initial allocation, and
// that additional allocation should only occur when the length of the
// difference is greater than half the length of l. This logic is predicated on
// the fact that differences will likely be much smaller than the original hash
// set.
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the ContainsPntr method on l and r. In big-O it might look something like
// this, O(O(r.ContainsPntr)*O(l.ContainsPntr)), where O(r.ContainsPntr)
// represents the time complexity of the containsPntr method on r and
// O(l.ContainsPntr) represents the time complexity of the containsPntr method
// on l.
func (h *HashSet[T, U]) Difference(
	l containerTypes.ComparisonsOtherConstraint[T],
	r containerTypes.ComparisonsOtherConstraint[T],
) {
	newH := HashSet[T, U]{
		internalHashSetImpl: make(
			internalHashSetImpl[T], len(h.internalHashSetImpl)/2,
		),
	}
	addressableSafeValIter[T](
		l,
		func(index int, val *T) (iter.IteratorFeedback, error) {
			if !r.ContainsPntr(val) {
				newH.appendOp(val)
			}
			return iter.Continue, nil
		},
	)
	h.Clear()
	*h = newH
}

// Description: Places a write lock on the underlying hash set and then calls
// the underlying hash sets [hash set.Difference] method. Attempts to place a
// read lock on l and r but whether or not that happens is implementation
// dependent.
//
// Lock Type: Write on this hash set, read on l and r
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the ContainsPntr method on l and r. In big-O it might look something like
// this, O(O(r.ContainsPntr)*O(l.ContainsPntr)), where O(r.ContainsPntr)
// represents the time complexity of the containsPntr method on r and
// O(l.ContainsPntr) represents the time complexity of the containsPntr method
// on l.
func (h *SyncedHashSet[T, U]) Difference(
	l containerTypes.ComparisonsOtherConstraint[T],
	r containerTypes.ComparisonsOtherConstraint[T],
) {
	r.RLock()
	l.RLock()
	h.Lock()
	defer r.RUnlock()
	defer l.RUnlock()
	defer h.Unlock()
	h.HashSet.Difference(l, r)
}

// Description: Returns true if this hash set is a superset to other.
//
// Time Complexity: O(m), where m is the numbe rof values in other.
func (h *HashSet[T, U]) IsSuperset(
	other containerTypes.ComparisonsOtherConstraint[T],
) bool {
	rv := (len(h.internalHashSetImpl) >= other.Length())
	if !rv {
		return false
	}
	addressableSafeValIter[T](
		other,
		func(index int, val *T) (iter.IteratorFeedback, error) {
			if rv = h.ContainsPntr(val); !rv {
				return iter.Break, nil
			}
			return iter.Continue, nil
		},
	)
	return rv
}

// Description: Places a read lock on the underlying hash set and then calls the
// underlying hash sets [HashSet.IsSuperset] method. Attempts to place a read
// lock on other but whether or not that happens is implementation dependent.
//
// Lock Type: Read on this hash set, read on other
//
// Time Complexity: O(m), where m is the numbe rof values in other.
func (h *SyncedHashSet[T, U]) IsSuperset(
	other containerTypes.ComparisonsOtherConstraint[T],
) bool {
	h.RLock()
	other.RLock()
	defer h.RUnlock()
	defer other.RUnlock()
	return h.HashSet.IsSuperset(other)
}

// Description: Returns true if this hash set is a subset to other.
//
// Time Complexity: Dependent on the ContainsPntr method of other. In big-O
// terms it may look somwthing like this: O(n*O(other.ContainsPntr)), where n is
// the number of elements in the current hash set and other.ContainsPntr
// represents the time complexity of the containsPntr method on other.
func (h *HashSet[T, U]) IsSubset(
	other containerTypes.ComparisonsOtherConstraint[T],
) bool {
	rv := (len(h.internalHashSetImpl) <= other.Length())
	for _, iterV := range h.internalHashSetImpl {
		rv = other.ContainsPntr(&iterV)
		if !rv {
			break
		}
	}
	return rv
}

// Description: Places a read lock on the underlying hash set and then calls the
// underlying hash sets [HashSet.IsSubset] method. Attempts to place a read lock
// on other but whether or not that happens is implementation dependent.
//
// Lock Type: Read on this hash set, read on other
//
// Time Complexity: Dependent on the ContainsPntr method of other. In big-O
// terms it may look somwthing like this: O(n*O(other.ContainsPntr)), where n is
// the number of elements in the current hash set and other.ContainsPntr
// represents the time complexity of the containsPntr method on other.
func (h *SyncedHashSet[T, U]) IsSubset(
	other containerTypes.ComparisonsOtherConstraint[T],
) bool {
	h.RLock()
	other.RLock()
	defer h.RUnlock()
	defer other.RUnlock()
	return h.HashSet.IsSubset(other)
}

// An equality function that implements the [algo.widget.WidgetInterface]
// interface. Internally this is equivalent to [HashSet.UnorderedEq]. Returns
// true if l==r, false otherwise.
func (h *HashSet[T, U]) Eq(l *HashSet[T, U], r *HashSet[T, U]) bool {
	return l.UnorderedEq(r)
}

// An equality function that implements the [algo.widget.WidgetInterface]
// interface. Internally this is equivalent to [SyncedHashSet.UnorderedEq].
// Returns true if l==r, false otherwise.
func (h *SyncedHashSet[T, U]) Eq(
	l *SyncedHashSet[T, U],
	r *SyncedHashSet[T, U],
) bool {
	h.RLock()
	defer h.RUnlock()
	return l.UnorderedEq(r)
}

// Panics, sets cannot be compared for order.
func (h *HashSet[T, U]) Lt(l *HashSet[T, U], r *HashSet[T, U]) bool {
	panic("Sets cannot be compared relative to each other.")
}

// Panics, sets cannot be compared for order.
func (h *SyncedHashSet[T, U]) Lt(
	l *SyncedHashSet[T, U],
	r *SyncedHashSet[T, U],
) bool {
	panic("Sets cannot be compared relative to each other.")
}

// A function that returns a hash of a hash set. To do this all of the individual
// hashes that are produced from the elements of the hash set are combined in a
// way that maintains identity, making it so the hash will represent the same
// equality operation that [HashSet.KeyedEq] and [HashSet.Eq] provide.
func (h *HashSet[T, U]) Hash(other *HashSet[T, U]) hash.Hash {
	cntr := 0
	var rv hash.Hash
	w := widgets.NewWidget[T, U]()
	for _, v := range other.internalHashSetImpl {
		if cntr == 0 {
			rv = w.Hash(&v)
			cntr++
		} else {
			rv = rv.CombineUnordered(w.Hash(&v))
		}
	}
	return rv
}

// Places a read lock on the underlying hash set of other and then calls others
// underlying hash set [hash set.IsSubset] method.
func (h *SyncedHashSet[T, U]) Hash(other *SyncedHashSet[T, U]) hash.Hash {
	other.RLock()
	defer other.RUnlock()
	return h.HashSet.Hash(&other.HashSet)
}

// An zero function that implements the [algo.widget.WidgetInterface] interface.
// Internally this is equivalent to [HashSet.Clear].
func (h *HashSet[T, U]) Zero(other *HashSet[T, U]) {
	other.Clear()
}

// An zero function that implements the [algo.widget.WidgetInterface] interface.
// Internally this is equivalent to [SyncedHashSet.Clear].
func (v *SyncedHashSet[T, U]) Zero(other *SyncedHashSet[T, U]) {
	other.Clear()
}
