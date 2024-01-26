package dynamic

//This file should never import anything other than the std library and the 
// parrent types package. If anything else is imported the risk of import loops
//is very high.
import "github.com/barbell-math/util/dataStruct/types"


// All types in this file must support the dynamic capacity interface.
type ReadVector[T any] interface {
    types.DynCapacity
    types.Read[T,int]
}
type WriteVector[T any] interface {
    types.SyncPassThrough
    types.DynCapacity
    types.Write[T,int]
    types.Delete[T,int]
}
type Vector[T any] interface {
    ReadVector[T]
    WriteVector[T]
};

type ReadQueue[T any] interface {
    types.DynCapacity
    types.FirstElemRead[T]
}
type WriteQueue[T any] interface {
    types.SyncPassThrough
    types.DynCapacity
    types.FirstElemRemove[T]
    types.LastElemWrite[T]
}
type Queue[T any] interface {
    ReadQueue[T]
    WriteQueue[T]
};

type ReadStack[T any] interface {
    types.DynCapacity
    types.FirstElemRead[T]
}
type WriteStack[T any] interface {
    types.SyncPassThrough
    types.DynCapacity
    types.FirstElemWrite[T]
    types.FirstElemRemove[T]
}
type Stack[T any] interface {
    ReadStack[T]
    WriteStack[T]
};

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
    types.DynCapacity
    types.Read[T,U]
}
type WriteMap[T any, U any] interface {
    types.SyncPassThrough
    types.DynCapacity
    types.Write[T,U]
    types.Delete[T,U]
}
type Map[T any, U any] interface {
    ReadMap[T,U]
    WriteMap[T,U]
}

type ReadSet[T ~uint32 | ~uint64, U any] interface {
    types.Read[T,U]
}
type WriteSet[T ~uint32 | ~uint64, U any] interface {
    types.SyncPassThrough
    types.Write[T,U]
}
type Set[T ~uint32 | ~uint64, U any] interface {
    ReadSet[T,U]
    WriteSet[T,U]
}
