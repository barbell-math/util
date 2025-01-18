package appactions

import (
	"errors"
	"testing"

	testenum "github.com/barbell-math/util/src/appActions/testEnum"
	"github.com/barbell-math/util/src/test"
)

type (
	testActions struct {
		errStart error
		errRun   error
		errStop  error

		started bool
		ran     bool
		stopped bool
	}
)

var testErr = errors.New("test err")
var testErr2 = errors.New("test err 2")

func (t *testActions) Setup() error {
	t.started = true
	return t.errStart
}
func (t *testActions) Run() error {
	t.ran = true
	return t.errRun
}
func (t *testActions) Teardown() error {
	t.stopped = true
	return t.errStop
}

func TestUnknwonAction(t *testing.T) {
	a := ActionRegistry{}
	err := a.PerformAction(testenum.AppActionOne)
	test.ContainsError(UnknownActionErr, err, t)
}

func TestKnownActionAllPassing(t *testing.T) {
	a := ActionRegistry{
		testenum.AppActionOne: &testActions{},
	}
	err := a.PerformAction(testenum.AppActionOne)
	test.Nil(err, t)
	test.True(a[testenum.AppActionOne].(*testActions).started, t)
	test.True(a[testenum.AppActionOne].(*testActions).ran, t)
	test.True(a[testenum.AppActionOne].(*testActions).stopped, t)
}

func TestKnownActionSetupError(t *testing.T) {
	a := ActionRegistry{
		testenum.AppActionOne: &testActions{
			errStart: testErr,
		},
	}
	err := a.PerformAction(testenum.AppActionOne)
	test.ContainsError(testErr, err, t)
	test.ContainsError(AppSetupErr, err, t)
	test.True(a[testenum.AppActionOne].(*testActions).started, t)
	test.False(a[testenum.AppActionOne].(*testActions).ran, t)
	test.True(a[testenum.AppActionOne].(*testActions).stopped, t)
}

func TestKnownActionRunError(t *testing.T) {
	a := ActionRegistry{
		testenum.AppActionOne: &testActions{
			errRun: testErr,
		},
	}
	err := a.PerformAction(testenum.AppActionOne)
	test.ContainsError(testErr, err, t)
	test.ContainsError(AppRunErr, err, t)
	test.True(a[testenum.AppActionOne].(*testActions).started, t)
	test.True(a[testenum.AppActionOne].(*testActions).ran, t)
	test.True(a[testenum.AppActionOne].(*testActions).stopped, t)
}

func TestKnownActionStopError(t *testing.T) {
	a := ActionRegistry{
		testenum.AppActionOne: &testActions{
			errStop: testErr,
		},
	}
	err := a.PerformAction(testenum.AppActionOne)
	test.ContainsError(testErr, err, t)
	test.ContainsError(AppTeardownErr, err, t)
	test.True(a[testenum.AppActionOne].(*testActions).started, t)
	test.True(a[testenum.AppActionOne].(*testActions).ran, t)
	test.True(a[testenum.AppActionOne].(*testActions).stopped, t)
}

func TestKnownActionSetupAndTeardownErr(t *testing.T) {
	a := ActionRegistry{
		testenum.AppActionOne: &testActions{
			errStart: testErr,
			errStop:  testErr2,
		},
	}
	err := a.PerformAction(testenum.AppActionOne)
	test.ContainsError(testErr, err, t)
	test.ContainsError(testErr2, err, t)
	test.ContainsError(AppSetupErr, err, t)
	test.ContainsError(AppTeardownErr, err, t)
	test.True(a[testenum.AppActionOne].(*testActions).started, t)
	test.False(a[testenum.AppActionOne].(*testActions).ran, t)
	test.True(a[testenum.AppActionOne].(*testActions).stopped, t)
}

func TestKnownActionRunAndTeardownErr(t *testing.T) {
	a := ActionRegistry{
		testenum.AppActionOne: &testActions{
			errRun:  testErr,
			errStop: testErr2,
		},
	}
	err := a.PerformAction(testenum.AppActionOne)
	test.ContainsError(testErr, err, t)
	test.ContainsError(testErr2, err, t)
	test.ContainsError(AppRunErr, err, t)
	test.ContainsError(AppTeardownErr, err, t)
	test.True(a[testenum.AppActionOne].(*testActions).started, t)
	test.True(a[testenum.AppActionOne].(*testActions).ran, t)
	test.True(a[testenum.AppActionOne].(*testActions).stopped, t)
}
