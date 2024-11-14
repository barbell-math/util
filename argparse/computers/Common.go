package computers

type (
	// An interface that defines how computed arguments will be evaluated. The
	// upderlying value is expected to have some way to access the data it needs
	// to compute the value T. Normally, the underlying value will be a struct
	// with a pointer to value(s) that it needs.
	Computer[T any] interface {
		ComputeVals() (T, error)
		// Resets the state of the Computers's underlying value.
		Reset()
	}
)
