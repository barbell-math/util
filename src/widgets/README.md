# widgets

A type safe way to augment standard types with more information for common
operations such as equality and hashing.

## What is a Widget?

A widget can be though of as a stateless value that implements methods that
perform operations on a given type. Why stateless? Because arbitrary types need
to be supported. This includes any builtin types and builtin types cannot be
modified with additional methods. For example, you cannot add a method to the
builtin int type. The below code will not compile.

```golang
func (i int) Eq(other int) bool { ... }
```

Since the type itself cannot be modified, an external _stateless_ value will be
used in concert with a given value to provide the necessary missing information.
This external value will need to be stateless to avoid any confusion between it
and the actual value. However, it will still need to act on values of the type
it is working with. So, then why not use generics to define a sub type that the
original type can be casted to and from freely? Because go forbids creating
non-instantiated generic types. The code shown below tries to do this. If you
try to compile this code it will fail.

```golang
type Widget[T any] T
```

So, that brings us back to having an external type to agument another type to
"add methods" to it.

The next question mught be "why on earth would you do this?" This is a valid,
question but can be simply answered by looking at the containers package. Each
container needs to know specific information about a given arbitrary type as
well as how to perform basic operations with that type. This is not just for the
containers package either, it is easy to image other packages that could make
use of this idea (algorithims, generic code in general, ...).

## Available Widgets

There are four kinds of widgets, each with an associated interface:

| Widget            | Associated Interface |
|-------------------|----------------------|
| Base              | https://github.com/barbell-math/util/blob/54112c5c3f04919e84cf583b531e7266fa948823/src/widgets/Common.go#L8-L21 |
| PartialOrder      | https://github.com/barbell-math/util/blob/eaa36c190ebe2d5aefe6f27ec3c588f8ea68458b/src/widgets/Common.go#L23-L31 |
| Arith             | https://github.com/barbell-math/util/blob/eaa36c190ebe2d5aefe6f27ec3c588f8ea68458b/src/widgets/Common.go#L33-L61 |
| PartialOrderArith | https://github.com/barbell-math/util/blob/eaa36c190ebe2d5aefe6f27ec3c588f8ea68458b/src/widgets/Common.go#L63-L70 |

Given these widgets, they have been defined to be used like the following:

https://github.com/barbell-math/util/blob/03fed1ba47c9d901ae66c53a1dce2ec2f0859866/src/widgets/Examples_test.go#L9-L19

## Further Reading:

1. [containers](../container/README.md)
