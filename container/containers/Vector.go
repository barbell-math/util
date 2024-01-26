package containers

import (
	"sync"

	"github.com/barbell-math/util/customerr"
)


type (
	Vector[T any, U any , CONSTRAINT WidgetConstraint[T,U]] []T

	SyncedVector[T any, U any, CONSTRAINT WidgetConstraint[T,U]] struct {
		*sync.RWMutex
		Vector[T,U,CONSTRAINT]
	}
)

func NewVector[T any, U any, CONSTRAINT WidgetConstraint[T,U]](
	size int,
) (Vector[T,U,CONSTRAINT],error) {
	if size<0 {
		return Vector[T,U,CONSTRAINT]{}, customerr.Wrap(
			customerr.ValOutsideRange,
			"Size must be >=0. Got: %d",size,
		)
	}
	return make(Vector[T,U,CONSTRAINT],size),nil
}

func SliceToVector[T any, U any, CONSTRAINT WidgetConstraint[T,U]](
	s []T,
) Vector[T,U,CONSTRAINT] {
	return Vector[T, U, CONSTRAINT](s)
}
