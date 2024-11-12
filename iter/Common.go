// This package defines an iterator framework comprized of lazily evaluated,
// pull style iterators.
package iter

// A type that defines the valid states that an iterator chain can use.
type IteratorFeedback int

const (
	Continue IteratorFeedback = iota // Signaling to continue iteration, consume the next value
	Break                            // Signaling to stop iteration. The current value is no longer valid.
	Iterate                          // Signaling to iterate again, get the next value from the iterator.
)

// Iter is the base type that the entire package is built from. This type defines
// the function that has methods defined on it such that they can be chained
// together to form iterator chains. The returned values are as follows:
//  - T: the value that will be produced in the iterator sequence
//  - error: any error that was generated when attempting to get the next value
// in the iterator sequence
//  - bool: a flag to indicate whether or not to continue iteration
type Iter[T any] func(f IteratorFeedback) (T, error, bool)
