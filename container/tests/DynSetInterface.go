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
		"Code did not panic when casting a dynamic set to a static set.", t,
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
		test.BasicTest(true,container.ContainsPntr(&i),
			"ContainsPntr returned a false negative",t,
		)
	}
	v:=5
	test.BasicTest(false,container.ContainsPntr(&v),
		"ContainsPntr returned a false positive.",t,
	)
	v=-1
	test.BasicTest(false,container.ContainsPntr(&v),
		"ContainsPntr returned a false positive.",t,
	)
	container.Pop(0,1)
	v=0
	test.BasicTest(false,container.ContainsPntr(&v),
		"ContainsPntr returned a false positive.",t,
	)
	for i:=1; i<5; i++ {
		test.BasicTest(true,container.ContainsPntr(&i),
			"ContainsPntr returned a false negative",t,
		)
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
		test.BasicTest(true,container.Contains(i),
			"Contains returned a false negative",t,
		)
	}
	test.BasicTest(false,container.Contains(5),
		"Contains returned a false positive.",t,
	)
	test.BasicTest(false,container.Contains(-1),
		"Contains returned a false positive.",t,
	)
	container.Pop(0,1)
	test.BasicTest(false,container.Contains(0),
		"Contains returned a false positive.",t,
	)
	for i:=1; i<5; i++ {
		test.BasicTest(true,container.Contains(i),
			"Contains returned a false negative",t,
		)
	}
}

// Tests the Clear method functionality of a dynamic set.
func SetInterfaceClear(
	factory func() dynamicContainers.Set[int],
	t *testing.T,
){
	container:=factory()
	test.BasicTest(0,container.Length(),
		"Container was initilized with values in it.",t,
	)
	for i:=0; i<5; i++ {
		container.AppendUnique(i)
	}
	test.BasicTest(5,container.Length(),
		"Container did not save all values.",t,
	)
	container.Clear()
	test.BasicTest(0,container.Length(),
		"Container did not remove all values and set length to 0.",t,
	)
}

// Tests the AppendUnique method functionality of a dynamic set.
func SetInterfaceAppendUnique(
	factory func() dynamicContainers.Set[int],
	t *testing.T,
){
	container:=factory()
	test.BasicTest(0,container.Length(),
		"Container was initilized with values in it.",t,
	)
	for i:=0; i<5; i++ {
		err:=container.AppendUnique(i)
		test.BasicTest(nil,err,
			"AppendUnique returned an error when it shouldn't have.",t,
		)
	}
	for i:=0; i<5; i++ {
		test.BasicTest(true,container.Contains(i),
			"Appending a value did not place it in the container.",t,
		)
	}
	for i:=0; i<5; i++ {
		container.AppendUnique(i)
		test.BasicTest(5,container.Length(),
			"Container had non-unique values added to it.",t,
		)
	}
	container=factory()
	test.BasicTest(0,container.Length(),
		"Container was initilized with values in it.",t,
	)
	for i:=0; i<6; i+=2 {
		container.AppendUnique(i,i+1)
	}
	for i:=0; i<6; i++ {
		test.BasicTest(true,container.Contains(i),
			"Appending a value did not place it in the container.",t,
		)
	}
}

// Tests the Pop method functionality of a dynamic set.
func SetInterfacePop(
	factory func() dynamicContainers.Set[int],
	t *testing.T,
){
	container:=factory()
	test.BasicTest(0,container.Length(),
		"Container was initilized with values in it.",t,
	)
	for i:=0; i<5; i++ {
		container.AppendUnique(i)
	}
	for i:=0; i<5; i++ {
		test.BasicTest(true,container.Contains(i),
			"The container contain the value originally.",t,
		)
		container.Pop(i,1)
		test.BasicTest(false,container.Contains(i),
			"The container did not remove the value.",t,
		)
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
	test.BasicTest(true, v.UnorderedEq(v2), 
		"UnorderedEq returned a false negative.", t,
	)
	test.BasicTest(true, v2.UnorderedEq(v), 
		"UnorderedEq returned a false negative.", t,
	)
	v.Pop(3,1)
	test.BasicTest(false, v.UnorderedEq(v2), 
		"UnorderedEq returned a false positive.", t,
	)
	test.BasicTest(false, v2.UnorderedEq(v), 
		"UnorderedEq returned a false positive.", t,
	)
	v.AppendUnique(3)
	v2 = factory()
	v2.AppendUnique(3, 1, 2)
	test.BasicTest(true, v.UnorderedEq(v2), 
		"UnorderedEq returned a false negative.", t,
	)
	test.BasicTest(true, v2.UnorderedEq(v), 
		"UnorderedEq returned a false negative.", t,
	)
	v.Pop(3,1)
	test.BasicTest(false, v.UnorderedEq(v2), 
		"UnorderedEq returned a false positive.", t,
	)
	test.BasicTest(false, v2.UnorderedEq(v), 
		"UnorderedEq returned a false positive.", t,
	)
	v.AppendUnique(3)
	v2 = factory()
	v2.AppendUnique(2, 3, 1)
	test.BasicTest(true, v.UnorderedEq(v2), 
		"UnorderedEq returned a false negative.", t,
	)
	test.BasicTest(true, v2.UnorderedEq(v), 
		"UnorderedEq returned a false negative.", t,
	)
	v.Pop(3,1)
	test.BasicTest(false, v.UnorderedEq(v2), 
		"UnorderedEq returned a false positive.", t,
	)
	test.BasicTest(false, v2.UnorderedEq(v), 
		"UnorderedEq returned a false positive.", t,
	)
	v = factory()
	v.AppendUnique(0)
	v2 = factory()
	v2.AppendUnique(0)
	test.BasicTest(true, v.UnorderedEq(v2), 
		"UnorderedEq returned a false negative.", t,
	)
	test.BasicTest(true, v2.UnorderedEq(v), 
		"UnorderedEq returned a false negative.", t,
	)
	v.Pop(0,1)
	test.BasicTest(false, v.UnorderedEq(v2), 
		"UnorderedEq returned a false positive.", t,
	)
	test.BasicTest(false, v2.UnorderedEq(v), 
		"UnorderedEq returned a false positive.", t,
	)
	v = factory()
	v2 = factory()
	test.BasicTest(true, v.UnorderedEq(v2), 
		"UnorderedEq returned a false negative.", t,
	)
	test.BasicTest(true, v2.UnorderedEq(v), 
		"UnorderedEq returned a false negative.", t,
	)
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
		test.BasicTest(len(exp),c.Length(),
			"Intersection produced a set of the wrong length.",t,
		)
		for _,v:=range(exp) {
			test.BasicTest(true,c.Contains(v),
				"Intersection did not contain the correct values.",t,
			)
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
		test.BasicTest(len(exp),c.Length(),
			"Union produced a set of the wrong length.",t,
		)
		for _,v:=range(exp) {
			test.BasicTest(true,c.Contains(v),
				"Union did not contain the correct values.",t,
			)
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
	test.BasicTest(len(exp),res.Length(),
		"Difference produced a set of the wrong length.",t,
	)
	for _,v:=range(exp) {
		test.BasicTest(true,res.Contains(v),
			"Difference did not contain the correct values.",t,
		)
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
	test.BasicTest(true,v.IsSuperset(v2),
		"Superset returned a false negative.",t,
	)
	v.AppendUnique(1)
	test.BasicTest(true,v.IsSuperset(v2),
		"Superset returned a false negative.",t,
	)
	test.BasicTest(false,v2.IsSuperset(v),
		"Superset returned a false positive.",t,
	)
	v2.AppendUnique(1)
	test.BasicTest(true,v.IsSuperset(v2),
		"Superset returned a false negative.",t,
	)
	test.BasicTest(true,v2.IsSuperset(v),
		"Superset returned a false negative.",t,
	)
	v2.AppendUnique(2)
	test.BasicTest(false,v.IsSuperset(v2),
		"Superset returned a false positive.",t,
	)
	test.BasicTest(true,v2.IsSuperset(v),
		"Superset returned a false negative.",t,
	)
	v2.AppendUnique(3)
	test.BasicTest(false,v.IsSuperset(v2),
		"Superset returned a false positive.",t,
	)
	test.BasicTest(true,v2.IsSuperset(v),
		"Superset returned a false negative.",t,
	)
	v.AppendUnique(2,3)
	test.BasicTest(true,v.IsSuperset(v2),
		"Superset returned a false negative.",t,
	)
	test.BasicTest(true,v2.IsSuperset(v),
		"Superset returned a false negative.",t,
	)
	v.AppendUnique(4,5)
	test.BasicTest(true,v.IsSuperset(v2),
		"Superset returned a false negative.",t,
	)
	test.BasicTest(false,v2.IsSuperset(v),
		"Superset returned a false positive.",t,
	)
}

func SetInterfaceIsSubset(
	factory func() dynamicContainers.Set[int],
	t *testing.T,
){
	v:=factory()
	v2:=factory()
	test.BasicTest(true,v.IsSubset(v2),
		"Subset returned a false negative.",t,
	)
	v.AppendUnique(1)
	test.BasicTest(false,v.IsSubset(v2),
		"Subset returned a false positive.",t,
	)
	test.BasicTest(true,v2.IsSubset(v),
		"Subset returned a false negative.",t,
	)
	v2.AppendUnique(1)
	test.BasicTest(true,v.IsSubset(v2),
		"Subset returned a false negative.",t,
	)
	test.BasicTest(true,v2.IsSubset(v),
		"Subset returned a false negative.",t,
	)
	v2.AppendUnique(2)
	test.BasicTest(true,v.IsSubset(v2),
		"Subset returned a false negative.",t,
	)
	test.BasicTest(false,v2.IsSubset(v),
		"Subset returned a false positive.",t,
	)
	v2.AppendUnique(3)
	test.BasicTest(true,v.IsSubset(v2),
		"Subset returned a false negative.",t,
	)
	test.BasicTest(false,v2.IsSubset(v),
		"Subset returned a false positive.",t,
	)
	v.AppendUnique(2,3)
	test.BasicTest(true,v.IsSubset(v2),
		"Subset returned a false negative.",t,
	)
	test.BasicTest(true,v2.IsSubset(v),
		"Subset returned a false negative.",t,
	)
	v.AppendUnique(4,5)
	test.BasicTest(false,v.IsSubset(v2),
		"Subset returned a false positive.",t,
	)
	test.BasicTest(true,v2.IsSubset(v),
		"Subset returned a false negative.",t,
	)
}
