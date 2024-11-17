package widgets

import (
	"fmt"
	"testing"
)

func Example_BaseWidget(t *testing.T) {
	v1, v2:=0, 1
	w:=Base[int, BuiltinInt]{}
	fmt.Println("Equality:", w.Eq(&v1, &v2))
	fmt.Println("Hash:", w.Hash(&v2))
	w.Zero(&v2)
	fmt.Println("Zeroed:", v2)

	//Output:
	//Equality: false
	//Hash: 1
	//Zeroed: 0
}

func Example_PartialOrderWidget(t *testing.T) {
	v1, v2:=0, 1
	w:=PartialOrder[int, BuiltinInt]{}
	fmt.Println("Lt:", w.Eq(&v1, &v2))

	//Output:
	//Lt: true
}

func Example_ArithWidget(t *testing.T) {
	v1, v2:=0, -1
	res:=0
	w:=Arith[int, BuiltinInt]{}
	fmt.Println("ZeroVal:", w.ZeroVal())
	fmt.Println("UnitVal:", w.UnitVal())
	w.Neg(&v2)
	fmt.Println("Neg:", res)
	w.Add(&res, &v1, &v2)
	fmt.Println("Add:", res)
	w.Mul(&res, &v1, &v2)
	fmt.Println("Mul:", res)

	//Output:
	//ZeroVal: 0
	//UnitVal: 1
	//Neg: 1
	//Add: 1
	//Mul: 1
	//Add: 1
}
