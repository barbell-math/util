package reflect

import (
	"reflect"
	"testing"

	"github.com/barbell-math/util/test"
)

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
