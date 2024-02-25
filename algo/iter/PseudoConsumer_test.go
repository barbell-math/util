package iter

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/barbell-math/util/test"
)

func collectIterHelper[T any](vals []T, t *testing.T) {
	collected, err := SliceElems(vals).Collect()
	test.SlicesMatch(vals, collected, t)
	test.Nil(err, t)
}
func TestCollect(t *testing.T) {
	collectIterHelper([]int{1, 2, 3, 4}, t)
	collectIterHelper([]int{1}, t)
	collectIterHelper([]int{}, t)
}

func appendIterHelper[T any](orig []T, vals []T, t *testing.T) {
	origLen := len(orig)
	expLen := len(orig) + len(vals)
	tmp, err := SliceElems(vals).AppendTo(&orig)
	test.Eq(expLen, len(orig), t)
	test.Eq(tmp, len(vals), t)
	test.Nil(err, t)
	for i := origLen; i < len(orig); i++ {
		test.Eq(vals[i-origLen], orig[i], t)
	}
}
func TestAppendTo(t *testing.T) {
	appendIterHelper([]int{1, 2, 3, 4}, []int{1, 2, 3, 4}, t)
	appendIterHelper([]int{1, 2, 3, 4}, []int{1}, t)
	appendIterHelper([]int{1, 2, 3, 4}, []int{}, t)
	appendIterHelper([]int{1}, []int{1, 2, 3, 4}, t)
	appendIterHelper([]int{}, []int{1, 2, 3, 4}, t)
	appendIterHelper([]int{}, []int{}, t)
}

func TestAll(t *testing.T) {
	res, err := SliceElems([]int{1, 2, 3, 4}).All(func(val int) (bool, error) {
		return val > 0, nil
	})
	test.True(res, t)
	test.Nil(err, t)
	res, err = SliceElems([]int{1, 2, 3, 4}).All(func(val int) (bool, error) {
		return val < 0, nil
	})
	test.False(res, t)
	test.Nil(err, t)
	res, err = SliceElems([]int{1, 2, 3, 4}).All(func(val int) (bool, error) {
		return val < 2, nil
	})
	test.False(res, t)
	test.Nil(err, t)
}

func TestAny(t *testing.T) {
	found, err := SliceElems([]int{1, 2, 3, 4}).Any(func(val int) (bool, error) {
		return val > 0, nil
	})
	test.True(found, t)
	test.Nil(err, t)
	found, err = SliceElems([]int{1, 2, 3, 4}).Any(func(val int) (bool, error) {
		return val < 2, nil
	})
	test.True(found, t)
	test.Nil(err, t)
	found, err = SliceElems([]int{1, 2, 3, 4}).Any(func(val int) (bool, error) {
		return val < 0, nil
	})
	test.False(found, t)
	test.Nil(err, t)
}

func findIterHelperFound[T comparable](elems []T, lookingFor T, t *testing.T) {
	v, err, ok := SliceElems(elems).Find(func(val T) (bool, error) {
		return val == lookingFor, nil
	})
	test.Eq(lookingFor, v, t)
	test.True(ok, t)
	test.Nil(err, t)
}
func findIterHelperNotFound[T comparable](elems []T, lookingFor T, t *testing.T) {
	var tmp T
	v, err, ok := SliceElems(elems).Find(func(val T) (bool, error) {
		return val == lookingFor, nil
	})
	test.Eq(tmp, v, t)
	test.False(ok, t)
	test.Nil(err, t)
}
func TestFind(t *testing.T) {
	findIterHelperFound([]int{1, 2, 3, 4}, 1, t)
	findIterHelperFound([]int{1, 2, 3, 4}, 4, t)
	findIterHelperNotFound([]int{1, 2, 3, 4}, 5, t)
	findIterHelperFound([]int{1}, 1, t)
	findIterHelperNotFound([]int{1}, 5, t)
}

func indexIterHelper[T comparable](elems []T,
	lookingFor T,
	expectedIndex int,
	t *testing.T) {
	v, err := SliceElems(elems).Index(func(val T) (bool, error) {
		return val == lookingFor, nil
	})
	test.Eq(expectedIndex, v, t)
	test.Nil(err, t)
}
func TestIndex(t *testing.T) {
	indexIterHelper([]int{1, 2, 3, 4}, 1, 0, t)
	indexIterHelper([]int{1, 2, 3, 4}, 3, 2, t)
	indexIterHelper([]int{1, 2, 3, 4}, 4, 3, t)
	indexIterHelper([]int{1, 2, 3, 4}, 5, -1, t)
	indexIterHelper([]int{1}, 1, 0, t)
	indexIterHelper([]int{1}, 2, -1, t)
	indexIterHelper([]int{}, 1, -1, t)
}

func TestIndexErrorFound(t *testing.T) {
	cntr := 0
	v, err := SliceElems([]int{1, 2, 3, 4}).Index(func(val int) (bool, error) {
		cntr++
		if val == 3 {
			return true, errors.New("")
		}
		return false, nil
	})
	test.Eq(3, cntr, t)
	test.Eq(2, v, t)
	test.NotNil(err, t)
}

func TestIndexErrorNotFound(t *testing.T) {
	cntr := 0
	v, err := SliceElems([]int{1, 2, 3, 4}).Index(func(val int) (bool, error) {
		cntr++
		if val == 3 {
			return false, errors.New("")
		}
		return false, nil
	})
	test.Eq(3, cntr, t)
	test.Eq(-1, v, t)
	test.NotNil(err, t)
}

func nthIterHelper[T any](
	vals []T,
	index int,
	expectedVal T,
	expectedError bool,
	t *testing.T) {
	val, err, ok := SliceElems(vals).Nth(index)
	test.Eq(expectedVal, val, t)
	test.Eq(expectedError, ok, t)
	test.Nil(err, t)
}
func TestNth(t *testing.T) {
	nthIterHelper([]int{1, 2, 3, 4}, 0, 1, true, t)
	nthIterHelper([]int{1, 2, 3, 4}, 2, 3, true, t)
	nthIterHelper([]int{1, 2, 3, 4}, 3, 4, true, t)
	nthIterHelper([]int{1, 2, 3, 4}, 4, 0, false, t)
	nthIterHelper([]int{1}, 0, 1, true, t)
	nthIterHelper([]int{1}, 1, 0, false, t)
	nthIterHelper([]int{}, 0, 0, false, t)
}

func TestCount(t *testing.T) {
	c, err := SliceElems([]int{1, 2, 3, 4}).Count()
	test.Eq(4, c, t)
	test.Nil(err, t)
	c, err = SliceElems([]int{1}).Count()
	test.Eq(1, c, t)
	test.Nil(err, t)
	c, err = SliceElems([]int{}).Count()
	test.Eq(0, c, t)
	test.Nil(err, t)
}

func toChanIterHelper[T any](vals []T, t *testing.T) {
	cntr := 0
	c := make(chan T)
	go func() {
		for val := range c {
			test.Eq(cntr, val, t)
			cntr++
		}
	}()
	SliceElems(vals).ToChan(c)
	close(c)
}
func TestToChan(t *testing.T) {
	vals := make([]int, 200)
	for i := 0; i < 200; i++ {
		vals[i] = i
	}
	toChanIterHelper([]int{}, t)
	toChanIterHelper([]int{0}, t)
	toChanIterHelper(vals, t)
}

func toFileIterHelperWithNewline(numVals int, src string, t *testing.T) {
	vals := make([]int, numVals)
	for i, _ := range vals {
		vals[i] = i
	}
	f, err := os.OpenFile(src, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	SliceElems(vals).ToWriter(f, true)
	f.Close()
	f, err = os.Open(src)
	test.Nil(err, t)
	w := bufio.NewScanner(f)
	for i := 0; w.Scan(); i++ {
		test.Eq(fmt.Sprintf("%d", i), w.Text(), t)
	}
	err = os.Remove(src)
	test.Nil(err, t)
}
func toFileIterHelperNoNewline(numVals int, src string, t *testing.T) {
	correctVal := ""
	vals := make([]int, numVals)
	for i, _ := range vals {
		vals[i] = i
		correctVal = fmt.Sprintf("%s%d", correctVal, i)
	}
	f, err := os.OpenFile(src, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	SliceElems(vals).ToWriter(f, false)
	f.Close()
	f, err = os.Open(src)
	test.Nil(err, t)
	w := bufio.NewScanner(f)
	for i := 0; w.Scan(); i++ {
		test.Eq(correctVal, w.Text(), t)
	}
	err = os.Remove(src)
	test.Nil(err, t)
}
func TestToFile(t *testing.T) {
	toFileIterHelperWithNewline(0, "emptyFileTest.txt", t)
	toFileIterHelperWithNewline(1, "oneLineFileTest.txt", t)
	toFileIterHelperWithNewline(5, "fiveLinesFileTest.txt", t)
	toFileIterHelperWithNewline(10, "tenLinesFileTest.txt", t)
	toFileIterHelperNoNewline(0, "emptyFileTest.txt", t)
	toFileIterHelperNoNewline(1, "oneLineFileTest.txt", t)
	toFileIterHelperNoNewline(5, "fiveLinesFileTest.txt", t)
	toFileIterHelperNoNewline(10, "tenLinesFileTest.txt", t)
}

func TestReduce(t *testing.T) {
	i := 0
	newErr := fmt.Errorf("NEW ERROR")
	tmp, err := SliceElems([]int{1, 2, 3, 4}).Reduce(0, func(accum *int, iter int) error {
		*accum = *accum + iter
		return nil
	})
	test.Eq(10, tmp, t)
	test.Nil(err, t)
	tmp, err = SliceElems([]int{1, 2, 3, 4}).Reduce(0, func(accum *int, iter int) error {
		i++
		return newErr
	})
	test.Eq(1, i, t)
	test.Eq(newErr, err, t)
}
