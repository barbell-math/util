package iter

import (
	"github.com/barbell-math/util/src/customerr"
	"github.com/barbell-math/util/src/test"
	"testing"
)

func parallelIterHelper(numVals int, numThreads int, t *testing.T) {
	vals := make([]int, numVals)
	for i := 0; i < numVals; i++ {
		vals[i] = i
	}

	cpy := make([]int, len(vals))
	rv := SliceElems(vals).Parallel(
		func(val int) (int, error) {
			return val + 1, nil
		},
		func(val int, res int, err error) {
			test.Eq(val+1, res, t)
			test.Nil(err, t)
			cpy[val] = res + 1
		},
		numThreads,
	)
	test.Nil(rv, t)
	for i, v := range cpy {
		test.Eq(vals[i]+2, v, t)
	}
}
func TestParallel(t *testing.T) {
	rv := SliceElems([]int{1, 2, 3, 4}).Parallel(func(val int) (int, error) {
		return 0, nil
	}, NoOp[int, int], 0)
	test.ContainsError(customerr.ValOutsideRange, rv, t)

	for _, i := range []int{1, 25, 50, 75, 100} {
		parallelIterHelper(0, i, t)
		parallelIterHelper(1, i, t)
		parallelIterHelper(200, i, t)
	}
}
