package containers

import (
	"testing"

	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/container/staticContainers"
	"github.com/barbell-math/util/iter"
	"github.com/barbell-math/util/test"
	"github.com/barbell-math/util/widgets"
)

func TestSlidingWindowEmpty(t *testing.T) {
	q, _ := NewCircularBuffer[int, widgets.BuiltinInt](2)
	cnt, err := SlidingWindow[int](iter.SliceElems([]int{}), &q, true).Count()
	test.Eq(0, cnt, t)
	test.Nil(err, t)
}

func TestSlidingWindowNoPartialsNoSlidingWindowValues(t *testing.T) {
	cntr := 0
	q, _ := NewCircularBuffer[int, widgets.BuiltinInt](101)
	vals := make([]int, 100)
	for i := 0; i < len(vals); i++ {
		vals[i] = i
	}
	cntr, err := SlidingWindow[int](iter.SliceElems(vals), &q, false).Count()
	test.Eq(0, cntr, t)
	test.Nil(err, t)
}

func TestSlidingWindowNoPartials(t *testing.T) {
	cntr := 0
	q, _ := NewCircularBuffer[int, widgets.BuiltinInt](2)
	vals := make([]int, 100)
	for i := 0; i < len(vals); i++ {
		vals[i] = i
	}
	err := SlidingWindow[int](iter.SliceElems(vals), &q, false).ForEach(
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
		},
	)
	test.Eq(99, cntr, t)
	test.Nil(err, t)
}

func TestSlidingWindowPartials(t *testing.T) {
	cntr := 0
	q, _ := NewCircularBuffer[int, widgets.BuiltinInt](2)
	vals := make([]int, 100)
	for i := 0; i < len(vals); i++ {
		vals[i] = i
	}
	err := SlidingWindow[int](iter.SliceElems(vals), &q, true).ForEach(
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
		},
	)
	test.Eq(100, cntr, t)
	test.Nil(err, t)
}

func TestSteppingWindowEmpty(t *testing.T) {
	q, _ := NewCircularBuffer[int, widgets.BuiltinInt](2)
	cnt, err := SteppingWindow[int](iter.SliceElems([]int{}), &q).Count()
	test.Eq(0, cnt, t)
	test.Nil(err, t)
}

func TestSteppingWindowNoStepValues(t *testing.T) {
	cntr := 0
	q, _ := NewCircularBuffer[int, widgets.BuiltinInt](101)
	vals := make([]int, 100)
	for i := 0; i < len(vals); i++ {
		vals[i] = i
	}
	cntr, err := SteppingWindow[int](iter.SliceElems(vals), &q).Count()
	test.Eq(0, cntr, t)
	test.Nil(err, t)
}

func TestSteppingWindow(t *testing.T) {
	cntr := 0
	q, _ := NewCircularBuffer[int, widgets.BuiltinInt](2)
	vals := make([]int, 100)
	for i := 0; i < len(vals); i++ {
		vals[i] = i
	}
	err := SteppingWindow[int](iter.SliceElems(vals), &q).ForEach(
		func(index int, val staticContainers.Vector[int]) (iter.IteratorFeedback, error) {
			test.Eq(2, q.Length(), t)
			if v, err := q.PeekFront(); err == nil {
				test.Eq(cntr*2, v, t)
			}
			if v, err := q.Get(1); err == nil {
				test.Eq(cntr*2+1, v, t)
			}
			cntr++
			return iter.Continue, nil
		},
	)
	test.Eq(50, cntr, t)
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
