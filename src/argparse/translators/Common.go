package translators

type (
	// An interface that defines what actions can be performed when translating
	// a string argument to a typed value. The translator is expected to perform
	// all validation required to ensure a correct value is returned. It is also
	// expected to return an error if a value is found to be invalid.
	Translater[T any] interface {
		Translate(arg string) (T, error)
		// Resets the state of the Translater's underlying value.
		Reset()
	}
)
