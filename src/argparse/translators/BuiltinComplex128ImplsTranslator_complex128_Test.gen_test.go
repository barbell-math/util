package translators

// Code generated by ../../../bin/ifaceImplCheck - DO NOT EDIT.
import (
	"testing"
)

func TestBuiltinComplex128ValueImplementsTranslator_complex128_(t *testing.T) {
	var typeThing BuiltinComplex128
	var iFaceThing Translator[complex128] = typeThing
	_ = iFaceThing
}

func TestBuiltinComplex128PntrImplementsTranslator_complex128_(t *testing.T) {
	var typeThing BuiltinComplex128
	var iFaceThing Translator[complex128] = &typeThing
	_ = iFaceThing
}
