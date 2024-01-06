package testingstupidideasthatmybraindecidestopoopout

import "fmt"

type IteratorFeedback int;
const (
    Continue IteratorFeedback=iota
    Break
    Iterate
);
type iter[T any] func(f IteratorFeedback)(T,error,bool);
type Iter[T any, U any] struct {
	iter[T]
}

func Mapable[T any, U any](i iter[T]) Iter[T,U] {
	return Iter[T,U]{i}
}


func SliceElems[T any](s []T) iter[T] {
    i:=-1;
    return func(f IteratorFeedback) (T,error,bool) {
        var rv T;
        i++;
        if i<len(s) && f!=Break {
            return s[i],nil,true;
        }
        return rv,nil,false;
    }
}

func (i *Iter[T,U])Map(op func(v T) U) iter[U] {
	return func(f IteratorFeedback) (U, error, bool) {
        if f==Break {
            var tmp U;
			return tmp,nil,false
        }
		v,err,cont:=(*(*i).iter)(f)
        tmp:=op(v);
        return tmp,err,cont
	}
}

func (i *iter[T])Collect() []T {
	l:=make([]T,0)
	for v,err,cont:=(*i)(Continue); err==nil && cont; v,err,cont=(*i)(Continue) {
		l = append(l, v)
	}
	return l
}

func main() {
	s:=[]int{0,1,2,3,4,5}
	s2:=Iter[int,string](SliceElems[int](
		s,
	)).Map[string,[]string](
		func(v int) string { return fmt.Sprintf("%d",v) },
	).Map[[]string,[]string]()
}

type RealIter[T any, U any] struct {
	iter[T]
}
