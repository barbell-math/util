package reflect

import (
	"reflect"
	"testing"

	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/iter"
	"github.com/barbell-math/util/test"
)

func TestIsMapVal(t *testing.T) {
	v := 0
	test.False(IsMapVal[int](&v), t)
	v2 := reflect.ValueOf(v)
	test.False(IsMapVal[int](v2), t)
	v2 = reflect.ValueOf(&v)
	test.False(IsMapVal[int](v2), t)
	a := map[int]int{}
	test.True(IsMapVal[map[int]int](&a), t)
	a2 := reflect.ValueOf(a)
	test.True(IsMapVal[map[int]int](a2), t)
	a2 = reflect.ValueOf(&a)
	test.True(IsMapVal[map[int]int](a2), t)
}

func TestNonMapElemKeys(t *testing.T) {
	v := 0
	vals, err := MapElemKeys[int](&v).Collect()
	test.Eq(0, len(vals), t)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonMapElemKeysReflectVal(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(v)
	vals, err := MapElemKeys[reflect.Value](v2).Collect()
	test.Eq(0, len(vals), t)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonMapElemKeysReflectValPntr(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(&v)
	vals, err := MapElemKeys[reflect.Value](v2).Collect()
	test.Eq(0, len(vals), t)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestMapElemKeys(t *testing.T) {
	v := map[int]string{0: "zero", 1: "one", 2: "two"}
	vals, err := MapElemKeys[map[int]string](&v).Collect()
	test.Eq(3, len(vals), t)
	test.Nil(err, t)
	for _, iterV := range vals {
		_, ok := v[iterV.(int)]
		test.True(ok, t)
	}
}

func TestMapElemKeysReflectVal(t *testing.T) {
	v := map[int]string{0: "zero", 1: "one", 2: "two"}
	v2 := reflect.ValueOf(v)
	vals, err := MapElemKeys[reflect.Value](v2).Collect()
	test.Eq(3, len(vals), t)
	test.Nil(err, t)
	for _, iterV := range vals {
		_, ok := v[iterV.(int)]
		test.True(ok, t)
	}
}

func TestMapElemKeysReflectValPntr(t *testing.T) {
	v := map[int]string{0: "zero", 1: "one", 2: "two"}
	v2 := reflect.ValueOf(&v)
	vals, err := MapElemKeys[reflect.Value](v2).Collect()
	test.Eq(3, len(vals), t)
	test.Nil(err, t)
	for _, iterV := range vals {
		_, ok := v[iterV.(int)]
		test.True(ok, t)
	}
}

func TestNonMapElemVals(t *testing.T) {
	v := 0
	vals, err := MapElemVals[int](&v).Collect()
	test.Eq(0, len(vals), t)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonMapElemValsReflectVal(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(v)
	vals, err := MapElemVals[reflect.Value](v2).Collect()
	test.Eq(0, len(vals), t)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonMapElemValsReflectValPntr(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(&v)
	vals, err := MapElemVals[reflect.Value](v2).Collect()
	test.Eq(0, len(vals), t)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestMapElemVals(t *testing.T) {
	v := map[int]string{0: "zero", 1: "one", 2: "two"}
	vals, err := MapElemVals[map[int]string](&v).Collect()
	test.Eq(3, len(vals), t)
	test.Nil(err, t)
	for _, iterV := range vals {
		_, _, found := iter.MapVals(v).Find(func(val string) (bool, error) {
			return val == iterV.(string), nil
		})
		test.True(found, t)
	}
}

func TestMapElemValsReflectVal(t *testing.T) {
	v := map[int]string{0: "zero", 1: "one", 2: "two"}
	v2 := reflect.ValueOf(v)
	vals, err := MapElemVals[reflect.Value](v2).Collect()
	test.Eq(3, len(vals), t)
	test.Nil(err, t)
	for _, iterV := range vals {
		_, _, found := iter.MapVals(v).Find(func(val string) (bool, error) {
			return val == iterV.(string), nil
		})
		test.True(found, t)
	}
}

func TestMapElemValsReflectValPntr(t *testing.T) {
	v := map[int]string{0: "zero", 1: "one", 2: "two"}
	v2 := reflect.ValueOf(&v)
	vals, err := MapElemVals[reflect.Value](v2).Collect()
	test.Eq(3, len(vals), t)
	test.Nil(err, t)
	for _, iterV := range vals {
		_, _, found := iter.MapVals(v).Find(func(val string) (bool, error) {
			return val == iterV.(string), nil
		})
		test.True(found, t)
	}
}

func TestNonMapElems(t *testing.T) {
	v := 0
	vals, err := MapElems[int](&v).Collect()
	test.Eq(0, len(vals), t)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonMapElemsReflectVal(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(v)
	vals, err := MapElemVals[reflect.Value](v2).Collect()
	test.Eq(0, len(vals), t)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonMapElemsReflectValPntr(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(&v)
	vals, err := MapElemVals[reflect.Value](v2).Collect()
	test.Eq(0, len(vals), t)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestMapElems(t *testing.T) {
	v := map[int]string{0: "zero", 1: "one", 2: "two"}
	vals, err := MapElems[map[int]string](&v).Collect()
	test.Eq(3, len(vals), t)
	test.Nil(err, t)
	for _, iterV := range vals {
		_, ok := v[iterV.A.(int)]
		test.True(ok, t)
	}
	for _, iterV := range vals {
		_, _, found := iter.MapVals(v).Find(func(val string) (bool, error) {
			return val == iterV.B, nil
		})
		test.True(found, t)
	}
}

func TestMapElemsReflectVal(t *testing.T) {
	v := map[int]string{0: "zero", 1: "one", 2: "two"}
	v2 := reflect.ValueOf(v)
	vals, err := MapElems[reflect.Value](v2).Collect()
	test.Eq(3, len(vals), t)
	test.Nil(err, t)
	for _, iterV := range vals {
		_, ok := v[iterV.A.(int)]
		test.True(ok, t)
	}
	for _, iterV := range vals {
		_, _, found := iter.MapVals(v).Find(func(val string) (bool, error) {
			return val == iterV.B, nil
		})
		test.True(found, t)
	}
}

func TestMapElemsReflectValPntr(t *testing.T) {
	v := map[int]string{0: "zero", 1: "one", 2: "two"}
	v2 := reflect.ValueOf(&v)
	vals, err := MapElems[reflect.Value](v2).Collect()
	test.Eq(3, len(vals), t)
	test.Nil(err, t)
	for _, iterV := range vals {
		_, ok := v[iterV.A.(int)]
		test.True(ok, t)
	}
	for _, iterV := range vals {
		_, _, found := iter.MapVals(v).Find(func(val string) (bool, error) {
			return val == iterV.B, nil
		})
		test.True(found, t)
	}
}

func TestNonMapKeyType(t *testing.T) {
	v := 0
	_, err := MapKeyType[int](&v)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonMapKeyTypeReflectVal(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(v)
	_, err := MapKeyType[int](v2)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonMapKeyTypeReflectValPntr(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(&v)
	_, err := MapKeyType[int](v2)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestMapKeyType(t *testing.T) {
	v := map[int]string{0: "zero", 1: "one", 2: "two"}
	_t, err := MapKeyType[map[int]string](&v)
	test.Nil(err, t)
	test.Eq(reflect.TypeOf(int(0)), _t, t)
}

func TestMapKeyTypeReflectVal(t *testing.T) {
	v := map[int]string{0: "zero", 1: "one", 2: "two"}
	v2 := reflect.ValueOf(v)
	_t, err := MapKeyType[map[int]string](v2)
	test.Nil(err, t)
	test.Eq(reflect.TypeOf(int(0)), _t, t)
}

func TestMapKeyTypeReflectValPntr(t *testing.T) {
	v := map[int]string{0: "zero", 1: "one", 2: "two"}
	v2 := reflect.ValueOf(&v)
	_t, err := MapKeyType[map[int]string](v2)
	test.Nil(err, t)
	test.Eq(reflect.TypeOf(int(0)), _t, t)
}

func TestNonMapKeyKind(t *testing.T) {
	v := 0
	_, err := MapKeyKind[int](&v)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonMapKeyKindReflectVal(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(v)
	_, err := MapKeyKind[int](v2)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonMapKeyKindReflectValPntr(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(&v)
	_, err := MapKeyKind[int](v2)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestMapKeyKind(t *testing.T) {
	v := map[int]string{0: "zero", 1: "one", 2: "two"}
	_t, err := MapKeyKind[map[int]string](&v)
	test.Nil(err, t)
	test.Eq(reflect.Int, _t, t)
}

func TestMapKeyKindReflectVal(t *testing.T) {
	v := map[int]string{0: "zero", 1: "one", 2: "two"}
	v2 := reflect.ValueOf(v)
	_t, err := MapKeyKind[map[int]string](v2)
	test.Nil(err, t)
	test.Eq(reflect.Int, _t, t)
}

func TestMapKeyKindReflectValPntr(t *testing.T) {
	v := map[int]string{0: "zero", 1: "one", 2: "two"}
	v2 := reflect.ValueOf(&v)
	_t, err := MapKeyKind[map[int]string](v2)
	test.Nil(err, t)
	test.Eq(reflect.Int, _t, t)
}

func TestNonMapValType(t *testing.T) {
	v := 0
	_, err := MapValType[int](&v)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonMapValTypeReflectVal(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(v)
	_, err := MapValType[int](v2)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonMapValTypeReflectValPntr(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(&v)
	_, err := MapValType[int](v2)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestMapValType(t *testing.T) {
	v := map[int]string{0: "zero", 1: "one", 2: "two"}
	_t, err := MapValType[map[int]string](&v)
	test.Nil(err, t)
	test.Eq(reflect.TypeOf(""), _t, t)
}

func TestMapValTypeReflectVal(t *testing.T) {
	v := map[int]string{0: "zero", 1: "one", 2: "two"}
	v2 := reflect.ValueOf(v)
	_t, err := MapValType[map[int]string](v2)
	test.Nil(err, t)
	test.Eq(reflect.TypeOf(""), _t, t)
}

func TestMapValTypeReflectValPntr(t *testing.T) {
	v := map[int]string{0: "zero", 1: "one", 2: "two"}
	v2 := reflect.ValueOf(&v)
	_t, err := MapValType[map[int]string](v2)
	test.Nil(err, t)
	test.Eq(reflect.TypeOf(""), _t, t)
}

func TestNonMapValKind(t *testing.T) {
	v := 0
	_, err := MapValKind[int](&v)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonMapValKindReflectVal(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(v)
	_, err := MapValKind[int](v2)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonMapValKindReflectValPntr(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(&v)
	_, err := MapValKind[int](v2)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestMapValKind(t *testing.T) {
	v := map[int]string{0: "zero", 1: "one", 2: "two"}
	_t, err := MapValKind[map[int]string](&v)
	test.Nil(err, t)
	test.Eq(reflect.String, _t, t)
}

func TestMapValKindReflectVal(t *testing.T) {
	v := map[int]string{0: "zero", 1: "one", 2: "two"}
	v2 := reflect.ValueOf(v)
	_t, err := MapValKind[map[int]string](v2)
	test.Nil(err, t)
	test.Eq(reflect.String, _t, t)
}

func TestMapValKindReflectValPntr(t *testing.T) {
	v := map[int]string{0: "zero", 1: "one", 2: "two"}
	v2 := reflect.ValueOf(&v)
	_t, err := MapValKind[map[int]string](v2)
	test.Nil(err, t)
	test.Eq(reflect.String, _t, t)
}

func TestNonMapElemInfo(t *testing.T) {
	v := 0
	_, err := MapElemInfo[int](&v, true).Collect()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonMapElemInfoReflectVal(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(v)
	_, err := MapElemInfo[int](v2, true).Collect()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonMapElemInfoReflectValPntr(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(&v)
	_, err := MapElemInfo[int](v2, true).Collect()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func mapElemInfoHelper(v map[int]string, info []KeyValInfo, err error, t *testing.T) {
	test.Nil(err, t)
	test.Eq(len(v), len(info), t)
	for _, iterV := range info {
		test.Eq(reflect.TypeOf(int(0)), iterV.A.Type, t)
		test.Eq(reflect.TypeOf(""), iterV.B.Type, t)
		test.Eq(reflect.Int, iterV.A.Kind, t)
		test.Eq(reflect.String, iterV.B.Kind, t)
		k, ok := iterV.A.Val()
		test.True(ok, t)
		actV, ok := v[k.(int)]
		test.True(ok, t)
		_v, ok := iterV.B.Val()
		test.True(ok, t)
		test.Eq(actV, _v.(string), t)
	}
}

func TestMapElemInfo(t *testing.T) {
	v := map[int]string{0: "zero", 1: "one", 2: "two"}
	info, err := MapElemInfo[map[int]string](&v, true).Collect()
	mapElemInfoHelper(v, info, err, t)
}

func TestMapElemInfoReflectVal(t *testing.T) {
	v := map[int]string{0: "zero", 1: "one", 2: "two"}
	v2 := reflect.ValueOf(v)
	info, err := MapElemInfo[map[int]string](v2, true).Collect()
	mapElemInfoHelper(v, info, err, t)
}

func TestMapElemInfoReflectValPntr(t *testing.T) {
	v := map[int]string{0: "zero", 1: "one", 2: "two"}
	v2 := reflect.ValueOf(&v)
	info, err := MapElemInfo[map[int]string](v2, true).Collect()
	mapElemInfoHelper(v, info, err, t)
}

func TestNonRecursiveMapElemInfo(t *testing.T) {
	v := 0
	_, err := RecursiveMapElemInfo[int](&v, true).Collect()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonRecursiveMapElemInfoReflectVal(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(v)
	_, err := RecursiveMapElemInfo[int](v2, true).Collect()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonRecursiveMapElemInfoReflectValPntr(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(&v)
	_, err := RecursiveMapElemInfo[int](v2, true).Collect()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestRecursiveMapElemInfo(t *testing.T) {
	v := map[int]any{0: "zero", 1: "one", 2: "two", 3: map[int]string{4: "four"}}
	info, err := RecursiveMapElemInfo[map[int]any](&v, true).Collect()
	test.Nil(err, t)
	test.Eq(4, len(info), t)
	for _, iterV := range info {
		k, ok := iterV.A.Val()
		test.True(ok, t)
		actV, ok := v[k.(int)]
		test.True(ok, t)
		_v, ok := iterV.B.Val()
		test.True(ok, t)
		switch _v.(type) {
		case string:
			test.Eq(actV, _v.(string), t)
			test.Eq(reflect.TypeOf(int(0)), iterV.A.Type, t)
			// test.Eq(reflect.TypeOf(any("")),iterV.B.Type,t)
			test.Eq(reflect.Int, iterV.A.Kind, t)
			test.Eq(reflect.Interface, iterV.B.Kind, t)
		case map[int]string:
			test.Eq(len(v[3].(map[int]string)), len(_v.(map[int]string)), t)
			test.Eq(v[3].(map[int]string)[4], _v.(map[int]string)[4], t)
			test.Eq(reflect.TypeOf(int(0)), iterV.A.Type, t)
			// test.Eq(reflect.TypeOf(any(map[int]string{})),iterV.B.Type,t)
			test.Eq(reflect.Int, iterV.A.Kind, t)
			test.Eq(reflect.Interface, iterV.B.Kind, t)
		}
	}
}
