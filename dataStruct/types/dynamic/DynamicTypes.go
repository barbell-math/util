package dynamic

//This file should never import anything other than the std library and the 
// parrent types package. If anything else is imported the risk of import loops
//is very high.
import "github.com/barbell-math/util/dataStruct/types"


// All types in this file must support the dynamic capacity interface.
type ReadVector[T any] interface {
    types.DynCapacity[T]
    types.RandomRead[T]
}
type WriteVector[T any] interface {
    types.DynCapacity[T]
    types.RandomWrite[T]
    types.RandomDelete[T]
}
type Vector[T any] interface {
    ReadVector[T]
    WriteVector[T]
};

type ReadQueue[T any] interface {
    types.DynCapacity[T]
    types.FirstElemRead[T]
}
type WriteQueue[T any] interface {
    types.DynCapacity[T]
    types.FirstElemRemove[T]
    types.LastElemWrite[T]
}
type Queue[T any] interface {
    ReadQueue[T]
    WriteQueue[T]
};

type ReadStack[T any] interface {
    types.DynCapacity[T]
    types.FirstElemRead[T]
}
type WriteStack[T any] interface {
    types.DynCapacity[T]
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
