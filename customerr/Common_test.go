package err

import (
	"fmt"
	"testing"

	"github.com/barbell-math/util/test"
)

func TestAssert(t *testing.T){
    test.Panics(
        func() {
            Assert(func() error {
                return ValOutsideRange
            })
        },
        "Assert did not panic when it should have.",t,
    )
    test.NoPanic(
        func() {
            Assert(func() error {
                return nil
            })
        },
        "Assert paniced when it should not have.",t,
    )
}

func TestWrap(t *testing.T) {
    e:=Wrap(ValOutsideRange,"%d",100)
    test.BasicTest(fmt.Sprintf("%s\n  |- 100\n",ValOutsideRange.Error()),e.Error(),
        "Wrap did not properly format the error.",t,
    )
}

func TestUnwrap(t *testing.T){
    e:=Wrap(ValOutsideRange,"%d",100)
    test.BasicTest(ValOutsideRange,Unwrap(e),
        "Unwrap did not return the original value.",t,
    )
    if ValOutsideRange!=Unwrap(e) {
        test.FormatError(ValOutsideRange.Error(),e.Error(),
            "Bare equality test failed.",t,
        )
    }
}

func TestAppendErrorTwoErrors(t *testing.T) {
    e:=AppendError(ValOutsideRange,DimensionsDoNotAgree)
    test.BasicTest(fmt.Sprintf(
        "Multiple errors have occurred.\nFirst error: %s\nSecond error: %s\n",
        ValOutsideRange,DimensionsDoNotAgree,
    ), e.Error(), "Appending errors did not format the final error properly.",t,)
}

func TestAppendErrorOnlyFirst(t *testing.T) {
    e:=AppendError(ValOutsideRange,nil)
    test.BasicTest(ValOutsideRange.Error(),e.Error(),
        "Only passing the first error changed the error text.",t,
    )
    test.BasicTest(ValOutsideRange,e,
        "Only passing the first error did not return the original error.",t,
    )
}

func TestAppendErrorOnlySecond(t *testing.T) {
    e:=AppendError(nil,ValOutsideRange)
    test.BasicTest(ValOutsideRange.Error(),e.Error(),
        "Only passing the second error changed the error text.",t,
    )
    test.BasicTest(ValOutsideRange,e,
        "Only passing the second error did not return the original error.",t,
    )
}

func TestAppendErrorBothNil(t *testing.T){
    e:=AppendError(nil,nil)
    test.BasicTest(nil,e,
        "Passing two nil errors did not return a nil error.",t,
    )
}
