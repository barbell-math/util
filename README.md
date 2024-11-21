# util

A golang utility library that aims to add useful features to the std lib.

## Documentation

This package contains many sub-packages that each deserve more space than can
be given in one combined README. Please refer to each packages README for
information about each sub-package:

1. [Iterator framework](./src/iter/README.md)
1. [Containers (a.k.a. data structures)](./src/container/README.md)
1. [CLI argument parser](./src/argparse/README.md)
1. [Widgets](./src/widgets/README.md)

## Code Generator Installation

If your project would benifit from using the code generators defined in this
utility library then you can install them using the command shown below. The
generator executables will be placed in the `$GOPATH/bin` directory.

```
go install github.com/barbell-math/util/generators/...
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
