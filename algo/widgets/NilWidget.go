package widgets

import "github.com/barbell-math/util/algo/hash"

type (
    NilWidget[T any] struct{}
)

func (n NilWidget[T])Eq(l *T, r *T) bool { return false }
func (n NilWidget[T])Lt(l *T, r *T) bool { return false }
func (n NilWidget[T])Hash(other *T) hash.Hash { return hash.Hash(0) }
func (n NilWidget[T])Zero(other *T) {}
