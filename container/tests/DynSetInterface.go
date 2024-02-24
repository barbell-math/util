package tests

import (
	"testing"

	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/container/dynamicContainers"
	"github.com/barbell-math/util/test"
)

func setReadInterface[U any](c dynamicContainers.ReadSet[U])   {}
func setWriteInterface[U any](c dynamicContainers.WriteSet[U]) {}
func setInterface[U any](c dynamicContainers.Set[U])           {}

// Tests that the value supplied by the factory implements the 
// [containerTypes.Length] interface.
func SetInterfaceLengthInterface[V any](
	factory func() dynamicContainers.Set[V],
	t *testing.T,
) {
	var container containerTypes.Length = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.Clear] interface.
func SetInterfaceClearInterface[V any](
	factory func() dynamicContainers.Set[V],
	t *testing.T,
) {
	var container containerTypes.Clear = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.WriteUniqueOps] interface.
func SetInterfaceWriteUniqueOpsInterface[V any](
	factory func() dynamicContainers.Set[V],
	t *testing.T,
) {
	var container containerTypes.WriteUniqueOps[uint64, V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.ReadOps] interface.
func SetInterfaceReadOpsInterface[V any](
	factory func() dynamicContainers.Set[V],
	t *testing.T,
) {
	var container containerTypes.ReadOps[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.DeleteOps] interface.
func SetInterfaceDeleteOpsInterface[V any](
	factory func() dynamicContainers.Set[V],
	t *testing.T,
) {
	var container containerTypes.DeleteOps[uint64, V] = factory()
	_ = container
}

// Tests that the value supplied by the factory does not implement the 
// [containerTypes.StaticCapacity] interface.
func SetInterfaceStaticCapacityInterface[V any](
	factory func() dynamicContainers.Set[V],
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

// Tests the ContainsPntr method functionality of a dynamic set.
func SetInterfaceContainsPntr(
	factory func() dynamicContainers.Set[int],
	t *testing.T,
){
	container:=factory()
	for i:=0; i<5; i++ {
		container.AppendUnique(i)
	}
	for i:=0; i<5; i++ {
		test.True(container.ContainsPntr(&i),t)
	}
	v:=5
	test.False(container.ContainsPntr(&v),t)
	v=-1
	test.False(container.ContainsPntr(&v),t)
	container.Pop(0,1)
	v=0
	test.False(container.ContainsPntr(&v),t)
	for i:=1; i<5; i++ {
		test.True(container.ContainsPntr(&i),t)
	}
}

// Tests the Contains method functionality of a dynamic set.
func SetInterfaceContains(
	factory func() dynamicContainers.Set[int],
	t *testing.T,
){
	container:=factory()
	for i:=0; i<5; i++ {
		container.AppendUnique(i)
	}
	for i:=0; i<5; i++ {
		test.True(container.Contains(i),t)
	}
	test.False(container.Contains(5),t)
	test.False(container.Contains(-1),t)
	container.Pop(0,1)
	test.False(container.Contains(0),t)
	for i:=1; i<5; i++ {
		test.True(container.Contains(i),t)
	}
}

// Tests the Clear method functionality of a dynamic set.
func SetInterfaceClear(
	factory func() dynamicContainers.Set[int],
	t *testing.T,
){
	container:=factory()
	test.Eq(0,container.Length(),t)
	for i:=0; i<5; i++ {
		container.AppendUnique(i)
	}
	test.Eq(5,container.Length(),t)
	container.Clear()
	test.Eq(0,container.Length(),t)
}

// Tests the AppendUnique method functionality of a dynamic set.
func SetInterfaceAppendUnique(
	factory func() dynamicContainers.Set[int],
	t *testing.T,
){
	container:=factory()
	test.Eq(0,container.Length(),t)
	for i:=0; i<5; i++ {
		err:=container.AppendUnique(i)
		test.Nil(err,t)
	}
	for i:=0; i<5; i++ {
		test.True(container.Contains(i),t)
	}
	for i:=0; i<5; i++ {
		container.AppendUnique(i)
		test.Eq(5,container.Length(),t)
	}
	container=factory()
	test.Eq(0,container.Length(),t)
	for i:=0; i<6; i+=2 {
		container.AppendUnique(i,i+1)
	}
	for i:=0; i<6; i++ {
		test.True(container.Contains(i),t)
	}
}

// Tests the Pop method functionality of a dynamic set.
func SetInterfacePop(
	factory func() dynamicContainers.Set[int],
	t *testing.T,
){
	container:=factory()
	test.Eq(0,container.Length(),t)
	for i:=0; i<5; i++ {
		container.AppendUnique(i)
	}
	for i:=0; i<5; i++ {
		test.True(container.Contains(i),t)
		container.Pop(i,1)
		test.False(container.Contains(i),t)
	}
}

// Tests the UnorderedEq method functionality of a dynamic set.
func SetInterfaceUnorderedEq(
	factory func() dynamicContainers.Set[int],
	t *testing.T,
) {
	v := factory()
	v.AppendUnique(1, 2, 3)
	v2 := factory()
	v2.AppendUnique(1, 2, 3)
	test.True(v.UnorderedEq(v2),t)
	test.True(v2.UnorderedEq(v),t)
	v.Pop(3,1)
	test.False(v.UnorderedEq(v2),t)
	test.False(v2.UnorderedEq(v),t)
	v.AppendUnique(3)
	v2 = factory()
	v2.AppendUnique(3, 1, 2)
	test.True(v.UnorderedEq(v2),t)
	test.True(v2.UnorderedEq(v),t)
	v.Pop(3,1)
	test.False(v.UnorderedEq(v2),t)
	test.False(v2.UnorderedEq(v),t)
	v.AppendUnique(3)
	v2 = factory()
	v2.AppendUnique(2, 3, 1)
	test.True(v.UnorderedEq(v2),t)
	test.True(v2.UnorderedEq(v),t)
	v.Pop(3,1)
	test.False(v.UnorderedEq(v2),t)
	test.False(v2.UnorderedEq(v),t)
	v = factory()
	v.AppendUnique(0)
	v2 = factory()
	v2.AppendUnique(0)
	test.True(v.UnorderedEq(v2),t)
	test.True(v2.UnorderedEq(v),t)
	v.Pop(0,1)
	test.False(v.UnorderedEq(v2),t)
	test.False(v2.UnorderedEq(v),t)
	v = factory()
	v2 = factory()
	test.True(v.UnorderedEq(v2),t)
	test.True(v2.UnorderedEq(v),t)
}

func setIntersectionHelper(
	res dynamicContainers.Set[int],
	l dynamicContainers.Set[int],
	r dynamicContainers.Set[int],
	exp []int,
	factory func() dynamicContainers.Set[int],
	t *testing.T,
){
	tester:=func(c dynamicContainers.Set[int]) {
		test.Eq(len(exp),c.Length(),t)
		for _,v:=range(exp) {
			test.True(c.Contains(v),t)
		}
	}
	res.Intersection(l,r)
	tester(res)
	res.Intersection(r,l)
	tester(res)
}

// Tests the Intersection method functionality of a dynamic set.
func SetInterfaceIntersection(
	factory func() dynamicContainers.Set[int],
	t *testing.T,
) {
	v:=factory()
	v2:=factory()
	setIntersectionHelper(factory(),v,v2,[]int{},factory,t)
	v.AppendUnique(1)
	setIntersectionHelper(factory(),v,v2,[]int{},factory,t)
	v2.AppendUnique(1)
	setIntersectionHelper(factory(),v,v2,[]int{1},factory,t)
	v2.AppendUnique(2)
	setIntersectionHelper(factory(),v,v2,[]int{1},factory,t)
	v.AppendUnique(2)
	setIntersectionHelper(factory(),v,v2,[]int{1,2},factory,t)
	v.AppendUnique(3)
	setIntersectionHelper(factory(),v,v2,[]int{1,2},factory,t)
	v2.AppendUnique(3)
	setIntersectionHelper(factory(),v,v2,[]int{1,2,3},factory,t)

	if !v.IsSynced() {
		v=factory()
		v2=factory()
		v.AppendUnique(1,2,3,4)
		v2.AppendUnique(2,4)
		setIntersectionHelper(v,v,v2,[]int{2,4},factory,t)
	}
}

func setUnionHelper(
	res dynamicContainers.Set[int],
	l dynamicContainers.Set[int],
	r dynamicContainers.Set[int],
	exp []int,
	factory func() dynamicContainers.Set[int],
	t *testing.T,
){
	tester:=func(c dynamicContainers.Set[int]) {
		test.Eq(len(exp),c.Length(),t)
		for _,v:=range(exp) {
			test.True(c.Contains(v),t)
		}
	}
	res.Union(l,r)
	tester(res)
	res.Union(r,l)
	tester(res)
}

// Tests the Union method functionality of a dynamic set.
func SetInterfaceUnion(
	factory func() dynamicContainers.Set[int],
	t *testing.T,
) {
	v:=factory()
	v2:=factory()
	setUnionHelper(factory(),v,v2,[]int{},factory,t)
	v.AppendUnique(1)
	setUnionHelper(factory(),v,v2,[]int{1},factory,t)
	v2.AppendUnique(1)
	setUnionHelper(factory(),v,v2,[]int{1},factory,t)
	v2.AppendUnique(2)
	setUnionHelper(factory(),v,v2,[]int{1,2},factory,t)
	v.AppendUnique(2)
	setUnionHelper(factory(),v,v2,[]int{1,2},factory,t)
	v.AppendUnique(3)
	setUnionHelper(factory(),v,v2,[]int{1,2,3},factory,t)
	v2.AppendUnique(3)
	setUnionHelper(factory(),v,v2,[]int{1,2,3},factory,t)

	if !v.IsSynced() {
		v=factory()
		v2=factory()
		v.AppendUnique(1,2,3,4)
		v2.AppendUnique(2,4,5,6)
		setUnionHelper(v,v,v2,[]int{1,2,3,4,5,6},factory,t)
	}
}

func setDifferenceHelper(
	res dynamicContainers.Set[int],
	l dynamicContainers.Set[int],
	r dynamicContainers.Set[int],
	exp []int,
	factory func() dynamicContainers.Set[int],
	t *testing.T,
){
	res.Difference(l,r)
	test.Eq(len(exp),res.Length(),t)
	for _,v:=range(exp) {
		test.True(res.Contains(v),t)
	}
}

// Tests the Difference method functionality of a dynamic set.
func SetInterfaceDifference(
	factory func() dynamicContainers.Set[int],
	t *testing.T,
) {
	v:=factory()
	v2:=factory()
	setDifferenceHelper(factory(),v,v2,[]int{},factory,t)
	setDifferenceHelper(factory(),v2,v,[]int{},factory,t)
	v.AppendUnique(1)
	setDifferenceHelper(factory(),v,v2,[]int{1},factory,t)
	setDifferenceHelper(factory(),v2,v,[]int{},factory,t)
	v2.AppendUnique(1)
	setDifferenceHelper(factory(),v,v2,[]int{},factory,t)
	setDifferenceHelper(factory(),v2,v,[]int{},factory,t)
	v2.AppendUnique(2)
	setDifferenceHelper(factory(),v,v2,[]int{},factory,t)
	setDifferenceHelper(factory(),v2,v,[]int{2},factory,t)
	v.AppendUnique(2)
	setDifferenceHelper(factory(),v,v2,[]int{},factory,t)
	setDifferenceHelper(factory(),v2,v,[]int{},factory,t)
	v.AppendUnique(3)
	setDifferenceHelper(factory(),v,v2,[]int{3},factory,t)
	setDifferenceHelper(factory(),v2,v,[]int{},factory,t)
	v2.AppendUnique(3)
	setDifferenceHelper(factory(),v,v2,[]int{},factory,t)
	setDifferenceHelper(factory(),v2,v,[]int{},factory,t)
	v.AppendUnique(4,5,6)
	setDifferenceHelper(factory(),v,v2,[]int{4,5,6},factory,t)
	setDifferenceHelper(factory(),v2,v,[]int{},factory,t)

	if !v.IsSynced() {
		v=factory()
		v2=factory()
		v.AppendUnique(1,2,3,4)
		v2.AppendUnique(2,4)
		setDifferenceHelper(v,v,v2,[]int{1,3},factory,t)
	}
}

func SetInterfaceIsSuperset(
	factory func() dynamicContainers.Set[int],
	t *testing.T,
){
	v:=factory()
	v2:=factory()
	test.True(v.IsSuperset(v2),t)
	v.AppendUnique(1)
	test.True(v.IsSuperset(v2),t)
	test.False(v2.IsSuperset(v),t)
	v2.AppendUnique(1)
	test.True(v.IsSuperset(v2),t)
	test.True(v2.IsSuperset(v),t)
	v2.AppendUnique(2)
	test.False(v.IsSuperset(v2),t)
	test.True(v2.IsSuperset(v),t)
	v2.AppendUnique(3)
	test.False(v.IsSuperset(v2),t)
	test.True(v2.IsSuperset(v),t)
	v.AppendUnique(2,3)
	test.True(v.IsSuperset(v2),t)
	test.True(v2.IsSuperset(v),t)
	v.AppendUnique(4,5)
	test.True(v.IsSuperset(v2),t)
	test.False(v2.IsSuperset(v),t)
}

func SetInterfaceIsSubset(
	factory func() dynamicContainers.Set[int],
	t *testing.T,
){
	v:=factory()
	v2:=factory()
	test.True(v.IsSubset(v2),t)
	v.AppendUnique(1)
	test.False(v.IsSubset(v2),t)
	test.True(v2.IsSubset(v),t)
	v2.AppendUnique(1)
	test.True(v.IsSubset(v2),t)
	test.True(v2.IsSubset(v),t)
	v2.AppendUnique(2)
	test.True(v.IsSubset(v2),t)
	test.False(v2.IsSubset(v),t)
	v2.AppendUnique(3)
	test.True(v.IsSubset(v2),t)
	test.False(v2.IsSubset(v),t)
	v.AppendUnique(2,3)
	test.True(v.IsSubset(v2),t)
	test.True(v2.IsSubset(v),t)
	v.AppendUnique(4,5)
	test.False(v.IsSubset(v2),t)
	test.True(v2.IsSubset(v),t)
}
