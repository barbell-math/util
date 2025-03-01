package translators

//go:generate ../../../bin/ifaceImplCheck -typeToCheck=Noop

type (
	// Used to represent a cmd line argument that will always be the default
	// zero-value initilized T type.
	//gen:ifaceImplCheck generics [bool]
	//gen:ifaceImplCheck ifaceName Translator[bool]
	//gen:ifaceImplCheck valOrPntr both
	Noop[T any] struct{}
)

func (_ Noop[T]) Translate(arg string) (T, error) {
	var tmp T
	return tmp, nil
}

func (_ Noop[T]) Reset() {
	// intentional noop - Noop has no state
}
