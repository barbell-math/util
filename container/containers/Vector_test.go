package containers

import (
	"testing"

	"github.com/barbell-math/util/container/builtinWidgets"
)

func TestNewVector(t *testing.T){
    v,_:=NewVector[int,builtinWidgets.BuiltinInt](0)
    v3:=[]int(v)

    v2:=make([]int,0)
    v=SliceToVector[int,builtinWidgets.BuiltinInt](v2)
    v=Vector[int,builtinWidgets.BuiltinInt,*builtinWidgets.BuiltinInt](v2)
}
