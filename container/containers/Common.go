package containers

import (
	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/customerr"
)

func getIndexOutOfBoundsError(idx int, upper int, lower int) error {
    return customerr.AppendError(
        customerr.Wrap(
            containerTypes.KeyError,
            "The supplied key (%d) was not found in the container.",
            idx,
        ),
        customerr.Wrap(
            customerr.ValOutsideRange,
            "Index must be >=%d and < %d. Idx: %d",
            lower,upper,idx,
        ),
    )
}
