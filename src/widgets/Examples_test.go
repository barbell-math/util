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
