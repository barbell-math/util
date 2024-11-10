package argparse

import (
	"testing"

	"github.com/barbell-math/util/argparse/translators"
	"github.com/barbell-math/util/test"
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

// TODO - delete if subparsers work
// func TestArgBuilderToParserReservedShortName(t *testing.T) {
// 	res := ""
// 	b := ArgBuilder{}
// 
// 	AddArg[string, translators.BuiltinString](
// 		&res,
// 		&b,
// 		"str",
// 		NewOpts[string, translators.BuiltinString]().
// 			SetShortName('h'),
// 	)
// 
// 	_, err := b.ToParser("", "")
// 	test.ContainsError(ParserConfigErr, err, t)
// 	test.ContainsError(ReservedShortNameErr, err, t)
// }
// 
// func TestArgBuilderToParserReservedLongName(t *testing.T) {
// 	res := ""
// 	b := ArgBuilder{}
// 
// 	AddArg[string, translators.BuiltinString](
// 		&res,
// 		&b,
// 		"help",
// 		NewOpts[string, translators.BuiltinString]().
// 			SetShortName('s'),
// 	)
// 
// 	_, err := b.ToParser("", "")
// 	test.ContainsError(ParserConfigErr, err, t)
// 	test.ContainsError(ReservedLongNameErr, err, t)
// }

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
