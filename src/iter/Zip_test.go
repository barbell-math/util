package iter

import (
	"fmt"
	"testing"

	"github.com/barbell-math/util/src/container/basic"
	"github.com/barbell-math/util/src/test"
)

func TestZipBothNoElements(t *testing.T) {
	count, err := Zip[int, string](
		NoElem[int](),
		NoElem[string](),
		func() basic.Pair[int, string] { return basic.Pair[int, string]{} },
	).Count()
	test.Eq(0, count, t)
	test.Nil(err, t)
}

func TestZipEmptyLeftAndNonEmptyRight(t *testing.T) {
	count, err := Zip[int, string](
		SliceElems[int]([]int{0, 1, 2}),
		NoElem[string](),
		func() basic.Pair[int, string] { return basic.Pair[int, string]{} },
	).Count()
	test.Eq(0, count, t)
	test.Nil(err, t)
}

func TestZipEmptyRightAndNonEmptyLeft(t *testing.T) {
	count, err := Zip[string, int](
		NoElem[string](),
		SliceElems[int]([]int{0, 1, 2}),
		func() basic.Pair[string, int] { return basic.Pair[string, int]{} },
	).Count()
	test.Eq(0, count, t)
	test.Nil(err, t)
}

func TestZipRightLessThanLeft(t *testing.T) {
	vals, err := Zip[string, int](
		SliceElems[string]([]string{"0", "1"}),
		SliceElems[int]([]int{0, 1, 2}),
		func() basic.Pair[string, int] { return basic.Pair[string, int]{} },
	).Collect()
	test.Eq(2, len(vals), t)
	test.Nil(err, t)
	for i := 0; i < 2; i++ {
		test.Eq(fmt.Sprintf("%d", i), vals[i].A, t)
		test.Eq(i, vals[i].B, t)
	}
}

func TestZipLeftLessThanRight(t *testing.T) {
	vals, err := Zip[string, int](
		SliceElems[string]([]string{"0", "1", "2"}),
		SliceElems[int]([]int{0, 1}),
		func() basic.Pair[string, int] { return basic.Pair[string, int]{} },
	).Collect()
	test.Eq(2, len(vals), t)
	test.Nil(err, t)
	for i := 0; i < 2; i++ {
		test.Eq(fmt.Sprintf("%d", i), vals[i].A, t)
		test.Eq(i, vals[i].B, t)
	}
}

func TestZipLeftEqualsRight(t *testing.T) {
	vals, err := Zip[string, int](
		SliceElems[string]([]string{"0", "1", "2"}),
		SliceElems[int]([]int{0, 1, 2}),
		func() basic.Pair[string, int] { return basic.Pair[string, int]{} },
	).Collect()
	test.Eq(3, len(vals), t)
	test.Nil(err, t)
	for i := 0; i < 3; i++ {
		test.Eq(fmt.Sprintf("%d", i), vals[i].A, t)
		test.Eq(i, vals[i].B, t)
	}
}
