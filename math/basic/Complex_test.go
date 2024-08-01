package basic

import (
	"testing"

	"github.com/barbell-math/util/test"
)

func TestRealPart64(t *testing.T) {
	var c complex64
	test.Eq(float32(0), RealPart[complex64, float32](c), t)
	test.Eq(float64(0), RealPart[complex64, float64](c), t)

	c = complex(5, 0)
	test.Eq(float32(5), RealPart[complex64, float32](c), t)
	test.Eq(float64(5), RealPart[complex64, float64](c), t)
}

func TestRealPart128(t *testing.T) {
	var c complex128
	test.Eq(float32(0), RealPart[complex128, float32](c), t)
	test.Eq(float64(0), RealPart[complex128, float64](c), t)

	c = complex(5, 0)
	test.Eq(float32(5), RealPart[complex128, float32](c), t)
	test.Eq(float64(5), RealPart[complex128, float64](c), t)
}

func TestImaginaryPart64(t *testing.T) {
	var c complex64
	test.Eq(float32(0), ImaginaryPart[complex64, float32](c), t)
	test.Eq(float64(0), ImaginaryPart[complex64, float64](c), t)

	c = complex(0, 5)
	test.Eq(float32(5), ImaginaryPart[complex64, float32](c), t)
	test.Eq(float64(5), ImaginaryPart[complex64, float64](c), t)
}

func TestImaginaryPart128(t *testing.T) {
	var c complex128
	test.Eq(float32(0), ImaginaryPart[complex128, float32](c), t)
	test.Eq(float64(0), ImaginaryPart[complex128, float64](c), t)

	c = complex(0, 5)
	test.Eq(float32(5), ImaginaryPart[complex128, float32](c), t)
	test.Eq(float64(5), ImaginaryPart[complex128, float64](c), t)
}
