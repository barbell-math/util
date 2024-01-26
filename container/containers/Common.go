package containers

import (
	"github.com/barbell-math/util/container/widgets"
	"github.com/barbell-math/util/customerr"
)

func getWidgetIFaceImpl[T any, U widgets.WidgetInterface[T]]() U {
    var rv U
    return rv
}

func getIndexOutOfBoundsError(idx int, upper int, lower int) error {
    return customerr.Wrap(
        customerr.ValOutsideRange,
        "Index must be >=%d and < %d. Idx: %d",
        lower,upper,idx,
    )
}
