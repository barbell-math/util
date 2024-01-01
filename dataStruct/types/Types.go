package types;

//This file should never import anything other than the std library. If anything
// else is imported the risk of import loops is very high.

type SyncPassThrough interface {
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

type Read[T any, U any] interface {
    Get(idx U) (T,error);
    GetPntr(idx U) (*T,error);
}

type Write[T any, U any] interface {
    Set(idx U, v T) error;
    Insert(idx U, v ...T) error;
    Append(vals ...T) error
}

type Delete[T any, U any] interface {
    Delete(idx U) error
    Clear()
}

type FirstElemRead[T any] interface {
    PeekFront() (T,error);
    PeekPntrFront() (*T,error);
}

type FirstElemWrite[T any] interface {
    PushFront(v T) error;
    ForcePushFront(v T)
}

type FirstElemRemove[T any] interface {
    PopFront() (T,error);
}

type LastElemRead[T any] interface {
    PeekBack() (T,error);
    PeekPntrBack() (*T,error);
}

type LastElemWrite[T any] interface {
    PushBack(v T) (error);
    ForcePushBack(v T)
}

type LastElemRemove[T any] interface {
    PopBack() (T,error);
}
