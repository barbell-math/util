package containers

import (
	"sync"

	"github.com/barbell-math/util/algo/hash"
	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/algo/widgets"
	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/customerr"
)

type (
	wrapingIndex int

	// A container that holds a fixed number of values in such a way that makes
	// stack and queue operations extremely efficient. Because the length of the
	// container is fixed (it will not dynamically expand to add more elements
	// as needed) the values in the underlying array will 'rotate' around the
	// array as operations are performed making it so no allocations are ever
	// performed beyond the initial creation of the underlying array.
	CircularBuffer[T any, U widgets.WidgetInterface[T]] struct {
		vals     []T
		numElems int
		start    wrapingIndex
	}

	// A synchronized version of CircularBuffer. All operations will be wrapped
	// in the appropriate calls the embedded RWMutex. A pointer to a RWMutex is
	// embedded rather than a value to avoid copying the lock value.
	SyncedCircularBuffer[T any, U widgets.WidgetInterface[T]] struct {
		*sync.RWMutex
		CircularBuffer[T, U]
	}
)

func (w wrapingIndex) normalize(wrapThreshold int) wrapingIndex {
	rv := int(w) % wrapThreshold // Takes care of positive bounds
	for ; rv < 0; rv += wrapThreshold {
	} // Takes care of negative bounds
	return wrapingIndex(rv)
}

func (w wrapingIndex) Add(amnt int, wrapThreshold int) wrapingIndex {
	rv := w + wrapingIndex(amnt)
	return rv.normalize(wrapThreshold)
}

func (w wrapingIndex) Sub(amnt int, wrapThreshold int) wrapingIndex {
	rv := w - wrapingIndex(amnt)
	return rv.normalize(wrapThreshold)
}

func (start wrapingIndex) GetProperIndex(idx int, wrapThreshold int) wrapingIndex {
	return start.Add(idx, wrapThreshold) % wrapingIndex(wrapThreshold)
}

// Creates a new CircularBuffer initialized with size zero valued elements. Size
// must be greater than 0, an error will be returned if it is not.
func NewCircularBuffer[T any, U widgets.WidgetInterface[T]](
	size int,
) (CircularBuffer[T, U], error) {
	if size <= 0 {
		return CircularBuffer[T, U]{}, customerr.Wrap(
			customerr.ValOutsideRange,
			"Size of buffer must be >0 | Have: %d", size,
		)
	}
	return CircularBuffer[T, U]{
		vals:  make([]T, size),
		start: 0,
	}, nil
}

// Creates a new synced CircularBuffer initialized with size zero valued
// elements. Size must be greater than 0, an error will be returned if it is not.
// The underlying RWMutex value will be fully unlocked upon initialization.
func NewSyncedCircularBuffer[T any, U widgets.WidgetInterface[T]](
	size int,
) (SyncedCircularBuffer[T, U], error) {
	rv, err := NewCircularBuffer[T, U](size)
	return SyncedCircularBuffer[T, U]{
		CircularBuffer: rv,
		RWMutex:        &sync.RWMutex{},
	}, err
}

// Converts the supplied map to a syncronized map. Beware: The original
// non-synced circular buffer will remain useable.
func (c *CircularBuffer[T, U]) ToSynced() SyncedCircularBuffer[T, U] {
	return SyncedCircularBuffer[T, U]{
		RWMutex:        &sync.RWMutex{},
		CircularBuffer: *c,
	}
}

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (c *CircularBuffer[T, U]) Lock() {}

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (c *CircularBuffer[T, U]) Unlock() {}

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (c *CircularBuffer[T, U]) RLock() {}

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (c *CircularBuffer[T, U]) RUnlock() {}

// The SyncedCircularBuffer method to override the HashMap pass through function
// and actually apply the mutex operation.
func (c *SyncedCircularBuffer[T, U]) Lock() { c.RWMutex.Lock() }

// The SyncedCircularBuffer method to override the HashMap pass through function
// and actually apply the mutex operation.
func (c *SyncedCircularBuffer[T, U]) Unlock() { c.RWMutex.Unlock() }

// The SyncedCircularBuffer method to override the HashMap pass through function
// and actually apply the mutex operation.
func (c *SyncedCircularBuffer[T, U]) RLock() { c.RWMutex.RLock() }

// The SyncedCircularBuffer method to override the HashMap pass through function
// and actually apply the mutex operation.
func (c *SyncedCircularBuffer[T, U]) RUnlock() { c.RWMutex.RUnlock() }

// Returns true, a circular biffer is addressable.
func (c *CircularBuffer[T, U]) IsAddressable() bool { return true }

// Returns false, a circular buffer is not synced.
func (c *CircularBuffer[T, U]) IsSynced() bool { return false }

// Returns true, a synced circular buffer is synced.
func (c *SyncedCircularBuffer[T, U]) IsSynced() bool { return true }

// Description: Returns true if the circular buffer has reached its capacity,
// false otherwise.
//
// Time Complexity: O(1)
func (c *CircularBuffer[T, U]) Full() bool {
	return c.numElems == len(c.vals)
}

// Description: Places a read lock on the underlying circular buffer and then
// calls the underlying circular buffers [CircularBuffer.Length] method.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (c *SyncedCircularBuffer[T, U]) Full() bool {
	c.RLock()
	defer c.RUnlock()
	return c.CircularBuffer.Full()
}

// Description: Returns the length of the circular buffer.
//
// Time Complexity: O(1)
func (c *CircularBuffer[T, U]) Length() int {
	return c.numElems
}

// Description: Returns the length of the underlying circular buffer.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (c *SyncedCircularBuffer[T, U]) Length() int {
	c.RLock()
	defer c.RUnlock()
	return c.CircularBuffer.numElems
}

// Description: Returns the capacity of the circular buffer.
//
// Time Complexity: O(1)
func (c *CircularBuffer[T, U]) Capacity() int {
	return len(c.vals)
}

// Description: Returns the capacity of the underlying circular buffer.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (c *SyncedCircularBuffer[T, U]) Capacity() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.vals)
}

// Description: Pushes an element to the front of the circular buffer.
// Equivalent to inserting a single value at the front of the circular buffer.
// Values will be pushed to the front in the order that they are given. For
// example, calling push front on [0,1,2] with vals of [3,4] will result in
// [3,4,0,1,2].
//
// Time Complexity: O(n+m), where m=len(vals)
func (c *CircularBuffer[T, U]) PushFront(v ...T) error {
	return c.pushFrontImpl(v)
}

// Description: Places a write lock on the underlying circular buffer and then
// pushes values to the front of the circular buffer. Exhibits the same behavior
// as [CircularBuffer.PusFront]. The underlying [CircularBuffer.PushFront]
// method is not called to avoid copying the list of values twice, which could
// be inefficient with a large type for the T generic or many values.
//
// Lock Type: Write
//
// Time Complexity: O(n+m), where m=len(vals)
func (c *SyncedCircularBuffer[T, U]) PushFront(v ...T) error {
	c.Lock()
	defer c.Unlock()
	return c.CircularBuffer.pushFrontImpl(v)
}

func (c *CircularBuffer[T, U]) pushFrontImpl(vals []T) error {
	for i := len(vals) - 1; i >= 0; i-- {
		if c.numElems >= len(c.vals) {
			return getFullError(len(c.vals))
		}
		c.start = c.start.Sub(1, len(c.vals))
		c.vals[c.start] = vals[i]
		c.numElems++
	}
	return nil
}

// Description: Pushes an element to the back of the circular buffer. Equivalent
// to appending values to the end of the circular buffer. Values will be pushed
// back in the order that they are given. For example, calling push back on
// [0,1,2] with vals of [3,4] will result in [0,1,2,3,4].
//
// Time Complexity: best case O(m), where m=len(vals)
func (c *CircularBuffer[T, U]) PushBack(v ...T) error {
	return c.pushBackImpl(v)
}

// Description: Places a write lock on the underlying circular buffer and then
// appends values to the end of the circular buffer. Exhibits the same behavior
// as [CircularBuffer.PushBack]. The underlying [CircularBuffer.PushBack] method
// is not called to avoid copying the list of values twice, which could be
// inefficient with a large type for the T generic or many values.
//
// Lock Type: Write
//
// Time Complexity: best case O(m), where m=len(vals)
func (c *SyncedCircularBuffer[T, U]) PushBack(v ...T) error {
	c.Lock()
	defer c.Unlock()
	return c.CircularBuffer.pushBackImpl(v)
}

func (c *CircularBuffer[T, U]) pushBackImpl(vals []T) error {
	for i := 0; i < len(vals); i++ {
		if c.numElems >= len(c.vals) {
			return getFullError(len(c.vals))
		}
		c.vals[c.start.Add(c.numElems, len(c.vals))] = vals[i]
		c.numElems++
	}
	return nil
}

// Description: Pushes an element to the back of the circular buffer. Equivalent
// to appending a single value to the end of the circular buffer.
//
// Time Complexity: O(m), where m=len(vals)
func (c *CircularBuffer[T, U]) ForcePushBack(v ...T) {
	c.forcePushBackImpl(v)
}

// Description: Places a write lock on the underlying circular buffer and then
// pushes values to the front of the circular buffer. Exhibits the same behavior
// as [CircularBuffer.ForcePushBack]. The underlying
// [CircularBuffer.ForcePushBack] method is not called to avoid copying the list
// of values twice, which could be inefficient with a large type for the T
// generic or many values.
//
// Lock Type: Write
//
// Time Complexity: O(m), where m=len(vals)
func (c *SyncedCircularBuffer[T, U]) ForcePushBack(v ...T) {
	c.Lock()
	defer c.Unlock()
	c.CircularBuffer.forcePushBackImpl(v)
}

func (c *CircularBuffer[T, U]) forcePushBackImpl(vals []T) {
	w := widgets.Widget[T, U]{}
	maxVals := min(len(c.vals), len(vals))
	leftoverSpace := len(c.vals) - c.numElems
	if maxVals > leftoverSpace {
		for i := 0; i < maxVals-leftoverSpace; i++ {
			w.Zero(&c.vals[c.start])
			c.start = c.start.Add(1, len(c.vals))
			c.numElems--
		}
	}
	for i := maxVals - 1; i >= 0; i-- {
		c.numElems++
		c.vals[c.inclusiveEnd()] = vals[len(vals)-i-1]
	}
}

// Description: Pushes an element to the back of the circular buffer, poping an
// element from the front of the buffer if necessary to make room for the new
// element. If the circular buffer is full then this is equavilent to poping and
// then pushing, but more efficient.
//
// Time Complexity: O(n+m), where m=len(vals)
func (c *CircularBuffer[T, U]) ForcePushFront(v ...T) {
	c.forcePushFrontImpl(v)
}

// Description: Places a write lock on the underlying circular buffer and then
// pushes values to the front of the circular buffer. Exhibits the same behavior
// as [CircularBuffer.ForcePushFront]. The underlying
// [CircularBuffer.ForcePushFront] method is not called to avoid copying the
// list of values twice, which could be inefficient with a large type for the T
// generic or many values.
//
// Lock Type: Write
//
// Time Complexity: O(n+m), where m=len(vals)
func (c *SyncedCircularBuffer[T, U]) ForcePushFront(v ...T) {
	c.Lock()
	defer c.Unlock()
	c.CircularBuffer.forcePushFrontImpl(v)
}

func (c *CircularBuffer[T, U]) forcePushFrontImpl(vals []T) {
	w := widgets.Widget[T, U]{}
	maxVals := min(len(c.vals), len(vals))
	leftoverSpace := len(c.vals) - c.numElems
	if maxVals > leftoverSpace {
		for i := 0; i < maxVals-leftoverSpace; i++ {
			w.Zero(&c.vals[c.inclusiveEnd()])
			c.numElems--
		}
	}
	for i := maxVals - 1; i >= 0; i-- {
		c.start = c.start.Sub(1, len(c.vals))
		c.numElems++
		c.vals[c.start.GetProperIndex(0, len(c.vals))] = vals[i]
	}
}

// Description: Returns the value at index 0 if one is present. If the circular
// buffer has no elements then an error is returned.
//
// Time Complexity: O(1)
func (c *CircularBuffer[T, U]) PeekFront() (T, error) {
	if c.numElems > 0 {
		return c.vals[c.start], nil
	}
	var tmp T
	return tmp, getIndexOutOfBoundsError(0, 0, c.numElems)
}

// Description: Places a read lock on the underlying circular buffer and then
// attempts to return the value at index 0 if one is present. Exhibits the same
// behavior as the [CircularBuffer.PeekFront] method. The underlying
// [CircularBuffer.PeekFront] method is not called to avoid copying the value
// twice, which could be inefficient with a large type for the T generic.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (c *SyncedCircularBuffer[T, U]) PeekFront() (T, error) {
	c.RLock()
	defer c.RUnlock()
	if c.numElems > 0 {
		return c.vals[c.start], nil
	}
	var tmp T
	return tmp, getIndexOutOfBoundsError(0, 0, c.numElems)
}

// Description: Returns a pointer to the value at index 0 if one is present. If
// the circular buffer has no elements then an error is returned.
//
// Time Complexity: O(1)
func (c *CircularBuffer[T, U]) PeekPntrFront() (*T, error) {
	if c.numElems > 0 {
		return &c.vals[c.start], nil
	}
	return nil, getIndexOutOfBoundsError(0, 0, c.numElems)
}

// Description: Places a read lock on the underlying circular buffer and then
// calls the underlying circular buffers [CircularBuffer.PeekPntrFront] method.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (c *SyncedCircularBuffer[T, U]) PeekPntrFront() (*T, error) {
	c.RLock()
	defer c.RUnlock()
	return c.CircularBuffer.PeekPntrFront()
}

// Description: Returns the value at index len(v)-1 if one is present. If the
// circular buffer has no elements then an error is returned.
//
// Time Complexity: O(1)
func (c *CircularBuffer[T, U]) PeekBack() (T, error) {
	if c.numElems > 0 {
		return c.vals[c.inclusiveEnd()], nil
	}
	var tmp T
	return tmp, getIndexOutOfBoundsError(0, 0, c.numElems)
}

// Description: Places a read lock on the underlying circular buffer and then
// attempts to return the value at index len(v)-1 if one is present. Exhibits
// the same behavior as the [circular buffer.PeekBack] method. The underlying
// [Vector.PeekBack] method is not called to avoid copying the value twice,
// which could be inefficient with a large type for the T generic.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (c *SyncedCircularBuffer[T, U]) PeekBack() (T, error) {
	c.RLock()
	defer c.RUnlock()
	if c.numElems > 0 {
		return c.vals[c.inclusiveEnd()], nil
	}
	var tmp T
	return tmp, getIndexOutOfBoundsError(0, 0, c.numElems)
}

// Description: Returns a pointer to the value at index len(v)-1 if one is
// present. If the circular buffer has no elements then an error is returned.
//
// Time Complexity: O(1)
func (c *CircularBuffer[T, U]) PeekPntrBack() (*T, error) {
	if c.numElems > 0 {
		return &c.vals[c.inclusiveEnd()], nil
	}
	return nil, getIndexOutOfBoundsError(0, 0, c.numElems)
}

// Description: Places a read lock on the underlying circular buffer and then
// calls the underlying circular buffer [CircularBuffer.PeekPntrBack] method.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (c *SyncedCircularBuffer[T, U]) PeekPntrBack() (*T, error) {
	c.RLock()
	defer c.RUnlock()
	return c.CircularBuffer.PeekPntrBack()
}

// Description: Gets the value at the specified index. Returns an error if the
// index is >= the length of the circular buffer.
//
// Time Complexity: O(1)
func (c *CircularBuffer[T, U]) Get(idx int) (T, error) {
	if idx >= 0 && idx < c.numElems && c.numElems > 0 {
		return c.vals[c.start.GetProperIndex(idx, len(c.vals))], nil
	}
	var tmp T
	return tmp, getIndexOutOfBoundsError(idx, 0, c.numElems)
}

// Description: Places a read lock on the underlying circular buffer and then
// gets the value at the specified index. Exhibits the same behavior as the
// [CircularBuffer.Get] method. The underlying [CircularBuffer.Get] method is
// not called to avoid copying the return value twice, which could be
// inefficient with a large value for the T generic.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (c *SyncedCircularBuffer[T, U]) Get(idx int) (T, error) {
	c.RLock()
	defer c.RUnlock()
	if idx >= 0 && idx < c.numElems && c.numElems > 0 {
		return c.vals[c.start.GetProperIndex(idx, len(c.vals))], nil
	}
	var tmp T
	return tmp, getIndexOutOfBoundsError(idx, 0, c.numElems)
}

// Description: Gets a pointer to the value at the specified index. Returns an
// error if the index is >= the length of the circular buffer.
//
// Time Complexity: O(1)
func (c *CircularBuffer[T, U]) GetPntr(idx int) (*T, error) {
	if idx >= 0 && idx < c.numElems && c.numElems > 0 {
		return &c.vals[c.start.GetProperIndex(idx, len(c.vals))], nil
	}
	return nil, getIndexOutOfBoundsError(idx, 0, c.numElems)
}

// Description: Places a read lock on the underlying circular buffer and then
// calls the underlying circular buffer [CircularBuffer.GetPntr] method.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (c *SyncedCircularBuffer[T, U]) GetPntr(idx int) (*T, error) {
	c.RLock()
	defer c.RUnlock()
	return c.CircularBuffer.GetPntr(idx)
}

// Description: Contains will return true if the supplied value is in the
// circular buffer, false otherwise. All equality comparisons are performed by
// the generic U widget type that the circular buffer was initialized with.
//
// Time Complexity: O(n) (linear search)
func (c *CircularBuffer[T, U]) Contains(val T) bool {
	return c.ContainsPntr(&val)
}

// Description: Places a read lock on the underlying circular buffer and then
// calls the underlying circular buffers [CircularBuffer.ContainsPntr] method.
//
// Lock Type: Read
//
// Time Complexity: O(n) (linear search)
func (c *SyncedCircularBuffer[T, U]) Contains(val T) bool {
	c.RLock()
	defer c.RUnlock()
	return c.CircularBuffer.ContainsPntr(&val)
}

// Description: ContainsPntr will return true if the supplied value is in the
// circular buffer, false otherwise. All equality comparisons are performed by
// the generic U widget type that the circular buffer was initialized with.
//
// Time Complexity: O(n) (linear search)
func (c *CircularBuffer[T, U]) ContainsPntr(val *T) bool {
	found := false
	w := widgets.Widget[T, U]{}
	for i := 0; i < c.numElems && !found; i++ {
		found = w.Eq(val, &c.vals[c.start.GetProperIndex(i, len(c.vals))])
	}
	return found
}

// Description: Places a read lock on the underlying circular buffer and then
// calls the underlying circular buffer [CircularBuffer.ContainsPntr] method.
//
// Lock Type: Read
//
// Time Complexity: O(n) (linear search)
func (c *SyncedCircularBuffer[T, U]) ContainsPntr(val *T) bool {
	c.RLock()
	defer c.RUnlock()
	return c.CircularBuffer.ContainsPntr(val)
}

// Description: KeyOf will return the index of the first occurrence of the
// supplied value in the circular buffer. If the value is not found then the
// returned index will be -1 and the boolean flag will be set to false. If the
// value is found then the boolean flag will be set to true. All equality
// comparisons are performed by the generic U widget type that the circular buffer was
// initialized with.
//
// Time Complexity: O(n) (linear search)
func (c *CircularBuffer[T, U]) KeyOf(val T) (int, bool) {
	return c.keyOfImpl(&val)
}

// Description: Places a read lock on the underlying circular buffer and then
// calls the underlying circular buffers [CircularBuffer.KeyOf] implemenation
// method. The [CircularBuffer.KeyOf] method is not called directly to avoid
// copying the val variable twice, which could be expensive with a large type
// for the T generic.
//
// Lock Type: Read
//
// Time Complexity: O(n) (linear search)
func (c *SyncedCircularBuffer[T, U]) KeyOf(val T) (int, bool) {
	c.RLock()
	defer c.RUnlock()
	return c.CircularBuffer.keyOfImpl(&val)
}

func (c *CircularBuffer[T, U]) keyOfImpl(val *T) (int, bool) {
	w := widgets.Widget[T, U]{}
	for i := 0; i < c.numElems; i++ {
		if w.Eq(val, &c.vals[c.start.GetProperIndex(i, len(c.vals))]) {
			return i, true
		}
	}
	return -1, false
}

// Description: Sets the values at the specified indexes. Returns an error if
// the index is >= the length of the circular buffer. Stops setting values as
// soon as an error is encountered.
//
// Time Complexity: O(m), where m=len(vals)
func (c *CircularBuffer[T, U]) Set(vals ...basic.Pair[int, T]) error {
	return c.setImpl(vals)
}

// Description: Places a write lock on the underlying circular buffer and then
// calls the underlying circular buffers [CircularBuffer.Set] implementaiton
// method. The [CircularBuffer.Set] method is not called directly to avoid
// copying the vals varargs twice, which could be expensive with a large type
// for the T generic or a large number of values.
//
// Lock Type: Write
//
// Time Complexity: O(m), where m=len(vals)
func (c *SyncedCircularBuffer[T, U]) Set(vals ...basic.Pair[int, T]) error {
	c.Lock()
	defer c.Unlock()
	return c.CircularBuffer.setImpl(vals)
}

func (c *CircularBuffer[T, U]) setImpl(vals []basic.Pair[int, T]) error {
	for _, iterV := range vals {
		if iterV.A >= 0 && iterV.A < c.numElems && len(c.vals) > 0 {
			c.vals[c.start.GetProperIndex(iterV.A, len(c.vals))] = iterV.B
		} else {
			return getIndexOutOfBoundsError(iterV.A, 0, c.numElems)
		}
	}
	return nil
}

// Description: Sets the supplied values sequentially starting at the supplied
// index and continuing sequentailly after that. Returns and error if any index
// that is attempted to be set is >= the length of the circular buffer. If an
// error occurs, all values will be set up until the value that caused the error.
//
// Time Complexity: O(m), where m=len(vals)
func (c *CircularBuffer[T, U]) SetSequential(idx int, vals ...T) error {
	return c.setSequentialImpl(idx, vals)
}

// Description: Places a write lock on the underlying circular buffer and then
// calls the underlying circular buffers [CircularBuffer.SetSequential]
// implementation method. The [circular buffer.SetSequential] method is not
// called directly to avoid copying the vals varargs twice, which could be
// expensive with a large type for the T generic or a large number of values.
//
// Lock Type: Write
//
// Time Complexity: O(m), where m=len(vals)
func (c *SyncedCircularBuffer[T, U]) SetSequential(idx int, vals ...T) error {
	c.Lock()
	defer c.Unlock()
	return c.CircularBuffer.setSequentialImpl(idx, vals)
}

func (c *CircularBuffer[T, U]) setSequentialImpl(idx int, vals []T) error {
	if idx >= len(c.vals) {
		return getIndexOutOfBoundsError(idx, 0, len(c.vals))
	}
	numCopyableVals := min(len(c.vals)-idx, len(vals))
	copy((c.vals)[idx:idx+numCopyableVals], vals[0:numCopyableVals])
	if idx+len(vals) > len(c.vals) {
		return getIndexOutOfBoundsError(len(c.vals), 0, len(c.vals))
	}
	return nil
}

// Description: Appends the supplied values to the circular buffer. This function
// will never return an error.
//
// Time Complexity: Best case O(m), where m=len(vals).
func (c *CircularBuffer[T, U]) Append(vals ...T) error {
	return c.appendImpl(vals)
}

// Description: Places a write lock on the underlying circular buffer and then
// calls the underlying circular buffers [CircularBuffer.Append] implementation
// method. The [CircularBuffer.Append] method is not called directly to avoid
// copying the vals varargs twice, which could be expensive with a large type
// for the T generic or a large number of values.
//
// Lock Type: Write
//
// Time Complexity: Best case O(m), where m=len(vals).
func (c *SyncedCircularBuffer[T, U]) Append(vals ...T) error {
	c.Lock()
	defer c.Unlock()
	return c.CircularBuffer.appendImpl(vals)
}

func (c *CircularBuffer[T, U]) appendImpl(vals []T) error {
	for i := 0; i < len(vals); i++ {
		if c.numElems < len(c.vals) {
			c.numElems++
			c.vals[c.inclusiveEnd()] = vals[i]
		} else {
			return getFullError(len(c.vals))
		}
	}
	return nil
}

// Description: AppendUnique will append the supplied values to the circular
// buffer if they are not already present in the circular buffer (unique).
// Non-unique values will not be appended. This function will never return an
// error.
//
// Time Complexity: Best case O(m), where m=len(vals).
func (c *CircularBuffer[T, U]) AppendUnique(vals ...T) error {
	var rv error
	for i := 0; i < len(vals) && rv == nil; i++ {
		rv = c.appendUniqueImpl(&vals[i])
	}
	return rv
}

// Description: Places a write lock on the underlying circular buffer and then
// calls the underlying circular buffer [CircularBuffer.AppendUnique]
// implementaion method. The [CircularBuffer.AppendUnique] method is not called
// directly to avoid copying the vals varargs twice, which could be expensive
// with a large type for the T generic or a large number of values.
//
// Lock Type: Write
//
// Time Complexity: Best case O(m), where m=len(vals).
func (c *SyncedCircularBuffer[T, U]) AppendUnique(vals ...T) error {
	c.Lock()
	defer c.Unlock()
	var rv error
	for i := 0; i < len(vals) && rv == nil; i++ {
		rv = c.CircularBuffer.appendUniqueImpl(&vals[i])
	}
	return rv
}

func (c *CircularBuffer[T, U]) appendUniqueImpl(val *T) error {
	if c.ContainsPntr(val) {
		return nil
	}
	if c.numElems >= len(c.vals) {
		return getFullError(len(c.vals))
	}
	c.numElems++
	c.vals[c.inclusiveEnd()] = *val
	return nil
}

// Description: Insert will insert the supplied values into the circular buffer.
// The values will be inserted in the order that they are given.
//
// Time Complexity: O(n*m), where m=len(vals)
func (c *CircularBuffer[T, U]) Insert(vals ...basic.Pair[int, T]) error {
	return c.insertImpl(vals)
}

// Description: Places a write lock on the underlying circular buffer and then
// calls the underlying circular buffer [CircularBuffer.Insert] implementation
// method. The [CircularBuffer.Insert] method is not called directly to avoid
// copying the vals varargs twice, which could be expensive with a large type
// for the T generic or a large number of values.
//
// Lock Type: Write
//
// Time Complexity: O(n*m), where m=len(vals)
func (c *SyncedCircularBuffer[T, U]) Insert(vals ...basic.Pair[int, T]) error {
	c.Lock()
	defer c.Unlock()
	return c.CircularBuffer.insertImpl(vals)
}

func (c *CircularBuffer[T, U]) insertImpl(vals []basic.Pair[int, T]) error {
	for _, iterV := range vals {
		if iterV.A < 0 || iterV.A > c.numElems {
			return getIndexOutOfBoundsError(iterV.A, 0, c.numElems)
		} else if c.numElems == len(c.vals) {
			return getFullError(len(c.vals))
		}
		if c.distanceFromBack(iterV.A) > c.distanceFromFront(iterV.A) {
			c.insertMoveFront(&iterV)
		} else {
			c.insertMoveBack(&iterV)
		}
	}
	return nil
}

func (c *CircularBuffer[T, U]) insertMoveFront(val *basic.Pair[int, T]) {
	c.numElems += 1
	c.start = c.start.Sub(1, len(c.vals))
	for j, i := 0, 1; i < val.A+1; i++ {
		c.vals[c.start.GetProperIndex(j, len(c.vals))] = c.vals[c.start.GetProperIndex(i, len(c.vals))]
		j++
	}
	c.vals[c.start.GetProperIndex(val.A, len(c.vals))] = val.B
}

func (c *CircularBuffer[T, U]) insertMoveBack(val *basic.Pair[int, T]) {
	c.numElems += 1
	for j, i := c.numElems-1, c.numElems-2; i >= val.A; i-- {
		c.vals[c.start.GetProperIndex(j, len(c.vals))] = c.vals[c.start.GetProperIndex(i, len(c.vals))]
		j--
	}
	c.vals[c.start.GetProperIndex(val.A, len(c.vals))] = val.B
}

// Description: Inserts the supplied values at the given index. Returns an error
// if the index is >= the length of the circular buffer.
//
// Time Complexity: O(n+m), where m=len(vals)
func (c *CircularBuffer[T, U]) InsertSequential(idx int, vals ...T) error {
	return c.insertSequentailImpl(idx, vals)
}

// Description: Places a write lock on the underlying circular buffer and then
// calls the underlying circular buffer [CircularBuffer.InsertSequential]
// implementation method. The [CircularBuffer.InsertSequential] method is not
// called directly to avoid copying the vals varargs twice, which could be
// expensive with a large type for the T generic or a large number of values.
//
// Lock Type: Write
//
// Time Complexity: O(n+m), where m=len(vals)
func (c *SyncedCircularBuffer[T, U]) InsertSequential(idx int, vals ...T) error {
	c.Lock()
	defer c.Unlock()
	return c.CircularBuffer.insertSequentailImpl(idx, vals)
}

func (c *CircularBuffer[T, U]) insertSequentailImpl(idx int, vals []T) error {
	if idx < 0 || idx > c.numElems {
		return getIndexOutOfBoundsError(idx, 0, c.numElems)
	} else if c.numElems == len(c.vals) {
		return getFullError(len(c.vals))
	}
	maxVals := len(vals)
	if c.numElems+maxVals > len(c.vals) {
		maxVals = len(c.vals) - c.numElems
	}
	if c.distanceFromBack(idx) > c.distanceFromFront(idx) {
		c.insertSequentialMoveFront(idx, vals, maxVals)
	} else {
		c.insertSequentialMoveBack(idx, vals, maxVals)
	}
	if maxVals < len(vals) {
		return getFullError(len(c.vals))
	}
	return nil
}

func (c *CircularBuffer[T, U]) insertSequentialMoveFront(idx int, v []T, maxVals int) {
	c.numElems += maxVals
	c.start = c.start.Sub(maxVals, len(c.vals))
	for j, i := 0, maxVals-1; i < idx+maxVals; i++ {
		c.vals[c.start.GetProperIndex(j, len(c.vals))] = c.vals[c.start.GetProperIndex(i, len(c.vals))]
		j++
	}
	for i := idx; i < idx+maxVals; i++ {
		c.vals[c.start.GetProperIndex(i, len(c.vals))] = v[i-idx]
	}
}

func (c *CircularBuffer[T, U]) insertSequentialMoveBack(idx int, v []T, maxVals int) {
	c.numElems += maxVals
	for j, i := c.numElems-1, c.numElems-maxVals-1; i >= idx; i-- {
		c.vals[c.start.GetProperIndex(j, len(c.vals))] = c.vals[c.start.GetProperIndex(i, len(c.vals))]
		j--
	}
	for i := idx; i < idx+maxVals; i++ {
		c.vals[c.start.GetProperIndex(i, len(c.vals))] = v[i-idx]
	}
}

// Description: Returns and removes the element at the front of the circular
// buffer. Returns an error if the vector has no elements.
//
// Time Complexity: O(n)
func (c *CircularBuffer[T, U]) PopFront() (T, error) {
	var rv T
	return rv, c.popFrontImpl(&rv)
}

// Description: Places a write lock on the underlying circular buffer and then
// calls the underlying circular buffer [CircularBuffer.PopFront] implementation
// method. The [CircularBuffer.PopFront] method is not called directly to avoid
// copying the return value twice, which could be expensive with a large type
// for the T generic.
//
// Lock Type: Write
//
// Time Complexity: O(n)
func (c *SyncedCircularBuffer[T, U]) PopFront() (T, error) {
	c.Lock()
	defer c.Unlock()
	var rv T
	return rv, c.CircularBuffer.popFrontImpl(&rv)
}

func (c *CircularBuffer[T, U]) popFrontImpl(rv *T) error {
	w := widgets.Widget[T, U]{}
	if c.numElems > 0 {
		*rv = c.vals[c.start]
		w.Zero(&c.vals[c.start])
		c.start = c.start.Add(1, len(c.vals))
		c.numElems--
		return nil
	}
	return containerTypes.Empty
}

// Description: Returns and removes the element at the back of the circular
// buffer. Returns an error if the vector has no elements.
//
// Time Complexity: O(1)
func (c *CircularBuffer[T, U]) PopBack() (T, error) {
	var rv T
	return rv, c.popBackImpl(&rv)
}

// Description: Places a write lock on the underlying circular buffer and then
// calls the underlying circular buffers [CircularBuffer.PopFront]
// implementation method. The [CircularBuffer.PopBack] method is not called
// directly to avoid copying the return value twice, which could be expensive
// with a large type for the T generic.
//
// Lock Type: Write
//
// Time Complexity: O(1)
func (c *SyncedCircularBuffer[T, U]) PopBack() (T, error) {
	c.Lock()
	defer c.Unlock()
	var rv T
	return rv, c.popBackImpl(&rv)
}

func (c *CircularBuffer[T, U]) popBackImpl(rv *T) error {
	w := widgets.Widget[T, U]{}
	if c.numElems > 0 {
		*rv = c.vals[c.inclusiveEnd()]
		w.Zero(&c.vals[c.inclusiveEnd()])
		c.numElems--
		return nil
	}
	return containerTypes.Empty
}

// Description: Deletes the value at the specified index. Returns an error if
// the index is >= the length of the circular buffer.
//
// Time Complexity: O(n)
func (c *CircularBuffer[T, U]) Delete(idx int) error {
	if idx < 0 || idx >= c.numElems {
		return getIndexOutOfBoundsError(idx, 0, c.numElems)
	}
	w := widgets.Widget[T, U]{}
	w.Zero(&c.vals[c.start.GetProperIndex(idx, len(c.vals))])
	if c.numElems == 1 && idx == 0 {
		c.numElems--
	} else if c.distanceFromBack(idx) > c.distanceFromFront(idx) {
		c.deleteMoveFront(idx)
	} else {
		c.deleteMoveBack(idx)
	}
	return nil
}

// Description: Places a write lock on the underlying circular buffer and then
// calls the underlying circular buffers [CircularBuffer.Delete] method.
//
// Lock Type: Write
//
// Time Complexity: O(n)
func (c *SyncedCircularBuffer[T, U]) Delete(idx int) error {
	c.Lock()
	defer c.Unlock()
	return c.CircularBuffer.Delete(idx)
}

func (c *CircularBuffer[T, U]) deleteMoveFront(idx int) {
	for i := idx - 1; i >= 0; i-- {
		c.vals[c.start.GetProperIndex(i+1, len(c.vals))] = c.vals[c.start.GetProperIndex(i, len(c.vals))]
	}
	c.numElems--
	c.start = c.start.Add(1, len(c.vals))
}

func (c *CircularBuffer[T, U]) deleteMoveBack(idx int) {
	for i := idx; i < c.numElems; i++ {
		c.vals[c.start.GetProperIndex(i, len(c.vals))] = c.vals[c.start.GetProperIndex(i+1, len(c.vals))]
	}
	c.numElems--
}

// Description: Deletes the values in the index range [start,end). Returns an
// error if the start index is < 0, the end index is >= the length of the
// circular buffer, or the end index is < the start index.
//
// Time Complexity: O(n)
func (c *CircularBuffer[T, U]) DeleteSequential(start int, end int) error {
	if start < 0 {
		return getIndexOutOfBoundsError(start, 0, c.numElems)
	}
	if end >= c.numElems {
		return getIndexOutOfBoundsError(end, 0, c.numElems)
	}
	if end <= start {
		return getStartEndIndexError(start, end)
	}
	w := widgets.Widget[T, U]{}
	for i := start; i < end; i++ {
		w.Zero(&c.vals[c.start.GetProperIndex(i, c.numElems)])
	}
	if c.numElems == 1 && start == 0 && end < 0 {
		c.numElems--
	} else if c.distanceFromBack((start+end)/2) > c.distanceFromFront((start+end)/2) {
		c.deleteSequentialMoveFront(start, end)
	} else {
		c.deleteSequentialMoveBack(start, end)
	}
	return nil
}

// Description: Places a write lock on the underlying circular buffer and then
// calls the underlying circular buffer [CircularBuffer.DeleteSequential] method.
//
// Lock Type: Write
//
// Time Complexity: O(n)
func (c *SyncedCircularBuffer[T, U]) DeleteSequential(start int, end int) error {
	c.Lock()
	defer c.Unlock()
	return c.CircularBuffer.DeleteSequential(start, end)
}

func (c *CircularBuffer[T, U]) deleteSequentialMoveFront(start int, end int) {
	for i, j := start-1, 0; i >= 0; i-- {
		c.vals[c.start.GetProperIndex(end-1-j, len(c.vals))] = c.vals[c.start.GetProperIndex(i, len(c.vals))]
		j++
	}
	c.numElems -= (end - start)
	c.start = c.start.Add(end-start, len(c.vals))
}

func (c *CircularBuffer[T, U]) deleteSequentialMoveBack(start int, end int) {
	for i, j := end, 0; i < c.numElems; i++ {
		c.vals[c.start.GetProperIndex(start+j, len(c.vals))] = c.vals[c.start.GetProperIndex(i, len(c.vals))]
		j++
	}
	c.numElems -= (end - start)
}

// Description: Pop will remove all occurrences of val in the circular buffer.
// All equality comparisons are performed by the generic U widget type that the
// circular buffer was initialized with.
//
// Time Complexity: O(n)
func (c *CircularBuffer[T, U]) Pop(val T) int {
	return c.popSequentialImpl(&val, containerTypes.PopAll)
}

// Description: Places a write lock on the underlying circular buffer and then
// calls the underlying circular buffers [CircularBuffer.Pop] implementation
// method. The [CircularBuffer.Pop] method is not called directly to avoid
// copying the value twice, which could be expensive with a large type for the
// T generic or a large number of values.
//
// Lock Type: Write
//
// Time Complexity: O(n)
func (c *SyncedCircularBuffer[T, U]) Pop(val T) int {
	c.Lock()
	defer c.Unlock()
	return c.CircularBuffer.popSequentialImpl(&val, containerTypes.PopAll)
}

// Description: PopPntr will remove all occurrences of val in the circular
// buffer. All equality comparisons are performed by the generic U widget type
// that the circular buffer was initialized with.
//
// Time Complexity: O(n)
func (c *CircularBuffer[T, U]) PopPntr(val *T) int {
	return c.popSequentialImpl(val, containerTypes.PopAll)
}

// Description: Places a write lock on the underlying circular buffer and then
// calls the underlying circular buffers [CircularBuffer.PopPntr] implementation
// method. The [CircularBuffer.PopPntr] method is not called directly to avoid
// copying the value twice, which could be expensive with a large type for the
// T generic or a large number of values.
//
// Lock Type: Write
//
// Time Complexity: O(n)
func (c *SyncedCircularBuffer[T, U]) PopPntr(val *T) int {
	c.Lock()
	defer c.Unlock()
	return c.CircularBuffer.popSequentialImpl(val, containerTypes.PopAll)
}

// Description: PopSequential will remove the first num occurrences of val in
// the circular buffer. All equality comparisons are performed by the generic U
// widget type that the circular buffer was initialized with. If num is <=0 then
// no values will be poped and the circular buffer will not change.
//
// Time Complexity: O(n)
func (c *CircularBuffer[T, U]) PopSequential(val T, num int) int {
	if num <= 0 {
		return 0
	}
	return c.popSequentialImpl(&val, num)
}

// Description: Places a write lock on the underlying circular buffer and then
// calls the underlying circular buffers [CircularBuffer.PopSequential]
// implementation method. The [CircularBuffer.PopSequential] method is not
// called directly to avoid copying the value twice, which could be expensive
// with a large type for the T generic or a large number of values.
//
// Lock Type: Write
//
// Time Complexity: O(n)
func (c *SyncedCircularBuffer[T, U]) PopSequential(val T, num int) int {
	if num <= 0 {
		return 0
	}
	c.Lock()
	defer c.Unlock()
	return c.CircularBuffer.popSequentialImpl(&val, num)
}

func (c *CircularBuffer[T, U]) popSequentialImpl(val *T, num int) int {
	rv := 0
	curIdx := -1
	w := widgets.Widget[T, U]{}
	for i := 0; i < c.numElems; i++ {
		if w.Eq(&c.vals[c.start.GetProperIndex(i, len(c.vals))], val) && rv < num {
			if rv == 0 {
				curIdx = 0
			}
			w.Zero(&c.vals[c.start.GetProperIndex(i, len(c.vals))])
			rv++
		} else if curIdx != -1 {
			c.vals[c.start.GetProperIndex(curIdx, len(c.vals))] = c.vals[c.start.GetProperIndex(i, len(c.vals))]
			curIdx++
		}
	}
	c.numElems -= rv
	return rv
}

// Description: Clears all values from the circular buffer.
//
// Time Complexity: O(n) (Because of zeroing)
func (c *CircularBuffer[T, U]) Clear() {
	w := widgets.Widget[T, U]{}
	for i := 0; i < c.numElems; i++ {
		w.Zero(&c.vals[c.start.GetProperIndex(i, len(c.vals))])
	}
	c.vals = make([]T, len(c.vals))
	c.numElems = 0
	c.start = 0
}

// Description: Places a write lock on the underlying circular buffer and then
// calls the underlying circular buffers [CircularBuffer.Clear] method.
//
// Lock Type: Write
//
// Time Complexity: O(n)
func (c *SyncedCircularBuffer[T, U]) Clear() {
	c.Lock()
	defer c.Unlock()
	c.CircularBuffer.Clear()
}

// Description: Returns an iterator that iterates over the values in the
// circular buffer.
//
// Time Complexity: O(n)
func (c *CircularBuffer[T, U]) Vals() iter.Iter[T] {
	return iter.SequentialElems[T](c.numElems, c.Get)
}

// Description: Modifies the iterator chain returned by the unerlying
// [CircularBuffer.Vals] method such that a read lock will be placed on the
// underlying circular buffer when the iterator is consumed. The circular buffer
// will have a read lock the entire time the iteration is being performed. The
// lock will not be applied until the iterator starts to be consumed.
//
// Lock Type: Read
//
// Time Complexity: O(n)
func (c *SyncedCircularBuffer[T, U]) Vals() iter.Iter[T] {
	return c.CircularBuffer.Vals().SetupTeardown(
		func() error { c.RLock(); return nil },
		func() error { c.RUnlock(); return nil },
	)
}

// Description: Returns an iterator that iterates over the pointers to the
// values in the circular buffer. The circular buffer will have a read lock the
// entire time the iteration is being performed. The lock will not be applied
// until the iterator is consumed.
//
// Time Complexity: O(n)
func (c *CircularBuffer[T, U]) ValPntrs() iter.Iter[*T] {
	return iter.SequentialElems[*T](c.numElems, c.GetPntr)
}

// Description: Modifies the iterator chain returned by the unerlying
// [CircularBuffer.ValPntrs] method such that a read lock will be placed on the
// underlying circular buffer when the iterator is consumed. The circular buffer
// will have a read lock the entire time the iteration is being performed. The
// lock will not be applied until the iterator starts to be consumed.
//
// Lock Type: Read
//
// Time Complexity: O(n)
func (c *SyncedCircularBuffer[T, U]) ValPntrs() iter.Iter[*T] {
	return c.CircularBuffer.ValPntrs().SetupTeardown(
		func() error { c.RLock(); return nil },
		func() error { c.RUnlock(); return nil },
	)
}

// Description: Returns an iterator that iterates over the keys (indexes) of the
// circular buffer. The circular buffer will have a read lock the entire time
// the iteration is being performed. The lock will not be applied until the
// iterator starts to be consumed.
//
// Time Complexity: O(n)
func (c *CircularBuffer[T, U]) Keys() iter.Iter[int] {
	return iter.Range[int](0, c.numElems, 1)
}

// Description: Modifies the iterator chain returned by the unerlying
// [CircularBuffer.Keys] method such that a read lock will be placed on the
// underlying circular buffer when iterator is consumed. The circular buffer
// will have a read lock the entire time the iteration is being performed. The
// lock will not be applied until the iterator starts to be consumed.
//
// Lock Type: Read
//
// Time Complexity: O(n)
func (c *SyncedCircularBuffer[T, U]) Keys() iter.Iter[int] {
	return c.CircularBuffer.Keys().SetupTeardown(
		func() error { c.RLock(); return nil },
		func() error { c.RUnlock(); return nil },
	)
}

func (c *CircularBuffer[T, U]) inclusiveEnd() wrapingIndex {
	return c.start.Add(c.numElems-1, len(c.vals))
}

func (c *CircularBuffer[T, U]) distanceFromFront(idx int) int {
	return idx
}
func (c *CircularBuffer[T, U]) distanceFromBack(idx int) int {
	return c.numElems - idx
}

// Description: Returns true if all the key value pairs in v are all contained
// in other and the key value pairs in other are all contained in v. Returns
// false otherwise.
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the Get/GetPntr method on other. In big-O it might look something like this,
// O(n*O(other.GetPntr))), where n is the number of elements in c and
// O(other.ContainsPntr) represents the time complexity of the containsPntr
// method on other.
func (c *CircularBuffer[T, U]) KeyedEq(
	other containerTypes.KeyedComparisonsOtherConstraint[int, T],
) bool {
	w := widgets.Widget[T, U]{}
	rv := (c.numElems == other.Length())
	for i := 0; i < c.numElems && rv; i++ {
		if otherV, err := addressableSafeGet[int, T](other, i); err == nil {
			rv = w.Eq(otherV, &(c.vals)[c.start.GetProperIndex(i, len(c.vals))])
		} else {
			rv = false
		}
	}
	return rv
}

// Description: Places a read lock on the underlying circular buffer and then
// calls the underlying circular buffers [CircularBuffer.KeyedEq] method.
// Attempts to place a read lock on other but whether or not that happens is
// implementation dependent.
//
// Lock Type: Read on this circular buffer, read on other
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the GetPntr method on other. In big-O it might look something like this,
// O(n*O(other.GetPntr))), where n is the number of elements in c and
// O(other.ContainsPntr) represents the time complexity of the containsPntr
// method on other.
func (c *SyncedCircularBuffer[T, U]) KeyedEq(
	other containerTypes.KeyedComparisonsOtherConstraint[int, T],
) bool {
	c.RLock()
	other.RLock()
	defer c.RUnlock()
	defer other.RUnlock()
	return c.CircularBuffer.KeyedEq(other)
}

// Description: Returns true if the elements in v are all contained in other and
// the elements of other are all contained in v, regardless of position. Returns
// false otherwise.
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the ContainsPntr method on other. In big-O it might look something like this,
// O(n*O(other.ContainsPntr))), where O(other.ContainsPntr) represents the time
// complexity of the ContainsPntr method on other with m values.
func (c *CircularBuffer[T, U]) UnorderedEq(
	other containerTypes.ComparisonsOtherConstraint[T],
) bool {
	rv := (c.numElems == other.Length())
	for i := 0; i < c.numElems && rv; i++ {
		rv = other.ContainsPntr(&(c.vals)[c.start.GetProperIndex(i, len(c.vals))])
	}
	return rv
}

// Description: Places a read lock on the underlying circular buffer and then
// calls the underlying circular buffer [CircularBuffer.UnorderedEq] method.
// Attempts to place a read lock on other but whether or not that happens is
// implementation dependent.
//
// Lock Type: Read on this circular buffer, read on other
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the ContainsPntr method on other. In big-O it might look something like this,
// O(n*O(other.ContainsPntr))), where O(other.ContainsPntr) represents the time
// complexity of the ContainsPntr method on other with m values.
func (c *SyncedCircularBuffer[T, U]) UnorderedEq(
	other containerTypes.ComparisonsOtherConstraint[T],
) bool {
	c.RLock()
	other.RLock()
	defer c.RUnlock()
	defer other.RUnlock()
	return c.CircularBuffer.UnorderedEq(other)
}

// Description: Populates the circular buffer with the intersection of values from the l
// and r containers. This circular buffer will be cleared before storing the
// result. When clearing, the new resulting circular buffer will be initialized
// with length equivalent to l.Length()+r.Length(). This is necessary to
// guarintee that all values will fit in the resulting statically allocated
// circular buffer.
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the ContainsPntr method on l and r. In big-O it might look something like
// this, O(O(r.ContainsPntr)*O(l.ContainsPntr)), where O(r.ContainsPntr)
// represents the time complexity of the containsPntr method on r and
// O(l.ContainsPntr) represents the time complexity of the containsPntr method
// on l.
func (c *CircularBuffer[T, U]) Intersection(
	l containerTypes.ComparisonsOtherConstraint[T],
	r containerTypes.ComparisonsOtherConstraint[T],
) {
	newC, _ := NewCircularBuffer[T, U](l.Length() + r.Length())
	addressableSafeValIter[T](l).ForEach(
		func(index int, val *T) (iter.IteratorFeedback, error) {
			if r.ContainsPntr(val) {
				newC.Append(*val)
			}
			return iter.Continue, nil
		},
	)
	c.Clear()
	*c = newC
}

// Description: Places a write lock on the underlying circular buffer and then
// calls the underlying circular buffer [CircularBuffer.Intersection] method.
// Attempts to place a read lock on l and r but whether or not that happens is
// implementation dependent.
//
// Lock Type: Write on this circular buffer, read on l and r
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the ContainsPntr method on l and r. In big-O it might look something like
// this, O(O(r.ContainsPntr)*O(l.ContainsPntr)), where O(r.ContainsPntr)
// represents the time complexity of the containsPntr method on r and
// O(l.ContainsPntr) represents the time complexity of the containsPntr method
// on l.
func (c *SyncedCircularBuffer[T, U]) Intersection(
	l containerTypes.ComparisonsOtherConstraint[T],
	r containerTypes.ComparisonsOtherConstraint[T],
) {
	r.RLock()
	l.RLock()
	c.Lock()
	defer r.RUnlock()
	defer l.RUnlock()
	defer c.Unlock()
	c.CircularBuffer.Intersection(l, r)
}

// Description: Populates the circular buffer with the union of values from the
// l and r containers. This circular buffer will be cleared before storing the
// result. When clearing, the new resulting circular buffer will be initialized
// with length equivalent to l.Length()+r.Length(). This is necessary to
// guarintee that all values will fit in the resulting statically allocated
// circular buffer.
//
// Time Complexity: O((n+m)*(n+m)), where n is the number of values in l and m
// is the number of values in r.
func (c *CircularBuffer[T, U]) Union(
	l containerTypes.ComparisonsOtherConstraint[T],
	r containerTypes.ComparisonsOtherConstraint[T],
) {
	maxLen := l.Length() + r.Length()
	newC, _ := NewCircularBuffer[T, U](maxLen)
	oper := func(index int, val *T) (iter.IteratorFeedback, error) {
		newC.AppendUnique(*val)
		return iter.Continue, nil
	}
	addressableSafeValIter[T](l).ForEach(oper)
	addressableSafeValIter[T](r).ForEach(oper)
	c.Clear()
	*c = newC
}

// Description: Places a write lock on the underlying cirular buffer and then
// calls the underlying cirular buffer [CircularBuffer.Union] method. Attempts
// to place a read lock on l and r but whether or not that happens is
// implementation dependent.
//
// Lock Type: Write on this circular buffer, read on l and r
//
// Time Complexity: O((n+m)*(n+m)), where n is the number of values in l and m
// is the number of values in r.
func (c *SyncedCircularBuffer[T, U]) Union(
	l containerTypes.ComparisonsOtherConstraint[T],
	r containerTypes.ComparisonsOtherConstraint[T],
) {
	r.RLock()
	l.RLock()
	c.Lock()
	defer r.RUnlock()
	defer l.RUnlock()
	defer c.Unlock()
	c.CircularBuffer.Union(l, r)
}

// Description: Populates the circular buffer with the result of taking the
// difference of r from l. This circular buffer will be cleared before storing
// the result. When clearing, the new resulting circular buffer will be
// initialized with zero capacity and enough backing memory to store the length
// of l.
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the ContainsPntr method on l and r. In big-O it might look something like
// this, O(O(r.ContainsPntr)*O(l.ContainsPntr)), where O(r.ContainsPntr)
// represents the time complexity of the containsPntr method on r and
// O(l.ContainsPntr) represents the time complexity of the containsPntr method
// on l.
func (c *CircularBuffer[T, U]) Difference(
	l containerTypes.ComparisonsOtherConstraint[T],
	r containerTypes.ComparisonsOtherConstraint[T],
) {
	newC, _ := NewCircularBuffer[T, U](l.Length())
	addressableSafeValIter[T](l).ForEach(
		func(index int, val *T) (iter.IteratorFeedback, error) {
			if !r.ContainsPntr(val) {
				newC.Append(*val)
			}
			return iter.Continue, nil
		},
	)
	c.Clear()
	*c = newC
}

// Description: Places a write lock on the underlying circular buffer and then
// calls the underlying circular buffer [CircularBuffer.Difference] method.
// Attempts to place a read lock on l and r but whether or not that happens is
// implementation dependent.
//
// Lock Type: Write on this circular buffer, read on l and r
//
// Time Complexity: Dependent on the time complexity of the implementation of
// the ContainsPntr method on l and r. In big-O it might look something like
// this, O(O(r.ContainsPntr)*O(l.ContainsPntr)), where O(r.ContainsPntr)
// represents the time complexity of the containsPntr method on r and
// O(l.ContainsPntr) represents the time complexity of the containsPntr method
// on l.
func (c *SyncedCircularBuffer[T, U]) Difference(
	l containerTypes.ComparisonsOtherConstraint[T],
	r containerTypes.ComparisonsOtherConstraint[T],
) {
	r.RLock()
	l.RLock()
	c.Lock()
	defer r.RUnlock()
	defer l.RUnlock()
	defer c.Unlock()
	c.CircularBuffer.Difference(l, r)
}

// Description: Returns true if this circular buffer is a superset to other.
//
// Time Complexity: O(n*m), where n is the number of values in this circular
// buffer and m is the number of values in other.
func (c *CircularBuffer[T, U]) IsSuperset(
	other containerTypes.ComparisonsOtherConstraint[T],
) bool {
	rv := (c.numElems >= other.Length())
	if !rv {
		return false
	}
	addressableSafeValIter[T](other).ForEach(
		func(index int, val *T) (iter.IteratorFeedback, error) {
			if rv = c.ContainsPntr(val); !rv {
				return iter.Break, nil
			}
			return iter.Continue, nil
		},
	)
	return rv
}

// Description: Places a read lock on the underlying circular buffer and then
// calls the underlying circular buffer [CircularBuffer.IsSuperset] method.
// Attempts to place a read lock on other but whether or not that happens is
// implementation dependent.
//
// Lock Type: Read on this circular buffer, read on other
//
// Time Complexity: O(n*m), where n is the number of values in this circular
// buffer and m is the number of values in other.
func (c *SyncedCircularBuffer[T, U]) IsSuperset(
	other containerTypes.ComparisonsOtherConstraint[T],
) bool {
	c.RLock()
	other.RLock()
	defer c.RUnlock()
	defer other.RUnlock()
	return c.CircularBuffer.IsSuperset(other)
}

// Description: Returns true if this circular buffer is a subset to other.
//
// Time Complexity: Dependent on the ContainsPntr method of other. In big-O
// terms it may look somwthing like this: O(n*O(other.ContainsPntr)), where n is
// the number of elements in the current circular buffer and other.ContainsPntr
// represents the time complexity of the containsPntr method on other.
func (c *CircularBuffer[T, U]) IsSubset(
	other containerTypes.ComparisonsOtherConstraint[T],
) bool {
	rv := (c.numElems <= other.Length())
	for i := 0; i < c.numElems && rv; i++ {
		rv = other.ContainsPntr(&c.vals[c.start.GetProperIndex(i, len(c.vals))])
	}
	return rv
}

// Description: Places a read lock on the underlying circular buffer and then
// calls the underlying circular buffer [CircularBuffer.IsSubset] method.
// Attempts to place a read lock on other but whether or not that happens is
// implementation dependent.
//
// Lock Type: Read on this circular buffer, read on other
//
// Time Complexity: Dependent on the ContainsPntr method of other. In big-O
// terms it may look somwthing like this: O(n*O(other.ContainsPntr)), where n is
// the number of elements in the current circular buffer and other.ContainsPntr
// represents the time complexity of the containsPntr method on other.
func (c *SyncedCircularBuffer[T, U]) IsSubset(
	other containerTypes.ComparisonsOtherConstraint[T],
) bool {
	c.RLock()
	other.RLock()
	defer c.RUnlock()
	defer other.RUnlock()
	return c.CircularBuffer.IsSubset(other)
}

// An equality function that implements the [algo.widget.WidgetInterface]
// interface. Internally this is equivalent to [CircularBuffer.KeyedEq]. Returns
// true if l==r, false otherwise.
func (_ *CircularBuffer[T, U]) Eq(
	l *CircularBuffer[T, U],
	r *CircularBuffer[T, U],
) bool {
	return l.KeyedEq(r)
}

// An equality function that implements the [algo.widget.WidgetInterface]
// interface. Internally this is equivalent to [SyncedCircularBuffer.KeyedEq].
// Returns true if l==r, false otherwise.
func (_ *SyncedCircularBuffer[T, U]) Eq(
	l *SyncedCircularBuffer[T, U],
	r *SyncedCircularBuffer[T, U],
) bool {
	return l.KeyedEq(r)
}

// A function that implements the [algo.widget.WidgetInterface] less than
// operation on circular buffers. The l and r vectors will be compared
// lexographically.
func (_ *CircularBuffer[T, U]) Lt(
	l *CircularBuffer[T, U],
	r *CircularBuffer[T, U],
) bool {
	w := widgets.Widget[T, U]{}
	for i := 0; i < min(r.numElems, l.numElems); i++ {
		if w.Lt(
			&l.vals[l.start.GetProperIndex(i, len(l.vals))],
			&r.vals[r.start.GetProperIndex(i, len(r.vals))],
		) {
			return true
		} else if w.Gt(
			&l.vals[l.start.GetProperIndex(i, len(l.vals))],
			&r.vals[r.start.GetProperIndex(i, len(r.vals))],
		) {
			return false
		}
	}
	if l.numElems >= r.numElems {
		return false
	}
	return true
}

// A function that implements the [algo.widget.WidgetInterface] less than
// operation on circular buffers. The l and r vectors will be compared
// lexographically. Read locks are placed on l and r before calling the
// underlying circular buffers [CircularBuffer.Lt] method.
func (_ *SyncedCircularBuffer[T, U]) Lt(
	l *SyncedCircularBuffer[T, U],
	r *SyncedCircularBuffer[T, U],
) bool {
	l.RLock()
	r.RLock()
	defer l.RUnlock()
	defer r.RUnlock()
	return l.CircularBuffer.Lt(&l.CircularBuffer, &r.CircularBuffer)
}

// A function that returns a hash of a circular buffer to implement the
// [algo.widget.WidgetInterface]. To do this all of the individual hashes that
// are produced from the elements of the circular buffer are combined in a way
// that maintains identity, making it so the hash will represent the same
// equality operation that [CircularBuffer.KeyedEq] and [CircularBuffer.Eq]
// provide.
func (_ *CircularBuffer[T, U]) Hash(other *CircularBuffer[T, U]) hash.Hash {
	var rv hash.Hash
	w := widgets.Widget[T, U]{}
	if other.numElems > 0 {
		rv = w.Hash(&other.vals[other.start])
		for i := 1; i < other.numElems; i++ {
			rv = rv.Combine(w.Hash(
				&other.vals[other.start.GetProperIndex(i, len(other.vals))],
			))
		}
	}
	return rv
}

// Places a read lock on the underlying circular buffer of other and then calls
// others underlying circular buffer [CircularBuffer.IsSubset] method.
// Implements the [algo.widget.WidgetInterface].
func (_ *SyncedCircularBuffer[T, U]) Hash(
	other *SyncedCircularBuffer[T, U],
) hash.Hash {
	other.RLock()
	defer other.RUnlock()
	return other.CircularBuffer.Hash(&other.CircularBuffer)
}

// An zero function that implements the [algo.widget.WidgetInterface] interface.
// Internally this is equivalent to [CircularBuffer.Clear].
func (_ *CircularBuffer[T, U]) Zero(other *CircularBuffer[T, U]) {
	other.Clear()
}

// An zero function that implements the [algo.widget.WidgetInterface] interface.
// Internally this is equivalent to [SyncedCircularBuffer.Clear].
func (_ *SyncedCircularBuffer[T, U]) Zero(other *SyncedCircularBuffer[T, U]) {
	other.Clear()
}
