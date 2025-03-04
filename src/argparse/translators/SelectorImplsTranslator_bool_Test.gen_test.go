package translators

// Code generated by ../../../bin/ifaceImplCheck - DO NOT EDIT.
import (
	"github.com/barbell-math/util/src/widgets"
	"testing"
)

func TestSelectorValueImplementsTranslator_bool_(t *testing.T) {
	var typeThing Selector[BuiltinBool, widgets.BuiltinBool, bool]
	var iFaceThing Translator[bool] = typeThing
	_ = iFaceThing
}

func TestSelectorPntrImplementsTranslator_bool_(t *testing.T) {
	var typeThing Selector[BuiltinBool, widgets.BuiltinBool, bool]
	var iFaceThing Translator[bool] = &typeThing
	_ = iFaceThing
}
