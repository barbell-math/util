package iter

import (
    "bufio"
    "os"

    staticType "github.com/barbell-math/util/dataStruct/types/static"
    customerr "github.com/barbell-math/util/err"
)

// NoElem provides an iterator that returns no elements. NoElem returns an empty
// iterator.
func NoElem[T any]() Iter[T] {
    return func(f IteratorFeedback) (T,error,bool) {
        var tmp T;
        return tmp,nil,false;
    }
}

// ValElem returns an iterator that produces the supplied value and error the 
// supplied number of times. The same value and error will be returned so any
// modifications made to the value and error will be visible on subsiquent
// iterations. Any error other than nil will cause iteration to stop after the
// first value due to how the intermediaries and consumers are implemented.
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

// SliceElems returns an iterator that iterates over the supplied slices 
// elements. No error will ever be returned by this producer. This producer is 
// not thread safe. If the underlying slice is modified while it is being 
// iterated over the behavior will be undefined and there are no order 
// guarantees if this function is passed to multiple threads. For a thread safe 
// implementation of SliceElems use the Vector.Elems method from the datastruct
// package.
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

// StrElems returns an iterator that iterates over the supplied strings 
// characters. No error will ever be returned by this producer.
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

// SequentialElems returns an iterator that iterates over a general container 
// using the get function in combination with the length function. Note that
// unlike the SliceElems and StrElems producers this producer can return an 
// error.
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

// MapElems returns an iterator that iterates over a maps key,value pairs. Do
// not confuse this with the Map intermediary function. This producer will never
// return an error. This producer is not thread safe. If the underlying map value
// it changed while being iterated over behavior is undefined.
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

// MapElems returns an iterator that iterates over a maps key values. Do not 
// confuse this with the Map intermediary function. This producer will never
// return an error. This producer is not thread safe. If the underlying map value
// it changed while being iterated over behavior is undefined.
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

// MapElems returns an iterator that iterates over a maps values. Do not confuse
// this with the Map intermediary function. This producer will never return an
// error. This producer is not thread safe. If the underlying map value it
// changed while being iterated over behavior is undefined.
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

// ChanElems returns an iterator that iterates over the elements in an unbuffered
// channel. Calling this function will block until the channel receives a value.
// This producer will never return an error.
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

// FileLines returns an iterator that iterates over the lines in a file. If an
// error occurs opening the file then no lines will be iterated over and the
// error will be returned upon the first iteration of the producer.
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

// Zip will take two iterators and return an iterator that iterates over pairs
// of values where each pair contains a value from each supplied iterator. The
// number of values returned by this iterator will equal the number of elements
// from the supplied iterator that produces the least number values. Excess 
// values will be ignored. Errors from the supplied iterators will be returned 
// from this iterator.
func Zip[T any, U any](
    i1 Iter[T], 
    i2 Iter[U], 
    factory func() staticType.Pair[T,U],
) Iter[staticType.Pair[T,U]] {
    return func(f IteratorFeedback) (staticType.Pair[T, U], error, bool) {
        if f==Break {
            return nil,customerr.AppendError(i1.Stop(),i2.Stop()),false
        }
        iVal1,err1,cont1:=i1(f)
        if err1!=nil {
            return nil,err1,false
        }
        iVal2,err2,cont2:=i2(f)
        if err2!=nil {
            return nil,err2,false
        }
        p:=factory()
        p.SetA(iVal1)
        p.SetB(iVal2)
        return p, nil, (cont1 && cont2)
    }
}

// Join takes two iterators and a decider function and returns an iterator that
// consumes both supplied iterators, returning a single value at a time based on
// the return value from the decider function. The number of values returned will
// equal the total number of values returned from both of the supplied iterators.
// Errors from the supplied iterators will be returned by this iterator.
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

// JoinSame takes two iterators and a decider function and returns an iterator 
// that consumes both supplied iterators, returning a single value at a time 
// based on the return value from the decider function. The number of values 
// returned will equal the total number of values returned from both of the 
// supplied iterators. Errors from the supplied iterators will be returned by 
// this iterator.
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

// Recurse will return an iterator that recursively returns values from the 
// supplied iterator. This iterator will enforce root-left-right traversal. This
// order is the only available order because once an iterator has produced a 
// value there is no way to "push" it back.
// 
// Recurse takes a root iterator where the recursion begins. This iterator can 
// return as many values as it needs, there is no limitation holding to only 
// produce one value. 
//
// The shouldRecurse function should return true if the current value needs to 
// be recursed upon. 
//
// The iterValToIter takes a value from an iterator and returns an iterator over
// that value. This is where the recursion happens.
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
            err=customerr.AppendError(err,v.Stop())
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
