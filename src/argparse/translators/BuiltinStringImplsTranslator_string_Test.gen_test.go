package translators

// Code generated by ../../../bin/ifaceImplCheck - DO NOT EDIT.
import (
	"testing"
)

func TestBuiltinStringValueImplementsTranslator_string_(t *testing.T) {
	var typeThing BuiltinString
	var iFaceThing Translator[string] = typeThing
	_ = iFaceThing
}

func TestBuiltinStringPntrImplementsTranslator_string_(t *testing.T) {
	var typeThing BuiltinString
	var iFaceThing Translator[string] = &typeThing
	_ = iFaceThing
}
