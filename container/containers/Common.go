// This package hold the concrete implementations of the containers.
package containers

import (
	"fmt"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/customerr"
)


func addressableSafeValIter[T any](
    other containerTypes.ComparisonsOtherConstraint[T],
    iterOp func(index int, val *T) (iter.IteratorFeedback,error),
) {
    if other.IsAddressable() {
        other.ValPntrs().ForEach(iterOp)
    } else {
        other.Vals().ForEach(func(index int, val T) (iter.IteratorFeedback, error) {
            return iterOp(index,&val)
        })
    }
}

func getNonAddressablePanicText(thingName string) string {
    return fmt.Sprintf("A %s is not addressable!",thingName)
}

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
