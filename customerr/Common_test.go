package customerr

import (
	"fmt"
	"testing"

	"github.com/barbell-math/util/test"
)

func TestAssert(t *testing.T) {
	test.Panics(
		func() {
			Assert(func() error {
				return ValOutsideRange
			})
		},
		t,
	)
	test.NoPanic(
		func() {
			Assert(func() error {
				return nil
			})
		},
		t,
	)
}

func TestWrap(t *testing.T) {
	e := Wrap(ValOutsideRange, "%d", 100)
	test.Eq(fmt.Sprintf("%s\n  |- 100", ValOutsideRange.Error()), e.Error(), t)
}

func TestWrapValueList(t *testing.T) {
	e:=WrapValueList(ValOutsideRange, "Val outside range", []WrapListVal{
		{ItemName: "value", Item: 10},
	})
	test.Eq(
		fmt.Sprintf(
			"%s\n  |- Description: Val outside range\n  |- value (int): 10",
			ValOutsideRange.Error(),
		),
		e.Error(),
		t,
	)
}

func TestUnwrap(t *testing.T) {
	e := Wrap(ValOutsideRange, "%d", 100)
	test.Eq(ValOutsideRange, Unwrap(e), t)
	test.ContainsError(ValOutsideRange, Unwrap(e), t)
}

func TestAppendErrorTwoErrors(t *testing.T) {
	e := AppendError(ValOutsideRange, DimensionsDoNotAgree)
	test.Eq(
		fmt.Sprintf(
			"%s\n%s",
			ValOutsideRange.Error(),
			DimensionsDoNotAgree.Error(),
		),
		e.Error(),
		t,
	)
}

func TestAppendErrorOnlyFirst(t *testing.T) {
	e := AppendError(ValOutsideRange, nil)
	test.Eq(ValOutsideRange.Error(), e.Error(), t)
	test.Eq(ValOutsideRange, e, t)
}

func TestAppendErrorOnlySecond(t *testing.T) {
	e := AppendError(nil, ValOutsideRange)
	test.Eq(ValOutsideRange.Error(), e.Error(), t)
	test.Eq(ValOutsideRange, e, t)
}

func TestAppendErrorBothNil(t *testing.T) {
	e := AppendError(nil, nil)
	test.Nil(e, t)
}
