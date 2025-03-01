package translators

import (
	"github.com/barbell-math/util/src/customerr"
	"github.com/barbell-math/util/src/math/basic"
)

//go:generate ../../../bin/ifaceImplCheck -typeToCheck=Range

type (
	// A translator that imposes a range on the supplied cmd line argument.
	//gen:ifaceImplCheck generics [int, BuiltinInt]
	//gen:ifaceImplCheck ifaceName Translator[int]
	//gen:ifaceImplCheck valOrPntr both
	Range[T basic.Number, U Translator[T]] struct {
		// Inclusive min
		Min T
		// Exclusive max
		Max           T
		NumTranslator U
	}
)

func (r Range[T, U]) Translate(arg string) (T, error) {
	rv, err := r.NumTranslator.Translate(arg)
	if err != nil {
		return rv, err
	}
	if rv < r.Min || rv >= r.Max {
		return rv, customerr.WrapValueList(
			customerr.ValOutsideRange,
			"",
			[]customerr.WrapListVal{
				{ItemName: "Min (inclusive)", Item: r.Min},
				{ItemName: "Max (exclusive)", Item: r.Max},
			},
		)
	}
	return rv, nil
}

func (r Range[T, U]) Reset() {
	r.NumTranslator.Reset()
}
