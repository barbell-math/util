package translators

import "strconv"

//go:generate ../../../bin/ifaceImplCheck -typeToCheck=BuiltinBool

type (
	// Represents a cmd line argument that will be translated to a bool type.
	//gen:ifaceImplCheck ifaceName Translator[bool]
	//gen:ifaceImplCheck valOrPntr both
	BuiltinBool struct{}
)

func (_ BuiltinBool) Translate(arg string) (bool, error) {
	c128, err := strconv.ParseBool(arg)
	return c128, err
}

func (_ BuiltinBool) Reset() {
	// intentional noop - Bool has no state
}
