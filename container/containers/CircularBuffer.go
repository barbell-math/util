package containers

import (
	"sync"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/algo/widgets"
	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/container/containerTypes"
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
        startEnd basic.Pair[int,int];
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
        return CircularBuffer[T,U]{},customerr.Wrap(
            customerr.ValOutsideRange,
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
        startEnd: basic.Pair[int, int]{A: 0, B: size-1},
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

// Converts the supplied map to a syncronized map. Beware: The original 
// non-synced circular buffer will remain useable.
func (c *CircularBuffer[T, U])ToSynced() SyncedCircularBuffer[T,U] {
    return SyncedCircularBuffer[T, U]{
        RWMutex: &sync.RWMutex{},
        CircularBuffer: *c,
    }
}

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (c *CircularBuffer[T, U])Lock() { }

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (c *CircularBuffer[T, U])Unlock() { }

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (c *CircularBuffer[T, U])RLock() { }

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (c *CircularBuffer[T, U])RUnlock() { }

// The SyncedCircularBuffer method to override the HashMap pass through function 
// and actually apply the mutex operation.
func (c *SyncedCircularBuffer[T,U])Lock() { c.RWMutex.Lock() }

// The SyncedCircularBuffer method to override the HashMap pass through function 
// and actually apply the mutex operation.
func (c *SyncedCircularBuffer[T,U])Unlock() { c.RWMutex.Unlock() }

// The SyncedCircularBuffer method to override the HashMap pass through function 
// and actually apply the mutex operation.
func (c *SyncedCircularBuffer[T, U])RLock() { c.RWMutex.RLock() }

// The SyncedCircularBuffer method to override the HashMap pass through function 
// and actually apply the mutex operation.
func (c *SyncedCircularBuffer[T, U])RUnlock() { c.RWMutex.RUnlock() }

// Returns true, a circular biffer is addressable.
func (c *CircularBuffer[T, U])IsAddressable() bool { return true }

// Returns false, a circular buffer is not synced.
func (c *CircularBuffer[T, U])IsSynced() bool { return false }

// Returns true, a synced circular buffer is synced.
func (c *SyncedCircularBuffer[T,U])IsSynced() bool { return true }

// Description: Returns true if the circular buffer has reached its capacity,
// false otherwise.
//
// Time Complexity: O(1)
func (c *CircularBuffer[T,U])Full() bool {
    return c.numElems==len(c.vals)
}
// Description: Places a read lock on the underlying circular buffer and then 
// calls the underlying circular buffers [CircularBuffer.Length] method.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (c *SyncedCircularBuffer[T, U])Full() bool {
    c.RLock()
    defer c.RUnlock()
    return c.CircularBuffer.Full()
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

// Description: Gets the value at the specified index. Returns an error if the 
// index is >= the length of the circular buffer.
//
// Time Complexity: O(1)
func (c *CircularBuffer[T,U])Get(idx int) (T,error){
    if idx>=0 && idx<c.numElems && c.numElems>0 {
        properIndex:=c.getProperIndex(idx)
        return c.vals[properIndex],nil;
    }
    var tmp T
    return tmp,getIndexOutOfBoundsError(idx,0,c.numElems)
}
// Description: Places a read lock on the underlying circular buffer and then 
// gets the value at the specified index. Exhibits the same behavior as the 
// [CircularBuffer.Get] method. The underlying [CircularBuffer.Get] method is 
// not called to avoid copying the return value twice, which could be 
// inefficient with a large value for the T generic.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (c *SyncedCircularBuffer[T, U])Get(idx int) (T,error) {
    c.RLock()
    defer c.RUnlock()
    if idx>=0 && idx<c.numElems && c.numElems>0 {
        properIndex:=c.getProperIndex(idx)
        return c.vals[properIndex],nil;
    }
    var tmp T
    return tmp,getIndexOutOfBoundsError(idx,0,c.numElems)
}

// Description: Gets a pointer to the value at the specified index. Returns an 
// error if the index is >= the length of the circular buffer.
//
// Time Complexity: O(1)
func (c *CircularBuffer[T,U])GetPntr(idx int) (*T,error) {
    if idx>=0 && idx<c.numElems && c.numElems>0 {
        properIndex:=c.getProperIndex(idx)
        return &c.vals[properIndex],nil;
    }
    return nil,getIndexOutOfBoundsError(idx,0,c.numElems)
}
// Description: Places a read lock on the underlying circular buffer and then 
// calls the underlying circular buffer [CircularBuffer.GetPntr] method.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (c *SyncedCircularBuffer[T, U])GetPntr(idx int) (*T,error) {
    c.RLock()
    defer c.RUnlock()
    return c.CircularBuffer.GetPntr(idx)
}

// Description: Contains will return true if the supplied value is in the 
// circular buffer, false otherwise. All equality comparisons are performed by 
// the generic U widget type that the circular buffer was initialized with. 
//
// Time Complexity: O(n) (linear search)
func (c *CircularBuffer[T, U])Contains(val T) bool {
    return c.ContainsPntr(&val)
}
// Description: Places a read lock on the underlying circular buffer and then 
// calls the underlying circular buffers [CircularBuffer.ContainsPntr] method.
//
// Lock Type: Read
//
// Time Complexity: O(n) (linear search)
func (c *SyncedCircularBuffer[T, U])Contains(val T) bool {
    c.RLock()
    defer c.RUnlock()
    return c.CircularBuffer.ContainsPntr(&val)
}

// Description: ContainsPntr will return true if the supplied value is in the 
// circular buffer, false otherwise. All equality comparisons are performed by 
// the generic U widget type that the circular buffer was initialized with.
//
// Time Complexity: O(n) (linear search)
func (c *CircularBuffer[T, U])ContainsPntr(val *T) bool {
    found:=false
    w:=widgets.NewWidget[T,U]()
    for i:=0; i<c.numElems && !found; i++ {
        properIndex:=c.getProperIndex(i)
        found=w.Eq(val,&c.vals[properIndex])
    }
    return found
}
// Description: Places a read lock on the underlying circular buffer and then 
// calls the underlying circular buffer [CircularBuffer.ContainsPntr] method.
//
// Lock Type: Read
//
// Time Complexity: O(n) (linear search)
func (c *SyncedCircularBuffer[T, U])ContainsPntr(val *T) bool {
    c.RLock()
    defer c.RUnlock()
    return c.CircularBuffer.ContainsPntr(val)
}

// Description: KeyOf will return the index of the first occurrence of the 
// supplied value in the circular buffer. If the value is not found then the 
// returned index will be -1 and the boolean flag will be set to false. If the 
// value is found then the boolean flag will be set to true. All equality 
// comparisons are performed by the generic U widget type that the circular buffer was 
// initialized with.
//
// Time Complexity: O(n) (linear search)
func (c *CircularBuffer[T, U])KeyOf(val T) (int,bool) {
    return c.keyOfImpl(&val)
}
// Description: Places a read lock on the underlying circular buffer and then 
// calls the underlying circular buffers [CircularBuffer.KeyOf] implemenation 
// method. The [CircularBuffer.KeyOf] method is not called directly to avoid 
// copying the val variable twice, which could be expensive with a large type 
// for the T generic.
//
// Lock Type: Read
//
// Time Complexity: O(n) (linear search)
func (c *SyncedCircularBuffer[T, U])KeyOf(val T) (int,bool) {
    c.RLock()
    defer c.RUnlock()
    return c.CircularBuffer.keyOfImpl(&val)
}

func (c *CircularBuffer[T,U])keyOfImpl(val *T) (int,bool) {
    w:=widgets.NewWidget[T,U]()
    for i:=0; i<c.numElems; i++ {
        properIndex:=c.getProperIndex(i)
        if w.Eq(val,&c.vals[properIndex]) {
            return i,true
        }
    }
    return -1,false
}

// Description: Sets the values at the specified indexes. Returns an error if 
// the index is >= the length of the circular buffer. Stops setting values as 
// soon as an error is encountered.
//
// Time Complexity: O(m), where m=len(vals)
func (c *CircularBuffer[T,U])Set(vals ...basic.Pair[int,T]) error {
    return c.setImpl(vals)
}
// Description: Places a write lock on the underlying circular buffer and then 
// calls the underlying circular buffers [CircularBuffer.Set] implementaiton 
// method. The [CircularBuffer.Set] method is not called directly to avoid 
// copying the vals varargs twice, which could be expensive with a large type 
// for the T generic or a large number of values.
//
// Lock Type: Write
//
// Time Complexity: O(m), where m=len(vals)
func (c *SyncedCircularBuffer[T, U])Set(vals ...basic.Pair[int,T]) error {
    c.Lock()
    defer c.Unlock()
    return c.CircularBuffer.setImpl(vals)
}

func (c *CircularBuffer[T, U])setImpl(vals []basic.Pair[int,T]) error {
    for _,iterV:=range(vals) {
        if iterV.A>=0 && iterV.A<c.numElems && len(c.vals)>0 {
            c.vals[c.getProperIndex(iterV.A)]=iterV.B
        } else {
            return getIndexOutOfBoundsError(iterV.A,0,c.numElems)
        }
    }
    return nil
}

// Description: Sets the supplied values sequentially starting at the supplied 
// index and continuing sequentailly after that. Returns and error if any index 
// that is attempted to be set is >= the length of the circular buffer. If an 
// error occurs, all values will be set up until the value that caused the error.
//
// Time Complexity: O(m), where m=len(vals)
func (c *CircularBuffer[T, U])SetSequential(idx int, vals ...T) error {
    return c.setSequentialImpl(idx,vals)
}
// Description: Places a write lock on the underlying circular buffer and then 
// calls the underlying circular buffers [CircularBuffer.SetSequential] 
// implementation method. The [circular buffer.SetSequential] method is not 
// called directly to avoid copying the vals varargs twice, which could be 
// expensive with a large type for the T generic or a large number of values.
//
// Lock Type: Write
//
// Time Complexity: O(m), where m=len(vals)
func (c *SyncedCircularBuffer[T, U])SetSequential(idx int, vals ...T) error {
    c.Lock()
    defer c.Unlock()
    return c.CircularBuffer.setSequentialImpl(idx,vals)
}

func (c *CircularBuffer[T, U])setSequentialImpl(idx int, vals []T) error {
    if idx>=len(c.vals) {
        return getIndexOutOfBoundsError(idx,0,len(c.vals))
    }
    numCopyableVals:=min(len(c.vals)-idx,len(vals))
    copy((c.vals)[idx:idx+numCopyableVals],vals[0:numCopyableVals])
    if idx+len(vals)>len(c.vals) {
        return getIndexOutOfBoundsError(len(c.vals),0,len(c.vals)) 
    }
    return nil
}

// Description: Append the supplied values to the circular buffer. This function
// will never return an error.
//
// Time Complexity: Best case O(m), where m=len(vals).
func (c *CircularBuffer[T,U])Append(vals ...T) error {
    return c.appendImpl(vals)
}
// Description: Places a write lock on the underlying circular buffer and then 
// calls the underlying circular buffers [CircularBuffer.Append] implementation 
// method. The [CircularBuffer.Append] method is not called directly to avoid 
// copying the vals varargs twice, which could be expensive with a large type 
// for the T generic or a large number of values.
//
// Lock Type: Write
//
// Time Complexity: Best case O(m), where m=len(vals).
func (c *SyncedCircularBuffer[T, U])Append(vals ...T) error {
    c.Lock()
    defer c.Unlock()
    return c.CircularBuffer.appendImpl(vals)
}

func (c *CircularBuffer[T,U])appendImpl(vals []T) error {
    for i:=0; i<len(vals); i++ {
        if c.numElems<len(c.vals) {
            c.numElems++;
            c.startEnd.B=c.incIndex(c.startEnd.B,1);
            c.vals[c.startEnd.B]=vals[i];
        } else {
            return c.getFullError()
        }
    }
    return nil
}


// Pushes (inserts) the supplied values at the given index. Returns an error if 
// the index is >= the length of the circular buffer.
func (c *CircularBuffer[T,U])Insert(vals ...basic.Pair[int,T]) error {
    // c.Lock()
    // defer c.Unlock()
    // if idx<0 || idx>c.numElems {
    //     return getIndexOutOfBoundsError(idx,0,c.numElems)
    // } else if c.numElems==len(c.vals) {
    //     return c.getFullError()
    // }
    // maxVals:=len(v)
    // if c.numElems+maxVals>len(c.vals) {
    //     maxVals=len(c.vals)-c.numElems
    // }
    // if c.distanceFromBack(idx)>c.distanceFromFront(idx) {
    //     c.insertMoveFront(v,idx,maxVals)
    // } else {
    //     c.insertMoveBack(v,idx,maxVals)
    // }
    // if maxVals<len(v) {
    //     return c.getFullError()
    // }
    return nil
}

func (c *CircularBuffer[T, U])InsertSequential(idx int, vals ...int) error {
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

// TODO - impl and test
func (c *CircularBuffer[T, U])DeleteSequential(idx int, num int) error {
    return nil
}

// TODO - implement
func (c *CircularBuffer[T, U])Pop(val T) int {
    return -1
}

func (c *CircularBuffer[T, U])PopSequential(val T, num int) int {
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
    c.startEnd=basic.Pair[int, int]{A: 0, B: len(c.vals)-1}
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

func (c *CircularBuffer[T, U])Keys() iter.Iter[int] {
    return iter.NoElem[int]()
}

func (c *CircularBuffer[T, U])Vals() iter.Iter[T] {
    return iter.NoElem[T]()
}

func (c *CircularBuffer[T, U])ValPntrs() iter.Iter[*T] {
    return iter.NoElem[*T]()
}

// TODO - impl and test
func (c *CircularBuffer[T,U])KeyedEq(other containerTypes.KeyedComparisonsOtherConstraint[int,T]) bool { return false }
func (c *CircularBuffer[T,U])UnorderedEq(other containerTypes.ComparisonsOtherConstraint[T]) bool { return false }
func (c *CircularBuffer[T,U])Intersection(l containerTypes.ComparisonsOtherConstraint[T], r containerTypes.ComparisonsOtherConstraint[T]) {}
func (c *CircularBuffer[T,U])Union(l containerTypes.ComparisonsOtherConstraint[T], r containerTypes.ComparisonsOtherConstraint[T]) {}
func (c *CircularBuffer[T,U])Difference(l containerTypes.ComparisonsOtherConstraint[T], r containerTypes.ComparisonsOtherConstraint[T]) {}
func (c *CircularBuffer[T,U])IsSuperset(other containerTypes.ComparisonsOtherConstraint[T]) bool { return false }
func (c *CircularBuffer[T,U])IsSubset(other containerTypes.ComparisonsOtherConstraint[T]) bool { return false }
