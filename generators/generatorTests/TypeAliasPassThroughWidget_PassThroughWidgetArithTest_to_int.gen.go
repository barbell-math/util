package generatortests

// Code generated by ../../bin/passThroughWidget - DO NOT EDIT.

import (
	"github.com/barbell-math/util/src/hash"
	"github.com/barbell-math/util/src/widgets"
)

// Returns true if l equals r. Uses the Eq operator provided by the
// widgets.BuiltinInt widget internally.
func (_ *PassThroughWidgetArithTest) Eq(
	l *PassThroughWidgetArithTest,
	r *PassThroughWidgetArithTest,
) bool {
	var tmp widgets.BuiltinInt
	return tmp.Eq((*int)(l), (*int)(r))
}

// Returns a hash to represent other. The hash that is returned will be supplied
// by the widgets.BuiltinInt widget internally.
func (_ *PassThroughWidgetArithTest) Hash(
	other *PassThroughWidgetArithTest,
) hash.Hash {
	var tmp widgets.BuiltinInt
	return tmp.Hash((*int)(other))
}

// Zeros the supplied value. The operation that is performed will be determined
// by the widgets.BuiltinInt widget internally.
func (_ *PassThroughWidgetArithTest) Zero(
	other *PassThroughWidgetArithTest,
) {
	var tmp widgets.BuiltinInt
	tmp.Zero((*int)(other))
}

// Returns the zero value for the underlying type. The value that is performed
// will be determined by the widgets.BuiltinInt widget internally.
func (_ *PassThroughWidgetArithTest) ZeroVal() PassThroughWidgetArithTest {
	var tmp widgets.BuiltinInt
	return (PassThroughWidgetArithTest)(tmp.ZeroVal())
}

// Returns the value that represent "1" for the underlying type. The value that
// is performed will be determined by the widgets.BuiltinInt widget internally.
func (_ *PassThroughWidgetArithTest) UnitVal() PassThroughWidgetArithTest {
	var tmp widgets.BuiltinInt
	return (PassThroughWidgetArithTest)(tmp.ZeroVal())
}

// Negates the value that is supplied to it. The value that is returned will be
// determined by the widgets.BuiltinInt widget internally.
func (_ *PassThroughWidgetArithTest) Neg(
	v *PassThroughWidgetArithTest,
) {
	var tmp widgets.BuiltinInt
	tmp.Neg((*int)(v))
}

// Adds l and r and places the results in res. The value that is returned will
// be determined by the widgets.BuiltinInt widget internally.
func (_ *PassThroughWidgetArithTest) Add(
	res *PassThroughWidgetArithTest,
	l *PassThroughWidgetArithTest,
	r *PassThroughWidgetArithTest,
) {
	var tmp widgets.BuiltinInt
	tmp.Add(
		(*int)(res),
		(*int)(l),
		(*int)(r),
	)
}

// Subtracts l and r and places the results in res. The value that is returned
// will be determined by the widgets.BuiltinInt widget internally.
func (_ *PassThroughWidgetArithTest) Sub(
	res *PassThroughWidgetArithTest,
	l *PassThroughWidgetArithTest,
	r *PassThroughWidgetArithTest,
) {
	var tmp widgets.BuiltinInt
	tmp.Sub(
		(*int)(res),
		(*int)(l),
		(*int)(r),
	)
}

// Multiplys l and r and places the results in res. The value that is returned
// will be determined by the widgets.BuiltinInt widget internally.
func (_ *PassThroughWidgetArithTest) Mul(
	res *PassThroughWidgetArithTest,
	l *PassThroughWidgetArithTest,
	r *PassThroughWidgetArithTest,
) {
	var tmp widgets.BuiltinInt
	tmp.Mul(
		(*int)(res),
		(*int)(l),
		(*int)(r),
	)
}

// Divides l and r and places the results in res. The value that is returned
// will be determined by the widgets.BuiltinInt widget internally.
func (_ *PassThroughWidgetArithTest) Div(
	res *PassThroughWidgetArithTest,
	l *PassThroughWidgetArithTest,
	r *PassThroughWidgetArithTest,
) {
	var tmp widgets.BuiltinInt
	tmp.Div(
		(*int)(res),
		(*int)(l),
		(*int)(r),
	)
}
