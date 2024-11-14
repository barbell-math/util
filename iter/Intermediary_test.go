package iter

import (
	"errors"
	"testing"

	"github.com/barbell-math/util/test"
)

func TestNext(t *testing.T) {
	//Should produce sequence:
	//  (1,nil,true)
	//  (3,nil,true)
	//  (5,nil,true)
	//  (5,err,false)
	n := SliceElems([]int{1, 2, 3, 4, 5, 6, 7}).Next(
		func(index, val int, status IteratorFeedback) (IteratorFeedback, int, error) {
			if status == Break {
				return Break, 0, nil
			}
			if val%2 == 0 {
				return Continue, val + 1, nil
			} else if val == 3 {
				return Iterate, val, nil
			} else if val == 5 {
				return Continue, val, errors.New("NEW ERROR")
			}
			return Continue, val, nil
		})
	next, err, cont := n(Iterate)
	test.Eq(1, next, t)
	test.Nil(err, t)
	test.True(cont, t)
	next, err, cont = n(Iterate)
	test.Eq(3, next, t)
	test.Nil(err, t)
	test.True(cont, t)
	next, err, cont = n(Iterate)
	test.Eq(5, next, t)
	test.Nil(err, t)
	test.True(cont, t)
	next, err, cont = n(Iterate)
	test.Eq(5, next, t)
	test.NotNil(err, t)
	test.False(cont, t)
	next, err, cont = n(Break)
	test.Eq(0, next, t)
	test.Nil(err, t)
	test.False(cont, t)
}

func TestNextReachesBreak(t *testing.T) {
	breakReached := false
	breakReached2 := false
	n := SliceElems([]int{1, 2, 3, 4, 5, 6, 7}).Next(
		func(index, val int, status IteratorFeedback) (IteratorFeedback, int, error) {
			if status == Break {
				breakReached = true
			}
			return Continue, val, nil
		}).Next(func(index, val int, status IteratorFeedback) (IteratorFeedback, int, error) {
		if status == Break {
			breakReached2 = true
		}
		return Continue, val, nil
	})
	err := n.Consume()
	test.True(breakReached, t)
	test.True(breakReached2, t)
	test.Nil(err, t)
}

func TestNextReachesBreakParentErr(t *testing.T) {
	expectedErr := errors.New("NEW ERROR")
	breakReached := false
	breakReached2 := false
	n := SliceElems([]int{1, 2, 3, 4, 5, 6, 7}).Next(
		func(index, val int, status IteratorFeedback) (IteratorFeedback, int, error) {
			if status == Break {
				breakReached = true
			}
			if val == 5 {
				return Break, val, expectedErr
			}
			return Continue, val, nil
		}).Next(func(index, val int, status IteratorFeedback) (IteratorFeedback, int, error) {
		if status == Break {
			breakReached2 = true
		}
		return Continue, val, nil
	})
	err := n.Consume()
	test.True(breakReached, t)
	test.True(breakReached2, t)
	test.Eq(expectedErr, err, t)
}

func TestNextReachesBreakParentCleanUpErr(t *testing.T) {
	expectedErr := errors.New("NEW ERROR")
	breakReached := false
	breakReached2 := false
	n := SliceElems([]int{1, 2, 3, 4, 5, 6, 7}).Next(
		func(index, val int, status IteratorFeedback) (IteratorFeedback, int, error) {
			if status == Break {
				breakReached = true
				return Continue, val, expectedErr
			}
			return Continue, val, nil
		}).Next(func(index, val int, status IteratorFeedback) (IteratorFeedback, int, error) {
		if status == Break {
			breakReached2 = true
		}
		return Continue, val, nil
	})
	err := n.Consume()
	test.True(breakReached, t)
	test.True(breakReached2, t)
	test.Eq(expectedErr, err, t)
}

func TestNextReachesBreakParentErrAndCleanUpErr(t *testing.T) {
	expectedErr := errors.New("NEW ERROR")
	breakReached := false
	breakReached2 := false
	n := SliceElems([]int{1, 2, 3, 4, 5, 6, 7}).Next(
		func(index, val int, status IteratorFeedback) (IteratorFeedback, int, error) {
			if status == Break {
				breakReached = true
				return Continue, val, expectedErr
			}
			return Continue, val, nil
		}).Next(func(index, val int, status IteratorFeedback) (IteratorFeedback, int, error) {
		if status == Break {
			breakReached2 = true
		}
		if val == 5 {
			return Break, val, expectedErr
		}
		return Continue, val, nil
	})
	err := n.Consume()
	test.True(breakReached, t)
	test.True(breakReached2, t)
	test.ContainsError(expectedErr, err, t)
}

func TestSetupTeardownNoElems(t *testing.T) {
	setupRan := false
	teardownRan := false
	setup := func() error { setupRan = true; return nil }
	teardown := func() error { teardownRan = true; return nil }
	cnt, err := NoElem[int]().SetupTeardown(setup, teardown).Count()
	test.True(setupRan, t)
	test.True(teardownRan, t)
	test.Eq(0, cnt, t)
	test.Nil(err, t)
}

func TestSetupTeardownWithElems(t *testing.T) {
	setupRan := false
	teardownRan := false
	setup := func() error { setupRan = true; return nil }
	teardown := func() error { teardownRan = true; return nil }
	cnt, err := SliceElems[int]([]int{0, 1, 2, 3}).SetupTeardown(setup, teardown).Count()
	test.True(setupRan, t)
	test.True(teardownRan, t)
	test.Eq(4, cnt, t)
	test.Nil(err, t)
}

func TestSetupTeardownWithSetupError(t *testing.T) {
	setupRan := false
	teardownRan := false
	expectedError := errors.New("ERROR")
	setup := func() error { setupRan = true; return expectedError }
	teardown := func() error { teardownRan = true; return nil }
	cnt, err := SliceElems[int]([]int{0, 1, 2, 3}).SetupTeardown(setup, teardown).Count()
	test.True(setupRan, t)
	test.True(teardownRan, t)
	test.Eq(0, cnt, t)
	test.ContainsError(expectedError, err, t)
}

func TestSetupTeardownWithTeardownError(t *testing.T) {
	setupRan := false
	teardownRan := false
	expectedError := errors.New("ERROR")
	setup := func() error { setupRan = true; return nil }
	teardown := func() error { teardownRan = true; return expectedError }
	err := SliceElems[int]([]int{0, 1, 2, 3}).SetupTeardown(setup, teardown).Consume()
	test.True(setupRan, t)
	test.True(teardownRan, t)
	test.ContainsError(expectedError, err, t)
}

func injectIterHelper[T any](
	initialVals []T,
	desiredSeq []T,
	op func(idx int, val T, injectedPrev bool) (T, error, bool),
	t *testing.T,
) {
	result, err := SliceElems(initialVals).Inject(op).Collect()
	test.Nil(err, t)
	test.Eq(len(desiredSeq), len(result), t)
	for i, v := range desiredSeq {
		if i < len(result) {
			test.Eq(v, result[i], t)
		}
	}
}
func TestInjectSingleValue(t *testing.T) {
	injectIterHelper([]int{}, []int{},
		func(idx, val int, injectedPrev bool) (int, error, bool) { return 0, nil, idx == 1 }, t,
	)
	injectIterHelper([]int{1, 2, 3, 4}, []int{1, 2, 3, 4},
		func(idx, val int, injectedPrev bool) (int, error, bool) { return 0, nil, idx == 5 }, t,
	)
	injectIterHelper([]int{}, []int{0},
		func(idx, val int, injectedPrev bool) (int, error, bool) { return 0, nil, idx == 0 }, t,
	)
	injectIterHelper([]int{1}, []int{0, 1},
		func(idx, val int, injectedPrev bool) (int, error, bool) { return 0, nil, idx == 0 }, t,
	)
	injectIterHelper([]int{1, 2, 3, 4}, []int{0, 1, 2, 3, 4},
		func(idx, val int, injectedPrev bool) (int, error, bool) { return 0, nil, idx == 0 }, t,
	)
	injectIterHelper([]int{1}, []int{1, 0},
		func(idx, val int, injectedPrev bool) (int, error, bool) { return 0, nil, idx == 1 }, t,
	)
	injectIterHelper([]int{1, 2, 3, 4}, []int{1, 2, 3, 4, 0},
		func(idx, val int, injectedPrev bool) (int, error, bool) { return 0, nil, idx == 4 }, t,
	)
	injectIterHelper([]int{1, 2, 3, 4}, []int{1, 2, 0, 3, 4},
		func(idx, val int, injectedPrev bool) (int, error, bool) { return 0, nil, idx == 2 }, t,
	)
}

func multiValueInjectHelper[T any](
	initialVals []T,
	injectables []T,
	startIdx int,
	exp []T,
	t *testing.T,
) {
	op := func(idx int, val T, injectedPrev bool) (T, error, bool) {
		if idx >= startIdx && len(injectables) > 0 {
			v := injectables[0]
			injectables = injectables[1:]
			return v, nil, true
		}
		var tmp T
		return tmp, nil, false
	}
	injectIterHelper[T](initialVals, exp, op, t)
}
func TestInjectMultiValue(t *testing.T) {
	multiValueInjectHelper[int]([]int{}, []int{}, 0, []int{}, t)
	multiValueInjectHelper[int]([]int{}, []int{1, 2}, 0, []int{1, 2}, t)
	multiValueInjectHelper[int]([]int{1}, []int{2, 3, 4}, 2, []int{1}, t)
	multiValueInjectHelper[int]([]int{3}, []int{1, 2}, 0, []int{1, 2, 3}, t)
	multiValueInjectHelper[int]([]int{4}, []int{1, 2, 3}, 0, []int{1, 2, 3, 4}, t)
	multiValueInjectHelper[int]([]int{1}, []int{2, 3, 4}, 1, []int{1, 2, 3, 4}, t)
	multiValueInjectHelper[int]([]int{1, 2}, []int{3, 4}, 2, []int{1, 2, 3, 4}, t)
	multiValueInjectHelper[int]([]int{1, 2, 3}, []int{4}, 3, []int{1, 2, 3, 4}, t)
	multiValueInjectHelper[int]([]int{1, 2, 3, 4}, []int{}, 4, []int{1, 2, 3, 4}, t)
	multiValueInjectHelper[int]([]int{1, 2, 3, 4}, []int{5}, 4, []int{1, 2, 3, 4, 5}, t)
	multiValueInjectHelper[int]([]int{1, 2, 3, 4}, []int{5, 6}, 4, []int{1, 2, 3, 4, 5, 6}, t)
}
