# argparse

A type safe, extensible CLI argument parsing utility package.

https://github.com/barbell-math/util/blob/22385fc610ab0b22b8f16fc828ef3a43b300ceb9/argparse/examples/SimpleExamples_test.go#L52-L72

## Usage

The general usage of this package is as follows:

https://github.com/barbell-math/util/blob/4d4a582428ccf312f8b2faa016626d6f35246b53/argparse/examples/SimpleExamples_test.go#L18-L31
<sup>Example usage of the argparse package</sup>

In this example there are three main parts to consider:

1. A `ArgBuilder` is created and it it populated with arguments. This stage is
where the vast majority of your interaction with the package will take place.
1. The `ArgBuilder` makes a `Parser`. Optionally this is where subparsers could
be added if you were using them.
1. The parser is then used to parse a slice of strings which is translated into
a sequence of tokens.

Steps 2 and 3 are mostly self explanitory, but the options available as step 1
need further explination.

## Design

## Benchmarking
