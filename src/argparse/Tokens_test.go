package argparse

import (
	"testing"

	"github.com/barbell-math/util/src/argparse/translators"
	"github.com/barbell-math/util/src/iter"
	"github.com/barbell-math/util/src/test"
	"github.com/barbell-math/util/src/widgets"
)

func TestToTokensNoArguments(t *testing.T) {
	tokens, err := ArgvIterFromSlice([]string{}).ToTokens().ToIter().Collect()
	test.Nil(err, t)
	test.Eq(0, len(tokens), t)
}

func TestToTokensShortSpaceFlag(t *testing.T) {
	tokens, err := ArgvIterFromSlice([]string{"-t"}).ToTokens().ToIter().Collect()
	test.Nil(err, t)
	test.SlicesMatch[token](
		tokens,
		[]token{token{value: "t", _type: shortFlagToken}},
		t,
	)

	tokens, err = ArgvIterFromSlice([]string{"-tt"}).ToTokens().ToIter().Collect()
	test.Nil(err, t)
	test.SlicesMatch[token](
		tokens,
		[]token{
			token{value: "t", _type: shortFlagToken},
			token{value: "t", _type: shortFlagToken},
		},
		t,
	)

	tokens, err = ArgvIterFromSlice([]string{"-order"}).
		ToTokens().ToIter().Collect()
	test.Nil(err, t)
	test.SlicesMatch[token](
		tokens,
		[]token{
			token{value: "o", _type: shortFlagToken},
			token{value: "r", _type: shortFlagToken},
			token{value: "d", _type: shortFlagToken},
			token{value: "e", _type: shortFlagToken},
			token{value: "r", _type: shortFlagToken},
		},
		t,
	)
}

func TestToTokensShortEqualsFlag(t *testing.T) {
	tokens, err := ArgvIterFromSlice([]string{"-t=123"}).
		ToTokens().ToIter().Collect()
	test.Nil(err, t)
	test.SlicesMatch[token](
		tokens,
		[]token{
			token{value: "t", _type: shortFlagToken},
			token{value: "123", _type: valueToken},
		},
		t,
	)

	tokens, err = ArgvIterFromSlice([]string{"-tt=123"}).ToTokens().ToIter().Collect()
	test.Nil(err, t)
	test.SlicesMatch[token](
		tokens,
		[]token{
			token{value: "tt", _type: shortFlagToken},
			token{value: "123", _type: valueToken},
		},
		t,
	)
}

func TestToTokensLongSpaceFlag(t *testing.T) {
	tokens, err := ArgvIterFromSlice([]string{"--time"}).
		ToTokens().ToIter().Collect()
	test.Nil(err, t)
	test.SlicesMatch[token](
		tokens,
		[]token{token{value: "time", _type: longFlagToken}},
		t,
	)
}

func TestToTokensLongEqualsFlag(t *testing.T) {
	tokens, err := ArgvIterFromSlice([]string{"--time=123"}).
		ToTokens().ToIter().Collect()
	test.Nil(err, t)
	test.SlicesMatch[token](
		tokens,
		[]token{
			{value: "time", _type: longFlagToken},
			{value: "123", _type: valueToken},
		},
		t,
	)
}

func TestToTokensValue(t *testing.T) {
	tokens, err := ArgvIterFromSlice([]string{"value"}).
		ToTokens().ToIter().Collect()
	test.Nil(err, t)
	test.SlicesMatch[token](
		tokens,
		[]token{
			{value: "value", _type: valueToken},
		},
		t,
	)
}

func TestToTokensAllFlags(t *testing.T) {
	tokens, err := ArgvIterFromSlice(
		[]string{"-t", "--time", "123", "--time=123"},
	).ToTokens().ToIter().Collect()
	test.Nil(err, t)
	test.SlicesMatch[token](
		tokens,
		[]token{
			{value: "t", _type: shortFlagToken},
			{value: "time", _type: longFlagToken},
			{value: "123", _type: valueToken},
			{value: "time", _type: longFlagToken},
			{value: "123", _type: valueToken},
		},
		t,
	)
}

func TestToArgValPairsInvalidTokenType(t *testing.T) {
	res := struct{ S string }{}

	b := ArgBuilder{}
	AddArg[translators.BuiltinString](
		&res.S,
		&b,
		"str",
		NewOpts[translators.BuiltinString]().SetShortName('s'),
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	_, err = tokenIter(iter.SliceElems[token]([]token{
		{_type: unknownTokenType},
	})).toArgValPairs(&p).ToIter().Collect()
	test.ContainsError(InvalidTokenType, err, t)
}

func TestToArgValPairsExpectedArgument(t *testing.T) {
	res := struct{ S string }{}

	b := ArgBuilder{}
	AddArg[translators.BuiltinString](
		&res.S,
		&b,
		"str",
		NewOpts[translators.BuiltinString]().
			SetShortName('s'),
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	_, err = ArgvIterFromSlice([]string{"value"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.ContainsError(ExpectedArgumentErr, err, t)
}

func TestToArgValPairsUnrecognizedShortFlag(t *testing.T) {
	res := struct{ S string }{}

	b := ArgBuilder{}
	AddArg[translators.BuiltinString](
		&res.S,
		&b,
		"str",
		NewOpts[translators.BuiltinString]().
			SetShortName('s'),
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	_, err = ArgvIterFromSlice([]string{"-t"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.ContainsError(UnrecognizedShortArgErr, err, t)

	_, err = ArgvIterFromSlice([]string{"-t", "123"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.ContainsError(UnrecognizedShortArgErr, err, t)
}

func TestToArgValPairsUnrecognizedLongFlag(t *testing.T) {
	res := struct{ S string }{}

	b := ArgBuilder{}
	AddArg[translators.BuiltinString](
		&res.S,
		&b,
		"str",
		NewOpts[translators.BuiltinString]().
			SetShortName('s'),
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	_, err = ArgvIterFromSlice([]string{"--time"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.ContainsError(UnrecognizedLongArgErr, err, t)

	_, err = ArgvIterFromSlice([]string{"--time", "123"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.ContainsError(UnrecognizedLongArgErr, err, t)

	_, err = ArgvIterFromSlice([]string{"--time=123"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.ContainsError(UnrecognizedLongArgErr, err, t)
}

func TestToArgValPairsInvalidArgType(t *testing.T) {
	res := struct{ S string }{}

	b := ArgBuilder{}
	AddArg[translators.BuiltinString](
		&res.S,
		&b,
		"str",
		NewOpts[translators.BuiltinString]().
			SetArgType(ArgType(-1)).
			SetShortName('s'),
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	_, err = ArgvIterFromSlice([]string{"--str"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.ContainsError(InvalidArgType, err, t)
}

func TestToArgValPairsEndOfTokenStream(t *testing.T) {
	res := struct{ S string }{}

	b := ArgBuilder{}
	AddArg[translators.BuiltinString](
		&res.S,
		&b,
		"str",
		NewOpts[translators.BuiltinString]().
			SetShortName('s'),
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	_, err = ArgvIterFromSlice([]string{"--str"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.ContainsError(ExpectedValueErr, err, t)
	test.ContainsError(EndOfTokenStreamErr, err, t)
}

func TestToArgValPairsExpectedValue(t *testing.T) {
	res := struct{ S string }{}

	b := ArgBuilder{}
	AddArg[translators.BuiltinString](
		&res.S,
		&b,
		"str",
		NewOpts[translators.BuiltinString]().
			SetShortName('s'),
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	_, err = ArgvIterFromSlice([]string{"--str", "--str"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.ContainsError(ExpectedValueErr, err, t)

	_, err = ArgvIterFromSlice([]string{"--str", "--time"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.ContainsError(ExpectedValueErr, err, t)

	_, err = ArgvIterFromSlice([]string{"--str", "-s"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.ContainsError(ExpectedValueErr, err, t)

	_, err = ArgvIterFromSlice([]string{"--str", "-t"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.ContainsError(ExpectedValueErr, err, t)
}

func TestToArgValPairsPassingShortFlag(t *testing.T) {
	res := struct{ S string }{}

	b := ArgBuilder{}
	AddArg[translators.BuiltinString](
		&res.S,
		&b,
		"str",
		NewOpts[translators.BuiltinString]().
			SetShortName('s'),
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	pairs, err := ArgvIterFromSlice([]string{"-s", "123"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.Nil(err, t)
	test.Eq(len(pairs), 1, t)
	test.Eq(pairs[0].A.longFlag, "str", t)
	test.Eq(pairs[0].A.argType, ValueArgType, t)
	test.Eq(pairs[0].B, "123", t)

	pairs, err = ArgvIterFromSlice([]string{"-s=123"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.Nil(err, t)
	test.Eq(len(pairs), 1, t)
	test.Eq(pairs[0].A.longFlag, "str", t)
	test.Eq(pairs[0].A.argType, ValueArgType, t)
	test.Eq(pairs[0].B, "123", t)
}

func TestToArgValPairsPassingLongFlag(t *testing.T) {
	res := struct{ S string }{}

	b := ArgBuilder{}
	AddArg[translators.BuiltinString](
		&res.S,
		&b,
		"str",
		NewOpts[translators.BuiltinString]().
			SetShortName('s'),
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	pairs, err := ArgvIterFromSlice([]string{"--str", "123"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.Nil(err, t)
	test.Eq(len(pairs), 1, t)
	test.Eq(pairs[0].A.longFlag, "str", t)
	test.Eq(pairs[0].A.argType, ValueArgType, t)
	test.Eq(pairs[0].B, "123", t)

	pairs, err = ArgvIterFromSlice([]string{"--str=123"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.Nil(err, t)
	test.Eq(len(pairs), 1, t)
	test.Eq(pairs[0].A.longFlag, "str", t)
	test.Eq(pairs[0].A.argType, ValueArgType, t)
	test.Eq(pairs[0].B, "123", t)
}

func TestToArgValPairsArgumentOnFlag(t *testing.T) {
	res := struct{ B bool }{}

	b := ArgBuilder{}
	AddFlag(
		&res.B,
		&b,
		"bool",
		NewOpts[translators.Flag]().SetShortName('b'),
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	_, err = ArgvIterFromSlice([]string{"--bool=123"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.ContainsError(ExpectedArgumentErr, err, t)

	_, err = ArgvIterFromSlice([]string{"-b=123"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.ContainsError(ExpectedArgumentErr, err, t)
}

func TestToArgValPairsArgumentOnFlagCntr(t *testing.T) {
	res := struct{ I int }{}

	b := ArgBuilder{}
	AddFlagCntr[int](
		&res.I,
		&b,
		"int",
		NewOpts[*translators.FlagCntr[int]]().SetShortName('i'),
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	_, err = ArgvIterFromSlice([]string{"--int=123"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.ContainsError(ExpectedArgumentErr, err, t)

	_, err = ArgvIterFromSlice([]string{"-i=123"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.ContainsError(ExpectedArgumentErr, err, t)
}

func TestToArgValPairsCombinedValueArgsPassing(t *testing.T) {
	res := struct {
		I int
		B bool
	}{}

	b := ArgBuilder{}
	AddFlag(
		&res.B,
		&b,
		"bool",
		NewOpts[translators.Flag]().SetShortName('b'),
	)
	AddFlagCntr[int](
		&res.I,
		&b,
		"int",
		NewOpts[*translators.FlagCntr[int]]().SetShortName('i'),
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	pairs, err := ArgvIterFromSlice([]string{"-iibib"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.Nil(err, t)
	test.Eq(len(pairs), 5, t)
	test.Eq(pairs[0].A.longFlag, "int", t)
	test.Eq(pairs[1].A.longFlag, "int", t)
	test.Eq(pairs[2].A.longFlag, "bool", t)
	test.Eq(pairs[3].A.longFlag, "int", t)
	test.Eq(pairs[4].A.longFlag, "bool", t)
}

func TestToArgValPairsCombinedValueArgsPassingWithLongArg(t *testing.T) {
	res := struct {
		I int
		B bool
	}{}

	b := ArgBuilder{}
	AddFlag(
		&res.B,
		&b,
		"bool",
		NewOpts[translators.Flag]().SetShortName('b'),
	)
	AddFlagCntr[int](
		&res.I,
		&b,
		"int",
		NewOpts[*translators.FlagCntr[int]]().SetShortName('i'),
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	pairs, err := ArgvIterFromSlice([]string{"-iibi", "--int"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.Nil(err, t)
	test.Eq(len(pairs), 5, t)
	test.Eq(pairs[0].A.longFlag, "int", t)
	test.Eq(pairs[1].A.longFlag, "int", t)
	test.Eq(pairs[2].A.longFlag, "bool", t)
	test.Eq(pairs[3].A.longFlag, "int", t)
	test.Eq(pairs[4].A.longFlag, "int", t)
}

func TestToArgValPairsMultiArgPassing(t *testing.T) {
	res := struct{ S []int }{}

	b := ArgBuilder{}
	AddListArg[translators.BuiltinInt, widgets.BuiltinInt](
		&res.S,
		&b,
		"list",
		NewOpts[*translators.ListValues[translators.BuiltinInt, widgets.BuiltinInt, int]]().
			SetShortName('l').
			SetTranslator(&translators.ListValues[
				translators.BuiltinInt,
				widgets.BuiltinInt,
				int,
			]{
				ValueTranslator: translators.BuiltinInt{Base: 10},
			}),
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	pairs, err := ArgvIterFromSlice([]string{"-l=1", "2", "3"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.Nil(err, t)
	test.Eq(len(pairs), 3, t)
	test.Eq(pairs[0].A.longFlag, "list", t)
	test.Eq(pairs[0].B, "1", t)
	test.Eq(pairs[1].A.longFlag, "list", t)
	test.Eq(pairs[1].B, "2", t)
	test.Eq(pairs[2].A.longFlag, "list", t)
	test.Eq(pairs[2].B, "3", t)

	pairs, err = ArgvIterFromSlice([]string{"-l", "1", "2", "3"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.Nil(err, t)
	test.Eq(len(pairs), 3, t)
	test.Eq(pairs[0].A.longFlag, "list", t)
	test.Eq(pairs[0].B, "1", t)
	test.Eq(pairs[1].A.longFlag, "list", t)
	test.Eq(pairs[1].B, "2", t)
	test.Eq(pairs[2].A.longFlag, "list", t)
	test.Eq(pairs[2].B, "3", t)

	pairs, err = ArgvIterFromSlice([]string{"-list=1", "2", "3"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.Nil(err, t)
	test.Eq(len(pairs), 3, t)
	test.Eq(pairs[0].A.longFlag, "list", t)
	test.Eq(pairs[0].B, "1", t)
	test.Eq(pairs[1].A.longFlag, "list", t)
	test.Eq(pairs[1].B, "2", t)
	test.Eq(pairs[2].A.longFlag, "list", t)
	test.Eq(pairs[2].B, "3", t)

	pairs, err = ArgvIterFromSlice([]string{"--list", "1", "2", "3"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.Nil(err, t)
	test.Eq(len(pairs), 3, t)
	test.Eq(pairs[0].A.longFlag, "list", t)
	test.Eq(pairs[0].B, "1", t)
	test.Eq(pairs[1].A.longFlag, "list", t)
	test.Eq(pairs[1].B, "2", t)
	test.Eq(pairs[2].A.longFlag, "list", t)
	test.Eq(pairs[2].B, "3", t)
}

func TestToArgValPairsMultiArgMissingValue(t *testing.T) {
	res := struct{ S []int }{}

	b := ArgBuilder{}
	AddListArg[translators.BuiltinInt, widgets.BuiltinInt](
		&res.S,
		&b,
		"list",
		NewOpts[*translators.ListValues[translators.BuiltinInt, widgets.BuiltinInt, int]]().
			SetShortName('l').
			SetTranslator(&translators.ListValues[
				translators.BuiltinInt,
				widgets.BuiltinInt,
				int,
			]{
				ValueTranslator: translators.BuiltinInt{Base: 10},
			}),
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	_, err = ArgvIterFromSlice([]string{"-list", "-t"}).
		ToTokens().
		toArgValPairs(&p).
		ToIter().Collect()
	test.ContainsError(ExpectedValueErr, err, t)
}
