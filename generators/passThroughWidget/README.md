# Pass Through Widget

This is a generator program that is used to create pass through widgets. A pass
through widget is a widget that simply calls the methods of another widget
without adding any additional logic. It is assumed that the two widget types can
be cast to and from each other.

## Usage

To execute this program use a go generate command of the following structure:

```
//go:generate <path to exe>/passThroughWidget -type=<type>
```

Given this generate command, the ```passThroughWidget``` program will search the
ast of any non-generated code in the current directory for the type definition
of the supplied type. Once the type definition has been found it will use the
doc string associated with the type to get several more arguments, an example of
which are shown below.

```
type (
	//gen:passThroughWidget widgetType <widget type>
	//gen:passThroughWidget package <package>
	//gen:passThroughWidget baseTypeWidget <base type widget>
	//gen:passThroughWidget widgetPackage <base type widget package>
	<type> <base type>
)
```

The comment arguments are listed below:

1. widgetType (string) (required): the type of widget to generate. One of Base,
PartialOrder, Arith, or PartialOrderArith.
1. package (string) (required): the package the supplied type is in.
1. baseTypeWidget (string) (required): the widget that is associated with the
current types underlying type.
1. widgetPackage (string) (required): the package the base types widget is in.

Given the example used throughout this file, with the information from the 
inline arguments and the struct tags the following code, and assuming that a
base type widget was specified, the following code will be generated:

```
package <package>

import (
	"github.com/barbell-math/util/hash"
)

// Returns true if l equals r. Uses the Eq operator provided by the
// *HashSetHash widget internally.
func (_ *<type>) Eq(l *<type>, r *<type>) bool {
	var tmp *<base type widget>
	return tmp.Eq((*<base type widget>)(l), (*<base type widget>)(r))
}

// Returns a hash to represent other. The hash that is returned will be supplied
// by the *HashSetHash widget internally.
func (_ *<type>) Hash(other *<type>) hash.Hash {
	var tmp *<base type widget>
	return tmp.Hash((*<base type widget>)(other))
}

// Zeros the supplied value. The operation that is performed will be determined
// by the *HashSetHash widget internally.
func (_ *<type>) Zero(other *<type>) {
	var tmp *<base type widget>
	tmp.Zero((*<base type widget>)(other))
}
```

At a high level the following code is generated:

1. Widget pass through methods are generated that will match the interface type
of the supplied supplied widget type.
