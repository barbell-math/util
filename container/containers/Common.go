// This package holds the concrete implementations of the containers.
package containers

import (
	"fmt"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/customerr"
)

func addressableSafeGet[K any, V any](
	other interface {
		containerTypes.Addressable
		containerTypes.ReadKeyedOps[K, V]
	},
	k K,
) (*V, error) {
	if other.IsAddressable() {
		return other.GetPntr(k)
	} else {
		tmp, err := other.Get(k)
		otherV := &tmp
		return otherV, err
	}
}

func addressableSafeValIter[T any](
	other interface {
		containerTypes.Addressable
		containerTypes.ReadOps[T]
	},
) iter.Iter[*T] {
	if other.IsAddressable() {
		return other.ValPntrs()
	}
	return iter.ValToPntr[T](other.Vals())
}

func addressableSafeVerticesIter[V any, E any](
	other containerTypes.GraphComparisonsConstraint[V, E],
) iter.Iter[*V] {
	if other.IsAddressable() {
		return other.VerticePntrs()
	}
	return iter.ValToPntr[V](other.Vertices())
}

func addressableSafeOutVerticesAndEdgesIter[V any, E any](
	other containerTypes.GraphComparisonsConstraint[V, E],
	fromV V,
) iter.Iter[basic.Pair[*E, *V]] {
	if other.IsAddressable() {
		return other.OutEdgesAndVerticePntrs(&fromV)
	}
	return iter.Map[basic.Pair[E,V], basic.Pair[*E,*V]](
		other.OutEdgesAndVertices(fromV),
		func(index int, val basic.Pair[E, V]) (basic.Pair[*E, *V], error) {
			return basic.Pair[*E,*V]{A: &val.A, B: &val.B}, nil
		},
	)
}

func getNonAddressablePanicText(thingName string) string {
	return fmt.Sprintf("A %s is not addressable!", thingName)
}

func getSizeError(size int) error {
	return customerr.Wrap(
		customerr.ValOutsideRange,
		"Size must be >=0. Got: %d", size,
	)
}

func getKeyError[K any](k *K) error {
	return customerr.Wrap(
		containerTypes.KeyError,
		"The supplied key (%v) was not found in the container.",
		*k,
	)
}

func getFullError(maxSize int) error {
	return customerr.Wrap(
		containerTypes.Full,
		"The container has reached its capacity (%d elems) and will not grow.",
		maxSize,
	)
}

func getIndexOutOfBoundsError(idx int, upper int, lower int) error {
	return customerr.AppendError(
		customerr.Wrap(
			containerTypes.KeyError,
			"The supplied key (%d) was not found in the container.",
			idx,
		),
		customerr.Wrap(
			customerr.ValOutsideRange,
			"Index must be >=%d and < %d. Supplied idx: %d",
			lower, upper, idx,
		),
	)
}

func getDuplicateValueError[T any](v T) error {
	return customerr.Wrap(
		containerTypes.Duplicate,
		"The supplied value (%v) is already in the container.",
		v,
	)
}

func getStartEndIndexError(start int, end int) error {
	return customerr.Wrap(
		customerr.InvalidValue,
		"The end index (%d) must be > the start index (%d).",
		end, start,
	)
}

func getVertexError[V any](v *V) error {
	return customerr.AppendError(
		customerr.Wrap(
			customerr.InvalidValue,
			"The supplied vertex was not found.",
		),
		getKeyError[V](v),
	)
}

func getEdgeError[E any](e *E) error {
	return customerr.AppendError(
		customerr.Wrap(
			customerr.InvalidValue,
			"The supplied edge was not found.",
		),
		getKeyError[E](e),
	)
}
