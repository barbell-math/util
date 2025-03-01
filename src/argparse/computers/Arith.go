package computers

import "github.com/barbell-math/util/src/math/basic"

//go:generate ../../../bin/ifaceImplCheck -typeToCheck=Add
//go:generate ../../../bin/ifaceImplCheck -typeToCheck=Sub
//go:generate ../../../bin/ifaceImplCheck -typeToCheck=Mul
//go:generate ../../../bin/ifaceImplCheck -typeToCheck=Div

type (
	//gen:ifaceImplCheck generics [int]
	//gen:ifaceImplCheck ifaceName Computer[int]
	//gen:ifaceImplCheck valOrPntr both
	Add[T basic.Number] struct {
		L *T
		R *T
	}

	//gen:ifaceImplCheck generics [int]
	//gen:ifaceImplCheck ifaceName Computer[int]
	//gen:ifaceImplCheck valOrPntr both
	Sub[T basic.Number] struct {
		L *T
		R *T
	}

	//gen:ifaceImplCheck generics [int]
	//gen:ifaceImplCheck ifaceName Computer[int]
	//gen:ifaceImplCheck valOrPntr both
	Mul[T basic.Number] struct {
		L *T
		R *T
	}

	//gen:ifaceImplCheck generics [int]
	//gen:ifaceImplCheck ifaceName Computer[int]
	//gen:ifaceImplCheck valOrPntr both
	Div[T basic.Number] struct {
		L *T
		R *T
	}
)

func (a Add[T]) ComputeVals() (T, error) {
	return *a.L + *a.R, nil
}

func (a Add[T]) Reset() {
	// intentional noop - Add has no state that needs to be reset
}

func (a Sub[T]) ComputeVals() (T, error) {
	return *a.L - *a.R, nil
}

func (a Sub[T]) Reset() {
	// intentional noop - Sub has no state that needs to be reset
}

func (a Mul[T]) ComputeVals() (T, error) {
	return *a.L * *a.R, nil
}

func (a Mul[T]) Reset() {
	// intentional noop - Mul has no state that needs to be reset
}

func (a Div[T]) ComputeVals() (T, error) {
	return *a.L / *a.R, nil
}

func (a Div[T]) Reset() {
	// intentional noop - Div has no state that needs to be reset
}
