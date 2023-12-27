package dataStruct;

import (
    "fmt"
    "sync"
    "github.com/barbell-math/util/algo/iter"
    customerr "github.com/barbell-math/util/err"
)

type CircularBuffer[T any] struct {
    m *sync.RWMutex
    vals []T;
    numElems int;
    startEnd Pair[int,int];
};

func NewCircularBuffer[T any](size int) (CircularBuffer[T],error) {
    if size<=0 {
        return CircularBuffer[T]{},customerr.ValOutsideRange(
            fmt.Sprintf("Size of queue must be >=0 | Have: %d",size),
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
        m: &sync.RWMutex{},
    },nil;
}

func (c *CircularBuffer[T])Full() bool {
    return c.Length()==c.Capacity()
}

func (c *CircularBuffer[T])Length() int {
    c.m.RLock()
    defer c.m.RUnlock()
    return c.numElems;
}

func (c *CircularBuffer[T])Capacity() int {
    return len(c.vals);
}

func (c *CircularBuffer[T])PushFront(v T) error {
    c.m.Lock()
    defer c.m.Unlock()
    if c.numElems<len(c.vals) {
        c.numElems++;
        c.startEnd.A=c.decIndex(c.startEnd.A);
        c.vals[c.startEnd.A]=v;
        return nil;
    }
    return c.getQueueFullError()
}

func (c *CircularBuffer[T])PushBack(v T) error {
    c.m.Lock()
    defer c.m.Unlock()
    if c.numElems<len(c.vals) {
        c.numElems++;
        c.startEnd.B=c.incIndex(c.startEnd.B);
        c.vals[c.startEnd.B]=v;
        return nil;
    }
    return c.getQueueFullError()
}

func (c *CircularBuffer[T])ForcePushBack(v T) {
    c.m.Lock()
    defer c.m.Unlock()
    if c.numElems==len(c.vals) {
        c.startEnd.A=c.incIndex(c.startEnd.A);
        c.numElems--;
    }
    c.numElems++;
    c.startEnd.B=c.incIndex(c.startEnd.B);
    c.vals[c.startEnd.B]=v;
}

func (c *CircularBuffer[T])ForcePushFront(v T) {
    c.m.Lock()
    defer c.m.Unlock()
    if c.numElems==len(c.vals) {
        c.startEnd.B=c.decIndex(c.startEnd.B);
        c.numElems--;
    }
    c.numElems++;
    c.startEnd.A=c.decIndex(c.startEnd.A);
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
    c.m.RLock()
    defer c.m.RUnlock()
    if c.numElems>0 {
        return &c.vals[c.startEnd.A],nil
    }
    return nil,c.getIndexOutOfBoundsError(0)
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
    c.m.RLock()
    defer c.m.RUnlock()
    if idx>=0 && idx<c.numElems && c.numElems>0 {
        properIndex:=c.getProperIndex(idx)
        return &c.vals[properIndex],nil;
    }
    return nil,c.getIndexOutOfBoundsError(idx)
}

func (c *CircularBuffer[T])Set(v T, idx int) error {
    c.m.Lock()
    defer c.m.Unlock()
    if idx>=0 && idx<c.numElems && c.numElems>0 {
        properIndex:=c.getProperIndex(idx)
        c.vals[properIndex]=v
        return nil
    }
    return c.getIndexOutOfBoundsError(idx)
}

func (c *CircularBuffer[T])Insert(v T, idx int) error {
    c.m.Lock()
    defer c.m.Unlock()
    if idx<0 || idx>c.numElems {
        return c.getIndexOutOfBoundsError(idx)
    } else if c.numElems==len(c.vals) {
        return c.getQueueFullError()
    }
    if c.numElems==0 && idx==0 {
        c.numElems++;
        c.startEnd.B=c.incIndex(c.startEnd.B);
        c.vals[c.startEnd.B]=v;
    } else if c.distanceFromBack(idx)>c.distanceFromFront(idx) {
        c.insertMoveFront(&v,idx)
    } else {
        c.insertMoveBack(&v,idx)
    }
    return nil
}

func (c *CircularBuffer[T])insertMoveFront(v *T, idx int) {
    c.numElems++;
    c.startEnd.A=c.decIndex(c.startEnd.A);
    for i:=0; i<=idx-1; i++ { 
        c.vals[c.getProperIndex(i)]=c.vals[c.getProperIndex(i+1)]
    }
    c.vals[c.getProperIndex(idx)]=*v
}

func (c *CircularBuffer[T])insertMoveBack(v *T, idx int) {
    c.numElems++;
    c.startEnd.B=c.incIndex(c.startEnd.B);
    for i:=c.numElems-2; i>=idx; i-- { 
        c.vals[c.getProperIndex(i+1)]=c.vals[c.getProperIndex(i)]
    }
    c.vals[c.getProperIndex(idx)]=*v
}

func (c *CircularBuffer[T])Append(v T) error {
    return c.PushBack(v)
}
// Static Vector (Array?)
// Dynamic Vector (name?)
// wtf to do about the decrepid double linked list?
// Dynamic Deque

func (c *CircularBuffer[T])PopFront() (T,error) {
    c.m.Lock()
    defer c.m.Unlock()
    if c.numElems>0 {
        rv:=c.vals[c.startEnd.A];
        c.startEnd.A=c.incIndex(c.startEnd.A);
        c.numElems--;
        return rv,nil;
    }
    var tmp T;
    return tmp,QueueEmpty("Nothing to pop!");
}

func (c *CircularBuffer[T])Delete(idx int) error {
    c.m.Lock()
    defer c.m.Unlock()
    if idx<0 || idx>=c.numElems {
        return c.getIndexOutOfBoundsError(idx)
    }
    if c.numElems==1 && idx==0 {
        c.numElems--;
        c.startEnd.B=c.decIndex(c.startEnd.B);
    } else if c.distanceFromBack(idx)>c.distanceFromFront(idx) {
        c.deleteMoveFront(idx)
    } else {
        c.deleteMoveBack(idx)
    }
    return nil
}

func (c *CircularBuffer[T])deleteMoveFront(idx int) {
    c.numElems--;
    c.startEnd.A=c.incIndex(c.startEnd.A);
    for i:=0; i<=idx-1; i++ { 
        c.vals[c.getProperIndex(i+1)]=c.vals[c.getProperIndex(i)]
    }
}

func (c *CircularBuffer[T])deleteMoveBack(idx int) {
    c.numElems--;
    c.startEnd.B=c.decIndex(c.startEnd.B);
    for i:=c.numElems-2; i>=idx; i-- { 
        c.vals[c.getProperIndex(i)]=c.vals[c.getProperIndex(i+1)]
    }
}

func (c *CircularBuffer[T])Clear() {
    c.m.Lock()
    defer c.m.Unlock()
    c.vals=make([]T,len(c.vals))
    c.numElems=0
    c.startEnd=Pair[int, int]{A: 0, B: len(c.vals)-1}
}

func (c *CircularBuffer[T])Elems() iter.Iter[T] {
    i:=-1;
    return func(f iter.IteratorFeedback) (T,error,bool) {
        i++;
        var rv T;
        if i<c.numElems && f!=iter.Break {
            if i==0 {
                c.m.RLock()
            }
            v,err:=c.Get(i);
            return v,err,true;
        } else if i==c.numElems && f!=iter.Break {
            return rv,nil,false;
        }
        if i>0 {
            c.m.RUnlock()
        }
        return rv,nil,false;
    }
}

func (c *CircularBuffer[T])PntrElems() iter.Iter[*T] {
    i:=-1;
    return func(f iter.IteratorFeedback) (*T,error,bool) {
        i++;
        if i<c.numElems && f!=iter.Break {
            if i==0 {
                c.m.RLock()
            }
            v,err:=c.GetPntr(i);
            return v,err,true;
        } else if i==c.numElems && f!=iter.Break {
            return nil,nil,false;
        }
        if i>0 {
            c.m.RUnlock()
        }
        return nil,nil,false;
    }
}

func (c *CircularBuffer[T])incIndex(idx int) int {
    if idx+1>=len(c.vals) {
        return 0
    }
    return idx+1;
}

func (c *CircularBuffer[T])decIndex(idx int) int {
    if idx-1<0 {
        return len(c.vals)-1;
    }
    return idx-1
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

func (c *CircularBuffer[T])getIndexOutOfBoundsError(idx int) error {
    return customerr.ValOutsideRange(fmt.Sprintf(
        "Index out of bounds. | NumElems: %d Index: %d",
        c.numElems,idx,
    ));
}
func (c *CircularBuffer[T])getQueueFullError() error {
    return QueueFull(fmt.Sprintf("Queue size: %d",len(c.vals)));
}
func (c *CircularBuffer[T])getQueueEmptyError() error {
    return QueueEmpty(fmt.Sprintf("Queue size: %d",len(c.vals)));
}
