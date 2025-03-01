package translators

import (
	"strings"

	"github.com/barbell-math/util/src/customerr"
)

//go:generate ../../../bin/ifaceImplCheck -typeToCheck=Bits

type (
	// Represents a cmd line argument that is expected to be a sequence of bits.
	// The bits will be stored in a byte slice. The number of bits does not need
	// to be a multiple of 4, the first value will be padded with 0's as
	// necessary. Arbitrary spaces may be included in the string input and will
	// be ignored.
	// The bits will be stored in the slice in the order that they were given.
	//     1010 1110 -> []byte{0b1010, 0b1110}
	//   1 1010 1110 -> []byte{0b0001, 0b1010, 0b1110}
	//  10 1010 1110 -> []byte{0b0010, 0b1010, 0b1110}
	// 100 1010 1110 -> []byte{0b0100, 0b1010, 0b1110}
	//gen:ifaceImplCheck ifaceName Translator[[]byte]
	//gen:ifaceImplCheck valOrPntr both
	Bits struct{}
)

func (_ Bits) Translate(arg string) ([]byte, error) {
	arg = strings.ReplaceAll(arg, " ", "")
	if len(arg)%4 != 0 {
		padding := strings.Repeat("0", 4-len(arg)%4)
		arg = padding + arg
	}

	rv := make([]byte, len(arg)/4)

	var i int = len(arg) - 1
	var c byte
	var iterVal byte
	for ; i >= 0; i-- {
		c = arg[i]
		if c != '1' && c != '0' {
			return rv, customerr.Wrap(
				BitsTranslationErr,
				"All bit values must be either '1' or '0'. Got: '%c' at index %d",
				c, i,
			)
		}

		iterVal >>= 1
		if c == '1' {
			iterVal |= 0b1000
		}

		if i%4 == 0 {
			rv[int(i/4)] = iterVal
			iterVal = 0
		}
	}

	return rv, nil
}

func (_ Bits) Reset() {
	// intentional noop - Bits has no state that needs to be reset
}
