package tests

import (
	"testing"

	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/container/staticContainers"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/test"
)

func staticVectorReadInterface[U any](c staticContainers.ReadVector[U])   {}
func staticVectorWriteInterface[U any](c staticContainers.WriteVector[U]) {}
func staticVectorInterface[U any](c staticContainers.Vector[U])           {}

// Tests that the value supplied by the factory implements the
// [containerTypes.Addressable] interface.
func StaticVectorInterfaceAddressableInterface[V any](
	factory func(capacity int) staticContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.Addressable = factory(1)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Length] interface.
func StaticVectorInterfaceLengthInterface[V any](
	factory func(capacity int) staticContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.Length = factory(1)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Capacity] interface.
func StaticVectorInterfaceCapacityInterface[V any](
	factory func(capacity int) staticContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.Capacity = factory(1)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Clear] interface.
func StaticVectorInterfaceClearInterface[V any](
	factory func(capacity int) staticContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.Clear = factory(1)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.WriteOps] interface.
func StaticVectorInterfaceWriteOpsInterface[V any](
	factory func(capacity int) staticContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.WriteOps[V] = factory(1)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.WriteKeyedOps] interface.
func StaticVectorInterfaceWriteKeyedOpsInterface[V any](
	factory func(capacity int) staticContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.WriteKeyedOps[int, V] = factory(1)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.WriteKeyedSequentialOps] interface.
func StaticVectorInterfaceWriteKeyedSequentialOpsInterface[V any](
	factory func(capacity int) staticContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.WriteKeyedSequentialOps[int, V] = factory(1)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.KeyedWriteOps] interface.
func StaticVectorInterfaceWriteDynKeyedOpsInterface[V any](
	factory func(capacity int) staticContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.WriteDynKeyedOps[int, V] = factory(1)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.ReadOps] interface.
func StaticVectorInterfaceReadOpsInterface[V any](
	factory func(capacity int) staticContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.ReadOps[V] = factory(1)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.KeyedReadOps] interface.
func StaticVectorInterfaceReadKeyedOpsInterface[V any](
	factory func(capacity int) staticContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.ReadKeyedOps[int, V] = factory(1)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.DeleteOps] interface.
func StaticVectorInterfaceDeleteOpsInterface[V any](
	factory func(capacity int) staticContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.DeleteOps[int, V] = factory(1)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.KeyedDeleteOps] interface.
func StaticVectorInterfaceDeleteKeyedOpsInterface[V any](
	factory func(capacity int) staticContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.DeleteKeyedOps[int, V] = factory(1)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.KeyedDeleteOps] interface.
func StaticVectorInterfaceDeleteSequentialOpsInterface[V any](
	factory func(capacity int) staticContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.DeleteSequentialOps[int, V] = factory(1)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.KeyedDeleteOps] interface.
func StaticVectorInterfaceDeleteKeyedSequentialOpsInterface[V any](
	factory func(capacity int) staticContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.DeleteKeyedSequentialOps[int, V] = factory(1)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [staticContainers.VectorRead] interface.
func ReadStaticVectorInterface[V any](
	factory func(capacity int) staticContainers.Vector[V],
	t *testing.T,
) {
	staticVectorReadInterface[V](factory(1))
}

// Tests that the value supplied by the factory implements the
// [staticContainers.WriteVector] interface.
func WriteStaticVectorInterface[V any](
	factory func(capacity int) staticContainers.Vector[V],
	t *testing.T,
) {
	staticVectorWriteInterface[V](factory(1))
}

// Tests that the value supplied by the factory implements the
// [staticContainers.Vector] interface.
func StaticVectorInterfaceInterface[V any](
	factory func(capacity int) staticContainers.Vector[V],
	t *testing.T,
) {
	staticVectorInterface[V](factory(1))
}

// Tests the Get method functionality of a static vector.
func StaticVectorInterfaceGet(
	factory func(capacity int) staticContainers.Vector[int],
	t *testing.T,
) {
	container := factory(5)
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

// Tests the GetPntr method functionality of a static vector.
func StaticVectorInterfaceGetPntr(
	factory func(capacity int) staticContainers.Vector[int],
	t *testing.T,
) {
	container := factory(5)
	if container.IsAddressable() {
		v, err := container.GetPntr(0)
		test.ContainsError(customerr.ValOutsideRange, err,t)
		test.NilPntr[int](v,t)
		for i := 0; i < 5; i++ {
			container.Append(i)
		}
		for i := 0; i < 5; i++ {
			_v, err := container.GetPntr(i)
			test.Eq(nil, err,t)
			test.Eq(i, *_v,t)
		}
		v, err = container.GetPntr(-1)
		test.ContainsError(customerr.ValOutsideRange, err,t)
		test.NilPntr[int](v,t)
		v, err = container.GetPntr(6)
		test.ContainsError(customerr.ValOutsideRange, err,t)
		test.NilPntr[int](v,t)
	} else {
		test.Panics(
			func() {
				container:=factory(0)
				container.GetPntr(1)
			},
			t,
		)
	}
}

func staticVectorContainsHelper(
	v staticContainers.Vector[int],
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

// Tests the Contains method functionality of a static vector.
func StaticVectorInterfaceContains(
	factory func(capacity int) staticContainers.Vector[int],
	t *testing.T,
) {
	staticVectorContainsHelper(factory(0), 0, t)
	staticVectorContainsHelper(factory(10), 0, t)
	staticVectorContainsHelper(factory(1), 1, t)
	staticVectorContainsHelper(factory(10), 1, t)
	staticVectorContainsHelper(factory(2), 2, t)
	staticVectorContainsHelper(factory(10), 2, t)
	staticVectorContainsHelper(factory(5), 5, t)
	staticVectorContainsHelper(factory(10), 5, t)
}

func staticVectorContainsPntrHelper(
	v staticContainers.Vector[int],
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

// Tests the ContainsPntr method functionality of a static vector.
func StaticVectorInterfaceContainsPntr(
	factory func(capacity int) staticContainers.Vector[int],
	t *testing.T,
) {
	staticVectorContainsPntrHelper(factory(0), 0, t)
	staticVectorContainsPntrHelper(factory(10), 0, t)
	staticVectorContainsPntrHelper(factory(1), 1, t)
	staticVectorContainsPntrHelper(factory(10), 1, t)
	staticVectorContainsPntrHelper(factory(2), 2, t)
	staticVectorContainsPntrHelper(factory(10), 2, t)
	staticVectorContainsPntrHelper(factory(5), 5, t)
	staticVectorContainsPntrHelper(factory(10), 5, t)
}

func staticVectorKeyOfHelper(
	v staticContainers.Vector[int],
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
	_, found = v.KeyOf(l)
	test.False(found,t)
	test.False(v.Contains(l),t)
}

// Tests the KeyOf method functionality of a static vector.
func StaticVectorInterfaceKeyOf(
	factory func(capacity int) staticContainers.Vector[int],
	t *testing.T,
) {
	staticVectorKeyOfHelper(factory(0), 0, t)
	staticVectorKeyOfHelper(factory(10), 0, t)
	staticVectorKeyOfHelper(factory(1), 1, t)
	staticVectorKeyOfHelper(factory(10), 1, t)
	staticVectorKeyOfHelper(factory(2), 2, t)
	staticVectorKeyOfHelper(factory(10), 2, t)
	staticVectorKeyOfHelper(factory(5), 5, t)
	staticVectorKeyOfHelper(factory(10), 5, t)
}

// Tests the Set method functionality of a static vector.
func StaticVectorInterfaceSet(
	factory func(capacity int) staticContainers.Vector[int],
	t *testing.T,
) {
	container := factory(5)
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
		iterV, err := container.Get(i)
		test.Nil(err,t)
		test.Eq(i+1, iterV,t)
	}
	err = container.Set(basic.Pair[int,int]{-1, 6})
	test.ContainsError(customerr.ValOutsideRange, err,t)
	err = container.Set(basic.Pair[int,int]{6, 6})
	test.ContainsError(customerr.ValOutsideRange, err,t)
}

func staticVectorSetSequentialHelper(
	factory func(capacity int) staticContainers.Vector[int],
	idx int,
	vals []int,
	l int,
	t *testing.T,
){
	container:=factory(l)
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
// Tests the SetSequential method functionality of a static vector.
func StaticVectorInterfaceSetSequential(
	factory func(capacity int) staticContainers.Vector[int],
	t *testing.T,
){
	for i:=0; i<20; i++ {
		for j:=0; j<i; j++ {
			vals:=[]int{}
			for k:=0; k<i-j; k++ {
				vals = append(vals, k)
				staticVectorSetSequentialHelper(factory,j,vals,i,t)
			}
		}
	}
}

// Tests the Append method functionality of a static vector.
func StaticVectorInterfaceAppend(
	factory func(capacity int) staticContainers.Vector[int],
	t *testing.T,
) {
	container := factory(0)
	err:=container.Append(0)
	test.ContainsError(containerTypes.Full,err,t)
	container=factory(5)
	for i := 0; i < 5; i++ {
		err:=container.Append(i)
		test.Nil(err,t)
	}
	for i := 0; i < 5; i++ {
		iterV, err := container.Get(i)
		test.Nil(err,t)
		test.Eq(i, iterV,t)
	}
	err=container.Append(5, 6, 7)
	test.ContainsError(containerTypes.Full,err,t)
	test.Eq(5,container.Length(),t)
	for i := 0; i < 5; i++ {
		iterV, err := container.Get(i)
		test.Nil(err,t)
		test.Eq(i, iterV,t)
	}
}

func staticVectorInsertHelper(
	factory func(capacity int) staticContainers.Vector[int],
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
	container:=factory(l)
	for i:=0; i<l; i++ {
		err:=container.Append(i)
		test.Nil(err,t)
	}
	err:=container.Insert(vals...)
	if len(vals)>0 {
		test.ContainsError(containerTypes.Full,err,t)
	} else {
		test.Nil(err,t)
	}
	container=factory(l+len(vals))
	for i:=0; i<l; i++ {
		container.Append(i)
	}
	err=container.Insert(vals...)
	test.Nil(err,t)
	test.Eq(l+len(vals),container.Length(),t)
	offset:=0
	for i:=0; i<container.Length(); i++ {
		iterV,err:=container.Get(i)
		test.Nil(err,t)
		if v,ok:=indexContained(i); ok {
			test.Eq(v.B,iterV,t)
			offset++
		} else {
			test.Eq(i-offset,iterV,t)
		}
	}
	err=container.Insert(basic.Pair[int, int]{l+len(vals)+1,-1})
	test.ContainsError(customerr.ValOutsideRange,err,t)
	err=container.Insert(basic.Pair[int, int]{-1,-1})
	test.ContainsError(customerr.ValOutsideRange,err,t)
}
// Tests the Insert method functionality of a static vector.
func StaticVectorInterfaceInsert(
	factory func(capacity int) staticContainers.Vector[int],
	t *testing.T,
){
	container:=factory(0)
	err:=container.Insert(basic.Pair[int, int]{0,0})
	test.ContainsError(containerTypes.Full,err,t)
	err=container.Insert(basic.Pair[int, int]{1,0})
	test.ContainsError(containerTypes.KeyError,err,t)
	for i:=0; i<20; i++ {
		vals:=[]basic.Pair[int,int]{}
		staticVectorInsertHelper(factory,vals,i,t)
		for j:=0; j<20; j++ {
			vals = append(vals, basic.Pair[int, int]{j,j})
			staticVectorInsertHelper(factory,vals,i,t)
		}
	}
	for i:=0; i<20; i++ {
		vals:=[]basic.Pair[int,int]{}
		staticVectorInsertHelper(factory,vals,i,t)
		for j:=0; j<i; j+=2 {
			vals = append(vals, basic.Pair[int, int]{j,j})
			staticVectorInsertHelper(factory,vals,i,t)
		}
	}
}

func staticVectorInsertSequentialHelper(
	factory func(capacity int) staticContainers.Vector[int],
	idx int,
	l int,
	t *testing.T,
) {
	container := factory(l)
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
	err=container.InsertSequential(l+1,-1)
	test.ContainsError(customerr.ValOutsideRange,err,t)
	test.Eq(l,container.Length(),t)
	err=container.InsertSequential(-1,-1)
	test.ContainsError(customerr.ValOutsideRange,err,t)
	test.Eq(l,container.Length(),t)
	err=container.InsertSequential(l,-1)
	test.ContainsError(containerTypes.Full,err,t)
	test.Eq(l,container.Length(),t)
}
// Tests the InsertSequential method functionality of a static vector.
func StaticVectorInterfaceInsertSequential(
	factory func(capacity int) staticContainers.Vector[int],
	t *testing.T,
) {
	container := factory(5)
	for i := 2; i >= 0; i-- {
		err:=container.InsertSequential(0,i)
		test.Nil(err,t)
	}
	for i := 3; i < 5; i++ {
		err:=container.InsertSequential(container.Length(), i)
		test.Nil(err,t)
	}
	for i := 0; i < 5; i++ {
		iterV, err := container.Get(i)
		test.Nil(err,t)
		test.Eq(i, iterV,t)
	}
	container = factory(6)
	container.InsertSequential(0,0,1,2)
	container.InsertSequential(3,4,5)
	container.InsertSequential(3,3)
	for i := 0; i < 6; i++ {
		iterV, err := container.Get(i)
		test.Nil(err,t)
		test.Eq(i, iterV,t)
	}
	for i := 0; i < 5; i++ {
		staticVectorInsertSequentialHelper(factory, i, 5, t)
	}
}

func staticVectorPopSequentialHelper(
	factory func(capacity int) staticContainers.Vector[int],
	l int,
	num int,
	t *testing.T,
) {
	permutation:=func(container staticContainers.Vector[int]) {
		// fmt.Println("Init:   ",container)
		n := container.PopSequential(-1, num)
		cntr := 0
		expLen:=0
		for i := 0; i < l; i++ {
			if i%4 == 0 {
				if cntr < num {
					cntr++
					continue
				} else {
					v,err:=container.Get(i-cntr)
					test.Nil(err,t)
					test.Eq(-1,v,t)
					expLen++
				}
			} else {
				v,err:=container.Get(i-cntr)
				test.Nil(err,t)
				test.Eq(i,v,t)
				expLen++
			}
		}
		test.Eq(expLen, container.Length(),t)
		test.Eq(cntr, n,t)
	}
	// fmt.Println("Permutation: l: ",l," num: ",num)
	container := factory(l)
	for i := 0; i < l; i++ {
		if i%4 == 0 {
			container.Append(-1)
		} else {
			container.Append(i)
		}
	}
	permutation(container)
	container = factory(l+1)
	for i := 0; i < l; i++ {
		if i%4 == 0 {
			container.Append(-1)
		} else {
			container.Append(i)
		}
	}
	permutation(container)
	container = factory(l+10)
	for i := 0; i < l; i++ {
		if i%4 == 0 {
			container.Append(-1)
		} else {
			container.Append(i)
		}
	}
	permutation(container)
}
// Tests the PopSequential method functionality of a static vector.
func StaticVectorInterfacePopSequential(
	factory func(capacity int) staticContainers.Vector[int],
	t *testing.T,
) {
	for i := 0; i < 13; i++ {
		for j := 0; j < 13; j++ {
			staticVectorPopSequentialHelper(factory, i, j, t)
		}
	}
}

func staticVectorPopHelper(
	factory func(capacity int) staticContainers.Vector[int],
	l int,
	t *testing.T,
) {
	permutation:=func(container staticContainers.Vector[int]){
		// fmt.Println("Init:   ",v)
		n := container.Pop(-1)
		cntr := 0
		expLength:=0
		for i := 0; i < l; i++ {
			if i%4 != 0 {
				v,err:=container.Get(i-cntr)
				test.Nil(err,t)
				test.Eq(i,v,t)
				expLength++
			} else {
				cntr++
			}
		}
		test.Eq(expLength, container.Length(),t)
		test.Eq(cntr, n,t)
	}
	// fmt.Println("Permutation: l: ",l," num: ",num)
	container := factory(l)
	for i := 0; i < l; i++ {
		if i%4 == 0 {
			container.Append(-1)
		} else {
			container.Append(i)
		}
	}
	permutation(container)
	container = factory(l+1)
	for i := 0; i < l; i++ {
		if i%4 == 0 {
			container.Append(-1)
		} else {
			container.Append(i)
		}
	}
	permutation(container)
	container = factory(l+10)
	for i := 0; i < l; i++ {
		if i%4 == 0 {
			container.Append(-1)
		} else {
			container.Append(i)
		}
	}
	permutation(container)
}
// Tests the Pop method functionality of a static vector.
func StaticVectorInterfacePop(
	factory func(capacity int) staticContainers.Vector[int],
	t *testing.T,
) {
	for i := 0; i < 13; i++ {
		staticVectorPopHelper(factory, i, t)
	}
}

// Tests the Delete method functionality of a static vector.
func StaticVectorInterfaceDelete(
	factory func(capacity int) staticContainers.Vector[int],
	t *testing.T,
) {
	container := factory(6)
	err:=container.Delete(0)
	test.ContainsError(customerr.ValOutsideRange,err,t)
	for i := 0; i < 6; i++ {
		container.Append(i)
	}
	err=container.Delete(-1)
	test.ContainsError(customerr.ValOutsideRange,err,t)
	err=container.Delete(7)
	test.ContainsError(customerr.ValOutsideRange,err,t)
	for i := container.Length() - 1; i >= 0; i-- {
		container.Delete(i)
		test.Eq(i, container.Length(), t)
		for j := 0; j < i; j++ {
			iterV, err := container.Get(j)
			test.Nil(err,t)
			test.Eq(j, iterV, t)
		}
	}
	err = container.Delete(0)
	test.ContainsError(customerr.ValOutsideRange, err,t)
}

// func staticVectorDeleteSequentialHelper(
// 	factory func(capacity int) staticContainers.Vector[int],
// 	start int,
// 	end int,
// 	l int,
// 	t *testing.T,
// ){
// 	container:=factory(0)
// 	for i:=0; i<l; i++ {
// 		container.Append(i)
// 	}
// 	container.DeleteSequential(start,end)
// 	test.Eq(l-(end-start),container.Length(),t)
// 	for i:=0; i<l; i++ {
// 		if i<start {
// 			v,err:=container.Get(i)
// 			test.Nil(err,nil)
// 			test.Eq(i,v,t)
// 		} else if i>=end {
// 			v,err:=container.Get(i-(end-start))
// 			test.Nil(err,nil)
// 			test.Eq(i,v,t)
// 		} else {
// 			test.False(container.Contains(i),t)
// 		}
// 	}
// }
// // Tests the DeleteSequential method functionality of a static vector.
// func StaticVectorInterfaceDeleteSequential(
// 	factory func(capacity int) staticContainers.Vector[int],
// 	t *testing.T,
// ){
// 	container:=factory(0)
// 	container.Append(0,1,2,3)
// 	err:=container.DeleteSequential(-1,3)
// 	test.ContainsError(customerr.ValOutsideRange,err,t)
// 	test.Eq(4,container.Length(),t)
// 	err=container.DeleteSequential(0,4)
// 	test.ContainsError(customerr.ValOutsideRange,err,t)
// 	test.Eq(4,container.Length(),t)
// 	err=container.DeleteSequential(2,1)
// 	test.ContainsError(customerr.InvalidValue,err,t)
// 	test.Eq(4,container.Length(),t)
// 	for i:=0; i<20; i++ {
// 		for j:=0; j<i; j++ {
// 			for k:=j; k<i; k++ {
// 				staticVectorDeleteSequentialHelper(factory,j,k,i,t)
// 			}
// 		}
// 	}
// }
