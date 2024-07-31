package containers

import (
	"github.com/barbell-math/util/container/dynamicContainers"
	"github.com/barbell-math/util/container/staticContainers"
	"github.com/barbell-math/util/iter"
)

// This function is a producer.
//
// Window will take the parent iterator and return a window of it's cached
// values of length equal to the allowed capacity of the supplied queue (q).
// Note that a static queue is expected to be passed instead of a dynamic one.
// If allowPartials is true then windows that are not full will be returned.
// Setting allowPartials to false will enforce all returned windows to have
// length equal to the allowed capacity of the supplied queue. An error will
// stop iteration.
func Window[T any](
	i iter.Iter[T],
	q interface {
		staticContainers.Queue[T]
		staticContainers.Vector[T]
	},
	allowPartials bool,
) iter.Iter[staticContainers.Vector[T]] {
	return iter.Next(
		i,
		func(
			index int, val T, status iter.IteratorFeedback,
		) (iter.IteratorFeedback, staticContainers.Vector[T], error) {
			if status == iter.Break {
				return iter.Break, q, nil
			}
			q.ForcePushBack(val)
			if !allowPartials && q.Length() != q.Capacity() {
				return iter.Iterate, q, nil
			}
			return iter.Continue, q, nil
		})
}

// This function is a consumer.
//
// Unique will consume all values from it's parent iterator and will collect all
// unique values into a set. The errOnDuplicate argument determines whether or
// not an error will be returned if a duplicate value is found. Iteration will
// stop if errOnDuplicate is set to true and a duplicate value is found.
func Unique[T any](
	i iter.Iter[T],
	s dynamicContainers.Set[T],
	errOnDuplicate bool,
) error {
	return i.ForEach(func(index int, val T) (iter.IteratorFeedback, error) {
		if errOnDuplicate && s.ContainsPntr(&val) {
			return iter.Break, getDuplicateValueError[T](val)
		}
		s.AppendUnique(val)
		return iter.Continue, nil
	})
}
