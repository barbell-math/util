package csv

// Code generated by ../bin/structDefaultInit - DO NOT EDIT.
import (
	"time"
)

// Returns a new options struct initialized with the default values.
func NewOptions() *options {
	return &options{
		optionsFlag:    DefaultOptionsFlag(),
		comment:        '#',
		delimiter:      ',',
		structTagName:  "csv",
		dateTimeFormat: time.DateTime,
		headers:        []string{},
	}
}

func (o *options) OptionsFlag(v optionsFlag) *options {
	o.optionsFlag = v
	return o
}

// Description: determines what character is considered to be a comment
//
// Used by: [Parse]
//
// Default: '#'
func (o *options) Comment(v rune) *options {
	o.comment = v
	return o
}

// Description: determines what character is considered to be the
// delimiter that separates fields
//
// Used by: [Parse], [Flatten]
//
// Default: ','
func (o *options) Delimiter(v rune) *options {
	o.delimiter = v
	return o
}

// Description: set to the desired struct tag name to use when mapping
// values to the appropriate fields in the struct
//
// Used by: [ToStructs], [FromStructs]
//
// Default: "csv"
func (o *options) StructTagName(v string) *options {
	o.structTagName = v
	return o
}

// Description: the date time format to use when attempting to parse. No
// correctness checking is performed on the date time format string. Any
// errors from incorrect date time formats will become apparent when
// parsing the CSV file.
//
// Used by: [ToStructs], [FromStructs]
//
// Default: [time.DateTime]
func (o *options) DateTimeFormat(v string) *options {
	o.dateTimeFormat = v
	return o
}

// Description: determines what character is considered to be a comment
//
// Used by: [Parse]
//
// Default: '#'
func (o *options) GetComment() rune {
	return o.comment
}

// Description: determines what character is considered to be the
// delimiter that separates fields
//
// Used by: [Parse], [Flatten]
//
// Default: ','
func (o *options) GetDelimiter() rune {
	return o.delimiter
}

// Description: set to the desired struct tag name to use when mapping
// values to the appropriate fields in the struct
//
// Used by: [ToStructs], [FromStructs]
//
// Default: "csv"
func (o *options) GetStructTagName() string {
	return o.structTagName
}

// Description: the date time format to use when attempting to parse. No
// correctness checking is performed on the date time format string. Any
// errors from incorrect date time formats will become apparent when
// parsing the CSV file.
//
// Used by: [ToStructs], [FromStructs]
//
// Default: [time.DateTime]
func (o *options) GetDateTimeFormat() string {
	return o.dateTimeFormat
}

// Description: the list of headers to use should you want them to be
// different from the options supplied by the struct field names or
// tag names.
//
// Used by: [FromStructs]
//
// Default: true
func (o *options) GetHeaders() []string {
	return o.headers
}
