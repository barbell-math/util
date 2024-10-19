# Flags

A generator program that creates methods surrounding a bit-flag type.

## Usage

To execute this program use a go generate command of the following structure:

```
//go:generate <path to exe>/flags -type=<type> -package=<package>
```

Given this go generate command, the ```flags``` program will perform two
tasks.

The first task ss to search the ast of any non-generated code in the current
directory for constant instances of the supplied type. These constant
definitions will comprise the enum's set of values. The ```flags``` program
does not perform any value checking - it is up to the user and compiler to
guarintee that there are no duplicate values in the enum definition. (This could
be a feature or a bug, depends on how you look at it.) With each constant
declaration there will be a couple more arguments in the doc string of the value
, as shown in the example below.

```
const (
    //gen:flags noSetter
    //gen:flags string <enum value 1 string>
    <enum value 1> <enum type>=<value 1>

    // This is a comment
    //gen:flags string <enum value 2 string>
    <enum value 2> <enum type>=<value 2>

    //gen:flags noSetter
    <enum value 3> <enum type>=<value 3>
)
```

The arguments are as follows:

1. noSetter (boolean) (optional): true to not generate a setter function for
this flag, false or missing to generate a setter function for this flag.

Given the example used throughout this file, with the information from the
inline arguments and the comment arguments, and assuming that the default
unknown type was set to <enum value 3>, the following code would be generated.

```
package <package>

var (
    Invalid<enum type> = errors.New("Invalid <enum type>")
    <enum type> []<enum type>=[]<enum type>{
        <enum value 1>,
        <enum value 2>,
        <enum value 3>,
    }
)

// Returns the supplied flags status
func (o <enum type>) GetFlag(flag <enum type>) bool {
	return o&flag > 0
}

// This is a comment
//gen:flags string <enum value 2 string>
func (o <enum type>) <enum value 1>(b bool) <enum type> {
    if b {
        o |= <enum value 1>
    } else {
        o &= ^<enum value 1>
    }
    return o
}
```

At a high level the following code is generated:

1. Setter functions are generated for each enum value that allows. The doc
string of this function is the same as the doc string of the enum value.
1. A GetFlag function is generated that return true if the supplied flag is true
in the bit map, false otherwise.
