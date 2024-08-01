package common

import "errors"

var NotEnoughArgs = errors.New("Not enough arguments.")
var MissingRequiredArgs = errors.New("Not all required flags were passed.")
