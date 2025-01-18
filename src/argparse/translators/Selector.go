package translators

import (
	"github.com/barbell-math/util/src/container/containers"
	"github.com/barbell-math/util/src/customerr"
	"github.com/barbell-math/util/src/widgets"
)

type (
	// A translator that imposes a set of specific values on a cmd line
	// argument.
	Selector[T any, U Translater[T], W widgets.BaseInterface[T]] struct {
		AllowedVals     containers.HashSet[T, W]
		ValueTranslator Translater[T]
	}
)

func (s Selector[T, U, W]) Translate(arg string) (T, error) {
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

func (s Selector[T, U, W]) Reset() {
	s.ValueTranslator.Reset()
}
