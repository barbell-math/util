package widgets

import (
	"testing"

	"github.com/barbell-math/util/algo/hash"
	"github.com/barbell-math/util/test"
)

type customWidget struct {
    a int
    b float32
}

func (c *customWidget)Eq(l *customWidget, r *customWidget) bool {
    return l.a==r.a
}
func (c *customWidget)Lt(l *customWidget, r *customWidget) bool {
    return l.a<r.a
}
func (c *customWidget)Hash(v *customWidget) hash.Hash {
    return hash.Hash(v.a)
}
func (c *customWidget)Zero(v *customWidget) {
    *v=customWidget{}
}

func testCustomWidgetHelper(
    v customWidget,
    v2 customWidget,
    t *testing.T,
) {
    w:=NewWidget[customWidget,*customWidget]()
    test.True(w.Eq(&v,&v),t)
    test.False(w.Eq(&v,&v2),t)
    test.False(w.Lt(&v,&v),t)
    test.False(w.Lt(&v2,&v2),t)
    test.False(w.Lt(&v,&v2),t)
    test.True(w.Lt(&v2,&v),t)
    test.Eq(hash.Hash(10),v.Hash(&v),t)
    test.Eq(hash.Hash(9),v.Hash(&v2),t)
}
func TestCustomWidget(t *testing.T){
    v:=customWidget{a: 10, b: 20}
    v2:=customWidget{a: 9, b: 20}
    testCustomWidgetHelper(v,v2,t)
}

func TestCustomWidgetPntr(t *testing.T) {
    vImpl:=&customWidget{a: 10, b: 20}
    v:=NewWidget[*customWidget,Pntr[customWidget,*customWidget]]()
    v2Impl:=&customWidget{a: 9, b: 20}
    test.True(v.Eq(&vImpl,&vImpl),t)
    test.False(v.Eq(&vImpl,&v2Impl),t)
    test.False(v.Lt(&vImpl,&vImpl),t)
    test.False(v.Lt(&v2Impl,&v2Impl),t)
    test.False(v.Lt(&vImpl,&v2Impl),t)
    test.True(v.Lt(&v2Impl,&vImpl),t)
    test.Eq(hash.Hash(10),v.Hash(&vImpl),t)
    test.Eq(hash.Hash(9),v.Hash(&v2Impl),t)
}
