# Flags

A generator program that creates methods surrounding a bit-flag type.

## Usage

To execute this program use a go generate command of the following structure:

```
//go:generate <path to exe>/flags -type=<type> -package=<package>
```

The type argument is implied to be an some kind of integer value, but that is
not checked by the generator because of the potential complexities of resolving
type aliases and custom types. Whatever type is supplied must implement the
standard bitwise operators, or compiler errors will be present in the generated
code. (This is a form of a check that the supplied type is some kind of an
integer value.)

Given this go generate command, the ```flags``` program will perform the 
following action.

The action is to search the ast of any non-generated code in the current
directory for constant instances of the supplied type. These constant
definitions will comprise the flags set of values. The ```flags``` program
does not perform any value checking - it is up to the user and compiler to
guarintee that there are no duplicate values in the flag definitions. (This
could be a feature or a bug, depends on how you look at it.) With each constant
declaration there will be a couple more arguments in the doc string of the value,
as shown in the example below.

```
const (
    //gen:flags noSetter
    <flag value 1> <flag type>=<value 1>

    // This is a comment
    <flag value 2> <flag type>=<value 2>

    //gen:flags noSetter
    <flag value 3> <flag type>=<value 3>
)
```

The arguments are as follows:

1. noSetter (boolean) (optional): true to not generate a setter function for
this flag, false or missing to generate a setter function for this flag.

Given the example used throughout this file, with the information from the
inline arguments and the comment arguments, and assuming that the default
unknown type was set to <flag value 3>, the following code would be generated.

```
package <package>

var (
    Invalid<flag type> = errors.New("Invalid <flag type>")
    <flag type> []<flag type>=[]<flag type>{
        <flag value 1>,
        <flag value 2>,
        <flag value 3>,
    }
)

// Returns the supplied flags status
func (o <flag type>) GetFlag(flag <flag type>) bool {
	return o&flag > 0
}

// This is a comment
func (o <flag type>) <flag value 1>(b bool) <flag type> {
    if b {
        o |= <flag value 1>
    } else {
        o &= ^<flag value 1>
    }
    return o
}
```

At a high level the following code is generated:

1. Setter functions are generated for each flag value that allows. The doc
string of this function is the same as the doc string of the flag value.
1. A GetFlag function is generated that return true if the supplied flag is true
in the bit map, false otherwise.
