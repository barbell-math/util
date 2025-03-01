package translators

//go:generate ../../../bin/ifaceImplCheck -typeToCheck=Flag

type (
	// Used to represent a flag that can only be supplied once, returning a
	// boolean value indicating the flags presence.
	//gen:ifaceImplCheck ifaceName Translator[bool]
	//gen:ifaceImplCheck valOrPntr both
	Flag struct{}
)

func (_ Flag) Translate(arg string) (bool, error) {
	return true, nil
}

func (_ Flag) Reset() {
	// intentional noop - Flag has no state
}
