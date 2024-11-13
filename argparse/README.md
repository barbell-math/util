# argparse

A type safe, extensible CLI argument parsing utility package.

https://github.com/barbell-math/util/blob/22385fc610ab0b22b8f16fc828ef3a43b300ceb9/argparse/examples/SimpleExamples_test.go#L47-L72
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

https://github.com/barbell-math/util/blob/eac51e0e6c5bb37d6e8db26e13e7b4e31e76c475/argparse/examples/SimpleExamples_test.go#L18-L19
<sup>Integer argument</sup>

The above example provides an argument with no options, meaning all the default
options will be used. Shown below is how to provide your own set of options.

https://github.com/barbell-math/util/blob/eac51e0e6c5bb37d6e8db26e13e7b4e31e76c475/argparse/examples/SimpleExamples_test.go#L53-L59
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

https://github.com/barbell-math/util/blob/eac51e0e6c5bb37d6e8db26e13e7b4e31e76c475/argparse/examples/SimpleExamples_test.go#L92-L98

2. Flag counter argument. This will return an integer equal to the number of
times that the flag was provided. It does not accept any values.

https://github.com/barbell-math/util/blob/eac51e0e6c5bb37d6e8db26e13e7b4e31e76c475/argparse/examples/SimpleExamples_test.go#L131-L136

3. List argument. This will build up a list of all the values that were provided
with the argument. Many values can be provided with a single argument or many
flags can be provided with a single argument, as shown in the example arguments
below the argument example.

https://github.com/barbell-math/util/blob/6d29be63cbb049ae12b61a43efaab4d9cd930984/argparse/examples/SimpleExamples_test.go#L171-L186
https://github.com/barbell-math/util/blob/6d29be63cbb049ae12b61a43efaab4d9cd930984/argparse/examples/SimpleExamples_test.go#L192

4. List argument with a predefined set of allowed values. This will build up a
list of all the values that were provided with the argument, provided that they
are in the allowed list of values. Many values can be provided with a single
argument or many flags can be provided with a single argument, as shown in the
example arguments below the argument example.

https://github.com/barbell-math/util/blob/6d29be63cbb049ae12b61a43efaab4d9cd930984/argparse/examples/SimpleExamples_test.go#L211-L231
https://github.com/barbell-math/util/blob/6d29be63cbb049ae12b61a43efaab4d9cd930984/argparse/examples/SimpleExamples_test.go#L239

5. Selector argument. This will accept a single value as long as that values is
in the predefined set of allowed values.

https://github.com/barbell-math/util/blob/6d29be63cbb049ae12b61a43efaab4d9cd930984/argparse/examples/SimpleExamples_test.go#L258-L275

## Argument Builder: Custom Types

Due to using generics, the argument builder can accept arguments of custom types
as long as there is a translator for that type. For examples of how to implement
translators refer to the
[stateless translator example](./examples/CustomStatelessTranslator_test.go)
as well as the 
[stateful translator example](./examples/CustomStatefulTranslator_test.go).

To support a custom type the `Translate` method on the translator will simply
need to return the custom type. Support for custom translators and types allows
for a completely type safe translation of the cmd line arguments into values
that your program can work with.

## Computed Arguments


## Design

## Benchmarking
