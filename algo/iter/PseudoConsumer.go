package iter

import (
	"fmt"
	"io"
)

// This function is a Consumer.
//
// Consume will consume its parent iterators entirely, only stopping early if it
// hits an error. The values from the parent iterator are not saved or used in
// any way.
func (i Iter[T])Consume() error {
    return i.ForEach(func(index int, val T) (IteratorFeedback, error) { 
        return Continue,nil;
    });
}

// This function is a Consumer.
//
// ToChan will take all the values from it's parent iterator and will push them
// to the specified channel. Errors are not propogated to the channel. Iteration
// will stop if an error occurs. This function will not close the channel that 
// is supplied to it.
func (i Iter[T])ToChan(c chan T) {
    i.ForEach(func(i int, val T) (IteratorFeedback,error) {
        c <- val;
        return Continue,nil;
    });
}

// This function is a Consumer.
//
// ToFile will take all the values from it's parent iterator and will write them
// to the specified writer. All values are written to the writer using the standard
// fmt.Sprintf %v formatting directive. If you want a format other than that then
// map the values to a the formatted string before passing values to this iterator.
// 
// If addNewLine is true then newlines will be appended to every value after it
// is formatted.
func (i Iter[T])ToWriter(src io.Writer, addNewLine bool) error {
    return i.ForEach(func(index int, val T) (IteratorFeedback, error) {
        src.Write([]byte(fmt.Sprintf("%v",val)))
        if addNewLine {
            src.Write([]byte("\n"))
        }
        return Continue,nil;
    });
}

// This function is a Consumer.
//
// Count will consumer it's parent iterator, counting the number of values that
// it recieves. The values from the parent iterator are not saved or used in
// any way. Iteration will stop if an error occurs.
func (i Iter[T])Count() (int,error) {
    rv:=0;
    err:=i.ForEach(func(i int, val T) (IteratorFeedback,error) {
        rv++;
        return Continue,nil;
    });
    return rv,err;
}

// This function is a Consumer.
//
// Collect will collect all of it's parent iterators values into a slice and
// return it. Iteration will stop if an error occurs. If an error occurs then
// the slice will contain all the values that were recieved up until that point.
func (i Iter[T])Collect() ([]T,error) {
    rv:=make([]T,0);
    err:=i.ForEach(func(index int, val T) (IteratorFeedback, error) {
        rv=append(rv,val);
        return Continue,nil;
    });
    return rv,err;
}

// This function is a Consumer.
//
// AppendTo will append all of it's parent iterators values into a slice.
// Iteration will stop if an error occurs. If an error occurs then
// the slice will contain all the values that were already present in it plus the
// values that were recieved up until the error occured.
func (i Iter[T])AppendTo(orig *[]T) (int,error) {
    j:=0;
    err:=i.ForEach(func(index int, val T) (IteratorFeedback,error) {
        *orig=append(*orig,val);
        j++;
        return Continue,nil;
    });
    return j,err;
}


// TODO -reimplement
// // This function is a Consumer.
// //
// // CollectInto will collect all of it's parent iterators values into the 
// // supplied container. The action that is used to place the value in the container
// // is specified by the collectOp parameter, which will be passed the container
// // and each iteration value to add to the container. The container is locked
// // for the entire duration of iteration, meaning reads will not be possible
// // until iteration is done. Note that there are no constraints on what the 
// // container contains before iteration starts, meaning CollectInto can also be
// // used to append values to an existing collection. Iteration will stop if an 
// // error occurs. If an error occurs then the slice will contain all the values 
// // that were recieved up until that point.
// func (i Iter[T])CollectInto(
//     container interface { types.Write[int,T]; types.Syncable },
//     collectOp func(
//         container interface { types.Write[int,T]; types.Syncable }, 
//         val T,
//     ) error,
// ) (int,error) {
//     j:=0
//     err:=i.SetupTeardown(
//         func() error { container.Lock(); return nil },
//         func() error { container.Unlock(); return nil },
//     ).ForEach(func(index int, val T) (IteratorFeedback, error) {
//         j++
//         if err:=collectOp(container,val); err!=nil {
//             return Break,err
//         }
//         return Continue,nil 
//     })
//     return j,err
// }

// This function is a Consumer.
//
// All will apply the supplied operation function (op) to each value from it's
// parent iterator and will only return true if the operation function returns
// true for every value, otherwise it will return false. If an error occurs
// iteration will stop. If an error occurs, the state of the boolean value that 
// All returns will reflect the result of applying the operation function to 
// every value from the parent iterator before the error occured, excluding the
// value returned with the error.
func (i Iter[T])All(op func(val T) (bool,error)) (bool,error) {
    rv:=true;
    err:=i.ForEach(func(index int, val T) (IteratorFeedback,error) {
        res,err:=op(val);
        if !res {
            rv=false;
            return Break,err;
        }
        return Continue,err;
    });
    return rv,err;
}

// This function is a Consumer.
//
// Any will apply the supplied operation function (op) to each value from it's
// parent iterator and will only return true if the operation function returns
// true for any value, otherwise it will return false. If an error occurs
// iteration will stop. If an error occurs, the state of the boolean value that 
// Any returns will reflect the result of applying the operation function to 
// every value from the parent iterator before the error occured, excluding the
// value returned with the error.
func (i Iter[T])Any(op func(val T) (bool,error)) (bool,error) {
    rv:=false;
    err:=i.ForEach(func(index int, val T) (IteratorFeedback, error) {
        res,err:=op(val);
        if res {
            rv=true;
            return Break,err;
        }
        return Continue,err;
    });
    return rv,err;
}

// This function is a Consumer.
//
// Find will search for a value in it's parent iterators values using the 
// supplied operation function (op) as an equality comparison. Find returns three
// values:
// 
// 1. The value that was found 
// 2. An error status that represents errors raised during iteration or from the
// equality comparison function.
// 3. A boolean value that indicates if the value was found or not.
//
// If an error occurs iteration stops.
func (i Iter[T])Find(op func(val T) (bool,error)) (T,error,bool) {
    var rv T;
    iterState:=Continue;
    err:=i.ForEach(func(i int, val T) (IteratorFeedback,error){
        found,err:=op(val);
        if found {
            rv=val;
            iterState=Break;
        }
        return iterState,err;
    });
    if iterState==Break {
        return rv,err,true;
    }
    return rv,err,false;
}

// This function is a Consumer.
//
// Index will search for a value in it's parent iterators values using the 
// supplied operation function (op) as an equality comparison. The index of the
// value will be returned, or -1 if the value is not found. If an error occurs 
// iteration will stop.
func (i Iter[T])Index(op func(val T) (bool,error)) (int,error) {
    rv:=-1;
    err:=i.ForEach(func(i int, val T) (IteratorFeedback,error) {
        found,err:=op(val);
        if found {
            rv=i;
            return Break,err;
        }
        return Continue,err;
    });
    return rv,err;
}

// This function is a Consumer.
//
// Nth will return the value at the nth index from it's parent consumer. Nth
// returns three values:
// 
// 1. The value that was found 
// 2. An error status that represents errors raised during iteration
// 3. A boolean value that indicates if the nth index was reached or not
//
// If an error occurs iteration will stop.
func (i Iter[T])Nth(idx int) (T,error,bool) {
    var valRv T;
    found:=false;
    err:=i.ForEach(func(i int, val T) (IteratorFeedback,error) {
        if i==idx {
            valRv=val;
            found=true;
            return Break,nil;
        }
        return Continue,nil;
    });
    return valRv,err,found;
}

// This function is a Consumer.
//
// Reduce will take all of the values from it's parent iteratior and will
// combine them using the logic from the operation function (op), returning
// the combined value. The start value is specified as an argument. If an error
// occurs then iteration will stop and the current accumulated value will be
// returned without applying the operation function to the value that returned
// an error.
func (i Iter[T])Reduce(start T, op func(accum *T, iter T) error) (T,error) {
    accum:=start;
    err:=i.ForEach(func(index int, val T) (IteratorFeedback, error) {
        return Continue,op(&accum,val);
    });
    return accum,err;
}
