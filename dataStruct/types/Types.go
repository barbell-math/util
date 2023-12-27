package types;

//This file should never import anything other than the std library. If anything
// else is imported the risk of import loops is very high.

type Capacity[T any] interface {
    Length() int;
    Capacity() int;
}

type DynCapacity[T any] interface {
    Capacity[T]
    SetCapacity(s int) error
}

type StaticCapacity[T any] interface {
    Capacity[T]
    Full() bool
}

type RandomRead[T any] interface {
    Get(idx int) (T,error);
    GetPntr(idx int) (*T,error);
}

type RandomWrite[T any] interface {
    Set(v T, idx int) error;
    Insert(v T, idx int) error;
    Append(v T) error
}

type RandomDelete[T any] interface {
    Delete(idx int) error
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
