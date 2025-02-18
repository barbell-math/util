package translators

import (
	stdtime "time"
)

type (
	time struct {
		Format string
	}

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
