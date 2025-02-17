package translators

import (
	"testing"

	testenum "github.com/barbell-math/util/src/argparse/testEnum"
	"github.com/barbell-math/util/src/test"
)

func TestEnum(t *testing.T) {
	e := Enum[*testenum.TestEnum, testenum.TestEnum]{}

	v, err := e.Translate("asdf")
	test.Eq(v, testenum.UnknownTestEnum, t)
	test.ContainsError(testenum.InvalidTestEnum, err, t)

	v, err = e.Translate("unknownTestEnum")
	test.Eq(v, testenum.UnknownTestEnum, t)
	test.Nil(err, t)

	v, err = e.Translate("oneTestEnum")
	test.Eq(v, testenum.OneTestEnum, t)
	test.Nil(err, t)

	v, err = e.Translate("twoTestEnum")
	test.Eq(v, testenum.TwoTestEnum, t)
	test.Nil(err, t)
}
