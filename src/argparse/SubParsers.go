package argparse

import (
	"github.com/barbell-math/util/src/argparse/translators"
)

// Creates a parser that will display the help menu when either -h or --help are
// supplied.
func NewHelpParser() Parser {
	res := struct{}{}
	b := ArgBuilder{}
	AddArg[translators.Stopper[struct{}]](
		&res,
		&b,
		"help",
		NewOpts[translators.Stopper[struct{}]]().
			SetArgType(FlagArgType).
			SetShortName('h').
			SetRequired(false).
			SetDescription("Prints this help menu.").
			SetTranslator(translators.Stopper[struct{}]{Err: HelpErr}),
	)
	rv, err := b.ToParser("", "")
	if err != nil {
		panic(err)
	}
	return rv
}
