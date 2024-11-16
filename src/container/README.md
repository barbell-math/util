# container

A type safe collection of generic data structures. Within this package there are
several sub-packages. They will all be referenced here.

1. `containers`: The concrete implementation of all the generic containers. This
is where all the heavy lifting is done.
1. `staticContainers`: Defines interfaces that define what it means to be
various kinds of static containers. Static is meant in the sense that the 
containers won't reallocate once created, their capicity is fixed.
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

The following containers are available in the `containers` package:

| Name             | Category | Desc |
|------------------|----------|------|
| `Pair`           | Baisc    | A pair of values. |
| `Triple`         | Baisc    | A triplet of values. |
| `WidgerPair`     | Basic    | A super set of `Pair` that implements the widget interface. |
| `WidgerTriple`   | Basic    | A super set of `Triple` that implements the widget interface. |
| `Variant`        | Baisc    | Represents a value that can be one of two values of differing types. |
| `CircularBuffer` | Static   | A static array of values that wrap around as values are added/removed. Creates efficient queue and stack operations. |
| `Vector`         | Dynamic  | A wrapper for a slice that implements the necessary interfaces. |
| `HashSet`        | Dynamic  | A hash set that can contain any(!) type as long as there is a widget interface for it. |
| `HashMap`        | Dynamic  | A hash map that can use any(!) type for keys as long as there is a widget interface for it. |
| `HashGraph`      | Dynamic  | graph data structure that relies on hashing to create efficient access and modifications to the graph structure. |
| `HookedHashSet`  | Dynamic  | A super set of a `HashSet` that provides callbacks for when hashes are being updated internally in the hash set. Mainly used for efficiency gains in other data structures. |

## Static and Dynamic Interfaces

All data structures in the `containers` package will implement one, if not
several, of the interfaces listed below.

| Name | Category(s) | Desc |
|------|-------------|------|

Each of these interfaces have been established to provide consistency between
the containers and allow for interface types to be passed in code rather than
the explicit concrete container types, giving the code greater flexibility. Each
of the interfaces has also been broken down into read only/write only parts. The
intent of this was to allow for the mutibility and access of the containers to
be controlled in any code that uses the interface values.