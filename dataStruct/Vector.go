// Array/slice reflect
// Static Vector (Array?) - no way to dynamically represent, probably no?
// Dynamic Deque - chunks? realloc based on f/b ratio, seed ratio with New
// Tree, graph...
// Precompiled nfa/dfa - use go:generate comments!
package dataStruct

import (
	"fmt"
	"sync"

	"github.com/barbell-math/util/algo/iter"
	customerr "github.com/barbell-math/util/err"
)


type (
    Vector[T any] []T
    
    SyncedVector[T any] struct {
        *sync.RWMutex
        Vector[T]
    }
)

func NewVector[T any](size int) (Vector[T],error) {
    if size<0 {
        return nil,customerr.ValOutsideRange(
            fmt.Sprintf("Size of vector must be >=0 | Have: %d",size),
        );
    }
    return Vector[T](make(Vector[T],size)),nil
}

func NewSyncedVector[T any](size int) (SyncedVector[T],error) {
    rv,err:=NewVector[T](size)
    return SyncedVector[T]{
        Vector: rv,
        RWMutex: &sync.RWMutex{},
    }, err
}

func (v *Vector[T])Lock() { }
func (v *Vector[T])Unlock() { }
func (v *Vector[T])RLock() { }
func (v *Vector[T])RUnlock() { }

func (v *SyncedVector[T])Lock() { v.RWMutex.Lock() }
func (v *SyncedVector[T])Unlock() { v.RWMutex.Unlock() }
func (v *SyncedVector[T])RLock() { v.RWMutex.RLock() }
func (v *SyncedVector[T])RUnlock() { v.RWMutex.RUnlock() }

func (v *Vector[T])Length() int {
    v.RLock()
    defer v.RUnlock()
    return len(*v)
}

func (v *Vector[T])Capacity() int {
    v.RLock()
    defer v.RUnlock()
    return cap(*v)
}

func (v *Vector[T])SetCapacity(c int) error {
    v.Lock()
    defer v.Unlock()
    tmp:=make(Vector[T],c)
    copy(tmp,*v)
    *v=tmp
    return nil
}

func (v *Vector[T])Get(idx int) (T,error){
    if _v,err:=v.GetPntr(idx); err==nil {
        return *_v,nil
    } else {
        var tmp T
        return tmp,err
    }
}

func (v *Vector[T])GetPntr(idx int) (*T,error){
    v.RLock()
    defer v.RUnlock()
    if idx>=0 && idx<len(*v) && len(*v)>0 {
        return &(*v)[idx],nil
    }
    return nil,getIndexOutOfBoundsError(idx,len(*v))
}

func (v *Vector[T])Set(idx int, val T) error {
    v.Lock()
    defer v.Unlock()
    if idx>=0 && idx<len(*v) && len(*v)>0 {
        (*v)[idx]=val
        return nil
    }
    return getIndexOutOfBoundsError(idx,len(*v))
}

func (v *Vector[T])Append(vals ...T) error {
    v.Lock()
    defer v.Unlock()
    *v=append(*v, vals...)
    return nil
}

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

func (v *Vector[T])Delete(idx int) error {
    v.Lock()
    defer v.Unlock()
    if idx>=0 && idx<len(*v) && len(*v)>0 {
        *v=append((*v)[:idx],(*v)[idx+1:]...)
    }
    return nil
}

func (v *Vector[T])Clear() {
    v.Lock()
    defer v.Unlock()
    *v=make(Vector[T], 0)
}

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

func (v *Vector[T])PeekPntrFront() (*T,error) {
    v.RLock()
    defer v.RUnlock()
    if len(*v)>0 {
        return &(*v)[0],nil
    }
    return nil,getIndexOutOfBoundsError(0,len(*v))
}

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

func (v *Vector[T])PeekPntrBack() (*T,error) {
    v.RLock()
    defer v.RUnlock()
    if len(*v)>0 {
        return &(*v)[len(*v)-1],nil
    }
    return nil,getIndexOutOfBoundsError(0,len(*v))
}

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

func (v *Vector[T])PushBack(val T) error {
    v.Lock()
    defer v.Unlock()
    *v=append(*v, val)
    return nil
}

func (v *Vector[T])PushFront(val T) error {
    v.Lock()
    defer v.Unlock()
    *v=append(Vector[T]{val}, (*v)...)
    return nil
}

func (v *Vector[T])ForcePushBack(val T) {
    v.Lock()
    defer v.Unlock()
    *v=append(*v, val)
}

func (v *Vector[T])ForcePushFront(val T) {
    v.Lock()
    defer v.Unlock()
    *v=append(Vector[T]{val}, (*v)...)
}

// TODO -test
func (v *Vector[T])Elems() iter.Iter[T] {
    return iter.SequentialElems[T](
        len(*v),
        func(i int) (T, error) { return (*v)[i],nil },
    ).SetupTeardown(
        func() error { v.RLock(); return nil },
        func() error { v.RUnlock(); return nil },
    )
}

// TODO -test
func (v *Vector[T])PntrElems() iter.Iter[*T] {
    return iter.SequentialElems[*T](
        len(*v),
        func(i int) (*T, error) { return &(*v)[i],nil },
    ).SetupTeardown(
        func() error { v.RLock(); return nil },
        func() error { v.RUnlock(); return nil },
    )
    // i:=-1;
    // return func(f iter.IteratorFeedback) (*T,error,bool) {
    //     var rv T;
    //     i++;
    //     if i<len(*v) && f!=iter.Break {
    //         if i==0 {
    //             v.RLock()
    //         }
    //         return &(*v)[i],nil,true;
    //     }
    //     if i>0 {
    //         v.RUnlock()
    //     }
    //     return &rv,nil,false;
    // }
}
