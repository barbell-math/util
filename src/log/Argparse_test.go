package log

import (
	"testing"

	"github.com/barbell-math/util/src/argparse"
	"github.com/barbell-math/util/src/argparse/translators"
	"github.com/barbell-math/util/src/test"
)

func TestVerbositySubParser(t *testing.T) {
	res := struct{ I int }{}

	p, err := (&argparse.ArgBuilder{}).ToParser("", "")
	test.Nil(err, t)
	err = p.AddSubParsers(NewVerbosityParser[int](&res.I, -1, 6))
	test.Nil(err, t)

	res.I = 0
	err = p.Parse(argparse.ArgvIterFromSlice([]string{}).ToTokens())
	test.Nil(err, t)
	test.Eq(res.I, -1, t)

	res.I = 0
	err = p.Parse(argparse.ArgvIterFromSlice([]string{"-v"}).ToTokens())
	test.Nil(err, t)
	test.Eq(res.I, 1, t)

	res.I = 0
	err = p.Parse(argparse.ArgvIterFromSlice([]string{"-v", "-v"}).ToTokens())
	test.Nil(err, t)
	test.Eq(res.I, 2, t)

	res.I = 0
	err = p.Parse(argparse.ArgvIterFromSlice([]string{"-v", "--verbose"}).ToTokens())
	test.Nil(err, t)
	test.Eq(res.I, 2, t)

	res.I = 0
	err = p.Parse(argparse.ArgvIterFromSlice([]string{"-vv"}).ToTokens())
	test.Nil(err, t)
	test.Eq(res.I, 2, t)

	res.I = 0
	err = p.Parse(argparse.ArgvIterFromSlice([]string{"-vvv", "--verbose", "-vv"}).
		ToTokens())
	test.Nil(err, t)
	test.Eq(res.I, 6, t)

	res.I = 0
	err = p.Parse(argparse.ArgvIterFromSlice([]string{"-vvv", "--verbose", "-vvv"}).
		ToTokens())
	test.ContainsError(translators.FlagProvidedToManyTimesErr, err, t)
	test.Eq(res.I, 6, t)
}
