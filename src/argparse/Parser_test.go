package argparse

import (
	"errors"
	"fmt"
	"testing"

	"github.com/barbell-math/util/src/argparse/computers"
	"github.com/barbell-math/util/src/argparse/translators"
	"github.com/barbell-math/util/src/container/containerTypes"
	"github.com/barbell-math/util/src/test"
)

func TestParserAddSubParsersEmptyParser(t *testing.T) {
	res := struct{ S string }{}

	b := ArgBuilder{}
	AddArg[translators.BuiltinString](
		&res.S,
		&b,
		"str",
		NewOpts[translators.BuiltinString]().SetShortName('s'),
	)
	p1, err := b.ToParser("", "")
	test.Nil(err, t)

	p2, err := (&ArgBuilder{}).ToParser("", "")

	err = p1.AddSubParsers(p2)
	test.Nil(err, t)
	test.Eq(len(p1.subParsers), 1, t)
	test.Eq(len(p1.subParsers[0]), 1, t)
	test.Eq(p1.subParsers[0][0].longFlag, "str", t)
}

func TestParserAddSubParsersNonEmptyValid(t *testing.T) {
	res := struct {
		S string
		I int
	}{}

	b := ArgBuilder{}
	AddArg[translators.BuiltinString](
		&res.S,
		&b,
		"str",
		NewOpts[translators.BuiltinString]().SetShortName('s'),
	)
	p1, err := b.ToParser("", "")
	test.Nil(err, t)

	err = p1.AddSubParsers(NewHelpParser())
	test.Nil(err, t)
	test.Eq(len(p1.subParsers), 2, t)
	test.Eq(len(p1.subParsers[0]), 1, t)
	test.Eq(len(p1.subParsers[1]), 1, t)
	test.Eq(p1.subParsers[0][0].longFlag, "str", t)
	test.Eq(p1.subParsers[1][0].longFlag, "help", t)
}

func TestParserAddSubParsersNonEmptyDuplicateLongNames(t *testing.T) {
	res := struct {
		S string
		I int
	}{}

	b := ArgBuilder{}
	AddArg[translators.BuiltinString](
		&res.S,
		&b,
		"str",
		NewOpts[translators.BuiltinString]().SetShortName('s'),
	)
	p1, err := b.ToParser("", "")
	test.Nil(err, t)

	b2 := ArgBuilder{}
	AddArg[translators.BuiltinString](
		&res.S,
		&b2,
		"str",
		NewOpts[translators.BuiltinString]().SetShortName('S'),
	)
	p2, err := b2.ToParser("", "")
	test.Nil(err, t)

	err = p1.AddSubParsers(p2)
	test.ContainsError(ParserCombinationErr, err, t)
	test.ContainsError(DuplicateLongNameErr, err, t)
	test.ContainsError(containerTypes.Duplicate, err, t)

	fmt.Println(err)
}

func TestParserAddSubParsersNonEmptyDuplicateShortNames(t *testing.T) {
	res := struct {
		S string
		I int
	}{}

	b := ArgBuilder{}
	AddArg[translators.BuiltinString](
		&res.S,
		&b,
		"str",
		NewOpts[translators.BuiltinString]().SetShortName('s'),
	)
	p1, err := b.ToParser("", "")
	test.Nil(err, t)

	b2 := ArgBuilder{}
	AddArg[translators.BuiltinString](
		&res.S,
		&b2,
		"str2",
		NewOpts[translators.BuiltinString]().SetShortName('s'),
	)
	p2, err := b2.ToParser("", "")
	test.Nil(err, t)

	err = p1.AddSubParsers(p2)
	test.ContainsError(ParserCombinationErr, err, t)
	test.ContainsError(DuplicateShortNameErr, err, t)
	test.ContainsError(containerTypes.Duplicate, err, t)
}

func TestParserAddSubParsers(t *testing.T) {
	res := struct {
		A int
		B int
		C int
		D int
		E int
		F int
		G int
	}{}

	b1 := ArgBuilder{}
	AddArg[translators.BuiltinInt](
		&res.A,
		&b1,
		"aa",
		NewOpts[translators.BuiltinInt]().
			SetShortName('a').
			SetRequired(true),
	)
	p1, err := b1.ToParser("", "")
	test.Nil(err, t)

	b2 := ArgBuilder{}
	AddArg[translators.BuiltinInt](
		&res.B,
		&b2,
		"bb",
		NewOpts[translators.BuiltinInt]().
			SetShortName('b').
			SetRequired(true),
	)
	p2, err := b2.ToParser("", "")
	test.Nil(err, t)

	b3 := ArgBuilder{}
	AddArg[translators.BuiltinInt](
		&res.C,
		&b3,
		"cc",
		NewOpts[translators.BuiltinInt]().
			SetShortName('c').
			SetRequired(true),
	)
	p3, err := b3.ToParser("", "")
	test.Nil(err, t)

	b4 := ArgBuilder{}
	AddArg[translators.BuiltinInt](
		&res.D,
		&b4,
		"dd",
		NewOpts[translators.BuiltinInt]().
			SetShortName('d').
			SetRequired(true),
	)
	p4, err := b4.ToParser("", "")
	test.Nil(err, t)

	b5 := ArgBuilder{}
	AddComputedArg[computers.Add[int]](
		&res.E,
		&b5,
		computers.Add[int]{L: &res.A, R: &res.B},
	)
	p5, err := b5.ToParser("", "")
	test.Nil(err, t)

	b6 := ArgBuilder{}
	AddComputedArg[computers.Sub[int]](
		&res.F,
		&b6,
		computers.Sub[int]{L: &res.C, R: &res.D},
	)
	p6, err := b6.ToParser("", "")
	test.Nil(err, t)

	b7 := ArgBuilder{}
	AddComputedArg[computers.Mul[int]](
		&res.G,
		&b7,
		computers.Mul[int]{L: &res.E, R: &res.F},
	)
	p7, err := b7.ToParser("", "")
	test.Nil(err, t)

	p5.AddSubParsers(p1, p2)
	p6.AddSubParsers(p3, p4)
	p7.AddSubParsers(p5, p6)

	longKeys, err := p7.longArgs.Keys().Collect()
	test.Nil(err, t)
	test.SlicesMatchUnordered[string](
		longKeys,
		[]string{"aa", "bb", "cc", "dd"},
		t,
	)

	shortKeys, err := p7.shortArgs.Keys().Collect()
	test.Nil(err, t)
	test.SlicesMatchUnordered[byte](
		shortKeys,
		[]byte{'a', 'b', 'c', 'd'},
		t,
	)
}

func TestParserParseDuplicateValArg(t *testing.T) {
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

	err = p.Parse(ArgvIterFromSlice([]string{"--str=123", "-s=456"}).ToTokens())
	test.ContainsError(ParsingErr, err, t)
	test.ContainsError(ArgumentPassedMultipleTimesErr, err, t)
}

func TestParserParseDuplicateFlagArg(t *testing.T) {
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

	err = p.Parse(ArgvIterFromSlice([]string{"--bool", "-b"}).ToTokens())
	test.ContainsError(ParsingErr, err, t)
	test.ContainsError(ArgumentPassedMultipleTimesErr, err, t)
}

func TestParserParseMissingRequiredArgs(t *testing.T) {
	res := struct{ B bool }{}

	b := ArgBuilder{}
	AddFlag(
		&res.B,
		&b,
		"bool",
		NewOpts[translators.Flag]().SetShortName('b').SetRequired(true),
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	err = p.Parse(ArgvIterFromSlice([]string{}).ToTokens())
	test.ContainsError(ParsingErr, err, t)
	test.ContainsError(MissingRequiredArgErr, err, t)
}

func TestParserParseDefaultValue(t *testing.T) {
	res := struct{ S string }{}

	b := ArgBuilder{}
	AddArg[translators.BuiltinString](
		&res.S,
		&b,
		"str",
		NewOpts[translators.BuiltinString]().
			SetShortName('s').
			SetDefaultVal("default"),
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	err = p.Parse(ArgvIterFromSlice([]string{}).ToTokens())
	test.Nil(err, t)
	test.Eq(res.S, "default", t)
}

func TestParserParseWithComputedArgumentErr(t *testing.T) {
	res := struct {
		L   int
		R   int
		Res int
	}{}

	b := ArgBuilder{}
	AddArg[translators.BuiltinInt](
		&res.L,
		&b,
		"left",
		NewOpts[translators.BuiltinInt]().
			SetShortName('l').
			SetRequired(true),
	)
	AddArg[translators.BuiltinInt](
		&res.R,
		&b,
		"right",
		NewOpts[translators.BuiltinInt]().
			SetShortName('r').
			SetRequired(true),
	)
	AddComputedArg[computers.Stopper[int]](
		&res.Res,
		&b,
		computers.Stopper[int]{Err: errors.New("ERROR")},
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	err = p.Parse(ArgvIterFromSlice([]string{"-l=3", "-r=5"}).ToTokens())
	test.ContainsError(ParsingErr, err, t)
	test.ContainsError(ComputedArgumentErr, err, t)
	test.Eq(res.L, 3, t)
	test.Eq(res.R, 5, t)
	test.Eq(res.Res, 0, t)
}

func TestParserParseWithComputedArgsSimple(t *testing.T) {
	res := struct {
		L   int
		R   int
		Res int
	}{}

	b := ArgBuilder{}
	AddArg[translators.BuiltinInt](
		&res.L,
		&b,
		"left",
		NewOpts[translators.BuiltinInt]().
			SetShortName('l').
			SetRequired(true),
	)
	AddArg[translators.BuiltinInt](
		&res.R,
		&b,
		"right",
		NewOpts[translators.BuiltinInt]().
			SetShortName('r').
			SetRequired(true),
	)
	AddComputedArg[computers.Add[int]](
		&res.Res,
		&b,
		computers.Add[int]{L: &res.L, R: &res.R},
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	err = p.Parse(ArgvIterFromSlice([]string{"-l=3", "-r=5"}).ToTokens())
	test.Eq(res.Res, 8, t)
}

func TestParserParseWithComputedArgsComplex(t *testing.T) {
	res := struct {
		A int
		B int
		C int
		D int
		E int
		F int
		G int
	}{}

	b1 := ArgBuilder{}
	AddArg[translators.BuiltinInt](
		&res.A,
		&b1,
		"aa",
		NewOpts[translators.BuiltinInt]().
			SetShortName('a').
			SetRequired(true),
	)
	p1, err := b1.ToParser("", "")
	test.Nil(err, t)

	b2 := ArgBuilder{}
	AddArg[translators.BuiltinInt](
		&res.B,
		&b2,
		"bb",
		NewOpts[translators.BuiltinInt]().
			SetShortName('b').
			SetRequired(true),
	)
	p2, err := b2.ToParser("", "")
	test.Nil(err, t)

	b3 := ArgBuilder{}
	AddArg[translators.BuiltinInt](
		&res.C,
		&b3,
		"cc",
		NewOpts[translators.BuiltinInt]().
			SetShortName('c').
			SetRequired(true),
	)
	p3, err := b3.ToParser("", "")
	test.Nil(err, t)

	b4 := ArgBuilder{}
	AddArg[translators.BuiltinInt](
		&res.D,
		&b4,
		"dd",
		NewOpts[translators.BuiltinInt]().
			SetShortName('d').
			SetRequired(true),
	)
	p4, err := b4.ToParser("", "")
	test.Nil(err, t)

	b5 := ArgBuilder{}
	AddComputedArg[computers.Add[int]](
		&res.E,
		&b5,
		computers.Add[int]{L: &res.A, R: &res.B},
	)
	p5, err := b5.ToParser("", "")
	test.Nil(err, t)

	b6 := ArgBuilder{}
	AddComputedArg[computers.Sub[int]](
		&res.F,
		&b6,
		computers.Sub[int]{L: &res.C, R: &res.D},
	)
	p6, err := b6.ToParser("", "")
	test.Nil(err, t)

	b7 := ArgBuilder{}
	AddComputedArg[computers.Mul[int]](
		&res.G,
		&b7,
		computers.Mul[int]{L: &res.E, R: &res.F},
	)
	p7, err := b7.ToParser("", "")
	test.Nil(err, t)

	p5.AddSubParsers(p1, p2)
	p6.AddSubParsers(p3, p4)
	p7.AddSubParsers(p5, p6)

	err = p7.Parse(ArgvIterFromSlice([]string{
		"-a=3",
		"-b=5",
		"-c=7",
		"-d=9",
	}).ToTokens())
	test.Nil(err, t)
	test.Eq(res.A, 3, t)
	test.Eq(res.B, 5, t)
	test.Eq(res.C, 7, t)
	test.Eq(res.D, 9, t)
	test.Eq(res.E, 8, t)
	test.Eq(res.F, -2, t)
	test.Eq(res.G, -16, t)
}

func TestParserParseConditionallyRequiredArguments(t *testing.T) {
	res := struct {
		S1 string
		S2 string
	}{}

	b := ArgBuilder{}
	AddArg[translators.BuiltinString](
		&res.S1,
		&b,
		"str1",
		NewOpts[translators.BuiltinString]().
			SetConditionallyRequired([]ArgConditionality[string]{
				ArgConditionality[string]{
					Requires: []string{"str2"},
					When:     ArgEquals[string]("foo"),
				},
			}),
	)
	AddArg[translators.BuiltinString](
		&res.S2,
		&b,
		"str2",
		NewOpts[translators.BuiltinString]().SetDefaultVal("foobar"),
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	err = p.Parse(ArgvIterFromSlice([]string{"--str2=bar"}).ToTokens())
	test.Nil(err, t)
	test.Eq("", res.S1, t)
	test.Eq("bar", res.S2, t)

	err = p.Parse(ArgvIterFromSlice([]string{"--str1=bar"}).ToTokens())
	test.Nil(err, t)
	test.Eq("bar", res.S1, t)
	test.Eq("foobar", res.S2, t)

	res.S1 = ""
	res.S2 = ""
	err = p.Parse(ArgvIterFromSlice([]string{"--str1=foo"}).ToTokens())
	test.ContainsError(ParsingErr, err, t)
	test.ContainsError(MissingConditionallyRequiredArgErr, err, t)

	res.S1 = ""
	res.S2 = ""
	err = p.Parse(ArgvIterFromSlice([]string{"--str1=foo", "--str2=asdf"}).ToTokens())
	test.Nil(err, t)
	test.Eq("foo", res.S1, t)
	test.Eq("asdf", res.S2, t)
}

func TestParserParseUnrecognizedConditionallyRequiredArg(t *testing.T) {
	res := ""
	b := ArgBuilder{}

	AddArg[translators.BuiltinString](
		&res,
		&b,
		"str",
		NewOpts[translators.BuiltinString]().
			SetShortName('t').
			SetConditionallyRequired([]ArgConditionality[string]{
				ArgConditionality[string]{
					Requires: []string{"foo"},
					When:     ArgSupplied[string],
				},
			}),
	)

	p, err := b.ToParser("", "")
	test.Nil(err, t)

	err = p.Parse(ArgvIterFromSlice([]string{}).ToTokens())
	test.ContainsError(ParserConfigErr, err, t)
	test.ContainsError(UnrecognizedConditionallyRequiredArgErr, err, t)
}

func TestParserParseConditionallyRequiredArgumentsWithDefaults(t *testing.T) {
	res := struct {
		S1 string
		S2 string
	}{}

	b := ArgBuilder{}
	AddArg[translators.BuiltinString](
		&res.S1,
		&b,
		"str1",
		NewOpts[translators.BuiltinString]().
			SetDefaultVal("string1").
			SetConditionallyRequired([]ArgConditionality[string]{
				ArgConditionality[string]{
					Requires: []string{"str2"},
					When:     ArgSupplied[string],
				},
			}),
	)
	AddArg[translators.BuiltinString](
		&res.S2,
		&b,
		"str2",
		NewOpts[translators.BuiltinString]().SetDefaultVal("string2"),
	)
	p, err := b.ToParser("", "")
	test.Nil(err, t)

	// Reaches err because default value was provided.
	err = p.Parse(ArgvIterFromSlice([]string{}).ToTokens())
	test.ContainsError(ParsingErr, err, t)
	test.ContainsError(MissingConditionallyRequiredArgErr, err, t)

	err = p.Parse(ArgvIterFromSlice([]string{"--str2=bar"}).ToTokens())
	test.Nil(err, t)
	test.Eq("string1", res.S1, t)
	test.Eq("bar", res.S2, t)

	err = p.Parse(ArgvIterFromSlice([]string{"--str1=bar"}).ToTokens())
	test.ContainsError(ParsingErr, err, t)
	test.ContainsError(MissingConditionallyRequiredArgErr, err, t)

	res.S1 = ""
	res.S2 = ""
	err = p.Parse(ArgvIterFromSlice([]string{"--str1=foo", "--str2=asdf"}).ToTokens())
	test.Nil(err, t)
	test.Eq("foo", res.S1, t)
	test.Eq("asdf", res.S2, t)
}
