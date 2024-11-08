package log

import "errors"

var (
	LogFileNotSpecified = errors.New("The log file was not specified.")
	LogFileMalformed = errors.New("Log line malformed.")
	NoLogStatus = errors.New("No log status present.")
	NoLogTime = errors.New("No log time present.")
	NoLogObj = errors.New("No object present.")
	InvalidLogStatus = errors.New("Invalid log status.")
)
