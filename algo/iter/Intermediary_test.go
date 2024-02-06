package iter

import (
	"errors"
	"fmt"
	"testing"

	"github.com/barbell-math/util/customerr"
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
	test.BasicTest(1, next, "Next did not return the correct value.", t)
	test.BasicTest(nil, err,
		"Next returned an error when it was not supposed to.", t,
	)
	test.BasicTest(true, cont, "Next did not reutrn correct cont status.", t)
	next, err, cont = n(Iterate)
	test.BasicTest(3, next, "Next did not return the correct value.", t)
	test.BasicTest(nil, err,
		"Next returned an error when it was not supposed to.", t,
	)
	test.BasicTest(true, cont, "Next did not reutrn correct cont status.", t)
	next, err, cont = n(Iterate)
	test.BasicTest(5, next, "Next did not return the correct value.", t)
	test.BasicTest(nil, err,
		"Next returned an error when it was not supposed to.", t,
	)
	test.BasicTest(true, cont, "Next did not reutrn correct cont status.", t)
	next, err, cont = n(Iterate)
	test.BasicTest(5, next, "Next did not return the correct value.", t)
	if err == nil {
		test.FormatError("!nil", err,
			"Next did not return an error when it was supposed to.", t,
		)
	}
	test.BasicTest(false, cont, "Next did not reutrn correct cont status.", t)
	next, err, cont = n(Break)
	test.BasicTest(0, next, "Next did not return the correct value.", t)
	test.BasicTest(nil, err,
		"Next returned an error when it was not supposed to.", t,
	)
	test.BasicTest(false, cont, "Next did not reutrn correct cont status.", t)
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
	test.BasicTest(true, breakReached,
		"Next did not properly call parrent iterators with break flag.", t,
	)
	test.BasicTest(true, breakReached2,
		"Next did not properly call parrent iterators with break flag.", t,
	)
	test.BasicTest(nil, err,
		"Next returned an error when it shouldn't have.", t,
	)
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
	test.BasicTest(true, breakReached,
		"Next did not properly call parrent iterators with break flag.", t,
	)
	test.BasicTest(true, breakReached2,
		"Next did not properly call parrent iterators with break flag.", t,
	)
	test.BasicTest(expectedErr, err,
		"Next returned an error when it shouldn't have.", t,
	)
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
	test.BasicTest(true, breakReached,
		"Next did not properly call parrent iterators with break flag.", t,
	)
	test.BasicTest(true, breakReached2,
		"Next did not properly call parrent iterators with break flag.", t,
	)
	test.BasicTest(expectedErr, err,
		"Next returned an error when it shouldn't have.", t,
	)
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
	test.BasicTest(true, breakReached,
		"Next did not properly call parrent iterators with break flag.", t,
	)
	test.BasicTest(true, breakReached2,
		"Next did not properly call parrent iterators with break flag.", t,
	)
	tmp := fmt.Sprintf("%s", customerr.AppendError(expectedErr, expectedErr))
	if fmt.Sprintf("%s", expectedErr) == tmp {
		test.FormatError(tmp, err,
			"Next returned an error when it shouldn't have.", t,
		)
	}
}

func TestSetupTeardownNoElems(t *testing.T) {
	setupRan := false
	teardownRan := false
	setup := func() error { setupRan = true; return nil }
	teardown := func() error { teardownRan = true; return nil }
	cnt, err := NoElem[int]().SetupTeardown(setup, teardown).Count()
	test.BasicTest(true, setupRan,
		"Setup was not run when it was supposed to be.", t,
	)
	test.BasicTest(true, teardownRan,
		"Teardown was not run when it was supposed to be.", t,
	)
	test.BasicTest(0, cnt,
		"SetupTeardown iterated the wrong number of times.", t,
	)
	test.BasicTest(nil, err,
		"SetupTeardown returned an error when it should not have.", t,
	)
}

func TestSetupTeardownWithElems(t *testing.T) {
	setupRan := false
	teardownRan := false
	setup := func() error { setupRan = true; return nil }
	teardown := func() error { teardownRan = true; return nil }
	cnt, err := SliceElems[int]([]int{0, 1, 2, 3}).SetupTeardown(setup, teardown).Count()
	test.BasicTest(true, setupRan,
		"Setup was not run when it was supposed to be.", t,
	)
	test.BasicTest(true, teardownRan,
		"Teardown was not run when it was supposed to be.", t,
	)
	test.BasicTest(4, cnt,
		"SetupTeardown iterated the wrong number of times.", t,
	)
	test.BasicTest(nil, err,
		"SetupTeardown returned an error when it should not have.", t,
	)
}

func TestSetupTeardownWithSetupError(t *testing.T) {
	setupRan := false
	teardownRan := false
	setup := func() error { setupRan = true; return errors.New("ERROR") }
	teardown := func() error { teardownRan = true; return nil }
	cnt, err := SliceElems[int]([]int{0, 1, 2, 3}).SetupTeardown(setup, teardown).Count()
	test.BasicTest(true, setupRan,
		"Setup was not run when it was supposed to be.", t,
	)
	test.BasicTest(true, teardownRan,
		"Teardown was not run when it was supposed to be.", t,
	)
	test.BasicTest(0, cnt,
		"SetupTeardown iterated the wrong number of times.", t,
	)
	if err.Error() != "ERROR" {
		test.FormatError(errors.New("ERROR"), err,
			"Setupteardown did not return the correct error.", t,
		)
	}
}

func TestSetupTeardownWithTeardownError(t *testing.T) {
	setupRan := false
	teardownRan := false
	setup := func() error { setupRan = true; return nil }
	teardown := func() error { teardownRan = true; return errors.New("ERROR") }
	err := SliceElems[int]([]int{0, 1, 2, 3}).SetupTeardown(setup, teardown).Consume()
	test.BasicTest(true, setupRan,
		"Setup was not run when it was supposed to be.", t,
	)
	test.BasicTest(true, teardownRan,
		"Teardown was not run when it was supposed to be.", t,
	)
	if err.Error() != "ERROR" {
		test.FormatError(errors.New("ERROR"), err,
			"Setupteardown did not return the correct error.", t,
		)
	}
}

func injectIterHelper[T any](
	initialVals []T,
	desiredSeq []T,
	op func(idx int, val T, injectedPrev bool) (T, error, bool),
	t *testing.T,
) {
	result, err := SliceElems(initialVals).Inject(op).Collect()
	test.BasicTest(nil, err, "Inject created an error when it should not have.", t)
	test.BasicTest(len(desiredSeq), len(result),
		"Using inject did not return the correct number of elements.", t,
	)
	for i, v := range desiredSeq {
		if i < len(result) {
			test.BasicTest(v, result[i],
				"Inject incorrectly modified values.", t,
			)
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
