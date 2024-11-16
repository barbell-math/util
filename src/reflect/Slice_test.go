package reflect

import (
	"reflect"
	"testing"

	"github.com/barbell-math/util/src/customerr"
	"github.com/barbell-math/util/src/test"
)

func TestIsSliceVal(t *testing.T) {
	v := [3]int{}
	test.False(IsSliceVal[[3]int](&v), t)
	v2 := reflect.ValueOf(v)
	test.False(IsSliceVal[int](v2), t)
	v2 = reflect.ValueOf(&v)
	test.False(IsSliceVal[int](v2), t)
	a := []int{}
	test.True(IsSliceVal[[]int](&a), t)
	a2 := reflect.ValueOf(a)
	test.True(IsSliceVal[[]int](a2), t)
	a2 = reflect.ValueOf(&a)
	test.True(IsSliceVal[[]int](a2), t)
}

func TestNonSliceElemVals(t *testing.T) {
	v := 0
	vals, err := SliceElemVals[int](&v).Collect()
	test.Eq(0, len(vals), t)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonSliceElemValsReflectVal(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(v)
	vals, err := SliceElemVals[reflect.Value](v2).Collect()
	test.Eq(0, len(vals), t)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonSliceElemValsReflectValPntr(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(&v)
	vals, err := SliceElemVals[reflect.Value](v2).Collect()
	test.Eq(0, len(vals), t)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestSliceElemVals(t *testing.T) {
	v := []int{0, 1, 2}
	vals, err := SliceElemVals[[]int](&v).Collect()
	test.Eq(3, len(vals), t)
	test.Nil(err, t)
	for i, v := range vals {
		test.Eq(i, v.(int), t)
	}
}

func TestSliceElemValsReflectVal(t *testing.T) {
	v := []int{0, 1, 2}
	v2 := reflect.ValueOf(v)
	vals, err := SliceElemVals[reflect.Value](v2).Collect()
	test.Eq(3, len(vals), t)
	test.Nil(err, t)
	for i, v := range vals {
		test.Eq(i, v.(int), t)
	}
}

func TestSliceElemValsReflectValPntr(t *testing.T) {
	v := []int{0, 1, 2}
	v2 := reflect.ValueOf(&v)
	vals, err := SliceElemVals[reflect.Value](v2).Collect()
	test.Eq(3, len(vals), t)
	test.Nil(err, t)
	for i, v := range vals {
		test.Eq(i, v.(int), t)
	}
}

func TestNonSliceElemPntrs(t *testing.T) {
	v := 0
	vals, err := SliceElemPntrs[int](&v).Collect()
	test.Eq(0, len(vals), t)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonSliceElemPntrsReflectVal(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(v)
	vals, err := SliceElemPntrs[reflect.Value](v2).Collect()
	test.Eq(0, len(vals), t)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonSliceElemPntrsReflectValPntr(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(&v)
	vals, err := SliceElemPntrs[reflect.Value](v2).Collect()
	test.Eq(0, len(vals), t)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestSliceElemPntrs(t *testing.T) {
	v := []int{0, 1, 2}
	vals, err := SliceElemPntrs[[]int](&v).Collect()
	test.Eq(3, len(vals), t)
	test.Nil(err, t)
	for i, iterV := range vals {
		test.Eq(&v[i], iterV.(*int), t)
	}
}

func TestSliceElemPntrsReflectVal(t *testing.T) {
	v := []int{0, 1, 2}
	v2 := reflect.ValueOf(v)
	vals, err := SliceElemPntrs[reflect.Value](v2).Collect()
	test.Eq(3, len(vals), t)
	test.Nil(err, t)
	for i, iterV := range vals {
		test.Eq(&v[i], iterV.(*int), t)
	}
}

func TestSliceElemPntrsReflectValPntr(t *testing.T) {
	v := []int{0, 1, 2}
	v2 := reflect.ValueOf(&v)
	vals, err := SliceElemPntrs[reflect.Value](v2).Collect()
	test.Eq(3, len(vals), t)
	test.Nil(err, t)
	for i, iterV := range vals {
		test.Eq(&v[i], iterV.(*int), t)
	}
}

func TestNonSliceElemType(t *testing.T) {
	v := 0
	_, err := SliceElemType[int](&v)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonSliceElemTypeReflectVal(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(v)
	_, err := SliceElemType[reflect.Value](v2)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonSliceElemTypeReflectValPntr(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(&v)
	_, err := SliceElemType[reflect.Value](v2)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestSliceElemType(t *testing.T) {
	v := []int{0, 1, 2}
	_type, err := SliceElemType[[]int](&v)
	test.Eq(reflect.TypeOf(int(0)), _type, t)
	test.Nil(err, t)
}

func TestSliceElemTypeReflectVal(t *testing.T) {
	v := []int{0, 1, 2}
	v2 := reflect.ValueOf(v)
	_type, err := SliceElemType[[]int](v2)
	test.Eq(reflect.TypeOf(int(0)), _type, t)
	test.Nil(err, t)
}

func TestSliceElemTypeReflectValPntr(t *testing.T) {
	v := []int{0, 1, 2}
	v2 := reflect.ValueOf(&v)
	_type, err := SliceElemType[[]int](v2)
	test.Eq(reflect.TypeOf(int(0)), _type, t)
	test.Nil(err, t)
}

func TestNonSliceElemKind(t *testing.T) {
	v := 0
	_, err := SliceElemKind[int](&v)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonSliceElemKindReflectVal(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(v)
	_, err := SliceElemKind[reflect.Value](v2)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonSliceElemKindReflectValPntr(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(&v)
	_, err := SliceElemKind[reflect.Value](v2)
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestSliceElemKind(t *testing.T) {
	v := []int{0, 1, 2}
	_type, err := SliceElemKind[[]int](&v)
	test.Eq(reflect.Int, _type, t)
	test.Nil(err, t)
}

func TestSliceElemKindReflectVal(t *testing.T) {
	v := []int{0, 1, 2}
	v2 := reflect.ValueOf(v)
	_type, err := SliceElemKind[[]int](v2)
	test.Eq(reflect.Int, _type, t)
	test.Nil(err, t)
}

func TestSliceElemKindReflectValPntr(t *testing.T) {
	v := []int{0, 1, 2}
	v2 := reflect.ValueOf(&v)
	_type, err := SliceElemKind[[]int](v2)
	test.Eq(reflect.Int, _type, t)
	test.Nil(err, t)
}

func TestNonSliceElemInfo(t *testing.T) {
	v := 0
	_, err := SliceElemInfo[int](&v, true).Collect()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonSliceElemInfoReflectVal(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(v)
	_, err := SliceElemInfo[reflect.Value](v2, true).Collect()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonSliceElemInfoReflectValPntr(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(&v)
	_, err := SliceElemInfo[reflect.Value](v2, true).Collect()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestSliceElemInfo(t *testing.T) {
	v := []int{0, 1, 2}
	info, err := SliceElemInfo[[]int](&v, true).Collect()
	test.Nil(err, t)
	test.Eq(3, len(info), t)
	for i, iterV := range v {
		_v, ok := info[i].Val()
		test.Eq(iterV, _v.(int), t)
		test.True(ok, t)
		test.Eq(reflect.TypeOf(int(0)), info[i].Type, t)
		test.Eq(reflect.Int, info[i].Kind, t)
		pntr, err := info[i].Pntr()
		test.Eq(&v[i], pntr.(*int), t)
		test.Nil(err, t)
	}
}

func TestSliceElemInfoReflectVal(t *testing.T) {
	v := []int{0, 1, 2}
	v2 := reflect.ValueOf(v)
	info, err := SliceElemInfo[[]int](v2, true).Collect()
	test.Nil(err, t)
	test.Eq(3, len(info), t)
	for i, iterV := range v {
		_v, ok := info[i].Val()
		test.Eq(iterV, _v.(int), t)
		test.True(ok, t)
		test.Eq(reflect.TypeOf(int(0)), info[i].Type, t)
		test.Eq(reflect.Int, info[i].Kind, t)
		pntr, err := info[i].Pntr()
		test.Eq(&v[i], pntr.(*int), t)
		test.Nil(err, t)
	}
}

func TestSliceElemInfoReflectValPntr(t *testing.T) {
	v := []int{0, 1, 2}
	v2 := reflect.ValueOf(&v)
	info, err := SliceElemInfo[[]int](v2, true).Collect()
	test.Nil(err, t)
	test.Eq(3, len(info), t)
	for i, iterV := range v {
		_v, ok := info[i].Val()
		test.Eq(iterV, _v.(int), t)
		test.True(ok, t)
		test.Eq(reflect.TypeOf(int(0)), info[i].Type, t)
		test.Eq(reflect.Int, info[i].Kind, t)
		pntr, err := info[i].Pntr()
		test.Eq(&v[i], pntr.(*int), t)
		test.Nil(err, t)
	}
}

func TestNonSliceRecursiveElemInfo(t *testing.T) {
	v := 0
	_, err := RecursiveSliceElemInfo[int](&v, true).Collect()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonSliceRecursiveElemInfoReflectVal(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(v)
	_, err := RecursiveSliceElemInfo[reflect.Value](v2, true).Collect()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestNonSliceRecursiveElemInfoReflectValPntr(t *testing.T) {
	v := 0
	v2 := reflect.ValueOf(&v)
	_, err := RecursiveSliceElemInfo[reflect.Value](v2, true).Collect()
	test.ContainsError(customerr.IncorrectType, err, t)
}

func TestRecursiveSliceElemInfo(t *testing.T) {
	v := [][]int{[]int{1, 2, 3}, []int{5, 6, 7}, []int{9, 10, 11}}
	vals, err := RecursiveSliceElemInfo[[][]int](&v, true).Collect()
	test.Nil(err, t)
	test.Eq(12, len(vals), t)
	for i, iterV := range vals {
		if i == 0 || i == 4 || i == 8 {
			tmp := []int{i + 1, i + 2, i + 3}
			for j, tmpV := range tmp {
				_v, ok := iterV.Val()
				test.Eq(tmpV, _v.([]int)[j], t)
				test.True(ok, t)
			}
			test.Eq(reflect.TypeOf([]int{}), iterV.Type, t)
			test.Eq(reflect.Slice, iterV.Kind, t)
			p, err := iterV.Pntr()
			test.Nil(err, t)
			test.Eq(&v[i/4], p, t)
		} else {
			_v, ok := iterV.Val()
			test.Eq(i, _v.(int), t)
			test.True(ok, t)
			test.Eq(reflect.TypeOf(int(0)), iterV.Type, t)
			test.Eq(reflect.Int, iterV.Kind, t)
			p, err := iterV.Pntr()
			test.Nil(err, t)
			test.Eq(&v[i/4][i%4-1], p, t)
		}
	}
}

func TestRecursiveSliceElemInfoReflectVal(t *testing.T) {
	v := [][]int{[]int{1, 2, 3}, []int{5, 6, 7}, []int{9, 10, 11}}
	v2 := reflect.ValueOf(v)
	vals, err := RecursiveSliceElemInfo[[][]int](v2, true).Collect()
	test.Nil(err, t)
	test.Eq(12, len(vals), t)
	for i, iterV := range vals {
		if i == 0 || i == 4 || i == 8 {
			tmp := []int{i + 1, i + 2, i + 3}
			for j, tmpV := range tmp {
				_v, ok := iterV.Val()
				test.Eq(tmpV, _v.([]int)[j], t)
				test.True(ok, t)
			}
			test.Eq(reflect.TypeOf([]int{}), iterV.Type, t)
			test.Eq(reflect.Slice, iterV.Kind, t)
			p, err := iterV.Pntr()
			test.Nil(err, t)
			test.Eq(&v[i/4], p, t)
		} else {
			_v, ok := iterV.Val()
			test.Eq(i, _v.(int), t)
			test.True(ok, t)
			test.Eq(reflect.TypeOf(int(0)), iterV.Type, t)
			test.Eq(reflect.Int, iterV.Kind, t)
			p, err := iterV.Pntr()
			test.Nil(err, t)
			test.Eq(&v[i/4][i%4-1], p, t)
		}
	}
}

func TestRecursiveSliceElemInfoReflectValPntr(t *testing.T) {
	v := [][]int{[]int{1, 2, 3}, []int{5, 6, 7}, []int{9, 10, 11}}
	v2 := reflect.ValueOf(&v)
	vals, err := RecursiveSliceElemInfo[[][]int](v2, true).Collect()
	test.Nil(err, t)
	test.Eq(12, len(vals), t)
	for i, iterV := range vals {
		if i == 0 || i == 4 || i == 8 {
			tmp := []int{i + 1, i + 2, i + 3}
			for j, tmpV := range tmp {
				_v, ok := iterV.Val()
				test.Eq(tmpV, _v.([]int)[j], t)
				test.True(ok, t)
			}
			test.Eq(reflect.TypeOf([]int{}), iterV.Type, t)
			test.Eq(reflect.Slice, iterV.Kind, t)
			p, err := iterV.Pntr()
			test.Nil(err, t)
			test.Eq(&v[i/4], p, t)
		} else {
			_v, ok := iterV.Val()
			test.Eq(i, _v.(int), t)
			test.True(ok, t)
			test.Eq(reflect.TypeOf(int(0)), iterV.Type, t)
			test.Eq(reflect.Int, iterV.Kind, t)
			p, err := iterV.Pntr()
			test.Nil(err, t)
			test.Eq(&v[i/4][i%4-1], p, t)
		}
	}
}
