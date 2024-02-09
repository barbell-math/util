package containerTypes

import "errors"

var Empty=errors.New("The container is empty.")
var Full=errors.New("The container is full.")
var Duplicate=errors.New("The container already contains the supplied value.")
var KeyError=errors.New("The supplied key was invalid.")
