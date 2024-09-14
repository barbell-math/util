package common

import "errors"

var MissingRequiredArgs = errors.New("Not all required flags were passed.")
