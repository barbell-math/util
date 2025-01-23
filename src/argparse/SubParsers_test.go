package argparse

import (
	"testing"

	"github.com/barbell-math/util/src/argparse/translators"
	"github.com/barbell-math/util/src/test"
)

func TestHelpSubParser(t *testing.T) {
	res := struct {
		S  string
		S1 string
		S2 string
	}{}

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
	AddArg[string, translators.BuiltinString](
		&res.S,
		&b,
		"str1",
		NewOpts[string, translators.BuiltinString]().
			SetDescription("this is a long description that needs to break 80 characters to I can see how the split works").
			SetDefaultVal("default"),
	)
	AddArg[string, translators.BuiltinString](
		&res.S,
		&b,
		"str2",
		NewOpts[string, translators.BuiltinString]().
			SetDescription("this is a long description that needs to break 80 characters to I can see how the split works").
			SetDefaultVal("default").
			SetConditionallyRequired([]ArgConditionality[string]{
				ArgConditionality[string]{
					Requires: []string{"str1"},
				},
			}),
	)
	p, err := b.ToParser("testProg", "this is a long description")
	test.Nil(err, t)
	err = p.AddSubParsers(NewHelpParser())
	test.Nil(err, t)

	err = p.Parse(ArgvIterFromSlice([]string{"-h, -s=123"}).ToTokens())
	test.ContainsError(HelpErr, err, t)
}
