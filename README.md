# util

A [zero dependency](./go.mod) golang utility library that aims to add useful
features to the std lib.

## Supported Versions and OS's

Supported Go versions: 1.21, 1.22, and 1.23

The latest versions of windows, linux, and macos (as defined by github actions)
are supported.

## Documentation

This package contains many sub-packages that each deserve more space than can
be given in one combined README. Please refer to each packages README for
information about each sub-package:

1. [Iterator framework](./src/iter/README.md)
1. [Containers (a.k.a. data structures)](./src/container/README.md)
1. [CLI argument parser](./src/argparse/README.md)
1. [Widgets](./src/widgets/README.md)
1. [Generators](./generators/README.md)

## Code Generator Installation

If your project would benefit from using the code generators defined in this
utility library then you can install them using the command shown below. The
generator executables will be placed in the `$GOPATH/bin` directory.

```
go install github.com/barbell-math/util/generators/...
```

To install a single generator rather than all of the use the command shown
below.

```
go install github.com/barbell-math/util/generators/<generator name>
```

## Package Install

If your project would benefit from using the this package directly as a library
then you can use the command shown below to add the library to your project.

```
go get github.com/barbell-math/util
```

## Helpful Commands

### Building and Running Generators

```sh
go build -o ./bin/ ./generators/...
go generate ./src/...
```

### Running Unit Tests

```sh
go build -o ./bin/ ./generators/...     # unnecessary if no generated code changed
go generate ./src/...                   # unnecessary if no generated code changed
go test ./src/...
```
