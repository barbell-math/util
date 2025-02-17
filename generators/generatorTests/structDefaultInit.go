package generatortests

import "fmt"

//go:generate ../../bin/structDefaultInit -struct=structDefaultInitTest
//go:generate ../../bin/structDefaultInit -struct=genericStructDefaultInitTest
//go:generate ../../bin/structDefaultInit -struct=pointerStructDefaultInitTest

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

	genericStructDefaultInitTest[T ~int, U any] struct {
		// Tests setting a generic default value
		//gen:structDefaultInit default generics.ZeroVal[T]()
		//gen:structDefaultInit imports github.com/barbell-math/util/src/generics
		field1 T

		// Tests adding getter and setter with a generic type
		//gen:structDefaultInit default generics.ZeroVal[U]()
		//gen:structDefaultInit imports github.com/barbell-math/util/src/generics
		//gen:structDefaultInit getter
		//gen:structDefaultInit setter
		field2 U
	}

	pointerStructDefaultInitTest struct {
		// Tests adding a pointer setter method to a pointer field
		//gen:structDefaultInit default nil
		//gen:structDefaultInit getter
		//gen:structDefaultInit setter
		//gen:structDefaultInit pointerSetter
		field1 *int
	}
)

func (_ structDefaultInitTest) String() string { return "" }
