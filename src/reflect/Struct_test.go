package reflect

import (
	"reflect"
	"testing"

	"github.com/barbell-math/util/src/customerr"
	"github.com/barbell-math/util/src/test"
)

type (
	customString string
	structTest   struct {
		One   int    `json:"one"`
		Two   string `json:"two"`
		Three customString
	}

	structTest2 struct {
		Four float64
		Five structTest
	}
)

func TestIsStructVal(t *testing.T) {
	v := 0
	test.False(IsStructVal[int](&v), t)
	v2 := reflect.ValueOf(v)
	test.False(IsStructVal[int](v2), t)
	v2 = reflect.ValueOf(&v)
	test.False(IsStructVal[int](v2), t)
	s := structTest{}
	test.True(IsStructVal[structTest](&s), t)
	s2 := reflect.ValueOf(s)
	test.True(IsStructVal[structTest](s2), t)
	s2 = reflect.ValueOf(&s)
	test.True(IsStructVal[structTest](s2), t)
}

func TestNonStructGetName(t *testing.T) {
	v := 0
	name, err := GetStructName[int](&v)
	test.Eq("", name, t)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonStructGetNameFromReflectVal(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(v)
	name, err := GetStructName[reflect.Value](v2)
	test.Eq("", name, t)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonStructGetNameFromReflectValPntr(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(&v)
	name, err := GetStructName[reflect.Value](v2)
	test.Eq("", name, t)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestGetStructName(t *testing.T) {
	var s structTest
	name, err := GetStructName[structTest](&s)
	test.Nil(err, t)
	test.Eq("structTest", name, t)
}

func TestGetStructNameFromReflectVal(t *testing.T) {
	var s structTest
	s2 := reflect.ValueOf(s)
	name, err := GetStructName[reflect.Value](s2)
	test.Nil(err, t)
	test.Eq("structTest", name, t)
}

func TestGetStructNameFromReflectValPntr(t *testing.T) {
	var s structTest
	s2 := reflect.ValueOf(&s)
	name, err := GetStructName[reflect.Value](s2)
	test.Nil(err, t)
	test.Eq("structTest", name, t)
}

func TestNonStructStructFieldNames(t *testing.T) {
	v := 0
	err := StructFieldNames[int](&v).Consume()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonStructStructFieldNamesFromReflectVal(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(v)
	err := StructFieldNames[reflect.Value](v2).Consume()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonStructStructFieldNamesFromReflectValPntr(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(&v)
	err := StructFieldNames[reflect.Value](v2).Consume()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestStructFieldNames(t *testing.T) {
	var s structTest
	vals, err := StructFieldNames[structTest](&s).Collect()
	test.Eq(3, len(vals), t)
	test.Eq("One", vals[0], t)
	test.Eq("Two", vals[1], t)
	test.Eq("Three", vals[2], t)
	test.Nil(err, t)
}

func TestStructFieldNamesFromReflectVal(t *testing.T) {
	var s structTest
	s2 := reflect.ValueOf(s)
	vals, err := StructFieldNames[reflect.Value](s2).Collect()
	test.Eq(3, len(vals), t)
	test.Eq("One", vals[0], t)
	test.Eq("Two", vals[1], t)
	test.Eq("Three", vals[2], t)
	test.Nil(err, t)
}

func TestStructFieldNamesFromReflectValPntr(t *testing.T) {
	var s structTest
	s2 := reflect.ValueOf(&s)
	vals, err := StructFieldNames[reflect.Value](s2).Collect()
	test.Eq(3, len(vals), t)
	test.Eq("One", vals[0], t)
	test.Eq("Two", vals[1], t)
	test.Eq("Three", vals[2], t)
	test.Nil(err, t)
}

func TestNonStructStructFieldVals(t *testing.T) {
	v := 0
	err := StructFieldVals[int](&v).Consume()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonStructStructFieldValsFromReflectVal(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(v)
	err := StructFieldVals[reflect.Value](v2).Consume()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonStructStructFieldValsFromReflectValPntr(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(&v)
	err := StructFieldVals[reflect.Value](v2).Consume()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestStructFieldVals(t *testing.T) {
	var s structTest = structTest{One: 1, Two: "2", Three: "3"}
	vals, err := StructFieldVals[structTest](&s).Collect()
	test.Eq(3, len(vals), t)
	test.Eq(1, vals[0], t)
	test.Eq("2", vals[1], t)
	test.Eq(customString("3"), vals[2], t)
	test.Nil(err, t)
}

func TestStructFieldValsFromReflectVal(t *testing.T) {
	var s structTest = structTest{One: 1, Two: "2", Three: "3"}
	s2 := reflect.ValueOf(s)
	vals, err := StructFieldVals[reflect.Value](s2).Collect()
	test.Eq(3, len(vals), t)
	test.Eq(1, vals[0], t)
	test.Eq("2", vals[1], t)
	test.Eq(customString("3"), vals[2], t)
	test.Nil(err, t)
}

func TestStructFieldValsFromReflectValPntr(t *testing.T) {
	var s structTest = structTest{One: 1, Two: "2", Three: "3"}
	s2 := reflect.ValueOf(&s)
	vals, err := StructFieldVals[reflect.Value](s2).Collect()
	test.Eq(3, len(vals), t)
	test.Eq(1, vals[0], t)
	test.Eq("2", vals[1], t)
	test.Eq(customString("3"), vals[2], t)
	test.Nil(err, t)
}

func TestNonStructStructFieldPntrs(t *testing.T) {
	v := 0
	err := StructFieldPntrs[int](&v).Consume()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonStructStructFieldPntrsFromReflectVal(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(v)
	err := StructFieldPntrs[reflect.Value](v2).Consume()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonStructStructFieldPntrsFromReflectValPntr(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(&v)
	err := StructFieldPntrs[reflect.Value](v2).Consume()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestStructFieldPntrs(t *testing.T) {
	s := structTest{
		One: 1,
		Two: "two",
	}
	vals, err := StructFieldPntrs[structTest](&s).Collect()
	test.Eq(3, len(vals), t)
	test.Eq(&s.One, vals[0].(*int), t)
	test.Eq(&s.Two, vals[1].(*string), t)
	test.Eq(&s.Three, vals[2].(*customString), t)
	test.Nil(err, t)
}

func TestStructFieldPntrsFromReflectVal(t *testing.T) {
	s := structTest{
		One: 1,
		Two: "two",
	}
	s2 := reflect.ValueOf(s)
	vals, err := StructFieldPntrs[reflect.Value](s2).Collect()
	test.Eq(0, len(vals), t)
	test.ContainsError(InAddressableField, err, t)
}

func TestStructFieldPntrsFromReflectValPntr(t *testing.T) {
	s := structTest{
		One: 1,
		Two: "two",
	}
	s2 := reflect.ValueOf(&s)
	vals, err := StructFieldPntrs[reflect.Value](s2).Collect()
	test.Eq(3, len(vals), t)
	test.Eq(&s.One, vals[0].(*int), t)
	test.Eq(&s.Two, vals[1].(*string), t)
	test.Eq(&s.Three, vals[2].(*customString), t)
	test.Nil(err, t)
}

func TestNonStructStructFieldTypes(t *testing.T) {
	v := 0
	err := StructFieldTypes[int](&v).Consume()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonStructStructFieldTypesFromReflectVal(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(v)
	err := StructFieldTypes[reflect.Value](v2).Consume()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonStructStructFieldTypesFromReflectValPntr(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(&v)
	err := StructFieldTypes[reflect.Value](v2).Consume()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestStructFieldTypes(t *testing.T) {
	var s structTest
	vals, err := StructFieldTypes[structTest](&s).Collect()
	test.Eq(3, len(vals), t)
	test.Eq(reflect.TypeOf(s.One).String(), vals[0].String(), t)
	test.Eq(reflect.TypeOf(s.Two).String(), vals[1].String(), t)
	test.Eq(reflect.TypeOf(s.Three).String(), vals[2].String(), t)
	test.Nil(err, t)
}

func TestStructFieldTypesFromReflectVal(t *testing.T) {
	var s structTest
	s2 := reflect.ValueOf(&s)
	vals, err := StructFieldTypes[reflect.Value](s2).Collect()
	test.Eq(3, len(vals), t)
	test.Eq(reflect.TypeOf(s.One).String(), vals[0].String(), t)
	test.Eq(reflect.TypeOf(s.Two).String(), vals[1].String(), t)
	test.Eq(reflect.TypeOf(s.Three).String(), vals[2].String(), t)
	test.Nil(err, t)
}

func TestStructFieldTypesFromReflectValPntr(t *testing.T) {
	var s structTest
	s2 := reflect.ValueOf(&s)
	vals, err := StructFieldTypes[reflect.Value](s2).Collect()
	test.Eq(3, len(vals), t)
	test.Eq(reflect.TypeOf(s.One).String(), vals[0].String(), t)
	test.Eq(reflect.TypeOf(s.Two).String(), vals[1].String(), t)
	test.Eq(reflect.TypeOf(s.Three).String(), vals[2].String(), t)
	test.Nil(err, t)
}

func TestNonStructStructFieldKinds(t *testing.T) {
	v := 0
	err := StructFieldKinds[int](&v).Consume()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonStructStructFieldKindsFromReflectVal(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(v)
	err := StructFieldKinds[reflect.Value](v2).Consume()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonStructStructFieldKindsFromReflectValPntr(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(&v)
	err := StructFieldKinds[reflect.Value](v2).Consume()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestStructFieldKinds(t *testing.T) {
	var s structTest
	vals, err := StructFieldKinds[structTest](&s).Collect()
	test.Eq(3, len(vals), t)
	test.Eq(reflect.TypeOf(s.One).Kind(), vals[0], t)
	test.Eq(reflect.TypeOf(s.Two).Kind(), vals[1], t)
	test.Eq(reflect.TypeOf(s.Three).Kind(), vals[2], t)
	test.Nil(err, t)
}

func TestStructFieldKindsFromReflectVal(t *testing.T) {
	var s structTest
	s2 := reflect.ValueOf(s)
	vals, err := StructFieldKinds[reflect.Value](s2).Collect()
	test.Eq(3, len(vals), t)
	test.Eq(reflect.TypeOf(s.One).Kind(), vals[0], t)
	test.Eq(reflect.TypeOf(s.Two).Kind(), vals[1], t)
	test.Eq(reflect.TypeOf(s.Three).Kind(), vals[2], t)
	test.Nil(err, t)
}

func TestStructFieldKindsFromReflectValPntr(t *testing.T) {
	var s structTest
	s2 := reflect.ValueOf(&s)
	vals, err := StructFieldKinds[reflect.Value](s2).Collect()
	test.Eq(3, len(vals), t)
	test.Eq(reflect.TypeOf(s.One).Kind(), vals[0], t)
	test.Eq(reflect.TypeOf(s.Two).Kind(), vals[1], t)
	test.Eq(reflect.TypeOf(s.Three).Kind(), vals[2], t)
	test.Nil(err, t)
}

func TestNonStructStructFieldTags(t *testing.T) {
	v := 0
	err := StructFieldTags[int](&v).Consume()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonStructStructFieldTagsFromReflectVal(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(v)
	err := StructFieldTags[reflect.Value](v2).Consume()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonStructStructFieldTagsFromReflectValPntr(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(&v)
	err := StructFieldTags[reflect.Value](v2).Consume()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestStructFieldTags(t *testing.T) {
	var s structTest
	vals, err := StructFieldTags[structTest](&s).Collect()
	test.Eq(3, len(vals), t)
	test.Eq("one", vals[0].Get("json"), t)
	test.Eq("two", vals[1].Get("json"), t)
	test.Eq("", vals[2].Get("json"), t)
	test.Nil(err, t)
}

func TestStructFieldTagsFromReflectVal(t *testing.T) {
	var s structTest
	s2 := reflect.ValueOf(s)
	vals, err := StructFieldTags[reflect.Value](s2).Collect()
	test.Eq(3, len(vals), t)
	test.Eq("one", vals[0].Get("json"), t)
	test.Eq("two", vals[1].Get("json"), t)
	test.Eq("", vals[2].Get("json"), t)
	test.Nil(err, t)
}

func TestStructFieldTagsFromReflectValPntr(t *testing.T) {
	var s structTest
	s2 := reflect.ValueOf(&s)
	vals, err := StructFieldTags[reflect.Value](s2).Collect()
	test.Eq(3, len(vals), t)
	test.Eq("one", vals[0].Get("json"), t)
	test.Eq("two", vals[1].Get("json"), t)
	test.Eq("", vals[2].Get("json"), t)
	test.Nil(err, t)
}

func TestNonStructStructFieldInfo(t *testing.T) {
	v := 0
	err := StructFieldInfo[int](&v, true).Consume()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonStructStructFieldInfoFromReflectVal(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(v)
	err := StructFieldInfo[reflect.Value](v2, true).Consume()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonStructStructFieldInfoFromReflectValPntr(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(&v)
	err := StructFieldInfo[reflect.Value](v2, true).Consume()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestStructFieldInfo(t *testing.T) {
	var s structTest = structTest{One: 1, Two: "2", Three: "3"}
	vals, err := StructFieldInfo[structTest](&s, true).Collect()
	test.Nil(err, t)
	test.Eq(3, len(vals), t)
	test.Eq("One", vals[0].Name, t)
	test.Eq("Two", vals[1].Name, t)
	test.Eq("Three", vals[2].Name, t)
	test.Eq("one", vals[0].Tag.Get("json"), t)
	test.Eq("two", vals[1].Tag.Get("json"), t)
	test.Eq("", vals[2].Tag.Get("json"), t)
	v, ok := vals[0].Val()
	test.Eq(1, v.(int), t)
	test.True(ok, t)
	v, ok = vals[1].Val()
	test.Eq("2", v.(string), t)
	test.True(ok, t)
	v, ok = vals[2].Val()
	test.Eq(customString("3"), v.(customString), t)
	test.True(ok, t)
	test.Eq(reflect.TypeOf(s.One), vals[0].Type, t)
	test.Eq(reflect.TypeOf(s.Two), vals[1].Type, t)
	test.Eq(reflect.TypeOf(s.Three), vals[2].Type, t)
	test.Eq(reflect.Int, vals[0].Kind, t)
	test.Eq(reflect.String, vals[1].Kind, t)
	test.Eq(reflect.String, vals[2].Kind, t)
	p, err := vals[0].Pntr()
	test.Eq(&s.One, p.(*int), t)
	test.Nil(err, t)
	p, err = vals[1].Pntr()
	test.Eq(&s.Two, p.(*string), t)
	test.Nil(err, t)
	p, err = vals[2].Pntr()
	test.Eq(&s.Three, p.(*customString), t)
	test.Nil(err, t)
}

func TestStructFieldInfoFromReflectVal(t *testing.T) {
	var s structTest = structTest{One: 1, Two: "2", Three: "3"}
	s2 := reflect.ValueOf(s)
	vals, err := StructFieldInfo[reflect.Value](s2, true).Collect()
	test.Nil(err, t)
	test.Eq(3, len(vals), t)
	test.Eq("One", vals[0].Name, t)
	test.Eq("Two", vals[1].Name, t)
	test.Eq("Three", vals[2].Name, t)
	test.Eq("one", vals[0].Tag.Get("json"), t)
	test.Eq("two", vals[1].Tag.Get("json"), t)
	test.Eq("", vals[2].Tag.Get("json"), t)
	v, ok := vals[0].Val()
	test.Eq(1, v.(int), t)
	test.True(ok, t)
	v, ok = vals[1].Val()
	test.Eq("2", v.(string), t)
	test.True(ok, t)
	v, ok = vals[2].Val()
	test.Eq(customString("3"), v.(customString), t)
	test.True(ok, t)
	test.Eq(reflect.TypeOf(s.One), vals[0].Type, t)
	test.Eq(reflect.TypeOf(s.Two), vals[1].Type, t)
	test.Eq(reflect.TypeOf(s.Three), vals[2].Type, t)
	test.Eq(reflect.Int, vals[0].Kind, t)
	test.Eq(reflect.String, vals[1].Kind, t)
	test.Eq(reflect.String, vals[2].Kind, t)
	p, err := vals[0].Pntr()
	test.Nil(p, t)
	test.ContainsError(InAddressableField, err, t)
	p, err = vals[1].Pntr()
	test.Nil(p, t)
	test.ContainsError(InAddressableField, err, t)
	p, err = vals[2].Pntr()
	test.Nil(p, t)
	test.ContainsError(InAddressableField, err, t)
}

func TestStructFieldInfoFromReflectValPntr(t *testing.T) {
	var s structTest = structTest{One: 1, Two: "2", Three: "3"}
	s2 := reflect.ValueOf(&s)
	vals, err := StructFieldInfo[reflect.Value](s2, true).Collect()
	test.Nil(err, t)
	test.Eq(3, len(vals), t)
	test.Eq("One", vals[0].Name, t)
	test.Eq("Two", vals[1].Name, t)
	test.Eq("Three", vals[2].Name, t)
	test.Eq("one", vals[0].Tag.Get("json"), t)
	test.Eq("two", vals[1].Tag.Get("json"), t)
	test.Eq("", vals[2].Tag.Get("json"), t)
	v, ok := vals[0].Val()
	test.Eq(1, v.(int), t)
	test.True(ok, t)
	v, ok = vals[1].Val()
	test.Eq("2", v.(string), t)
	test.True(ok, t)
	v, ok = vals[2].Val()
	test.Eq(customString("3"), v.(customString), t)
	test.True(ok, t)
	test.Eq(reflect.TypeOf(s.One), vals[0].Type, t)
	test.Eq(reflect.TypeOf(s.Two), vals[1].Type, t)
	test.Eq(reflect.TypeOf(s.Three), vals[2].Type, t)
	test.Eq(reflect.Int, vals[0].Kind, t)
	test.Eq(reflect.String, vals[1].Kind, t)
	test.Eq(reflect.String, vals[2].Kind, t)
	p, err := vals[0].Pntr()
	test.Eq(&s.One, p.(*int), t)
	test.Nil(err, t)
	p, err = vals[1].Pntr()
	test.Eq(&s.Two, p.(*string), t)
	test.Nil(err, t)
	p, err = vals[2].Pntr()
	test.Eq(&s.Three, p.(*customString), t)
}

func TestNonStructRecursiveStructFieldInfo(t *testing.T) {
	v := 0
	vals, err := RecursiveStructFieldInfo[int](&v, true).Collect()
	test.Eq(0, len(vals), t)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonStructRecursiveStructFieldInfoFromReflectVal(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(v)
	vals, err := RecursiveStructFieldInfo[reflect.Value](v2, true).Collect()
	test.Eq(0, len(vals), t)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonStructRecursiveStructFieldInfoFromReflectValPntr(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(&v)
	vals, err := RecursiveStructFieldInfo[reflect.Value](v2, true).Collect()
	test.Eq(0, len(vals), t)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestRecursiveStructFieldInfo(t *testing.T) {
	var s structTest = structTest{One: 1, Two: "2", Three: "3"}
	vals, err := RecursiveStructFieldInfo[structTest](&s, true).Collect()
	test.Nil(err, t)
	test.Eq(3, len(vals), t)
	test.Eq("One", vals[0].Name, t)
	test.Eq("Two", vals[1].Name, t)
	test.Eq("Three", vals[2].Name, t)
	test.Eq("one", vals[0].Tag.Get("json"), t)
	test.Eq("two", vals[1].Tag.Get("json"), t)
	test.Eq("", vals[2].Tag.Get("json"), t)
	v, ok := vals[0].Val()
	test.Eq(1, v.(int), t)
	test.True(ok, t)
	v, ok = vals[1].Val()
	test.Eq("2", v.(string), t)
	test.True(ok, t)
	v, ok = vals[2].Val()
	test.Eq(customString("3"), v.(customString), t)
	test.True(ok, t)
	test.Eq(reflect.TypeOf(s.One), vals[0].Type, t)
	test.Eq(reflect.TypeOf(s.Two), vals[1].Type, t)
	test.Eq(reflect.TypeOf(s.Three), vals[2].Type, t)
	test.Eq(reflect.Int, vals[0].Kind, t)
	test.Eq(reflect.String, vals[1].Kind, t)
	test.Eq(reflect.String, vals[2].Kind, t)
	p, err := vals[0].Pntr()
	test.Eq(&s.One, p.(*int), t)
	test.Nil(err, t)
	p, err = vals[1].Pntr()
	test.Eq(&s.Two, p.(*string), t)
	test.Nil(err, t)
	p, err = vals[2].Pntr()
	test.Eq(&s.Three, p.(*customString), t)
	test.Nil(err, t)
}

func TestRecursiveStructFieldInfo2(t *testing.T) {
	var s structTest2 = structTest2{
		Four: 4.0,
		Five: structTest{One: 1, Two: "2", Three: "3"},
	}
	vals, err := RecursiveStructFieldInfo[structTest2](&s, true).Collect()
	test.Nil(err, t)
	test.Eq(5, len(vals), t)
	test.Eq("Four", vals[0].Name, t)
	test.Eq("Five", vals[1].Name, t)
	test.Eq("One", vals[2].Name, t)
	test.Eq("Two", vals[3].Name, t)
	test.Eq("Three", vals[4].Name, t)
	test.Eq("", vals[0].Tag.Get("json"), t)
	test.Eq("", vals[1].Tag.Get("json"), t)
	test.Eq("one", vals[2].Tag.Get("json"), t)
	test.Eq("two", vals[3].Tag.Get("json"), t)
	test.Eq("", vals[4].Tag.Get("json"), t)
	v, ok := vals[0].Val()
	test.Eq(4.0, v.(float64), t)
	test.True(ok, t)
	v, ok = vals[1].Val()
	test.Eq(structTest{One: 1, Two: "2", Three: "3"}, v.(structTest), t)
	test.True(ok, t)
	v, ok = vals[2].Val()
	test.Eq(1, v.(int), t)
	v, ok = vals[3].Val()
	test.Eq("2", v.(string), t)
	test.True(ok, t)
	v, ok = vals[4].Val()
	test.Eq(customString("3"), v.(customString), t)
	test.True(ok, t)
	test.Eq(reflect.TypeOf(s.Four), vals[0].Type, t)
	test.Eq(reflect.TypeOf(s.Five), vals[1].Type, t)
	test.Eq(reflect.TypeOf(s.Five.One), vals[2].Type, t)
	test.Eq(reflect.TypeOf(s.Five.Two), vals[3].Type, t)
	test.Eq(reflect.TypeOf(s.Five.Three), vals[4].Type, t)
	test.Eq(reflect.Float64, vals[0].Kind, t)
	test.Eq(reflect.Struct, vals[1].Kind, t)
	test.Eq(reflect.Int, vals[2].Kind, t)
	test.Eq(reflect.String, vals[3].Kind, t)
	test.Eq(reflect.String, vals[4].Kind, t)
	p, err := vals[0].Pntr()
	test.Eq(&s.Four, p.(*float64), t)
	test.Nil(err, t)
	p, err = vals[1].Pntr()
	test.Eq(&s.Five, p.(*structTest), t)
	test.Nil(err, t)
	p, err = vals[2].Pntr()
	test.Eq(&s.Five.One, p.(*int), t)
	test.Nil(err, t)
	p, err = vals[3].Pntr()
	test.Eq(&s.Five.Two, p.(*string), t)
	test.Nil(err, t)
	p, err = vals[4].Pntr()
	test.Eq(&s.Five.Three, p.(*customString), t)
	test.Nil(err, t)
}

func TestRecursiveStructFieldInfo2FromReflectValue(t *testing.T) {
	var s structTest2 = structTest2{
		Four: 4.0,
		Five: structTest{One: 1, Two: "2", Three: "3"},
	}
	v2 := reflect.ValueOf(s)
	vals, err := RecursiveStructFieldInfo[reflect.Value](v2, true).Collect()
	test.ContainsError(InAddressableField, err, t)
	test.Eq(2, len(vals), t)
	test.Eq("Four", vals[0].Name, t)
	test.Eq("Five", vals[1].Name, t)
	test.Eq("", vals[0].Tag.Get("json"), t)
	test.Eq("", vals[1].Tag.Get("json"), t)
	v, ok := vals[0].Val()
	test.Eq(4.0, v.(float64), t)
	test.True(ok, t)
	v, ok = vals[1].Val()
	test.Eq(structTest{One: 1, Two: "2", Three: "3"}, v.(structTest), t)
	test.True(ok, t)
	test.Eq(reflect.TypeOf(s.Four), vals[0].Type, t)
	test.Eq(reflect.TypeOf(s.Five), vals[1].Type, t)
	test.Eq(reflect.Float64, vals[0].Kind, t)
	test.Eq(reflect.Struct, vals[1].Kind, t)
	p, err := vals[0].Pntr()
	test.Nil(p, t)
	test.ContainsError(InAddressableField, err, t)
	p, err = vals[1].Pntr()
	test.Nil(p, t)
	test.ContainsError(InAddressableField, err, t)
}

func TestRecursiveStructFieldInfo2FromReflectValuePntr(t *testing.T) {
	var s structTest2 = structTest2{
		Four: 4.0,
		Five: structTest{One: 1, Two: "2", Three: "3"},
	}
	s2 := reflect.ValueOf(&s)
	vals, err := RecursiveStructFieldInfo[reflect.Value](s2, true).Collect()
	test.Nil(err, t)
	test.Eq(5, len(vals), t)
	test.Eq("Four", vals[0].Name, t)
	test.Eq("Five", vals[1].Name, t)
	test.Eq("One", vals[2].Name, t)
	test.Eq("Two", vals[3].Name, t)
	test.Eq("Three", vals[4].Name, t)
	test.Eq("", vals[0].Tag.Get("json"), t)
	test.Eq("", vals[1].Tag.Get("json"), t)
	test.Eq("one", vals[2].Tag.Get("json"), t)
	test.Eq("two", vals[3].Tag.Get("json"), t)
	test.Eq("", vals[4].Tag.Get("json"), t)
	v, ok := vals[0].Val()
	test.Eq(4.0, v.(float64), t)
	test.True(ok, t)
	v, ok = vals[1].Val()
	test.Eq(structTest{One: 1, Two: "2", Three: "3"}, v.(structTest), t)
	test.True(ok, t)
	v, ok = vals[2].Val()
	test.Eq(1, v.(int), t)
	test.True(ok, t)
	v, ok = vals[3].Val()
	test.Eq("2", v.(string), t)
	test.True(ok, t)
	v, ok = vals[4].Val()
	test.Eq(customString("3"), v.(customString), t)
	test.True(ok, t)
	test.Eq(reflect.TypeOf(s.Four), vals[0].Type, t)
	test.Eq(reflect.TypeOf(s.Five), vals[1].Type, t)
	test.Eq(reflect.TypeOf(s.Five.One), vals[2].Type, t)
	test.Eq(reflect.TypeOf(s.Five.Two), vals[3].Type, t)
	test.Eq(reflect.TypeOf(s.Five.Three), vals[4].Type, t)
	test.Eq(reflect.Float64, vals[0].Kind, t)
	test.Eq(reflect.Struct, vals[1].Kind, t)
	test.Eq(reflect.Int, vals[2].Kind, t)
	test.Eq(reflect.String, vals[3].Kind, t)
	test.Eq(reflect.String, vals[4].Kind, t)
	p, err := vals[0].Pntr()
	test.Eq(&s.Four, p.(*float64), t)
	test.Nil(err, t)
	p, err = vals[1].Pntr()
	test.Eq(&s.Five, p.(*structTest), t)
	test.Nil(err, t)
	p, err = vals[2].Pntr()
	test.Eq(&s.Five.One, p.(*int), t)
	test.Nil(err, t)
	p, err = vals[3].Pntr()
	test.Eq(&s.Five.Two, p.(*string), t)
	test.Nil(err, t)
	p, err = vals[4].Pntr()
	test.Eq(&s.Five.Three, p.(*customString), t)
	test.Nil(err, t)
}
