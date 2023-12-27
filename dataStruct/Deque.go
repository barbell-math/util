package dataStruct

type Deque[T any] struct {
    NumElems int;
    chunkSize int
}

func (d *Deque[T])Length() int {
    return d.NumElems
}

func (d *Deque[T])Capacity() int {
    //return d.chunks.Length()*d.chunkSize
    return 0
}

// func (d *Deque[T])PopFront() (T,error) {
// }
