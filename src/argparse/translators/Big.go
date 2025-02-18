package translators

import (
	"math/big"
)

type (
	// Represents a cmd line argument that will be translated to a [big.Int] type.
	BigInt struct{}

	// Represents a cmd line argument that will be translated to a [big.Rat] type.
	BigRat struct{}

	// Represents a cmd line argument that will be translated to a [big.Float] type.
	BigFloat struct{}
)

func (_ BigInt) Translate(arg string) (big.Int, error) {
	rv := big.Int{}
	_, ok := rv.SetString(arg, 10)
	if !ok {
		return big.Int{}, CouldNotParseBigIntErr
	}
	return rv, nil
}

func (_ BigInt) Reset() {
	// intentional noop - BigInt has no state that needs to be reset
}

func (_ BigFloat) Translate(arg string) (big.Float, error) {
	rv := big.Float{}
	_, ok := rv.SetString(arg)
	if !ok {
		return big.Float{}, CouldNotParseBigFloatErr
	}
	return rv, nil
}

func (_ BigFloat) Reset() {
	// intentional noop - BigFloat has no state that needs to be reset
}

func (_ BigRat) Translate(arg string) (big.Rat, error) {
	rv := big.Rat{}
	_, ok := rv.SetString(arg)
	if !ok {
		return big.Rat{}, CouldNotParseBigFloatErr
	}
	return rv, nil
}

func (_ BigRat) Reset() {
	// intentional noop - BigFloat has no state that needs to be reset
}
