package translators

import (
	"testing"

	"github.com/barbell-math/util/src/customerr"
	"github.com/barbell-math/util/src/test"
)

func TestRange(t *testing.T) {
	r := Range[int, BuiltinInt]{
		Min:           3,
		Max:           5,
		NumTranslator: BuiltinInt{Base: 10},
	}

	_, err := r.Translate("2")
	test.ContainsError(customerr.ValOutsideRange, err, t)

	v, err := r.Translate("3")
	test.Nil(err, t)
	test.Eq(v, 3, t)
	v, err = r.Translate("4")
	test.Nil(err, t)
	test.Eq(v, 4, t)

	_, err = r.Translate("5")
	test.ContainsError(customerr.ValOutsideRange, err, t)
	_, err = r.Translate("6")
	test.ContainsError(customerr.ValOutsideRange, err, t)
}
