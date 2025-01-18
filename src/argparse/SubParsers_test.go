package argparse

import (
	"testing"

	testenum "github.com/barbell-math/util/src/argparse/testEnum"
	"github.com/barbell-math/util/src/argparse/translators"
	"github.com/barbell-math/util/src/test"
)

func TestHelpSubParser(t *testing.T) {
	res := struct{ S string }{}

	b := ArgBuilder{}
	AddArg[string, translators.BuiltinString](
		&res.S,
		&b,
		"str",
		NewOpts[string, translators.BuiltinString]().
			SetShortName('s').
			SetDescription("this is a long description that needs to break 80 characters to I can see how the split works").
			SetDefaultVal("default").
			SetRequired(true),
	)
	p, err := b.ToParser("testProg", "this is a long description")
	test.Nil(err, t)
	err = p.AddSubParsers(NewHelpParser())
	test.Nil(err, t)

	err = p.Parse(ArgvIterFromSlice([]string{"-h, -s=123"}).ToTokens())
	test.ContainsError(HelpErr, err, t)
}

func TestVerbositySubParser(t *testing.T) {
	res := struct{ I int }{}

	p, err := (&ArgBuilder{}).ToParser("", "")
	test.Nil(err, t)
	err = p.AddSubParsers(NewVerbosityParser[int](&res.I))
	test.Nil(err, t)

	res.I = 0
	err = p.Parse(ArgvIterFromSlice([]string{"-v"}).ToTokens())
	test.Nil(err, t)
	test.Eq(res.I, 1, t)

	res.I = 0
	err = p.Parse(ArgvIterFromSlice([]string{"-v", "-v"}).ToTokens())
	test.Nil(err, t)
	test.Eq(res.I, 2, t)

	res.I = 0
	err = p.Parse(ArgvIterFromSlice([]string{"-v", "--verbose"}).ToTokens())
	test.Nil(err, t)
	test.Eq(res.I, 2, t)

	res.I = 0
	err = p.Parse(ArgvIterFromSlice([]string{"-vv"}).ToTokens())
	test.Nil(err, t)
	test.Eq(res.I, 2, t)

	res.I = 0
	err = p.Parse(ArgvIterFromSlice([]string{"-vvv", "--verbose", "-vv"}).
		ToTokens())
	test.Nil(err, t)
	test.Eq(res.I, 6, t)
}

func TestAppActionParser(t *testing.T) {
	res := struct{ Action testenum.TestEnum }{}

	p, err := (&ArgBuilder{}).ToParser("", "")
	test.Nil(err, t)
	err = p.AddSubParsers(
		NewAppActionParser[testenum.TestEnum, *testenum.TestEnum](&res.Action),
	)
	test.Nil(err, t)

	res.Action = testenum.UnknownTestEnum
	err = p.Parse(ArgvIterFromSlice([]string{"-a", "asdf"}).ToTokens())
	test.ContainsError(testenum.InvalidTestEnum, err, t)
	test.Eq(testenum.UnknownTestEnum, res.Action, t)

	res.Action = testenum.UnknownTestEnum
	err = p.Parse(ArgvIterFromSlice([]string{"-a", "unknownTestEnum"}).ToTokens())
	test.Nil(err, t)
	test.Eq(testenum.UnknownTestEnum, res.Action, t)

	res.Action = testenum.UnknownTestEnum
	err = p.Parse(ArgvIterFromSlice([]string{"-a", "oneTestEnum"}).ToTokens())
	test.Nil(err, t)
	test.Eq(testenum.OneTestEnum, res.Action, t)

	res.Action = testenum.UnknownTestEnum
	err = p.Parse(ArgvIterFromSlice([]string{"-a", "twoTestEnum"}).ToTokens())
	test.Nil(err, t)
	test.Eq(testenum.TwoTestEnum, res.Action, t)
}
