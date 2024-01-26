package containers

import (
	"fmt"
	"sync"

	"github.com/barbell-math/util/container/containerTypes"
)


type (
	Vector[W containerTypes.Widget[T], T any] []containerTypes.Widget[T]

	SyncedVector[W containerTypes.Widget[T], T any] struct {
		*sync.RWMutex
		Vector[W,T]
	}
)

func NewVector[W containerTypes.Widget[T], T any](size int) (Vector[W,T],error) {
	if size<0 {
		return Vector[W, T]{},fmt.Errorf(
			"%w | Size must be >=0. Got: %d",ValOutsideRange,size,
		)
	}
}
