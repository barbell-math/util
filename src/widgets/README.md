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

So, that brings us back to having an external type to augment another type to
"add methods" to it.

The next question might be "why on earth would you do this?" This is a valid,
question but can be simply answered by looking at the containers package. Each
container needs to know specific information about a given arbitrary type as
well as how to perform basic operations with that type. This is not just for the
containers package either, it is easy to image other packages that could make
use of this idea (algorithms, generic code in general, ...).

## Available Widgets

There are four kinds of widgets, each with an associated interface:

| Widget            | Associated Interface |
|-------------------|----------------------|
| Base              | https://github.com/barbell-math/util/blob/54112c5c3f04919e84cf583b531e7266fa948823/src/widgets/Common.go#L8-L21 |
| PartialOrder      | https://github.com/barbell-math/util/blob/eaa36c190ebe2d5aefe6f27ec3c588f8ea68458b/src/widgets/Common.go#L23-L31 |
| Arith             | https://github.com/barbell-math/util/blob/eaa36c190ebe2d5aefe6f27ec3c588f8ea68458b/src/widgets/Common.go#L33-L61 |
| PartialOrderArith | https://github.com/barbell-math/util/blob/eaa36c190ebe2d5aefe6f27ec3c588f8ea68458b/src/widgets/Common.go#L63-L70 |

The widgets interfaces are built off of each other, meaning that they can be
down casted as needed.

| Widget Interface           | Allowed Down Cast Targets |
|----------------------------|---------------------------|
| BaseInterface              | |
| PartialOrderInterface      | BaseInterface |
| ArithInterface             | BaseInterface |
| PartialOrderArithInterface | BaseInterface, PartialOrderInterface, ArithInterface |

Besides these widget types, all builtin types have an analogous widget defined
within this package. All of these widgets will follow the `Builtin<type>` naming
format.

## Example Usage

The previously established widget types have been defined to be used like the
following examples:

1. Base widget example

https://github.com/barbell-math/util/blob/03fed1ba47c9d901ae66c53a1dce2ec2f0859866/src/widgets/Examples_test.go#L8-L18

2. Partial order widget example

https://github.com/barbell-math/util/blob/b49d7d40cf8df570522ac49ea0903c30d89cd990/src/widgets/Examples_test.go#L22-L27

3. Arith widget example

https://github.com/barbell-math/util/blob/b49d7d40cf8df570522ac49ea0903c30d89cd990/src/widgets/Examples_test.go#L31-L48

## Widgets on Custom Types

Custom user defined types do not suffer from the problem of not being able to
add methods. As such, a custom type may choose to implement the widget interface
on itself. In this case, the type and the stateless type collapse into one.
However, they must still be considered to be separate. The method receiver must
not be used in any implementation where a type implements the widget interface
on itself. Doing so would violate the stateless nature that predicated the
design of widgets. All values that are needed to perform the requested
operations will be passed as arguments to the required methods.

The example below shows a custom type implementing the partial order widget
interface on itself. Note how the method receivers are not used and the `_`
character is used as the method receiver. This indicates that it should not be
used and is a recommended pattern to follow.

https://github.com/barbell-math/util/blob/b49d7d40cf8df570522ac49ea0903c30d89cd990/src/widgets/CustomWidget_test.go#L10-L26
<sup>A custom type implementing the partial order widget on itself.</sup>

The example below demonstrates how to use the widget associated with this custom
type. This example also demonstrates how the custom type can be used with both
the base and partial order widgets. This is due to the down casting rules that
were discussed in the previous section. Given this behavior, it is recommended
that every custom type implements the largest allowable widget interface which
allows it to be used in the greatest number of scenarios.

https://github.com/barbell-math/util/blob/b49d7d40cf8df570522ac49ea0903c30d89cd990/src/widgets/CustomWidget_test.go#L50-L63

## What About Pointers??

Pointers are a common thing to use. As such, there needs to be a way for a
widget to perform it's operations on the underlying value rather than the
pointer value. To do this the pointer widgets can be used as wrappers to peel
back the pointer and return the result of the required operation on the
underlying value. There are pointer widgets for each of the widget types.

1. BasePntr
1. PartialOrderPntr
1. ArithPntr
1. PartialOrderArithPntr

The example below shows how to use these pointer widgets.

https://github.com/barbell-math/util/blob/b49d7d40cf8df570522ac49ea0903c30d89cd990/src/widgets/CustomWidget_test.go#L43-L56

## Further Reading:

1. [containers](../container/README.md)
