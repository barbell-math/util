# widgets

A type safe way to augment standard types with more information for common
operations such as equality and hashing.

## What is a Widget?

A widget can be though of as a stateless value that implements methods that
perform operations on a given type. Why stateless? Because arbitrary types need
to be supported. This includes any builtin types and builtin types cannot be
modified with additional methods. For example, you cannot add a method to the
builtin int type. The below code will not compile.

```
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

```
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

| Widget | Interface |
|--------|-----------|
| Base   | https://github.com/barbell-math/util/blob/54112c5c3f04919e84cf583b531e7266fa948823/src/widgets/Common.go#L8-L21 |

1. Base: the base widget type that only provides equality, hashing, and zeroing methods.
1. PartialOrder: an extension of the Base widget that adds comparison methods.
1. Arith: an extenstion of the Base widget that add common arithmatic operations.
1. PartialOrderArith: an extenstion of both the ParitialOrder and Arith widgets, combining all of there methods.

Each widget is associated with an interface.

1. Base Widget

2. 

## Further Reading:

1. [containers](../container/README.md)
