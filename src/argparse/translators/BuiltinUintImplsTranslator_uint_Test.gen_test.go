package translators

// Code generated by ../../../bin/ifaceImplCheck - DO NOT EDIT.
import (
	"testing"
)

func TestBuiltinUintValueImplementsTranslator_uint_(t *testing.T) {
	var typeThing BuiltinUint
	var iFaceThing Translator[uint] = typeThing
	_ = iFaceThing
}

func TestBuiltinUintPntrImplementsTranslator_uint_(t *testing.T) {
	var typeThing BuiltinUint
	var iFaceThing Translator[uint] = &typeThing
	_ = iFaceThing
}
