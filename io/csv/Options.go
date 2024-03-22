package csv

import "time"

type (
    options struct {
        // Description: determines what character is considered to be a comment
        //
        // Used by: [Parse]
        //
        // Default: '#'
        comment rune
        // Description: determines what character is considered to be the
        // delimiter that separates fields
        //
        // Used by: [Parse]
        //
        // Default: ','
        delimiter rune

        // Description: set to true if the incoming iterator stream has
        // headers in the first row
        //
        // Used by: [ToStruct]
        //
        // Default: true
        hasHeaders bool
        // Description: set to true to skip the headers from the incoming
        // iterator stream and instead determine field ordering by the order of
        // the fields in the struct
        //
        // Used by: [ToStruct]
        //
        // Default: false
        ignoreHeaders bool
        // Description: set to true to use struct field tags instead of the
        // field name when a tag is present and has the same name as defined by
        // the structTagName option
        //
        // Used by: [ToStruct]
        //
        // Default: true
        useStructTags bool
        // Description: set to the desired struct tag name to use when mapping
        // values to the appropriate fields in the struct
        //
        // Used by: [ToStruct]
        //
        // Default: "csv"
        structTagName string
        // Description: the date time format to use when attempting to parse
        // date time fields
        //
        // Used by: [ToStruct]
        //
        // Default: [time.DateTime]
        dateTimeFormat string


        writeHeaders bool
        headers []string
        headersSupplied bool
    }
)

// Returns a new options struct initialized with the default values that can be
// passed to the other functions in this file that require options.
func NewOptions() *options {
    return &options{
        comment: '#',
        delimiter: ',',
        hasHeaders: true,
        ignoreHeaders: false,
        useStructTags: true,
        structTagName: "csv",
        dateTimeFormat: time.DateTime,
        writeHeaders: true,
        headers: []string{},
        headersSupplied: false,
    }
}

// Sets the comment option in an options struct
func (o *options)Comment(c rune) *options {
    o.comment=c
    return o
}

// Sets the delimiter option in an options struct
func (o *options)Delimiter(d rune) *options {
    o.delimiter=d
    return o
}

// Sets the has headers option in an options struct
func (o *options)HasHeaders(b bool) *options {
    o.hasHeaders=b
    return o
}

// Sets the ignore headers option in an options struct
func (o *options)IgnoreHeaders(b bool) *options {
    o.ignoreHeaders=b
    return o
}

// Sets the use struct tags option in an options struct
func (o *options)UseStructTags(b bool) *options {
    o.useStructTags=b
    return o
}

// Sets the struct tag name option in an options struct
func (o *options)StructTagName(s string) *options {
    o.structTagName=s
    return o
}

// Sets the date time format in an options struct. No correctness checking is
// performed on the date time format string. Any errors from incorrect date time
// formats will become apparent when parsing the CSV file.
func (o *options)DateTimeFormat(f string) *options {
    o.dateTimeFormat=f
    return o
}

func (o *options)WriteHeaders(b bool) *options {
    o.writeHeaders=b
    return o
}

func (o *options)Headers(h []string) *options {
    o.headers=h
    o.headersSupplied=true
    return o
}
