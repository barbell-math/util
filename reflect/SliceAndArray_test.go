package reflect

import (
	"reflect"
	"testing"

	"github.com/barbell-math/util/test"
)

func TestIsArrayVal(t *testing.T){
    v:=[]int{}
    test.BasicTest(false,IsArrayVal(&v),
        "IsArrayVal returned false positive.",t,
    )
    v2:=reflect.ValueOf(v)
    test.BasicTest(false,IsArrayVal(&v2),
        "IsArrayVal returned false positive.",t,
    )
    v2=reflect.ValueOf(&v)
    test.BasicTest(false,IsArrayVal(&v2),
        "IsArrayVal returned false positive.",t,
    )
    a:=[3]int{}
    test.BasicTest(true,IsArrayVal(&a),
        "IsArrayVal returned false negative.",t,
    )
    a2:=reflect.ValueOf(a)
    test.BasicTest(true,IsArrayVal(&a2),
        "IsArrayVal returned false negative.",t,
    )
    a2=reflect.ValueOf(&a)
    test.BasicTest(true,IsArrayVal(&a2),
        "IsArrayVal returned false negative.",t,
    )
}

func TestIsSliceVal(t *testing.T){
    v:=[3]int{}
    test.BasicTest(false,IsSliceVal(&v),
        "IsSliceVal returned false positive.",t,
    )
    v2:=reflect.ValueOf(v)
    test.BasicTest(false,IsSliceVal(&v2),
        "IsSliceVal returned false positive.",t,
    )
    v2=reflect.ValueOf(&v)
    test.BasicTest(false,IsSliceVal(&v2),
        "IsSliceVal returned false positive.",t,
    )
    a:=[]int{}
    test.BasicTest(true,IsSliceVal(&a),
        "IsSliceVal returned false negative.",t,
    )
    a2:=reflect.ValueOf(a)
    test.BasicTest(true,IsSliceVal(&a2),
        "IsSliceVal returned false negative.",t,
    )
    a2=reflect.ValueOf(&a)
    test.BasicTest(true,IsSliceVal(&a2),
        "IsSliceVal returned false negative.",t,
    )
}
