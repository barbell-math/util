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
// [containerTypes.RWSyncable] interface.
func SetInterfaceSyncableInterface[V any](
	factory func() dynamicContainers.Set[V],
	t *testing.T,
) {
	var container containerTypes.RWSyncable = factory()
	_ = container
}

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
	var container containerTypes.ReadOps[uint64, V] = factory()
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
		err:=container.AppendUnique(i)
		test.ContainsError(containerTypes.Duplicate,err,
			"AppendUnique allowed a duplicate value to be appended.",t,
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