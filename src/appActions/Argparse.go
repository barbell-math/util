package appactions

import (
	"github.com/barbell-math/util/src/argparse"
	"github.com/barbell-math/util/src/argparse/translators"
	"github.com/barbell-math/util/src/enum"
)

// Creates a parser that has -a and --action falgs. These flags will be used to
// set an enum value that should be use to tell your application what action it
// needs to perform.
func NewAppActionParser[E enum.Value, EP enum.Pntr[E]](
	val *E,
	defaultVal E,
	conditionalArgs []argparse.ArgConditionality[E],
) argparse.Parser {
	b := argparse.ArgBuilder{}
	argparse.AddEnum(
		val, &b, "action",
		argparse.NewOpts[E, translators.Enum[E, EP]]().
			SetArgType(argparse.ValueArgType).
			SetShortName('a').
			SetRequired(false).
			SetDefaultVal(defaultVal).
			SetDescription("The action that the application should perform.").
			SetTranslator(translators.Enum[E, EP]{}).
			SetConditionallyRequired(conditionalArgs),
	)
	rv, err := b.ToParser("", "")
	if err != nil {
		panic(err)
	}
	return rv
}
