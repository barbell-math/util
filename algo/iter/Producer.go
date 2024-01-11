package iter

import (
    "bufio"
    "os"

    staticType "github.com/barbell-math/util/dataStruct/types/static"
    customerr "github.com/barbell-math/util/err"
)

func NoElem[T any]() Iter[T] {
    return func(f IteratorFeedback) (T,error,bool) {
        var tmp T;
        return tmp,nil,false;
    }
}

func ValElem[T any](val T, err error, repeat int) Iter[T] {
    cntr:=0;
    return func(f IteratorFeedback) (T,error,bool) {
        var rv T;
        if cntr<repeat && f!=Break {
            cntr++;
            return val,err,true;
        }
        return rv,nil,false;    
    }
}

func SliceElems[T any](s []T) Iter[T] {
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

func StrElems(s string) Iter[byte] {
    i:=-1;
    return func(f IteratorFeedback) (byte,error,bool) {
        i++;
        if i<len(s) && f!=Break {
            return s[i],nil,true;
        }
        return ' ',nil,false;
    }
}

func SequentialElems[T any](_len int, get func(i int) (T,error)) Iter[T] {
    i:=-1
    return func(f IteratorFeedback) (T, error, bool) {
        i++;
        if i<_len && f!=Break {
            v,err:=get(i)
            return v,err,(err==nil);
        }
        var tmp T
        return tmp,nil,false;
    }
}

// TODO -test
func SetupTeardownSequentialElems[T any](
    _len int,
    get func(i int) (T,error),
    setup func() error,
    teardown func() error,
) Iter[T] {
    i:=-1
    return func(f IteratorFeedback) (T, error, bool) {
        i++;
        if i>=_len || f==Break {
            var err error
            if i>0 {
                err=teardown()
            }
            var tmp T
            return tmp,err,false;
        }
        if i==0 {
            if err:=setup(); err!=nil {
                var tmp T
                return tmp,err,false
            }
        }
        v,err:=get(i)
        return v,err,(err==nil);
    }
}

func MapElems[K comparable, V any](
    m map[K]V, 
    v staticType.Pair[K,V],
) Iter[staticType.Pair[K,V]] {
    c:=make(chan K)
    go func(){
        for k,_:=range(m) {
            c <- k
        }
    }()
    i:=-1;
    return func(f IteratorFeedback) (staticType.Pair[K,V],error,bool) {
        i++;
        if i<len(m) && f!=Break {
            v.SetA(<-c)
            v.SetB(m[v.GetA()])
            return v,nil,true;
        }
        close(c)
        return nil,nil,false
    }
}

func MapKeys[K comparable, V any](m map[K]V) Iter[K] {
    c:=make(chan K)
    go func(){
        for k,_:=range(m) {
            c <- k
        }
    }()
    i:=-1;
    return func(f IteratorFeedback) (K,error,bool) {
        i++;
        if i<len(m) && f!=Break {
            return (<-c),nil,true;
        }
        close(c)
        var tmp K
        return tmp,nil,false
    }
}

func MapVals[K comparable, V any](m map[K]V) Iter[V] {
    c:=make(chan V)
    go func(){
        for _,v:=range(m) {
            c <- v
        }
    }()
    i:=-1;
    return func(f IteratorFeedback) (V,error,bool) {
        i++;
        if i<len(m) && f!=Break {
            return (<-c),nil,true;
        }
        close(c)
        var tmp V
        return tmp,nil,false
    }
}

func ChanElems[T any](c <-chan T) Iter[T] {
    return func(f IteratorFeedback) (T,error,bool) {
        if f!=Break {
            next,ok:=<-c;
            return next,nil,ok;
        }
        var rv T;
        return rv,nil,false;
    }
}

func FileLines(path string) Iter[string] {
    var scanner *bufio.Scanner;
    file,err:=os.Open(path);
    if err==nil {
        scanner=bufio.NewScanner(file);
        scanner.Split(bufio.ScanLines);
    }
    return func(f IteratorFeedback) (string,error,bool) {
        if f==Break || err!=nil || !scanner.Scan() {
            file.Close();
            return "",err,false;
        }
        return scanner.Text(),nil,true;
    }
}

func Zip[T any, U any](
    i1 Iter[T], 
    i2 Iter[U], 
    factory func() staticType.Pair[T,U],
) Iter[staticType.Pair[T,U]] {
    return func(f IteratorFeedback) (staticType.Pair[T, U], error, bool) {
        if f==Break {
            return nil,nil,false
        }
        iVal1,err1,cont1:=i1(f)
        iVal2,err2,cont2:=i2(f)
        p:=factory()
        p.SetA(iVal1)
        p.SetB(iVal2)
        if err1!=nil || err2!=nil {
            return nil,customerr.AppendError(err1,err2), false
        }
        return p, nil, (cont1 && cont2)
    }
}

func Join[T any, U any](
    i1 Iter[T],
    i2 Iter[U],
    factory func() staticType.Variant[T,U],
    decider func(left T, right U) bool,
) Iter[staticType.Variant[T,U]] {
    var i1Val T;
    var i2Val U;
    var err1, err2 error;
    cont1, cont2:=true, true;
    getI1Val, getI2Val:=true, true;
    return func(f IteratorFeedback) (staticType.Variant[T,U], error, bool) {
        if f==Break {
            return nil, customerr.AppendError(i1.Stop(),i2.Stop()), false;
        }
        if getI1Val && cont1 && err1==nil {
            i1Val,err1,cont1=i1(f);
        }
        if getI2Val && cont2 && err2==nil {
            i2Val,err2,cont2=i2(f);
        }
        if err1!=nil || err2!=nil {
            return nil,customerr.AppendError(err1,err2),false;
        }
        if cont1 && cont2 {
            d:=decider(i1Val,i2Val);
            getI1Val=d;
            getI2Val=!d;
            if d {
                return factory().SetValA(i1Val),err1,cont1 && cont2;
            } else {
                return factory().SetValB(i2Val),err2,cont1 && cont2;
            }
        } else if cont1 && !cont2 {
            getI1Val=true;
            getI2Val=false;
            return factory().SetValA(i1Val),err1,cont1;
        } else { // !cont1 && cont2
            getI1Val=false;
            getI2Val=true;
            return factory().SetValB(i2Val),err2,cont2;
        }
    }
}

func JoinSame[T any](
    i1 Iter[T],
    i2 Iter[T],
    factory func() staticType.Variant[T,T],
    decider func(left T, right T) bool,
) Iter[T] {
    var tmp T;
    realJoiner:=Join(i1,i2,factory,decider);
    return func(f IteratorFeedback) (T, error, bool) {
        val,err,cont:=realJoiner(f);
        if cont && err==nil {
            if val.HasA() {
                return val.ValA(),err,cont;
            } else if val.HasB() {
                return val.ValB(),err,cont;
            }
        }
        return tmp,err,cont;
    }
}

func Recurse[T any](
    root Iter[T],
    shouldRecurse func(v T) bool,
    iterValToIter func (v T) Iter[T],
) Iter[T] {
    levels:=make([]Iter[T],1)
    levels[0]=root
    levelsBreakOp:=func() (T,error,bool) {
        var err error
        for _,v:=range(levels) {
            _,err2,_:=v(Break)
            err=customerr.AppendError(err,err2)
        }
        var tmp T
        return tmp, err, false
    }
    return func(f IteratorFeedback) (T, error, bool) {
        if f==Break {
            return levelsBreakOp()
        }
        for len(levels)>0 {
            v,err,cont:=levels[len(levels)-1](Continue)
            if !cont {
                levels=levels[0:len(levels)-1]
                continue
            }
            if err!=nil {
                var tmp T
                return  tmp,err,false
            }
            if shouldRecurse(v) {
                levels = append(levels, iterValToIter(v))
                return v,nil,true
            } else {
                return v,nil,true
            }
        }
        var tmp T
        return tmp, nil, false
    }
}
