package containers

import (
	"github.com/barbell-math/util/hash"
	"github.com/barbell-math/util/widgets"
)

type (
	// An interface defining the operations that can be performed with a hooked
	// hash set. These functions are call backs that the hooked hash set will
	// call as values are modified.
	HashSetHooks interface {
		addOp(hashLoc HashSetHash)
		deleteOp(
			deletedHash HashSetHash,
			updatedHashes map[OldHashSetHash]NewHashSetHash,
		)
		clearOp()
	}

	// A type that represents a hash set with call backs to notify a given value
	// of any changes to the set. Several functions are also added that allow
	// the internal state of the hash set to be viewed, including the hash
	// values that the set uses internally. All of the implementation logic is
	// provided by the [HashSet] type. The type constraints on the generics
	// define the logic for how value specific operations, such as equality
	// comparisons, will be handled.
	HookedHashSet[T any, U widgets.WidgetInterface[T]] struct {
		HashSet[T, U]
		hooks HashSetHooks
	}

	// A synchronized version of [HookedHashSet]. All operations will be wrapped
	// in the appropriate calls to the embedded RWMutex. A pointer to a RWMutex
	// is embedded rather than a value to avoid copying the lock value.
	SyncedHookedHashSet[T any, U widgets.WidgetInterface[T]] struct {
		SyncedHashSet[T, U]
		hooks HashSetHooks
	}
)

// Creates a new hooked hash set with enough memory to store size number of
// elements. Size must be >=0, an error will be returned if it is not. If size
// is 0 the hooked hash set will be initialized with 0 elements.
func NewHookedHashSet[T any, U widgets.WidgetInterface[T]](
	hooks HashSetHooks,
	size int,
) (HookedHashSet[T, U], error) {
	hs, err := NewHashSet[T, U](size)
	if err != nil {
		return HookedHashSet[T, U]{}, err
	}
	return HookedHashSet[T, U]{
		HashSet: hs,
		hooks:   hooks,
	}, nil
}

// Creates a new synced hooked hash set initialized with enough memory to hold
// size elements. Size must be >= 0, an error will be returned if it is not. If
// size is 0 the hash set will be initialized with 0 elements. The underlying
// RWMutex value will be fully unlocked upon initialization.
func NewSyncedHookedHashSet[T any, U widgets.WidgetInterface[T]](
	hooks HashSetHooks,
	size int,
) (SyncedHookedHashSet[T, U], error) {
	hs, err := NewSyncedHashSet[T, U](size)
	if err != nil {
		return SyncedHookedHashSet[T, U]{}, err
	}
	return SyncedHookedHashSet[T, U]{
		SyncedHashSet: hs,
		hooks:         hooks,
	}, nil
}

// Converts the supplied hooked hash set to a synchronized hash set. Beware: The 
// original non-synced hash set will remain useable.
func (h *HookedHashSet[T, U])ToSynced() SyncedHookedHashSet[T,U] {
	return SyncedHookedHashSet[T, U]{
		SyncedHashSet: h.HashSet.ToSynced(),
		hooks: h.hooks,
	}
}

// Description: Gets the underlying hash position for the supplied value. The
// boolean flag indicates if the value was found or not.
//
// Time Complexity: O(1)
func (h *HookedHashSet[T, U]) GetHashPosition(v *T) (HashSetHash, bool) {
	return h.getHashPosition(v)
}

// Description: Places a read lock on the underlying hash set and then attempts
// to get the underlying hash position for the supplied value. The boolean flag
// indicates if the value was found or not.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (h *SyncedHookedHashSet[T, U])GetHashPosition(v *T) (HashSetHash, bool) {
	h.RLock()
	h.RUnlock()
	return h.SyncedHashSet.HashSet.getHashPosition(v)
}

// Description: Gets the value from the supplied hash, returning an error if the
// supplied hash is not present in the set.
//
// Time Complexity: O(1)
func (h *HookedHashSet[T, U]) GetFromHash(internalHash HashSetHash) (T, error) {
	if v, ok := h.HashSet.internalHashSetImpl[internalHash]; ok {
		return v, nil
	}
	var tmp T
	return tmp, getKeyError[HashSetHash](&internalHash)
}

// Description: Places a read lock on the underlying hash set and then gets the
// value from the supplied hash, returning an error if the supplied hash is not
// present in the set.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (h *SyncedHookedHashSet[T, U]) GetFromHash(
	internalHash HashSetHash,
) (T,error) {
	if v, ok := h.SyncedHashSet.HashSet.internalHashSetImpl[internalHash]; ok {
		return v, nil
	}
	var tmp T
	return tmp, getKeyError[HashSetHash](&internalHash)
}

// Description: Appends the values to the underlying hash set, calling the addOp
// hook each time a value is successfully added to the set.
//
// Time Complexity: O(n), where n=len(vals)
func (h *HookedHashSet[T, U]) AppendUnique(vals ...T) error {
	for _, v := range vals {
		h.HashSet.AppendUnique(v)
		if vHash, ok := h.getHashPosition(&v); ok {
			h.hooks.addOp(vHash)
		}
	}
	// The equivalent method on a normal hash set will never return an error.
	return nil
}

// Description: Places a write lock on the underlying hash set before appending
// the values to the underlying hash set, calling the addOp hook each time a
// value is successfully added to the set.
//
// Lock Type: Write
//
// Time Complexity: O(n), where n=len(vals)
func (h *SyncedHookedHashSet[T, U]) AppendUnique(vals ...T) error {
	h.Lock()
	defer h.Unlock()
	for _, v := range vals {
		h.SyncedHashSet.HashSet.AppendUnique(v)
		if vHash, ok := h.getHashPosition(&v); ok {
			h.hooks.addOp(vHash)
		}
	}
	// The equivalent method on a normal hash set will never return an error.
	return nil
}

// Description: Removes the supplied element from the set if it is present. No
// action is performed if the value is not in the set. The deleteOp hook is
// called only if a value was removed from the set.
//
// Time Complexity: O(1)
func (h *HookedHashSet[T, U]) Pop(v T) int {
	if deletedHash, affectedHashes, cnt := h.HashSet.popAndGetAffectedHashes(&v); cnt > 0 {
		h.hooks.deleteOp(deletedHash, affectedHashes)
		return cnt
	}
	return 0
}

// Description: Places a write lock on the underlying hash set before removing
// the supplied element from the set if it is present. No action is performed if
// the value is not in the set. The deleteOp hook is called only if a value was
// removed from the set.
//
// Lock Type: Write
//
// Time Complexity: O(1)
func (h *SyncedHookedHashSet[T, U]) Pop(v T) int {
	h.Lock()
	defer h.Unlock()
	if deletedHash, affectedHashes, cnt := h.SyncedHashSet.HashSet.popAndGetAffectedHashes(&v); cnt > 0 {
		h.hooks.deleteOp(deletedHash, affectedHashes)
		return cnt
	}
	return 0
}

// Description: Removes the supplied element from the set if it is present. No
// action is performed if the value is not in the set. The deleteOp hook is
// called only if a value was removed from the set.
//
// Time Complexity: O(1)
func (h *HookedHashSet[T, U]) PopPntr(v *T) int {
	if deletedHash, affectedHashes, cnt := h.HashSet.popAndGetAffectedHashes(v); cnt > 0 {
		h.hooks.deleteOp(deletedHash, affectedHashes)
		return cnt
	}
	return 0
}

// Description: Places a write lock on the underlying hash set before removing
// the supplied element from the set if it is present. No action is performed if
// the value is not in the set. The deleteOp hook is called only if a value was
// removed from the set.
//
// Lock Type: Write
//
// Time Complexity: O(1)
func (h *SyncedHookedHashSet[T, U]) PopPntr(v *T) int {
	h.Lock()
	defer h.Unlock()
	if deletedHash, affectedHashes, cnt := h.SyncedHashSet.HashSet.popAndGetAffectedHashes(v); cnt > 0 {
		h.hooks.deleteOp(deletedHash, affectedHashes)
		return cnt
	}
	return 0
}

// Description: Calls the clearOp hook before clearing all values from the
// underlying hash set.
//
// Time Complexity: O(n)
func (h *HookedHashSet[T, U]) Clear() {
	h.hooks.clearOp()
	h.HashSet.Clear()
}

// Description: Places a write lock on the underlying hash set before calling
// the clearOp hook before clearing all values from the underlying hash set.
//
// Time Complexity: O(n)
func (h *SyncedHookedHashSet[T, U]) Clear() {
	h.Lock()
	defer h.Unlock()
	h.hooks.clearOp()
	h.SyncedHashSet.HashSet.Clear()
}

// An equality function that implements the [algo.widget.WidgetInterface]
// interface. Returns true if l==r, false otherwise.
func (_ *HookedHashSet[T, U])Eq(
	l *HookedHashSet[T,U],
	r *HookedHashSet[T,U],
) bool {
	return l.HashSet.Eq(&l.HashSet, &r.HashSet)
}

// An equality function that implements the [algo.widget.WidgetInterface]
// interface. Returns true if l==r, false otherwise.
func (_ *SyncedHookedHashSet[T, U])Eq(
	l *SyncedHookedHashSet[T,U],
	r *SyncedHookedHashSet[T,U],
) bool {
	return l.HashSet.Eq(&l.HashSet, &r.HashSet)
}

// An equality function that implements the [algo.widget.WidgetInterface]
// interface. Returns true if l<r, false otherwise.
func (_ *HookedHashSet[T, U])Lt(
	l *HookedHashSet[T,U],
	r *HookedHashSet[T,U],
) bool {
	return l.HashSet.Lt(&l.HashSet, &r.HashSet)
}

// An equality function that implements the [algo.widget.WidgetInterface]
// interface. Returns true if l<r, false otherwise.
func (_ *SyncedHookedHashSet[T, U])Lt(
	l *SyncedHookedHashSet[T,U],
	r *SyncedHookedHashSet[T,U],
) bool {
	return l.HashSet.Lt(&l.HashSet, &r.HashSet)
}

// A function that returns a hash of a vector to implement the
// [algo.widget.WidgetInterface]. To do this all of the individual hashes that
// are produced from the elements of the set are combined in a way that
// maintains identity, making it so the hash will represent the same equality
// operation that [HookedHashSet.KeyedEq] and [HookedHashSet.Eq] provide.
func (_ *HookedHashSet[T, U])Hash(other *HookedHashSet[T,U]) hash.Hash {
	return other.HashSet.Hash(&other.HashSet)
}

// A function that returns a hash of a vector to implement the
// [algo.widget.WidgetInterface]. To do this all of the individual hashes that
// are produced from the elements of the set are combined in a way that
// maintains identity, making it so the hash will represent the same equality
// operation that [SyncedHookedHashSet.KeyedEq] and [SyncedHookedHashSet.Eq]
// provide.
func (_ *SyncedHookedHashSet[T, U])Hash(
	other *SyncedHookedHashSet[T,U],
) hash.Hash {
	return other.HashSet.Hash(&other.HashSet)
}

// An zero function that implements the [algo.widget.WidgetInterface] interface.
// Internally this is equivalent to [HookedHashSet.Clear].
func (_ *HookedHashSet[T, U])Zero(other *HookedHashSet[T,U]) {
	other.HashSet.Zero(&other.HashSet)
}

// An zero function that implements the [algo.widget.WidgetInterface] interface.
// Internally this is equivalent to [HookedHashSet.Clear].
func (_ *SyncedHookedHashSet[T, U])Zero(other *SyncedHookedHashSet[T,U]) {
	other.HashSet.Zero(&other.HashSet)
}
