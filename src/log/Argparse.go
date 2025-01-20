package log

import (
	"os"

	"github.com/barbell-math/util/src/argparse"
	"github.com/barbell-math/util/src/argparse/translators"
)

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
