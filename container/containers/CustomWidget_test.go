package containers

import (
	"testing"

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
func (c *customWidget)Hash(v *customWidget) uint64 {
    return uint64(v.a)
}

func TestCustomWidgetInVector(t *testing.T){
    v:=make(Vector[customWidget,*customWidget],0)
    v.Append(customWidget{a: 10, b: 20})
    v.Append(customWidget{a: 9, b: 20})
    test.BasicTest(true,v[0].Eq(&v[0],&v[0]),
        "The correct equals function was not called.",t,
    )
    test.BasicTest(false,v[0].Eq(&v[0],&v[1]),
        "The correct equals function was not called.",t,
    )
    test.BasicTest(false,v[0].Lt(&v[0],&v[0]),
        "The correct less than function was not called.",t,
    )
    test.BasicTest(false,v[0].Lt(&v[1],&v[1]),
        "The correct less than function was not called.",t,
    )
    test.BasicTest(false,v[0].Lt(&v[0],&v[1]),
        "The correct less than function was not called.",t,
    )
    test.BasicTest(true,v[0].Lt(&v[1],&v[0]),
        "The correct less than function was not called.",t,
    )
    test.BasicTest(uint64(10),v[0].Hash(&v[0]),
        "The correct hash function was not called.",t,
    )
    test.BasicTest(uint64(9),v[0].Hash(&v[1]),
        "The correct hash function was not called.",t,
    )
}
