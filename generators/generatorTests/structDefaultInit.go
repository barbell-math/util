package tests

import "fmt"

//go:generate ../../bin/structDefaultInit -struct=structDefaultInitTest

type (
	structDefaultInitTest struct {
		// Tests setting the default value
		//gen:structDefaultInit default 3
		field1 int

		// Tests adding an import
		//gen:structDefaultInit default structDefaultInitTest{}
		//gen:structDefaultInit imports fmt
		field2 fmt.Stringer

		// Tests adding a getter and setter
		//gen:structDefaultInit default 3
		//gen:structDefaultInit setter
		//gen:structDefaultInit getter
		field3 float32

		// Tests when duplicate import paths are provided
		//gen:structDefaultInit default structDefaultInitTest{}
		//gen:structDefaultInit setter
		//gen:structDefaultInit getter
		//gen:structDefaultInit imports fmt
		field4 fmt.Stringer

		// Tests a field that is a pointer
		//gen:structDefaultInit default nil
		//gen:structDefaultInit setter
		//gen:structDefaultInit getter
		field5 *int
	}
)

func (_ structDefaultInitTest) String() string { return "" }
