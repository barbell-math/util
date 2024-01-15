package dataStruct;

import (
    "fmt"
    "sync"
    "github.com/barbell-math/util/algo/iter"
    customerr "github.com/barbell-math/util/err"
)

type (
    CircularBuffer[T any] struct {
        vals []T;
        numElems int;
        startEnd Pair[int,int];
    };
    
    SyncedCircularBuffer[T any] struct {
        *sync.RWMutex
        CircularBuffer[T]
    }
)

func NewCircularBuffer[T any](size int) (CircularBuffer[T],error) {
    if size<=0 {
        return CircularBuffer[T]{},customerr.ValOutsideRange(
            fmt.Sprintf("Size of queue must be >0 | Have: %d",size),
        );
    }
    // B will always inc when adding values to the end
    // A will always dec when adding values to the beginning
    // B will always dec when removing values from the end
    // A will always inc when removing values from the beginning
    // These conventions are the only thing that determines the direction that
    // the buffer wraps around. You have been warned.
    return CircularBuffer[T]{
        vals: make([]T,size),
        startEnd: Pair[int, int]{A: 0, B: size-1},
    },nil;
}

func NewSyncedCircularBuffer[T any](size int) (SyncedCircularBuffer[T],error) {
    rv,err:=NewCircularBuffer[T](size)
    return SyncedCircularBuffer[T]{
        CircularBuffer: rv,
        RWMutex: &sync.RWMutex{},
    }, err
}

func (c *CircularBuffer[T])Lock() { }
func (c *CircularBuffer[T])Unlock() { }
func (c *CircularBuffer[T])RLock() { }
func (c *CircularBuffer[T])RUnlock() { }

func (c *SyncedCircularBuffer[T])Lock() { c.RWMutex.Lock() }
func (c *SyncedCircularBuffer[T])Unlock() { c.RWMutex.Unlock() }
func (c *SyncedCircularBuffer[T])RLock() { c.RWMutex.RLock() }
func (c *SyncedCircularBuffer[T])RUnlock() { c.RWMutex.RUnlock() }

func (c *CircularBuffer[T])Full() bool {
    c.RLock()
    defer c.RUnlock()
    return c.numElems==len(c.vals)
}

func (c *CircularBuffer[T])Length() int {
    c.RLock()
    defer c.RUnlock()
    return c.numElems;
}

func (c *CircularBuffer[T])Capacity() int {
    c.RLock()
    defer c.RUnlock()
    return len(c.vals);
}

func (c *CircularBuffer[T])PushFront(v T) error {
    c.Lock()
    defer c.Unlock()
    if c.numElems<len(c.vals) {
        c.numElems++;
        c.startEnd.A=c.decIndex(c.startEnd.A,1);
        c.vals[c.startEnd.A]=v;
        return nil;
    }
    return c.getQueueFullError()
}

func (c *CircularBuffer[T])PushBack(v T) error {
    c.Lock()
    defer c.Unlock()
    if c.numElems<len(c.vals) {
        c.numElems++;
        c.startEnd.B=c.incIndex(c.startEnd.B,1);
        c.vals[c.startEnd.B]=v;
        return nil;
    }
    return c.getQueueFullError()
}

func (c *CircularBuffer[T])ForcePushBack(v T) {
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

func (c *CircularBuffer[T])ForcePushFront(v T) {
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

func (c *CircularBuffer[T])PeekFront() (T,error) {
    v,err:=c.PeekPntrFront();
    if v!=nil {
        return *v,err;
    }
    var tmp T;
    return tmp,err;
}

func (c *CircularBuffer[T])PeekPntrFront() (*T,error) {
    c.RLock()
    defer c.RUnlock()
    if c.numElems>0 {
        return &c.vals[c.startEnd.A],nil
    }
    return nil,getIndexOutOfBoundsError(0,c.numElems)
}

func (c *CircularBuffer[T])Get(idx int) (T,error){
    v,err:=c.GetPntr(idx);
    if v!=nil {
        return *v,err;
    }
    var tmp T;
    return tmp,err;
}

func (c *CircularBuffer[T])GetPntr(idx int) (*T,error) {
    c.RLock()
    defer c.RUnlock()
    if idx>=0 && idx<c.numElems && c.numElems>0 {
        properIndex:=c.getProperIndex(idx)
        return &c.vals[properIndex],nil;
    }
    return nil,getIndexOutOfBoundsError(idx,c.numElems)
}

func (c *CircularBuffer[T])Set(v T, idx int) error {
    c.Lock()
    defer c.Unlock()
    if idx>=0 && idx<c.numElems && c.numElems>0 {
        properIndex:=c.getProperIndex(idx)
        c.vals[properIndex]=v
        return nil
    }
    return getIndexOutOfBoundsError(idx,c.numElems)
}

func (c *CircularBuffer[T])Insert(idx int, v ...T) error {
    c.Lock()
    defer c.Unlock()
    if idx<0 || idx>c.numElems {
        return getIndexOutOfBoundsError(idx,c.numElems)
    } else if c.numElems==len(c.vals) {
        return c.getQueueFullError()
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
        return c.getQueueFullError()
    }
    return nil
}

func (c *CircularBuffer[T])insertMoveFront(v []T, idx int, maxVals int) {
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

func (c *CircularBuffer[T])insertMoveBack(v []T, idx int, maxVals int) {
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

func (c *CircularBuffer[T])Append(v ...T) error {
    c.Lock()
    defer c.Unlock()
    for i:=0; i<len(v); i++ {
        if c.numElems<len(c.vals) {
            c.numElems++;
            c.startEnd.B=c.incIndex(c.startEnd.B,1);
            c.vals[c.startEnd.B]=v[i];
        } else {
            return c.getQueueFullError()
        }
    }
    return nil
}

func (c *CircularBuffer[T])PopFront() (T,error) {
    c.Lock()
    defer c.Unlock()
    if c.numElems>0 {
        rv:=c.vals[c.startEnd.A];
        c.startEnd.A=c.incIndex(c.startEnd.A,1);
        c.numElems--;
        return rv,nil;
    }
    var tmp T;
    return tmp,Empty("Nothing to pop!");
}

func (c *CircularBuffer[T])Delete(idx int) error {
    c.Lock()
    defer c.Unlock()
    if idx<0 || idx>=c.numElems {
        return getIndexOutOfBoundsError(idx,c.numElems)
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

func (c *CircularBuffer[T])deleteMoveFront(idx int) {
    for i:=idx-1; i>=0; i-- {
        c.vals[c.getProperIndex(i+1)]=c.vals[c.getProperIndex(i)]
    }
    c.numElems--;
    c.startEnd.A=c.incIndex(c.startEnd.A,1);
}

func (c *CircularBuffer[T])deleteMoveBack(idx int) {
    for i:=idx; i<c.numElems; i++ {
        c.vals[c.getProperIndex(i)]=c.vals[c.getProperIndex(i+1)]
    }
    c.numElems--;
    c.startEnd.B=c.decIndex(c.startEnd.B,1);
}

func (c *CircularBuffer[T])Clear() {
    c.Lock()
    defer c.Unlock()
    c.vals=make([]T,len(c.vals))
    c.numElems=0
    c.startEnd=Pair[int, int]{A: 0, B: len(c.vals)-1}
}

func (c *CircularBuffer[T])Elems() iter.Iter[T] {
    return iter.SequentialElems[T](
        c.numElems,
        c.Get,
    ).SetupTeardown(
        func() error { c.RLock(); return nil },
        func() error { c.RUnlock(); return nil },
    )
}

func (c *CircularBuffer[T])PntrElems() iter.Iter[*T] {
    return iter.SequentialElems[*T](
        c.numElems,
        c.GetPntr,
    ).SetupTeardown(
        func() error { c.RLock(); return nil },
        func() error { c.RUnlock(); return nil },
    )
}

// This function only works for an index that is <2n when n is the capacity of 
// the underlying array
func (c *CircularBuffer[T])incIndex(idx int, amnt int) int {
    rv:=idx+amnt
    if rv>=len(c.vals) {
        rv-=len(c.vals)
    }
    return rv
}

// This function only works for an index that is <2n when n is the capacity of 
// the underlying array
func (c *CircularBuffer[T])decIndex(idx int, amnt int) int {
    rv:=idx-amnt
    if rv<0 {
        rv+=len(c.vals)
    }
    return rv
}

func (c *CircularBuffer[T])getProperIndex(idx int) int {
    properIndex:=idx+c.startEnd.A;
    if properIndex>=len(c.vals) {
        properIndex-=len(c.vals);
    }
    return properIndex
}

func (c *CircularBuffer[T])distanceFromFront(idx int) int {
    return idx
}
func (c *CircularBuffer[T])distanceFromBack(idx int) int{
    return c.numElems-idx
}

func (c *CircularBuffer[T])getQueueFullError() error {
    return QueueFull(fmt.Sprintf("Queue size: %d",len(c.vals)));
}
