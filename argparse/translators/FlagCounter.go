package translators

import (
	"github.com/barbell-math/util/math/basic"
)

type (
	// Used to represent a flag that can be supplied many times, with a counter
	// incrementing each time the flag is encountered.
	FlagCntr[T basic.Int | basic.Uint] struct {
		val T
	}
)

func (f *FlagCntr[T]) Translate(arg string) (T, error) {
	f.val++
	return f.val, nil
}

func (f *FlagCntr[T]) Reset() {
	f.val = 0
}