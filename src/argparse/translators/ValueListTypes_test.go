package translators

import (
	"testing"

	"github.com/barbell-math/util/src/container/containers"
	"github.com/barbell-math/util/src/customerr"
	"github.com/barbell-math/util/src/test"
	"github.com/barbell-math/util/src/widgets"
)

func TestListValuesNilAllowedVals(t *testing.T) {
	l := ListValues[BuiltinInt, widgets.BuiltinInt, int]{
		ValueTranslator: BuiltinInt{Base: 10},
	}

	res, err := l.Translate("1")
	test.Nil(err, t)
	test.SlicesMatch[int](res, []int{1}, t)
}

func TestListValuesAllowedValsPassing(t *testing.T) {
	l := ListValues[BuiltinInt, widgets.BuiltinInt, int]{
		ValueTranslator: BuiltinInt{Base: 10},
		AllowedVals:     containers.HashSetValInit[int, widgets.BuiltinInt](1),
	}

	res, err := l.Translate("1")
	test.Nil(err, t)
	test.SlicesMatch[int](res, []int{1}, t)
}

func TestListValuesAllowedValsFailing(t *testing.T) {
	l := ListValues[BuiltinInt, widgets.BuiltinInt, int]{
		ValueTranslator: BuiltinInt{Base: 10},
		AllowedVals:     containers.HashSetValInit[int, widgets.BuiltinInt](1, 2, 3),
	}

	_, err := l.Translate("4")
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(ValNotInAllowedListErr, err, t)
}
