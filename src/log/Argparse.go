package log

import (
	"os"

	"github.com/barbell-math/util/src/argparse"
	"github.com/barbell-math/util/src/argparse/translators"
	"github.com/barbell-math/util/src/math/basic"
)

// Creates a parser that has the --logFile flag. This flag will be used to
// specify a single log file. The argument parser will return an open file
// handle to the file for future writing.
func NewSingleLogFileParser(f *os.File) *argparse.Parser {
	b := argparse.ArgBuilder{}
	argparse.AddArg[*os.File, *translators.OpenFile](
		&f, &b, "logFile",
		argparse.NewOpts[*os.File, *translators.OpenFile]().
			SetDescription("The file to send all log messages to").
			SetTranslator(
				translators.NewOpenFile().
					SetFlags(os.O_CREATE).
					SetPermissions(0777),
			),
	)
	rv, _ := b.ToParser("", "")
	return &rv
}

// Creates a parser that has -v and --verbose flags. These flags can be
// supplied many times and the total count of the number of times the argument
// was supplied will be placed in val.
func NewVerbosityParser[T basic.Int | basic.Uint](val *T) *argparse.Parser {
	b := argparse.ArgBuilder{}
	argparse.AddFlagCntr[T](
		val,
		&b,
		"verbose",
		argparse.NewOpts[T, *translators.FlagCntr[T]]().
			SetArgType(argparse.MultiFlagArgType).
			SetShortName('v').
			SetRequired(false).
			SetDescription("Sets the verbosity level. Specify multiple times to increase.").
			SetTranslator(&translators.FlagCntr[T]{}),
	)
	rv, _ := b.ToParser("", "")
	return &rv
}
