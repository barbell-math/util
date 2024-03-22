package iter

import (
	"errors"
	"fmt"
	"github.com/barbell-math/util/test"
	"testing"
)

func TestStop(t *testing.T) {
	cntr := 0
	newErr := fmt.Errorf("Reached break")
	err := SliceElems([]int{1, 2, 3, 4}).Next(
		func(index, val int, status IteratorFeedback) (IteratorFeedback, int, error) {
			cntr++
			if status == Break {
				return Break, 0, newErr
			}
			return Continue, val, nil
		},
	).Stop()
	test.Eq(1, cntr, t)
	test.Eq(newErr, err, t)
}

func TestPullOne(t *testing.T) {
	_iter:=SliceElems[int]([]int{1,2,3,4})
	for i:=0; i<4; i++ {
		v,err,ok:=_iter.PullOne()
		test.Eq(i+1,v,t)
		test.Nil(err,t)
		test.True(ok,t)
	}
	v,err,ok:=_iter.PullOne()
	test.Eq(0,v,t)
	test.Nil(err,t)
	test.False(ok,t)
}

func TestPull(t *testing.T) {
	_iter:=SliceElems[int]([]int{1,2,3,4})
	for i:=0; i<4; i++ {
		v,err,ok:=_iter.Pull(1)
		test.Eq(i+1,v[0],t)
		test.Nil(err,t)
		test.True(ok,t)
	}
	v,err,ok:=_iter.Pull(1)
	test.Eq(0,v[0],t)
	test.Nil(err,t)
	test.False(ok,t)

	_iter=SliceElems[int]([]int{1,2,3,4})
	for i:=0; i<4; i+=2 {
		v,err,ok:=_iter.Pull(2)
		test.Eq(i+1,v[0],t)
		test.Eq(i+2,v[1],t)
		test.Nil(err,t)
		test.True(ok,t)
	}
	v,err,ok=_iter.Pull(2)
	test.Eq(0,v[0],t)
	test.Eq(0,v[1],t)
	test.Nil(err,t)
	test.False(ok,t)

	_iter=SliceElems[int]([]int{1,2,3,4})
	v,err,ok=_iter.Pull(4)
	for i:=0; i<4; i++ {
		test.Eq(i+1,v[i],t)
	}
	test.Nil(err,t)
	test.True(ok,t)
	v,err,ok=_iter.Pull(4)
	for i:=0; i<4; i++ {
		test.Eq(0,v[i],t)
	}
	test.Nil(err,t)
	test.False(ok,t)

	_iter=SliceElems[int]([]int{1,2,3,4})
	v,err,ok=_iter.Pull(5)
	for i:=0; i<4; i++ {
		test.Eq(i+1,v[i],t)
	}
	test.Eq(0,v[4],t)
	test.Nil(err,t)
	test.False(ok,t)
	v,err,ok=_iter.Pull(5)
	for i:=0; i<5; i++ {
		test.Eq(0,v[i],t)
	}
	test.Nil(err,t)
	test.False(ok,t)
}

func forEachIterHelper[T any](
	vals []T,
	op func(index int, val T) T,
	t *testing.T) {
	i := 0
	cpy := append([]T{}, vals...)
	err := SliceElems(vals).ForEach(
		func(index int, val T) (IteratorFeedback, error) {
			vals[i] = op(index, val)
			i++
			return Continue, nil
		})
	test.Nil(err, t)
	test.Eq(len(cpy), len(vals), t)
	for i, v := range cpy {
		test.Eq(op(i, v), vals[i], t)
	}
}
func TestForEach(t *testing.T) {
	forEachIterHelper([]int{1, 2, 3, 4}, func(index int, val int) int {
		return val + 1
	}, t)
	forEachIterHelper([]int{1}, func(index int, val int) int {
		return val + 1
	}, t)
	forEachIterHelper([]int{}, func(index int, val int) int {
		return val + 1
	}, t)
	forEachIterHelper([]int{5, 5, 5, 5}, func(index int, val int) int {
		return index + 1
	}, t)
}

func TestForEachEarlyStopBool(t *testing.T) {
	cntr := 0
	err := SliceElems([]int{0, 1, 2, 3, 4}).ForEach(
		func(i int, v int) (IteratorFeedback, error) {
			cntr++
			if v == 3 {
				return Break, nil
			}
			return Continue, nil
		})
	test.Eq(4, cntr, t)
	test.Nil(err, t)
}

func TestForEachEarlyStopErr(t *testing.T) {
	cntr := 0
	err := SliceElems([]int{0, 1, 2, 3, 4}).ForEach(
		func(i int, v int) (IteratorFeedback, error) {
			cntr++
			if v == 3 {
				return Continue, errors.New("NEW ERROR")
			}
			return Continue, nil
		})
	test.Eq(4, cntr, t)
	if err == nil {
		fmt.Println(
			"ForEach did not return an error when it was supposed to.", t,
		)
	}
}
