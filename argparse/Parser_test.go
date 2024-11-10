package argparse

import (
	"testing"

	"github.com/barbell-math/util/argparse/translators"
	"github.com/barbell-math/util/test"
)

func TestParserParseDuplicateValArg(t *testing.T) {
	res:=struct { S string }{}

	b := ArgBuilder{}
	AddArg[string, translators.BuiltinString](
		&res.S,
		&b, 
		"str",
		NewOpts[string, translators.BuiltinString]().SetShortName('s'),
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	err=p.Parse(ArgvIterFromSlice([]string{"--str=123", "-s=456"}).ToTokens())
	test.ContainsError(ParsingErr, err, t)
	test.ContainsError(ArgumentPassedMultipleTimesErr, err, t)
}

func TestParserParseDuplicateFlagArg(t *testing.T) {
	res:=struct { B bool }{}

	b := ArgBuilder{}
	AddFlag(
		&res.B,
		&b, 
		"bool",
		NewOpts[bool, translators.Flag]().SetShortName('b'),
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	err=p.Parse(ArgvIterFromSlice([]string{"--bool", "-b"}).ToTokens())
	test.ContainsError(ParsingErr, err, t)
	test.ContainsError(ArgumentPassedMultipleTimesErr, err, t)
}

func TestParserParseMissingRequiredArgs(t *testing.T) {
	res:=struct { B bool }{}

	b := ArgBuilder{}
	AddFlag(
		&res.B,
		&b, 
		"bool",
		NewOpts[bool, translators.Flag]().SetShortName('b').SetRequired(true),
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	err=p.Parse(ArgvIterFromSlice([]string{}).ToTokens())
	test.ContainsError(ParsingErr, err, t)
	test.ContainsError(MissingRequiredArgErr, err, t)
}

func TestParserParseDefaultValue(t *testing.T) {
	res:=struct { S string }{}

	b := ArgBuilder{}
	AddArg[string, translators.BuiltinString](
		&res.S,
		&b, 
		"str",
		NewOpts[string, translators.BuiltinString]().
			SetShortName('s').
			SetDefaultVal("default"),
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	err=p.Parse(ArgvIterFromSlice([]string{}).ToTokens())
	test.Nil(err, t)
	test.Eq(res.S, "default", t)
}

func TestParserParseHelpPrintout(t*testing.T) {
	res:=struct { S string }{}

	b := ArgBuilder{}
	AddArg[string, translators.BuiltinString](
		&res.S,
		&b, 
		"str",
		NewOpts[string, translators.BuiltinString]().
			SetShortName('s').
			SetDefaultVal("default"),
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	err=p.Parse(ArgvIterFromSlice([]string{"-h"}).ToTokens())
	test.Nil(err, t)
}

// TODO - test combine

// TODO - load args from file (JSON)
// TODO - translators generator (strconv)
// TODO - doc strings
