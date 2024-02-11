package containers

import (
	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/customerr"
)


func getSizeError(size int) error {
    return customerr.Wrap(
	customerr.ValOutsideRange,
	"Size must be >=0. Got: %d",size,
    )
}

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

func getDuplicateValueError[T any](v T) error {
    return customerr.Wrap(
        containerTypes.Duplicate,
        "The supplied value is already in the container: %v",
        v,
    )
}
