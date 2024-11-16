package translators

import (
	"github.com/barbell-math/util/src/container/containers"
	"github.com/barbell-math/util/src/customerr"
	"github.com/barbell-math/util/src/widgets"
)

type (
	// A translator that collects all supplied values into a slice.
	ListValues[T any, U Translater[T], W widgets.BaseInterface[T]] struct {
		vals            []T
		ValueTranslator U
		AllowedVals     containers.HashSet[T, W]
	}
)

func (l *ListValues[T, U, W]) Translate(arg string) ([]T, error) {
	v, err := l.ValueTranslator.Translate(arg)
	if err != nil {
		return l.vals, err
	}
	if l.AllowedVals.Length() > 0 && !l.AllowedVals.ContainsPntr(&v) {
		return l.vals, customerr.AppendError(
			customerr.InvalidValue,
			customerr.WrapValueList(
				ValNotInAllowedListErr,
				"The supplied value must be found in the list shown below",
				[]customerr.WrapListVal{
					{"Supplied value", v},
					{"Allowed values", &l.AllowedVals},
				},
			),
		)
	}
	l.vals = append(l.vals, v)
	return l.vals, nil
}

func (l *ListValues[T, U, W]) Reset() {
	l.ValueTranslator.Reset()
	l.vals = []T{}
}
