package translators

// Code generated by ../../../bin/ifaceImplCheck - DO NOT EDIT.
import (
	"os"
	"testing"
)

func TestOpenFileValueImplementsTranslator__os_File_(t *testing.T) {
	var typeThing OpenFile
	var iFaceThing Translator[*os.File] = typeThing
	_ = iFaceThing
}

func TestOpenFilePntrImplementsTranslator__os_File_(t *testing.T) {
	var typeThing OpenFile
	var iFaceThing Translator[*os.File] = &typeThing
	_ = iFaceThing
}
