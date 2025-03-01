# Struct Default Init

A generator program that provides boiler plate code for struct initialization
and basic methods.

## Usage

To execute this program use a go generate command of the following structure:

```
//go:generate <path to exe>/structDefaultInit -struct=<struct type>
```

Given this generate command, the ```structDefaultInit``` program will search the
ast of any non-generated code in the current directory for the supplied struct
type definition. Once the supplied struct type definition is found it will use
the comments on each of the struct fields to generate several pieces of code.
Before going over the code this is generated, the code below shows the expected
struct format:

```
//gen:structDefaultInit newReturns [val | pntr]
type <struct type> struct {
    // This is an example embeded field - they are supported
    //gen:structDefaultInit default <default value 1>
    <field 1>

    // Another comment for this field
    //gen:structDefaultInit default <default value 2>
    //gen:structDefaultInit setter
    //gen:structDefaultInit getter
    <field 2> <field 2 type> `default:"<default value 2>" setter:"t" getter:"t"`
    //gen:structDefaultInit default <default value 3>
    //gen:structDefaultInit setter
    //gen:structDefaultInit imports <imported package 1> <desired name>-><imported package2>
    <field 3> <field 3 type>

    //gen:structDefaultInit default <default value 3>
    //gen:structDefaultInit pointerSetter
    <field 4> *<field 4 type>
}
```

The following comment arguments are supported for the struct definition:

1. newReturns (string) (required): whether or not the generated new function for
the associated struct returns a pointer to the associated struct or the
associated struct as a value.

The following comment arguments are supported for each field:

1. default (string) (required): the value that the field should be initialized 
with. This value is treated as a string, meaning whatever text is in the string
will be what is placed in the generated code. Expressions are not evaluated.
1. setter (bool) (optional): true to make a setter function for this field,
false to not add one
1. pointerSetter (bool) (optional): true to make a pointer setter function for
this field, false to not add one. A pointer setter will assume that the
associated field is a pointer to an underlying value and will set the pointers
underlying value rather than the pointer itself.
1. getter (bool) (required): true to make a getter function for this field,
false to not add one
1. imports (string) (optional): a space separated list of imports to include in 
the generated code file. This is useful when the default value is derived from a
value in an external package. Every import will be automatically wrapped in
quotes. To import something under a different name use the following syntax: 
`<desired name>-><import path>`

Given the example used throughout this file, with the information from the 
inline arguments and the comment args the following code will be generated:

```
package <package>

import (
    "<imported package 1>"
    <desired name> "<imported package 2>"
)

// returns a new <struct type> struct initialized with the default values.
func new<struct type>() <struct type> {
    return <struct type>{
        <field 1>: <default value 1>,
        <field 2>: <default value 2>,
        <field 3>: <default value 3>,
        <field 4>: <default value 4>,
    }
}

// another comment for this field
//
//gen:structDefaultInit default <default value 2>
//gen:structDefaultInit setter
//gen:structDefaultInit getter
func (o *<struct type>) Set<field 2>(v <field 2 type>) *<struct type> {
    o.<field 2>=v
    return o
}

// another comment for this field
//
//gen:structDefaultInit default <default value 2>
//gen:structDefaultInit setter
//gen:structDefaultInit getter
func (o *<struct type>) Get<field 2>() *<field 2 type> {
    return o.<field 2>
}

//gen:structDefaultInit default <default value 3>
//gen:structDefaultInit setter
//gen:structDefaultInit imports <imported package 1> <imported package2>
func (o *<struct type>) Set<field 3>(v <field 3 type>) *<struct type> {
    o.<field 3>=v
    return o
}

//gen:structDefaultInit default <default value 3>
//gen:structDefaultInit pointerSetter
func (o *<struct type>) Set<field 4>PntrVal(v *<field 4 type>) *<struct type> {
    *o.<field 4>=*v
    return o
}
```

At a high level the following code is generated:

1. A ```New<struct type>``` function. This function will return a new struct
with the default values that were specified in the struct tags.
1. Setter functions for any fields that had the setter tag set to true.
1. Pointer setter functions for any fields that had the pointer setter tag set
to true.
1. Getter functions for any fields that had the getter tag set to true.
1. Every getter and setter function will have the same doc string comment as
the comment that was on the associated struct field.
