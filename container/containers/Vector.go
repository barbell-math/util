package containers

import (
	"sync"

	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/hash"
	"github.com/barbell-math/util/iter"
	"github.com/barbell-math/util/widgets"
)

type (
	// A type to represent an array that dynamically grows as elements are added.
	// This is nothing more than a generically initialized slice with methods
	// attached to it so it can be passed to functions that use the interfaces
	// defined in the [containerTypes], [staticContainers], or
	// [dynamicContainers] packages. The type constraints on the generics
	// define the logic for how value specific operations, such as equality
	// comparisons, will be handled.
	Vector[T any, U widgets.BaseInterface[T]] []T

	// A synchronized version of [Vector]. All operations will be wrapped in the
	// appropriate calls to the embedded RWMutex. A pointer to a RWMutex is
	// embedded rather than a value to avoid copying the lock value.
	SyncedVector[T any, U widgets.BaseInterface[T]] struct {
		*sync.RWMutex
		Vector[T, U]
	}
)

// Creates a new vector initialized with size zero valued elements. Size must be
// >= 0, an error will be returned if it is not. If size is 0 the vector will be
// initialized with 0 elements. A vector can also be created by type casting a
// standard slice, as shown below.
//
//	// Vector to slice.
//	v,_:=NewVector[string,builtinBases.BuiltinString](3)
//	s:=[]string(v)
//	// Slice to vector.
//	s2:=make([]string,4)
//	v2:=Vector[string,builtinBases.BuiltinString](s2)
//
// Note that by performing the above type casts the operations provided by the
// widget, including equality, are not preserved.
func NewVector[T any, U widgets.BaseInterface[T]](size int) (Vector[T, U], error) {
	if size < 0 {
		return Vector[T, U]{}, getSizeError(size)
	}
	return make(Vector[T, U], size), nil
}

// Creates a new synced vector initialized with size zero valued elements. Size
// must be >= 0, an error will be returned if it is not. If size is 0 the vector
// will be initialized with 0 elements. The underlying RWMutex value will be
// fully unlocked upon initialization.
func NewSyncedVector[T any, U widgets.BaseInterface[T]](
	size int,
) (SyncedVector[T, U], error) {
	rv, err := NewVector[T, U](size)
	return SyncedVector[T, U]{
		RWMutex: &sync.RWMutex{},
		Vector:  rv,
	}, err
}

// Creates a new vector and populates it with the supplied values.
func VectorValInit[T any, U widgets.BaseInterface[T]](vals ...T) Vector[T, U] {
	return Vector[T, U](vals)
}

// Creates a new synced vector and populates it with the supplied values.
func SyncedVectorValInit[T any, U widgets.BaseInterface[T]](
	vals ...T,
) SyncedVector[T, U] {
	return SyncedVector[T, U]{
		Vector:  vals,
		RWMutex: &sync.RWMutex{},
	}
}

// Converts the supplied vector to a syncronized vector. Beware: The original
// non-synced vector will remain useable.
func (v *Vector[T, U]) ToSynced() SyncedVector[T, U] {
	return SyncedVector[T, U]{
		RWMutex: &sync.RWMutex{},
		Vector:  *v,
	}
}

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (v *Vector[T, U]) Lock() {}

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (v *Vector[T, U]) Unlock() {}

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (v *Vector[T, U]) RLock() {}

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (v *Vector[T, U]) RUnlock() {}

// The SyncedVector method to override the Vector pass through function and
// actually apply the mutex operation.
func (v *SyncedVector[T, U]) Lock() { v.RWMutex.Lock() }

// The SyncedVector method to override the Vector pass through function and
// actually apply the mutex operation.
func (v *SyncedVector[T, U]) Unlock() { v.RWMutex.Unlock() }

// The SyncedVector method to override the Vector pass through function and
// actually apply the mutex operation.
func (v *SyncedVector[T, U]) RLock() { v.RWMutex.RLock() }

// The SyncedVector method to override the Vector pass through function and
// actually apply the mutex operation.
func (v *SyncedVector[T, U]) RUnlock() { v.RWMutex.RUnlock() }

// Returns true, vectors are addressable.
func (v *Vector[T, U]) IsAddressable() bool { return true }

// Returns false, a vector is not synced.
func (h *Vector[T, U]) IsSynced() bool { return false }

// Returns true, a synced vector is synced. :O
func (h *SyncedVector[T, U]) IsSynced() bool { return true }

// Description: Returns the length of the vector.
//
// Time Complexity: O(1)
func (v *Vector[T, U]) Length() int {
	return len(*v)
}

// Description: Places a read lock on the underlying vector and then calls the
// underlying vectors [Vector.Length] method.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (v *SyncedVector[T, U]) Length() int {
	v.RLock()
	defer v.RUnlock()
	return v.Vector.Length()
}

// Description: Returns the capacity of the vector.
//
// Time Complexity: O(1)
func (v *Vector[T, U]) Capacity() int {
	return cap(*v)
}

// Description: Places a read lock on the underlying vector and then calls the
// underlying vectors [Vector.Capacity] method.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (v *SyncedVector[T, U]) Capacity() int {
	v.RLock()
	defer v.RUnlock()
	return v.Vector.Capacity()
}

// Description: Sets the capacity of the underlying slice. If the new capacity
// is less than the old capacity then values at the end of the slice will be
// dropped.
//
// Time Complexity: O(n) because a copy operation is performed.
func (v *Vector[T, U]) SetCapacity(c int) error {
	w := widgets.Base[T, U]{}
	for i := c + 1; i < len(*v); i++ {
		w.Zero(&(*v)[i])
	}
	tmp := make(Vector[T, U], c)
	copy(tmp, *v)
	*v = tmp
	return nil
}

// Description: Places a write lock on the underlying vector and then calls the
// underlying vectors [Vector.SetCapacity] method.
//
// Lock Type: Write
//
// Time Complexity: O(n), same as [Vector.SetCapacity].
func (v *SyncedVector[T, U]) SetCapacity(c int) error {
	v.Lock()
	defer v.Unlock()
	return v.Vector.SetCapacity(c)
}

// Description: Gets the value at the specified index. Returns an error if the
// index is >= the length of the vector.
//
// Time Complexity: O(1)
func (v *Vector[T, U]) Get(idx int) (T, error) {
	if idx >= 0 && idx < len(*v) && len(*v) > 0 {
		return (*v)[idx], nil
	}
	var tmp T
	return tmp, getIndexOutOfBoundsError(idx, 0, len(*v))
}

// Description: Places a read lock on the underlying vector and then gets the
// value at the specified index. Exhibits the same behavior as the [Vector.Get]
// method. The underlying [Vector.Get] method is not called to avoid copying the
// return value twice, which could be inefficient with a large value for the T
// generic.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (v *SyncedVector[T, U]) Get(idx int) (T, error) {
	v.RLock()
	defer v.RUnlock()
	if idx >= 0 && idx < len(v.Vector) && len(v.Vector) > 0 {
		return (v.Vector)[idx], nil
	}
	var tmp T
	return tmp, getIndexOutOfBoundsError(idx, 0, len(v.Vector))
}

// Description: Gets a pointer to the value at the specified index. Returns an
// error if the index is >= the length of the vector.
//
// Time Complexity: O(1)
func (v *Vector[T, U]) GetPntr(idx int) (*T, error) {
	if idx >= 0 && idx < len(*v) && len(*v) > 0 {
		return &(*v)[idx], nil
	}
	return nil, getIndexOutOfBoundsError(idx, 0, len(*v))
}

// Description: Places a read lock on the underlying vector and then calls the
// underlying vectors [Vector.GetPntr] method.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (v *SyncedVector[T, U]) GetPntr(idx int) (*T, error) {
	v.RLock()
	defer v.RUnlock()
	return v.Vector.GetPntr(idx)
}

// Description: Populates the supplied value with the value that is in the
// container. This is useful when storing structs and the structs identity as
// defined by the U widget only depends on a subset of the structs fields. This
// function allows for getting the entire value based on just the part of the
// struct that defines it's identity. Returns a value error if the value is not
// found in the vector.
//
// Time complexity: O(n)
func (v *Vector[T, U]) GetUnique(val *T) error {
	w := widgets.Base[T, U]{}
	for i := 0; i < len(*v); i++ {
		if w.Eq(val, &(*v)[i]) {
			*val = (*v)[i]
			return nil
		}
	}
	return getValueError[T](val)
}

// Description: Places a read lock on the underlying vector and then calls the
// underlying vectors [Vector.GetUnique] method.
//
// Lock Type: Read
//
// Time Complexity: O(n)
func (v *SyncedVector[T, U]) GetUnique(val *T) error {
	v.RLock()
	defer v.RUnlock()
	return v.Vector.GetUnique(val)
}

// Description: Contains will return true if the supplied value is in the
// vector, false otherwise. All equality comparisons are performed by the
// generic U widget type that the vector was initialized with.
//
// Time Complexity: O(n) (linear search)
func (v *Vector[T, U]) Contains(val T) bool {
	return v.ContainsPntr(&val)
}

// Description: Places a read lock on the underlying vector and then calls the
// underlying vectors [Vector.ContainsPntr] method.
//
// Lock Type: Read
//
// Time Complexity: O(n) (linear search)
func (v *SyncedVector[T, U]) Contains(val T) bool {
	v.RLock()
	defer v.RUnlock()
	return v.Vector.ContainsPntr(&val)
}

// Description: ContainsPntr will return true if the supplied value is in the
// vector, false otherwise. All equality comparisons are performed by the
// generic U widget type that the vector was initialized with.
//
// Time Complexity: O(n) (linear search)
func (v *Vector[T, U]) ContainsPntr(val *T) bool {
	found := false
	w := widgets.Base[T, U]{}
	for i := 0; i < len(*v) && !found; i++ {
		found = w.Eq(val, &(*v)[i])
	}
	return found
}

// Description: Places a read lock on the underlying vector and then calls the
// underlying vectors [Vector.ContainsPntr] method.
//
// Lock Type: Read
//
// Time Complexity: O(n) (linear search)
func (v *SyncedVector[T, U]) ContainsPntr(val *T) bool {
	v.RLock()
	defer v.RUnlock()
	return v.Vector.ContainsPntr(val)
}

// Description: KeyOf will return the index of the first occurrence of the
// supplied value in the vector. If the value is not found then the returned
// index will be -1 and the boolean flag will be set to false. If the value is
// found then the boolean flag will be set to true. All equality comparisons are
// performed by the generic U widget type that the vector was initialized with.
//
// Time Complexity: O(n) (linear search)
func (v *Vector[T, U]) KeyOf(val T) (int, bool) {
	return v.KeyOfPntr(&val)
}

// Description: Places a read lock on the underlying vector and then calls the
// underlying vectors [Vector.KeyOf] implementation method. The [Vector.KeyOf]
// method is not called directly to avoid copying the val variable twice, which
// could be expensive with a large type for the T generic.
//
// Lock Type: Read
//
// Time Complexity: O(n) (linear search)
func (v *SyncedVector[T, U]) KeyOf(val T) (int, bool) {
	v.RLock()
	defer v.RUnlock()
	return v.Vector.KeyOfPntr(&val)
}

// Description: KeyOfPntr will return the index of the first occurrence of the
// supplied value in the vector. If the value is not found then the returned
// index will be -1 and the boolean flag will be set to false. If the value is
// found then the boolean flag will be set to true. All equality comparisons are
// performed by the generic U widget type that the vector was initialized with.
//
// Time Complexity: O(n) (linear search)
func (v *Vector[T, U]) KeyOfPntr(val *T) (int, bool) {
	rv := -1
	found := false
	w := widgets.Base[T, U]{}
	for i := 0; i < len(*v) && !found; i++ {
		if found = w.Eq(val, &(*v)[i]); found {
			rv = i
		}
	}
	return rv, found
}

// Description: Places a read lock on the underlying vector and then calls the
// underlying vectors [Vector.KeyOfPntr] method.
//
// Lock Type: Read
//
// Time Complexity: O(n) (linear search)
func (v *SyncedVector[T, U]) KeyOfPntr(val *T) (int, bool) {
	v.RLock()
	defer v.RUnlock()
	return v.Vector.KeyOfPntr(val)
}

// Description: Sets the values at the specified indexes. Returns an error if
// the index is >= the length of the vector. Stops setting values as soon as an
// error is encountered.
//
// Time Complexity: O(m), where m=len(vals)
func (v *Vector[T, U]) Set(vals ...basic.Pair[int, T]) error {
	for _, iterV := range vals {
		if iterV.A >= 0 && iterV.A < len(*v) && len(*v) > 0 {
			(*v)[iterV.A] = iterV.B
		} else {
			return getIndexOutOfBoundsError(iterV.A, 0, len(*v))
		}
	}
	return nil
}

// Description: Places a write lock on the underlying vector and then calls the
// underlying vectors [Vector.Set] method.
//
// Lock Type: Write
//
// Time Complexity: O(m), where m=len(vals)
func (v *SyncedVector[T, U]) Set(vals ...basic.Pair[int, T]) error {
	v.Lock()
	defer v.Unlock()
	return v.Vector.Set(vals...)
}

// Description: Sets the supplied values sequentially starting at the supplied
// index and continuing sequentailly after that. Returns and error if any index
// that is attempted to be set is >= the length of the vector. If an error
// occurs, all values will be set up until the value that caused the error.
//
// Time Complexity: O(m), where m=len(vals)
func (v *Vector[T, U]) SetSequential(idx int, vals ...T) error {
	if idx >= len(*v) {
		return getIndexOutOfBoundsError(idx, 0, len(*v))
	}
	numCopyableVals := min(len(*v)-idx, len(vals))
	copy((*v)[idx:idx+numCopyableVals], vals[0:numCopyableVals])
	if idx+len(vals) > len(*v) {
		return getIndexOutOfBoundsError(len(*v), 0, len(*v))
	}
	return nil
}

// Description: Places a write lock on the underlying vector and then calls the
// underlying vectors [Vector.SetSequential] method.
//
// Lock Type: Write
//
// Time Complexity: O(m), where m=len(vals)
func (v *SyncedVector[T, U]) SetSequential(idx int, vals ...T) error {
	v.Lock()
	defer v.Unlock()
	return v.Vector.SetSequential(idx, vals...)
}

// Description: Append the supplied values to the vector. This function will
// never return an error.
//
// Time Complexity: Best case O(m), worst case O(n+m) (reallocation), where
// m=len(vals).
func (v *Vector[T, U]) Append(vals ...T) error {
	*v = append(*v, vals...)
	return nil
}

// Description: Places a write lock on the underlying vector and then calls the
// underlying vectors [Vector.Append] method.
//
// Lock Type: Write
//
// Time Complexity: Best case O(m) (no reallocation), worst case O(n+m)
// (reallocation), where m=len(vals).
func (v *SyncedVector[T, U]) Append(vals ...T) error {
	v.Lock()
	defer v.Unlock()
	return v.Vector.Append(vals...)
}

// Description: AppendUnique will append the supplied values to the vector if
// they are not already present in the vector (unique). Non-unique values will
// not be appended. This function will never return an error.
//
// Time Complexity: Best case O(m) (no reallocation), worst case O(n+m)
// (reallocation), where m=len(vals).
func (v *Vector[T, U]) AppendUnique(vals ...T) error {
	found := false
	w := widgets.Base[T, U]{}
	for i := 0; i < len(vals); i++ {
		for j := 0; j < len(*v) && !found; j++ {
			found = w.Eq(&vals[i], &(*v)[j])
		}
		if !found {
			*v = append(*v, vals[i])
		}
	}
	return nil
}

// Description: updates the supplied value in the underlying vector set,
// assuming that it is present in the vector already. The hash must not change
// from the update and the updated value must compare equal to the original
// value. If these rules are broken then an update violation error will be
// returned. This method is useful when you are storing struct values and want
// to update a field that is not utilized when calculating the hash and is also
// ignored when comparing for equality. This assumes congruency is present
// between the hash and equality methods defined in the U widget. If the value
// is not found then a key error will be returned.
//
// Time Complexity: O(n)
func (v *Vector[T, U]) UpdateUnique(orig T, updateOp func(orig *T)) error {
	return v.updateUniqueOp(&orig, updateOp)
}

// Description: Places a write lock on the underlying vector and then calls
// the underlying vectors [Vector.UpdateUnique] implementation method. The
// [Vector.UpdateUnique] method is not called directly to avoid copying the
// supplied value, which could be expensive with a large type for the T generic.
//
// Lock Type: Write
//
// Time Complexity: O(n)
func (v *SyncedVector[T, U]) UpdateUnique(orig T, updateOp func(orig *T)) error {
	v.Lock()
	defer v.Unlock()
	return v.Vector.updateUniqueOp(&orig, updateOp)
}

func (v *Vector[T, U]) updateUniqueOp(orig *T, updateOp func(orig *T)) error {
	w := widgets.Base[T, U]{}
	idx, found := v.KeyOfPntr(orig)
	if !found {
		return getValueError[T](orig)
	}
	updateOp(orig)
	newHash := w.Hash(orig)
	oldHash := w.Hash(&(*v)[idx])
	if newHash != oldHash {
		return getUpdateViolationHashError[T](
			&(*v)[idx], orig, hash.Hash(oldHash), hash.Hash(newHash),
		)
	}
	if !w.Eq(&(*v)[idx], orig) {
		return getUpdateViolationEqError[T](&(*v)[idx], orig)
	}
	(*v)[idx] = *orig
	return nil
}

// Description: Places a write lock on the underlying vector and then calls the
// underlying vectors [Vector.AppendUnique] implementaion method. The
// [Vector.AppendUnique] method is not called directly to avoid copying the vals
// varargs twice, which could be expensive with a large type for the T generic
// or a large number of values.
//
// Lock Type: Write
//
// Time Complexity: Best case O(m) (no reallocation), worst case O(n+m)
// (reallocation), where m=len(vals).
func (v *SyncedVector[T, U]) AppendUnique(vals ...T) error {
	v.Lock()
	defer v.Unlock()
	return v.Vector.AppendUnique(vals...)
}

// Description: Insert will insert the supplied values into the vector. The
// values will be inserted in the order that they are given.
//
// Time Complexity: O(n*m), where m=len(vals)
func (v *Vector[T, U]) Insert(vals ...basic.Pair[int, T]) error {
	for i := 0; i < len(vals); i++ {
		if vals[i].A >= 0 && vals[i].A < len(*v) && len(*v) > 0 {
			var tmp T
			*v = append(*v, tmp)
			copy((*v)[vals[i].A+1:], (*v)[vals[i].A:])
			(*v)[vals[i].A] = vals[i].B
		} else if vals[i].A == len(*v) {
			*v = append(*v, vals[i].B)
		} else {
			return getIndexOutOfBoundsError(vals[i].A, 0, len(*v))
		}
	}
	return nil
}

// Description: Places a write lock on the underlying vector and then calls the
// underlying vectors [Vector.Insert] method.
//
// Lock Type: Write
//
// Time Complexity: O(n*m), where m=len(vals)
func (v *SyncedVector[T, U]) Insert(vals ...basic.Pair[int, T]) error {
	v.Lock()
	defer v.Unlock()
	return v.Vector.Insert(vals...)
}

// Description: Inserts the supplied values at the given index. Returns an error
// if the index is >= the length of the vector.
//
// Time Complexity: O(n+m), where m=len(vals). For time complexity see the
// InsertVector section of https://go.dev/wiki/SliceTricks
func (v *Vector[T, U]) InsertSequential(idx int, vals ...T) error {
	if idx >= 0 && idx < len(*v) && len(*v) > 0 {
		*v = append((*v)[:idx], append(vals, (*v)[idx:]...)...)
		return nil
	} else if idx == len(*v) {
		*v = append(*v, vals...)
		return nil
	}
	return getIndexOutOfBoundsError(idx, 0, len(*v))
}

// Description: Places a write lock on the underlying vector and then calls the
// underlying vectors [Vector.InsertSequential] method.
//
// Lock Type: Write
//
// Time Complexity: O(n+m), where m=len(vals). For time complexity see the
// InsertVector section of https://go.dev/wiki/SliceTricks
func (v *SyncedVector[T, U]) InsertSequential(idx int, vals ...T) error {
	v.Lock()
	defer v.Unlock()
	return v.Vector.InsertSequential(idx, vals...)
}

// Description: Pop will remove all occurrences of val in the vector. All
// equality comparisons are performed by the generic U widget type that the
// vector was initialized with.
//
// Time Complexity: O(n)
func (v *Vector[T, U]) Pop(val T) int {
	return v.popSequentialImpl(&val, containerTypes.PopAll)
}

// Description: Places a write lock on the underlying vector and then calls the
// underlying vectors [Vector.Pop] implementation method.
//
// Lock Type: Write
//
// Time Complexity: O(n)
func (v *SyncedVector[T, U]) Pop(val T) int {
	v.Lock()
	defer v.Unlock()
	return v.Vector.popSequentialImpl(&val, containerTypes.PopAll)
}

// Description: PopPntr will remove all occurrences of val in the vector. All
// equality comparisons are performed by the generic U widget type that the
// vector was initialized with.
//
// Time Complexity: O(n)
func (v *Vector[T, U]) PopPntr(val *T) int {
	v.Lock()
	defer v.Unlock()
	return v.popSequentialImpl(val, containerTypes.PopAll)
}

// Description: Places a write lock on the underlying vector and then calls the
// underlying vectors [Vector.PopPntr] implementation method.
//
// Lock Type: Write
//
// Time Complexity: O(n)
func (v *SyncedVector[T, U]) PopPntr(val *T) int {
	v.Lock()
	defer v.Unlock()
	return v.Vector.popSequentialImpl(val, containerTypes.PopAll)
}

// Description: PopSequential will remove the first num occurrences of val in
// the vector. All equality comparisons are performed by the generic U widget
// type that the vector was initialized with. If num is <=0 then no values will
// be poped and the vector will not change.
//
// Time Complexity: O(n)
func (v *Vector[T, U]) PopSequential(val T, num int) int {
	if num <= 0 {
		return 0
	}
	return v.popSequentialImpl(&val, num)
}

// Description: Places a write lock on the underlying vector and then calls the
// underlying vectors [Vector.PopSequential] implementation method.
//
// Lock Type: Write
//
// Time Complexity: O(n)
func (v *SyncedVector[T, U]) PopSequential(val T, num int) int {
	if num <= 0 {
		return 0
	}
	v.Lock()
	defer v.Unlock()
	return v.Vector.popSequentialImpl(&val, num)
}

func (v *Vector[T, U]) popSequentialImpl(val *T, num int) int {
	cntr := 0
	prevIndex := -1
	w := widgets.Base[T, U]{}
	for i := 0; i < len(*v); i++ {
		if w.Eq(val, &(*v)[i]) && cntr+1 <= num {
			if prevIndex == -1 { // Initial value found
				prevIndex = i
			} else {
				w.Zero(&(*v)[i])
				copy((*v)[prevIndex-cntr+1:i], (*v)[prevIndex+1:i])
				prevIndex = i
			}
			cntr++
			if cntr >= num {
				break
			}
		}
	}
	if prevIndex != -1 {
		copy((*v)[prevIndex-cntr+1:len(*v)], (*v)[prevIndex+1:len(*v)])
	}
	*v = (*v)[:len(*v)-cntr]
	return cntr
}

// Description: Deletes the value at the specified index. Returns an error if
// the index is >= the length of the vector.
//
// Time Complexity: O(n)
func (v *Vector[T, U]) Delete(idx int) error {
	if idx < 0 || idx >= len(*v) {
		return getIndexOutOfBoundsError(idx, 0, len(*v))
	} else if idx >= 0 && idx < len(*v) && len(*v) > 0 {
		w := widgets.Base[T, U]{}
		w.Zero(&(*v)[idx])
		*v = append((*v)[:idx], (*v)[idx+1:]...)
	}
	return nil
}

// Description: Places a write lock on the underlying vector and then calls the
// underlying vectors [Vector.Delete] method.
//
// Lock Type: Write
//
// Time Complexity: O(n)
func (v *SyncedVector[T, U]) Delete(idx int) error {
	v.Lock()
	defer v.Unlock()
	return v.Vector.Delete(idx)
}

// Description: Deletes the values in the index range [start,end). Returns an
// error if the start index is < 0, the end index is > the length of the
// vector, or the end index is < the start index.
//
// Time Complexity: O(n)
func (v *Vector[T, U]) DeleteSequential(start int, end int) error {
	if start < 0 {
		return getIndexOutOfBoundsError(start, 0, len(*v))
	}
	if end > len(*v) {
		return getIndexOutOfBoundsError(end, 0, len(*v))
	}
	if end <= start {
		return getStartEndIndexError(start, end)
	}
	w := widgets.Base[T, U]{}
	for i := start; i < end; i++ {
		w.Zero(&(*v)[i])
	}
	*v = append((*v)[0:start], (*v)[end:]...)
	return nil
}

// Description: Places a write lock on the underlying vector and then calls the
// underlying vectors [Vector.DeleteSequential] method.
//
// Lock Type: Write
//
// Time Complexity: O(n)
func (v *SyncedVector[T, U]) DeleteSequential(start int, end int) error {
	v.Lock()
	defer v.Unlock()
	return v.Vector.DeleteSequential(start, end)
}

// Description: Clears all values from the vector.
//
// Time Complexity: O(n) (Because of zeroing)
func (v *Vector[T, U]) Clear() {
	w := widgets.Base[T, U]{}
	for i := 0; i < len(*v); i++ {
		w.Zero(&(*v)[i])
	}
	*v = make(Vector[T, U], 0)
}

// Description: Places a write lock on the underlying vector and then calls the
// underlying vectors [Vector.Clear] method.
//
// Lock Type: Write
//
// Time Complexity: O(n)
func (v *SyncedVector[T, U]) Clear() {
	v.Lock()
	defer v.Unlock()
	v.Vector.Clear()
}

// Description: Returns the value at index 0 if one is present. If the vector
// has no elements then an error is returned.
//
// Time Complexity: O(1)
func (v *Vector[T, U]) PeekFront() (T, error) {
	if len(*v) > 0 {
		return (*v)[0], nil
	}
	var tmp T
	return tmp, getIndexOutOfBoundsError(0, 0, len(*v))
}

// Description: Places a read lock on the underlying vector and then attempts to
// return the value at index 0 if one is present. Exhibits the same behavior as
// the [Vector.PeekFront] method. The underlying [Vector.PeekFront] method is
// not called to avoid copying the value twice, which could be inefficient with
// a large type for the T generic.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (v *SyncedVector[T, U]) PeekFront() (T, error) {
	v.RLock()
	defer v.RUnlock()
	if len(v.Vector) > 0 {
		return (v.Vector)[0], nil
	}
	var tmp T
	return tmp, getIndexOutOfBoundsError(0, 0, len(v.Vector))
}

// Description: Returns a pointer to the value at index 0 if one is present. If
// the vector has no elements then an error is returned.
//
// Time Complexity: O(1)
func (v *Vector[T, U]) PeekPntrFront() (*T, error) {
	if len(*v) > 0 {
		return &(*v)[0], nil
	}
	return nil, getIndexOutOfBoundsError(0, 0, len(*v))
}

// Description: Places a read lock on the underlying vector and then calls the
// underlying vectors [Vector.PeekPntrFront] method.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (v *SyncedVector[T, U]) PeekPntrFront() (*T, error) {
	v.RLock()
	defer v.RUnlock()
	return v.Vector.PeekPntrFront()
}

// Description: Returns the value at index len(v)-1 if one is present. If the
// vector has no elements then an error is returned.
//
// Time Complexity: O(1)
func (v *Vector[T, U]) PeekBack() (T, error) {
	if len(*v) > 0 {
		return (*v)[len(*v)-1], nil
	}
	var tmp T
	return tmp, getIndexOutOfBoundsError(0, 0, len(*v))
}

// Description: Places a read lock on the underlying vector and then attempts to
// return the value at index len(v)-1 if one is present. Exhibits the same
// behavior as the [Vector.PeekBack] method. The underlying [Vector.PeekBack]
// method is not called to avoid copying the value twice, which could be
// inefficient with a large type for the T generic.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (v *SyncedVector[T, U]) PeekBack() (T, error) {
	v.RLock()
	defer v.RUnlock()
	if len(v.Vector) > 0 {
		return (v.Vector)[len(v.Vector)-1], nil
	}
	var tmp T
	return tmp, getIndexOutOfBoundsError(0, 0, len(v.Vector))
}

// Description: Returns a pointer to the value at index len(v)-1 if one is
// present. If the vector has no elements then an error is returned.
//
// Time Complexity: O(1)
func (v *Vector[T, U]) PeekPntrBack() (*T, error) {
	if len(*v) > 0 {
		return &(*v)[len(*v)-1], nil
	}
	return nil, getIndexOutOfBoundsError(0, 0, len(*v))
}

// Description: Places a read lock on the underlying vector and then calls the
// underlying vectors [Vector.PeekPntrBack] method.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (v *SyncedVector[T, U]) PeekPntrBack() (*T, error) {
	v.RLock()
	defer v.RUnlock()
	return v.Vector.PeekPntrBack()
}

// Description: Returns and removes the element at the front of the vector.
// Returns an error if the vector has no elements.
//
// Time Complexity: O(n)
func (v *Vector[T, U]) PopFront() (T, error) {
	var rv T
	return rv, v.popFontImpl(&rv)
}

// Description: Places a write lock on the underlying vector and then calls the
// underlying vectors [Vector.PopFront] implementation method.
//
// Lock Type: Write
//
// Time Complexity: O(n)
func (v *SyncedVector[T, U]) PopFront() (T, error) {
	v.Lock()
	defer v.Unlock()
	var rv T
	return rv, v.Vector.popFontImpl(&rv)
}

func (v *Vector[T, U]) popFontImpl(rv *T) error {
	w := widgets.Base[T, U]{}
	if len(*v) > 0 {
		*rv = (*v)[0]
		w.Zero(&(*v)[0])
		*v = (*v)[1:]
		return nil
	}
	return customerr.Wrap(containerTypes.Empty, "Nothing to pop!")
}

// Description: Returns and removes the element at the back of the vector.
// Returns an error if the vector has no elements.
//
// Time Complexity: O(1)
func (v *Vector[T, U]) PopBack() (T, error) {
	var rv T
	return rv, v.popBackImpl(&rv)
}

// Description: Places a write lock on the underlying vector and then calls the
// underlying vectors [Vector.PopFront] implementation method.
//
// Lock Type: Write
//
// Time Complexity: O(1)
func (v *SyncedVector[T, U]) PopBack() (T, error) {
	v.Lock()
	defer v.Unlock()
	var rv T
	return rv, v.Vector.popBackImpl(&rv)
}

func (v *Vector[T, U]) popBackImpl(rv *T) error {
	w := widgets.Base[T, U]{}
	if len(*v) > 0 {
		*rv = (*v)[len(*v)-1]
		w.Zero(&(*v)[len(*v)-1])
		*v = (*v)[:len(*v)-1]
		return nil
	}
	return customerr.Wrap(containerTypes.Empty, "Nothing to pop!")
}

// Description: Pushes an element to the back of the vector. Equivalent to
// appending values to the end of the vector. Values will be pushed back in the
// order that they are given. For example, calling push back on [0,1,2] with
// vals of [3,4] will result in [0,1,2,3,4].
//
// Time Complexity: best case O(m) (no reallocation), worst case O(n+m) (with
// reallocation), where m=len(vals)
func (v *Vector[T, U]) PushBack(vals ...T) error {
	*v = append(*v, vals...)
	return nil
}

// Description: Places a write lock on the underlying vector and then appends
// values to the end of the vector. Exhibits the same behavior as
// [Vector.PushBack]. The underlying [Vector.PushBack] method is not called to
// avoid copying the list of values twice, which could be inefficient with a
// large type for the T generic or many values.
//
// Lock Type: Write
//
// Time Complexity: best case O(m) (no reallocation), worst case O(n+m) (with
// reallocation), where m=len(vals)
func (v *SyncedVector[T, U]) PushBack(vals ...T) error {
	v.Lock()
	defer v.Unlock()
	v.Vector = append(v.Vector, vals...)
	return nil
}

// Description: Pushes an element to the front of the vector. Equivalent to
// inserting a single value at the front of the vector. Values will be pushed to
// the front in the order that they are given. For example, calling push front
// on [0,1,2] with vals of [3,4] will result in [3,4,0,1,2].
//
// Time Complexity: O(n+m), where m=len(vals)
func (v *Vector[T, U]) PushFront(vals ...T) error {
	*v = append(vals, (*v)...)
	return nil
}

// Description: Places a write lock on the underlying vector and then pushes
// values to the front of the vector. Exhibits the same behavior as
// [Vector.PusFront]. The underlying [Vector.PushFront] method is not called to
// avoid copying the list of values twice, which could be inefficient with a
// large type for the T generic or many values.
//
// Lock Type: Write
//
// Time Complexity: O(n+m), where m=len(vals)
func (v *SyncedVector[T, U]) PushFront(vals ...T) error {
	v.Lock()
	defer v.Unlock()
	v.Vector = append(vals, v.Vector...)
	return nil
}

// Description: Pushes an element to the back of the vector. Equivalent to
// appending a single value to the end of the vector. Has the same behavior as
// PushBack because the underlying vector grows as needed.
//
// Time Complexity: O(m), where m=len(vals)
func (v *Vector[T, U]) ForcePushBack(vals ...T) {
	*v = append(*v, vals...)
}

// Description: Places a write lock on the underlying vector and then pushes
// values to the front of the vector. Exhibits the same behavior as
// [Vector.ForcePushBack]. The underlying [Vector.ForcePushBack] method is not
// called to avoid copying the list of values twice, which could be inefficient
// with a large type for the T generic or many values.
//
// Lock Type: Write
//
// Time Complexity: O(m), where m=len(vals)
func (v *SyncedVector[T, U]) ForcePushBack(vals ...T) {
	v.Lock()
	defer v.Unlock()
	v.Vector = append(v.Vector, vals...)
}

// Description: Pushes an element to the front of the vector. Equivalent to
// inserting a single value at the front of the vector. Has the same behavior as
// PushBack because the underlying vector grows as needed.
//
// Time Complexity: O(n+m), where m=len(vals)
func (v *Vector[T, U]) ForcePushFront(vals ...T) {
	*v = append(vals, (*v)...)
}

// Description: Places a write lock on the underlying vector and then pushes
// values to the front of the vector. Exhibits the same behavior as
// [Vector.ForcePushFront]. The underlying [Vector.ForcePushFront] method is not
// called to avoid copying the list of values twice, which could be inefficient
// with a large type for the T generic or many values.
//
// Lock Type: Write
//
// Time Complexity: O(n+m), where m=len(vals)
func (v *SyncedVector[T, U]) ForcePushFront(vals ...T) {
	v.Lock()
	defer v.Unlock()
	v.Vector = append(vals, v.Vector...)
}

// Description: Returns an iterator that iterates over the values in the vector.
//
// Time Complexity: O(n)
func (v *Vector[T, U]) Vals() iter.Iter[T] {
	return iter.SequentialElems[T](
		len(*v),
		func(i int) (T, error) { return (*v)[i], nil },
	)
}

// Description: Modifies the iterator chain returned by the unerlying
// [Vector.Vals] method such that a read lock will be placed on the underlying
// vector when the iterator is consumed. The vector will have a read lock the
// entire time the iteration is being performed. The lock will not be applied
// until the iterator starts to be consumed.
//
// Lock Type: Read
//
// Time Complexity: O(n)
func (v *SyncedVector[T, U]) Vals() iter.Iter[T] {
	return v.Vector.Vals().SetupTeardown(
		func() error { v.RLock(); return nil },
		func() error { v.RUnlock(); return nil },
	)
}

// Description: Returns an iterator that iterates over the pointers to the
// values in the vector. The vector will have a read lock the entire time the
// iteration is being performed. The lock will not be applied until the iterator
// is consumed.
//
// Time Complexity: O(n)
func (v *Vector[T, U]) ValPntrs() iter.Iter[*T] {
	return iter.SequentialElems[*T](
		len(*v),
		func(i int) (*T, error) { return &(*v)[i], nil },
	)
}

// Description: Modifies the iterator chain returned by the unerlying
// [Vector.ValPntrs] method such that a read lock will be placed on the
// underlying vector when the iterator is consumed. The vector will have a read
// lock the entire time the iteration is being performed. The lock will not be
// applied until the iterator starts to be consumed.
//
// Lock Type: Read
//
// Time Complexity: O(n)
func (v *SyncedVector[T, U]) ValPntrs() iter.Iter[*T] {
	return v.Vector.ValPntrs().SetupTeardown(
		func() error { v.RLock(); return nil },
		func() error { v.RUnlock(); return nil },
	)
}

// Description: Returns an iterator that iterates over the keys (indexes) of the
// vector. The vector will have a read lock the entire time the iteration is
// being performed. The lock will not be applied until the iterator starts to be
// consumed.
//
// Time Complexity: O(n)
func (v *Vector[T, U]) Keys() iter.Iter[int] {
	return iter.Range[int](0, len(*v), 1)
}

// Description: Modifies the iterator chain returned by the unerlying
// [Vector.Keys] method such that a read lock will be placed on the underlying
// vector when iterator is consumed. The vector will have a read lock the entire
// time the iteration is being performed. The lock will not be applied until the
// iterator starts to be consumed.
//
// Lock Type: Read
//
// Time Complexity: O(n)
func (v *SyncedVector[T, U]) Keys() iter.Iter[int] {
	return v.Vector.Keys().SetupTeardown(
		func() error { v.RLock(); return nil },
		func() error { v.RUnlock(); return nil },
	)
}

// Description: Returns true if the elements in v are all contained in other and
// the elements of other are all contained in v, regardless of position. Returns
// false otherwise.
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the ContainsPntr method on other. In big-O it might look something like this,
// O(n*O(other.ContainsPntr))), where O(other.ContainsPntr) represents the time
// complexity of the ContainsPntr method on other with m values.
func (v *Vector[T, U]) UnorderedEq(
	other containerTypes.ComparisonsOtherConstraint[T],
) bool {
	rv := (len(*v) == other.Length())
	for i := 0; i < len(*v) && rv; i++ {
		rv = other.ContainsPntr(&(*v)[i])
	}
	return rv
}

// Description: Places a read lock on the underlying vector and then calls the
// underlying vectors [Vector.UnorderedEq] method. Attempts to place a read lock
// on other but whether or not that happens is implementation dependent.
//
// Lock Type: Read on this vector, read on other
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the ContainsPntr method on other. In big-O it might look something like this,
// O(n*O(other.ContainsPntr))), where O(other.ContainsPntr) represents the time
// complexity of the ContainsPntr method on other with m values.
func (v *SyncedVector[T, U]) UnorderedEq(
	other containerTypes.ComparisonsOtherConstraint[T],
) bool {
	v.RLock()
	other.RLock()
	defer v.RUnlock()
	defer other.RUnlock()
	return v.Vector.UnorderedEq(other)
}

// Description: Returns true if all the key value pairs in v are all contained
// in other and the key value pairs in other are all contained in v. Returns
// false otherwise.
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the Get/GetPntr method on other. In big-O it might look something like this,
// O(n*O(other.GetPntr))), where n is the number of elements in v and
// O(other.ContainsPntr) represents the time complexity of the containsPntr
// method on other.
func (v *Vector[T, U]) KeyedEq(
	other containerTypes.KeyedComparisonsOtherConstraint[int, T],
) bool {
	w := widgets.Base[T, U]{}
	rv := (len(*v) == other.Length())
	for i := 0; i < len(*v) && rv; i++ {
		if otherV, err := addressableSafeGet[int, T](other, i); err == nil {
			rv = w.Eq(otherV, &(*v)[i])
		} else {
			rv = false
		}
	}
	return rv
}

// Description: Places a read lock on the underlying vector and then calls the
// underlying vectors [Vector.KeyedEq] method. Attempts to place a read lock on
// other but whether or not that happens is implementation dependent.
//
// Lock Type: Read on this vector, read on other
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the GetPntr method on other. In big-O it might look something like this,
// O(n*O(other.GetPntr))), where n is the number of elements in v and
// O(other.ContainsPntr) represents the time complexity of the containsPntr
// method on other.
func (v *SyncedVector[T, U]) KeyedEq(
	other containerTypes.KeyedComparisonsOtherConstraint[int, T],
) bool {
	v.RLock()
	other.RLock()
	defer v.RUnlock()
	defer other.RUnlock()
	return v.Vector.KeyedEq(other)
}

// Description: Populates the vector with the intersection of values from the l
// and r containers. This vector will be cleared before storing the result. When
// clearing, the new resulting vector will be initialized with zero length and
// enough backing capacity to store (l.Length()+r.Length())/2 elements before
// reallocating. This means that there should be at most 1 reallocation beyond
// this initial allocation, and that additional allocation should only occur
// when the length of the intersection is greater than the average length of the
// l and r vectors. This logic is predicated on the fact that intersections will
// likely be much smaller than the original vectors.
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the ContainsPntr method on l and r. In big-O it might look something like
// this, O(O(r.ContainsPntr)*O(l.ContainsPntr)), where O(r.ContainsPntr)
// represents the time complexity of the containsPntr method on r and
// O(l.ContainsPntr) represents the time complexity of the containsPntr method
// on l.
func (v *Vector[T, U]) Intersection(
	l containerTypes.ComparisonsOtherConstraint[T],
	r containerTypes.ComparisonsOtherConstraint[T],
) {
	newV := make(Vector[T, U], 0, (l.Length()+r.Length())/2)
	addressableSafeValIter[T](l).ForEach(
		func(index int, val *T) (iter.IteratorFeedback, error) {
			if r.ContainsPntr(val) {
				newV = append(newV, *val)
			}
			return iter.Continue, nil
		},
	)
	v.Clear()
	*v = newV
}

// Description: Places a write lock on the underlying vector and then calls the
// underlying vectors [Vector.Intersection] method. Attempts to place a read
// lock on l and r but whether or not that happens is implementation dependent.
//
// Lock Type: Write on this vector, read on l and r
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the ContainsPntr method on l and r. In big-O it might look something like
// this, O(O(r.ContainsPntr)*O(l.ContainsPntr)), where O(r.ContainsPntr)
// represents the time complexity of the containsPntr method on r and
// O(l.ContainsPntr) represents the time complexity of the containsPntr method
// on l.
func (v *SyncedVector[T, U]) Intersection(
	l containerTypes.ComparisonsOtherConstraint[T],
	r containerTypes.ComparisonsOtherConstraint[T],
) {
	r.RLock()
	l.RLock()
	v.Lock()
	defer r.RUnlock()
	defer l.RUnlock()
	defer v.Unlock()
	v.Vector.Intersection(l, r)
}

// Description: Populates the vector with the union of values from the l and r
// containers. This vector will be cleared before storing the result. When
// clearing, the new resulting vector will be initialized with zero capacity and
// enough backing memory to store the average of the maximum and minimum
// possible union sizes before reallocating. This means that there should be at
// most 1 reallocation beyond this initial allocation, and that additional
// allocation should only occur when the length of the union is greater than the
// average length of the minimum and maximum possible union sizes. This logic is
// predicated on the fact that unions will likely be much smaller than the
// original vectors.
//
// Time Complexity: O((n+m)*(n+m)), where n is the number of values in l and m
// is the number of values in r.
func (v *Vector[T, U]) Union(
	l containerTypes.ComparisonsOtherConstraint[T],
	r containerTypes.ComparisonsOtherConstraint[T],
) {
	minLen := max(l.Length(), r.Length())
	maxLen := l.Length() + r.Length()
	newV := make(Vector[T, U], 0, (maxLen+minLen)/2)
	oper := func(index int, val *T) (iter.IteratorFeedback, error) {
		if !newV.ContainsPntr(val) {
			newV = append(newV, *val)
		}
		return iter.Continue, nil
	}
	addressableSafeValIter[T](l).ForEach(oper)
	addressableSafeValIter[T](r).ForEach(oper)
	v.Clear()
	*v = newV
}

// Description: Places a write lock on the underlying vector and then calls the
// underlying vectors [Vector.Union] method. Attempts to place a read lock on l
// and r but whether or not that happens is implementation dependent.
//
// Lock Type: Write on this vector, read on l and r
//
// Time Complexity: O((n+m)*(n+m)), where n is the number of values in l and m
// is the number of values in r.
func (v *SyncedVector[T, U]) Union(
	l containerTypes.ComparisonsOtherConstraint[T],
	r containerTypes.ComparisonsOtherConstraint[T],
) {
	r.RLock()
	l.RLock()
	v.Lock()
	defer r.RUnlock()
	defer l.RUnlock()
	defer v.Unlock()
	v.Vector.Union(l, r)
}

// Description: Populates the vector with the result of taking the difference of
// r from l. This vector will be cleared before storing the result. When
// clearing, the new resulting vector will be initialized with zero capacity and
// enough backing memory to store half the length of l. This means that there
// should be at most 1 reallocation beyond this initial allocation, and that
// additional allocation should only occur when the length of the difference is
// greater than half the length of l. This logic is predicated on the fact that
// differences will likely be much smaller than the original vector.
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the ContainsPntr method on l and r. In big-O it might look something like
// this, O(O(r.ContainsPntr)*O(l.ContainsPntr)), where O(r.ContainsPntr)
// represents the time complexity of the containsPntr method on r and
// O(l.ContainsPntr) represents the time complexity of the containsPntr method
// on l.
func (v *Vector[T, U]) Difference(
	l containerTypes.ComparisonsOtherConstraint[T],
	r containerTypes.ComparisonsOtherConstraint[T],
) {
	newV := make(Vector[T, U], 0, l.Length()/2)
	addressableSafeValIter[T](l).ForEach(
		func(index int, val *T) (iter.IteratorFeedback, error) {
			if !r.ContainsPntr(val) {
				newV = append(newV, *val)
			}
			return iter.Continue, nil
		},
	)
	v.Clear()
	*v = newV
}

// Description: Places a write lock on the underlying vector and then calls the
// underlying vectors [Vector.Difference] method. Attempts to place a read lock
// on l and r but whether or not that happens is implementation dependent.
//
// Lock Type: Write on this vector, read on l and r
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the ContainsPntr method on l and r. In big-O it might look something like
// this, O(O(r.ContainsPntr)*O(l.ContainsPntr)), where O(r.ContainsPntr)
// represents the time complexity of the containsPntr method on r and
// O(l.ContainsPntr) represents the time complexity of the containsPntr method
// on l.
func (v *SyncedVector[T, U]) Difference(
	l containerTypes.ComparisonsOtherConstraint[T],
	r containerTypes.ComparisonsOtherConstraint[T],
) {
	r.RLock()
	l.RLock()
	v.Lock()
	defer r.RUnlock()
	defer l.RUnlock()
	defer v.Unlock()
	v.Vector.Difference(l, r)
}

// Description: Returns true if this vector is a superset to other.
//
// Time Complexity: O(n*m), where n is the number of values in this vector and
// m is the number of values in other.
func (v *Vector[T, U]) IsSuperset(
	other containerTypes.ComparisonsOtherConstraint[T],
) bool {
	rv := (len(*v) >= other.Length())
	if !rv {
		return false
	}
	addressableSafeValIter[T](other).ForEach(
		func(index int, val *T) (iter.IteratorFeedback, error) {
			if rv = v.ContainsPntr(val); !rv {
				return iter.Break, nil
			}
			return iter.Continue, nil
		},
	)
	return rv
}

// Description: Places a read lock on the underlying vector and then calls the
// underlying vectors [Vector.IsSuperset] method. Attempts to place a read lock
// on other but whether or not that happens is implementation dependent.
//
// Lock Type: Read on this vector, read on other
//
// Time Complexity: O(n*m), where n is the number of values in this vector and
// m is the number of values in other.
func (v *SyncedVector[T, U]) IsSuperset(
	other containerTypes.ComparisonsOtherConstraint[T],
) bool {
	v.RLock()
	other.RLock()
	defer v.RUnlock()
	defer other.RUnlock()
	return v.Vector.IsSuperset(other)
}

// Description: Returns true if this vector is a subset to other.
//
// Time Complexity: Dependent on the ContainsPntr method of other. In big-O
// terms it may look somwthing like this: O(n*O(other.ContainsPntr)), where n is
// the number of elements in the current vector and other.ContainsPntr
// represents the time complexity of the containsPntr method on other.
func (v *Vector[T, U]) IsSubset(
	other containerTypes.ComparisonsOtherConstraint[T],
) bool {
	rv := (len(*v) <= other.Length())
	for i := 0; i < len(*v) && rv; i++ {
		rv = other.ContainsPntr(&(*v)[i])
	}
	return rv
}

// Description: Places a read lock on the underlying vector and then calls the
// underlying vectors [Vector.IsSubset] method. Attempts to place a read lock on
// other but whether or not that happens is implementation dependent.
//
// Lock Type: Read on this vector, read on other
//
// Time Complexity: Dependent on the ContainsPntr method of other. In big-O
// terms it may look somwthing like this: O(n*O(other.ContainsPntr)), where n is
// the number of elements in the current vector and other.ContainsPntr
// represents the time complexity of the containsPntr method on other.
func (v *SyncedVector[T, U]) IsSubset(
	other containerTypes.ComparisonsOtherConstraint[T],
) bool {
	v.RLock()
	other.RLock()
	defer v.RUnlock()
	defer other.RUnlock()
	return v.Vector.IsSubset(other)
}

// An equality function that implements the [widget.Base]
// interface. Internally this is equivalent to [Vector.KeyedEq]. Returns true
// if l==r, false otherwise.
func (_ *Vector[T, U]) Eq(l *Vector[T, U], r *Vector[T, U]) bool {
	return l.KeyedEq(r)
}

// An equality function that implements the [widget.Base]
// interface. Internally this is equivalent to [SyncedVector.KeyedEq]. Returns
// true if l==r, false otherwise.
func (_ *SyncedVector[T, U]) Eq(l *SyncedVector[T, U], r *SyncedVector[T, U]) bool {
	return l.KeyedEq(r)
}

// A function that returns a hash of a vector to implement the
// [widget.Base]. To do this all of the individual hashes that
// are produced from the elements of the vector are combined in a way that
// maintains identity, making it so the hash will represent the same equality
// operation that [Vector.KeyedEq] and [Vector.Eq] provide.
func (_ *Vector[T, U]) Hash(other *Vector[T, U]) hash.Hash {
	var rv hash.Hash
	w := widgets.Base[T, U]{}
	if len(*other) > 0 {
		rv = w.Hash(&(*other)[0])
		for i := 1; i < len(*other); i++ {
			rv = rv.Combine(w.Hash(&(*other)[i]))
		}
	}
	return rv
}

// Places a read lock on the underlying vector of other and then calls others
// underlying vector [Vector.Hash] method. Implements the
// [widget.Base].
func (_ *SyncedVector[T, U]) Hash(other *SyncedVector[T, U]) hash.Hash {
	other.RLock()
	defer other.RUnlock()
	return other.Vector.Hash(&other.Vector)
}

// An zero function that implements the [widget.Base] interface.
// Internally this is equivalent to [Vector.Clear].
func (_ *Vector[T, U]) Zero(other *Vector[T, U]) {
	other.Clear()
}

// An zero function that implements the [widget.Base] interface.
// Internally this is equivalent to [SyncedVector.Clear].
func (_ *SyncedVector[T, U]) Zero(other *SyncedVector[T, U]) {
	other.Clear()
}
