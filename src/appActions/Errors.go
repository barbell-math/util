package appactions

import "errors"

var UnknownActionErr = errors.New(
	"The supplied enum value has no associated action",
)
var AppSetupErr = errors.New("An error occurred setting up the application")
var AppRunErr = errors.New("An error occurred running the application")
var AppTeardownErr = errors.New("An error occurred tearing down the application")
