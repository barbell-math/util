package dataStruct

import (
	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/dataStruct/types"
)

type Iterable[T any] interface {
    Elems() iter.Iter[T]
    PntrElems() iter.Iter[*T]
    Collect(i iter.Iter[T]) error
}

func AppendCollector[T any, U any](
    v interface {types.Write[T,U]; types.SyncPassThrough}, 
    val U,
) error {
    return v.Append(val)
}

