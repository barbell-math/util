package translators

// Code generated by ../../../bin/ifaceImplCheck - DO NOT EDIT.
import (
	"testing"
)

func TestBuiltinFloat32ValueImplementsTranslator_float32_(t *testing.T) {
	var typeThing BuiltinFloat32
	var iFaceThing Translator[float32] = typeThing
	_ = iFaceThing
}

func TestBuiltinFloat32PntrImplementsTranslator_float32_(t *testing.T) {
	var typeThing BuiltinFloat32
	var iFaceThing Translator[float32] = &typeThing
	_ = iFaceThing
}
