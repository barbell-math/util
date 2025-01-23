package argparse

import (
	"testing"

	"github.com/barbell-math/util/src/argparse/translators"
	"github.com/barbell-math/util/src/test"
)

func TestArgBuilderToParserDuplicateShortNames(t *testing.T) {
	res := ""
	b := ArgBuilder{}

	AddArg[string, translators.BuiltinString](
		&res,
		&b,
		"str",
		NewOpts[string, translators.BuiltinString]().
			SetShortName('s'),
	)
	AddArg[string, translators.BuiltinString](
		&res,
		&b,
		"str2",
		NewOpts[string, translators.BuiltinString]().
			SetShortName('s'),
	)

	_, err := b.ToParser("", "")
	test.ContainsError(ParserConfigErr, err, t)
	test.ContainsError(DuplicateShortNameErr, err, t)
}

func TestArgBuilderToParserMissingShortNames(t *testing.T) {
	res := ""
	b := ArgBuilder{}

	AddArg[string, translators.BuiltinString](
		&res,
		&b,
		"str",
		NewOpts[string, translators.BuiltinString](),
	)
	AddArg[string, translators.BuiltinString](
		&res,
		&b,
		"str2",
		NewOpts[string, translators.BuiltinString](),
	)

	_, err := b.ToParser("", "")
	test.Nil(err, t)
}

func TestArgBuilderToParserDuplicateLongNames(t *testing.T) {
	res := ""
	b := ArgBuilder{}

	AddArg[string, translators.BuiltinString](
		&res,
		&b,
		"str",
		NewOpts[string, translators.BuiltinString]().
			SetShortName('s'),
	)
	AddArg[string, translators.BuiltinString](
		&res,
		&b,
		"str",
		NewOpts[string, translators.BuiltinString]().
			SetShortName('t'),
	)

	_, err := b.ToParser("", "")
	test.ContainsError(ParserConfigErr, err, t)
	test.ContainsError(DuplicateLongNameErr, err, t)
}

func TestArgBuilderToParserInvalidConfigLongName(t *testing.T) {
	res := ""
	b := ArgBuilder{}

	AddArg[string, translators.BuiltinString](
		&res,
		&b,
		"config",
		NewOpts[string, translators.BuiltinString]().
			SetShortName('c'),
	)

	_, err := b.ToParser("", "")
	test.ContainsError(ParserConfigErr, err, t)
	test.ContainsError(ReservedLongNameErr, err, t)
}

func TestArgBuilderToParserValidArgBuilder(t *testing.T) {
	res := ""
	b := ArgBuilder{}

	AddArg[string, translators.BuiltinString](
		&res,
		&b,
		"str",
		NewOpts[string, translators.BuiltinString]().
			SetShortName('s'),
	)
	AddArg[string, translators.BuiltinString](
		&res,
		&b,
		"str2",
		NewOpts[string, translators.BuiltinString]().
			SetShortName('t'),
	)

	_, err := b.ToParser("", "")
	test.Nil(err, t)
}

func TestArgBuilderLongNameToShort(t *testing.T) {
	res := ""
	b := ArgBuilder{}

	AddArg[string, translators.BuiltinString](
		&res,
		&b,
		"s",
		NewOpts[string, translators.BuiltinString]().
			SetShortName('t'),
	)

	_, err := b.ToParser("", "")
	test.ContainsError(ParserConfigErr, err, t)
	test.ContainsError(LongNameToShortErr, err, t)
}

func TestArgBuilderUnrecognizedConditionallyRequiredArg(t *testing.T) {
	res := ""
	b := ArgBuilder{}

	AddArg[string, translators.BuiltinString](
		&res,
		&b,
		"str",
		NewOpts[string, translators.BuiltinString]().
			SetShortName('t').
			SetConditionallyRequired([]ArgConditionality[string]{
				ArgConditionality[string]{
					Requires: []string{"foo"},
					When:     ArgSupplied[string],
				},
			}),
	)

	_, err := b.ToParser("", "")
	test.ContainsError(ParserConfigErr, err, t)
	test.ContainsError(UnrecognizedConditionallyRequiredArgErr, err, t)
}
