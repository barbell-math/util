package containerTypes

import "errors"

var Empty = errors.New("The container is empty.")
var Full = errors.New("The container is full.")
var Duplicate = errors.New("The container already contains the supplied value.")
var KeyError = errors.New("The supplied key was invalid.")
var ValueError = errors.New("The supplied value was invalid.")
var UpdateViolation = errors.New("Attempting to update the supplied value would violate uniqueness constraints.")
