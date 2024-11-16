package translators

import "strconv"

type (
	// Represents a cmd line argument that will be translated to a float32 type.
	BuiltinFloat32 struct{}

	// Represents a cmd line argument that will be translated to a floa64 type.
	BuiltinFloat64 struct{}
)

func (_ BuiltinFloat32) Translate(arg string) (float32, error) {
	f64, err := strconv.ParseFloat(arg, 32)
	return float32(f64), err
}

func (_ BuiltinFloat32) Reset() {
	// intentional noop - BuiltinFloat32 has no state
}

func (_ BuiltinFloat64) Translate(arg string) (float64, error) {
	f64, err := strconv.ParseFloat(arg, 32)
	return f64, err
}

func (_ BuiltinFloat64) Reset() {
	// intentional noop - BuiltinFloat64 has no state
}
