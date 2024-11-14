package translators

import (
	"testing"

	"github.com/barbell-math/util/container/containers"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/test"
	"github.com/barbell-math/util/widgets"
)

func TestSelectorPassing(t *testing.T) {
	l := Selector[int, BuiltinInt, widgets.BuiltinInt]{
		ValueTranslator: BuiltinInt{Base: 10},
		AllowedVals:     containers.HashSetValInit[int, widgets.BuiltinInt](1),
	}

	res, err := l.Translate("1")
	test.Nil(err, t)
	test.Eq(res, 1, t)
}

func TestSelectorFailing(t *testing.T) {
	l := Selector[int, BuiltinInt, widgets.BuiltinInt]{
		ValueTranslator: BuiltinInt{Base: 10},
		AllowedVals:     containers.HashSetValInit[int, widgets.BuiltinInt](1, 2, 3),
	}

	_, err := l.Translate("4")
	test.ContainsError(customerr.InvalidValue, err, t)
	test.ContainsError(ValNotInAllowedListErr, err, t)
}
