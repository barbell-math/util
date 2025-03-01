package computers

//go:generate ../../../bin/ifaceImplCheck -typeToCheck=Stopper

type (
	//gen:ifaceImplCheck generics [int]
	//gen:ifaceImplCheck ifaceName Computer[int]
	//gen:ifaceImplCheck valOrPntr both
	Stopper[T any] struct{ Err error }
)

func (s Stopper[T]) ComputeVals() (T, error) {
	var tmp T
	return tmp, s.Err
}

func (s Stopper[T]) Reset() {
	// intentional noop - Stopper has no state that needs to be reset
}
