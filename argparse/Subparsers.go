package argparse

import (
	"github.com/barbell-math/util/argparse/translators"
	"github.com/barbell-math/util/math/basic"
)

func NewHelpParser() *Parser {
	res:=struct{}{}
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
			SetDescription("Sets the verbosity level. Specify multiple times to increase."),
	)
	rv, _ := b.ToParser("", "")
	return &rv
}
