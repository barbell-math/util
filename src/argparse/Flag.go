package argparse

import "regexp"

//go:generate ../../bin/enum -type=flag -package=argparse
//go:generate ../../bin/passThroughWidget -type=tokenType

type (
	// Used to represent the different formats of flags that the arg parser
	// accepts. View the enum values for examples of these flag formats and the
	// variants that each one accepts.
	//gen:enum unknownValue unknownFlag
	//gen:enum default unknownFlag
	//gen:passThroughWidget widgetType Base
	//gen:passThroughWidget baseTypeWidget widgets.BuiltinInt
	//gen:passThroughWidget widgetPackage github.com/barbell-math/util/src/widgets
	flag int
)

const (
	// Represents an unknown flag value.
	//gen:enum string unknownFlag
	unknownFlag flag = iota
	// Represents a short flag, that is one with only one dash and one character.
	//
	// Example: -t 10:00
	//gen:enum string shortSpaceFlag
	shortSpaceFlag
	// Represents a short flag, that is one with only one dash and one character,
	// with an equals sign and a value.
	//
	// Example: -t=10:00
	//gen:enum string shortEqualsFlag
	shortEqualsFlag
	// Represents a long flag, that is one with two dashes and a full name,
	// without an equals sign.
	//
	// Example: --time 10:00
	//gen:enum string longSpaceFlag
	longSpaceFlag
	// Represents a long flag, that is one with two dashes and a full name,
	// with an equals sign and value.
	//
	// Example: --time=10:00
	//gen:enum string longEqualsFlag
	longEqualsFlag
	// Represents a config file flag with an equal sign. A config file flag
	// will have a single value that is a path to a file that contains more
	// arguments.
	//
	// Example: --config=/path/to/file
	//gen:enum string configEqualsFileFlag
	configEqualsFileFlag
	// Represents a config file flag with an space. A config file flag
	// will have a single value that is a path to a file that contains more
	// arguments.
	//
	// Example: --config /path/to/file
	//gen:enum string configEqualsFileFlag
	configSpaceFileFlag
)

var (
	regexes = map[flag]*regexp.Regexp{
		shortSpaceFlag:       regexp.MustCompile("^-.*$"),
		shortEqualsFlag:      regexp.MustCompile("^-.*=.*$"),
		longSpaceFlag:        regexp.MustCompile("^--.*$"),
		longEqualsFlag:       regexp.MustCompile("^--.*=.*$"),
		configEqualsFileFlag: regexp.MustCompile("^--config=.*$"),
		configSpaceFileFlag:  regexp.MustCompile("^--config$"),
	}
)
