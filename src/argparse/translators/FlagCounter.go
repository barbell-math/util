package translators

import (
	"github.com/barbell-math/util/src/customerr"
	"github.com/barbell-math/util/src/math/basic"
)

//go:generate ../../../bin/ifaceImplCheck -typeToCheck=FlagCntr
//go:generate ../../../bin/ifaceImplCheck -typeToCheck=LimitedFlagCntr

type (
	// Used to represent a flag that can be supplied many times, with a counter
	// incrementing each time the flag is encountered.
	//gen:ifaceImplCheck generics [int]
	//gen:ifaceImplCheck ifaceName Translator[int]
	//gen:ifaceImplCheck valOrPntr pntr
	FlagCntr[T basic.Int | basic.Uint] struct {
		cntr T
	}

	// Used to represent a flag that can be supplied many times up to the
	// provided maximum number of times. A counter will be incremented each
	// time the flag is encountered.
	//gen:ifaceImplCheck generics [int]
	//gen:ifaceImplCheck ifaceName Translator[int]
	//gen:ifaceImplCheck valOrPntr pntr
	LimitedFlagCntr[T basic.Int | basic.Uint] struct {
		FlagCntr[T]
		MaxTimes T
	}
)

func (f *FlagCntr[T]) Translate(arg string) (T, error) {
	f.cntr++
	return f.cntr, nil
}
func (f *LimitedFlagCntr[T]) Translate(arg string) (T, error) {
	if f.FlagCntr.cntr >= f.MaxTimes {
		return f.FlagCntr.cntr, customerr.Wrap(
			FlagProvidedToManyTimesErr,
			"Maximum allowed: %d", f.MaxTimes,
		)
	}
	rv, err := f.FlagCntr.Translate(arg)
	if err != nil {
		return rv, nil
	}
	return rv, nil
}

func (f *FlagCntr[T]) Reset() {
	f.cntr = 0
}
func (f *LimitedFlagCntr[T]) Reset() {
	f.FlagCntr.cntr = 0
}
