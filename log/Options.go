package log

import "log"

type (
	optionsFlag int
	options     struct {
		flags optionsFlag

		lFlags         int
		dateTimeFormat string
	}
)

const (
	_append optionsFlag = 1 << iota
)

// Returns a new options struct initialized with the default values that can be
// passed to the other functions in this package that require options.
func NewOptions() *options {
	return &options{
		flags:          0,
		lFlags:         log.LstdFlags,
		dateTimeFormat: "2006/01/02 15:04:05",
	}
}

// Description: whether or not to append to the log file. If not appending then
// the log file will be cleared before writing.
//
// Used by: [ValueLogger]
//
// Default: false
func (o *options) Append(b bool) *options {
	if b {
		o.flags |= _append
	} else {
		o.flags &= ^_append
	}
	return o
}

// Description: the flags that should be passed to the logger. These flags are
// specified in the standard library documentation for logging.
//
// Used by: [ValueLogger]
//
// Default: log.LstdFlags
func (o *options) LogFlags(flags int) *options {
	o.lFlags = flags
	return o
}

// Description: the date time format to use when attempting to parse. No
// correctness checking is performed on the date time format string. Any errors
// from incorrect date time formats will become apparent when parsing the log
// file.
//
// Used by: [ValueLogger]
//
// Default: "2006/01/02 15:04:05", to match the log.LstdFlags option
func (o *options) DateTimeFormat(f string) *options {
	o.dateTimeFormat = f
	return o
}

func (o *options) getFlag(flag optionsFlag) bool {
	return o.flags&flag > 0
}
