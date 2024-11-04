package csv

import "errors"

var (
	MalformedCSVFile = errors.New("The supplied CSV file is not valid.")
	MalformedCSVStruct = errors.New("The supplied struct was not valid.")
	DuplicateColName = errors.New("A column name was duplicated either in the struct definition or the file itself.")
	InvalidHeader = errors.New("A header specified in the CSV file or options could not be mapped to the supplied struct using the supplied options. Make sure the specified field is exported.")
	InvalidStructField = errors.New("An error occurred while trying to set the struct field. Make sure the field is settable.")
	InvalidHeaders = errors.New("The supplied list of headers is not valid with the supplied struct.")
)
