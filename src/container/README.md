# container

A type safe collection of generic data structures. Within this package there are
several sub-packages. They will all be referenced here.

1. `containers`: The concrete implementation of all the generic containers. This
is where all the heavy lifting is done.
1. `staticContainers`: Defines interfaces that define what it means to be
various kinds of static containers. Static is meant in the sense that the 
containers won't reallocate once created, their capacity is fixed.
1. `dynamicContainers`: Defines interfaces that define what it means to be
various kinds of dynamic containers. Dynamic is meany in the sense that the
containers will reallocate as needed, their capacity will grow or shrink.
1. `containerTypes`: The set of sub-types that make up the types in the
`staticContainers` and `dynamicContainers` packages.
1. `basic`: Implementations of simple generic containers that do not conform to
the interfaces defined in the `staticContainers` and `dynamicContainers`
packages. This package is intended to be able to be imported anywhere without
creating import cycles.

## Available Containers

The following containers are available in the `containers` package. For brevity
the synchronized version of each container is not listed, but all containers
with a `*` at the end of their name also have a synchronized version. The
synchronized version of the container will always follow the naming convention
of `Synced<container name>`.

| Name              | Category | Desc |
|-------------------|----------|------|
| `Pair`            | Basic    | A pair of values. |
| `Triple`          | Basic    | A triplet of values. |
| `WidgetPair`      | Basic    | A super set of `Pair` that implements the widget interface. |
| `WidgetTriple`    | Basic    | A super set of `Triple` that implements the widget interface. |
| `Variant`         | Basic    | Represents a value that can be one of two values of differing types. |
| `CircularBuffer*` | Static   | A static array of values that wrap around as values are added/removed. Creates efficient queue and stack operations. |
| `Vector*`         | Dynamic  | A wrapper for a slice that implements the necessary interfaces. |
| `HashSet*`        | Dynamic  | A hash set that can contain any(!) type as long as there is a widget interface for it. |
| `HashMap*`        | Dynamic  | A hash map that can use any(!) type for keys as long as there is a widget interface for it. |
| `HashGraph*`      | Dynamic  | graph data structure that relies on hashing to create efficient access and modifications to the graph structure. |
| `HookedHashSet*`  | Dynamic  | A super set of a `HashSet` that provides callbacks for when hashes are being updated internally in the hash set. Mainly used for efficiency gains in other data structures. |

## Static and Dynamic Interfaces

All data structures in the `containers` package will implement one, if not
several, of the interfaces listed below.

| Name   | `staticContainers` Interface | `dynamicContainers` Interface |
|--------|------------------------------|-------------------------------|
| Set    | &check;                      | &check;                       |
| Map    | &cross;                      | &check;                       |
| Queue  | &check;                      | &check;                       |
| Stack  | &check;                      | &check;                       |
| Deque  | &check;                      | &check;                       |
| Vector | &check;                      | &check;                       |
| Graph  | &cross;                      | &check;                       |

Each of these interfaces have been established to provide consistency between
the containers and allow for interface types to be passed in code rather than
the explicit concrete container types, giving the code greater flexibility. Each
of the interfaces has also been broken down into read only/write only parts. The
intent of this is to allow for the mutability and access of the containers to be
controlled in any code that uses interface values.

https://github.com/barbell-math/util/blob/9948c0045f7246acb5c2827ea4b34cb4919fcae9/src/container/staticContainers/Set.go#L5-L17
<sup>Read-only static set interface.</sup>

https://github.com/barbell-math/util/blob/9948c0045f7246acb5c2827ea4b34cb4919fcae9/src/container/staticContainers/Set.go#L19-L30
<sup>Write-only static set interface.</sup>

https://github.com/barbell-math/util/blob/9948c0045f7246acb5c2827ea4b34cb4919fcae9/src/container/staticContainers/Set.go#L32-L37
<sup>Full static set interface.</sup>

An example of using the interface types over the concrete types is shown in the
`SlidingWindow` function implementation. The function expects a container to be
passed to it that implements _both_ the static queue and static vector
interfaces and then it only returns the vector interface component of the same
value. This was done because the sliding window operations require efficient
queue operations, but all down stream iterators only need random access to the
elements within the container and not access to any of the queue operations.

https://github.com/barbell-math/util/blob/9948c0045f7246acb5c2827ea4b34cb4919fcae9/src/container/containers/Iterface.go#L22-L29
<sup>The sliding window function definition.</sup>

## Addressability

Any go programmer should be familiar with the
[idea behind addressability](https://go.dev/ref/spec#Address_operators). Given
how intertwined this idea is with the standard go slices and maps, it should
make sense that the concrete implementations of the containers in this package
are beholden to similar addressability constraints.

Any container that uses a map internally will not be addressable due to the
underlying properties of the map. Yet, these containers will still implement all
the same methods as ones who are addressable because they still need to
implement the appropriate static and dynamic container interfaces. Some methods
that are defined in these interfaces define return values that are pointers to
the data within the container. This presents a tension between the
addressability constraints imposed by go and the interface definitions. This is
likely the result of shortsighted design choices early on in the making of this
package. As a result the following conventions have been established:

1. Any container that is not addressable will panic when a method is called that
returns a pointer to the data within the container.
1. The doc string will clearly state if the method panics or not.

This is not the best decision, but the following counter ideas seem worse:

1. It might seem like a natural idea to completely remove the ability to return
pointers to the data within the container, but containers like `Vector` would be
severely impaired by removing that ability.
1. The static and dynamic container interfaces could be further divided into
addressable/non-addressable versions. This however only further complicates the
API exposed by those packages. This complication may seem trivial but it
complicates any users from writing generalized code because any generalized
code will not know if any application specific code will need to the
addressability or not.

## Value Initializers

Almost all of the containers in the `containers` package should be initialized
with their respective new function. However, if you just want to create a
container from a static list of values then the value initializers can be used.
The value initializers will take a sequence of varargs and return a container
initialized and containing those values.

https://github.com/barbell-math/util/blob/9948c0045f7246acb5c2827ea4b34cb4919fcae9/src/container/containers/Vector_test.go#L46-L48
<sup>Value initializer example with a Vector.</sup>

## Vector: The Special Case

The `Vector` container is slightly different from the other containers. It is
nothing more than a sub type of a slice.

https://github.com/barbell-math/util/blob/9948c0045f7246acb5c2827ea4b34cb4919fcae9/src/container/containers/Vector.go#L16-L23
<sup>The `Vector` type definition.</sup>

Due to this vectors can be type casted to and from a slice.

https://github.com/barbell-math/util/blob/9948c0045f7246acb5c2827ea4b34cb4919fcae9/src/container/containers/Vector_test.go#L33-L42
<sup>Example of type casting a slice and vector.</sup>

This is convenient for sure, but it allows for some other more interesting
properties. As was stated in the comment of the above code example, the widget
type information is lost when performing this transformation. This opens the
door to _changing_ the widget type with a sequence of type casting operations.
This may seem like a strange thing, but it actually can be used to modify how
the slice views it's data. Methods like `Contains` could return different values
depending on the equality operation defined by the different widget types. This
is actually used in the `HashGraph` implementation. Shown below are two sub
types of the underlying `graphLink` type.

https://github.com/barbell-math/util/blob/9948c0045f7246acb5c2827ea4b34cb4919fcae9/src/container/containers/HashGraph.go#L34-L37

As explained in the comments, each sub type is used in different scenarios
depending on what part of the `graphLink` is relevant. Each of these sub types
implements the `BaseWidget` interface as shown below. Notice how the equality
and hash operations depend on different attributes from the underlying
`graphLink` type.

https://github.com/barbell-math/util/blob/9948c0045f7246acb5c2827ea4b34cb4919fcae9/src/container/containers/HashGraph.go#L108-L126

Given these two types the code later is able to type cast a
`Vector[graphLink, *graphLink]` to `[]graphLink` and then to 
`Vector[graphLink, *vertexOnlyGraphLink]`. This changes the underlying equality
function, which is then subsequently used by the `Pop` method on the vector.

https://github.com/barbell-math/util/blob/9948c0045f7246acb5c2827ea4b34cb4919fcae9/src/container/containers/HashGraph.go#L1292-L1295

Kinda cool if you ask me.

Now, this only applies to `Vector`. No other container can do this. The reason
for this is because most other containers rely on a singular equality and hash
definition. Imagine if a hash map suddenly had the hash and equality operations
changed, all the data within the map would be jumbled up and inaccessible.
Vectors are unique in the sense that they are nothing more than a continuous
block of memory. The structure of the container does not depend on anything
else, allowing for the type casting operations explained above to be possible.

## Helpful Design Pattern: Sub-Typing Pointers

A similar idea to the previous section is sub typing pointers. The `argparse`
package uses this concept. If you look at the `argparse.Parser` type definition
you will see the following containers being used.

https://github.com/barbell-math/util/blob/9948c0045f7246acb5c2827ea4b34cb4919fcae9/src/argparse/Parser.go#L37-L48

These containers hold `*shortArg` and `*longArg` values. However, if you look at
the definition of those types you will find thy are sub types of `*Arg`.

https://github.com/barbell-math/util/blob/9948c0045f7246acb5c2827ea4b34cb4919fcae9/src/argparse/Arg.go#L56-L59

Similar to the previous section the only difference between these sub types are
the attributes of the `Arg` type that define equality and hashing.

https://github.com/barbell-math/util/blob/9948c0045f7246acb5c2827ea4b34cb4919fcae9/src/argparse/Arg.go#L110-L129

So, what you end up with are two separate `HashMap`'s that point to the same set
of underlying values. The only difference is how those hash maps _view_ the
underlying values. To add the same value to both hash maps one simply has to
type cast a `*Arg` as shown below.

https://github.com/barbell-math/util/blob/9948c0045f7246acb5c2827ea4b34cb4919fcae9/src/argparse/ArgBuilder.go#L181-L188

This design pattern provides a simple way to create many indexes on a single set
of underlying data.

## Further Reading:

1. [Widgets](../widgets/README.md)
1. [argparse](../argparse/README.md)
