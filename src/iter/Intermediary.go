package iter

import (
	"github.com/barbell-math/util/src/customerr"
)

// This function is an intermediary.
//
// Next will take it's parent iterator and consume its values. As it does this
// it will apply the operation (op) function to the value before passing on the
// transformed value to it's child iterator. If an error is generated iteration
// will stop.
func Next[T any, U any](i Iter[T],
	op func(index int, val T, status IteratorFeedback) (IteratorFeedback, U, error),
) Iter[U] {
	j := 0
	return func(f IteratorFeedback) (U, error, bool) {
		var tmp U
		next, err, cont := i(f)
		if f == Break {
			_, _, err2 := op(j, next, Break) //Clean up current iterator
			return tmp, customerr.AppendError(err, err2), false
		}
		var opErr error = nil
		for ; cont && err == nil && opErr == nil && f == Iterate; next, err, cont = i(f) {
			f, tmp, opErr = op(j, next, f)
			j++
			if f == Continue || f == Break || opErr != nil {
				return tmp, opErr, (opErr == nil && f == Continue)
			}
		}
		return tmp, err, cont
	}
}

// This function is an intermediary.
//
// This function is equivalent to [Next], the only difference is that the
// the inputs iterator type and output iterator type must be the same. It is
// offered as a convenience function.
func (i Iter[T]) Next(
	op func(index int, val T, status IteratorFeedback) (IteratorFeedback, T, error),
) Iter[T] {
	return Next(i, op)
}

// This function is an intermediary.
//
// SetupTeardown provides a way to have setup and teardown procedures. These
// setup and teardown procedures will be called once before the parent iterator
// is ever called and after once after the parent iterator has completed. The
// teardown procedure will always be called even if the parent iterator returned
// an error.
func (i Iter[T]) SetupTeardown(setup func() error, teardown func() error) Iter[T] {
	j := -1
	return func(f IteratorFeedback) (T, error, bool) {
		j++
		if f != Break && j == 0 {
			if err := setup(); err != nil {
				var tmp T
				return tmp, err, false
			}
		}
		if f == Break {
			val, err, cont := i(f)
			if j > 0 {
				err = customerr.AppendError(err, teardown())
			}
			return val, err, cont
		}
		return i(f)
	}
}

// This function is an intermediary.
//
// Inject will inject values into the iterator stream based on the operation (op)
// functions return values. If the operation function returns true then the value
// will be injected and the parent iterators current value will be cached to be
// returned as soon as the operation function returns false. This provides a way
// to arbitrarily change the sequence of values that is output. If an error is
// returned from either the parent iterator or the operation function
// then iteration will stop.
func (i Iter[T]) Inject(
	op func(idx int, val T, injectedPrev bool) (T, error, bool),
) Iter[T] {
	j := -1
	injected := false
	var prevErr error
	var prevCont bool
	var prevVal T
	return func(f IteratorFeedback) (T, error, bool) {
		if f == Break {
			var tmp T
			return tmp, nil, false
		}
		j++
		var next T
		var err error
		var cont bool
		if injected {
			next = prevVal
			err = prevErr
			cont = prevCont
		} else {
			next, err, cont = i(f)
		}
		if v, opErr, status := op(j, next, injected); status {
			if !injected {
				prevVal = next
				prevCont = cont
				prevErr = err
				injected = true
			}
			return v, opErr, (opErr == nil)
		} else {
			injected = false
			return next, err, cont
		}
	}
}
