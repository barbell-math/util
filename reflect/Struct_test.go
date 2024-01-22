package reflect

import (
	"reflect"
	"testing"

	"github.com/barbell-math/util/test"
	customerr "github.com/barbell-math/util/err"
)

type customString string
type testStruct struct {
    One int;
    Two string;
    Three customString
};

type testStruct2 struct {
    Four float64
    Five testStruct
}

func TestIsStructVal(t *testing.T){
    v:=0
    test.BasicTest(false,IsStructVal[int](&v),
        "IsStructVal returned a false positive.",t,
    )
    v2:=reflect.ValueOf(v)
    test.BasicTest(false,IsStructVal[int](v2),
        "IsStructVal returned a false positive.",t,
    )
    v2=reflect.ValueOf(&v)
    test.BasicTest(false,IsStructVal[int](v2),
        "IsStructVal returned a false positive.",t,
    )
    s:=testStruct{}
    test.BasicTest(true,IsStructVal[testStruct](&s),
        "IsStructVal returned a false negative.",t,
    )
    s2:=reflect.ValueOf(s)
    test.BasicTest(true,IsStructVal[testStruct](s2),
        "IsStructVal returned a false negative.",t,
    )
    s2=reflect.ValueOf(&s)
    test.BasicTest(true,IsStructVal[testStruct](s2),
        "IsStructVal returned a false negative.",t,
    )
}

func TestNonStructGetName(t *testing.T){
    v:=0;
    name,err:=GetStructName[int](&v);
    test.BasicTest("",name,"The name of a non struct type was returned.",t);
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructGetNameFromReflectVal(t *testing.T){
    v:=0;
    v2:=reflect.ValueOf(v)
    name,err:=GetStructName[reflect.Value](v2);
    test.BasicTest("",name,"The name of a non struct type was returned.",t);
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructGetNameFromReflectValPntr(t *testing.T){
    v:=0;
    v2:=reflect.ValueOf(&v)
    name,err:=GetStructName[reflect.Value](v2);
    test.BasicTest("",name,"The name of a non struct type was returned.",t);
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestGetStructName(t *testing.T){
    var s testStruct;
    name,err:=GetStructName[testStruct](&s);
    test.BasicTest(nil,err,
        "Getting name of struct returned error when it was not supposed to.",t,
    );
    test.BasicTest("testStruct",name,"Name of the struct was not correct.",t);
}

func TestGetStructNameFromReflectVal(t *testing.T){
    var s testStruct;
    s2:=reflect.ValueOf(s)
    name,err:=GetStructName[reflect.Value](s2);
    test.BasicTest(nil,err,
        "Getting name of struct returned error when it was not supposed to.",t,
    );
    test.BasicTest("testStruct",name,"Name of the struct was not correct.",t);
}

func TestGetStructNameFromReflectValPntr(t *testing.T){
    var s testStruct;
    s2:=reflect.ValueOf(&s)
    name,err:=GetStructName[reflect.Value](s2);
    test.BasicTest(nil,err,
        "Getting name of struct returned error when it was not supposed to.",t,
    );
    test.BasicTest("testStruct",name,"Name of the struct was not correct.",t);
}

func TestNonStructStructFieldNames(t *testing.T){
    v:=0
    err:=StructFieldNames[int](&v).Consume()
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructStructFieldNamesFromReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    err:=StructFieldNames[reflect.Value](v2).Consume()
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructStructFieldNamesFromReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    err:=StructFieldNames[reflect.Value](v2).Consume()
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestStructFieldNames(t *testing.T){
    var s testStruct
    vals,err:=StructFieldNames[testStruct](&s).Collect()
    test.BasicTest(3,len(vals),
        "The correct number of vals was not returned.",t,
    )
    test.BasicTest("One",vals[0],
        "First struct field name was not correct.",t,
    );
    test.BasicTest("Two",vals[1],
        "Second struct field name was not correct.",t,
    );
    test.BasicTest("Three",vals[2],
        "Third struct field name was not correct.",t,
    );
    test.BasicTest(nil,err,
        "Getting struct field vals returned error when it was not supposed to.",t,
    );
}

func TestStructFieldNamesFromReflectVal(t *testing.T){
    var s testStruct
    s2:=reflect.ValueOf(s)
    vals,err:=StructFieldNames[reflect.Value](s2).Collect()
    test.BasicTest(3,len(vals),
        "The correct number of vals was not returned.",t,
    )
    test.BasicTest("One",vals[0],
        "First struct field name was not correct.",t,
    );
    test.BasicTest("Two",vals[1],
        "Second struct field name was not correct.",t,
    );
    test.BasicTest("Three",vals[2],
        "Third struct field name was not correct.",t,
    );
    test.BasicTest(nil,err,
        "Getting struct field vals returned error when it was not supposed to.",t,
    );
}

func TestStructFieldNamesFromReflectValPntr(t *testing.T){
    var s testStruct
    s2:=reflect.ValueOf(&s)
    vals,err:=StructFieldNames[reflect.Value](s2).Collect()
    test.BasicTest(3,len(vals),
        "The correct number of vals was not returned.",t,
    )
    test.BasicTest("One",vals[0],
        "First struct field name was not correct.",t,
    );
    test.BasicTest("Two",vals[1],
        "Second struct field name was not correct.",t,
    );
    test.BasicTest("Three",vals[2],
        "Third struct field name was not correct.",t,
    );
    test.BasicTest(nil,err,
        "Getting struct field vals returned error when it was not supposed to.",t,
    );
}

func TestNonStructStructFieldVals(t *testing.T){
    v:=0
    err:=StructFieldVals[int](&v).Consume()
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructStructFieldValsFromReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    err:=StructFieldVals[reflect.Value](v2).Consume()
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructStructFieldValsFromReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    err:=StructFieldVals[reflect.Value](v2).Consume()
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestStructFieldVals(t *testing.T){
    var s testStruct=testStruct{One: 1, Two: "2", Three: "3"}
    vals,err:=StructFieldVals[testStruct](&s).Collect()
    test.BasicTest(3,len(vals),
        "The correct number of vals was not returned.",t,
    )
    test.BasicTest(1,vals[0],
        "First struct field val was not correct.",t,
    );
    test.BasicTest("2",vals[1],
        "Second struct field val was not correct.",t,
    );
    test.BasicTest(customString("3"),vals[2],
        "Second struct field val was not correct.",t,
    );
    test.BasicTest(nil,err,
        "Getting struct field vals returned error when it was not supposed to.",t,
    );
}

func TestStructFieldValsFromReflectVal(t *testing.T){
    var s testStruct=testStruct{One: 1, Two: "2", Three: "3"}
    s2:=reflect.ValueOf(s)
    vals,err:=StructFieldVals[reflect.Value](s2).Collect()
    test.BasicTest(3,len(vals),
        "The correct number of vals was not returned.",t,
    )
    test.BasicTest(1,vals[0],
        "First struct field val was not correct.",t,
    );
    test.BasicTest("2",vals[1],
        "Second struct field val was not correct.",t,
    );
    test.BasicTest(customString("3"),vals[2],
        "Second struct field val was not correct.",t,
    );
    test.BasicTest(nil,err,
        "Getting struct field vals returned error when it was not supposed to.",t,
    );
}

func TestStructFieldValsFromReflectValPntr(t *testing.T){
    var s testStruct=testStruct{One: 1, Two: "2", Three: "3"}
    s2:=reflect.ValueOf(&s)
    vals,err:=StructFieldVals[reflect.Value](s2).Collect()
    test.BasicTest(3,len(vals),
        "The correct number of vals was not returned.",t,
    )
    test.BasicTest(1,vals[0],
        "First struct field val was not correct.",t,
    );
    test.BasicTest("2",vals[1],
        "Second struct field val was not correct.",t,
    );
    test.BasicTest(customString("3"),vals[2],
        "Second struct field val was not correct.",t,
    );
    test.BasicTest(nil,err,
        "Getting struct field vals returned error when it was not supposed to.",t,
    );
}

func TestNonStructStructFieldPntrs(t *testing.T){
    v:=0
    err:=StructFieldPntrs[int](&v).Consume()
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructStructFieldPntrsFromReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    err:=StructFieldPntrs[reflect.Value](v2).Consume()
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructStructFieldPntrsFromReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    err:=StructFieldPntrs[reflect.Value](v2).Consume()
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestStructFieldPntrs(t *testing.T){
    var s testStruct
    vals,err:=StructFieldPntrs[testStruct](&s).Collect()
    test.BasicTest(3,len(vals),
        "The correct number of vals was not returned.",t,
    )
    test.BasicTest(&s.One,vals[0].(*int),
        "First struct field pntr was not correct.",t,
    );
    test.BasicTest(&s.Two,vals[1].(*string),
        "Second struct field pntr was not correct.",t,
    );
    test.BasicTest(&s.Three,vals[2].(*customString),
        "Third struct field pntr was not correct.",t,
    );
    test.BasicTest(nil,err,
        "Getting struct field pntrs returned error when it was not supposed to.",t,
    );
}

func TestStructFieldPntrsFromReflectVal(t *testing.T){
    var s testStruct
    s2:=reflect.ValueOf(s)
    vals,err:=StructFieldPntrs[reflect.Value](s2).Collect()
    test.BasicTest(0,len(vals),
        "The correct number of vals was not returned.",t,
    )
    if !IsInAddressableField(err) {
        test.FormatError(InAddressableField(""),err,
            "Getting struct field pointers from a non-pointer reflect value did not return the correct error.",t,
        )
    }
}

func TestStructFieldPntrsFromReflectValPntr(t *testing.T){
    var s testStruct
    s2:=reflect.ValueOf(&s)
    vals,err:=StructFieldPntrs[reflect.Value](s2).Collect()
    test.BasicTest(3,len(vals),
        "The correct number of vals was not returned.",t,
    )
    test.BasicTest(&s.One,vals[0].(*int),
        "First struct field pntr was not correct.",t,
    );
    test.BasicTest(&s.Two,vals[1].(*string),
        "Second struct field pntr was not correct.",t,
    );
    test.BasicTest(&s.Three,vals[2].(*customString),
        "Third struct field pntr was not correct.",t,
    );
    test.BasicTest(nil,err,
        "Getting struct field pntrs returned error when it was not supposed to.",t,
    );
}

func TestNonStructStructFieldTypes(t *testing.T){
    v:=0
    err:=StructFieldTypes[int](&v).Consume()
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructStructFieldTypesFromReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    err:=StructFieldTypes[reflect.Value](v2).Consume()
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructStructFieldTypesFromReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    err:=StructFieldTypes[reflect.Value](v2).Consume()
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestStructFieldTypes(t *testing.T){
    var s testStruct
    vals,err:=StructFieldTypes[testStruct](&s).Collect()
    test.BasicTest(3,len(vals),
        "The correct number of vals was not returned.",t,
    )
    test.BasicTest(reflect.TypeOf(s.One).String(),vals[0].String(),
        "First struct field type was not correct.",t,
    );
    test.BasicTest(reflect.TypeOf(s.Two).String(),vals[1].String(),
        "Second struct field type was not correct.",t,
    );
    test.BasicTest(reflect.TypeOf(s.Three).String(),vals[2].String(),
        "Second struct field type was not correct.",t,
    );
    test.BasicTest(nil,err,
        "Getting struct field types returned error when it was not supposed to.",t,
    );
}

func TestStructFieldTypesFromReflectVal(t *testing.T){
    var s testStruct
    s2:=reflect.ValueOf(&s)
    vals,err:=StructFieldTypes[reflect.Value](s2).Collect()
    test.BasicTest(3,len(vals),
        "The correct number of vals was not returned.",t,
    )
    test.BasicTest(reflect.TypeOf(s.One).String(),vals[0].String(),
        "First struct field type was not correct.",t,
    );
    test.BasicTest(reflect.TypeOf(s.Two).String(),vals[1].String(),
        "Second struct field type was not correct.",t,
    );
    test.BasicTest(reflect.TypeOf(s.Three).String(),vals[2].String(),
        "Second struct field type was not correct.",t,
    );
    test.BasicTest(nil,err,
        "Getting struct field types returned error when it was not supposed to.",t,
    );
}

func TestStructFieldTypesFromReflectValPntr(t *testing.T){
    var s testStruct
    s2:=reflect.ValueOf(&s)
    vals,err:=StructFieldTypes[reflect.Value](s2).Collect()
    test.BasicTest(3,len(vals),
        "The correct number of vals was not returned.",t,
    )
    test.BasicTest(reflect.TypeOf(s.One).String(),vals[0].String(),
        "First struct field type was not correct.",t,
    );
    test.BasicTest(reflect.TypeOf(s.Two).String(),vals[1].String(),
        "Second struct field type was not correct.",t,
    );
    test.BasicTest(reflect.TypeOf(s.Three).String(),vals[2].String(),
        "Second struct field type was not correct.",t,
    );
    test.BasicTest(nil,err,
        "Getting struct field types returned error when it was not supposed to.",t,
    );
}

func TestNonStructStructFieldKinds(t *testing.T){
    v:=0
    err:=StructFieldKinds[int](&v).Consume()
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructStructFieldKindsFromReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    err:=StructFieldKinds[reflect.Value](v2).Consume()
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructStructFieldKindsFromReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    err:=StructFieldKinds[reflect.Value](v2).Consume()
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestStructFieldKinds(t *testing.T){
    var s testStruct
    vals,err:=StructFieldKinds[testStruct](&s).Collect()
    test.BasicTest(3,len(vals),
        "The correct number of vals was not returned.",t,
    )
    test.BasicTest(reflect.TypeOf(s.One).Kind(),vals[0],
        "First struct field kind was not correct.",t,
    );
    test.BasicTest(reflect.TypeOf(s.Two).Kind(),vals[1],
        "Second struct field kind was not correct.",t,
    );
    test.BasicTest(reflect.TypeOf(s.Three).Kind(),vals[2],
        "Second struct field kind was not correct.",t,
    );
    test.BasicTest(nil,err,
        "Getting struct field types returned error when it was not supposed to.",t,
    );
}

func TestStructFieldKindsFromReflectVal(t *testing.T){
    var s testStruct
    s2:=reflect.ValueOf(s)
    vals,err:=StructFieldKinds[reflect.Value](s2).Collect()
    test.BasicTest(3,len(vals),
        "The correct number of vals was not returned.",t,
    )
    test.BasicTest(reflect.TypeOf(s.One).Kind(),vals[0],
        "First struct field kind was not correct.",t,
    );
    test.BasicTest(reflect.TypeOf(s.Two).Kind(),vals[1],
        "Second struct field kind was not correct.",t,
    );
    test.BasicTest(reflect.TypeOf(s.Three).Kind(),vals[2],
        "Second struct field kind was not correct.",t,
    );
    test.BasicTest(nil,err,
        "Getting struct field types returned error when it was not supposed to.",t,
    );
}

func TestStructFieldKindsFromReflectValPntr(t *testing.T){
    var s testStruct
    s2:=reflect.ValueOf(&s)
    vals,err:=StructFieldKinds[reflect.Value](s2).Collect()
    test.BasicTest(3,len(vals),
        "The correct number of vals was not returned.",t,
    )
    test.BasicTest(reflect.TypeOf(s.One).Kind(),vals[0],
        "First struct field kind was not correct.",t,
    );
    test.BasicTest(reflect.TypeOf(s.Two).Kind(),vals[1],
        "Second struct field kind was not correct.",t,
    );
    test.BasicTest(reflect.TypeOf(s.Three).Kind(),vals[2],
        "Second struct field kind was not correct.",t,
    );
    test.BasicTest(nil,err,
        "Getting struct field types returned error when it was not supposed to.",t,
    );
}

func TestNonStructStructFieldInfo(t *testing.T){
    v:=0
    err:=StructFieldInfo[int](&v,true).Consume()
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructStructFieldInfoFromReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    err:=StructFieldInfo[reflect.Value](v2,true).Consume()
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructStructFieldInfoFromReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    err:=StructFieldInfo[reflect.Value](v2,true).Consume()
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestStructFieldInfo(t *testing.T){
    var s testStruct=testStruct{One: 1, Two: "2", Three: "3"}
    vals,err:=StructFieldInfo[testStruct](&s,true).Collect()
    test.BasicTest(nil,err,
        "Getting struct field types returned error when it was not supposed to.",t,
    );
    test.BasicTest(3,len(vals),
        "The correct number of vals was not returned.",t,
    )
    test.BasicTest("One",vals[0].Name,
        "First struct name was not correct.",t,
    )
    test.BasicTest("Two",vals[1].Name,
        "Second struct name was not correct.",t,
    )
    test.BasicTest("Three",vals[2].Name,
        "Third struct name was not correct.",t,
    )
    v,ok:=vals[0].Val()
    test.BasicTest(1,v.(int),
        "First struct val was not correct.",t,
    )
    test.BasicTest(true,ok,
        "First struct val was not correct.",t,
    )
    v,ok=vals[1].Val()
    test.BasicTest("2",v.(string),
        "Second struct val was not correct.",t,
    )
    test.BasicTest(true,ok,
        "Second struct val was not correct.",t,
    )
    v,ok=vals[2].Val()
    test.BasicTest(customString("3"),v.(customString),
        "Third struct val was not correct.",t,
    )
    test.BasicTest(true,ok,
        "Third struct val was not correct.",t,
    )
    test.BasicTest(reflect.TypeOf(s.One),vals[0].Type,
        "First struct type was not correct.",t,
    )
    test.BasicTest(reflect.TypeOf(s.Two),vals[1].Type,
        "Second struct type was not correct.",t,
    )
    test.BasicTest(reflect.TypeOf(s.Three),vals[2].Type,
        "Third struct type was not correct.",t,
    )
    test.BasicTest(reflect.Int,vals[0].Kind,
        "First struct kind was not correct.",t,
    )
    test.BasicTest(reflect.String,vals[1].Kind,
        "Second struct kind was not correct.",t,
    )
    test.BasicTest(reflect.String,vals[2].Kind,
        "Third struct kind was not correct.",t,
    )
    p,err:=vals[0].Pntr()
    test.BasicTest(&s.One,p.(*int),
        "First struct kind was not correct.",t,
    )
    test.BasicTest(nil,err,
        "First struct pntr returned an error when it should not have.",t,
    )
    p,err=vals[1].Pntr()
    test.BasicTest(&s.Two,p.(*string),
        "Second struct kind was not correct.",t,
    )
    test.BasicTest(nil,err,
        "Second struct pntr returned an error when it should not have.",t,
    )
    p,err=vals[2].Pntr()
    test.BasicTest(&s.Three,p.(*customString),
        "Third struct kind was not correct.",t,
    )
    test.BasicTest(nil,err,
        "Third struct pntr returned an error when it should not have.",t,
    )
}

func TestStructFieldInfoFromReflectVal(t *testing.T){
    var s testStruct=testStruct{One: 1, Two: "2", Three: "3"}
    s2:=reflect.ValueOf(s)
    vals,err:=StructFieldInfo[reflect.Value](s2,true).Collect()
    test.BasicTest(nil,err,
        "Getting struct field types returned error when it was not supposed to.",t,
    );
    test.BasicTest(3,len(vals),
        "The correct number of vals was not returned.",t,
    )
    test.BasicTest("One",vals[0].Name,
        "First struct name was not correct.",t,
    )
    test.BasicTest("Two",vals[1].Name,
        "Second struct name was not correct.",t,
    )
    test.BasicTest("Three",vals[2].Name,
        "Third struct name was not correct.",t,
    )
    v,ok:=vals[0].Val()
    test.BasicTest(1,v.(int),
        "First struct val was not correct.",t,
    )
    test.BasicTest(true,ok,
        "First struct val was not correct.",t,
    )
    v,ok=vals[1].Val()
    test.BasicTest("2",v.(string),
        "Second struct val was not correct.",t,
    )
    test.BasicTest(true,ok,
        "Second struct val was not correct.",t,
    )
    v,ok=vals[2].Val()
    test.BasicTest(customString("3"),v.(customString),
        "Third struct val was not correct.",t,
    )
    test.BasicTest(true,ok,
        "Third struct val was not correct.",t,
    )
    test.BasicTest(reflect.TypeOf(s.One),vals[0].Type,
        "First struct type was not correct.",t,
    )
    test.BasicTest(reflect.TypeOf(s.Two),vals[1].Type,
        "Second struct type was not correct.",t,
    )
    test.BasicTest(reflect.TypeOf(s.Three),vals[2].Type,
        "Third struct type was not correct.",t,
    )
    test.BasicTest(reflect.Int,vals[0].Kind,
        "First struct kind was not correct.",t,
    )
    test.BasicTest(reflect.String,vals[1].Kind,
        "Second struct kind was not correct.",t,
    )
    test.BasicTest(reflect.String,vals[2].Kind,
        "Third struct kind was not correct.",t,
    )
    p,err:=vals[0].Pntr()
    test.BasicTest(nil,p,
        "First struct kind was not correct.",t,
    )
    if !IsInAddressableField(err) {
        test.FormatError(InAddressableField(""),err,
            "First struct pntr returned the wrong error.",t,
        )
    }
    p,err=vals[1].Pntr()
    test.BasicTest(nil,p,
        "Second struct kind was not correct.",t,
    )
    if !IsInAddressableField(err) {
        test.FormatError(InAddressableField(""),err,
            "Second struct pntr returned the wrong error.",t,
        )
    }
    p,err=vals[2].Pntr()
    test.BasicTest(nil,p,
        "Third struct kind was not correct.",t,
    )
    if !IsInAddressableField(err) {
        test.FormatError(InAddressableField(""),err,
            "Third struct pntr returned the wrong error.",t,
        )
    }
}

func TestStructFieldInfoFromReflectValPntr(t *testing.T){
    var s testStruct=testStruct{One: 1, Two: "2", Three: "3"}
    s2:=reflect.ValueOf(&s)
    vals,err:=StructFieldInfo[reflect.Value](s2,true).Collect()
    test.BasicTest(nil,err,
        "Getting struct field types returned error when it was not supposed to.",t,
    );
    test.BasicTest(3,len(vals),
        "The correct number of vals was not returned.",t,
    )
    test.BasicTest("One",vals[0].Name,
        "First struct name was not correct.",t,
    )
    test.BasicTest("Two",vals[1].Name,
        "Second struct name was not correct.",t,
    )
    test.BasicTest("Three",vals[2].Name,
        "Third struct name was not correct.",t,
    )
    v,ok:=vals[0].Val()
    test.BasicTest(1,v.(int),
        "First struct val was not correct.",t,
    )
    test.BasicTest(true,ok,
        "First struct val was not correct.",t,
    )
    v,ok=vals[1].Val()
    test.BasicTest("2",v.(string),
        "Second struct val was not correct.",t,
    )
    test.BasicTest(true,ok,
        "Second struct val was not correct.",t,
    )
    v,ok=vals[2].Val()
    test.BasicTest(customString("3"),v.(customString),
        "Third struct val was not correct.",t,
    )
    test.BasicTest(true,ok,
        "Third struct val was not correct.",t,
    )
    test.BasicTest(reflect.TypeOf(s.One),vals[0].Type,
        "First struct type was not correct.",t,
    )
    test.BasicTest(reflect.TypeOf(s.Two),vals[1].Type,
        "Second struct type was not correct.",t,
    )
    test.BasicTest(reflect.TypeOf(s.Three),vals[2].Type,
        "Third struct type was not correct.",t,
    )
    test.BasicTest(reflect.Int,vals[0].Kind,
        "First struct kind was not correct.",t,
    )
    test.BasicTest(reflect.String,vals[1].Kind,
        "Second struct kind was not correct.",t,
    )
    test.BasicTest(reflect.String,vals[2].Kind,
        "Third struct kind was not correct.",t,
    )
    p,err:=vals[0].Pntr()
    test.BasicTest(&s.One,p.(*int),
        "First struct kind was not correct.",t,
    )
    test.BasicTest(nil,err,
        "First struct pntr returned an error when it should not have.",t,
    )
    p,err=vals[1].Pntr()
    test.BasicTest(&s.Two,p.(*string),
        "Second struct kind was not correct.",t,
    )
    test.BasicTest(nil,err,
        "Second struct pntr returned an error when it should not have.",t,
    )
    p,err=vals[2].Pntr()
    test.BasicTest(&s.Three,p.(*customString),
        "Third struct kind was not correct.",t,
    )
}

func TestNonStructRecursiveStructFieldInfo(t *testing.T){
    v:=0;
    vals,err:=RecursiveStructFieldInfo[int](&v,true).Collect();
    test.BasicTest(0,len(vals),
        "Recursive struct field info returned values when it should not have.",t,
    )
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructRecursiveStructFieldInfoFromReflectVal(t *testing.T){
    v:=0;
    v2:=reflect.ValueOf(v)
    vals,err:=RecursiveStructFieldInfo[reflect.Value](v2,true).Collect()
    test.BasicTest(0,len(vals),
        "Recursive struct field info returned values when it should not have.",t,
    )
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructRecursiveStructFieldInfoFromReflectValPntr(t *testing.T){
    v:=0;
    v2:=reflect.ValueOf(&v)
    vals,err:=RecursiveStructFieldInfo[reflect.Value](v2,true).Collect()
    test.BasicTest(0,len(vals),
        "Recursive struct field info returned values when it should not have.",t,
    )
    if !customerr.IsIncorrectType(err) {
        test.FormatError(customerr.IncorrectType(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestRecursiveStructFieldInfo(t *testing.T){
    var s testStruct=testStruct{One: 1, Two: "2", Three: "3"};
    vals,err:=RecursiveStructFieldInfo[testStruct](&s,true).Collect()
    test.BasicTest(nil,err,
        "Recursive struct field info returned error when it was not supposed to.",t,
    );
    test.BasicTest(3,len(vals),
        "Recursive struct field info returned the wrong number of values.",t,     
    )
    test.BasicTest("One",vals[0].Name,
        "First struct name was not correct.",t,
    )
    test.BasicTest("Two",vals[1].Name,
        "Second struct name was not correct.",t,
    )
    test.BasicTest("Three",vals[2].Name,
        "Third struct name was not correct.",t,
    )
    v,ok:=vals[0].Val()
    test.BasicTest(1,v.(int),
        "First struct val was not correct.",t,
    )
    test.BasicTest(true,ok,
        "First struct val was not correct.",t,
    )
    v,ok=vals[1].Val()
    test.BasicTest("2",v.(string),
        "Second struct val was not correct.",t,
    )
    test.BasicTest(true,ok,
        "Second struct val was not correct.",t,
    )
    v,ok=vals[2].Val()
    test.BasicTest(customString("3"),v.(customString),
        "Third struct val was not correct.",t,
    )
    test.BasicTest(true,ok,
        "Third struct val was not correct.",t,
    )
    test.BasicTest(reflect.TypeOf(s.One),vals[0].Type,
        "First struct type was not correct.",t,
    )
    test.BasicTest(reflect.TypeOf(s.Two),vals[1].Type,
        "Second struct type was not correct.",t,
    )
    test.BasicTest(reflect.TypeOf(s.Three),vals[2].Type,
        "Third struct type was not correct.",t,
    )
    test.BasicTest(reflect.Int,vals[0].Kind,
        "First struct kind was not correct.",t,
    )
    test.BasicTest(reflect.String,vals[1].Kind,
        "Second struct kind was not correct.",t,
    )
    test.BasicTest(reflect.String,vals[2].Kind,
        "Third struct kind was not correct.",t,
    )
    p,err:=vals[0].Pntr()
    test.BasicTest(&s.One,p.(*int),
        "First struct kind was not correct.",t,
    )
    test.BasicTest(nil,err,
        "First struct pntr returned an error when it should not have.",t,
    )
    p,err=vals[1].Pntr()
    test.BasicTest(&s.Two,p.(*string),
        "Second struct kind was not correct.",t,
    )
    test.BasicTest(nil,err,
        "Second struct pntr returned an error when it should not have.",t,
    )
    p,err=vals[2].Pntr()
    test.BasicTest(&s.Three,p.(*customString),
        "Third struct kind was not correct.",t,
    )
    test.BasicTest(nil,err,
        "Third struct pntr returned an error when it should not have.",t,
    )
}

func TestRecursiveStructFieldInfo2(t *testing.T){
    var s testStruct2=testStruct2{
        Four: 4.0, 
        Five: testStruct{One: 1, Two: "2", Three: "3"},
    }
    vals,err:=RecursiveStructFieldInfo[testStruct2](&s,true).Collect()
    test.BasicTest(nil,err,
        "Recursive struct field info returned error when it was not supposed to.",t,
    );
    test.BasicTest(5,len(vals),
        "Recursive struct field info returned the wrong number of values.",t,     
    )
    test.BasicTest("Four",vals[0].Name,
        "First struct name was not correct.",t,
    )
    test.BasicTest("Five",vals[1].Name,
        "Second struct name was not correct.",t,
    )
    test.BasicTest("One",vals[2].Name,
        "Third struct name was not correct.",t,
    )
    test.BasicTest("Two",vals[3].Name,
        "Third struct name was not correct.",t,
    )
    test.BasicTest("Three",vals[4].Name,
        "Third struct name was not correct.",t,
    )
    v,ok:=vals[0].Val()
    test.BasicTest(4.0,v.(float64),
        "First struct val was not correct.",t,
    )
    test.BasicTest(true,ok,
        "First struct val was not correct.",t,
    )
    v,ok=vals[1].Val()
    test.BasicTest(testStruct{One: 1, Two: "2", Three: "3"}, v.(testStruct),
        "Second struct val was not correct.",t,
    )
    test.BasicTest(true,ok,
        "Second struct val was not correct.",t,
    )
    v,ok=vals[2].Val()
    test.BasicTest(1,v.(int),
        "First sub-struct val was not correct.",t,
    )
    v,ok=vals[3].Val()
    test.BasicTest("2",v.(string),
        "Second sub-struct val was not correct.",t,
    )
    test.BasicTest(true,ok,
        "Second sub-struct val was not correct.",t,
    )
    v,ok=vals[4].Val()
    test.BasicTest(customString("3"),v.(customString),
        "Third sub-struct val was not correct.",t,
    )
    test.BasicTest(true,ok,
        "Third sub-struct val was not correct.",t,
    )
    test.BasicTest(reflect.TypeOf(s.Four),vals[0].Type,
        "First struct type was not correct.",t,
    )
    test.BasicTest(reflect.TypeOf(s.Five),vals[1].Type,
        "Second struct type was not correct.",t,
    )
    test.BasicTest(reflect.TypeOf(s.Five.One),vals[2].Type,
        "First sub-struct type was not correct.",t,
    )
    test.BasicTest(reflect.TypeOf(s.Five.Two),vals[3].Type,
        "Second sub-struct type was not correct.",t,
    )
    test.BasicTest(reflect.TypeOf(s.Five.Three),vals[4].Type,
        "Third sub-struct type was not correct.",t,
    )
    test.BasicTest(reflect.Float64,vals[0].Kind,
        "First struct kind was not correct.",t,
    )
    test.BasicTest(reflect.Struct,vals[1].Kind,
        "Second struct kind was not correct.",t,
    )
    test.BasicTest(reflect.Int,vals[2].Kind,
        "First sub-struct kind was not correct.",t,
    )
    test.BasicTest(reflect.String,vals[3].Kind,
        "Second sub-struct kind was not correct.",t,
    )
    test.BasicTest(reflect.String,vals[4].Kind,
        "Third sub-struct kind was not correct.",t,
    )
    p,err:=vals[0].Pntr()
    test.BasicTest(&s.Four,p.(*float64),
        "First struct pntr was not correct.",t,
    )
    test.BasicTest(nil,err,
        "First struct pntr returned an error when it should not have.",t,
    )
    p,err=vals[1].Pntr()
    test.BasicTest(&s.Five,p.(*testStruct),
        "First struct pntr was not correct.",t,
    )
    test.BasicTest(nil,err,
        "Second struct pntr returned an error when it should not have.",t,
    )
    p,err=vals[2].Pntr()
    test.BasicTest(&s.Five.One,p.(*int),
        "First sub-struct pntr was not correct.",t,
    )
    test.BasicTest(nil,err,
        "First sub-struct pntr returned an error when it should not have.",t,
    )
    p,err=vals[3].Pntr()
    test.BasicTest(&s.Five.Two,p.(*string),
        "Second sub-struct pntr was not correct.",t,
    )
    test.BasicTest(nil,err,
        "Second sub-struct pntr returned an error when it should not have.",t,
    )
    p,err=vals[4].Pntr()
    test.BasicTest(&s.Five.Three,p.(*customString),
        "Third sub-struct pntr was not correct.",t,
    )
    test.BasicTest(nil,err,
        "Third sub-struct pntr returned an error when it should not have.",t,
    )
}

func TestRecursiveStructFieldInfo2FromReflectValue(t *testing.T){
    var s testStruct2=testStruct2{
        Four: 4.0, 
        Five: testStruct{One: 1, Two: "2", Three: "3"},
    }
    v2:=reflect.ValueOf(s)
    vals,err:=RecursiveStructFieldInfo[reflect.Value](v2,true).Collect()
    if !IsInAddressableField(err) {
        test.FormatError(InAddressableField(""),err,
            "Recursive struct field did not return the correct error.",t,
        )
    }
    test.BasicTest(2,len(vals),
        "Recursive struct field info returned the wrong number of values.",t,     
    )
    test.BasicTest("Four",vals[0].Name,
        "First struct name was not correct.",t,
    )
    test.BasicTest("Five",vals[1].Name,
        "Second struct name was not correct.",t,
    )
    v,ok:=vals[0].Val()
    test.BasicTest(4.0,v.(float64),
        "First struct val was not correct.",t,
    )
    test.BasicTest(true,ok,
        "First struct val was not correct.",t,
    )
    v,ok=vals[1].Val()
    test.BasicTest(testStruct{One: 1, Two: "2", Three: "3"}, v.(testStruct),
        "Second struct val was not correct.",t,
    )
    test.BasicTest(true,ok,
        "Second struct val was not correct.",t,
    )
    test.BasicTest(reflect.TypeOf(s.Four),vals[0].Type,
        "First struct type was not correct.",t,
    )
    test.BasicTest(reflect.TypeOf(s.Five),vals[1].Type,
        "Second struct type was not correct.",t,
    )
    test.BasicTest(reflect.Float64,vals[0].Kind,
        "First struct kind was not correct.",t,
    )
    test.BasicTest(reflect.Struct,vals[1].Kind,
        "Second struct kind was not correct.",t,
    )
    p,err:=vals[0].Pntr()
    test.BasicTest(nil,p,
        "First struct pntr was not correct.",t,
    )
    if !IsInAddressableField(err) {
        test.FormatError(InAddressableField(""),err,
            "Recursive struct field did not return the correct error.",t,
        )
    }
    p,err=vals[1].Pntr()
    test.BasicTest(nil,p,
        "First struct pntr was not correct.",t,
    )
    if !IsInAddressableField(err) {
        test.FormatError(InAddressableField(""),err,
            "Recursive struct field did not return the correct error.",t,
        )
    }
}

func TestRecursiveStructFieldInfo2FromReflectValuePntr(t *testing.T){
    var s testStruct2=testStruct2{
        Four: 4.0, 
        Five: testStruct{One: 1, Two: "2", Three: "3"},
    }
    s2:=reflect.ValueOf(&s)
    vals,err:=RecursiveStructFieldInfo[reflect.Value](s2,true).Collect()
    test.BasicTest(nil,err,
        "Recursive struct field info returned error when it was not supposed to.",t,
    );
    test.BasicTest(5,len(vals),
        "Recursive struct field info returned the wrong number of values.",t,     
    )
    test.BasicTest("Four",vals[0].Name,
        "First struct name was not correct.",t,
    )
    test.BasicTest("Five",vals[1].Name,
        "Second struct name was not correct.",t,
    )
    test.BasicTest("One",vals[2].Name,
        "Third struct name was not correct.",t,
    )
    test.BasicTest("Two",vals[3].Name,
        "Third struct name was not correct.",t,
    )
    test.BasicTest("Three",vals[4].Name,
        "Third struct name was not correct.",t,
    )
    v,ok:=vals[0].Val()
    test.BasicTest(4.0,v.(float64),
        "First struct val was not correct.",t,
    )
    test.BasicTest(true,ok,
        "First struct val was not correct.",t,
    )
    v,ok=vals[1].Val()
    test.BasicTest(testStruct{One: 1, Two: "2", Three: "3"}, v.(testStruct),
        "Second struct val was not correct.",t,
    )
    test.BasicTest(true,ok,
        "Second struct val was not correct.",t,
    )
    v,ok=vals[2].Val()
    test.BasicTest(1,v.(int),
        "First sub-struct val was not correct.",t,
    )
    test.BasicTest(true,ok,
        "First sub-struct val was not correct.",t,
    )
    v,ok=vals[3].Val()
    test.BasicTest("2",v.(string),
        "Second sub-struct val was not correct.",t,
    )
    test.BasicTest(true,ok,
        "Second sub-struct val was not correct.",t,
    )
    v,ok=vals[4].Val()
    test.BasicTest(customString("3"),v.(customString),
        "Third sub-struct val was not correct.",t,
    )
    test.BasicTest(true,ok,
        "Third sub-struct val was not correct.",t,
    )
    test.BasicTest(reflect.TypeOf(s.Four),vals[0].Type,
        "First struct type was not correct.",t,
    )
    test.BasicTest(reflect.TypeOf(s.Five),vals[1].Type,
        "Second struct type was not correct.",t,
    )
    test.BasicTest(reflect.TypeOf(s.Five.One),vals[2].Type,
        "First sub-struct type was not correct.",t,
    )
    test.BasicTest(reflect.TypeOf(s.Five.Two),vals[3].Type,
        "Second sub-struct type was not correct.",t,
    )
    test.BasicTest(reflect.TypeOf(s.Five.Three),vals[4].Type,
        "Third sub-struct type was not correct.",t,
    )
    test.BasicTest(reflect.Float64,vals[0].Kind,
        "First struct kind was not correct.",t,
    )
    test.BasicTest(reflect.Struct,vals[1].Kind,
        "Second struct kind was not correct.",t,
    )
    test.BasicTest(reflect.Int,vals[2].Kind,
        "First sub-struct kind was not correct.",t,
    )
    test.BasicTest(reflect.String,vals[3].Kind,
        "Second sub-struct kind was not correct.",t,
    )
    test.BasicTest(reflect.String,vals[4].Kind,
        "Third sub-struct kind was not correct.",t,
    )
    p,err:=vals[0].Pntr()
    test.BasicTest(&s.Four,p.(*float64),
        "First struct pntr was not correct.",t,
    )
    test.BasicTest(nil,err,
        "First struct pntr returned an error when it should not have.",t,
    )
    p,err=vals[1].Pntr()
    test.BasicTest(&s.Five,p.(*testStruct),
        "First struct pntr was not correct.",t,
    )
    test.BasicTest(nil,err,
        "Second struct pntr returned an error when it should not have.",t,
    )
    p,err=vals[2].Pntr()
    test.BasicTest(&s.Five.One,p.(*int),
        "First sub-struct pntr was not correct.",t,
    )
    test.BasicTest(nil,err,
        "First sub-struct pntr returned an error when it should not have.",t,
    )
    p,err=vals[3].Pntr()
    test.BasicTest(&s.Five.Two,p.(*string),
        "Second sub-struct pntr was not correct.",t,
    )
    test.BasicTest(nil,err,
        "Second sub-struct pntr returned an error when it should not have.",t,
    )
    p,err=vals[4].Pntr()
    test.BasicTest(&s.Five.Three,p.(*customString),
        "Third sub-struct pntr was not correct.",t,
    )
    test.BasicTest(nil,err,
        "Third sub-struct pntr returned an error when it should not have.",t,
    )
}
