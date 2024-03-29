package dataStruct;

import (
    "fmt"
    customerr "github.com/barbell-math/util/err"
)

var Full,IsFull=customerr.ErrorFactory(
    "The capacity of the container has been reached.",
);

var Empty,IsEmpty=customerr.ErrorFactory(
    "The container is empty.",
);

func getIndexOutOfBoundsError(idx int, _len int) error {
    return customerr.ValOutsideRange(fmt.Sprintf(
        "Index out of bounds. | NumElems: %d Index: %d",
        _len,idx,
    ));
}
