package widgets

import "github.com/barbell-math/util/src/hash"

type (
	// A widget value that is used when values do not necessarily need the
	// methods that the widget interface imposes but still need to be used in
	// scenarios that expect a widget type. The behavior of the widget interface
	// is as follows:
	//  - Widgets of this type will never be considered equal
	//  - Widgets of this type will never be considered less than another
	//  - Widgets of this type will always have a hash of 0
	//  - Widgets of this type will perform no action with the zero function
	NilWidget[T any] struct{}
	// A widget value that is used to represent explicitly zero values.
	// The behavior of the widget interface is as follows:
	//  - Widgets of this type will always be considered equal
	//  - Widgets of this type will never be considered less than another
	//  - Widgets of this type will always have a hash of 0
	//  - Widgets of this type will perform no action with the zero function
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
