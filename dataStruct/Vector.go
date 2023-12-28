// SyncedQueue/Stack/Deque/Vector
// - synced implies read+write so only need one synced struct per type
// - wont support cross over types :( (I think??)
// - remove sync from CircularBuffer
// Static Vector (Array?)
// Dynamic Vector (name?)
// wtf to do about the decrepid double linked list? - Delete
// Dynamic Deque - chunks? realloc based on f/b ratio, seed ratio with New
package dataStruct


type Vector[T any] []T

func NewVector[T any](size int) Vector[T] {
    return Vector[T](make(Vector[T],size))
}

func (c *Vector[T])Lock() { }
func (c *Vector[T])Unlock() { }
func (c *Vector[T])RLock() { }
func (c *Vector[T])RUnlock() { }

func (v Vector[T])Length() int {
    return 0
}

func (v Vector[T])Capacity() int {
    return 0
}

func (v Vector[T])SetCapacity(c int) error {
    return nil
}

func (v Vector[T])Get(idx int) (T,error){
    var tmp T
    return tmp,nil
}

func (v Vector[T])Set(val T, idx int) error {
    return nil
}

func (v Vector[T])GetPntr(idx int) (*T,error){
    return nil,nil
}

func (v Vector[T])Append(vals ...T) error {
    return nil
}

func (v Vector[T])Insert(val T, idx int) error {
    return nil
}

func (v Vector[T])Delete(idx int) error {
    return nil
}

func (v Vector[T])Clear() {

}
