package translators

import (
	"testing"

	"github.com/barbell-math/util/src/test"
)

func TestBits(t *testing.T) {
	b := Bits{}

	_, err := b.Translate("00asdf")
	test.ContainsError(BitsTranslationErr, err, t)

	res, err := b.Translate("00000000")
	test.Nil(err, t)
	test.SlicesMatch[byte]([]byte{0, 0}, res, t)

	res, err = b.Translate("10101010")
	test.Nil(err, t)
	test.SlicesMatch[byte]([]byte{0b1010, 0b1010}, res, t)

	res, err = b.Translate("10101010")
	test.Nil(err, t)
	test.SlicesMatch[byte]([]byte{0b1010, 0b1010}, res, t)

	res, err = b.Translate("10100101")
	test.Nil(err, t)
	test.SlicesMatch[byte]([]byte{0b1010, 0b0101}, res, t)

	res, err = b.Translate("101001011010")
	test.Nil(err, t)
	test.SlicesMatch[byte]([]byte{0b1010, 0b0101, 0b1010}, res, t)

	res, err = b.Translate("1 0100 1011 0101")
	test.Nil(err, t)
	test.SlicesMatch[byte]([]byte{0b0001, 0b0100, 0b1011, 0b0101}, res, t)

	res, err = b.Translate("10 0100 1011 0101")
	test.Nil(err, t)
	test.SlicesMatch[byte]([]byte{0b0010, 0b0100, 0b1011, 0b0101}, res, t)

	res, err = b.Translate("100 0100 1011 0101")
	test.Nil(err, t)
	test.SlicesMatch[byte]([]byte{0b0100, 0b0100, 0b1011, 0b0101}, res, t)

	res, err = b.Translate("")
	test.Nil(err, t)
	test.SlicesMatch[byte]([]byte{}, res, t)

	res, err = b.Translate("1")
	test.Nil(err, t)
	test.SlicesMatch[byte]([]byte{0b0001}, res, t)

	res, err = b.Translate("10")
	test.Nil(err, t)
	test.SlicesMatch[byte]([]byte{0b0010}, res, t)

	res, err = b.Translate("100")
	test.Nil(err, t)
	test.SlicesMatch[byte]([]byte{0b0100}, res, t)
}
