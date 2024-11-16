package iter

import (
	"github.com/barbell-math/util/src/customerr"
)

// ForEach is the most ubuquitous consumer, most other consumers can be expressed
// using ForEach making them pseudo-consumers. By using this pattern all
// pseudo-consumers are abstracted away from the complex looping logic.

// This funciton is a consumer.
//
// For each will take values from it's parent iterator and perform the supplied
// operation function on each value, until an error occurs. If an error occurs
// the operation function is not called and iteration stops.
func (i Iter[T]) ForEach(
	op func(index int, val T) (IteratorFeedback, error),
) error {
	j := 0
	f := Continue
	var next T
	var err error
	var cont bool = true
	var opErr error = nil
	for cont && err == nil && opErr == nil && f == Continue {
		next, err, cont = i(Iterate)
		if err == nil && cont {
			f, opErr = op(j, next)
			j++
		}
	}
	_, cleanUpErr, _ := i(Break)
	return customerr.AppendError(err, opErr, cleanUpErr)
}

// Why is stop not a pseudo consumer? It breaks the parent calling convention
// that ForEach uses. For each will always get a parent iterators value before
// the op function is consulted. Stop should just stop, and not call the parent
// iterators one last time meaning it has to be separate.

// This function is a consumer.
//
// Stop will stop all iteration without ever consuming a single value from it's
// parent iterator. Any errors returned will be generated from each iterator
// performing there respective teardown operations.
func (i Iter[T]) Stop() error {
	_, cleanUpErr, _ := i(Break)
	return cleanUpErr
}

// This function is a consumer.
//
// PullOne will pull one value from the iterator chain. It will return the value,
// an error if one occurred while obtaining the value, and a boolean flag to
// indicate success. The value that is returned should be assumed to be invalid
// if err is not nil or the boolean flag is false. This function *does not*
// clean up the iterator chain once the chains producer has reached the end of
// its stream of values. Stop must be called manually to do this. Continuing to
// call PullOne after the end of the iterator stream has been reached will not
// result in any undefined behavior.
func (i Iter[T]) PullOne() (T, error, bool) {
	return i(Iterate)
}

// This function is a consumer.
//
// Pull will pull num values from the iterator chain, until the end of the
// iterator chain has been reached. It will return the values, an error if one
// occurred while obtaining the values, and a boolean flag to indicate success.
// The values that is returned should be assumed to be valid, even if err is not
// nil or the boolean flag is false. This function *does not* clean up the
// iterator chain once the chains producer has reached the end of its stream of
// values. Stop must be called manually to do this. Continuing to call Pull
// after the end of the iterator stream has been reached will not result in any
// undefined behavior.
func (i Iter[T]) Pull(num int) ([]T, error, bool) {
	j := 0
	rv := make([]T, num)
	var next T
	var err error
	var cont bool = true
	for cont && err == nil && j < num {
		next, err, cont = i(Iterate)
		if err == nil && cont {
			rv[j] = next
			j++
		}
	}
	return rv, err, cont
}
