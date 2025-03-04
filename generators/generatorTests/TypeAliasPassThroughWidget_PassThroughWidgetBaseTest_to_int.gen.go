package generatortests

// Code generated by ../../bin/passThroughWidget - DO NOT EDIT.

import (
	"github.com/barbell-math/util/src/hash"
	"github.com/barbell-math/util/src/widgets"
)

// Returns true if l equals r. Uses the Eq operator provided by the
// widgets.BuiltinInt widget internally.
func (_ *PassThroughWidgetBaseTest) Eq(
	l *PassThroughWidgetBaseTest,
	r *PassThroughWidgetBaseTest,
) bool {
	var tmp widgets.BuiltinInt
	return tmp.Eq((*int)(l), (*int)(r))
}

// Returns a hash to represent other. The hash that is returned will be supplied
// by the widgets.BuiltinInt widget internally.
func (_ *PassThroughWidgetBaseTest) Hash(
	other *PassThroughWidgetBaseTest,
) hash.Hash {
	var tmp widgets.BuiltinInt
	return tmp.Hash((*int)(other))
}

// Zeros the supplied value. The operation that is performed will be determined
// by the widgets.BuiltinInt widget internally.
func (_ *PassThroughWidgetBaseTest) Zero(
	other *PassThroughWidgetBaseTest,
) {
	var tmp widgets.BuiltinInt
	tmp.Zero((*int)(other))
}
