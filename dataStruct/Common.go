package dataStruct

import "github.com/barbell-math/util/algo/iter"

type Iterable[T any] interface {
    Elems() iter.Iter[T]
    PntrElems() iter.Iter[*T]
}
