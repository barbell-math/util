# argparse

A type safe, extensible CLI argument parsing utility package.

https://github.com/barbell-math/util/blob/23d18a281acb287fe6253b150bae98433da10419/src/argparse/examples/SimpleExamples_test.go#L49-L74
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

https://github.com/barbell-math/util/blob/23d18a281acb287fe6253b150bae98433da10419/src/argparse/examples/SimpleExamples_test.go#L20-L21
<sup>Integer argument</sup>

The above example provides an argument with no options, meaning all the default
options will be used. Shown below is how to provide your own set of options.

https://github.com/barbell-math/util/blob/23d18a281acb287fe6253b150bae98433da10419/src/argparse/examples/SimpleExamples_test.go#L55-L61
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
1. `conditionallyRequired`: A way to conditionally require other arguments. See
the Argument Conditionally section for a more in dept explanation.

The available argument types are as follows:

1. `ValueArgType`: Represents a flag type that must accept a single value as an
argument and must only be supplied once.
1. `MultiValueArgType`: Represents a flag type that can accept many values as an
argument. At least one argument must be supplied.
1. `FlagArgType`: Represents a flag type that must not accept a value and must
only be supplied once.
1. `MultiFlagArgType`: Represents a flag type that must not accept a value and
may be supplied many times.

## Argument Builder: Out of the Box Argument Types

The `ArgBuilder` also has several helper functions and translators for common
CLI argument types.

1. Flag arguments. This will return true if the flag is provided. It does accept
any values.

https://github.com/barbell-math/util/blob/23d18a281acb287fe6253b150bae98433da10419/src/argparse/examples/SimpleExamples_test.go#L94-L100

2. Flag counter argument. This will return an integer equal to the number of
times that the flag was provided. It does not accept any values.

https://github.com/barbell-math/util/blob/23d18a281acb287fe6253b150bae98433da10419/src/argparse/examples/SimpleExamples_test.go#L133-L138

3. List argument. This will build up a list of all the values that were provided
with the argument. Many values can be provided with a single argument or many
flags can be provided with a single argument, as shown in the example arguments
below the argument example.The `ListValue` translator can work with any type as
long as it has a translator.

https://github.com/barbell-math/util/blob/23d18a281acb287fe6253b150bae98433da10419/src/argparse/examples/SimpleExamples_test.go#L171-L188
https://github.com/barbell-math/util/blob/23d18a281acb287fe6253b150bae98433da10419/src/argparse/examples/SimpleExamples_test.go#L193

4. List argument with a predefined set of allowed values. This will build up a
list of all the values that were provided with the argument, provided that they
are in the allowed list of values. Many values can be provided with a single
argument or many flags can be provided with a single argument, as shown in the
example arguments below the argument example. The `ListValue` translator can
work with any type as long as it has a translator.

https://github.com/barbell-math/util/blob/23d18a281acb287fe6253b150bae98433da10419/src/argparse/examples/SimpleExamples_test.go#L212-L232
https://github.com/barbell-math/util/blob/23d18a281acb287fe6253b150bae98433da10419/src/argparse/examples/SimpleExamples_test.go#L239

5. Selector argument. This will accept a single value as long as that value is
in the predefined set of allowed values. The `Selector` translator can work with
any type as long as it has a translator.

https://github.com/barbell-math/util/blob/23d18a281acb287fe6253b150bae98433da10419/src/argparse/examples/SimpleExamples_test.go#L258-L278

6. File argument. This will accept a single string value and verify that the
supplied string is a path that exists as a file.

https://github.com/barbell-math/util/blob/23d18a281acb287fe6253b150bae98433da10419/src/argparse/examples/SimpleExamples_test.go#L396-L401

7. Directory argument. This will accept a single string value and verify that
the supplied string is a path that exists as a directory.

https://github.com/barbell-math/util/blob/23d18a281acb287fe6253b150bae98433da10419/src/argparse/examples/SimpleExamples_test.go#L425-L430

8. File open argument. This will accept a single string value and will attempt
to make the supplied file with the given file mode and permissions.

https://github.com/barbell-math/util/blob/23d18a281acb287fe6253b150bae98433da10419/src/argparse/examples/SimpleExamples_test.go#L454-L466

9. Mkdir argument. This will accept a single string value and will attempt to
make the directory(s) that are denoted by the string value.

https://github.com/barbell-math/util/blob/23d18a281acb287fe6253b150bae98433da10419/src/argparse/examples/SimpleExamples_test.go#L491-L499

10. Enum argument. This will accept a single string value and will attempt to
translate it to the underlying enum value given the enum type it was supplied
with through the generic parameters.

https://github.com/barbell-math/util/blob/23d18a281acb287fe6253b150bae98433da10419/src/argparse/examples/SimpleExamples_test.go#L302-L310


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

https://github.com/barbell-math/util/blob/23d18a281acb287fe6253b150bae98433da10419/src/argparse/examples/ConditionallyRequiredArgs_test.go#L30-L44

In the above example the `uint` argument expects that if it is provided that the
`int` and `float` arguments are also provided. However, if the `uint` argument
is not provided then the `int` and `float` arguments do not have to be provided.
Hence, the `uint` argument conditionally requires the `int` and `float`
arguments.

The example below shows how to add conditional arguments based on the value of
the argument that is being added. It is very similar, the only difference is the
`ArgSupplied` function was swapped for a closure provided by the `ArgEquals`
function.

https://github.com/barbell-math/util/blob/23d18a281acb287fe6253b150bae98433da10419/src/argparse/examples/ConditionallyRequiredArgs_test.go#L114-L129

The previous two examples only had one conditionality rule, but a list of rules
can be provided. This allows different sets of arguments to be conditionally
required based on differing requirements (i.e. different enum values.) If
multiple conditionality rules are found to match then the resulting required
argument set will be the union of the required argument sets from all the
matching rules.

> Gotcha!
> Combining default values and conditionally required arguments can sometimes
> lead to seemingly weird behavior. Remember this: if a default value was
> provided then that value will be used when evaluating the argument
> conditionality rules. As such, if a default value matches a rule then the
> arguments specified by that rule will be required. This may seem weird but
> again referring to the motivating example of having an enum value control
> program behavior, it is perfectly reasonable to set a default action. That
> default action should still impose the argument conditionality as if it were
> provided on the CMD line.

## Argument Builder: Computed Arguments

Computed arguments provide a way for the argument parser to set values that were
not directly provided by the CLI, potentially computing values based on the
provided CLI arguments. The example below shows how to add a computed argument
to the parser.

https://github.com/barbell-math/util/blob/23d18a281acb287fe6253b150bae98433da10419/src/argparse/examples/SimpleExamples_test.go#L352-L355

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

https://github.com/barbell-math/util/blob/23d18a281acb287fe6253b150bae98433da10419/src/argparse/testData/ValidConfigFile.txt#L1-L17
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

A more practical example of using a config file is shown using the `db` packages
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

> Comments: The config file format also supports comments. Comments start with
> `//` and _must be on a line of their own_. Comments placed at the end of a
> line will be considered to be part of the arguments value. This was done
> because there is no way to reliably differentiate between a value and comment.
> A value can be anything, so there is no sequence of characters that could
> always define the start of a comment.

> Blank lines: Blank lines will be ignored. Feel free to add them as needed to
> increase the readability of your config files.

## Motivation for Making this Package

Other CMD line argument parsers exist. However none of them did what I wanted.
What I wanted was two fold:

1. A type safe parser that uses generics. No `Parse[Int|Bool|Float]` nonsense.
The type should be provided through generics. This opens the door to support
custom types rather than only supporting the types builtin to the language. So
ideally, the argument parser is extensible enough to support custom types.
1. An argument parser that _fully_ validates and sets up the state of my
program. This point requires more explanation.

Most other argument parsers provide basic validation of CMD line arguments by
attempting to coerce the value to a native type in the language, but what if I
wanted to do something outside of just parsing an integer? What if I wanted to
do something as simple as make sure that integer is within a predefined range?
That would require additional validation logic _after_ the argument parser
supposedly finished validating the input. That additional setup should be able
to be performed by the argument parser.

There is another concern of fully setting the applications state. A motivating
example is a database connection. The CMD line arguments define the connection
parameters, why does more work need to be done _after_ the argument parser is
already done creating values to finalize the applications state? Instead, have
the argument parser be smart enough to be able to create the database connection
once it is done parsing all of the supplied arguments.

By making an argument parser that is extensible enough to _fully_ validate and
_fully_ set the applications state it becomes easier to reason about the
programs execution. The argument parser will either succeed and you will know
that your application will be good to immediately start running with _zero_
additional logic or it will fail and your program will exit before even
starting.

## Further Reading:

1. [Widgets](./src/widgets/README.md)
1. [Database Package](./src/db/README.md)
