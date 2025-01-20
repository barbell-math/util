package translators

import (
	"testing"

	"github.com/barbell-math/util/src/test"
)

func TestLimitedFlagCntr(t *testing.T) {
	l := LimitedFlagCntr[int]{MaxTimes: 3}

	c, err := l.Translate("asdf")
	test.Nil(err, t)
	test.Eq(c, 1, t)
	c, err = l.Translate("asdf")
	test.Nil(err, t)
	test.Eq(c, 2, t)
	c, err = l.Translate("asdf")
	test.Nil(err, t)
	test.Eq(c, 3, t)
	c, err = l.Translate("asdf")
	test.ContainsError(FlagProvidedToManyTimesErr, err, t)
	test.Eq(c, 3, t)
}
