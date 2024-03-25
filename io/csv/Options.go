package csv

import "time"

type (
	optionsFlag int
	options struct {
		flags optionsFlag

		comment   rune
		delimiter rune

		structTagName  string
		dateTimeFormat string

		headers         []string
	}
)

const (
	hasHeaders optionsFlag=1<<iota
	ignoreHeaders
	useStructTags
	writeHeaders
	headersSupplied
	writeZeroValues
)

// Returns a new options struct initialized with the default values that can be
// passed to the other functions in this package that require options.
func NewOptions() *options {
	return &options{
		flags: 0 | hasHeaders | useStructTags | writeHeaders,
		comment:         '#',
		delimiter:       ',',
		structTagName:   "csv",
		dateTimeFormat:  time.DateTime,
		headers:         []string{},
	}
}

// Description: determines what character is considered to be a comment
//
// Used by: [Parse]
//
// Default: '#'
func (o *options) Comment(c rune) *options {
	o.comment = c
	return o
}

// Description: determines what character is considered to be the
// delimiter that separates fields
//
// Used by: [Parse], [Flatten]
//
// Default: ','
func (o *options) Delimiter(d rune) *options {
	o.delimiter = d
	return o
}

// Description: set to true if the incoming iterator stream has
// headers in the first row
//
// Used by: [ToStructs]
//
// Default: true
func (o *options) HasHeaders(b bool) *options {
	if b {
		o.flags|=hasHeaders
	} else {
		o.flags&=^hasHeaders
	}
	return o
}

// Description: set to true to skip the headers from the incoming
// iterator stream and instead determine field ordering by the order of
// the fields in the struct.
//
// Used by: [ToStructs]
//
// Default: false
func (o *options) IgnoreHeaders(b bool) *options {
	if b {
		o.flags|=ignoreHeaders
	} else {
		o.flags&=^ignoreHeaders
	}
	return o
}

// Description: set to true to use struct field tags instead of the
// field name when a tag is present and has the same name as defined by
// the structTagName option
//
// Used by: [ToStructs], [FromStructs]
//
// Default: true
func (o *options) UseStructTags(b bool) *options {
	if b {
		o.flags|=useStructTags
	} else {
		o.flags&=^useStructTags
	}
	return o
}

// Description: set to the desired struct tag name to use when mapping
// values to the appropriate fields in the struct
//
// Used by: [ToStructs], [FromStructs]
//
// Default: "csv"
func (o *options) StructTagName(s string) *options {
	o.structTagName = s
	return o
}

// Description: the date time format to use when attempting to parse. No
// correctness checking is performed on the date time format string. Any errors
// from incorrect date time formats will become apparent when parsing the CSV
// file.
//
// Used by: [ToStructs], [FromStructs]
//
// Default: [time.DateTime]
func (o *options) DateTimeFormat(f string) *options {
	o.dateTimeFormat = f
	return o
}

// Description: set to true to write the headers to the file
//
// Used by: [FromStructs]
//
// Default: true
func (o *options) WriteHeaders(b bool) *options {
	if b {
		o.flags|=writeHeaders
	} else {
		o.flags&=^writeHeaders
	}
	return o
}

// Description: the list of headers to use should you want them to be
// different from the options supplied by the struct field names or
// tag names.
//
// Used by: [FromStructs]
//
// Default: true
func (o *options) Headers(h []string) *options {
	o.headers = h
	o.flags|=headersSupplied
	return o
}

// Description: whether or not to write zero-values to the csv file. If false,
// any zero values will be left as blank fields.
//
// Used by: [FromStructs]
//
// Default: false
func (o *options) WriteZeroValues(b bool) *options {
	if b {
		o.flags|=writeZeroValues
	} else {
		o.flags&=^writeZeroValues
	}
	return o
}

func (o *options) getFlag(flag optionsFlag) bool {
	return o.flags&flag>0
}
