package containers

import (
	"sync"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/container/widgets"
	"github.com/barbell-math/util/customerr"
)

type (
    // A container that holds a fixed number of values in such a way that makes
    // stack and queue operations extremely efficient. Because the length of the
    // container is fixed (it will not dynamically expand to add more elements
    // as needed) the values in the underlying array will 'rotate' around the
    // array as operations are performed making it so no allocations are ever
    // performed beyond the initial creation of the underlying array.
    CircularBuffer[T any, U widgets.WidgetInterface[T]] struct {
        vals []T;
        numElems int;
        startEnd Pair[int,int];
    };
    
    // A synchronized version of CircularBuffer. All operations will be wrapped 
    // in the appropriate calls the embedded RWMutex. A pointer to a RWMutex is 
    // embedded rather than a value to avoid copying the lock value.
    SyncedCircularBuffer[T any, U widgets.WidgetInterface[T]] struct {
        *sync.RWMutex
        CircularBuffer[T, U]
    }
)

// Creates a new CircularBuffer initialized with size zero valued elements. Size 
// must be greater than 0, an error will be returned if it is not.
func NewCircularBuffer[T any, U widgets.WidgetInterface[T]](
    size int,
) (CircularBuffer[T,U],error) {
    if size<=0 {
        return CircularBuffer[T,U]{},customerr.Wrap(customerr.ValOutsideRange,
            "Size of buffer must be >0 | Have: %d",size,
        )
    }
    // B will always inc when adding values to the end
    // A will always dec when adding values to the beginning
    // B will always dec when removing values from the end
    // A will always inc when removing values from the beginning
    // These conventions are the only thing that determines the direction that
    // the buffer wraps around. You have been warned.
    return CircularBuffer[T,U]{
        vals: make([]T,size),
        startEnd: Pair[int, int]{A: 0, B: size-1},
    },nil;
}

// Creates a new synced CircularBuffer initialized with size zero valued 
// elements. Size must be greater than 0, an error will be returned if it is not.
// The underlying RWMutex value will be fully unlocked upon initialization.
func NewSyncedCircularBuffer[T any, U widgets.WidgetInterface[T]](
    size int,
) (SyncedCircularBuffer[T,U],error) {
    rv,err:=NewCircularBuffer[T,U](size)
    return SyncedCircularBuffer[T,U]{
        CircularBuffer: rv,
        RWMutex: &sync.RWMutex{},
    }, err
}

// A empty pass through function that performs no action. CircularBuffer will 
// call all the appropriate locking methods despite not being synced, just 
// nothing will happen. This is done so that SyncedCircularBuffer can simply 
// embed a CircularBuffer and override the appropriate locking methods to 
// implement the correct behavior without needing to make any additional changes 
// such as wrapping every single method from CircularBuffer.
func (c *CircularBuffer[T,U])Lock() { }

// A empty pass through function that performs no action. CircularBuffer will 
// call all the appropriate locking methods despite not being synced, just 
// nothing will happen. This is done so that SyncedCircularBuffer can simply 
// embed a CircularBuffer and override the appropriate locking methods to 
// implement the correct behavior without needing to make any additional changes 
// such as wrapping every single method from CircularBuffer.
func (c *CircularBuffer[T,U])Unlock() { }

// A empty pass through function that performs no action. CircularBuffer will 
// call all the appropriate locking methods despite not being synced, just 
// nothing will happen. This is done so that SyncedCircularBuffer can simply 
// embed a CircularBuffer and override the appropriate locking methods to 
// implement the correct behavior without needing to make any additional changes 
// such as wrapping every single method from CircularBuffer.
func (c *CircularBuffer[T,U])RLock() { }

// A empty pass through function that performs no action. CircularBuffer will 
// call all the appropriate locking methods despite not being synced, just 
// nothing will happen. This is done so that SyncedCircularBuffer can simply 
// embed a CircularBuffer and override the appropriate locking methods to 
// implement the correct behavior without needing to make any additional changes 
// such as wrapping every single method from CircularBuffer.
func (c *CircularBuffer[T,U])RUnlock() { }

// The SyncedCircularBuffer method to override the CircularBuffer pass through 
// function and actually apply the mutex operation.
func (c *SyncedCircularBuffer[T,U])Lock() { c.RWMutex.Lock() }

// The SyncedCircularBuffer method to override the CircularBuffer pass through 
// function and actually apply the mutex operation.
func (c *SyncedCircularBuffer[T,U])Unlock() { c.RWMutex.Unlock() }

// The SyncedCircularBuffer method to override the CircularBuffer pass through 
// function and actually apply the mutex operation.
func (c *SyncedCircularBuffer[T,U])RLock() { c.RWMutex.RLock() }

// The SyncedCircularBuffer method to override the CircularBuffer pass through 
// function and actually apply the mutex operation.
func (c *SyncedCircularBuffer[T,U])RUnlock() { c.RWMutex.RUnlock() }

// Returns true if the circular buffer has reached its capacity.
func (c *CircularBuffer[T,U])Full() bool {
    c.RLock()
    defer c.RUnlock()
    return c.numElems==len(c.vals)
}

// Returns the length of the circular buffer.
func (c *CircularBuffer[T,U])Length() int {
    c.RLock()
    defer c.RUnlock()
    return c.numElems;
}

// Returns the capacity of the circular buffer.
func (c *CircularBuffer[T,U])Capacity() int {
    c.RLock()
    defer c.RUnlock()
    return len(c.vals);
}

// Pushes an element to the front of the circular buffer. Equivalent to inserting 
// a single value at the front of the circular buffer. Returns an error if the
// circular buffer is full.
func (c *CircularBuffer[T,U])PushFront(v T) error {
    c.Lock()
    defer c.Unlock()
    if c.numElems<len(c.vals) {
        c.numElems++;
        c.startEnd.A=c.decIndex(c.startEnd.A,1);
        c.vals[c.startEnd.A]=v;
        return nil;
    }
    return c.getFullError()
}

// Pushes an element to the back of the circular buffer. Equivalent to appending 
// a single value to the circular buffer. Returns an error if the circular buffer 
// is full.
func (c *CircularBuffer[T,U])PushBack(v T) error {
    c.Lock()
    defer c.Unlock()
    if c.numElems<len(c.vals) {
        c.numElems++;
        c.startEnd.B=c.incIndex(c.startEnd.B,1);
        c.vals[c.startEnd.B]=v;
        return nil;
    }
    return c.getFullError()
}


// Pushes an element to the front of the circular buffer, poping an element from
// the back of the buffer if necessary to make room for the new element. If the 
// circular buffer is full then this is equavilent to poping and then pushing, 
// but more efficient.
func (c *CircularBuffer[T,U])ForcePushBack(v T) {
    c.Lock()
    defer c.Unlock()
    if c.numElems==len(c.vals) {
        c.startEnd.A=c.incIndex(c.startEnd.A,1);
        c.numElems--;
    }
    c.numElems++;
    c.startEnd.B=c.incIndex(c.startEnd.B,1);
    c.vals[c.startEnd.B]=v;
}

// Pushes an element to the back of the circular buffer, poping an element from
// the front of the buffer if necessary to make room for the new element. If the 
// circular buffer is full then this is equavilent to poping and then pushing, 
// but more efficient.
func (c *CircularBuffer[T,U])ForcePushFront(v T) {
    c.Lock()
    defer c.Unlock()
    if c.numElems==len(c.vals) {
        c.startEnd.B=c.decIndex(c.startEnd.B,1);
        c.numElems--;
    }
    c.numElems++;
    c.startEnd.A=c.decIndex(c.startEnd.A,1);
    c.vals[c.startEnd.A]=v;
}

// Returns the value at index 0 if one is present. If the circular buffer has no 
// elements then an error is returned.
func (c *CircularBuffer[T,U])PeekFront() (T,error) {
    v,err:=c.PeekPntrFront();
    if v!=nil {
        return *v,err;
    }
    var tmp T;
    return tmp,err;
}

// Returns a pointer to the value at index 0 if one is present. If the circular 
// buffer has no elements then an error is returned.
func (c *CircularBuffer[T,U])PeekPntrFront() (*T,error) {
    c.RLock()
    defer c.RUnlock()
    if c.numElems>0 {
        return &c.vals[c.startEnd.A],nil
    }
    return nil,getIndexOutOfBoundsError(0,0,c.numElems)
}

// Returns the value at index len(circular buffer)-1 if one is present. If the 
// circular buffer has no elements then an error is returned.
func (c *CircularBuffer[T, U])PeekBack() (T,error) {
    v,err:=c.PeekPntrBack();
    if v!=nil {
        return *v,err;
    }
    var tmp T;
    return tmp,err;
}

// Returns a pointer to the value at index len(circular buffer)-1 if one is 
// present. If the circular buffer has no elements then an error is returned.
func (c *CircularBuffer[T, U])PeekPntrBack() (*T,error) {
    c.RLock()
    defer c.RUnlock()
    if c.numElems>0 {
        return &c.vals[c.startEnd.B],nil
    }
    return nil,getIndexOutOfBoundsError(0,0,c.numElems)
}

// Gets the value at the specified index. Returns an error if the index is 
// >= the length of the circular buffer.
func (c *CircularBuffer[T,U])Get(idx int) (T,error){
    v,err:=c.GetPntr(idx);
    if v!=nil {
        return *v,err;
    }
    var tmp T;
    return tmp,err;
}

// Gets a pointer to the value at the specified index. Returns an error if the 
// index is >= the length of the circular buffer.
func (c *CircularBuffer[T,U])GetPntr(idx int) (*T,error) {
    c.RLock()
    defer c.RUnlock()
    if idx>=0 && idx<c.numElems && c.numElems>0 {
        properIndex:=c.getProperIndex(idx)
        return &c.vals[properIndex],nil;
    }
    return nil,getIndexOutOfBoundsError(idx,0,c.numElems)
}

// Contains will return true if the supplied value is in the vector, false
// otherwise. All equality comparisons are performed by the generic U widget
// type that the vector was initialized with.
func (c *CircularBuffer[T, U])Contains(val T) bool {
    c.RLock()
    defer c.RUnlock()
    found:=false
    w:=widgets.NewWidget[T,U]()
    for i:=0; i<c.numElems && !found; i++ {
        properIndex:=c.getProperIndex(i)
        found=w.Eq(&val,&c.vals[properIndex])
    }
    return found
}

// TODO - implement
func (c *CircularBuffer[T, U])KeyOf(val T) (int,bool) {
    return -1,false
}

// Emplaces (sets) the value at the specified index. Returns an error if the 
// index is >= the length of the circular buffer.
func (c *CircularBuffer[T,U])Emplace(idx int, v T) error {
    c.Lock()
    defer c.Unlock()
    if idx>=0 && idx<c.numElems && c.numElems>0 {
        properIndex:=c.getProperIndex(idx)
        c.vals[properIndex]=v
        return nil
    }
    return getIndexOutOfBoundsError(idx,0,c.numElems)
}

// Pushes (inserts) the supplied values at the given index. Returns an error if 
// the index is >= the length of the circular buffer.
func (c *CircularBuffer[T,U])Push(idx int, v ...T) error {
    c.Lock()
    defer c.Unlock()
    if idx<0 || idx>c.numElems {
        return getIndexOutOfBoundsError(idx,0,c.numElems)
    } else if c.numElems==len(c.vals) {
        return c.getFullError()
    }
    maxVals:=len(v)
    if c.numElems+maxVals>len(c.vals) {
        maxVals=len(c.vals)-c.numElems
    }
    if c.distanceFromBack(idx)>c.distanceFromFront(idx) {
        c.insertMoveFront(v,idx,maxVals)
    } else {
        c.insertMoveBack(v,idx,maxVals)
    }
    if maxVals<len(v) {
        return c.getFullError()
    }
    return nil
}

func (c *CircularBuffer[T,U])insertMoveFront(v []T, idx int, maxVals int) {
    c.numElems+=maxVals;
    c.startEnd.A=c.decIndex(c.startEnd.A,maxVals);
    for j,i:=0,maxVals-1; i<idx+maxVals; i++ {
        c.vals[c.getProperIndex(j)]=c.vals[c.getProperIndex(i)]
        j++
    }
    for i:=idx; i<idx+maxVals; i++ {
        c.vals[c.getProperIndex(i)]=v[i-idx]
    }
}

func (c *CircularBuffer[T,U])insertMoveBack(v []T, idx int, maxVals int) {
    c.numElems+=maxVals
    c.startEnd.B=c.incIndex(c.startEnd.B,maxVals)
    for j,i:=c.numElems-1,c.numElems-maxVals-1; i>=idx; i-- {
        c.vals[c.getProperIndex(j)]=c.vals[c.getProperIndex(i)]
        j--
    }
    for i:=idx; i<idx+maxVals; i++ {
        c.vals[c.getProperIndex(i)]=v[i-idx]
    }
}

// Appends the supplied values to the circular buffer. A [containerTypes.Full] error will be
// returned if the circular buffer reaches it's capacity.
func (c *CircularBuffer[T,U])Append(v ...T) error {
    c.Lock()
    defer c.Unlock()
    for i:=0; i<len(v); i++ {
        if c.numElems<len(c.vals) {
            c.numElems++;
            c.startEnd.B=c.incIndex(c.startEnd.B,1);
            c.vals[c.startEnd.B]=v[i];
        } else {
            return c.getFullError()
        }
    }
    return nil
}

// Returns and removes the element at the front of the circular buffer. Returns 
// an error if the circular buffer has no elements.
func (c *CircularBuffer[T,U])PopFront() (T,error) {
    c.Lock()
    defer c.Unlock()
    if c.numElems>0 {
        rv:=c.vals[c.startEnd.A];
        c.startEnd.A=c.incIndex(c.startEnd.A,1);
        c.numElems--;
        return rv,nil;
    }
    var tmp T;
    return tmp,containerTypes.Empty
}

// Returns and removes the element at the back of the circular buffer. Returns 
// an error if the circular buffer has no elements.
func (c *CircularBuffer[T, U])PopBack() (T,error) {
    c.Lock()
    defer c.Unlock()
    if c.numElems>0 {
        rv:=c.vals[c.startEnd.B];
        c.startEnd.B=c.decIndex(c.startEnd.B,1);
        c.numElems--;
        return rv,nil;
    }
    var tmp T;
    return tmp,containerTypes.Empty
}

// Deletes the value at the specified index. If the index is >= the length of the
// circular buffer then no action is taken and no error is returned. The function 
// will never return an error.
func (c *CircularBuffer[T,U])Delete(idx int) error {
    c.Lock()
    defer c.Unlock()
    if idx<0 || idx>=c.numElems {
        return getIndexOutOfBoundsError(idx,0,c.numElems)
    }
    if c.numElems==1 && idx==0 {
        c.numElems--;
        c.startEnd.B=c.decIndex(c.startEnd.B,1);
    } else if c.distanceFromBack(idx)>c.distanceFromFront(idx) {
        c.deleteMoveFront(idx)
    } else {
        c.deleteMoveBack(idx)
    }
    return nil
}

// TODO - implement
func (c *CircularBuffer[T, U])Pop(val T, num int) int {
    return -1
}

func (c *CircularBuffer[T,U])deleteMoveFront(idx int) {
    for i:=idx-1; i>=0; i-- {
        c.vals[c.getProperIndex(i+1)]=c.vals[c.getProperIndex(i)]
    }
    c.numElems--;
    c.startEnd.A=c.incIndex(c.startEnd.A,1);
}

func (c *CircularBuffer[T,U])deleteMoveBack(idx int) {
    for i:=idx; i<c.numElems; i++ {
        c.vals[c.getProperIndex(i)]=c.vals[c.getProperIndex(i+1)]
    }
    c.numElems--;
    c.startEnd.B=c.decIndex(c.startEnd.B,1);
}

// Clears all values from the circular buffer. Equivalent to making a new 
// circular buffer and setting it equal to the current one.
func (c *CircularBuffer[T,U])Clear() {
    c.Lock()
    defer c.Unlock()
    c.vals=make([]T,len(c.vals))
    c.numElems=0
    c.startEnd=Pair[int, int]{A: 0, B: len(c.vals)-1}
}

// Returns an iterator that iterates over the values in the circular buffer. The 
// circular buffer will have a read lock the entire time the iteration is being 
// performed. The lock will not be applied until the iterator is consumed.
func (c *CircularBuffer[T,U])Elems() iter.Iter[T] {
    return iter.SequentialElems[T](
        c.numElems,
        c.Get,
    ).SetupTeardown(
        func() error { c.RLock(); return nil },
        func() error { c.RUnlock(); return nil },
    )
}

// Returns an iterator that iterates over the pointers to ithe values in the 
// circular buffer. The circular buffer will have a read lock the entire time 
// the iteration is being performed. The lock will not be applied until the 
// iterator is consumed.
func (c *CircularBuffer[T,U])PntrElems() iter.Iter[*T] {
    return iter.SequentialElems[*T](
        c.numElems,
        c.GetPntr,
    ).SetupTeardown(
        func() error { c.RLock(); return nil },
        func() error { c.RUnlock(); return nil },
    )
}

// // Returns true if the circular buffers are equal. The supplied comparison 
// // function will be used when comparing values in the circular buffer.
// func (c *CircularBuffer[T,U])Eq(
//     other *CircularBuffer[T,U], 
//     comp func(l *T, r *T) bool,
// ) bool {
//     c.RLock()
//     defer c.RUnlock()
//     rv:=(c.Length()==other.Length())
//     for i:=0; i<c.Length() && rv; i++ {
//         l,_:=c.GetPntr(i)   // Ignoring index error because of loop bounds
//         r,_:=other.GetPntr(i)   // Ignoring index error because of loop bounds
//         rv=(rv && comp(l,r))
//     }
//     return rv
// }
// 
// // Returns true if the circular buffers are not equal. The supplied comparison 
// // function will be used when comparing values in the circular buffer.
// func (c *CircularBuffer[T,U])Neq(
//     other *CircularBuffer[T,U], 
//     comp func(l *T, r *T) bool,
// ) bool {
//     return !c.Eq(other,comp)
// }

// This function only works for an index that is <2n when n is the capacity of 
// the underlying array
func (c *CircularBuffer[T,U])incIndex(idx int, amnt int) int {
    rv:=idx+amnt
    if rv>=len(c.vals) {
        rv-=len(c.vals)
    }
    return rv
}

// This function only works for an index that is <2n when n is the capacity of 
// the underlying array
func (c *CircularBuffer[T,U])decIndex(idx int, amnt int) int {
    rv:=idx-amnt
    if rv<0 {
        rv+=len(c.vals)
    }
    return rv
}

func (c *CircularBuffer[T,U])getProperIndex(idx int) int {
    properIndex:=idx+c.startEnd.A;
    if properIndex>=len(c.vals) {
        properIndex-=len(c.vals);
    }
    return properIndex
}

func (c *CircularBuffer[T,U])distanceFromFront(idx int) int {
    return idx
}
func (c *CircularBuffer[T,U])distanceFromBack(idx int) int{
    return c.numElems-idx
}

func (c *CircularBuffer[T,U])getFullError() error {
    return customerr.Wrap(containerTypes.Full,"Circular buffer size: %d",len(c.vals))
}
