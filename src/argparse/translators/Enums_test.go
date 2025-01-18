package translators

import (
	"testing"

	"github.com/barbell-math/util/src/test"
)

func TestEnum(t *testing.T) {
	e := Enum[testEnum, *testEnum]{}

	v, err := e.Translate("asdf")
	test.Eq(v, unknownTestEnum, t)
	test.ContainsError(InvalidTestEnum, err, t)

	v, err = e.Translate("unknownTestEnum")
	test.Eq(v, unknownTestEnum, t)
	test.Nil(err, t)

	v, err = e.Translate("oneTestEnum")
	test.Eq(v, oneTestEnum, t)
	test.Nil(err, t)

	v, err = e.Translate("twoTestEnum")
	test.Eq(v, twoTestEnum, t)
	test.Nil(err, t)
}
