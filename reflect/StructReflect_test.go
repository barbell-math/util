package reflect

import (
	"reflect"
	"testing"

	"github.com/barbell-math/util/test"
)

type customString string
type testStruct struct {
    One int;
    Two string;
    Three customString
};

func TestIsStructVal(t *testing.T){
    v:=0
    test.BasicTest(false,IsStructVal(&v),
        "IsStructVal returned a false positive.",t,
    )
    v2:=reflect.ValueOf(v)
    test.BasicTest(false,IsStructVal(&v2),
        "IsStructVal returned a false positive.",t,
    )
    v2=reflect.ValueOf(v)
    test.BasicTest(false,IsStructVal(&v2),
        "IsStructVal returned a false positive.",t,
    )
    s:=testStruct{}
    test.BasicTest(true,IsStructVal(&s),
        "IsStructVal returned a false positive.",t,
    )
    s2:=reflect.ValueOf(s)
    test.BasicTest(true,IsStructVal(&s2),
        "IsStructVal returned a false positive.",t,
    )
    s2=reflect.ValueOf(&s)
    test.BasicTest(true,IsStructVal(&s2),
        "IsStructVal returned a false positive.",t,
    )
}

// func TestIsStructPntr(t *testing.T){
//     v:=0
//     v2:=&v
//     test.BasicTest(false,IsStructPntr(&v2),
//         "IsStructVal returned a false positive.",t,
//     )
//     v3:=reflect.ValueOf(v2)
//     test.BasicTest(false,IsStructPntr(&v3),
//         "IsStructVal returned a false positive.",t,
//     )
//     s:=testStruct{}
//     s2:=&s
//     test.BasicTest(true,IsStructPntr(&s2),
//         "IsStructVal returned a false positive.",t,
//     )
// }

func TestNonStructGetName(t *testing.T){
    v:=0;
    name,err:=GetStructName[int](&v);
    test.BasicTest("",name,"The name of a non struct type was returned.",t);
    if !IsNonStructValue(err) {
        test.FormatError(NonStructValue(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructGetNameFromReflectVal(t *testing.T){
    v:=0;
    v2:=reflect.ValueOf(v)
    name,err:=GetStructName[reflect.Value](v2);
    test.BasicTest("",name,"The name of a non struct type was returned.",t);
    if !IsNonStructValue(err) {
        test.FormatError(NonStructValue(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructGetNameFromReflectValPntr(t *testing.T){
    v:=0;
    v2:=reflect.ValueOf(&v)
    name,err:=GetStructName[reflect.Value](v2);
    test.BasicTest("",name,"The name of a non struct type was returned.",t);
    if !IsNonStructValue(err) {
        test.FormatError(NonStructValue(""),err,
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
    if !IsNonStructValue(err) {
        test.FormatError(NonStructValue(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructStructFieldNamesFromReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    err:=StructFieldNames[reflect.Value](v2).Consume()
    if !IsNonStructValue(err) {
        test.FormatError(NonStructValue(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructStructFieldNamesFromReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    err:=StructFieldNames[reflect.Value](v2).Consume()
    if !IsNonStructValue(err) {
        test.FormatError(NonStructValue(""),err,
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
    if !IsNonStructValue(err) {
        test.FormatError(NonStructValue(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructStructFieldValsFromReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    err:=StructFieldVals[reflect.Value](v2).Consume()
    if !IsNonStructValue(err) {
        test.FormatError(NonStructValue(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructStructFieldValsFromReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    err:=StructFieldVals[reflect.Value](v2).Consume()
    if !IsNonStructValue(err) {
        test.FormatError(NonStructValue(""),err,
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
    if !IsNonStructValue(err) {
        test.FormatError(NonStructValue(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructStructFieldPntrsFromReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    err:=StructFieldPntrs[reflect.Value](v2).Consume()
    if !IsNonStructValue(err) {
        test.FormatError(NonStructValue(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructStructFieldPntrsFromReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    err:=StructFieldPntrs[reflect.Value](v2).Consume()
    if !IsNonStructValue(err) {
        test.FormatError(NonStructValue(""),err,
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
    test.BasicTest(&s.One,vals[0].Interface().(*int),
        "First struct field pntr was not correct.",t,
    );
    test.BasicTest(&s.Two,vals[1].Interface().(*string),
        "Second struct field pntr was not correct.",t,
    );
    test.BasicTest(&s.Three,vals[2].Interface().(*customString),
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
    test.BasicTest(&s.One,vals[0].Interface().(*int),
        "First struct field pntr was not correct.",t,
    );
    test.BasicTest(&s.Two,vals[1].Interface().(*string),
        "Second struct field pntr was not correct.",t,
    );
    test.BasicTest(&s.Three,vals[2].Interface().(*customString),
        "Third struct field pntr was not correct.",t,
    );
    test.BasicTest(nil,err,
        "Getting struct field pntrs returned error when it was not supposed to.",t,
    );
}

func TestNonStructStructFieldTypes(t *testing.T){
    v:=0
    err:=StructFieldTypes[int](&v).Consume()
    if !IsNonStructValue(err) {
        test.FormatError(NonStructValue(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructStructFieldTypesFromReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    err:=StructFieldTypes[reflect.Value](v2).Consume()
    if !IsNonStructValue(err) {
        test.FormatError(NonStructValue(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructStructFieldTypesFromReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    err:=StructFieldTypes[reflect.Value](v2).Consume()
    if !IsNonStructValue(err) {
        test.FormatError(NonStructValue(""),err,
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
    if !IsNonStructValue(err) {
        test.FormatError(NonStructValue(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructStructFieldKindsFromReflectVal(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(v)
    err:=StructFieldKinds[reflect.Value](v2).Consume()
    if !IsNonStructValue(err) {
        test.FormatError(NonStructValue(""),err,
            "Non struct val did not raise appropriate error.",t,
        );
    }
}

func TestNonStructStructFieldKindsFromReflectValPntr(t *testing.T){
    v:=0
    v2:=reflect.ValueOf(&v)
    err:=StructFieldKinds[reflect.Value](v2).Consume()
    if !IsNonStructValue(err) {
        test.FormatError(NonStructValue(""),err,
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
