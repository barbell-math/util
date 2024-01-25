package dataStruct

import (
	"fmt"
	"sync"

	"github.com/barbell-math/util/algo/iter"
	customerr "github.com/barbell-math/util/err"
)


type (
    // A type to represent an array that dynamically grows as elements are added.
    // This is nothing more than a generically initialized slice with methods
    // attached to it so it can be passed to functions that use the interfaces
    // defined in the dataStruct/types, dataStruct/types/static, or
    // dataStruct/types/dynamic packages.
    Vector[T any] []T
    
    // A synchronized version of Vector. All operations will be wrapped in the
    // appropriate calls the embedded RWMutex. A pointer to a RWMutex is embedded
    // rather than a value to avoid copying the lock value.
    SyncedVector[T any] struct {
        *sync.RWMutex
        Vector[T]
    }
)

// Creates a new vector initialized with size zero valued elements. Size must be
// >= 0, an error will be returned if it is not. If size is 0 the vector will be 
// initialized with 0 elements. A vector can also be created by type casting a
// standard slice.
func NewVector[T any](size int) (Vector[T],error) {
    if size<0 {
        return nil,customerr.ValOutsideRange(
            fmt.Sprintf("Size of vector must be >=0 | Have: %d",size),
        );
    }
    return Vector[T](make(Vector[T],size)),nil
}

// Creates a new synced vector initialized with size zero valued elements. Size 
// must be >= 0, an error will be returned if it is not. If size is 0 the vector 
// will be initialized with 0 elements. The underlying RWMutex value will be 
// fully unlocked upon initialization.
func NewSyncedVector[T any](size int) (SyncedVector[T],error) {
    rv,err:=NewVector[T](size)
    return SyncedVector[T]{
        Vector: rv,
        RWMutex: &sync.RWMutex{},
    }, err
}

// A empty pass through function that performs no action. Vector will call all 
// the appropriate locking methods despite not being synced, just nothing will
// happen. This is done so that SyncedVector can simply embed a Vector and
// override the appropriate locking methods to implement the correct behavior
// without needing to make any additional changes such as wrapping every single
// method from Vector.
func (v *Vector[T])Lock() { }

// A empty pass through function that performs no action. Vector will call all 
// the appropriate locking methods despite not being synced, just nothing will
// happen. This is done so that SyncedVector can simply embed a Vector and
// override the appropriate locking methods to implement the correct behavior
// without needing to make any additional changes such as wrapping every single
// method from Vector.
func (v *Vector[T])Unlock() { }

// A empty pass through function that performs no action. Vector will call all 
// the appropriate locking methods despite not being synced, just nothing will
// happen. This is done so that SyncedVector can simply embed a Vector and
// override the appropriate locking methods to implement the correct behavior
// without needing to make any additional changes such as wrapping every single
// method from Vector.
func (v *Vector[T])RLock() { }

// A empty pass through function that performs no action. Vector will call all 
// the appropriate locking methods despite not being synced, just nothing will
// happen. This is done so that SyncedVector can simply embed a Vector and
// override the appropriate locking methods to implement the correct behavior
// without needing to make any additional changes such as wrapping every single
// method from Vector.
func (v *Vector[T])RUnlock() { }

// The SyncedVector method to override the Vector pass through function and 
// actually apply the mutex operation.
func (v *SyncedVector[T])Lock() { v.RWMutex.Lock() }

// The SyncedVector method to override the Vector pass through function and 
// actually apply the mutex operation.
func (v *SyncedVector[T])Unlock() { v.RWMutex.Unlock() }

// The SyncedVector method to override the Vector pass through function and 
// actually apply the mutex operation.
func (v *SyncedVector[T])RLock() { v.RWMutex.RLock() }

// The SyncedVector method to override the Vector pass through function and 
// actually apply the mutex operation.
func (v *SyncedVector[T])RUnlock() { v.RWMutex.RUnlock() }

// Returns the length of the vector.
func (v *Vector[T])Length() int {
    v.RLock()
    defer v.RUnlock()
    return len(*v)
}

// Returns the capacity of the vector.
func (v *Vector[T])Capacity() int {
    v.RLock()
    defer v.RUnlock()
    return cap(*v)
}

// Sets the capacity of the underlying slice. If the new capacity is less than
// the old capacity then values at the end of the list will be dropped. Performs
// a copy operations, making the time complexity O(N).
func (v *Vector[T])SetCapacity(c int) error {
    v.Lock()
    defer v.Unlock()
    tmp:=make(Vector[T],c)
    copy(tmp,*v)
    *v=tmp
    return nil
}

// Gets the value at the specified index. Returns an error if the index is 
// >= the length of the vector.
func (v *Vector[T])Get(idx int) (T,error){
    if _v,err:=v.GetPntr(idx); err==nil {
        return *_v,nil
    } else {
        var tmp T
        return tmp,err
    }
}

// Gets a pointer to the value at the specified index. Returns an error if the 
// index is >= the length of the vector.
func (v *Vector[T])GetPntr(idx int) (*T,error){
    v.RLock()
    defer v.RUnlock()
    if idx>=0 && idx<len(*v) && len(*v)>0 {
        return &(*v)[idx],nil
    }
    return nil,getIndexOutOfBoundsError(idx,len(*v))
}

// Sets the value at the specified index. Returns an error if the index is >= the
// length of the vector.
func (v *Vector[T])Set(idx int, val T) error {
    v.Lock()
    defer v.Unlock()
    if idx>=0 && idx<len(*v) && len(*v)>0 {
        (*v)[idx]=val
        return nil
    }
    return getIndexOutOfBoundsError(idx,len(*v))
}

// Appends the supplied values to the vector. This function will never return
// an error.
func (v *Vector[T])Append(vals ...T) error {
    v.Lock()
    defer v.Unlock()
    *v=append(*v, vals...)
    return nil
}

// Inserts the supplied values at the given index. Returns an error if the index
// is >= the length of the vector.
// For time complexity see the InsertVector section of:
// https://go.dev/wiki/SliceTricks
func (v *Vector[T])Insert(idx int, vals ...T) error {
    v.Lock()
    defer v.Unlock()
    if idx>=0 && idx<len(*v) && len(*v)>0 {
        *v=append((*v)[:idx],append(vals,(*v)[idx:]...)...) 
        return nil
    } else if idx==len(*v) {
        *v=append(*v,vals...)
        return nil
    }
    return getIndexOutOfBoundsError(idx,len(*v))
}

// Deletes the value at the specified index. Returns an error if the index is 
// >= the length of the vector.
func (v *Vector[T])Delete(idx int) error {
    v.Lock()
    defer v.Unlock()
    if idx<0 || idx>=len(*v) {
        return getIndexOutOfBoundsError(idx,len(*v))
    } else if idx>=0 && idx<len(*v) && len(*v)>0 {
        *v=append((*v)[:idx],(*v)[idx+1:]...)
    }
    return nil
}

// Clears all values from the vector. Equivalent to making a new vector and
// setting it equal to the current one.
func (v *Vector[T])Clear() {
    v.Lock()
    defer v.Unlock()
    *v=make(Vector[T], 0)
}

// Returns the value at index 0 if one is present. If the vector has no elements
// then an error is returned.
func (v *Vector[T])PeekFront() (T,error) {
    v.RLock()
    defer v.RUnlock()
    if _v,err:=v.PeekPntrFront(); err==nil {
        return *_v,err
    } else {
        var tmp T
        return tmp,err
    }
}

// Returns a pointer to the value at index 0 if one is present. If the vector 
// has no elements then an error is returned.
func (v *Vector[T])PeekPntrFront() (*T,error) {
    v.RLock()
    defer v.RUnlock()
    if len(*v)>0 {
        return &(*v)[0],nil
    }
    return nil,getIndexOutOfBoundsError(0,len(*v))
}

// Returns the value at index len(v)-1 if one is present. If the vector has no 
// elements then an error is returned.
func (v *Vector[T])PeekBack() (T,error) {
    v.RLock()
    defer v.RUnlock()
    if _v,err:=v.PeekPntrBack(); err==nil {
        return *_v,err
    } else {
        var tmp T
        return tmp,err
    }
}

// Returns a pointer to the value at index len(v)-1 if one is present. If the 
// vector has no elements then an error is returned.
func (v *Vector[T])PeekPntrBack() (*T,error) {
    v.RLock()
    defer v.RUnlock()
    if len(*v)>0 {
        return &(*v)[len(*v)-1],nil
    }
    return nil,getIndexOutOfBoundsError(0,len(*v))
}

// Returns and removes the element at the front of the vector. Returns an error
// if the vector has no elements.
func (v *Vector[T])PopFront() (T,error) {
    v.Lock()
    defer v.Unlock()
    if len(*v)>0 {
        rv:=(*v)[0]
        *v=(*v)[1:]
        return rv,nil
    }
    var tmp T
    return tmp,Empty("Nothing to pop!")
}

// Returns and removes the element at the back of the vector. Returns an error
// if the vector has no elements.
func (v *Vector[T])PopBack() (T,error) {
    v.Lock()
    defer v.Unlock()
    if len(*v)>0 {
        rv:=(*v)[len(*v)-1]
        *v=(*v)[:len(*v)-1]
        return rv,nil
    }
    var tmp T
    return tmp,Empty("Nothing to pop!")
}

// Pushes an element to the back of the vector. Equivalent to appending a single
// value to the end of the vector.
func (v *Vector[T])PushBack(val T) error {
    v.Lock()
    defer v.Unlock()
    *v=append(*v, val)
    return nil
}

// Pushes an element to the front of the vector. Equivalent to inserting a single
// value at the front of the vector.
func (v *Vector[T])PushFront(val T) error {
    v.Lock()
    defer v.Unlock()
    *v=append(Vector[T]{val}, (*v)...)
    return nil
}

// Pushes an element to the back of the vector. Equivalent to appending a single
// value to the end of the vector. Has the same behavior as PushBack because
// the underlying vector grows as needed.
func (v *Vector[T])ForcePushBack(val T) {
    v.Lock()
    defer v.Unlock()
    *v=append(*v, val)
}

// Pushes an element to the front of the vector. Equivalent to inserting a single
// value at the front of the vector. Has the same behavior as PushBack because
// the underlying vector grows as needed.
func (v *Vector[T])ForcePushFront(val T) {
    v.Lock()
    defer v.Unlock()
    *v=append(Vector[T]{val}, (*v)...)
}

// Returns an iterator that iterates over the values in the vector. The vector
// will have a read lock the entire time the iteration is being performed. The
// lock will not be applied until the iterator is consumed.
func (v *Vector[T])Elems() iter.Iter[T] {
    return iter.SequentialElems[T](
        len(*v),
        func(i int) (T, error) { return (*v)[i],nil },
    ).SetupTeardown(
        func() error { v.RLock(); return nil },
        func() error { v.RUnlock(); return nil },
    )
}

// Returns an iterator that iterates over the pointers to ithe values in the 
// vector. The vector will have a read lock the entire time the iteration is 
// being performed. The lock will not be applied until the iterator is consumed.
func (v *Vector[T])PntrElems() iter.Iter[*T] {
    return iter.SequentialElems[*T](
        len(*v),
        func(i int) (*T, error) { return &(*v)[i],nil },
    ).SetupTeardown(
        func() error { v.RLock(); return nil },
        func() error { v.RUnlock(); return nil },
    )
}

// Returns true if the vectors are equal. The supplied comparison function will
// be used when comparing values in the vector.
func (v *Vector[T])Eq(other Vector[T], comp func(l *T, r *T) bool) bool {
    v.RLock()
    defer v.RUnlock()
    rv:=(len(*v)==len(other))
    for i:=0; i<len(other) && rv; i++ {
        rv=(rv && comp(&(*v)[i],&other[i]))
    }
    return rv
}

// Returns true if the vectors are not equal. The supplied comparison function 
// will be used when comparing values in the vector.
func (v *Vector[T])Neq(other Vector[T], comp func(l *T, r *T) bool) bool {
    return !v.Eq(other,comp)
}
