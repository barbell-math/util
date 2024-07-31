package log

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/iter"
)

type (
	// Represents a logger that outputs log files in the format shown below.
	//
	//  <log status> | <date time> | <message> | <JSON formatted object>
	//
	// The message is set on a per-log-call basis. The object that is logged is
	// formatted in JSON to allow for marshalling and un-marshalling. This
	// logger should only be used when concerned with a value, or a small set of
	// values. Generic logging statements with this logger will not perform
	// optimally. For that consider using the std libs logging library.
	ValueLogger[T any] struct {
		file    string
		logFile *os.File
		logger  *log.Logger
		opts    *options
		Log     func(val T, fmt string, vals ...any)
	}
)

// Returns a new log, initialized with the supplied options. If an error occurs
// it will be returned and the logger will be in an invalid state and should not
// be used.
func NewValueLogger[T any](
	status LogStatus,
	file string,
	opts *options,
) (ValueLogger[T], error) {
	rv := ValueLogger[T]{file: file, opts: opts}
	var err error = nil
	mode := os.O_TRUNC
	if opts.getFlag(_append) {
		mode = os.O_APPEND
	}
	if rv.logFile, err = os.OpenFile(
		file,
		mode|os.O_CREATE|os.O_WRONLY,
		0644,
	); err == nil {
		rv.logger = log.New(rv.logFile, status.String(), opts.lFlags)
	} else {
		return rv, err
	}
	rv.Log = func(val T, message string, vals ...any) {
		if b, err := json.Marshal(val); err == nil && rv.logger != nil {
			rv.logger.Printf(
				"%s %s %s %s",
				LogPartSeparator,
				fmt.Sprintf(message, vals...),
				LogPartSeparator, b,
			)
		}
	}
	return rv, nil
}

// Creates a new logger object that is initialized such that it performs no
// action. The idea behind this is to allow for debugging logging to be
// conditionally turned off in production code, hopefully allowing to compiler
// to optimize away the logger entirely.
func NewBlankLog[T any]() ValueLogger[T] {
	return ValueLogger[T]{
		Log: func(val T, message string, vals ...any) {},
	}
}

// Sets the loggers status to the supplied log status.
func (l *ValueLogger[T]) SetStatus(s LogStatus) {
	l.logger.SetPrefix(s.String())
}

// Closes the logger and its associated file. Any writes to the Log method after
// calling this function will result in no action being performed.
func (l *ValueLogger[T]) Close() {
	if l.logFile != nil {
		l.logFile.Close()
		l.logFile = nil
		l.Log = func(val T, message string, vals ...any) {}
	}
}

// Clears all statements in the log and it's associated file. This will maintain
// the state of the log file, meaning if it was closed before it will be closed
// after the operation; same thing with open.
func (l *ValueLogger[T]) Clear() error {
	if len(l.file) == 0 {
		return customerr.Wrap(LogFileNotSpecified, "Nothing to clear.")
	}
	opened := false
	if l.logFile != nil {
		l.logFile.Close()
		opened = true
	}
	if logFile, err := os.OpenFile(
		l.file,
		os.O_TRUNC|os.O_CREATE|os.O_WRONLY,
		0644,
	); err == nil {
		l.logFile = logFile
	} else {
		return err
	}
	if !opened {
		l.logFile.Close()
		l.logFile = nil
	}
	return nil
}

// Retrieves the elements from the log, returning them as a stream of [LogEntry]
// structs. Any objects that were encoded in the log will be un-marshaled into
// each [LogEntry] struct allowing values to be retrieved from the log. The
// datetime format will be determined from the options that the [Logger] struct
// was initialized with.
func (l *ValueLogger[T]) LogElems() iter.Iter[LogEntry[T]] {
	var iterElem T
	return iter.Map(
		iter.FileLines(l.file),
		func(index int, val string) (LogEntry[T], error) {
			parts := strings.SplitN(val, LogPartSeparator, 4)
			s, serr := getStatus(parts)
			t, terr := getTime(parts, l.opts)
			verr := getObject(parts, &iterElem)
			var finalErr error = nil
			if rv := customerr.AppendError(serr, terr, verr); rv != nil {
				finalErr = customerr.Wrap(
					LogFileMalformed,
					"File: '%s' | Line: '%d' | %s",
					l.file, index+1, rv,
				)
			}
			return LogEntry[T]{
				Status:  s,
				Time:    t,
				Message: getMessage(parts),
				Val:     iterElem,
			}, finalErr
		},
	)
}

func getStatus(parts []string) (LogStatus, error) {
	if len(parts) > 0 {
		return LogStatusFromString(strings.TrimSpace(parts[0]))
	}
	return -1, NoLogStatus
}

func getTime(parts []string, opts *options) (time.Time, error) {
	if len(parts) >= 1 {
		if rv, err := time.Parse(
			opts.dateTimeFormat,
			strings.TrimSpace(parts[1]),
		); err == nil {
			return rv, err
		} else {
			return time.Time{}, err
		}
	}
	return time.Time{}, NoLogTime
}

func getMessage(parts []string) string {
	if len(parts) >= 2 {
		return strings.TrimSpace(parts[2])
	}
	return ""
}

func getObject[T any](parts []string, elem *T) error {
	if len(parts) >= 3 {
		return json.Unmarshal([]byte(parts[3]), elem)
	}
	return NoLogObj
}
