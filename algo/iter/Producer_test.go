package iter

import (
	"fmt"
	"testing"

	"github.com/barbell-math/util/test"
)

func valElemIterHelper[T any](val T, err error, r int, t *testing.T) {
	var tmp T
	iter := ValElem(val, err, r)
	for i := 0; i < r; i++ {
		vIter, eIter, contIter := iter(Continue)
		if i < r {
			test.Eq(val, vIter, t)
			test.Eq(err, eIter, t)
			test.True(contIter, t)
		} else {
			test.Eq(tmp, vIter, t)
			test.Nil(eIter, t)
			test.False(contIter, t)
		}
	}
}
func TestValElem(t *testing.T) {
	valElemIterHelper(1, nil, 1, t)
	valElemIterHelper(2, fmt.Errorf("NEW ERROR"), 1, t)
	valElemIterHelper(1, nil, 2, t)
	valElemIterHelper(2, fmt.Errorf("NEW ERROR"), 2, t)
	valElemIterHelper(1, nil, 5, t)
	valElemIterHelper(2, fmt.Errorf("NEW ERROR"), 5, t)
}

func TestRange(t *testing.T) {
	res, err := Range[int](0, 0, 0).Collect()
	test.SlicesMatch[int]([]int{}, res, t)
	test.Nil(err, t)
	res, err = Range[int](0, 0, 1).Collect()
	test.SlicesMatch[int]([]int{}, res, t)
	test.Nil(err, t)

	res, err = Range[int](0, 1, 1).Collect()
	test.SlicesMatch[int]([]int{0}, res, t)
	test.Nil(err, t)
	res, err = Range[int](0, 2, 1).Collect()
	test.SlicesMatch[int]([]int{0, 1}, res, t)
	test.Nil(err, t)
	res, err = Range[int](0, 5, 1).Collect()
	test.SlicesMatch[int]([]int{0, 1, 2, 3, 4}, res, t)
	test.Nil(err, t)
	res, err = Range[int](1, 0, -1).Collect()
	test.SlicesMatch[int]([]int{1}, res, t)
	test.Nil(err, t)
	res, err = Range[int](1, -1, -1).Collect()
	test.SlicesMatch[int]([]int{1, 0}, res, t)
	test.Nil(err, t)
	res, err = Range[int](1, -4, -1).Collect()
	test.SlicesMatch[int]([]int{1, 0, -1, -2, -3}, res, t)
	test.Nil(err, t)

	res, err = Range[int](0, 1, 2).Collect()
	test.SlicesMatch[int]([]int{0}, res, t)
	test.Nil(err, t)
	res, err = Range[int](1, 0, -2).Collect()
	test.SlicesMatch[int]([]int{1}, res, t)
	test.Nil(err, t)

	res, err = Range[int](0, 4, 2).Collect()
	test.SlicesMatch[int]([]int{0, 2}, res, t)
	test.Nil(err, t)
	res, err = Range[int](0, -4, -2).Collect()
	test.SlicesMatch[int]([]int{0, -2}, res, t)
	test.Nil(err, t)
}

func sliceElemsIterHelper[T any](vals []T, t *testing.T) {
	sIter := SliceElems(vals)
	for i := 0; i < len(vals); i++ {
		sV, sErr, sBool := sIter(Continue)
		test.Eq(vals[i], sV, t)
		test.Nil(sErr, t)
		test.True(sBool, t)
	}
	var tmp T
	sV, sErr, sBool := sIter(Continue)
	test.Eq(tmp, sV, t)
	test.Nil(sErr, t)
	test.False(sBool, t)
}
func TestSliceElems(t *testing.T) {
	sliceElemsIterHelper([]string{"one", "two", "three"}, t)
	sliceElemsIterHelper([]int{1, 2, 3}, t)
	sliceElemsIterHelper([]int{1}, t)
	sliceElemsIterHelper([]int{}, t)
}

func sliceElemPntrsIterHelper[T any](vals []T, t *testing.T) {
	sIter := SliceElemPntrs(vals)
	for i := 0; i < len(vals); i++ {
		sV, sErr, sBool := sIter(Continue)
		test.Eq(&vals[i], sV, t)
		test.Eq(vals[i], *sV, t)
		test.Nil(sErr, t)
		test.True(sBool, t)
	}
	sV, sErr, sBool := sIter(Continue)
	test.NilPntr[T](sV, t)
	test.Nil(sErr, t)
	test.False(sBool, t)
}
func TestSliceElemPntrs(t *testing.T) {
	sliceElemPntrsIterHelper([]string{"one", "two", "three"}, t)
	sliceElemPntrsIterHelper([]int{1, 2, 3}, t)
	sliceElemPntrsIterHelper([]int{1}, t)
	sliceElemPntrsIterHelper([]int{}, t)
}

func stringElemsIterHelper(vals string, t *testing.T) {
	sIter := StrElems(vals)
	for i := 0; i < len(vals); i++ {
		sV, sErr, sBool := sIter(Continue)
		test.Eq(vals[i], sV, t)
		test.Nil(sErr, t)
		test.True(sBool, t)
	}
	var tmp string
	sV, sErr, sBool := sIter(Continue)
	test.Eq(tmp, sV, t)
	test.Nil(sErr, t)
	test.False(sBool, t)
}
func TestStringElems(t *testing.T) {
	sliceElemsIterHelper([]string{"one", "two", "three"}, t)
	sliceElemsIterHelper([]int{1, 2, 3}, t)
	sliceElemsIterHelper([]int{1}, t)
	sliceElemsIterHelper([]int{}, t)
}

func sequentialElemsIterHelper[T any](vals []T, t *testing.T) {
	sIter := SequentialElems(len(vals), func(i int) (T, error) { return vals[i], nil })
	for i := 0; i < len(vals); i++ {
		sV, sErr, sBool := sIter(Continue)
		test.Eq(vals[i], sV, t)
		test.Nil(sErr, t)
		test.True(sBool, t)
	}
	var tmp T
	sV, sErr, sBool := sIter(Continue)
	test.Eq(tmp, sV, t)
	test.Nil(sErr, t)
	test.False(sBool, t)
}
func TestSequentialElems(t *testing.T) {
	sequentialElemsIterHelper([]string{"one", "two", "three"}, t)
	sequentialElemsIterHelper([]int{1, 2, 3}, t)
	sequentialElemsIterHelper([]int{1}, t)
	sequentialElemsIterHelper([]int{}, t)
}

func testChanIterHelper(chanNum int, t *testing.T) {
	c := make(chan int)
	go func(c chan int, numElems int) {
		for i := 0; i < numElems; i++ {
			c <- i
		}
		close(c)
	}(c, chanNum)
	cnt, err := ChanElems(c).Count()
	test.Eq(chanNum, cnt, t)
	test.Nil(err, t)
}
func TestChanElems(t *testing.T) {
	testChanIterHelper(0, t)
	testChanIterHelper(1, t)
	testChanIterHelper(5, t)
	testChanIterHelper(20, t)
}

func testFileLinesHelper(numLines int, path string, t *testing.T) {
	fIter := FileLines(fmt.Sprintf("./testData/%s", path))
	for i := 0; i < numLines; i++ {
		fV, fErr, fBool := fIter(Continue)
		test.Eq(fmt.Sprintf("%d", i+1), fV, t)
		test.Nil(fErr, t)
		test.True(fBool, t)
	}
	fV, fErr, fBool := fIter(Continue)
	test.Eq("", fV, t)
	test.Nil(fErr, t)
	test.False(fBool, t)
}
func TestFileLines(t *testing.T) {
	testFileLinesHelper(0, "emptyFile.txt", t)
	testFileLinesHelper(1, "oneLine.txt", t)
	testFileLinesHelper(3, "threeLines.txt", t)
}

func TestRecurseEmpty(t *testing.T) {
	v, err := Recurse[int](
		NoElem[int](),
		func(v int) bool { return true },
		func(v int) Iter[int] { return NoElem[int]() },
	).Collect()
	test.Eq(0, len(v), t)
	test.Nil(err, t)
}

func TestRecurseSingleValue(t *testing.T) {
	v, err := Recurse[int](
		ValElem[int](0, nil, 1),
		func(v int) bool { return false },
		func(v int) Iter[int] { return NoElem[int]() },
	).Collect()
	test.Eq(1, len(v), t)
	test.Eq(0, v[0], t)
	test.Nil(err, t)
}

func TestRecurseSingleValueWithEmptyRecurse(t *testing.T) {
	v, err := Recurse[int](
		ValElem[int](0, nil, 1),
		func(v int) bool { return true },
		func(v int) Iter[int] { return NoElem[int]() },
	).Collect()
	test.Eq(1, len(v), t)
	test.Eq(0, v[0], t)
	test.Nil(err, t)
}

func TestRecurseSingleValueWithSingleValueRecurse(t *testing.T) {
	v, err := Recurse[int](
		ValElem[int](0, nil, 1),
		func(v int) bool { return v == 0 },
		func(v int) Iter[int] { return ValElem[int](1, nil, 1) },
	).Collect()
	test.Eq(2, len(v), t)
	test.Eq(0, v[0], t)
	test.Eq(1, v[1], t)
	test.Nil(err, t)
}

func TestRecurse(t *testing.T) {
	vals, err := Recurse[int](
		SliceElems[int]([]int{0, 1, 2}),
		func(v int) bool { return v == 0 || v == 1 || v == 3 },
		func(v int) Iter[int] {
			if v == 0 {
				return SliceElems[int]([]int{3, 5})
			} else if v == 1 {
				return SliceElems[int]([]int{7, 9})
			} else if v == 3 {
				return SliceElems[int]([]int{11, 13})
			}
			return NoElem[int]()
		},
	).Collect()
	exp := []int{0, 3, 11, 13, 5, 1, 7, 9, 2}
	test.Eq(len(exp), len(vals), t)
	for i, v := range vals {
		test.Eq(v, exp[i], t)
	}
	test.Nil(err, t)
}
