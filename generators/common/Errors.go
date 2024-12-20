package common

import "errors"

var MissingRequiredArgs = errors.New("Not all required flags were passed.")
var InvalidGeneratorFileName = errors.New("Invalid generator file name.")
var UnrecognizedArgs = errors.New("Unrecognized arguments were passed.")
var CommentArgsMalformed = errors.New("Malformed comment arg.")
