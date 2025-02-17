package translators

import (
	"github.com/barbell-math/util/src/container/containers"
	"github.com/barbell-math/util/src/customerr"
	"github.com/barbell-math/util/src/widgets"
)

type (
	// A translator that imposes a set of specific values on a cmd line
	// argument.
	Selector[T Translater[U], W widgets.BaseInterface[U], U any] struct {
		AllowedVals     containers.HashSet[U, W]
		ValueTranslator Translater[U]
	}
)

func (s Selector[T, W, U]) Translate(arg string) (U, error) {
	rv, err := s.ValueTranslator.Translate(arg)
	if err != nil {
		return rv, err
	}
	if !s.AllowedVals.ContainsPntr(&rv) {
		return rv, customerr.AppendError(
			customerr.InvalidValue,
			customerr.WrapValueList(
				ValNotInAllowedListErr,
				"The supplied value must be found in the list shown below",
				[]customerr.WrapListVal{
					{"Supplied value", rv},
					{"Allowed values", &s.AllowedVals},
				},
			),
		)
	}
	return rv, nil
}

func (s Selector[T, W, U]) Reset() {
	s.ValueTranslator.Reset()
}
