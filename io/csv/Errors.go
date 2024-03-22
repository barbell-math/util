package csv

import "errors"

var MalformedCSVFile=errors.New("The supplied CSV file is not valid.")
var MalformedCSVStruct=errors.New("The supplied struct was not valid.")
var DuplicateColName=errors.New("A column name was duplicated either in the struct definition or the file itself.")
var InvalidHeader=errors.New("A header specified in the CSV file or options could not be mapped to the supplied struct using the supplied options. Make sure the specified field is exported.")
var InvalidStructField=errors.New("An error occurred while trying to set the struct field. Make sure the field is settable.")
var InvalidHeaders=errors.New("The supplied list of headers is not valid with the supplied struct.")
