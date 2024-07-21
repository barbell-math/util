package basic

import (
	"github.com/barbell-math/util/algo/hash"
	"github.com/barbell-math/util/algo/widgets"
)

type (
	Pair[T any, U any] struct {
		A T
		B U
	}

	WidgetPair[
		T any,
		U any,
		TI widgets.WidgetInterface[T],
		UI widgets.WidgetInterface[U],
	] Pair[T, U]
)

// This is occasionally useful for passing a factory to something.
func NewPair[T any, U any]() Pair[T, U] {
	return Pair[T, U]{}
}

func (_ *WidgetPair[T, U, TI, UI]) Eq(
	l *WidgetPair[T, U, TI, UI],
	r *WidgetPair[T, U, TI, UI],
) bool {
	tw := widgets.Widget[T, TI]{}
	uw := widgets.Widget[U, UI]{}
	return tw.Eq(&l.A, &r.A) && uw.Eq(&l.B, &r.B)
}

func (_ *WidgetPair[T, U, TI, UI]) Lt(
	l *WidgetPair[T, U, TI, UI],
	r *WidgetPair[T, U, TI, UI],
) bool {
	tw := widgets.Widget[T, TI]{}
	uw := widgets.Widget[U, UI]{}
	return tw.Lt(&l.A, &r.A) && uw.Lt(&l.B, &r.B)
}

func (_ *WidgetPair[T, U, TI, UI]) Hash(
	other *WidgetPair[T, U, TI, UI],
) hash.Hash {
	tw := widgets.Widget[T, TI]{}
	uw := widgets.Widget[U, UI]{}
	return tw.Hash(&other.A).Combine(uw.Hash(&other.B))
}

func (_ *WidgetPair[T, U, TI, UI]) Zero(other *WidgetPair[T, U, TI, UI]) {
	tw := widgets.Widget[T, TI]{}
	uw := widgets.Widget[U, UI]{}
	tw.Zero(&other.A)
	uw.Zero(&other.B)
}
