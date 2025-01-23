# argparse

A type safe, extensible CLI argument parsing utility package.

https://github.com/barbell-math/util/blob/d4b081dad4b35c30ca0bb67e6ed603ca04059a3b/src/argparse/examples/SimpleExamples_test.go#L49-L74
<sup>Example usage of the argparse package</sup>

## Usage

The above example shows the general usage of this package. In the example there
are three main parts to consider:

1. An `ArgBuilder` is created and it is populated with arguments. This stage is
where the vast majority of your interaction with the package will take place,
and is explained in more detail in the next section.
1. The `ArgBuilder` makes a `Parser`. Optionally this is where sub parsers could
be added if you were using them.
1. The `Parser` is then used to parse a slice of strings which is translated
into a sequence of tokens.

## Argument Builder: Primitive Types

Primitive types are the most straight forward. Shown below is how to add an
integer argument. The `translators.BuiltinInt` type is responsible for parsing
an integer from the string value supplied by the CLI. Analogous types are
available for all primitive types, all following the `Builtin<type>` format.

https://github.com/barbell-math/util/blob/d4b081dad4b35c30ca0bb67e6ed603ca04059a3b/src/argparse/examples/SimpleExamples_test.go#L20-L21
<sup>Integer argument</sup>

The above example provides an argument with no options, meaning all the default
options will be used. Shown below is how to provide your own set of options.

https://github.com/barbell-math/util/blob/d4b081dad4b35c30ca0bb67e6ed603ca04059a3b/src/argparse/examples/SimpleExamples_test.go#L55-L61
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
CLI argument types.

1. Flag arguments. This will return true if the flag is provided. It does accept
any values.

https://github.com/barbell-math/util/blob/d4b081dad4b35c30ca0bb67e6ed603ca04059a3b/src/argparse/examples/SimpleExamples_test.go#L94-L100

2. Flag counter argument. This will return an integer equal to the number of
times that the flag was provided. It does not accept any values.

https://github.com/barbell-math/util/blob/d4b081dad4b35c30ca0bb67e6ed603ca04059a3b/src/argparse/examples/SimpleExamples_test.go#L133-L138

3. List argument. This will build up a list of all the values that were provided
with the argument. Many values can be provided with a single argument or many
flags can be provided with a single argument, as shown in the example arguments
below the argument example.

https://github.com/barbell-math/util/blob/d4b081dad4b35c30ca0bb67e6ed603ca04059a3b/src/argparse/examples/SimpleExamples_test.go#L171-L188
https://github.com/barbell-math/util/blob/d4b081dad4b35c30ca0bb67e6ed603ca04059a3b/src/argparse/examples/SimpleExamples_test.go#L193

4. List argument with a predefined set of allowed values. This will build up a
list of all the values that were provided with the argument, provided that they
are in the allowed list of values. Many values can be provided with a single
argument or many flags can be provided with a single argument, as shown in the
example arguments below the argument example. Note that given the design of this
translator the list can contain any type, as long as the underlying type has a
translator of it's own.

https://github.com/barbell-math/util/blob/d4b081dad4b35c30ca0bb67e6ed603ca04059a3b/src/argparse/examples/SimpleExamples_test.go#L212-L232
https://github.com/barbell-math/util/blob/d4b081dad4b35c30ca0bb67e6ed603ca04059a3b/src/argparse/examples/SimpleExamples_test.go#L239

5. Selector argument. This will accept a single value as long as that value is
in the predefined set of allowed values. As with the list argument, the selector
translator can work with any type as long as it has an underlying translator of
it's own.

https://github.com/barbell-math/util/blob/d4b081dad4b35c30ca0bb67e6ed603ca04059a3b/src/argparse/examples/SimpleExamples_test.go#L260-L278

6. File argument. This will accept a single string value and verify that the
supplied string is a path that exists as a file.

https://github.com/barbell-math/util/blob/d4b081dad4b35c30ca0bb67e6ed603ca04059a3b/src/argparse/examples/SimpleExamples_test.go#L398-L403

7. Directory argument. This will accept a single string value and verify that
the supplied string is a path that exists as a directory.

https://github.com/barbell-math/util/blob/d4b081dad4b35c30ca0bb67e6ed603ca04059a3b/src/argparse/examples/SimpleExamples_test.go#L427-L432

8. File open argument. This will accept a single string value and will attempt
to make the supplied file with the given file mode and permissions.

https://github.com/barbell-math/util/blob/d4b081dad4b35c30ca0bb67e6ed603ca04059a3b/src/argparse/examples/SimpleExamples_test.go#L456-L468

9. Mkdir argument. This will accept a single string value and will attempt to
make the directory(s) that are denoted by the string value.

https://github.com/barbell-math/util/blob/d4b081dad4b35c30ca0bb67e6ed603ca04059a3b/src/argparse/examples/SimpleExamples_test.go#L493-L501

10. Enum argument. This will accept a single string value and will attempt to
translate it to the underlying enum value given the enum type it was supplied
with through the generic parameters.

https://github.com/barbell-math/util/blob/d4b081dad4b35c30ca0bb67e6ed603ca04059a3b/src/argparse/examples/SimpleExamples_test.go#L302-L312


## Argument Builder: Custom Types

Due to using generics, the argument builder can accept arguments of custom types
as long as there is a translator for that type. For examples of how to implement
translators refer to the
[stateless translator example](./examples/CustomStatelessTranslator_test.go)
as well as the 
[stateful translator example](./examples/CustomStatefulTranslator_test.go).
Any custom types defined outside of this package will use the `AddArg` function
to add arguments.

To support a custom type the `Translate` method on the translator will simply
need to return the custom type. Support for custom translators and types allows
for a completely type safe translation of the CLI arguments into values that
your program can work with.

## Argument Conditionality

In some circumstances it is not enough to simply set arguments as either
required or not. Sometimes certain arguments should be required if another
argument is supplied, or if another argument is a specific value. A motivating
example of this kind of scenario would be an enum argument that defines what
action your program should take. It is reasonable that your program would
require different arguments depending on the value supplied to this enum
argument.

To accommodate for this, this package provides the ability to conditionally
require arguments based an another argument. An example of this is shown below.

## Argument Builder: Computed Arguments

Computed arguments provide a way for the argument parser to set values that were
not directly provided by the CLI, potentially computing values based on the
provided CLI arguments. The example below shows how to add a computed argument
to the parser.

https://github.com/barbell-math/util/blob/d4b081dad4b35c30ca0bb67e6ed603ca04059a3b/src/argparse/examples/SimpleExamples_test.go#L354-L357

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
computed arguments are then evaluated bottom up, left to right. With this
organized evaluation ordering it is possible to evaluate arbitrary expressions.
This will likely never be useful and an argument could be made that this should
never be done, but the capability is there. For an example of this evaluation
refer to the [sub-parsers examples](./examples/SubParsers_test.go).

## Sub-Parsers

Several different parsers, each with there own set of arguments, can be
separately built and then combined into one larger parser. When adding one
parser to another the parsers are added as children, or sub parsers, to the main
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

## Argument Config Files

An argument config file format is provided out of the box for the case where the
list of arguments a program needs gets very large. This argument config file
provides a place to put arguments and allows for a basic form of grouping. The
`--config` long argument name is a reserved name and will be available to every
argparse parser. It is used to specify a path to an argument config file, as
shown below.

```
# Both an equals sign and a space are allowed
./<prog> --config /path/toFile
./<prog> --config=/path/toFile
```

The format of the config file is best shown by example.

https://github.com/barbell-math/util/blob/d4b081dad4b35c30ca0bb67e6ed603ca04059a3b/src/argparse/testData/ValidConfigFile.txt#L1-L17
<sup>An example argument config file</sup>

The above config file is equivalent to the cmd shown below:

```
./<prog> --Name0 Value0 --Group1Group2Name1 Value1 --Group1Group2Group3Name2 Value2 --Group1Group2Name3 Value3 --Group1Name4 Value4 --Group1Name5 Value5 --Name6 Value6
```

There are several points of note:

1. Notice how in the above example the group names get prepended to the argument
names to create the final argument name. Nested group names are appended in
the order they are nested.
1. All names are expected to be valid long names that will be recognized by the
argument parser after group names are done being prepended. Short names were
disallowed to make things clearer in the config file. There is no sense in
single letter names in a config file.
1. Arguments will be parsed in the top down order they are given in the file
regardless of how deeply nested they are. Duplicating single value arguments
will result in an error. Duplicating multi value arguments will not result in an
error.
1. Many argument config files can be specified, though this is discouraged as
it is easy to duplicate arguments between the files.
1. Standard cmd line arguments can be used in conjunction with config files. As
usual, if single value arguments are duplicated an error will be returned and if
multi value arguments are duplicated no error will be returned.

A more practical example of using a config file is shown using the [db] packages
argparse interface. The below config file and cmd line arguments are equivalent.

```
# ./Config.txt
db {
    User <user>
    EnvPswdVar <var>
    NetLoc <url>
    Port <port num>
    Name <db name>
}
```

```
./<prog> --dbUser <user> --dbEnvPswdVar <var> --dbNetLoc <url> --dbPort <port num> --dbName <db name>
# ... would be the same as ...
./<prog> --config ./Config.txt
```

## Further Reading:

1. [Widgets](./src/widgets/README.md)
1. [Database Package](./src/db/README.md)
