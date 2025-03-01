package translators

//go:generate ../../../bin/ifaceImplCheck -typeToCheck=BuiltinString

type (
	// Represents a cmd line argument that will be translated to a string type.
	//gen:ifaceImplCheck ifaceName Translator[string]
	//gen:ifaceImplCheck valOrPntr both
	BuiltinString struct{}
)

func (_ BuiltinString) Translate(arg string) (string, error) {
	return arg, nil
}

func (_ BuiltinString) Reset() {
	// intentional noop - BuiltinString has no state
}
