package translators

import "errors"

var (
	FlagProvidedToManyTimesErr = errors.New("A flag was provided to many times.")
	ValNotInAllowedListErr     = errors.New("Value was not found in the allowed list")
	EnvVarNotSetErr            = errors.New("The supplied environment variable was not set")
)
