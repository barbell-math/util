package appactions

import (
	"github.com/barbell-math/util/src/customerr"
	"github.com/barbell-math/util/src/enum"
)

type (
	Action interface {
		Setup() error
		Run() error
		Teardown() error
	}

	ActionRegistry map[enum.Value]Action
)

func (a ActionRegistry) PerformAction(e enum.Value) error {
	action, ok := a[e]
	if !ok {
		return customerr.Wrap(UnknownActionErr, "Enum Value: %s", e)
	}

	var appErr error
	if err := action.Setup(); err != nil {
		appErr = customerr.AppendError(AppSetupErr, err)
		goto teardown
	}
	if err := action.Run(); err != nil {
		appErr = customerr.AppendError(AppRunErr, err)
		goto teardown
	}

teardown:
	if err := action.Teardown(); err != nil {
		return customerr.AppendError(appErr, AppTeardownErr, err)
	}
	return appErr
}
