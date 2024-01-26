package iter;

import (
    "github.com/barbell-math/util/customerr"
)

// ForEach is the most ubuquitous consumer, most other consumers can be expressed 
// using ForEach making them pseudo-consumers. By using this pattern all 
// pseudo-consumers are abstracted away from the complex looping logic.

// This funciton is a consumer.
//
// For each will take values from it's parent iterator and perform the supplied
// operation function on each value, until an error occurs. If an error occurs
// the operation function is not called and iteration stops.
func (i Iter[T])ForEach(
        op func(index int, val T) (IteratorFeedback,error)) error {
    j:=0;
    f:=Continue;
    var next T;
    var err error;
    var cont bool=true;
    var opErr error=nil;
    for cont && err==nil && opErr==nil && f==Continue {
        next,err,cont=i(Iterate);
        if err==nil && cont {
            f,opErr=op(j,next);
            j++;
        }
    }
    _,cleanUpErr,_:=i(Break);
    return customerr.AppendError(opErr,
        customerr.AppendError(err,cleanUpErr),
    );
}

//Why is stop not a pseudo consumer? It breaks the parent calling convention
// that ForEach uses. For each will always get a parent iterators value before
// the op function is consulted. Stop should just stop, and not call the parent
// iterators one last time meaning it has to be separate.
func (i Iter[T])Stop() error {
    _,cleanUpErr,_:=i(Break);
    return cleanUpErr;
}
