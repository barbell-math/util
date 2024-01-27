package containers

import "errors"

var Empty=errors.New("The container is empty.")
var Full=errors.New("The container is full.")
