package translators

import (
	stdtime "time"
)

//go:generate ../../../bin/ifaceImplCheck -typeToCheck=time
//go:generate ../../../bin/ifaceImplCheck -typeToCheck=Duration

type (
	// Represents a cmd line argument that will be parsed to a time value with
	// a date and time component. Uses [stdtime.Parse] internally.
	//gen:ifaceImplCheck ifaceName Translator[stdTime.Time]
	//gen:ifaceImplCheck imports stdTime->time
	//gen:ifaceImplCheck valOrPntr both
	time struct {
		Format string
	}

	// Represents a cmd line argument that will be parsed as a duration. All
	// the rules that [stdtime.ParseDuration] use will be used here.
	//gen:ifaceImplCheck ifaceName Translator[stdTime.Duration]
	//gen:ifaceImplCheck imports stdTime->time
	//gen:ifaceImplCheck valOrPntr both
	Duration struct{}
)

func NewTime(format string) time {
	return time{Format: format}
}

func (t time) Translate(arg string) (stdtime.Time, error) {
	return stdtime.Parse(t.Format, arg)
}

func (_ time) Reset() {
	// intentional noop - time has no state that needs to be reset
}

func (_ Duration) Translate(arg string) (stdtime.Duration, error) {
	return stdtime.ParseDuration(arg)
}

func (_ Duration) Reset() {
	// intentional noop - time has no state that needs to be reset
}
