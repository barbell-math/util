package tests

//go:generate ../../bin/passThroughWidget -type=PassThroughWidgetBaseTest
//go:generate ../../bin/passThroughWidget -type=PassThroughWidgetPartialOrderTest
//go:generate ../../bin/passThroughWidget -type=PassThroughWidgetArithTest
//go:generate ../../bin/passThroughWidget -type=PassThroughWidgetArithAndPartialOrderTest

type (
	//gen:passThroughWidget widgetType Base
	//gen:passThroughWidget baseTypeWidget widgets.BuiltinInt
	//gen:passThroughWidget widgetPackage github.com/barbell-math/util/src/widgets
	PassThroughWidgetBaseTest int

	//gen:passThroughWidget widgetType PartialOrder
	//gen:passThroughWidget baseTypeWidget widgets.BuiltinInt
	//gen:passThroughWidget widgetPackage github.com/barbell-math/util/src/widgets
	PassThroughWidgetPartialOrderTest int

	//gen:passThroughWidget widgetType Arith
	//gen:passThroughWidget baseTypeWidget widgets.BuiltinInt
	//gen:passThroughWidget widgetPackage github.com/barbell-math/util/src/widgets
	PassThroughWidgetArithTest int

	//gen:passThroughWidget widgetType PartialOrderArith
	//gen:passThroughWidget baseTypeWidget widgets.BuiltinInt
	//gen:passThroughWidget widgetPackage github.com/barbell-math/util/src/widgets
	PassThroughWidgetArithAndPartialOrderTest int
)
