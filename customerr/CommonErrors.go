package customerr

import "errors"

var ValOutsideRange = errors.New("The specified value is outside the allowed range.")
var DimensionsDoNotAgree = errors.New("Dimensions do not agree.")
var InvalidValue = errors.New("The supplied value is not valid in the supplied context.")
var IncorrectType = errors.New("An incorrect type was recieved.")                 // Used when only one type is accepted
var UnsupportedType = errors.New("The type of the recieved value was not valid.") // Used when multiple types are accepted
var MultipleErrorsOccurred = errors.New("Multiple errors have occurred.")
