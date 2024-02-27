package tests

import (
	"testing"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/container/dynamicContainers"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/test"
)

func vectorReadInterface[U any](c dynamicContainers.ReadVector[U])   {}
func vectorWriteInterface[U any](c dynamicContainers.WriteVector[U]) {}
func vectorInterface[U any](c dynamicContainers.Vector[U])           {}

// Tests that the value supplied by the factory implements the
// [containerTypes.Addressable] interface.
func VectorInterfaceAddressableInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.Addressable = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Length] interface.
func VectorInterfaceLengthInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.Length = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Capacity] interface.
func VectorInterfaceCapacityInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.Capacity = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Clear] interface.
func VectorInterfaceClearInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.Clear = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.WriteOps] interface.
func VectorInterfaceWriteOpsInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.WriteOps[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.WriteKeyedOps] interface.
func VectorInterfaceWriteKeyedOpsInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.WriteKeyedOps[int, V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.WriteKeyedSequentialOps] interface.
func VectorInterfaceWriteKeyedSequentialOpsInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.WriteKeyedSequentialOps[int, V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.KeyedWriteOps] interface.
func VectorInterfaceWriteDynKeyedOpsInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.WriteDynKeyedOps[int, V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.ReadOps] interface.
func VectorInterfaceReadOpsInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.ReadOps[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.KeyedReadOps] interface.
func VectorInterfaceReadKeyedOpsInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.ReadKeyedOps[int, V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.DeleteOps] interface.
func VectorInterfaceDeleteOpsInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.DeleteOps[int, V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.KeyedDeleteOps] interface.
func VectorInterfaceDeleteKeyedOpsInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.DeleteKeyedOps[int, V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.KeyedDeleteOps] interface.
func VectorInterfaceDeleteSequentialOpsInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.DeleteSequentialOps[int, V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.KeyedDeleteOps] interface.
func VectorInterfaceDeleteKeyedSequentialOpsInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.DeleteKeyedSequentialOps[int, V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.VectorRead] interface.
func ReadVectorInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	vectorReadInterface[V](factory())
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.WriteVector] interface.
func WriteVectorInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	vectorWriteInterface[V](factory())
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.Vector] interface.
func VectorInterfaceInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	vectorInterface[V](factory())
}

// Tests that the value supplied by the factory does not implement the
// [staticContainers.Vector] interface.
func VectorInterfaceStaticCapacityInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	test.Panics(
		func() {
			var c any
			c = factory()
			c2 := c.(containerTypes.StaticCapacity)
			_ = c2
		},
		t,
	)
}

// Tests the Get method functionality of a dynamic vector.
func VectorInterfaceGet(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	container := factory()
	_, err := container.Get(0)
	test.ContainsError(customerr.ValOutsideRange, err,t)
	for i := 0; i < 5; i++ {
		container.Append(i)
	}
	for i := 0; i < 5; i++ {
		_v, err := container.Get(i)
		test.Eq(i, _v,t)
		test.Eq(nil, err,t)
	}
	_, err = container.Get(-1)
	test.ContainsError(customerr.ValOutsideRange, err,t)
	_, err = container.Get(6)
	test.ContainsError(customerr.ValOutsideRange, err,t)
}

// Tests the GetPntr method functionality of a dynamic vector.
func VectorInterfaceGetPntr(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	container := factory()
	if container.IsAddressable() {
		_, err := container.GetPntr(0)
		test.ContainsError(customerr.ValOutsideRange, err,t)
		for i := 0; i < 5; i++ {
			container.Append(i)
		}
		for i := 0; i < 5; i++ {
			_v, err := container.GetPntr(i)
			test.Eq(i, *_v,t)
			test.Eq(nil, err,t)
		}
		_, err = container.GetPntr(-1)
		test.ContainsError(customerr.ValOutsideRange, err,t)
		_, err = container.GetPntr(6)
		test.ContainsError(customerr.ValOutsideRange, err,t)
	} else {
		test.Panics(
			func() {
				container:=factory()
				container.GetPntr(1)
			},
			t,
		)
	}
}

func vectorContainsHelper(
	v dynamicContainers.Vector[int],
	l int,
	t *testing.T,
) {
	for i := 0; i < l; i++ {
		v.Append(i)
	}
	for i := 0; i < l; i++ {
		test.True(v.Contains(i),t)
	}
	test.False(v.Contains(-1),t)
	test.False(v.Contains(l),t)
}

// Tests the Contains method functionality of a dynamic vector.
func VectorInterfaceContains(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	vectorContainsHelper(factory(), 0, t)
	vectorContainsHelper(factory(), 1, t)
	vectorContainsHelper(factory(), 2, t)
	vectorContainsHelper(factory(), 5, t)
}

func vectorContainsPntrHelper(
	v dynamicContainers.Vector[int],
	l int,
	t *testing.T,
) {
	for i := 0; i < l; i++ {
		v.Append(i)
	}
	for i := 0; i < l; i++ {
		test.True(v.ContainsPntr(&i),t)
	}
	tmp:=-1
	test.False(v.ContainsPntr(&tmp),t)
	test.False(v.ContainsPntr(&l),t)
}

// Tests the ContainsPntr method functionality of a dynamic vector.
func VectorInterfaceContainsPntr(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	vectorContainsHelper(factory(), 0, t)
	vectorContainsHelper(factory(), 1, t)
	vectorContainsHelper(factory(), 2, t)
	vectorContainsHelper(factory(), 5, t)
}

func vectorKeyOfHelper(
	v dynamicContainers.Vector[int],
	l int,
	t *testing.T,
) {
	for i := 0; i < l; i++ {
		v.Append(i)
	}
	for i := 0; i < l; i++ {
		k, found := v.KeyOf(i)
		test.Eq(i, k,t)
		test.True(found,t)
	}
	_, found := v.KeyOf(-1)
	test.False(found,t)
	_, found = v.KeyOf(-1)
	test.False(v.Contains(l),t)
}

// Tests the KeyOf method functionality of a dynamic vector.
func VectorInterfaceKeyOf(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	vectorKeyOfHelper(factory(), 0, t)
	vectorKeyOfHelper(factory(), 1, t)
	vectorKeyOfHelper(factory(), 2, t)
	vectorKeyOfHelper(factory(), 5, t)
}

// Tests the Set method functionality of a dynamic vector.
func VectorInterfaceSet(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	container := factory()
	err := container.Set(basic.Pair[int, int]{0,6})
	test.ContainsError(customerr.ValOutsideRange, err,t)
	for i := 0; i < 5; i++ {
		container.Append(i)
	}
	for i := 0; i < 5; i++ {
		err := container.Set(basic.Pair[int,int]{i, i+1})
		test.Nil(err,t)
	}
	for i := 0; i < 5; i++ {
		iterV, _ := container.Get(i)
		test.Eq(i+1, iterV,t)
	}
	err = container.Set(basic.Pair[int,int]{-1, 6})
	test.ContainsError(customerr.ValOutsideRange, err,t)
	err = container.Set(basic.Pair[int,int]{6, 6})
	test.ContainsError(customerr.ValOutsideRange, err,t)
}

func vectorSetSequentialHelper(
	factory func() dynamicContainers.Vector[int],
	idx int,
	vals []int,
	l int,
	t *testing.T,
){
	container:=factory()
	for i:=0; i<l; i++ {
		container.Append(i)
	}
	err:=container.SetSequential(idx,vals...)
	test.Nil(err,t)
	for i:=0; i<container.Length(); i++ {
		iterV,_:=container.Get(i)
		if i>=idx && i<idx+len(vals) {
			test.Eq(vals[i-idx],iterV,t)
		} else {
			test.Eq(i,iterV,t)
		}
	}
}
// Tests the SetSequential method functionality of a dynamic vector.
func VectorInterfaceSetSequential(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
){
	for i:=0; i<20; i++ {
		for j:=0; j<i; j++ {
			vals:=[]int{}
			for k:=0; k<i-j; k++ {
				vals = append(vals, k)
				vectorSetSequentialHelper(factory,j,vals,i,t)
			}
		}
	}
}

// Tests the Append method functionality of a dynamic vector.
func VectorInterfaceAppend(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 5; i++ {
		container.Append(i)
	}
	for i := 0; i < 5; i++ {
		iterV, _ := container.Get(i)
		test.Eq(i, iterV,t)
	}
	container.Append(5, 6, 7)
	for i := 0; i < 8; i++ {
		iterV, _ := container.Get(i)
		test.Eq(i, iterV,t)
	}
}

func vectorInsertHelper(
	factory func() dynamicContainers.Vector[int],
	vals []basic.Pair[int,int],
	l int,
	t *testing.T,
){
	indexContained:=func(idx int) (basic.Pair[int,int],bool) {
		for _,v:=range(vals) {
			if idx==v.A {
				return v,true
			}
		}
		return basic.Pair[int,int]{},false
	}
	container:=factory()
	for i:=0; i<l; i++ {
		container.Append(i)
	}
	err:=container.Insert(vals...)
	test.Nil(err,t)
	test.Eq(l+len(vals),container.Length(),t)
	offset:=0
	for i:=0; i<container.Length(); i++ {
		iterV,_:=container.Get(i)
		if v,ok:=indexContained(i); ok {
			test.Eq(v.B,iterV,t)
			offset++
		} else {
			test.Eq(i-offset,iterV,t)
		}
	}
}
// Tests the Insert method functionality of a dynamic vector.
func VectorInterfaceInsert(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
){
	container:=factory()
	err:=container.Insert(basic.Pair[int, int]{1,0})
	test.ContainsError(containerTypes.KeyError,err,t)
	for i:=0; i<20; i++ {
		vals:=[]basic.Pair[int,int]{}
		vectorInsertHelper(factory,vals,i,t)
		for j:=0; j<20; j++ {
			vals = append(vals, basic.Pair[int, int]{j,j})
			vectorInsertHelper(factory,vals,i,t)
		}
	}
	for i:=0; i<20; i++ {
		vals:=[]basic.Pair[int,int]{}
		vectorInsertHelper(factory,vals,i,t)
		for j:=0; j<i; j+=2 {
			vals = append(vals, basic.Pair[int, int]{j,j})
			vectorInsertHelper(factory,vals,i,t)
		}
	}
}

func vectorInsertSequentialHelper(
	v func() dynamicContainers.Vector[int],
	idx int,
	l int,
	t *testing.T,
) {
	container := v()
	for i := 0; i < l-1; i++ {
		container.Append(i)
	}
	err := container.InsertSequential(idx, l-1)
	test.Nil(err,t)
	test.Eq(l, container.Length(),t)
	for i := 0; i < container.Length(); i++ {
		var exp int
		v, _ := container.Get(i)
		if i < idx {
			exp = i
		} else if i == idx {
			exp = l - 1
		} else {
			exp = i - 1
		}
		test.Eq(exp, v,t)
	}
}

// Tests the InsertSequential method functionality of a dynamic vector.
func VectorInterfaceInsertSequential(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	container := factory()
	for i := 2; i >= 0; i-- {
		container.InsertSequential(0,i)
	}
	for i := 3; i < 5; i++ {
		container.InsertSequential(container.Length(), i)
	}
	for i := 0; i < 5; i++ {
		iterV, _ := container.Get(i)
		test.Eq(i, iterV,t)
	}
	container = factory()
	container.InsertSequential(0,0,1,2)
	container.InsertSequential(3,4,5)
	container.InsertSequential(3,3)
	for i := 0; i < 6; i++ {
		iterV, _ := container.Get(i)
		test.Eq(i, iterV,t)
	}
	for i := 0; i < 5; i++ {
		vectorInsertSequentialHelper(factory, i, 5, t)
	}
}

func vectorPopSequentialHelper(
	factory func() dynamicContainers.Vector[int],
	l int,
	num int,
	t *testing.T,
) {
	// fmt.Println("Permutation: l: ",l," num: ",num)
	container := factory()
	for i := 0; i < l; i++ {
		if i%4 == 0 {
			container.Append(-1)
		} else {
			container.Append(i)
		}
	}
	// fmt.Println("Init:   ",v)
	n := container.PopSequential(-1, num)
	exp := factory()
	cntr := 0
	for i := 0; i < l; i++ {
		if i%4 == 0 {
			if cntr < num {
				cntr++
				continue
			} else {
				exp.Append(-1)
			}
		} else {
			exp.Append(i)
		}
	}
	test.Eq(exp.Length(), container.Length(),t)
	test.Eq(cntr, n,t)
	// fmt.Println("EXP:    ",exp)
	// fmt.Println("Final:  ",v)
	for i := 0; i < container.Length(); i++ {
		iterV, found := container.Get(i)
		expIterV, foundExp := exp.Get(i)
		test.Nil(found,t)
		test.Nil(foundExp,t)
		test.Eq(expIterV, iterV,t)
	}
}

// Tests the PopSequential method functionality of a dynamic vector.
func VectorInterfacePopSequential(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	for i := 0; i < 13; i++ {
		for j := 0; j < 13; j++ {
			vectorPopSequentialHelper(factory, i, j, t)
		}
	}
}

func vectorPopHelper(
	factory func() dynamicContainers.Vector[int],
	l int,
	t *testing.T,
) {
	// fmt.Println("Permutation: l: ",l," num: ",num)
	container := factory()
	for i := 0; i < l; i++ {
		if i%4 == 0 {
			container.Append(-1)
		} else {
			container.Append(i)
		}
	}
	// fmt.Println("Init:   ",v)
	n := container.Pop(-1)
	exp := factory()
	cntr := 0
	for i := 0; i < l; i++ {
		if i%4 != 0 {
			exp.Append(i)
		} else {
			cntr++
		}
	}
	test.Eq(exp.Length(), container.Length(),t)
	test.Eq(cntr, n,t)
	// fmt.Println("EXP:    ",exp)
	// fmt.Println("Final:  ",v)
	for i := 0; i < container.Length(); i++ {
		iterV, found := container.Get(i)
		expIterV, foundExp := exp.Get(i)
		test.Nil(found,t)
		test.Nil(foundExp,t)
		test.Eq(expIterV, iterV,t)
	}
}

// Tests the Pop method functionality of a dynamic vector.
func VectorInterfacePop(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	for i := 0; i < 13; i++ {
		vectorPopHelper(factory, i, t)
	}
}

// Tests the Delete method functionality of a dynamic vector.
func VectorInterfaceDelete(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 6; i++ {
		container.Append(i)
	}
	for i := container.Length() - 1; i >= 0; i-- {
		container.Delete(i)
		test.Eq(i, container.Length(), t)
		for j := 0; j < i; j++ {
			iterV, _ := container.Get(j)
			test.Eq(j, iterV, t)
		}
	}
	err := container.Delete(0)
	test.ContainsError(customerr.ValOutsideRange, err,t)
}

// Tests the Clear method functionality of a dynamic vector.
func VectorInterfaceClear(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 6; i++ {
		container.Append(i)
	}
	container.Clear()
	test.Eq(0, container.Length(), t)
	test.Eq(0, container.Capacity(), t)
}

func testVectorValsHelper(
    factory func() dynamicContainers.Vector[int],
    l int, 
    t *testing.T,
){
	container:=factory()
    for i:=0; i<l; i++ {
        container.Append(i);
    }
    cnt:=0
    container.Vals().ForEach(func(index, val int) (iter.IteratorFeedback, error) {
        cnt++
        test.Eq(index,val,t);
        return iter.Continue,nil;
    });
    test.Eq(l,cnt,t)
}
// Tests the Vals method functionality of a dynamic vector.
func VectorInterfaceVals(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
){
    testVectorValsHelper(factory,0,t);
    testVectorValsHelper(factory,1,t);
    testVectorValsHelper(factory,2,t);
    testVectorValsHelper(factory,5,t);
}

func testVectorPntrValsHelper(
	factory func() dynamicContainers.Vector[int],
    l int, 
    t *testing.T,
){
	container:=factory()
    for i:=0; i<l; i++ {
        container.Append(i);
    }
    cnt:=0
    container.ValPntrs().ForEach(func(index int, val *int) (iter.IteratorFeedback, error) {
        cnt++
        test.Eq(index,*val,t);
        *val=100;
        return iter.Continue,nil;
    });
    container.Vals().ForEach(func(index int, val int) (iter.IteratorFeedback, error) {
        test.Eq(100,val,t);
        return iter.Continue,nil;
    });
    test.Eq(l,cnt,t)
}
// Tests the ValPntrs method functionality of a dynamic vector.
func VectorInterfaceValPntrs(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
){
	container:=factory()
	if container.IsAddressable() {
		testVectorPntrValsHelper(factory,0,t);
    	testVectorPntrValsHelper(factory,1,t);
    	testVectorPntrValsHelper(factory,2,t);
	} else {
		test.Panics(
			func() { container.ValPntrs() },
			t,
		)
	}
}

func testVectorKeysHelper(
    factory func() dynamicContainers.Vector[int],
    l int, 
    t *testing.T,
){
	container:=factory()
    for i:=0; i<l; i++ {
        container.Append(i);
    }
    cnt:=0
    container.Keys().ForEach(func(index, val int) (iter.IteratorFeedback, error) {
        cnt++
        test.Eq(index,val,t);
        return iter.Continue,nil;
    });
    test.Eq(l,cnt,t)
}
func VectorInterfaceKeys(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
){
    testVectorKeysHelper(factory,0,t);
    testVectorKeysHelper(factory,1,t);
    testVectorKeysHelper(factory,2,t);
    testVectorKeysHelper(factory,5,t);
}

// Tests the KeyedEq method functionality of a dynamic vector.
func VectorInterfaceKeyedEq(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	v := factory()
	v.Append(1, 2, 3)
	v2 := factory()
	v2.Append(1, 2, 3)
	test.True(v.KeyedEq(v2),t)
	test.True(v2.KeyedEq(v),t)
	v.Pop(3)
	test.False(v.KeyedEq(v2), t)
	test.False(v2.KeyedEq(v), t)
	v.Append(3)
	v2 = factory()
	v2.Append(3, 1, 2)
	test.False(v.KeyedEq(v2), t)
	test.False(v2.KeyedEq(v), t)
	v.Pop(3)
	test.False(v.KeyedEq(v2), t)
	test.False(v2.KeyedEq(v), t)
	v = factory()
	v.Append(0)
	v2 = factory()
	v2.Append(0)
	test.True(v.KeyedEq(v2), t)
	test.True(v2.KeyedEq(v), t)
	v.Pop(0)
	test.False(v.KeyedEq(v2), t)
	test.False(v2.KeyedEq(v), t)
	v = factory()
	v2 = factory()
	test.True(v.KeyedEq(v2), t)
	test.True(v2.KeyedEq(v), t)
}
