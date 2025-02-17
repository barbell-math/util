package translators

import (
	"github.com/barbell-math/util/src/enum"
)

type (
	// A translator that forces a value to be a valid enum value for the given
	// enum type.
	Enum[EP enum.Pntr[E], E enum.Value] struct{}
)

func (_ Enum[EP, E]) Translate(arg string) (E, error) {
	var rv E
	var ei EP
	ei = &rv
	return rv, ei.FromString(arg)
}

func (_ Enum[E, EP]) Reset() {
	// intentional noop - Enum has no state to reset
}
