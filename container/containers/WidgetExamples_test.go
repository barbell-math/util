package containers

import "github.com/barbell-math/util/widgets"

func ExampleVector_OfPntrs() {
	// Lets say that you wanted to create a vector of pointers to integers.
	// You might try to create the vector like this:
	//
	//	v, err:=NewVector[
	// 		*int,
	// 		widgets.BuiltinInt,
	// 	](0)
	//
	// But when you do this you will see compile time errors due to mismatching
	// types of the arguments on the [widgets.BuiltinInt] methods. This is
	// caused by the vector storing pointers to ints rather than ints. To fix
	// this you will need to wrap the BuiltinInt widget in a pointer widget,
	// which will strip off the outter pointer before calling the widget type
	// that the pointer widget was passed. The below code demonstrates this and
	// will not have any compiler errors.
	NewVector[
		*int,
		widgets.BasePntr[int, widgets.BuiltinInt],
	](0)

	// Note that there are 4 kinds of widgets, each of which has it's own
	// pointer widget counterpart. You will need to make sure you choose the
	// correct type of widget for your code.

	// This example applies to all containers within this package.
}
