package log

import (
	"os"

	"github.com/barbell-math/util/src/argparse"
	"github.com/barbell-math/util/src/argparse/translators"
	"github.com/barbell-math/util/src/math/basic"
)

type (
	VerbosityTranslator[T basic.Int | basic.Uint] struct {
		translators.FlagCntr[T]
		MaxCnt T
	}
)

const (
	LogFileArg string = "logFile"
	VerboseArg string = "verbose"
)

// Creates a parser that has the --logFile flag. This flag will be used to
// specify a single log file. The argument parser will return an open file
// handle to the file for future writing.
func NewSingleLogFileParser(f *os.File) argparse.Parser {
	b := argparse.ArgBuilder{}
	argparse.AddArg[*translators.OpenFile](
		&f, &b, LogFileArg,
		argparse.NewOpts[*translators.OpenFile]().
			SetDescription("The file to send all log messages to").
			SetTranslator(
				translators.NewOpenFile().
					SetFlags(os.O_CREATE).
					SetPermissions(0777),
			),
	)
	rv, err := b.ToParser("", "")
	if err != nil {
		panic(err)
	}
	return rv
}

// Creates a parser that has -v and --verbose flags. These flags can be
// supplied many times and the total count of the number of times the argument
// was supplied will be placed in val.
func NewVerbosityParser[T basic.Int | basic.Uint](
	val *T,
	defaultVal T,
	maxLevel T,
) argparse.Parser {
	b := argparse.ArgBuilder{}
	argparse.AddArg[*translators.LimitedFlagCntr[T]](
		val, &b, VerboseArg,
		argparse.NewOpts[*translators.LimitedFlagCntr[T]]().
			SetArgType(argparse.MultiFlagArgType).
			SetShortName('v').
			SetRequired(false).
			SetDefaultVal(defaultVal).
			SetDescription("Sets the verbosity level. Specify multiple times to increase.").
			SetTranslator(&translators.LimitedFlagCntr[T]{MaxTimes: maxLevel}),
	)
	rv, err := b.ToParser("", "")
	if err != nil {
		panic(err)
	}
	return rv
}
