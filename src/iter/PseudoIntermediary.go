package iter

// This function is an intermediary.
//
// Take will consume the first num elements of it's parent iterator. It will
// stop iteraton after the first num elements have been consumed. If an error
// occurs iteraton will stop regardless of if num elements have been consumed.
func (i Iter[T]) Take(num int) Iter[T] {
	cntr := 0
	return i.Next(
		func(index int, val T, status IteratorFeedback) (IteratorFeedback, T, error) {
			if status != Break && cntr < num {
				cntr++
				return Continue, val, nil
			}
			return Break, val, nil
		},
	)
}

// This function is an intermediary.
//
// Take while will take elements while the supplied operation (op) returns true.
// Once the supplied operation returns false iteration will stop. If an error
// is returned from the parent iterator iteration will stop and the operation
// function will not be called on the value that errored.
func (i Iter[T]) TakeWhile(op func(val T) bool) Iter[T] {
	return i.Next(
		func(index int, val T, status IteratorFeedback) (IteratorFeedback, T, error) {
			if status != Break && op(val) {
				return Continue, val, nil
			}
			return Break, val, nil
		},
	)
}

// This function is an intermediary.
//
// Skip will skip the first num elements of it's parent iterator before
// propagating any further elements to it's child iterator. Skip will stop
// iteration if an error is returned from it's parent iterator regardless of if
// it has reached num elements or not.
func (i Iter[T]) Skip(num int) Iter[T] {
	return i.Filter(FilterToIndex[T](num))
}

// This function is an intermediary.
//
// Map will create a mapping between value of one iterator and values of another
// iterator. Iteration will stop if an error is generated.
func Map[T any, U any](
	i Iter[T],
	op func(index int, val T) (U, error),
) Iter[U] {
	return Next(
		i,
		func(index int, val T, status IteratorFeedback) (IteratorFeedback, U, error) {
			if status == Break {
				var tmp U
				return Break, tmp, nil
			}
			tmp, err := op(index, val)
			return Continue, tmp, err
		},
	)
}

// This function is an intermediary.
//
// ValToPntr will map a stream of values to a stream of pointers to the values.
// ValToPntr will stop iteration if an error is returned from it's parent
// iterator regardless of if it has fully consumed it's input stream of values.
// ValToPntr will never be the cause of an error.
func ValToPntr[T any](i Iter[T]) Iter[*T] {
	return Next[T, *T](
		i,
		func(index int, val T, status IteratorFeedback) (IteratorFeedback, *T, error) {
			if status == Break {
				return Break, nil, nil
			}
			return Continue, &val, nil
		},
	)
}

// This function is an intermediary.
//
// PntrToVal will map a stream of pointers to values to a stream of values.
// PntrToVal will stop iteration if an error is returned from it's parent
// iterator regardless of if it has fully consumed it's input stream of values.
// PntrToVal will never be the cause of an error.
func PntrToVal[T any](i Iter[*T]) Iter[T] {
	return Next[*T, T](
		i,
		func(index int, val *T, status IteratorFeedback) (IteratorFeedback, T, error) {
			if status == Break {
				var tmp T
				return Break, tmp, nil
			}
			return Continue, *val, nil
		},
	)
}

// This function is an intermediary.
//
// Map will create a mapping between two iterators of the same type. This is
// equivalent the calling the previous Map function and providing it with the
// same types. Iteration will stop if an error is generated.
func (i Iter[T]) Map(op func(index int, val T) (T, error)) Iter[T] {
	return Map(i, op)
}

// A helper function that can be passed to the [Filter] function that filters
// elements before a specified index value. Internally, this is how [Skip] is
// implemented.
func FilterToIndex[T any](num int) func(index int, val T) bool {
	return func(index int, val T) bool {
		return !(index < num)
	}
}

// This function is an intermediary.
//
// Filter will selectively pass on values from it's parent iterator to its child
// iterator based on the return value of the operatio (op) function. An error
// will stop iteration and be propogated to the child iterator regardless of
// the implementation of the operation function.
func (i Iter[T]) Filter(op func(index int, val T) bool) Iter[T] {
	return i.Next(
		func(index int, val T, status IteratorFeedback) (IteratorFeedback, T, error) {
			if status != Break && op(index, val) {
				return Continue, val, nil
			}
			return Iterate, val, nil
		},
	)
}

// This function is an intermediary.
//
// Setup will call the provided setup function before it calls its parent
// iterator. The setup function will only be called once. An error returned
// from the setup function will stop iteration and the parent iterator will
// never be called.
func (i Iter[T]) Setup(setup func() error) Iter[T] {
	return i.SetupTeardown(setup, func() error { return nil })
}

// This function is an intermediary.
//
// Teardown will call the provided teardown function once it's parent iterator
// has completed iteration. Teardown will only be called it iterator has started.
// If iteration never began due to an early error teardown will not be called.
// The teardown function will only be called once. An error returned from the
// teardown function will stop iteration.
func (i Iter[T]) Teardown(teardown func() error) Iter[T] {
	return i.SetupTeardown(func() error { return nil }, teardown)
}
