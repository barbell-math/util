package log

import (
	"fmt"
	"time"
)

type (
	LogStatus int

	LogEntry[T any] struct {
		Status  LogStatus
		Time    time.Time
		Message string
		Val     T
	}
)

const (
	// The separator character that is used to delineate different values in the
	// log
	LogPartSeparator string = "|"

	Error LogStatus = iota
	Warning
	Deprecation
	Info
	Debug
	Invalid
)

func (l LogStatus) String() string {
	switch l {
	case Error:
		return fmt.Sprintf("Error %s ", LogPartSeparator)
	case Warning:
		return fmt.Sprintf("Warning %s ", LogPartSeparator)
	case Deprecation:
		return fmt.Sprintf("Deprecation %s ", LogPartSeparator)
	case Info:
		return fmt.Sprintf("Info %s ", LogPartSeparator)
	case Debug:
		return fmt.Sprintf("Debug %s ", LogPartSeparator)
	default:
		return fmt.Sprintf("Invalid %s ", LogPartSeparator)
	}
}
func LogStatusFromString(s string) (LogStatus, error) {
	switch s {
	case "Error":
		return Error, nil
	case "Warning":
		return Warning, nil
	case "Deprecation":
		return Deprecation, nil
	case "Info":
		return Info, nil
	case "Debug":
		return Debug, nil
	default:
		return Invalid, InvalidLogStatus
	}
}

// A function that can be passed to [iter.Join] to allow it to sort the values
// from two logs, providing each log statement in the order that they were
// written.
func JoinLogByTimeInc[T any, U any](left LogEntry[T], right LogEntry[U]) bool {
	return left.Time.Before(right.Time)
}
