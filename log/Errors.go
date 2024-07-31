package log

import "errors"

var LogFileNotSpecified = errors.New("The log file was not specified.")
var LogFileMalformed = errors.New("Log line malformed.")
var NoLogStatus = errors.New("No log status present.")
var NoLogTime = errors.New("No log time present.")
var NoLogObj = errors.New("No object present.")
var InvalidLogStatus = errors.New("Invalid log status.")
