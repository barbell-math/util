package containers

import (
	"testing"

	"github.com/barbell-math/util/algo/widgets"
	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/test"
)

func TestWrapAroundIntAdd(t *testing.T) {
	v := wrapingIndex(0)
	v = v.Add(1, 5)
	test.Eq(wrapingIndex(1), v, t)
	v = v.Add(1, 5)
	test.Eq(wrapingIndex(2), v, t)
	v = v.Add(1, 5)
	test.Eq(wrapingIndex(3), v, t)
	v = v.Add(1, 5)
	test.Eq(wrapingIndex(4), v, t)
	v = v.Add(1, 5)
	test.Eq(wrapingIndex(0), v, t)
	v = v.Add(1, 5)
	test.Eq(wrapingIndex(1), v, t)
	v = v.Add(2, 5)
	test.Eq(wrapingIndex(3), v, t)
	v = v.Add(2, 5)
	test.Eq(wrapingIndex(0), v, t)
	v = v.Add(3, 5)
	test.Eq(wrapingIndex(3), v, t)
	v = v.Add(3, 5)
	test.Eq(wrapingIndex(1), v, t)
	v = v.Add(4, 5)
	test.Eq(wrapingIndex(0), v, t)
	v = v.Add(5, 5)
	test.Eq(wrapingIndex(0), v, t)
	v = v.Add(6, 5)
	test.Eq(wrapingIndex(1), v, t)
}

func TestWrapAroundIntSub(t *testing.T) {
	v := wrapingIndex(5)
	v = v.Sub(1, 5)
	test.Eq(wrapingIndex(4), v, t)
	v = v.Sub(1, 5)
	test.Eq(wrapingIndex(3), v, t)
	v = v.Sub(1, 5)
	test.Eq(wrapingIndex(2), v, t)
	v = v.Sub(1, 5)
	test.Eq(wrapingIndex(1), v, t)
	v = v.Sub(1, 5)
	test.Eq(wrapingIndex(0), v, t)
	v = v.Sub(1, 5)
	test.Eq(wrapingIndex(4), v, t)
	v = v.Sub(2, 5)
	test.Eq(wrapingIndex(2), v, t)
	v = v.Sub(2, 5)
	test.Eq(wrapingIndex(0), v, t)
	v = v.Sub(3, 5)
	test.Eq(wrapingIndex(2), v, t)
	v = v.Sub(3, 5)
	test.Eq(wrapingIndex(4), v, t)
	v = v.Sub(4, 5)
	test.Eq(wrapingIndex(0), v, t)
	v = v.Sub(4, 5)
	test.Eq(wrapingIndex(1), v, t)
	v = v.Sub(5, 5)
	test.Eq(wrapingIndex(1), v, t)
	v = v.Sub(6, 5)
	test.Eq(wrapingIndex(0), v, t)
}

func TestWrapAroundIntGetProperIndex(t *testing.T) {
	w := wrapingIndex(0)
	test.Eq(wrapingIndex(0), w.GetProperIndex(-5, 5), t)
	test.Eq(wrapingIndex(1), w.GetProperIndex(-4, 5), t)
	test.Eq(wrapingIndex(2), w.GetProperIndex(-3, 5), t)
	test.Eq(wrapingIndex(3), w.GetProperIndex(-2, 5), t)
	test.Eq(wrapingIndex(4), w.GetProperIndex(-1, 5), t)
	test.Eq(wrapingIndex(0), w.GetProperIndex(0, 5), t)
	test.Eq(wrapingIndex(1), w.GetProperIndex(1, 5), t)
	test.Eq(wrapingIndex(2), w.GetProperIndex(2, 5), t)
	test.Eq(wrapingIndex(3), w.GetProperIndex(3, 5), t)
	test.Eq(wrapingIndex(4), w.GetProperIndex(4, 5), t)
	test.Eq(wrapingIndex(0), w.GetProperIndex(5, 5), t)
	test.Eq(wrapingIndex(1), w.GetProperIndex(6, 5), t)
	test.Eq(wrapingIndex(2), w.GetProperIndex(7, 5), t)
	test.Eq(wrapingIndex(3), w.GetProperIndex(8, 5), t)
	test.Eq(wrapingIndex(4), w.GetProperIndex(9, 5), t)
}

//go:generate go run interfaceTest.go -type=CircularBuffer -category=static -interface=Vector -genericDecl=[int] -factory=generateCircularBuffer
//go:generate go run interfaceTest.go -type=SyncedCircularBuffer -category=static -interface=Vector -genericDecl=[int] -factory=generateSyncedCircularBuffer
//go:generate go run interfaceTest.go -type=CircularBuffer -category=static -interface=Set -genericDecl=[int] -factory=generateCircularBuffer
//go:generate go run interfaceTest.go -type=SyncedCircularBuffer -category=static -interface=Set -genericDecl=[int] -factory=generateSyncedCircularBuffer
//go:generate go run interfaceTest.go -type=CircularBuffer -category=static -interface=Deque -genericDecl=[int] -factory=generateCircularBuffer
//go:generate go run interfaceTest.go -type=SyncedCircularBuffer -category=static -interface=Deque -genericDecl=[int] -factory=generateSyncedCircularBuffer
//go:generate go run interfaceTest.go -type=CircularBuffer -category=static -interface=Stack -genericDecl=[int] -factory=generateCircularBuffer
//go:generate go run interfaceTest.go -type=SyncedCircularBuffer -category=static -interface=Stack -genericDecl=[int] -factory=generateSyncedCircularBuffer
//go:generate go run interfaceTest.go -type=CircularBuffer -category=static -interface=Queue -genericDecl=[int] -factory=generateCircularBuffer
//go:generate go run interfaceTest.go -type=SyncedCircularBuffer -category=static -interface=Queue -genericDecl=[int] -factory=generateSyncedCircularBuffer

func generateCircularBuffer(capacity int) CircularBuffer[int, widgets.BuiltinInt] {
	c, _ := NewCircularBuffer[int, widgets.BuiltinInt](capacity)
	return c
}

func generateSyncedCircularBuffer(capacity int) SyncedCircularBuffer[int, widgets.BuiltinInt] {
	c, _ := NewSyncedCircularBuffer[int, widgets.BuiltinInt](capacity)
	return c
}

func circularBufferDeleteFrontHelper(l int, startIdx int, t *testing.T) {
	tmp, err := NewCircularBuffer[int, widgets.BuiltinInt](5)
	tmp.start = wrapingIndex(startIdx)
	test.Nil(err, t)
	err = tmp.Delete(0)
	test.ContainsError(customerr.ValOutsideRange, err, t)
	for i := 0; i < 5; i++ {
		tmp.PushBack(i)
	}
	err = tmp.Delete(6)
	test.ContainsError(customerr.ValOutsideRange, err, t)
	for i := 0; i < 5; i++ {
		err := tmp.Delete(0)
		test.Nil(err, t)
		test.Eq(4-i, tmp.Length(), t)
		for j := 0; j < 5-i-1; j++ {
			v, err := tmp.Get(j)
			test.Nil(err, t)
			test.Eq(i+j+1, v, t)
		}
	}
	test.Eq(0, tmp.Length(), t)
}
func TestCircularBufferDeleteFront(t *testing.T) {
	for i := 0; i < 5; i++ {
		circularBufferDeleteFrontHelper(5, i, t)
	}
}

func circularBufferDeleteBackHelper(l int, startIdx int, t *testing.T) {
	tmp, err := NewCircularBuffer[int, widgets.BuiltinInt](5)
	test.Nil(err, t)
	tmp.start = wrapingIndex(startIdx)
	err = tmp.Delete(0)
	test.ContainsError(customerr.ValOutsideRange, err, t)
	for i := 0; i < 5; i++ {
		tmp.PushBack(i)
	}
	err = tmp.Delete(6)
	test.ContainsError(customerr.ValOutsideRange, err, t)
	for i := 0; i < 5; i++ {
		err := tmp.Delete(tmp.Length() - 1)
		test.Nil(err, t)
		test.Eq(4-i, tmp.Length(), t)
		for j := 0; j < 5-i-1; j++ {
			v, _ := tmp.Get(j)
			test.Eq(j, v, t)
		}
	}
	test.Eq(0, tmp.Length(), t)
}
func TestCircularBufferDeleteBack(t *testing.T) {
	for i := 0; i < 5; i++ {
		circularBufferDeleteBackHelper(5, i, t)
	}
}

func circularBufferDeleteHelper(idx int, l int, startIdx int, t *testing.T) {
	tmp, err := NewCircularBuffer[int, widgets.BuiltinInt](l)
	test.Nil(err, t)
	tmp.start = wrapingIndex(startIdx)
	for i := 0; i < l; i++ {
		tmp.PushBack(i)
	}
	err = tmp.Delete(idx)
	test.Nil(err, t)
	test.Eq(l-1, tmp.Length(), t)
	for i := 0; i < l-1; i++ {
		var exp int
		if i < idx {
			exp = i
		} else if i >= idx {
			exp = i + 1
		}
		v, err := tmp.Get(i)
		test.Nil(err, t)
		test.Eq(exp, v, t)
	}
}
func TestCircularBufferRandomDelete(t *testing.T) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			circularBufferDeleteHelper(i, 5, j, t)
		}
	}
}

func TestCircularBufferWidgetInterface(t *testing.T) {
	var widget widgets.WidgetInterface[CircularBuffer[string, widgets.BuiltinString]]
	v, _ := NewCircularBuffer[string, widgets.BuiltinString](0)
	widget = &v
	_ = widget
}

func TestCircularBufferEq(t *testing.T) {
	v1, _ := NewCircularBuffer[string, widgets.BuiltinString](5)
	v1.Append("one", "two", "three", "four")
	v2, _ := NewCircularBuffer[string, widgets.BuiltinString](5)
	v2.Append("one", "two", "three", "four")

	test.True(v1.Eq(&v1, &v2), t)
	test.True(v1.Eq(&v2, &v1), t)
	v1.Set(basic.Pair[int, string]{0, "not one"})
	test.False(v1.Eq(&v1, &v2), t)
	test.False(v1.Eq(&v2, &v1), t)
}

func TestCircularBufferLt(t *testing.T) {
	v1, _ := NewCircularBuffer[string, widgets.BuiltinString](5)
	v1.Append("a", "b", "c", "d")
	v2, _ := NewCircularBuffer[string, widgets.BuiltinString](5)
	v2.Append("a", "b", "c", "d")

	test.False(v1.Lt(&v1, &v2), t)
	test.False(v1.Lt(&v2, &v1), t)
	v1.Set(basic.Pair[int, string]{0, "A"})
	test.True(v1.Lt(&v1, &v2), t)
	test.False(v1.Lt(&v2, &v1), t)
	v1.Set(basic.Pair[int, string]{0, "a"})
	v1.Set(basic.Pair[int, string]{1, "B"})
	test.True(v1.Lt(&v1, &v2), t)
	test.False(v1.Lt(&v2, &v1), t)
	v1.Set(basic.Pair[int, string]{1, "b"})
	v1.Delete(3)
	test.True(v1.Lt(&v1, &v2), t)
	test.False(v1.Lt(&v2, &v1), t)
	v2.Clear()
	test.False(v1.Lt(&v1, &v2), t)
	test.True(v1.Lt(&v2, &v1), t)
}

func TestCircularBufferHash(t *testing.T) {
	v1, _ := NewCircularBuffer[string, widgets.BuiltinString](5)
	v1.Append("a", "b", "c", "d")
	v2, _ := NewCircularBuffer[string, widgets.BuiltinString](5)
	v2.Append("a", "b", "c", "d")

	test.Eq(v1.Hash(&v1), v2.Hash(&v2), t)
	v1.Set(basic.Pair[int, string]{0, "blah"})
	test.False(v1.Hash(&v1) == v2.Hash(&v2), t)
	h := v1.Hash(&v1)
	for i := 0; i < 100; i++ {
		test.Eq(h, v1.Hash(&v1), t)
	}
	v3, _ := NewCircularBuffer[int, widgets.BuiltinInt](5)
	v3.Append(500, 600, 700)
	v4, _ := NewCircularBuffer[int, widgets.BuiltinInt](5)
	v4.Append(700, 600, 500)
	test.False(v3.Hash(&v3) == v4.Hash(&v4), t)
}

func TestCircularBufferZero(t *testing.T) {
	v, _ := NewCircularBuffer[string, widgets.BuiltinString](5)
	v.Append("a", "b", "c", "d")
	v.Zero(&v)
	test.Eq(0, v.numElems, t)
	test.Eq(0, int(v.start), t)
	test.SlicesMatch[string]([]string{"", "", "", "", ""}, v.vals, t)
}
