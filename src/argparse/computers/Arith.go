package computers

import "github.com/barbell-math/util/src/math/basic"

type (
	Add[T basic.Number] struct {
		L *T
		R *T
	}

	Sub[T basic.Number] struct {
		L *T
		R *T
	}

	Mul[T basic.Number] struct {
		L *T
		R *T
	}

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
