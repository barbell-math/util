package iter

import (
	"github.com/barbell-math/util/customerr"
)

type forEachParallelResult[T any, U any] struct {
    val T;
    res U;
    err error;
};
func (f forEachParallelResult[T,U])Unpack() (T,U,error){
    return f.val,f.res,f.err;
}

func forEachWorker[T any, U any](
        jobs chan T,
        results chan forEachParallelResult[T,U],
        op func(val T) (U,error)){
    for val:=range(jobs) {
        tmp,err:=op(val);
        results <- forEachParallelResult[T, U]{
            val: val,
            res: tmp,
            err: err,
        };
    }
}

// A helper function that can be passed to [Parallel] as a result operation that 
// performs no action. Helpfull when you want to start a bunch of parallel jobs
// and do nothing with there results.
func NoOp[T any, U any](val T, res U, err error) { return; };

// This function is a consumer.
//
// ForEachParallel will consumer it's parent iterators values and start a new 
// go routine to process the result using the worker operation. As each worker 
// finishes the result operation will be called on the workers results. The
// result operation is syncronized, so this could be where you collect the results
// from the parallel workers. The number of go routines that are spawned will be 
// limited to numThreads, which must be >=1.
func Parallel[T any, U any](
    i Iter[T],
    workerOp func(val T) (U,error),
    resOp func(val T, res U, err error),
    numThreads int,
) error {
    if err:=numThreadsCheck(numThreads); err!=nil {
        return err;
    }
    taken,j:=0,0;
    jobs, results:=make(chan T), make(chan forEachParallelResult[T,U]);
    i.ForEach(func(index int, val T) (IteratorFeedback, error) {
        if j<numThreads {
            //If another worker can be created, make one
            go forEachWorker(jobs,results,workerOp);
        } else {
            //If all the worker threads are used wait for one to finish
            resOp((<-results).Unpack());
            taken++;
        }
        jobs <- val;
        j++;
        return Continue,nil;
    });
    close(jobs);
    //Need to wait for all the results!!
    for i:=0; i<j-taken; i++ { resOp((<-results).Unpack()); }
    close(results);
    return nil;
}

// This function is a consumer.
//
// This function is equivalent to [Parallel], the only difference is that the
// the inputs and outputs of the workers must be the same. It is offered as a
// convenience function.
func (i Iter[T])Parallel(
    workerOp func(val T) (T,error),
    resOp func(val T, res T, err error),
    numThreads int,
) error {
    return Parallel(i,workerOp,resOp,numThreads);
}

func numThreadsCheck(numThreads int) error {
    if numThreads<1 {
        return customerr.Wrap(
            customerr.ValOutsideRange,
            "Expected >0 | Got: %d",numThreads,
        )
    }
    return nil;
}
