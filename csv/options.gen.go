package csv

import (
	"time"
)

// Returns a new options struct initialized with the default values.
func NewOptions() *options {
	return &options{
		headers:	[]string{},
		flags:		0 | hasHeaders | useStructTags | writeHeaders,
		comment:	'#',
		delimiter:	',',
		structTagName:	"csv",
		dateTimeFormat:	time.DateTime,
	}
}

func (o *options) getFlag(flag optionsFlag) bool {
	return o.flags&flag > 0
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
		o.flags |= useStructTags
	} else {
		o.flags &= ^useStructTags
	}
	return o
}

// Description: set to true to write the headers to the file
//
// Used by: [FromStructs]
//
// Default: true
func (o *options) WriteHeaders(b bool) *options {
	if b {
		o.flags |= writeHeaders
	} else {
		o.flags &= ^writeHeaders
	}
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
		o.flags |= writeZeroValues
	} else {
		o.flags &= ^writeZeroValues
	}
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
		o.flags |= hasHeaders
	} else {
		o.flags &= ^hasHeaders
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
		o.flags |= ignoreHeaders
	} else {
		o.flags &= ^ignoreHeaders
	}
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
