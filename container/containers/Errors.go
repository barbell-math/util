package containers

import "errors"

var ValOutsideRange error=errors.New(
	"The specified value is outside the allowed range.",
)
