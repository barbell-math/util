package translators

//go:generate ../../../bin/ifaceImplCheck -typeToCheck=Stopper

type (
	// Used to represent a flag that when encountered should stop parsing of the
	// cmd line arguments. The error that the stopper is created with will be
	// returned when [Stopper.Translate] is called.
	//gen:ifaceImplCheck generics [bool]
	//gen:ifaceImplCheck ifaceName Translator[bool]
	//gen:ifaceImplCheck valOrPntr both
	Stopper[T any] struct{ Err error }
)

func (s Stopper[T]) Translate(arg string) (T, error) {
	var tmp T
	return tmp, s.Err
}

func (s Stopper[T]) Reset() {
	// intentional noop - Stopper has no state that needs to reset
}
