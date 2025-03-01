package translators

import (
	"math/big"
)

//go:generate ../../../bin/ifaceImplCheck -typeToCheck=BigInt
//go:generate ../../../bin/ifaceImplCheck -typeToCheck=BigRat
//go:generate ../../../bin/ifaceImplCheck -typeToCheck=BigFloat

type (
	// Represents a cmd line argument that will be translated to a [big.Int] type.
	//gen:ifaceImplCheck ifaceName Translator[big.Int]
	//gen:ifaceImplCheck imports math/big
	//gen:ifaceImplCheck valOrPntr both
	BigInt struct{}

	// Represents a cmd line argument that will be translated to a [big.Rat] type.
	//gen:ifaceImplCheck ifaceName Translator[big.Rat]
	//gen:ifaceImplCheck imports math/big
	//gen:ifaceImplCheck valOrPntr both
	BigRat struct{}

	// Represents a cmd line argument that will be translated to a [big.Float] type.
	//gen:ifaceImplCheck ifaceName Translator[big.Float]
	//gen:ifaceImplCheck imports math/big
	//gen:ifaceImplCheck valOrPntr both
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
