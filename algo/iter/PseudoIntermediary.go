package iter;

import (
    staticType "github.com/barbell-math/util/dataStruct/types/static"
)

func (i Iter[T])Take(num int) Iter[T] {
    cntr:=0;
    return i.Next(
    func(index int, val T, status IteratorFeedback) (IteratorFeedback, T, error) {
        if status!=Break && cntr<num {
            cntr++;
            return Continue,val,nil;
        }
        return Break,val,nil;
    });
}

func (i Iter[T])TakeWhile(op func(val T) bool) Iter[T] {
    return i.Next(
    func(index int, val T, status IteratorFeedback) (IteratorFeedback, T, error) {
        if status!=Break && op(val) {
            return Continue,val,nil;
        }
        return Break,val,nil;
    });
}

func (i Iter[T])Skip(num int) Iter[T] {
    return i.Filter(FilterToIndex[T](num));
}

func Map[T any, U any](
    i Iter[T],
    op func(index int, val T) (U,error),
) Iter[U] {
    return Next(i,
    func(index int, val T, status IteratorFeedback) (IteratorFeedback, U, error) {
        if status==Break {
            var tmp U;
            return Break,tmp,nil;
        }
        tmp,err:=op(index,val);
        return Continue,tmp,err;
    });
}
func (i Iter[T])Map(op func(index int, val T) (T,error)) Iter[T] {
    return Map(i,op);
}

func FilterToIndex[T any](num int) func(index int, val T) bool {
    return func(index int, val T) bool {
        return !(index<num);
    }
}
func (i Iter[T])Filter(op func(index int, val T) bool) Iter[T] {
    return i.Next(
    func(index int, val T, status IteratorFeedback) (IteratorFeedback, T, error) {
        if status!=Break && op(index,val) {
            return Continue,val,nil;
        }
        return Iterate,val,nil;
    });
}

//Window cannot be tested here because it would cause a circular import with the 
//dataStruct module. Testing would require importing a specific implementation
//of a queue. Using the types interface definition is the only thing preventing
//a circular import now.
func Window[T any](i Iter[T],
    q interface{ staticType.Queue[T]; staticType.Vector[T] },
    allowPartials bool,
) Iter[staticType.Vector[T]] {
    return Next(i,
    func(
        index int, val T, status IteratorFeedback,
    ) (IteratorFeedback, staticType.Vector[T], error) {
        if status==Break {
            return Break,q,nil;
        }
        q.ForcePushBack(val);
        if !allowPartials && q.Length()!=q.Capacity() {
            return Iterate,q,nil;
        }
        return Continue,q,nil;
    });
}
