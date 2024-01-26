package containers

import (
	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/customerr"
)

type (
    WidgetConstraint[T any, U any] interface {
        *U
        containerTypes.Widget[T]
    }
)

func getIndexOutOfBoundsError(idx int, upper int, lower int) error {
    return customerr.Wrap(
        customerr.ValOutsideRange,
        "Index must be >=%d and < %d. Idx: %d",
        lower,upper,idx,
    )
}
