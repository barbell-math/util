package containers

import (
	"sync"

	"github.com/barbell-math/util/algo/hash"
	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/algo/widgets"
	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/customerr"
)

type (
    // A type to represent an array that dynamically grows as elements are added.
    // This is nothing more than a generically initialized slice with methods
    // attached to it so it can be passed to functions that use the interfaces
    // defined in the [containerTypes], [staticContainers], or
    // [dynamicContainers] packages. The type constraints on the generics
    // define the logic for how value specific operations, such as equality 
    // comparisons, will be handled.
    Vector[T any, U widgets.WidgetInterface[T]] []T

    // A synchronized version of Vector. All operations will be wrapped in the
    // appropriate calls the embedded RWMutex. A pointer to a RWMutex is embedded
    // rather than a value to avoid copying the lock value.
    SyncedVector[T any, U widgets.WidgetInterface[T]] struct {
    	*sync.RWMutex
    	Vector[T,U]
    }
)

// Creates a new vector initialized with size zero valued elements. Size must be
// >= 0, an error will be returned if it is not. If size is 0 the vector will be 
// initialized with 0 elements. A vector can also be created by type casting a
// standard slice, as shown below.
//
//  // Vector to slice.
//  v,_:=NewVector[string,builtinWidgets.BuiltinString](3)
//  s:=[]string(v)
//  // Slice to vector.
//  s2:=make([]string,4)
//  v2:=Vector[string,builtinWidgets.BuiltinString](s2)
//
// Note that by performing the above type casts the operations provided by the
// widget, including equality, are not preserved.
func NewVector[T any, U widgets.WidgetInterface[T]](size int) (Vector[T,U],error) {
    if size<0 {
        return Vector[T, U]{}, getSizeError(size)
    }
    return make(Vector[T,U],size),nil
}

// Creates a new synced vector initialized with size zero valued elements. Size 
// must be >= 0, an error will be returned if it is not. If size is 0 the vector 
// will be initialized with 0 elements. The underlying RWMutex value will be 
// fully unlocked upon initialization.
func NewSyncedVector[T any, U widgets.WidgetInterface[T]](
    size int,
) (SyncedVector[T,U],error) {
    rv,err:=NewVector[T,U](size)
    return SyncedVector[T,U]{
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
func (v *Vector[T,U])Lock() { }

// A empty pass through function that performs no action. Vector will call all 
// the appropriate locking methods despite not being synced, just nothing will
// happen. This is done so that SyncedVector can simply embed a Vector and
// override the appropriate locking methods to implement the correct behavior
// without needing to make any additional changes such as wrapping every single
// method from Vector.
func (v *Vector[T,U])Unlock() { }

// A empty pass through function that performs no action. Vector will call all 
// the appropriate locking methods despite not being synced, just nothing will
// happen. This is done so that SyncedVector can simply embed a Vector and
// override the appropriate locking methods to implement the correct behavior
// without needing to make any additional changes such as wrapping every single
// method from Vector.
func (v *Vector[T,U])RLock() { }

// A empty pass through function that performs no action. Vector will call all 
// the appropriate locking methods despite not being synced, just nothing will
// happen. This is done so that SyncedVector can simply embed a Vector and
// override the appropriate locking methods to implement the correct behavior
// without needing to make any additional changes such as wrapping every single
// method from Vector.
func (v *Vector[T,U])RUnlock() { }

// The SyncedVector method to override the Vector pass through function and 
// actually apply the mutex operation.
func (v *SyncedVector[T,U])Lock() { v.RWMutex.Lock() }

// The SyncedVector method to override the Vector pass through function and 
// actually apply the mutex operation.
func (v *SyncedVector[T,U])Unlock() { v.RWMutex.Unlock() }

// The SyncedVector method to override the Vector pass through function and 
// actually apply the mutex operation.
func (v *SyncedVector[T,U])RLock() { v.RWMutex.RLock() }

// The SyncedVector method to override the Vector pass through function and 
// actually apply the mutex operation.
func (v *SyncedVector[T,U])RUnlock() { v.RWMutex.RUnlock() }

// Returns the length of the vector.
func (v *Vector[T,U])Length() int {
    v.RLock()
    defer v.RUnlock()
    return len(*v)
}

// Returns the capacity of the vector.
func (v *Vector[T,U])Capacity() int {
    v.RLock()
    defer v.RUnlock()
    return cap(*v)
}

// Sets the capacity of the underlying slice. If the new capacity is less than
// the old capacity then values at the end of the list will be dropped. Performs
// a copy operations, making the time complexity O(N).
func (v *Vector[T,U])SetCapacity(c int) error {
    v.Lock()
    defer v.Unlock()
    tmp:=make(Vector[T,U],c)
    copy(tmp,*v)
    *v=tmp
    return nil
}

// Gets the value at the specified index. Returns an error if the index is 
// >= the length of the vector.
func (v *Vector[T,U])Get(idx int) (T,error){
    if _v,err:=v.GetPntr(idx); err==nil {
        return *_v,nil
    } else {
        var tmp T
        return tmp,err
    }
}

// Gets a pointer to the value at the specified index. Returns an error if the 
// index is >= the length of the vector.
func (v *Vector[T,U])GetPntr(idx int) (*T,error){
    v.RLock()
    defer v.RUnlock()
    if idx>=0 && idx<len(*v) && len(*v)>0 {
        return &(*v)[idx],nil
    }
    return nil,getIndexOutOfBoundsError(idx,0,len(*v))
}

// Contains will return true if the supplied value is in the vector, false
// otherwise. All equality comparisons are performed by the generic U widget
// type that the vector was initialized with.
func (v *Vector[T, U])Contains(val T) bool {
    return v.ContainsPntr(&val)
}

// ContainsPntr will return true if the supplied value is in the vector, false
// otherwise. All equality comparisons are performed by the generic U widget
// type that the vector was initialized with.
func (v *Vector[T, U])ContainsPntr(val *T) bool {
    v.RLock()
    defer v.RUnlock()
    found:=false
    w:=widgets.NewWidget[T,U]()
    for i:=0; i<len(*v) && !found; i++ {
        found=w.Eq(val,&(*v)[i])
    }
    return found
}

// KeyOf will return the index of the first occurrence of the supplied value
// in the vector. If the value is not found then the returned index will be -1
// and the boolean flag will be set to false. If the value is found then the
// boolean flag will be set to true. All equality comparisons are performed by 
// the generic U widget type that the vector was initialized with.
func (v *Vector[T, U])KeyOf(val T) (int,bool) {
    v.RLock()
    defer v.RUnlock()
    rv:=-1
    found:=false
    w:=widgets.NewWidget[T,U]()
    for i:=0; i<len(*v) && !found; i++ {
        if found=w.Eq(&val,&(*v)[i]); found {
            rv=i
        }
    }
    return rv,found
}

// Sets the value at the specified index. Returns an error if the index is >= 
// the length of the vector.
func (v *Vector[T,U])Set(vals ...basic.Pair[int,T]) error {
    v.Lock()
    defer v.Unlock()
    for _,iterV:=range(vals) {
        if iterV.A>=0 && iterV.A<len(*v) && len(*v)>0 {
            (*v)[iterV.A]=iterV.B
        } else {
            return getIndexOutOfBoundsError(iterV.A,0,len(*v))
        }
    }
    return nil
}

// Sets the supplied values sequentially starting at the supplied index and
// continuing sequentailly after that. Returns and error if any index that is
// attempted to be set is >= the length of the vector. If an error occurs, all 
// values will be set up until the value that caused the error.
func (v *Vector[T,U])SetSequential(idx int, vals ...T) error {
    v.Lock()
    defer v.Unlock()
    if idx>=len(*v) {
        return getIndexOutOfBoundsError(idx,0,len(*v))
    }
    numCopyableVals:=min(len(*v)-idx,len(vals))
    copy((*v)[idx:idx+numCopyableVals],vals[0:numCopyableVals])
    if idx+len(vals)>len(*v) {
        return getIndexOutOfBoundsError(len(*v),0,len(*v)) 
    }
    return nil
}

// Append the supplied values to the vector. This function will never return
// an error.
func (v *Vector[T,U])Append(vals ...T) error {
    v.Lock()
    defer v.Unlock()
    *v=append(*v, vals...)
    return nil
}

// AppendUnique will append the supplied values to the vector if they are not
// already present in the vector (unique). Non-unique values will not be 
// appended. This function will never return an error. The time complexity of 
// AppendUnique is O(n*m) where n is the number of values in the vector and m 
// is the number of values to append. For a more efficient implementation of 
// this method use a different container, such as [Set].
func (v *Vector[T,U])AppendUnique(vals ...T) error {
    v.Lock()
    defer v.Unlock()
    found:=false
    w:=widgets.NewWidget[T,U]()
    for _,iterV:=range(vals) {
        for j:=0; j<len(*v) && !found; j++ {
            found=w.Eq(&iterV,&(*v)[j])
        }
        if !found {
            *v=append(*v,iterV)
        }
    }
    return nil
}

// Insert will insert the supplied values into the vector. The values will be
// inserted in the order that they are given. The time complexity of this
// insert method is O(n^2)
func (v *Vector[T, U])Insert(vals ...basic.Pair[int,T]) error {
    v.Lock()
    defer v.Unlock()
    for i:=0; i<len(vals); i++ {
        if vals[i].A>=0 && vals[i].A<len(*v) && len(*v)>0 {
            var tmp T
            *v=append(*v, tmp)
            copy((*v)[vals[i].A+1:], (*v)[vals[i].A:])
            (*v)[vals[i].A] = vals[i].B
        } else if vals[i].A==len(*v) {
            *v=append(*v,vals[i].B)
        } else {
            return getIndexOutOfBoundsError(vals[i].A,0,len(*v))
        }
    }
    return nil
}

// Inserts the supplied values at the given index. Returns an error if the index
// is >= the length of the vector.
// For time complexity see the InsertVector section of:
// https://go.dev/wiki/SliceTricks
func (v *Vector[T,U])InsertSequential(idx int, vals ...T) error {
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

// Pop will remove the first num occurrences of val in the vector. All equality 
// comparisons are performed by the generic U widget type that the vector was 
// initialized with. If num is <=0 then no values will be poped and the vector
// will not change.
func (v *Vector[T, U])Pop(val T, num int) int {
    if num<=0 {
        return 0
    }
    v.Lock()
    defer v.Unlock()
    cntr:=0
    prevIndex:=-1
    w:=widgets.NewWidget[T,U]()
    for i:=0; i<len(*v); i++ {
        if w.Eq(&val,&(*v)[i]) && cntr+1<=num {
            if prevIndex==-1 {  // Initial value found
                prevIndex=i
            } else {
                copy((*v)[prevIndex-cntr+1:i],(*v)[prevIndex+1:i])
                prevIndex=i
            }
            cntr++
            if cntr>=num {
                break
            }
        }
    }
    if prevIndex!=-1 {
        copy((*v)[prevIndex-cntr+1:len(*v)],(*v)[prevIndex+1:len(*v)])
    }
    *v=(*v)[:len(*v)-cntr]
    return cntr
}

// Deletes the value at the specified index. Returns an error if the index is 
// >= the length of the vector.
func (v *Vector[T,U])Delete(idx int) error {
    v.Lock()
    defer v.Unlock()
    if idx<0 || idx>=len(*v) {
        return getIndexOutOfBoundsError(idx,0,len(*v))
    } else if idx>=0 && idx<len(*v) && len(*v)>0 {
        w:=widgets.NewWidget[T,U]()
        w.Zero(&(*v)[idx])
        *v=append((*v)[:idx],(*v)[idx+1:]...)
    }
    return nil
}

// Clears all values from the vector. Equivalent to making a new vector and
// setting it equal to the current one.
func (v *Vector[T,U])Clear() {
    v.Lock()
    defer v.Unlock()
    w:=widgets.NewWidget[T,U]()
    for i:=0; i<len(*v); i++ {
        w.Zero(&(*v)[i])
    }
    *v=make(Vector[T,U], 0)
}

// Returns the value at index 0 if one is present. If the vector has no elements
// then an error is returned.
func (v *Vector[T,U])PeekFront() (T,error) {
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
func (v *Vector[T,U])PeekPntrFront() (*T,error) {
    v.RLock()
    defer v.RUnlock()
    if len(*v)>0 {
        return &(*v)[0],nil
    }
    return nil,getIndexOutOfBoundsError(0,0,len(*v))
}

// Returns the value at index len(v)-1 if one is present. If the vector has no 
// elements then an error is returned.
func (v *Vector[T,U])PeekBack() (T,error) {
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
func (v *Vector[T,U])PeekPntrBack() (*T,error) {
    v.RLock()
    defer v.RUnlock()
    if len(*v)>0 {
        return &(*v)[len(*v)-1],nil
    }
    return nil,getIndexOutOfBoundsError(0,0,len(*v))
}

// Returns and removes the element at the front of the vector. Returns an error
// if the vector has no elements.
func (v *Vector[T,U])PopFront() (T,error) {
    v.Lock()
    defer v.Unlock()
    if len(*v)>0 {
        rv:=(*v)[0]
        *v=(*v)[1:]
        return rv,nil
    }
    var tmp T
    return tmp,customerr.Wrap(containerTypes.Empty,"Nothing to pop!")
}

// Returns and removes the element at the back of the vector. Returns an error
// if the vector has no elements.
func (v *Vector[T,U])PopBack() (T,error) {
    v.Lock()
    defer v.Unlock()
    if len(*v)>0 {
        rv:=(*v)[len(*v)-1]
        *v=(*v)[:len(*v)-1]
        return rv,nil
    }
    var tmp T
    return tmp,customerr.Wrap(containerTypes.Empty,"Nothing to pop!")
}

// Pushes an element to the back of the vector. Equivalent to appending a single
// value to the end of the vector. Values will be pushed back in the order that
// they are given. For example, calling push back on [0,1,2] with vals of [3,4]
// will result in [0,1,2,3,4].
func (v *Vector[T,U])PushBack(vals ...T) error {
    v.Lock()
    defer v.Unlock()
    *v=append(*v, vals...)
    return nil
}

// Pushes an element to the front of the vector. Equivalent to inserting a single
// value at the front of the vector. Values will be pushed to the front in the 
// order that they are given. For example, calling push front on [0,1,2] with 
// vals of [3,4] will result in [3,4,0,1,2].
func (v *Vector[T,U])PushFront(vals ...T) error {
    v.Lock()
    defer v.Unlock()
    *v=append(vals, (*v)...)
    return nil
}

// Pushes an element to the back of the vector. Equivalent to appending a single
// value to the end of the vector. Has the same behavior as PushBack because
// the underlying vector grows as needed.
func (v *Vector[T,U])ForcePushBack(vals ...T) {
    v.Lock()
    defer v.Unlock()
    *v=append(*v, vals...)
}

// Pushes an element to the front of the vector. Equivalent to inserting a single
// value at the front of the vector. Has the same behavior as PushBack because
// the underlying vector grows as needed.
func (v *Vector[T,U])ForcePushFront(vals ...T) {
    v.Lock()
    defer v.Unlock()
    *v=append(vals, (*v)...)
}

// Returns an iterator that iterates over the values in the vector. The vector
// will have a read lock the entire time the iteration is being performed. The
// lock will not be applied until the iterator is consumed.
func (v *Vector[T,U])Vals() iter.Iter[T] {
    return iter.SequentialElems[T](
        len(*v),
        func(i int) (T, error) { return (*v)[i],nil },
    ).SetupTeardown(
        func() error { v.RLock(); return nil },
        func() error { v.RUnlock(); return nil },
    )
}

// Returns an iterator that iterates over the pointers to the values in the 
// vector. The vector will have a read lock the entire time the iteration is 
// being performed. The lock will not be applied until the iterator is consumed.
func (v *Vector[T,U])ValPntrs() iter.Iter[*T] {
    return iter.SequentialElems[*T](
        len(*v),
        func(i int) (*T, error) { return &(*v)[i],nil },
    ).SetupTeardown(
        func() error { v.RLock(); return nil },
        func() error { v.RUnlock(); return nil },
    )
}

// Returns an iterator that iterates over the keys (indexes) of the vector. The
// vector will have a read lock the entire time the iteration is being performed.
// The lock will not be applied until the iterator is consumed.
func (v *Vector[T,U])Keys() iter.Iter[int] {
    return iter.Range[int](0,len(*v),1).SetupTeardown(
        func() error { v.RLock(); return nil },
        func() error { v.RUnlock(); return nil },
    )
}

// Returns true if the elements in v are all contained in other and the elements
// of other are all contained in v, regardless of position. Returns false 
// otherwise. This implementation of UnorderedEq is dependent on the time 
// complexity of the implementation of the ContainsPntr method on other. In 
// big-O it might look something like this, O(n*O(other.ContainsPntr))), where n 
// is the number of elements in v and O(other.ContainsPntr) represents the 
// time complexity of the containsPntr method on other with m values. Read locks 
// will be placed on both this vector and the other vector.
func (v *Vector[T,U])UnorderedEq(
    other containerTypes.ComparisonsOtherConstraint[T],
) bool {
    v.RLock()
    other.RLock()
    defer v.RUnlock()
    defer other.RUnlock()
    rv:=(len(*v)==other.Length())
    for i:=0; i<len(*v) && rv; i++ {
        rv=other.ContainsPntr(&(*v)[i])
    }
    return rv
}

// Returns true if all the key value pairs in v are all contained in other and 
// the key value pairs are all contained in v. Returns false otherwise. This 
// implementation of KeyedEq is dependent on the time complexity of the 
// implementation of the GetPntr method on other. In big-O it might look 
// something like this, O(n*O(other.GetPntr))), where n is the number of 
// elements in v and O(other.ContainsPntr) represents the time complexity of 
// the containsPntr method on other with m values. Read locks will be placed on 
// both this vector and the other vector.
func (v *Vector[T,U])KeyedEq(
    other containerTypes.KeyedComparisonsOtherConstraint[int,T],
) bool {
    v.RLock()
    other.RLock()
    defer v.RUnlock()
    defer other.RUnlock()
    w:=widgets.NewWidget[T,U]()
    rv:=(len(*v)==other.Length())
    for i:=0; i<len(*v) && rv; i++ {
        if otherV,err:=other.GetPntr(i); err!=nil {
            rv=false
        } else {
            rv=w.Eq(&(*v)[i],otherV)
        }
    }
    return rv
}

// Populates the vector with the intersection of values from the l and r 
// containers. This implementation of intersection is dependent on the time 
// complexity of the implementation of the ContainsPntr method on l and r. In 
// big-O it might look something like this, O(n*O(r.ContainsPntr)), where 
// O(r.ContainsPntr) represents the time complexity of the containsPntr 
// method on r. Read locks will be placed on l and r and a write lock will be 
// placed on the vector that is being populated.
//
// This vector will be cleared before storing the result. When clearing, the
// new resulting vector will be initialized with zero capacity and enough
// backing memory to store (l.Length()+r.Length())/2 elements before 
// reallocating. This means that there should be at most 1 reallocation beyond
// this initial allocation, and that additional allocation should only occur 
// when the length of the intersection is greater than the average length of the 
// l and r vectors. This logic is predicated on the fact that intersections will
// likely be much smaller than the original vectors.
func (v *Vector[T,U])Intersection(
    l containerTypes.ComparisonsOtherConstraint[T],
    r containerTypes.ComparisonsOtherConstraint[T],
) {
    r.RLock()
    l.RLock()
    v.Lock()
    defer r.RUnlock()
    defer l.RUnlock()
    defer v.Unlock()
    w:=widgets.NewWidget[T,U]()
    for i:=0; i<len(*v); i++ {
        w.Zero(&(*v)[i])
    }
    *v=make(Vector[T, U], 0, (l.Length()+r.Length())/2)
    l.ValPntrs().ForEach(func(index int, val *T) (iter.IteratorFeedback, error) {
        if r.ContainsPntr(val) {
            // Need to copy because val is comming from another container.
            *v=append(*v, *val) 
        }
        return iter.Continue,nil
    })
}

// Populates the vector with the union of values from the l and r containers. 
// The time complexity of this union operation will look like this in big-O: 
// O((n+m)*(n+m)), where n is the number of values in l and m is the number of
// values in n. Read locks will be placed on l and r and a write lock will be 
// placed on the vector that is being populated.
//
// This vector will be cleared before storing the result. When clearing, the
// new resulting vector will be initialized with zero capacity and enough
// backing memory to store the average of the maximum and minimum possible
// union sizes before reallocating. This means that there should be at most 1 
// reallocation beyond this initial allocation, and that additional allocation 
// should only occur when the length of the union is greater than the average 
// length of the minumum and maximum possible union sizes. This logic is 
// predicated on the fact that unions will likely be much smaller than the 
// original vectors.
func (v *Vector[T,U])Union(
    l containerTypes.ComparisonsOtherConstraint[T],
    r containerTypes.ComparisonsOtherConstraint[T],
) {
    r.RLock()
    l.RLock()
    v.Lock()
    defer r.RUnlock()
    defer l.RUnlock()
    defer v.Unlock()
    w:=widgets.NewWidget[T,U]()
    for i:=0; i<len(*v); i++ {
        w.Zero(&(*v)[i])
    }
    minLen:=max(l.Length(),r.Length())
    maxLen:=l.Length()+r.Length()
    *v=make(Vector[T, U], 0, (maxLen+minLen)/2)
    l.ValPntrs().ForEach(func(index int, val *T) (iter.IteratorFeedback, error) {
        if !v.ContainsPntr(val) {
            // Need to copy because val is comming from another container.
            *v=append(*v, *val) 
        }
        return iter.Continue,nil
    })
    r.ValPntrs().ForEach(func(index int, val *T) (iter.IteratorFeedback, error) {
        if !v.ContainsPntr(val) {
            // Need to copy because val is comming from another container.
            *v=append(*v, *val) 
        }
        return iter.Continue,nil
    })
}

// Populates the vector with the result of taking the difference of r from l.
// This implementation of difference is dependent on the time complexity of the 
// implementation of the ContainsPntr method on r. In big-O it might look 
// something like this, O(n*O(r.ContainsPntr)), where O(r.ContainsPntr) 
// represents the time complexity of the containsPntr method on r. Read locks 
// will be placed on l and r and a write lock will be placed on the vector that 
// is being populated.
//
// This vector will be cleared before storing the result. When clearing, the
// new resulting vector will be initialized with zero capacity and enough
// backing memory to store half the length of l. This means that there should be 
// at most 1 reallocation beyond this initial allocation, and that additional 
// allocation should only occur when the length of the difference is greater 
// than half the length of l. This logic is predicated on the fact that 
// differences will likely be much smaller than the original vector.
func (v *Vector[T,U])Difference(
    l containerTypes.ComparisonsOtherConstraint[T],
    r containerTypes.ComparisonsOtherConstraint[T],
) {
    r.RLock()
    l.RLock()
    v.Lock()
    defer r.RUnlock()
    defer l.RUnlock()
    defer v.Unlock()
    w:=widgets.NewWidget[T,U]()
    for i:=0; i<len(*v); i++ {
        w.Zero(&(*v)[i])
    }
    *v=make(Vector[T, U], 0, l.Length()/2)
    l.ValPntrs().ForEach(func(index int, val *T) (iter.IteratorFeedback, error) {
        if !r.ContainsPntr(val) {
            // Need to copy because val is comming from another container.
            *v=append(*v, *val) 
        }
        return iter.Continue,nil
    })
}

// Returns true if this vector is a superset to other. This implementation has
// a time complexity of O(n*m), where n is the number of values in this vector
// and m is the number of values in other. Read locks will be placed on both
// this vector and the other vector.
func (v *Vector[T,U])IsSuperset(
    other containerTypes.ComparisonsOtherConstraint[T],
) bool {
    v.RLock()
    other.RLock()
    defer v.RUnlock()
    defer other.RUnlock()
    rv:=(len(*v)>=other.Length())
    if !rv {
        return false
    }
    other.ValPntrs().ForEach(func(index int, val *T) (iter.IteratorFeedback, error) {
        if rv=v.ContainsPntr(val); !rv {
            return iter.Break,nil
        }
        return iter.Continue,nil
    })
    return rv
}

// Returns true if this vector is a subset to other. This implementation has a
// time complexity that is dependent on the ContainsPntr method of other. In 
// big-O terms it may look somwthing like this: O(n*O(other.ContainsPntr)), 
// where n is the number of elements in the current vector and 
// other.ContainsPntr represents the time complexity of the containsPntr method
// on other. Read locks will be placed on both this vector and the other vector.
func (v *Vector[T,U])IsSubset(
    other containerTypes.ComparisonsOtherConstraint[T],
) bool {
    v.RLock()
    other.RLock()
    defer v.RUnlock()
    defer other.RUnlock()
    rv:=(len(*v)<=other.Length())
    for i:=0; i<len(*v) && rv; i++ {
        rv=other.ContainsPntr(&(*v)[i])
    }
    return rv
}

// An equality function that implements the [algo.widget.WidgetInterface] 
// interface. Internally this is equivalent to [Vector.KeyedEq]. Returns true
// if l==r, false otherwise.
func (v *Vector[T, U])Eq(l *Vector[T,U], r *Vector[T,U]) bool {
    return l.KeyedEq(r)
}

// A function that implements the less than operation on vectors. The l and r
// vectors will be compared lexographically.
func (v *Vector[T, U])Lt(l *Vector[T,U], r *Vector[T,U]) bool {
    w:=widgets.NewWidget[T,U]()
    for i:=0; i<min(len(*l),len(*r)); i++ {
        if w.Lt(&(*l)[i],&(*r)[i]) {
            return true
        } else if w.Gt(&(*l)[i],&(*r)[i]) {
            return false
        }
    }
    if len(*l)>=len(*r) {
        return false
    }
    return true
}

// A function that returns a hash of a vector. To do this all of the individual
// hashes that are produced from the elements of the vector are combined in a
// way that maintains identity, making it so the hash will represent the same
// equality operation that [Vector.KeyedEq] and [Vector.Eq] provide.
func (c *Vector[T, U])Hash(other *Vector[T,U]) hash.Hash {
    other.RLock()
    defer other.RUnlock()
    var rv hash.Hash=0
    w:=widgets.NewWidget[T,U]()
    if len(*other)>0 {
        rv=w.Hash(&(*other)[0])    
        for i:=1; i<len(*other); i++ {
            rv=rv.Combine(uint64(w.Hash(&(*other)[i])))
        }
    }
    return rv
}

// An zero function that implements the [algo.widget.WidgetInterface] interface.
// Internally this is equivalent to [vector.Clear].
func (v *Vector[T, U])Zero(other *Vector[T,U]) {
    other.Clear()
}
