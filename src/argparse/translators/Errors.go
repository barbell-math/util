package translators

import "errors"

var (
	FlagProvidedToManyTimesErr = errors.New("A flag was provided to many times.")
	ValNotInAllowedListErr     = errors.New("Value was not found in the allowed list")
	EnvVarNotSetErr            = errors.New("The supplied environment variable was not set")
	CouldNotParseBigIntErr     = errors.New("Could not parse the supplied string as a big int")
	CouldNotParseBigRatErr     = errors.New("Could not parse the supplied string as a big rat")
	CouldNotParseBigFloatErr   = errors.New("Could not parse the supplied string as a big float")
	BitsTranslationErr         = errors.New("Could not translate bits value")
)
