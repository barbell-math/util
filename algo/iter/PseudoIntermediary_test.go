package iter

import (
	"fmt"
	"github.com/barbell-math/util/test"
	"testing"
)

func TestTake(t *testing.T) {
	cnt, err := SliceElems([]int{1, 2, 3, 4}).Take(0).Count()
	test.Eq(0, cnt, t)
	test.Nil(err, t)
	cnt, err = SliceElems([]int{1, 2, 3, 4}).Take(1).Count()
	test.Eq(1, cnt, t)
	test.Nil(err, t)
	cnt, err = SliceElems([]int{1, 2, 3, 4}).Take(2).Count()
	test.Eq(2, cnt, t)
	test.Nil(err, t)
	cnt, err = SliceElems([]int{1, 2, 3, 4}).Take(4).Count()
	test.Eq(4, cnt, t)
	test.Nil(err, t)
	cnt, err = SliceElems([]int{1, 2, 3, 4}).Take(5).Count()
	test.Eq(4, cnt, t)
	test.Nil(err, t)
	cnt, err = SliceElems([]int{}).Take(1).Count()
	test.Eq(0, cnt, t)
	test.Nil(err, t)
	cnt, err = SliceElems([]int{}).Take(0).Count()
	test.Eq(0, cnt, t)
	test.Nil(err, t)
}

func TestTakeWhile(t *testing.T) {
	cnt, err := SliceElems([]int{1, 2, 3, 4}).TakeWhile(func(val int) bool {
		return val < 3
	}).Count()
	test.Eq(2, cnt, t)
	test.Nil(err, t)
	cnt, err = SliceElems([]int{1, 2, 3, 4}).TakeWhile(func(val int) bool {
		return val < 1
	}).Count()
	test.Eq(0, cnt, t)
	test.Nil(err, t)
	cnt, err = SliceElems([]int{1, 2, 3, 4}).TakeWhile(func(val int) bool {
		return val < 2
	}).Count()
	test.Eq(1, cnt, t)
	test.Nil(err, t)
	cnt, err = SliceElems([]int{1, 2, 3, 4}).TakeWhile(func(val int) bool {
		return val < 5
	}).Count()
	test.Eq(4, cnt, t)
	test.Nil(err, t)
}

func TestSkip(t *testing.T) {
	cnt, err := SliceElems([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).Skip(0).Count()
	test.Eq(9, cnt, t)
	test.Nil(err, t)
	cnt, err = SliceElems([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).Skip(1).Count()
	test.Eq(8, cnt, t)
	test.Nil(err, t)
	cnt, err = SliceElems([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).Skip(2).Count()
	test.Eq(7, cnt, t)
	test.Nil(err, t)
	cnt, err = SliceElems([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).Skip(8).Count()
	test.Eq(1, cnt, t)
	test.Nil(err, t)
	cnt, err = SliceElems([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).Skip(9).Count()
	test.Eq(0, cnt, t)
	test.Nil(err, t)
	cnt, err = SliceElems([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).Skip(10).Count()
	test.Eq(0, cnt, t)
	test.Nil(err, t)
}

func mapIterHelper[T any](elems []T, t *testing.T) {
	mapped, err := Map(SliceElems(elems), func(index int, val T) (string, error) {
		return fmt.Sprintf("%v", val), nil
	}).Collect()
	test.Nil(err, t)
	for i, v := range elems {
		test.Eq(fmt.Sprintf("%v", v), mapped[i], t)
	}
}
func TestMap(t *testing.T) {
	mapIterHelper([]int{1, 2, 3, 4}, t)
	mapIterHelper([]int{1}, t)
	mapIterHelper([]int{}, t)
}

func valToPntrHelper[T any](expected []T, t *testing.T) {
	mapped, err := ValToPntr[T](SliceElems(expected)).Collect()
	test.Nil(err, t)
	test.Eq(len(expected), len(mapped), t)
	for i, v := range expected {
		test.Eq(v, *mapped[i], t)
	}
}
func TestMapToPntr(t *testing.T) {
	valToPntrHelper[int]([]int{}, t)
	valToPntrHelper[int]([]int{0}, t)
	valToPntrHelper[int]([]int{0, 1}, t)
	valToPntrHelper[int]([]int{0, 1, 2, 3, 4}, t)
}

func pntrToValHelper[T any](expected []T, t *testing.T) {
	pntrs := make([]*T, len(expected))
	for i, _ := range expected {
		pntrs[i] = &expected[i]
	}
	mapped, err := PntrToVal[T](SliceElems(pntrs)).Collect()
	test.Nil(err, t)
	test.Eq(len(expected), len(mapped), t)
	for i, v := range expected {
		test.Eq(v, mapped[i], t)
	}
}
func TestMapFromPntr(t *testing.T) {
	pntrToValHelper[int]([]int{}, t)
	pntrToValHelper[int]([]int{0}, t)
	pntrToValHelper[int]([]int{0, 1}, t)
	pntrToValHelper[int]([]int{0, 1, 2, 3, 4}, t)
}

func TestFilter(t *testing.T) {
	cntr, err := SliceElems([]int{1, 2, 3, 4}).Filter(func(index int, val int) bool {
		return val < 3
	}).Count()
	test.Eq(2, cntr, t)
	test.Nil(err, t)
	cntr, err = SliceElems([]int{1, 2, 3, 4}).Filter(func(index int, val int) bool {
		return val < 5
	}).Count()
	test.Eq(4, cntr, t)
	test.Nil(err, t)
	cntr, err = SliceElems([]int{1, 2, 3, 4}).Filter(func(index int, val int) bool {
		return val < 2
	}).Count()
	test.Eq(1, cntr, t)
	test.Nil(err, t)
	cntr, err = SliceElems([]int{1, 2, 3, 4}).Filter(func(index int, val int) bool {
		return val < 1
	}).Count()
	test.Eq(0, cntr, t)
	test.Nil(err, t)
}
