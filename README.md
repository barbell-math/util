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
