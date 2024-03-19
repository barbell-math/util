package containers

import (
	"testing"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/algo/widgets"
	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/container/staticContainers"
	"github.com/barbell-math/util/test"
)

func TestWindowEmpty(t *testing.T) {
	q, _ := NewCircularBuffer[int, widgets.BuiltinInt](2)
	cnt, err := Window[int](iter.SliceElems([]int{}), &q, true).Count()
	test.Eq(0, cnt, t)
	test.Nil(err, t)
}

func TestWindowNoPartialsNoWindowValues(t *testing.T) {
	cntr := 0
	q, _ := NewCircularBuffer[int, widgets.BuiltinInt](101)
	vals := make([]int, 100)
	for i := 0; i < len(vals); i++ {
		vals[i] = i
	}
	cntr, err := Window[int](iter.SliceElems(vals), &q, false).Count()
	test.Eq(0, cntr, t)
	test.Nil(err, t)
}

func TestWindowNoPartials(t *testing.T) {
	cntr := 0
	q, _ := NewCircularBuffer[int, widgets.BuiltinInt](2)
	vals := make([]int, 100)
	for i := 0; i < len(vals); i++ {
		vals[i] = i
	}
	err := Window[int](iter.SliceElems(vals), &q, false).ForEach(
		func(index int, val staticContainers.Vector[int]) (iter.IteratorFeedback, error) {
			cntr++
			test.Eq(2, q.Length(), t)
			if v, err := q.PeekFront(); err == nil {
				test.Eq(index, v, t)
			}
			if v, err := q.Get(1); err == nil {
				test.Eq(index+1, v, t)
			}
			return iter.Continue, nil
		})
	test.Eq(99, cntr, t)
	test.Nil(err, t)
}

func TestWindowPartials(t *testing.T) {
	cntr := 0
	q, _ := NewCircularBuffer[int, widgets.BuiltinInt](2)
	vals := make([]int, 100)
	for i := 0; i < len(vals); i++ {
		vals[i] = i
	}
	err := Window[int](iter.SliceElems(vals), &q, true).ForEach(
		func(index int, val staticContainers.Vector[int]) (iter.IteratorFeedback, error) {
			cntr++
			if index == 0 || index == 100 {
				test.Eq(1, q.Length(), t)
				if v, err := q.PeekFront(); err == nil {
					test.Eq(index, v, t)
				}
			} else {
				test.Eq(2, q.Length(), t)
				if v, err := q.PeekFront(); err == nil {
					test.Eq(index-1, v, t)
				}
				if v, err := q.Get(1); err == nil {
					test.Eq(index, v, t)
				}
			}
			return iter.Continue, nil
		})
	test.Eq(100, cntr, t)
	test.Nil(err, t)
}

func TestUniqueEmpty(t *testing.T) {
	s, _ := NewHashSet[int, widgets.BuiltinInt](2)
	Unique[int](iter.SliceElems([]int{}), &s, true)
	test.Eq(0, s.Length(), t)
}

func TestUniqueNoErrorOnDup(t *testing.T) {
	s, _ := NewHashSet[int, widgets.BuiltinInt](2)
	test.Nil(Unique[int](iter.SliceElems([]int{1, 2, 3, 4, 5, 6}), &s, false), t)
	test.Eq(6, s.Length(), t)
	for i := 1; i < 7; i++ {
		test.True(s.Contains(i), t)
	}

	s.Clear()
	test.Nil(
		Unique[int](iter.SliceElems([]int{1, 1, 2, 2, 3, 4, 5, 6, 3, 4, 5, 6, 7}), &s, false),
		t,
	)
	test.Eq(7, s.Length(), t)
	for i := 1; i < 8; i++ {
		test.True(s.Contains(i), t)
	}
}

func TestUniqueErrorOnDup(t *testing.T) {
	s, _ := NewHashSet[int, widgets.BuiltinInt](2)
	test.Nil(Unique[int](iter.SliceElems([]int{1, 2, 3, 4, 5, 6}), &s, true), t)
	test.Eq(6, s.Length(), t)
	for i := 1; i < 7; i++ {
		test.True(s.Contains(i), t)
	}

	s.Clear()
	test.ContainsError(
		containerTypes.Duplicate,
		Unique[int](iter.SliceElems([]int{1, 1, 2, 2, 3, 4, 5, 6, 3, 4, 5, 6, 7}), &s, true),
		t,
	)
	test.Eq(1, s.Length(), t)
	test.True(s.Contains(1), t)
}
