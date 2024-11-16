# argparse

A type safe, extensible CLI argument parsing utility package.

https://github.com/barbell-math/util/blob/436b7e550a6ca905f239473b22a093506eddd798/src/argparse/examples/SimpleExamples_test.go#L47-L72
<sup>Example usage of the argparse package</sup>

## Usage

The above example shows the general usage of this package. In the example there
are three main parts to consider:

1. An `ArgBuilder` is created and it is populated with arguments. This stage is
where the vast majority of your interaction with the package will take place,
and is explained in more detail in the next section.
1. The `ArgBuilder` makes a `Parser`. Optionally this is where subparsers could
be added if you were using them.
1. The `Parser` is then used to parse a slice of strings which is translated
into a sequence of tokens.

## Argument Builder: Primitive Types

Primitive types are the most straight forward. Shown below is how to add an
integer argument. The `translators.BuiltinInt` type is responsible for parsing
an integer from the string value supplied by the CLI. Analgous types are
available for all primitive types, all following the `Builtin<type>` format.

https://github.com/barbell-math/util/blob/436b7e550a6ca905f239473b22a093506eddd798/src/argparse/examples/SimpleExamples_test.go#L18-L19
<sup>Integer argument</sup>

The above example provides an argument with no options, meaning all the default
options will be used. Shown below is how to provide your own set of options.

https://github.com/barbell-math/util/blob/436b7e550a6ca905f239473b22a093506eddd798/src/argparse/examples/SimpleExamples_test.go#L53-L59
<sup>Integer argument with options</sup>

The following options can be provided all through setter methods similar to the
ones shown in the example above:

1. `shortName`: The single character flag that the argument can be specified
with as a short hand for the provided long name.
1. `description`: The description that will be printed on the help menu.
1. `required`: A boolean indicating if the argument is required or not.
1. `defaultVal`: The value that should be returned if the argument is not
provided on the CLI.
1. `translator`: The translator that should be used when translating the CLI
string to the appropriately typed value.
1. `argType`: The type of argument. This value tells the parser what semantics
are valid for the argument.

The available argument types are as follows:

1. `ValueArgType`: Represents a flag type that must accept a single value as an
argument and must only be supplied once.
1. `MultiValueArgType`: Represents a flag type that can accept many values as an
argument and must only be supplied once. At least one argument must be supplied.
1. `FlagArgType`: Represents a flag type that must not accept a value and must
only be supplied once.
1. `MultiFlagArgType`: Represents a flag type that must not accept a value and
may be supplied many times.

## Argument Builder: Out of the Box Argument Types

The `ArgBuilder` also has several helper functions and translators for common
CLI argument types:

1. Flag arguments. This will return true if the flag is provided. It does accept
any values.

https://github.com/barbell-math/util/blob/436b7e550a6ca905f239473b22a093506eddd798/src/argparse/examples/SimpleExamples_test.go#L92-L98

2. Flag counter argument. This will return an integer equal to the number of
times that the flag was provided. It does not accept any values.

https://github.com/barbell-math/util/blob/436b7e550a6ca905f239473b22a093506eddd798/src/argparse/examples/SimpleExamples_test.go#L131-L136

3. List argument. This will build up a list of all the values that were provided
with the argument. Many values can be provided with a single argument or many
flags can be provided with a single argument, as shown in the example arguments
below the argument example.

https://github.com/barbell-math/util/blob/436b7e550a6ca905f239473b22a093506eddd798/src/argparse/examples/SimpleExamples_test.go#L171-L186
https://github.com/barbell-math/util/blob/436b7e550a6ca905f239473b22a093506eddd798/src/argparse/examples/SimpleExamples_test.go#L191

4. List argument with a predefined set of allowed values. This will build up a
list of all the values that were provided with the argument, provided that they
are in the allowed list of values. Many values can be provided with a single
argument or many flags can be provided with a single argument, as shown in the
example arguments below the argument example.

https://github.com/barbell-math/util/blob/436b7e550a6ca905f239473b22a093506eddd798/src/argparse/examples/SimpleExamples_test.go#L210-L230
https://github.com/barbell-math/util/blob/436b7e550a6ca905f239473b22a093506eddd798/src/argparse/examples/SimpleExamples_test.go#L237

5. Selector argument. This will accept a single value as long as that value is
in the predefined set of allowed values.

https://github.com/barbell-math/util/blob/436b7e550a6ca905f239473b22a093506eddd798/src/argparse/examples/SimpleExamples_test.go#L256-L274

## Argument Builder: Custom Types

Due to using generics, the argument builder can accept arguments of custom types
as long as there is a translator for that type. For examples of how to implement
translators refer to the
[stateless translator example](./examples/CustomStatelessTranslator_test.go)
as well as the 
[stateful translator example](./examples/CustomStatefulTranslator_test.go).

To support a custom type the `Translate` method on the translator will simply
need to return the custom type. Support for custom translators and types allows
for a completely type safe translation of the CLI arguments into values that
your program can work with.

## Argument Builder: Computed Arguments

Computed arguments provide a way for the argument parser to set values that were
not directly provided by the CLI, potentially computing values based on the
provided CLI arguments. The example below shows how to add a computed argument
to the parser.

https://github.com/barbell-math/util/blob/436b7e550a6ca905f239473b22a093506eddd798/src/argparse/examples/SimpleExamples_test.go#L315-L318

Much like translators for arguments, computers are needed to set computed
values. Also like translators, computers are expected to return a value, this
time the value that is returned is the result of a computation rather than a
translation. Yet another similarity is that computed arguments support custom
types in all the same ways that translators do. For an example of how to
implement a custom computer refer to the
[custom computer example](./examples/CustomComputer_test.go).

Computed arguments do have one advantage over normal (translated) arguments. All
computed arguments are added to a tree like data structure following the same
ordering that any sub-parsers were added in. Given this tree like data structure
computed arguments are then evaulated bottom up, left to right. With this
organized evaluation ordering it is possible to evauluate arbitrary expressions.
This will likely never be useful and an argument could be made that this should
never be done, but the capibility is there. For an example of this evaulation
refer to the [sub-parsers examples](./examples/SubParsers_test.go).

## Sub-Parsers

Several different parsers, each with there own set of arguments, can be
separately built and then combined into one larger parser. When adding one
parser to another the parsers are added as childen, or sub parsers, to the main
parser. A couple examples of sub-parsers are shown in the
[sub-parsers examples file](./examples/SubParsers_test.go). There are a couple
rules that dictate what happens when sub-parsers are added:

1. All non-computed arguments are added to one single global namespace. This
means that all long argument and short argument names must be unique among all
of the sub-parers.
1. Computed arguments are added in a tree like data structure that mirrors the
structure created by the function calls for adding sub-parsers. The advantage of
this is explained in the previous section.

There are several sub-parsers that are provided out-of-the-box for common CLI
needs.

1. Help: this adds the `-h` and `--help` flag arguments which when encountered
will stop all further parsing and print the help menu.
1. Verbosity: this adds the `-v` and `--verbose` flag counter arguments which
can be used to set a verbosity level for a running application.
