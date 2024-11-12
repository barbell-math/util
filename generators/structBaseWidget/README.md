# Struct Base Widget

A generator program that creates methods for a given struct that implement the
base widget interface.

## Usage

To execute this program use a go generate command of the following structure:

```
//go:generate <path to exe>/structBaseWidget -type=<type>
```

Given this go generate command, the ```structBaseWidget``` program will perform
one task.

This task is to search the ast of any non-generated code in the current
directory for the supplied struct types definition. Once this definition is
found, it is expected that at least one of the fields of the struct will contain
several more arguments in the doc string of each field.

```
type (
    <type> struct {
		//gen:structBaseWidget identity
		//gen:structBaseWidget baseTypeWidget <base type widget 1>
		//gen:structBaseWidget widgetPackage <base type widget package 1>
        <field 1>
		//gen:structBaseWidget identity
		//gen:structBaseWidget baseTypeWidget <base type widget 2>
		//gen:structBaseWidget widgetPackage <base type widget package 2>
        <field 2>
        <field 3>
    }
)
```

The arguments are as follows:

1. identity (falg) (required): include this flag if the associated field should
be considered to be part of the structs identity.
2. baseTypeWidget (string) (required if identity is provided): the widget that
should be associated with the struct field.
3. widgetPackage (string) (required if identity is provided): the package the
base type widget is in. Use '.' if it is in the current package.

Given the example used throughout this file, with the information from the
inline arguments and the doc string arguments, the following code would be
generated.

```
package <package>

import (
    <base type widget package 1>
    <base type widget package 2>
	"github.com/barbell-math/util/hash"
)

// Returns true if l equals r using the widgets associated with the following
// fields:
//   - <field 1>
//   - <field 2>
func (_ *<type>) Eq(l *<type>, r *<type>) bool {
	var (
		<field 1>Widget <base type widget 1>
		<field 2>Widget <base type widget 2>
		rv          bool = true
	)

	rv = rv && <field 1>Widget.Eq(&l.value, &r.value)
	rv = rv && <field 2>Widget.Eq(&l._type, &r._type)
	return rv
}

// Returns a hash that represents other by calling [hash.CombineIgnoreZero] with
// the values generated from using the widgets associated with the following
// fields:
//   - <field 1>
//   - <field 2>
func (_ *token) Hash(other *<type>) hash.Hash {
	var (
		<field 1>Widget <base type widget 1>
		<field 2>Widget <base type widget 2>
		rv          hash.Hash
	)

	rv = rv.CombineIgnoreZero(<field 1>Widget.Hash(&other.value))
	rv = rv.CombineIgnoreZero(<field 2>Widget.Hash(&other._type))
	return rv
}

// Zeros the following fields in the struct using the widgets associated with
// the following fields:
//   - <field 1>
//   - <field 2>
func (_ *token) Zero(other *<type>) {
	var (
		<field 1>Widget <base type widget 1>
		<field 2>Widget <base type widget 2>
	)

	valueWidget.Zero(&other.<field 1>)
	_typeWidget.Zero(&other.<field 2>)
}
```
