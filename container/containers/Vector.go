package containers

import (
	"sync"

	"github.com/barbell-math/util/customerr"
)


type (
    // A type to represent an array that dynamically grows as elements are added.
    // This is nothing more than a generically initialized slice with methods
    // attached to it so it can be passed to functions that use the interfaces
    // defined in the [containerTypes], [staticContainers], or
    // [dynamicContainers] packages. The type constraints on the generics
    // define the logic for how equality comparisons will be handled.
    Vector[T any, U any , CONSTRAINT WidgetConstraint[T,U]] []T

    // A synchronized version of Vector. All operations will be wrapped in the
    // appropriate calls the embedded RWMutex. A pointer to a RWMutex is embedded
    // rather than a value to avoid copying the lock value.
    SyncedVector[T any, U any, CONSTRAINT WidgetConstraint[T,U]] struct {
    	*sync.RWMutex
    	Vector[T,U,CONSTRAINT]
    }
)

// Type casts the given slice to a vector. There is nothing special happening
// with this function, a slice can be cast on it's own outside of this function
// just fine, this function just saves some typing.
func SliceToVector[T any, U any, CONSTRAINT WidgetConstraint[T,U]](
    s []T,
) Vector[T,U,CONSTRAINT] {
    return Vector[T, U, CONSTRAINT](s)
}

// Creates a new vector initialized with size zero valued elements. Size must be
// >= 0, an error will be returned if it is not. If size is 0 the vector will be 
// initialized with 0 elements. A vector can also be created by type casting a
// standard slice.
func NewVector[T any, U any, CONSTRAINT WidgetConstraint[T,U]](
    size int,
) (Vector[T,U,CONSTRAINT],error) {
    if size<0 {
	 return Vector[T,U,CONSTRAINT]{}, customerr.Wrap(
	    customerr.ValOutsideRange,
	    "Size must be >=0. Got: %d",size,
    	)	
    }
    return make(Vector[T,U,CONSTRAINT],size),nil
}

// Creates a new synced vector initialized with size zero valued elements. Size 
// must be >= 0, an error will be returned if it is not. If size is 0 the vector 
// will be initialized with 0 elements. The underlying RWMutex value will be 
// fully unlocked upon initialization.
func NewSyncedVector[T any, U any, CONSTRAINT WidgetConstraint[T,U]](
    size int,
) (SyncedVector[T,U,CONSTRAINT],error) {
    rv,err:=NewVector[T,U,CONSTRAINT](size)
    return SyncedVector[T,U,CONSTRAINT]{
	RWMutex: &sync.RWMutex{},
	Vector: rv,
    },err
}

// A empty pass through function that performs no action. Vector will call all 
// the appropriate locking methods despite not being synced, just nothing will
// happen. This is done so that SyncedVector can simply embed a Vector and
// override the appropriate locking methods to implement the correct behavior
// without needing to make any additional changes such as wrapping every single
// method from Vector.
func (v *Vector[T,U,CONSTRAINT])Lock() { }

// A empty pass through function that performs no action. Vector will call all 
// the appropriate locking methods despite not being synced, just nothing will
// happen. This is done so that SyncedVector can simply embed a Vector and
// override the appropriate locking methods to implement the correct behavior
// without needing to make any additional changes such as wrapping every single
// method from Vector.
func (v *Vector[T,U,CONSTRAINT])Unlock() { }

// A empty pass through function that performs no action. Vector will call all 
// the appropriate locking methods despite not being synced, just nothing will
// happen. This is done so that SyncedVector can simply embed a Vector and
// override the appropriate locking methods to implement the correct behavior
// without needing to make any additional changes such as wrapping every single
// method from Vector.
func (v *Vector[T,U,CONSTRAINT])RLock() { }

// A empty pass through function that performs no action. Vector will call all 
// the appropriate locking methods despite not being synced, just nothing will
// happen. This is done so that SyncedVector can simply embed a Vector and
// override the appropriate locking methods to implement the correct behavior
// without needing to make any additional changes such as wrapping every single
// method from Vector.
func (v *Vector[T,U,CONSTRAINT])RUnlock() { }

// The SyncedVector method to override the Vector pass through function and 
// actually apply the mutex operation.
func (v *SyncedVector[T,U,CONSTRAINT])Lock() { v.RWMutex.Lock() }

// The SyncedVector method to override the Vector pass through function and 
// actually apply the mutex operation.
func (v *SyncedVector[T,U,CONSTRAINT])Unlock() { v.RWMutex.Unlock() }

// The SyncedVector method to override the Vector pass through function and 
// actually apply the mutex operation.
func (v *SyncedVector[T,U,CONSTRAINT])RLock() { v.RWMutex.RLock() }

// The SyncedVector method to override the Vector pass through function and 
// actually apply the mutex operation.
func (v *SyncedVector[T,U,CONSTRAINT])RUnlock() { v.RWMutex.RUnlock() }

// Returns the length of the vector.
func (v *Vector[T,U,CONSTRAINT])Length() int {
    v.RLock()
    defer v.RUnlock()
    return len(*v)
}

// Returns the capacity of the vector.
func (v *Vector[T,U,CONSTRAINT])Capacity() int {
    v.RLock()
    defer v.RUnlock()
    return cap(*v)
}

// Sets the capacity of the underlying slice. If the new capacity is less than
// the old capacity then values at the end of the list will be dropped. Performs
// a copy operations, making the time complexity O(N).
func (v *Vector[T,U,CONSTRAINT])SetCapacity(c int) error {
    v.Lock()
    defer v.Unlock()
    tmp:=make(Vector[T,U,CONSTRAINT],c)
    copy(tmp,*v)
    *v=tmp
    return nil
}

// Gets the value at the specified index. Returns an error if the index is 
// >= the length of the vector.
func (v *Vector[T,U,CONSTRAINT])Get(idx int) (T,error){
    if _v,err:=v.GetPntr(idx); err==nil {
        return *_v,nil
    } else {
        var tmp T
        return tmp,err
    }
}

// Gets a pointer to the value at the specified index. Returns an error if the 
// index is >= the length of the vector.
func (v *Vector[T,U,CONSTRAINT])GetPntr(idx int) (*T,error){
    v.RLock()
    defer v.RUnlock()
    if idx>=0 && idx<len(*v) && len(*v)>0 {
        return &(*v)[idx],nil
    }
    return nil,getIndexOutOfBoundsError(idx,0,len(*v))
}

// Emplace (sets) the value at the specified index. Returns an error if the 
// index is >= the length of the vector.
func (v *Vector[T,U,CONSTRAINT])Emplace(idx int, val T) error {
    v.Lock()
    defer v.Unlock()
    if idx>=0 && idx<len(*v) && len(*v)>0 {
        (*v)[idx]=val
        return nil
    }
    return getIndexOutOfBoundsError(idx,0,len(*v))
}

// Appends the supplied values to the vector. This function will never return
// an error.
func (v *Vector[T,U,CONSTRAINT])Append(vals ...T) error {
    v.Lock()
    defer v.Unlock()
    *v=append(*v, vals...)
    return nil
}

// Pushes (inserts)the supplied values at the given index. Returns an error if 
// the index is >= the length of the vector.
// For time complexity see the InsertVector section of:
// https://go.dev/wiki/SliceTricks
func (v *Vector[T,U,CONSTRAINT])Push(idx int, vals ...T) error {
    v.Lock()
    defer v.Unlock()
    if idx>=0 && idx<len(*v) && len(*v)>0 {
        *v=append((*v)[:idx],append(vals,(*v)[idx:]...)...) 
        return nil
    } else if idx==len(*v) {
        *v=append(*v,vals...)
        return nil
    }
    return getIndexOutOfBoundsError(idx,0,len(*v))
}

// Deletes the value at the specified index. Returns an error if the index is 
// >= the length of the vector.
func (v *Vector[T,U,CONSTRAINT])Delete(idx int) error {
    v.Lock()
    defer v.Unlock()
    if idx<0 || idx>=len(*v) {
        return getIndexOutOfBoundsError(idx,0,len(*v))
    } else if idx>=0 && idx<len(*v) && len(*v)>0 {
        *v=append((*v)[:idx],(*v)[idx+1:]...)
    }
    return nil
}

// Clears all values from the vector. Equivalent to making a new vector and
// setting it equal to the current one.
func (v *Vector[T,U,CONSTRAINT])Clear() {
    v.Lock()
    defer v.Unlock()
    *v=make(Vector[T,U,CONSTRAINT], 0)
}

// Returns the value at index 0 if one is present. If the vector has no elements
// then an error is returned.
func (v *Vector[T,U,CONSTRAINT])PeekFront() (T,error) {
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
func (v *Vector[T,U,CONSTRAINT])PeekPntrFront() (*T,error) {
    v.RLock()
    defer v.RUnlock()
    if len(*v)>0 {
        return &(*v)[0],nil
    }
    return nil,getIndexOutOfBoundsError(0,0,len(*v))
}

// Returns the value at index len(v)-1 if one is present. If the vector has no 
// elements then an error is returned.
func (v *Vector[T,U,CONSTRAINT])PeekBack() (T,error) {
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
func (v *Vector[T,U,CONSTRAINT])PeekPntrBack() (*T,error) {
    v.RLock()
    defer v.RUnlock()
    if len(*v)>0 {
        return &(*v)[len(*v)-1],nil
    }
    return nil,getIndexOutOfBoundsError(0,0,len(*v))
}

// Returns and removes the element at the front of the vector. Returns an error
// if the vector has no elements.
func (v *Vector[T,U,CONSTRAINT])PopFront() (T,error) {
    v.Lock()
    defer v.Unlock()
    if len(*v)>0 {
        rv:=(*v)[0]
        *v=(*v)[1:]
        return rv,nil
    }
    var tmp T
    return tmp,customerr.Wrap(Empty,"Nothing to pop!")
}

// Returns and removes the element at the back of the vector. Returns an error
// if the vector has no elements.
func (v *Vector[T,U,CONSTRAINT])PopBack() (T,error) {
    v.Lock()
    defer v.Unlock()
    if len(*v)>0 {
        rv:=(*v)[len(*v)-1]
        *v=(*v)[:len(*v)-1]
        return rv,nil
    }
    var tmp T
    return tmp,customerr.Wrap(Empty,"Nothing to pop!")
}

// Pushes an element to the back of the vector. Equivalent to appending a single
// value to the end of the vector.
func (v *Vector[T,U,CONSTRAINT])PushBack(val T) error {
    v.Lock()
    defer v.Unlock()
    *v=append(*v, val)
    return nil
}

// Pushes an element to the front of the vector. Equivalent to inserting a single
// value at the front of the vector.
func (v *Vector[T,U,CONSTRAINT])PushFront(val T) error {
    v.Lock()
    defer v.Unlock()
    *v=append(Vector[T,U,CONSTRAINT]{val}, (*v)...)
    return nil
}

// Pushes an element to the back of the vector. Equivalent to appending a single
// value to the end of the vector. Has the same behavior as PushBack because
// the underlying vector grows as needed.
func (v *Vector[T,U,CONSTRAINT])ForcePushBack(val T) {
    v.Lock()
    defer v.Unlock()
    *v=append(*v, val)
}

// Pushes an element to the front of the vector. Equivalent to inserting a single
// value at the front of the vector. Has the same behavior as PushBack because
// the underlying vector grows as needed.
func (v *Vector[T,U,CONSTRAINT])ForcePushFront(val T) {
    v.Lock()
    defer v.Unlock()
    *v=append(Vector[T,U,CONSTRAINT]{val}, (*v)...)
}

// TODO - look into exposing iter through interface?
// // Returns an iterator that iterates over the values in the vector. The vector
// // will have a read lock the entire time the iteration is being performed. The
// // lock will not be applied until the iterator is consumed.
// func (v *Vector[T,U,CONSTRAINT])Elems() iter.Iter[T] {
//     return iter.SequentialElems[T](
//         len(*v),
//         func(i int) (T, error) { return (*v)[i],nil },
//     ).SetupTeardown(
//         func() error { v.RLock(); return nil },
//         func() error { v.RUnlock(); return nil },
//     )
// }
// 
// // Returns an iterator that iterates over the pointers to ithe values in the 
// // vector. The vector will have a read lock the entire time the iteration is 
// // being performed. The lock will not be applied until the iterator is consumed.
// func (v *Vector[T,U,CONSTRAINT])PntrElems() iter.Iter[*T] {
//     return iter.SequentialElems[*T](
//         len(*v),
//         func(i int) (*T, error) { return &(*v)[i],nil },
//     ).SetupTeardown(
//         func() error { v.RLock(); return nil },
//         func() error { v.RUnlock(); return nil },
//     )
// }

// TODO - use new generic types to implement equality
// // Returns true if the vectors are equal. The supplied comparison function will
// // be used when comparing values in the vector.
// func (v *Vector[T,U,CONSTRAINT])Eq(other *Vector[T,U,CONSTRAINT], comp func(l *T, r *T) bool) bool {
//     v.RLock()
//     defer v.RUnlock()
//     rv:=(len(*v)==len(*other))
//     for i:=0; i<len(*other) && rv; i++ {
//         rv=(rv && comp(&(*v)[i],&(*other)[i]))
//     }
//     return rv
// }
// 
// // Returns true if the vectors are not equal. The supplied comparison function 
// // will be used when comparing values in the vector.
// func (v *Vector[T,U,CONSTRAINT])Neq(
//     other *Vector[T,U,CONSTRAINT], 
//     comp func(l *T, r *T) bool,
// ) bool {
//     return !v.Eq(other,comp)
// }
