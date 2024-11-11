package translators

import "strconv"

type (
	// Represents a cmd line argument that will be translated to a complex64 type.
	BuiltinComplex64 struct{}

	// Represents a cmd line argument that will be translated to a complex128 type.
	BuiltinComplex128 struct{}
)

func (_ BuiltinComplex64) Translate(arg string) (complex64, error) {
	c64, err := strconv.ParseComplex(arg, 64)
	return complex64(c64), err
}

func (_ BuiltinComplex64) Reset() {
	// intentional noop - Complex64 has no state
}

func (_ BuiltinComplex128) Translate(arg string) (complex128, error) {
	c128, err := strconv.ParseComplex(arg, 32)
	return c128, err
}

func (_ BuiltinComplex128) Reset() {
	// intentional noop - Complex128 has no state
}