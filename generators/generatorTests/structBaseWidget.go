package tests

import "github.com/barbell-math/util/src/hash"

//go:generate ../../bin/structBaseWidget -type=structBaseWidgetTest

type (
	testCustomWidget  int
	testCustomWidget2 int

	structBaseWidgetTest struct {
		//gen:structBaseWidget identity
		//gen:structBaseWidget baseTypeWidget widgets.BuiltinString
		//gen:structBaseWidget widgetPackage github.com/barbell-math/util/src/widgets
		value1 string
		//gen:structBaseWidget identity
		//gen:structBaseWidget baseTypeWidget testCustomWidget
		//gen:structBaseWidget widgetPackage .
		value2 testCustomWidget
		//gen:structBaseWidget identity
		//gen:structBaseWidget baseTypeWidget testCustomWidget
		//gen:structBaseWidget widgetPackage .
		value3 testCustomWidget
		// Tests a field that does not contribute to the structs identity
		value4 string
	}
)

func (_ testCustomWidget) Eq(l *testCustomWidget, r *testCustomWidget) bool {
	return false
}
func (_ testCustomWidget) Hash(l *testCustomWidget) hash.Hash {
	return 0
}
func (_ testCustomWidget) Zero(l *testCustomWidget) {}

func (_ *testCustomWidget2) Eq(l *testCustomWidget2, r *testCustomWidget2) bool {
	return false
}
func (_ *testCustomWidget2) Hash(l *testCustomWidget2) hash.Hash {
	return 0
}
func (_ *testCustomWidget2) Zero(l *testCustomWidget2) {}
