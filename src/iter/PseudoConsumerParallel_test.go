package iter

import (
	"github.com/barbell-math/util/src/customerr"
	"github.com/barbell-math/util/src/test"
	"testing"
)

func filterParallelHelper(vals []int, numThreads int, t *testing.T) {
	rv, err := SliceElems(vals).FilterParallel(func(val int) bool {
		return val == 1 || val == 2
	}, numThreads)
	test.Nil(err, t)
	for _, v := range rv {
		test.False(v != 1 && v != 2, t)
	}
}
func TestFilterParallel(t *testing.T) {
	_, err := SliceElems([]int{1, 2, 3, 4}).FilterParallel(func(val int) bool {
		return false
	}, 0)
	test.ContainsError(customerr.ValOutsideRange, err, t)
	vals := make([]int, 200)
	for i := 0; i < 200; i++ {
		vals[i] = i
	}
	for _, i := range []int{1, 25, 50, 75, 100} {
		filterParallelHelper([]int{}, i, t)
		filterParallelHelper([]int{1}, i, t)
		filterParallelHelper(vals, i, t)
	}
}
