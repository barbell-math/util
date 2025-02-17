package appactions

import (
	"github.com/barbell-math/util/src/argparse"
	"github.com/barbell-math/util/src/argparse/translators"
	"github.com/barbell-math/util/src/enum"
)

const (
	ActionArg string = "action"
)

// Creates a parser that has -a and --action falgs. These flags will be used to
// set an enum value that should be use to tell your application what action it
// needs to perform.
func NewAppActionParser[EP enum.Pntr[E], E enum.Value](
	val *E,
	defaultVal E,
	conditionalArgs []argparse.ArgConditionality[E],
) argparse.Parser {
	b := argparse.ArgBuilder{}
	argparse.AddEnum(
		val, &b, ActionArg,
		argparse.NewOpts[translators.Enum[EP, E]]().
			SetArgType(argparse.ValueArgType).
			SetShortName('a').
			SetRequired(false).
			SetDefaultVal(defaultVal).
			SetDescription("The action that the application should perform.").
			SetTranslator(translators.Enum[EP, E]{}).
			SetConditionallyRequired(conditionalArgs),
	)
	rv, err := b.ToParser("", "")
	if err != nil {
		panic(err)
	}
	return rv
}
