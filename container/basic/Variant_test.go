package basic

import (
	"testing"

	"github.com/barbell-math/util/test"
)

type testStruct struct {
    a int
    b float32
}

func TestVariantA(t *testing.T){
    var tmp int=5;
    v:=Variant[int,testStruct]{};
    newV:=v.SetValA(tmp);
    test.BasicTest(true,newV.HasA(),"Variant did not claim to have correct value.",t);
    test.BasicTest(false,newV.HasB(),"Variant did not claim to have correct value.",t);
    test.BasicTest(5,newV.ValA(),"Variant did not return correct value.",t);
    test.BasicTest(5,newV.ValAOr(1),"Variant did not return correct value.",t);
    test.BasicTest(testStruct{},newV.ValBOr(testStruct{}),
        "Variant did not return correct value.",t,
    );
}

func TestVariantB(t *testing.T){
    var tmp testStruct=testStruct{a: 1, b: 2};
    v:=Variant[int,testStruct]{};
    newV:=v.SetValB(tmp);
    test.BasicTest(false,newV.HasA(),"Variant did not claim to have correct value.",t);
    test.BasicTest(true,newV.HasB(),"Variant did not claim to have correct value.",t);
    test.BasicTest(tmp,newV.ValB(),"Variant did not return correct value.",t);
    test.BasicTest(1,newV.ValAOr(1),"Variant did not return correct value.",t);
    test.BasicTest(tmp,newV.ValBOr(testStruct{}),
        "Variant did not return correct value.",t,
    );
}
