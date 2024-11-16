// This package holds the concrete implementations of the containers.
package containers

import (
	"fmt"

	"github.com/barbell-math/util/src/container/basic"
	"github.com/barbell-math/util/src/container/containerTypes"
	"github.com/barbell-math/util/src/customerr"
	"github.com/barbell-math/util/src/hash"
	"github.com/barbell-math/util/src/iter"
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

func addressableSafeEdgesIter[V any, E any](
	other containerTypes.GraphComparisonsConstraint[V, E],
) iter.Iter[*E] {
	if other.IsAddressable() {
		return other.EdgePntrs()
	}
	return iter.ValToPntr[E](other.Edges())
}

func addressableSafeOutVerticesAndEdgesIter[V any, E any](
	other containerTypes.GraphComparisonsConstraint[V, E],
	fromV V,
) iter.Iter[basic.Pair[*E, *V]] {
	if other.IsAddressable() {
		return other.OutEdgesAndVerticePntrs(&fromV)
	}
	return iter.Map[basic.Pair[E, V], basic.Pair[*E, *V]](
		other.OutEdgesAndVertices(fromV),
		func(index int, val basic.Pair[E, V]) (basic.Pair[*E, *V], error) {
			return basic.Pair[*E, *V]{A: &val.A, B: &val.B}, nil
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
	return customerr.WrapValueList(
		containerTypes.KeyError,
		"The supplied key was not found in the container.",
		[]customerr.WrapListVal{
			{ItemName: "Key", Item: *k},
		},
	)
}

func getValueError[V any](v *V) error {
	return customerr.WrapValueList(
		containerTypes.ValueError,
		"The supplied value was not found in the container.",
		[]customerr.WrapListVal{
			{ItemName: "Value", Item: *v},
		},
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
		getKeyError[int](&idx),
		customerr.Wrap(
			customerr.ValOutsideRange,
			"Index must be >=%d and < %d. Supplied idx: %d",
			lower, upper, idx,
		),
	)
}

func getDuplicateValueError[T any](v T) error {
	return customerr.WrapValueList(
		containerTypes.Duplicate,
		"The supplied value is already in the container.",
		[]customerr.WrapListVal{
			{ItemName: "Value", Item: v},
		},
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
	return customerr.WrapValueList(
		containerTypes.ValueError,
		"The supplied vertex was not found in the container.",
		[]customerr.WrapListVal{
			{ItemName: "Vertex", Item: *v},
		},
	)
}

func getEdgeError[E any](e *E) error {
	return customerr.WrapValueList(
		containerTypes.ValueError,
		"The supplied edge was not found in the container.",
		[]customerr.WrapListVal{
			{ItemName: "Edge", Item: *e},
		},
	)
}

func getUpdateViolationHashError[T any](
	orig *T,
	updated *T,
	origHash hash.Hash,
	newHash hash.Hash,
) error {
	return customerr.AppendError(
		containerTypes.UpdateViolation,
		customerr.WrapValueList(
			customerr.InvalidValue,
			"Updating the original value changed it's hash value.",
			[]customerr.WrapListVal{
				{ItemName: "Orig Value", Item: *orig},
				{ItemName: "New Value ", Item: *updated},
				{ItemName: "Orig Hash ", Item: origHash},
				{ItemName: "New Hash  ", Item: newHash},
			},
		),
	)
}

func getUpdateViolationEqError[T any](
	orig *T,
	updated *T,
) error {
	return customerr.AppendError(
		containerTypes.UpdateViolation,
		customerr.WrapValueList(
			customerr.InvalidValue,
			"Updating the original value changed it's identity.",
			[]customerr.WrapListVal{
				{ItemName: "Orig Value", Item: *orig},
				{ItemName: "New Value ", Item: *updated},
			},
		),
	)
}
