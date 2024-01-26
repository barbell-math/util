package types;

//This file should never import anything other than the std library. If anything
// else is imported the risk of import loops is very high.

type Widget[T any] interface {
    Eq(l *T, r *T) bool
    Lt(l *T, r *T) bool
    Unwrap() T
}

type Syncable interface {
    Lock()
    Unlock()
    RLock()
    RUnlock()
}

type Capacity interface {
    Length() int;
    Capacity() int;
}

type DynCapacity interface {
    Capacity
    SetCapacity(s int) error
}

type StaticCapacity interface {
    Capacity
    Full() bool
}

// TODO -rename to Writer/Reader???
// TODO - add contains - no, container should not 'act' on its elements. But then wtf is set??
type Read[K any, V any] interface {
    Get(idx K) (V,error);
    GetPntr(idx K) (*V,error);
    // Contains(v V) bool
}

type Write[K any, V any] interface {
    Set(idx K, v V) error;
    Push(idx K, v ...V) error;
    Append(vals ...V) error
}

// TODO - add Pop
type Delete[K any, V any] interface {
    // Pop(v V, n int) error
    Delete(idx K) error
    Clear()
}

type FirstElemRead[V any] interface {
    PeekFront() (V,error);
    PeekPntrFront() (*V,error);
}

type FirstElemWrite[V any] interface {
    PushFront(v V) error;
    ForcePushFront(v V)
}

type FirstElemRemove[V any] interface {
    PopFront() (V,error);
}

type LastElemRead[V any] interface {
    PeekBack() (V,error);
    PeekPntrBack() (*V,error);
}

type LastElemWrite[V any] interface {
    PushBack(v V) (error);
    ForcePushBack(v V)
}

type LastElemRemove[V any] interface {
    PopBack() (V,error);
}

type Equals[O any, V any] interface {
    Eq(other *O, comp func(l *V, r *V) bool) bool
    Neq(other *O, comp func(l *V, r *V) bool) bool
}
