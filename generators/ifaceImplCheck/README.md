# Interface Implementation Check

A generator that provides unit test boiler plate code that checks if the
supplied type implements the supplied interface.

## Usage

To execute this program use a go generate command of the following structure:

```
//go:generate <path to exe>/ifaceImplCheck -typeToCheck=<type>
```

Given this generate command, the ```structDefaultInit``` program will search the
ast of any non-generated code in the current directory for the supplied type
definition. Once the supplied type definition is found it will use the doc
string to generate the unit tests. Before going over the code this is generated,
the code below shows the expected type format:

```
//gen:ifaceImplCheck generics <generic vals 1>
//gen:ifaceImplCheck ifaceName <ifaceName 1>
//gen:ifaceImplCheck valOrPntr [val | pntr | both]
type <custom type 1> <type definition>

//gen:ifaceImplCheck generics <generic vals 1>
//gen:ifaceImplCheck ifaceName <ifaceName 2>
//gen:ifaceImplCheck imports <import path 1> <desired name>-><import path 2>
//gen:ifaceImplCheck valOrPntr [val | pntr | both]
type <custom type 2> <type definition>
```

The following comment arguments are supported for the type definition:

1. generics (string) (optional): if the supplied type requires generic type
parameters then supply them here. This should include the square brackets.
1. ifaceName (string) (required): the interface that the type should implement.
1. imports (string) (optional): a space separated list of imports to include in 
the generated code file. This is useful when the translator or a generic type
parameter value is derived from a type in an external package. Every import will
be automatically wrapped in quotes. To import something under a different name
use the following syntax: `<desired name>-><import path>`
1. valOrPntr (string) (required): whether the type implements the supplied
interface on a pointer receiver, value receiver, or both. The resulting unit
tests will only test the provided receiver type.

Given the example used throughout this file, with the information from the 
inline arguments and the comment args the following code will be generated:

```
package <package>

import (
    "testing"
    "<import path 1>"
    <desired name> "<import path 2>"
)

func Test<custom type 1>ValImplements<ifaceName 1>(t *testing.T) {
	var typeThing <custom type 1><generic vals 1>
	var iFaceThing <ifaceName 1> = typeThing
	_ = iFaceThing
}

func Test<custom type 1>PntrImplements<ifaceName 1>(t *testing.T) {
	var typeThing <custom type 1><generic vals 1>
	var iFaceThing <ifaceName > = &typeThing
	_ = iFaceThing
}

func Test<custom type 2>ValImplements<ifaceName 2>(t *testing.T) {
	var typeThing <custom type 2><generic vals 2>
	var iFaceThing <ifaceName 2> = typeThing
	_ = iFaceThing
}

func Test<custom type 2>PntrImplements<ifaceName 2>(t *testing.T) {
	var typeThing <custom type 2><generic vals 2>
	var iFaceThing <ifaceName 2> = &typeThing
	_ = iFaceThing
}
```

At a high level the following code is generated:

1. Unit test functions that check the required interface is implemented on the
required type receiver.
