package dataStruct

import (
	"fmt"
	"sync"

	"github.com/barbell-math/util/algo/iter"
	customerr "github.com/barbell-math/util/err"
)

type (
    Deque[T any] struct {
        vals []T
    }

    SyncedDeque[T any] struct {
        *sync.RWMutex
        Deque[T]
    }
)

func NewDeque[T any](size int) (Deque[T],error) {
    if size<0 {
        return Deque[T]{},customerr.ValOutsideRange(
            fmt.Sprintf("Size of deque must be >=0 | Have: %d",size),
        );
    }
    return Deque[T]{
        vals: make([]T,size),
    },nil
}

func NewSyncedDeque[T any](size int) (SyncedDeque[T],error) {
    rv,err:=NewDeque[T](size)
    return SyncedDeque[T]{
        Deque: rv,
        RWMutex: &sync.RWMutex{},
    }, err
}

func (d *Deque[T])Lock() { }
func (d *Deque[T])Unlock() { }
func (d *Deque[T])RLock() { }
func (d *Deque[T])RUnlock() { }

func (d *SyncedDeque[T])Lock() { d.RWMutex.Lock() }
func (d *SyncedDeque[T])Unlock() { d.RWMutex.Unlock() }
func (d *SyncedDeque[T])RLock() { d.RWMutex.RLock() }
func (d *SyncedDeque[T])RUnlock() { d.RWMutex.RUnlock() }

func (d *Deque[T])Length() int {
    d.RLock()
    defer d.RUnlock()
    return len(d.vals)
}

func (d *Deque[T])Capacity() int {
    d.RLock()
    defer d.RUnlock()
    return 0
}

func (d *Deque[T])SetCapacity(c int) error {
    d.Lock()
    defer d.Unlock()
    // tmp:=make(Vector[T],c)
    // copy(tmp,*v)
    // *v=tmp
    return nil
}

func (d *Deque[T])Get(idx int) (T,error){
    if _v,err:=d.GetPntr(idx); err==nil {
        return *_v,nil
    } else {
        var tmp T
        return tmp,err
    }
}

func (d *Deque[T])GetPntr(idx int) (*T,error){
    d.RLock()
    defer d.RUnlock()
    if idx>=0 && idx<len(d.vals) && len(d.vals)>0 {
        // return &(*v)[idx],nil
    }
    return nil,getIndexOutOfBoundsError(idx,len(d.vals))
}

func (d *Deque[T])Set(val T, idx int) error {
    d.Lock()
    defer d.Unlock()
    if idx>=0 && idx<len(d.vals) && len(d.vals)>0 {
        // (*v)[idx]=val
        // return nil
    }
    return getIndexOutOfBoundsError(idx,len(d.vals))
}

func (d *Deque[T])Append(vals ...T) error {
    d.Lock()
    defer d.Unlock()
    // *v=append(*v, vals...)
    return nil
}

func (d *Deque[T])Insert(idx int, vals ...T) error {
    d.Lock()
    defer d.Unlock()
    if idx>=0 && idx<len(d.vals) && len(d.vals)>0 {
        // *v=append((*v)[:idx],append(vals,(*v)[idx:]...)...) 
        return nil
    } else if idx==len(d.vals) {
        // *v=append(*v,vals...)
        return nil
    }
    return getIndexOutOfBoundsError(idx,len(d.vals))
}

func (d *Deque[T])Delete(idx int) error {
    d.Lock()
    defer d.Unlock()
    if idx>=0 && idx<len(d.vals) && len(d.vals)>0 {
        // *v=append((*v)[:idx],(*v)[idx+1:]...)
    }
    return nil
}

func (d *Deque[T])Clear() {
    d.Lock()
    defer d.Unlock()
    // *v=make(Vector[T], 0)
}

func (d *Deque[T])PeekFront() (T,error) {
    d.RLock()
    defer d.RUnlock()
    if _v,err:=d.PeekPntrFront(); err==nil {
        return *_v,err
    } else {
        var tmp T
        return tmp,err
    }
}

func (d *Deque[T])PeekPntrFront() (*T,error) {
    d.RLock()
    defer d.RUnlock()
    if len(d.vals)>0 {
        // return &(*v)[0],nil
    }
    return nil,getIndexOutOfBoundsError(0,len(d.vals))
}

func (d *Deque[T])PeekBack() (T,error) {
    d.RLock()
    defer d.RUnlock()
    if _v,err:=d.PeekPntrBack(); err==nil {
        return *_v,err
    } else {
        var tmp T
        return tmp,err
    }
}

func (d *Deque[T])PeekPntrBack() (*T,error) {
    d.RLock()
    defer d.RUnlock()
    if len(d.vals)>0 {
        // return &(*v)[len(*v)-1],nil
    }
    return nil,getIndexOutOfBoundsError(0,len(d.vals))
}

func (d *Deque[T])PopFront() (T,error) {
    d.Lock()
    defer d.Unlock()
    if len(d.vals)>0 {
        // rv:=(*v)[0]
        // *v=(*v)[1:]
        // return rv,nil
    }
    var tmp T
    return tmp,Empty("Nothing to pop!")
}

func (d *Deque[T])PopBack() (T,error) {
    d.Lock()
    defer d.Unlock()
    if len(d.vals)>0 {
        // rv:=(*v)[len(*v)-1]
        // *v=(*v)[:len(*v)-1]
        // return rv,nil
    }
    var tmp T
    return tmp,Empty("Nothing to pop!")
}

func (d *Deque[T])PushBack(val T) error {
    d.Lock()
    defer d.Unlock()
    // *v=append(*v, val)
    return nil
}

func (d *Deque[T])PushFront(val T) error {
    d.Lock()
    defer d.Unlock()
    // *v=append(Vector[T]{val}, (*v)...)
    return nil
}

func (d *Deque[T])ForcePushBack(val T) {
    d.Lock()
    defer d.Unlock()
    // *v=append(*v, val)
}

func (d *Deque[T])ForcePushFront(val T) {
    d.Lock()
    defer d.Unlock()
    // *v=append(Vector[T]{val}, (*v)...)
}

func (d *Deque[T])Elems() iter.Iter[T] {
    i:=-1;
    return func(f iter.IteratorFeedback) (T,error,bool) {
        var rv T;
        i++;
        if i<len(d.vals) && f!=iter.Break {
            if i==0 {
                d.RLock()
            }
            // return (*v)[i],nil,true;
        }
        if i>0 {
            d.RUnlock()
        }
        return rv,nil,false;
    }
}

func (d *Deque[T])PntrElems() iter.Iter[*T] {
    i:=-1;
    return func(f iter.IteratorFeedback) (*T,error,bool) {
        var rv T;
        i++;
        if i<len(d.vals) && f!=iter.Break {
            if i==0 {
                d.RLock()
            }
            // return &(*v)[i],nil,true;
        }
        if i>0 {
            d.RUnlock()
        }
        return &rv,nil,false;
    }
}
