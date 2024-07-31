package basic

import (
	"github.com/barbell-math/util/hash"
	"github.com/barbell-math/util/widgets"
)

type (
	Triple[T any, U any, V any] struct {
		A T
		B U
		C V
	}

	WidgetTriple[
		T any,
		U any,
		V any,
		TI widgets.WidgetInterface[T],
		UI widgets.WidgetInterface[U],
		VI widgets.WidgetInterface[V],
	] Triple[T, U, V]
)

// This is occasionally useful for passing a factory to something.
func NewTriple[T any, U any, V any]() Triple[T, U, V] {
	return Triple[T, U, V]{}
}

func (_ *WidgetTriple[T, U, V, TI, UI, VI]) Eq(
	l *WidgetTriple[T, U, V, TI, UI, VI],
	r *WidgetTriple[T, U, V, TI, UI, VI],
) bool {
	tw := widgets.Widget[T, TI]{}
	uw := widgets.Widget[U, UI]{}
	vw := widgets.Widget[V, VI]{}
	return tw.Eq(&l.A, &r.A) && uw.Eq(&l.B, &r.B) && vw.Eq(&l.C, &r.C)
}

func (_ *WidgetTriple[T, U, V, TI, UI, VI]) Lt(
	l *WidgetTriple[T, U, V, TI, UI, VI],
	r *WidgetTriple[T, U, V, TI, UI, VI],
) bool {
	tw := widgets.Widget[T, TI]{}
	uw := widgets.Widget[U, UI]{}
	vw := widgets.Widget[V, VI]{}
	return tw.Lt(&l.A, &r.A) && uw.Lt(&l.B, &r.B) && vw.Lt(&l.C, &r.C)
}

func (_ *WidgetTriple[T, U, V, TI, UI, VI]) Hash(
	other *WidgetTriple[T, U, V, TI, UI, VI],
) hash.Hash {
	tw := widgets.Widget[T, TI]{}
	uw := widgets.Widget[U, UI]{}
	vw := widgets.Widget[V, VI]{}
	return tw.Hash(&other.A).
		Combine(uw.Hash(&other.B)).
		Combine(vw.Hash(&other.C))
}

func (_ *WidgetTriple[T, U, V, TI, UI, VI]) Zero(
	other *WidgetTriple[T, U, V, TI, UI, VI],
) {
	tw := widgets.Widget[T, TI]{}
	uw := widgets.Widget[U, UI]{}
	vw := widgets.Widget[V, VI]{}
	tw.Zero(&other.A)
	uw.Zero(&other.B)
	vw.Zero(&other.C)
}
