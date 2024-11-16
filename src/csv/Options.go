package csv

//go:generate ../../bin/enum -type=optionsFlag -package=csv
//go:generate ../../bin/flags -type=optionsFlag -package=csv
//go:generate ../../bin/structDefaultInit -struct=options

type (
	//gen:enum unknownValue unknownOptionsFlag
	//gen:enum default 0 | hasHeaders | useStructTags | writeHeaders
	optionsFlag int
	options     struct {
		// Description: boolean options encoded in a bit flag enum.
		//
		// Used by: [Parse], [Flatten], [FromStructs]
		//
		// Default: See the [NewOptionsFlag] function.
		optionsFlag `default:"NewOptionsFlag()" setter:"t" getter:"f"`

		// Description: determines what character is considered to be a comment
		//
		// Used by: [Parse]
		//
		// Default: '#'
		comment rune `default:"'#'" setter:"t" getter:"t"`
		// Description: determines what character is considered to be the
		// delimiter that separates fields
		//
		// Used by: [Parse], [Flatten]
		//
		// Default: ','
		delimiter rune `default:"','" setter:"t" getter:"t"`
		// Description: set to the desired struct tag name to use when mapping
		// values to the appropriate fields in the struct
		//
		// Used by: [ToStructs], [FromStructs]
		//
		// Default: "csv"
		structTagName string `default:"\"csv\"" setter:"t" getter:"t"`
		// Description: the date time format to use when attempting to parse. No
		// correctness checking is performed on the date time format string. Any
		// errors from incorrect date time formats will become apparent when
		// parsing the CSV file.
		//
		// Used by: [ToStructs], [FromStructs]
		//
		// Default: [time.DateTime]
		dateTimeFormat string `default:"time.DateTime" import:"time" setter:"t" getter:"t"`
		// Description: the list of headers to use should you want them to be
		// different from the options supplied by the struct field names or
		// tag names.
		//
		// Used by: [FromStructs]
		//
		// Default: true
		headers []string `default:"[]string{}" setter:"f" getter:"t"`
	}
)

const (
	// Description: set to true if the incoming iterator stream has
	// headers in the first row
	//
	// Used by: [ToStructs]
	//
	// Default: true
	//gen:enum string hasHeaders
	hasHeaders optionsFlag = 1 << iota
	// Description: set to true to skip the headers from the incoming
	// iterator stream and instead determine field ordering by the order of
	// the fields in the struct.
	//
	// Used by: [ToStructs]
	//
	// Default: false
	//gen:enum string ignoreHeaders
	ignoreHeaders
	// Description: set to true to use struct field tags instead of the
	// field name when a tag is present and has the same name as defined by
	// the structTagName option
	//
	// Used by: [ToStructs], [FromStructs]
	//
	// Default: true
	//gen:enum string useStructTags
	useStructTags
	// Description: set to true to write the headers to the file
	//
	// Used by: [FromStructs]
	//
	// Default: true
	//gen:enum string writeHeaders
	writeHeaders
	//gen:flags	noSetter
	//gen:enum string headersSupplied
	headersSupplied
	// Description: whether or not to write zero-values to the csv file. If false,
	// any zero values will be left as blank fields.
	//
	// Used by: [FromStructs]
	//
	// Default: false
	//gen:enum string writeZeroValues
	writeZeroValues
	//gen:flags	noSetter
	//gen:enum string unknownOptionsFlag
	unknownOptionsFlag
)

// Description: the list of headers to use should you want them to be
// different from the options supplied by the struct field names or
// tag names.
//
// Used by: [FromStructs]
//
// Default: true
func (o *options) SetHeaders(h []string) *options {
	o.headers = h
	o.optionsFlag |= headersSupplied
	return o
}
