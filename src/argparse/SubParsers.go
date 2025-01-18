package argparse

import (
	"github.com/barbell-math/util/src/argparse/translators"
	"github.com/barbell-math/util/src/enum"
	"github.com/barbell-math/util/src/math/basic"
)

// Creates a parser that will display the help menu when either -h or --help are
// supplied.
func NewHelpParser() *Parser {
	res := struct{}{}
	b := ArgBuilder{}
	AddArg[struct{}, translators.Stopper[struct{}]](
		&res,
		&b,
		"help",
		NewOpts[struct{}, translators.Stopper[struct{}]]().
			SetArgType(FlagArgType).
			SetShortName('h').
			SetRequired(false).
			SetDescription("Prints this help menu.").
			SetTranslator(translators.Stopper[struct{}]{Err: HelpErr}),
	)
	rv, _ := b.ToParser("", "")
	return &rv
}

// Creates a parser that has -v and --verbose flags. These flags can be
// supplied many times and the total count of the number of times the argument
// was supplied will be placed in val.
func NewVerbosityParser[T basic.Int | basic.Uint](val *T) *Parser {
	b := ArgBuilder{}
	AddFlagCntr[T](
		val,
		&b,
		"verbose",
		NewOpts[T, *translators.FlagCntr[T]]().
			SetArgType(MultiFlagArgType).
			SetShortName('v').
			SetRequired(false).
			SetDescription("Sets the verbosity level. Specify multiple times to increase.").
			SetTranslator(&translators.FlagCntr[T]{}),
	)
	rv, _ := b.ToParser("", "")
	return &rv
}

// Creates a parser that has -a and --action falgs. These flags will be used to
// set an enum value that should be use to tell your application what action it
// needs to perform.
func NewAppActionParser[E enum.Value, EP enum.Pntr[E]](val *E) *Parser {
	b := ArgBuilder{}
	AddEnum(
		val,
		&b,
		"action",
		NewOpts[E, translators.Enum[E, EP]]().
			SetArgType(ValueArgType).
			SetShortName('a').
			SetRequired(true).
			SetDescription("The action that the application should perform.").
			SetTranslator(translators.Enum[E, EP]{}),
	)
	rv, _ := b.ToParser("", "")
	return &rv
}
