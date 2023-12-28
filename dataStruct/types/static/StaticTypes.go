package static

//This file should never import anything other than the std library and the 
// parrent types package. If anything else is imported the risk of import loops
//is very high.
import "github.com/barbell-math/util/dataStruct/types"


// All types in this file must support the static capacity interface, except for
// the Variant and Pair types.
type ReadVector[T any] interface {
    types.StaticCapacity
    types.Read[T,int]
}
type WriteVector[T any] interface {
    types.SyncPassThrough
    types.StaticCapacity
    types.Write[T,int]
    types.Delete[T,int]
}
type Vector[T any] interface {
    ReadVector[T]
    WriteVector[T]
};

type ReadQueue[T any] interface {
    types.StaticCapacity
    types.FirstElemRead[T]
}
type WriteQueue[T any] interface {
    types.SyncPassThrough
    types.StaticCapacity
    types.FirstElemRemove[T]
    types.LastElemWrite[T]
}
type Queue[T any] interface {
    ReadQueue[T]
    WriteQueue[T]
};

type ReadStack[T any] interface {
    types.StaticCapacity
    types.FirstElemRead[T]
}
type WriteStack[T any] interface {
    types.SyncPassThrough
    types.StaticCapacity
    types.FirstElemWrite[T]
    types.FirstElemRemove[T]
}
type Stack[T any] interface {
    ReadStack[T]
    WriteStack[T]
}

type ReadDeque[T any] interface {
    ReadQueue[T]
    ReadStack[T]
}
type WriteDeque[T any] interface {
    WriteQueue[T]
    WriteStack[T]
}
type Deque[T any] interface {
    ReadDeque[T]
    WriteDeque[T]
}

type ReadMap[T any, U any] interface {
    types.StaticCapacity
    types.Read[T,U]
}
type WriteMap[T any, U any] interface {
    types.SyncPassThrough
    types.StaticCapacity
    types.Write[T,U]
    types.Delete[T,U]
}
type Map[T any, U any] interface {
    ReadMap[T,U]
    WriteMap[T,U]
}

type Variant[T any, U any] interface {
    SetValA(newVal T) Variant[T,U];
    SetValB(newVal U) Variant[T,U];
    HasA() bool;
    HasB() bool;
    ValA() T;
    ValB() U;
    ValAOr(_default T) T;
    ValBOr(_default U) U;
};

type Pair[T any, U any] interface {
    GetA() T
    SetA(v T)
    GetB() U
    SetB(v U)
}
