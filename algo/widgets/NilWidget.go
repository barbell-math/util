package widgets

import "github.com/barbell-math/util/algo/hash"

type (
	NilWidget[T any] struct{}
	ZeroStructWidget struct{}
)

func (n NilWidget[T]) Eq(l *T, r *T) bool      { return false }
func (n NilWidget[T]) Lt(l *T, r *T) bool      { return false }
func (n NilWidget[T]) Hash(other *T) hash.Hash { return hash.Hash(0) }
func (n NilWidget[T]) Zero(other *T)           {}

func (z ZeroStructWidget) Eq(l *struct{}, r *struct{}) bool { return true }
func (z ZeroStructWidget) Lt(l *struct{}, r *struct{}) bool { return false }
func (z ZeroStructWidget) Hash(other *struct{}) hash.Hash   { return hash.Hash(0) }
func (z ZeroStructWidget) Zero(other *struct{})             {}
