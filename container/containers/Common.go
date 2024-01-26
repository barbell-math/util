package containers

import (
	"github.com/barbell-math/util/customerr"
)

func getIndexOutOfBoundsError(idx int, upper int, lower int) error {
    return customerr.Wrap(
        customerr.ValOutsideRange,
        "Index must be >=%d and < %d. Idx: %d",
        lower,upper,idx,
    )
}
