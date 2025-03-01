package translators

import (
	"github.com/barbell-math/util/src/container/containers"
	"github.com/barbell-math/util/src/customerr"
	"github.com/barbell-math/util/src/widgets"
)

//go:generate ../../../bin/ifaceImplCheck -typeToCheck=Selector

type (
	// A translator that imposes a set of specific values on a cmd line
	// argument.
	//gen:ifaceImplCheck generics [BuiltinBool, widgets.BuiltinBool, bool]
	//gen:ifaceImplCheck ifaceName Translator[bool]
	//gen:ifaceImplCheck imports github.com/barbell-math/util/src/widgets
	//gen:ifaceImplCheck valOrPntr both
	Selector[T Translator[U], W widgets.BaseInterface[U], U any] struct {
		AllowedVals     containers.HashSet[U, W]
		ValueTranslator Translator[U]
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
