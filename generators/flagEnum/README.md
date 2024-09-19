# Flag Enum

A generator program that creates methods surrounding a bit-flag enum type.

## Usage

To execute this program use a go generate command of the following structure:

```
//go:generate <path to exe>/flagEnum -type=<type> -package=<package>
```

Given this go generate command, the ```flagEnum``` program will perform two
tasks.

The first task is to search the ast of any non-generated code in the current
directory for the supplied types definition. Once this definition is found, it
is expected that there will be several more arguments in the doc string of the
type, as shown below.

```
type (
	//gen:flagEnum unknownValue <default unknown value>
	//gen:flagEnum default <default value>
	optionsFlag int
)
```

The arguments are as follows:

1. unknownValue (string) (required): the constant value that is part of the enum
that represents an unknown or invalid enum state. If a value is supplied that is
not found to be value for this enum the program will print an error and exit with
a non-zero exit status.
1. default (string) (required): The default value that a new instance of this
enum should be initilized with. This is useful for setting a default set of
flags.

The second task is to search the ast of any non-generated code in the current
directory for constant instances of the supplied type. These constant
definitions will comprise the enum's set of values. The ```flagEnum``` program
does not perform any value checking - it is up to the user and compiler to
guarintee that there are no duplicate values in the enum definition. (This could
be a feature or a bug, depends on how you look at it.) With each constant
declaration there will be a couple more arguments in the doc string of the value
, as shown in the example below.

```
const (
    //gen:flagEnum noSetter
    //gen:flagEnum string <enum value 1 string>
    <enum value 1> <enum type>=<value 1>

    // This is a comment
    //gen:flagEnum string <enum value 2 string>
    <enum value 2> <enum type>=<value 2>

    //gen:flagEnum noSetter
    //gen:flagEnum string <enum value 3 string>
    <enum value 3> <enum type>=<value 3>
)
```

The arguments are as follows:

1. noSetter (boolean) (optional): true to not generate a setter function for
this flag, false or missing to generate a setter function for this flag.
1. string (string) (required): the string representation of this enum value.
This is often useful when decoding config files where all valid keys/values can
be represented by a set of enum values.


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

func Default<enum type>() <enum type> {
    return <default value>
}

func (o <enum type>) String() string {
    switch o {
    case <enum value 1>: return <enum value 1 string>
    case <enum value 2>: return <enum value 2 string>
    case <enum value 3>: return <enum value 3 string>
    default: return <enum value 3 string>
    }
}

func (o *<enum type>) FromString(s string) error {
    switch s {
    case "<enum value 1 string>":
        *o = <enum value 1>
        return nil
    case "<enum value 2 string>":
        *o = <enum value 2>
        return nil
    case "<enum value 3 string>":
        *o = <enum value 3>
        return nil
    default:
        *o = <enum value 3>
        return fmt.Errorf("%w: %s", Invalid<enum type>, s)
    }
}

// Returns the supplied flags status
func (o <enum type>) GetFlag(flag <enum type>) bool {
	return o&flag > 0
}

// This is a comment
//gen:flagEnum string <enum value 2 string>
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

1. A error that represents invalid enum values is generated.
1. A list of enum values is generated.
1. String and FromString methods are generated that allow translating an enum
value into a string value.
1. Setter functions are generated for each enum value that allows. The doc
string of this function is the same as the doc string of the enum value.
1. A GetFlag function is generated that return true if the supplied flag is true
in the bit map, false otherwise.
