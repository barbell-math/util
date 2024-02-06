package iter

// A type that defines the valid states that an iterator chain can use.
type IteratorFeedback int;
const (
    Continue IteratorFeedback=iota // Signaling to continue iteration, consume the next value
    Break // Signaling to stop iteration. The current value is no longer valid.
    Iterate // Signaling to iterate again, get the next value from the iterator.
);

// Iter is the base type that the entire package is built from. This type defines
// the function that has methods defined on it such that they can be chained
// together to form iterator chains.
type Iter[T any] func(f IteratorFeedback)(T,error,bool)
