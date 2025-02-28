# generators

A package that provides several code generators that are meant to be used with
the rest of the code in the [src dir](../src/) and any projects that use this
utility package as a library.

https://github.com/barbell-math/util/blob/268ed2b9941d535decf206ee4e49b2442bba1262/src/reflect/StructHash.go#L11-L13
https://github.com/barbell-math/util/blob/268ed2b9941d535decf206ee4e49b2442bba1262/src/reflect/StructHash.go#L16-L18
https://github.com/barbell-math/util/blob/268ed2b9941d535decf206ee4e49b2442bba1262/src/reflect/StructHash.go#L24-L78

## Code Generator Installation

If your project would benefit from using the code generators defined in this
package then you can install them using the command shown below. The generator
executables will be placed in the `$GOPATH/bin` directory.

```
go install github.com/barbell-math/util/generators/...
```

To install a single generator rather than all of the use the command shown
below.

```
go install github.com/barbell-math/util/generators/<generator name>
```

## Usage

As the purpose of each generator is slightly different please refer to the
generator readme that pertains to your needs.

1. [enum](./enum/README.md)
1. [flags](./flags/README.md)
1. [ifaceImplChec](./ifaceImplCheck/README.md)
1. [passThroughWidget](./passThroughWidget/README.md)
1. [structBaseWidget](./structBaseWidget/README.md)
1. [structDefaultInit](./structDefaultInit/README.md)
1. [widgetInterfaceImpl](./widgetInterfaceImpl/README.md)
1. [clean](./clean/README.md)

## General Usage

In general each generator program takes arguments and uses them in combination
with the AST to generate code. Each generator has three potential sources of
arguments, listed below. Each generator program is not limited to any one single
source of input arguments.

1. Cmd-Line Arguments: These are arguments that are passed to the program
through the go:generate line.
1. Comment Arguments: These are arguments that are passed to the program through
the comment directly above a specific value in the AST (usually a type
definition.) These take the form of `//gen:<generator name> <arg name> <arg value>`.
Due to the line not having a space after the double slash any comment arguments
_will not_ be picked up as part of the go doc string.
1. Struct tag arguments: If a struct is being used it may place some arguments
in the struct field tags.

Using a generator is as simple as providing the proper `go:generate` command
and providing the correct arguments in the form of comment arguments or struct
field tag arguments. Refer to each generators readme for the arguments that it
expects.


## Helpful Tid-Bits:

1. If a generator program fails it will return a non-zero exit code, which
will stop any further code generation.
1. The code output from all generators is run through go fmt so you will not
have to worry about formatting issues if you decide to commit the generated code
to version control.
1. All generated code will either have a `.gen.go` or `.gen_test.go` file
extensions. If you don't want to check in any generated code then add the below
lines to your top level git ignore.

```
*.gen.go
*.gen_test.go
```

1. The `clean` generator is a bit different. It will recursively delete all
generated code files as identified by the file extensions listed in the
previous bullet point. Due to golang not guaranteeing the order of execution of
any `go:generate` commands it is not recommended to use this directly in your
code, but instead as a helper program that you can manually run.
1. If you have generated code commited to version control it may be helpful to
setup a mergegate pipeline step that checks that the generated code that is
checked into version control is the same code that is generated from the
generators. This will prevent any accidental cases where generated code is
manually modified an commited. To do this you can use the script below. The
script assumes that the generators are available and in the path and that it is
run at the root directory of your repository.

```sh
clean
go generate -v ./...
if [[ -z "$(git diff)" ]]; then
  echo "No changes detected"
else
  echo "Changes detected!"
  echo "Fix generated code formatting to match what the generators produce!"
  exit 1
fi
```
