package csv

//go:generate ../bin/optionsFlags -optionsStruct=options -optionsEnum=optionsFlag -package=csv -debug

type (
	optionsFlag int
	options     struct {
		flags optionsFlag `default:"0 | hasHeaders | useStructTags | writeHeaders"`

		// Description: determines what character is considered to be a comment
		//
		// Used by: [Parse]
		//
		// Default: '#'
		comment rune `auto:"t" default:"'#'"`
		// Description: determines what character is considered to be the
		// delimiter that separates fields
		//
		// Used by: [Parse], [Flatten]
		//
		// Default: ','
		delimiter rune `auto:"t" default:"','"`
		// Description: set to the desired struct tag name to use when mapping
		// values to the appropriate fields in the struct
		//
		// Used by: [ToStructs], [FromStructs]
		//
		// Default: "csv"
		structTagName string `auto:"t" default:"\"csv\""`
		// Description: the date time format to use when attempting to parse. No
		// correctness checking is performed on the date time format string. Any
		// errors from incorrect date time formats will become apparent when
		// parsing the CSV file.
		//
		// Used by: [ToStructs], [FromStructs]
		//
		// Default: [time.DateTime]
		dateTimeFormat string `auto:"t" default:"time.DateTime" import:"\"time\""`
		// Description: the list of headers to use should you want them to be
		// different from the options supplied by the struct field names or
		// tag names.
		//
		// Used by: [FromStructs]
		//
		// Default: true
		headers []string `auto:"f" default:"[]string{}"`
	}
)

const (
	// Description: set to true if the incoming iterator stream has
	// headers in the first row
	//
	// Used by: [ToStructs]
	//
	// Default: true
	hasHeaders optionsFlag = 1 << iota
	// Description: set to true to skip the headers from the incoming
	// iterator stream and instead determine field ordering by the order of
	// the fields in the struct.
	//
	// Used by: [ToStructs]
	//
	// Default: false
	ignoreHeaders
	// Description: set to true to use struct field tags instead of the
	// field name when a tag is present and has the same name as defined by
	// the structTagName option
	//
	// Used by: [ToStructs], [FromStructs]
	//
	// Default: true
	useStructTags
	// Description: set to true to write the headers to the file
	//
	// Used by: [FromStructs]
	//
	// Default: true
	writeHeaders
	headersSupplied //optionsFlags ignore
	// Description: whether or not to write zero-values to the csv file. If false,
	// any zero values will be left as blank fields.
	//
	// Used by: [FromStructs]
	//
	// Default: false
	writeZeroValues
)

// Description: the list of headers to use should you want them to be
// different from the options supplied by the struct field names or
// tag names.
//
// Used by: [FromStructs]
//
// Default: true
func (o *options) Headers(h []string) *options {
	o.headers = h
	o.flags |= headersSupplied
	return o
}
