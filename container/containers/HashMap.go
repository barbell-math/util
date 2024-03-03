package containers

import (
	"sync"

	"github.com/barbell-math/util/algo/hash"
	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/algo/widgets"
	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/container/containerTypes"
)

type (
    internalHashMapImpl[K any, V any] map[hash.Hash]basic.Pair[K,V]

    // A type to represent a map that dynamically grows as key value pairs are 
    // added. The set will maintain uniqueness and is internally implemented 
    // with a hashing method. The type constraints on the generics define the 
    // logic for how value specific operations, such as equality comparisons, 
    // will be handled.
    HashMap[
        K any,
        V any,
        KI widgets.WidgetInterface[K],
        VI widgets.WidgetInterface[V],
    ] struct {
        internalHashMapImpl[K,V]
    }

    // A synchronized version of HashMap. All operations will be wrapped in the
    // appropriate calls the embedded RWMutex. A pointer to a RWMutex is 
    // embedded rather than a value to avoid copying the lock value.
    SyncedHashMap[
        K any,
        V any,
        KI widgets.WidgetInterface[K],
        VI widgets.WidgetInterface[V],
    ] struct {
        *sync.RWMutex
        HashMap[K,V,KI,VI]
    }
)

// Creates a new map initialized with enough memory to hold size elements. 
// Size must be >= 0, an error will be returned if it is not. If size is 0 the 
// map will be initialized with 0 elements.
func NewHashMap[
    K any, 
    V any, 
    KI widgets.WidgetInterface[K], 
    VI widgets.WidgetInterface[V],
](size int) (HashMap[K,V,KI,VI],error) {
    if size<0 {
        return HashMap[K, V, KI, VI]{},getSizeError(size)
    }
    return HashMap[K, V, KI, VI]{
        internalHashMapImpl: make(internalHashMapImpl[K, V],size),
    },nil
}

// Creates a new synced map initialized with enough memory to hold size 
// elements. Size must be >= 0, an error will be returned if it is not. If size 
// is 0 the map will be initialized with 0 elements. The underlying RWMutex 
// value will be fully unlocked upon initialization.
func NewSyncedHashMap[
    K any, 
    V any, 
    KI widgets.WidgetInterface[K],
    VI widgets.WidgetInterface[V],
](size int) (SyncedHashMap[K,V,KI,VI],error) {
    rv,err:=NewHashMap[K,V,KI,VI](size)
    return SyncedHashMap[K, V, KI, VI]{
        RWMutex: &sync.RWMutex{},
        HashMap: rv,
    },err
}

// Converts the supplied map to a syncronized map. Beware: The original 
// non-synced map will remain useable.
func (m *HashMap[K, V, KI, VI])ToSynced() SyncedHashMap[K,V,KI,VI] {
    return SyncedHashMap[K, V, KI, VI]{
        RWMutex: &sync.RWMutex{},
        HashMap: *m,
    }
}

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (m *HashMap[K, V, KI, VI])Lock() { }

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (m *HashMap[K, V, KI, VI])Unlock() { }

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (m *HashMap[K, V, KI, VI])RLock() { }

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (m *HashMap[K, V, KI, VI])RUnlock() { }

// The SyncedHashMap method to override the HashMap pass through function and 
// actually apply the mutex operation.
func (m *SyncedHashMap[K, V, KI, VI])Lock() { m.RWMutex.Lock() }

// The SyncedHashMap method to override the HashMap pass through function and 
// actually apply the mutex operation.
func (m *SyncedHashMap[K, V, KI, VI])Unlock() { m.RWMutex.Unlock() }

// The SyncedHashMap method to override the HashMap pass through function and 
// actually apply the mutex operation.
func (m *SyncedHashMap[K, V, KI, VI])RLock() { m.RWMutex.RLock() }

// The SyncedHashMap method to override the HashMap pass through function and 
// actually apply the mutex operation.
func (m *SyncedHashMap[K, V, KI, VI])RUnlock() { m.RWMutex.RUnlock() }

// Returns false, maps are addressable.
func (m *HashMap[K, V, KI, VI])IsAddressable() bool { return false }

// Returns false, a map is not synced.
func (m *HashMap[K, V, KI, VI])IsSynced() bool { return false }

// Returns true, a synced map is synced.
func (m *SyncedHashMap[K, V, KI, VI])IsSynced() bool { return true }

// Description: Returns the number of elements in the hash map.
//
// Time Complexity: O(1)
func (m *HashMap[K, V, KI, VI])Length() int {
    return len(m.internalHashMapImpl)
}
// Description: Places a read lock on the underlying hash map and then calls the 
// underlying hash map [HashMap.Length] method.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (m *SyncedHashMap[K, V, KI, VI])Length() int {
    m.RLock()
    defer m.RUnlock()
    return m.HashMap.Length()
}

// Description: Contains will return true if the supplied value is in the 
// map, false otherwise. All equality comparisons are performed by the 
// generic VI widget type that the hash map was initialized with. 
//
// Time Complexity: O(n) (linear search)
func (m *HashMap[K, V, KI, VI])Contains(v V) bool {
    return m.ContainsPntr(&v)
}
// Description: Places a read lock on the underlying map and then calls the 
// underlying map [Map.ContainsPntr] method.
//
// Lock Type: Read
//
// Time Complexity: O(n) (linear search)
func (m *SyncedHashMap[K, V, KI, VI])Contains(v V) bool {
    m.RLock()
    defer m.RUnlock()
    return m.ContainsPntr(&v)
}

// Description: ContainsPntr will return true if the supplied value is in the 
// map, false otherwise. All equality comparisons are performed by the 
// generic VI widget type that the map was initialized with.
//
// Time Complexity: O(n) (linear search)
func (m *HashMap[K, V, KI, VI])ContainsPntr(v *V) bool {
    w:=widgets.NewWidget[V,VI]()
    for _,iterV:=range(m.internalHashMapImpl) {
        if w.Eq(v,&iterV.B) {
            return true
        }
    }
    return false
}
// Description: Places a read lock on the underlying map and then calls the 
// underlying maps [Map.ContainsPntr] method.
//
// Lock Type: Read
//
// Time Complexity: O(n) (linear search)
func (m *SyncedHashMap[K, V, KI, VI])ContainsPntr(v *V) bool {
    m.RLock()
    defer m.RUnlock()
    return m.HashMap.ContainsPntr(v)
}

// Description: Gets the value at the specified key. Returns a 
// [containerTypes.KeyError] if the key is not found in the hash map.
//
// Time Complexity: O(1)
func (m *HashMap[K, V, KI, VI])Get(k K) (V,error) {
    if h,ok:=m.getHashPosition(&k); ok {
        return m.internalHashMapImpl[h].B,nil
    }
    var tmp V
    return tmp,getKeyError[K](&k)
}
// Description: Places a read lock on the underlying hash map and then gets the 
// value at the specified key. Exhibits the same behavior as the [HashMap.Get]
// method. The underlying [HashMap.Get] method is not called to avoid copying 
// the return value twice, which could be inefficient with a large value for the 
// K generic.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (m *SyncedHashMap[K, V, KI, VI])Get(k K) (V,error) {
    m.RLock()
    defer m.RUnlock()
    if h,ok:=m.getHashPosition(&k); ok {
        return m.internalHashMapImpl[h].B,nil
    }
    var tmp V
    return tmp,getKeyError[K](&k)
}

func (h *HashMap[K, V, KI, VI])getHashPosition(k *K) (hash.Hash,bool) {
    w:=widgets.NewWidget[K,KI]()
    for i:=w.Hash(k); ; i++ {
        if iterV,found:=h.internalHashMapImpl[i]; found && w.Eq(k,&iterV.A) {
            return i,true 
        } else if !found {
            return hash.Hash(0),false
        }
    }
}

// Panics, hash maps are not addressable.
func (m *HashMap[K, V, KI, VI])GetPntr(k K) (*V,error) {
    panic(getNonAddressablePanicText("hash map"))
}

// Description: KeyOf will return the key of the first occurrence of the 
// supplied value in the map. If the value is not found then the returned 
// key will be a zero initialized key value and the boolean flag will be set to 
// false. If the value is found then the boolean flag will be set to true. All 
// equality comparisons are performed by the generic VI widget type that the map 
// was initialized with.
//
// Time Complexity: O(n) (linear search)
func (m *HashMap[K, V, KI, VI])KeyOf(v V) (K,bool) {
    var tmp K
    return tmp,m.keyOfImpl(&tmp,&v)
}
// Description: Places a read lock on the underlying hash map and then calls the 
// underlying hash map [HashMap.KeyOf] implemenation method. The [HashMap.KeyOf]
// method is not called directly to avoid copying the val variable twice, which 
// could be expensive with a large type for the V generic.
//
// Lock Type: Read
//
// Time Complexity: O(n) (linear search)
func (m *SyncedHashMap[K, V, KI, VI])KeyOf(v V) (K,bool) {
    m.RLock()
    defer m.RUnlock()
    var tmp K
    return tmp,m.keyOfImpl(&tmp,&v)
}

func (m *HashMap[K, V, KI, VI])keyOfImpl(k *K, v *V) bool {
    w:=widgets.NewWidget[V,VI]()
    for _,iterV:=range(m.internalHashMapImpl) {
        if w.Eq(v,&iterV.B) {
            *k=iterV.A
            return true
        }
    }
    return false
}

// Description: Sets the values at the specified keys. Returns an error if the 
// key is not in the map. Stops setting values as soon as an error is 
// encountered.
//
// Time Complexity: O(m), where m=len(vals)
func (m *HashMap[K, V, KI, VI])Set(kvPairs ...basic.Pair[K,V]) error {
    return m.setImpl(kvPairs)
}
// Description: Places a write lock on the underlying map and then calls the 
// underlying hash map [HashMap.Set] implementaiton method. The [HashMap.Set] 
// method is not called directly to avoid copying the vals varargs twice, which 
// could be expensive with a large types for the K or V generics or a large 
// number of values.
//
// Lock Type: Write
//
// Time Complexity: O(m), where m=len(vals)
func (m *SyncedHashMap[K, V, KI, VI])Set(kvPairs ...basic.Pair[K,V]) error {
    m.Lock()
    defer m.Unlock()
    return m.setImpl(kvPairs)
}

func (m *HashMap[K, V, KI, VI])setImpl(kvPairs []basic.Pair[K,V]) error {
    for i:=0; i<len(kvPairs); i++ {
        if h,found:=m.getHashPosition(&kvPairs[i].A); found {
            m.zeroKVPair(h)
            m.internalHashMapImpl[h]=kvPairs[i]
        } else {
            return getKeyError[K](&kvPairs[i].A)
        }
    }
    return nil
}

// Description: Emplace will insert the supplied values into the hash map if 
// they do not exist and will set they keys value if it already exists in the
// hash map. The values will be inserted in the order that they are given. 
//
// Time Complexity: O(m), where m=len(vals)
func (m *HashMap[K, V, KI, VI])Emplace(vals ...basic.Pair[K,V]) error {
    for i:=0; i<len(vals); i++ {
        m.emplaceImpl(&vals[i])
    }
    return nil
}
// Description: Places a write lock on the underlying hash map and then calls 
// the underlying hash maps [HashMap.Insert] implementation method. The 
// [HashMap.Insert] method is not called directly to avoid copying the vals 
// varargs twice, which could be expensive with a large types for the K or V 
// generics or a large number of values.
//
// Lock Type: Write
//
// Time Complexity: O(n*m), where m=len(vals)
func (m *SyncedHashMap[K, V, KI, VI])Emplace(vals ...basic.Pair[K,V]) error {
    m.Lock()
    defer m.Unlock()
    for i:=0; i<len(vals); i++ {
        m.emplaceImpl(&vals[i])
    }
    return nil
}

func (m *HashMap[K, V, KI, VI])emplaceImpl(v *basic.Pair[K,V]) {
    kw:=widgets.NewWidget[K,KI]()
    for i:=kw.Hash(&v.A); ; i++ {
        if iterV,found:=m.internalHashMapImpl[i]; !found {
            // Reached end of collision chain, insert
            m.internalHashMapImpl[i]=*v
            break
        } else if kw.Eq(&v.A,&iterV.A) {
            // Found value in collision chain, set value
            m.zeroKVPair(i)
            m.internalHashMapImpl[i]=*v
            break
        }
    }
}

// Description: Pop will remove the first num occurrences of val in the map. 
// All equality comparisons are performed by the generic VI widget type that the 
// map was initialized with. If num is <=0 then no values will be poped and 
// the map will not change.
//
// Time Complexity: O(n)
func (m *HashMap[K, V, KI, VI])Pop(v V) int {
    return m.popImpl(&v)
}
// Description: Places a write lock on the underlying map and then calls the 
// underlying map [Map.Pop] implementation method. The [Map.Pop] method is not 
// called directly to avoid copying the value twice, which could be expensive 
// with a large type for the V generic or a large number of values.
//
// Lock Type: Write
//
// Time Complexity: O(n)
func (m *SyncedHashMap[K, V, KI, VI])Pop(v V) int {
    m.Lock()
    defer m.Unlock()
    return m.popImpl(&v)
}

func (m *HashMap[K, V, KI, VI])popImpl(v *V) int {
    rv:=0
    vw:=widgets.NewWidget[V,VI]()
    for iterH,iterV:=range(m.internalHashMapImpl) {
        if vw.Eq(&iterV.B,v) {
            m.removeMultipleValues(iterH,v)
            rv++
        }
    }
    return rv
}

func (m *HashMap[K, V, KI, VI])removeMultipleValues(h hash.Hash, v *V) {
    kw:=widgets.NewWidget[K,KI]()
    vw:=widgets.NewWidget[V,VI]()
    m.zeroKVPair(h)
    delete(m.internalHashMapImpl,h)
    curPos:=h
    for j:=h+1; ; j++ {
        if iterV,found:=m.internalHashMapImpl[j]; found {
            // If the hash of iterV is > the curPos then do not move it, it 
            // is already in the correct position, but continue on this 
            // collision chain to check for other values that could move.
            if kw.Hash(&iterV.A)>curPos {
                continue
            }
            // If the value is eq to the value we are deleting then also skip it
            // because deleting it and then re-creating it does not guarintee
            // it will be iterated on later.
            // Scenario 1: The value is the last value of the collision chain
            //   The value will be orphaned an unaccessable using the standard
            //   Get method but will be deleted by a future iteration of the
            //   popImpl function. If it were moved this 're-teration' would not
            //   be guarinteed.
            // Scenario 2: The value is in the middle of a collision chain. in
            //   this case the hole that is created will be filled by an 
            //   appripriate value later in the collision chain.
            if vw.Eq(v,&iterV.B) {
                continue
            }
            m.internalHashMapImpl[curPos]=m.internalHashMapImpl[j]
            m.zeroKVPair(h)
            delete(m.internalHashMapImpl,j)
            curPos=j 
        } else {
            break
        }
    }
}

func (m *HashMap[K, V, KI, VI])zeroKVPair(h hash.Hash) {
    kw:=widgets.NewWidget[K,KI]()
    vw:=widgets.NewWidget[V,VI]()
    if v,ok:=m.internalHashMapImpl[h]; ok {
        kw.Zero(&v.A)
        vw.Zero(&v.B)
    }
}

// Description: Deletes the key value pair that has the specified key. Returns 
// an error if the key is not found in the hash map.
//
// Time Complexity: O(1)
func (m *HashMap[K, V, KI, VI])Delete(k K) error {
    return m.deleteImpl(&k)
}
// Description: Places a write lock on the underlying hash map and then calls 
// the underlying hash maps [HashMap.Delete] method.
//
// Lock Type: Write
//
// Time Complexity: O(1)
func (m *SyncedHashMap[K, V, KI, VI])Delete(k K) error {
    m.Lock()
    defer m.Unlock()
    return m.HashMap.deleteImpl(&k)
}

func (m *HashMap[K, V, KI, VI])deleteImpl(k *K) error {
    if h,found:=m.getHashPosition(k); found {
        m.removeSingleValue(h)
        return nil
    }
    return getKeyError[K](k)
}

func (m *HashMap[K, V, KI, VI])removeSingleValue(h hash.Hash) {
    kw:=widgets.NewWidget[K,KI]()
    m.zeroKVPair(h)
    delete(m.internalHashMapImpl,h)
    curPos:=h
    for j:=h+1; ; j++ {
        if iterV,found:=m.internalHashMapImpl[j]; found {
            // If the hash of iterV is > the curPos then do not move it, it 
            // is already in the correct position, but continue on this 
            // collision chain to check for other values that could move.
            if kw.Hash(&iterV.A)>curPos {
                continue
            }
            m.internalHashMapImpl[curPos]=m.internalHashMapImpl[j]
            m.zeroKVPair(h)
            delete(m.internalHashMapImpl,j)
            curPos=j 
        } else {
            break
        }
    }
}

// Description: Clears all values from the map. Equivalent to making a new hash 
// map and setting it equal to the current one.
//
// Time Complexity: O(1)
func (m *HashMap[K, V, KI, VI])Clear() {
    kw:=widgets.NewWidget[K,KI]()
    vw:=widgets.NewWidget[V,VI]()
    for _,v:=range(m.internalHashMapImpl) {
        kw.Zero(&v.A)
        vw.Zero(&v.B)
    }
    m.internalHashMapImpl=internalHashMapImpl[K,V]{}
}
// Description: Places a write lock on the underlying hash map and then calls 
// the underlying hash maps [HashMap.Clear] method.
//
// Lock Type: Write
//
// Time Complexity: O(1)
func (m *SyncedHashMap[K, V, KI, VI])Clear() {
    m.Lock()
    defer m.Unlock()
    m.HashMap.Clear()
}

// Description: Returns an iterator that iterates over the keys of the hash map.
// The hash map will have a read lock the entire time the iteration is being 
// performed. The lock will not be applied until the iterator starts to be 
// consumed.
//
// Time Complexity: O(n)
func (m *HashMap[K, V, KI, VI])Keys() iter.Iter[K] {
    return iter.Map[basic.Pair[K,V],K](
        iter.MapVals[hash.Hash,basic.Pair[K,V]](m.internalHashMapImpl),
        func(index int, val basic.Pair[K, V]) (K, error) { return val.A,nil },
    )
}
// Description: Modifies the iterator chain returned by the unerlying 
// [HashMap.Keys] method such that a read lock will be placed on the underlying 
// hash map when the iterator is consumed. The hash map will have a read lock the 
// entire time the iteration is being performed. The lock will not be applied 
// until the iterator starts to be consumed.
//
// Lock Type: Read
//
// Time Complexity: O(n)
func (m *SyncedHashMap[K, V, KI, VI])Keys() iter.Iter[K] {
    return m.HashMap.Keys().SetupTeardown(
        func() error { m.RLock(); return nil },
        func() error { m.RUnlock(); return nil },
    )
}

// Description: Returns an iterator that iterates over the values in the hash 
// map.
//
// Time Complexity: O(n)
func (m *HashMap[K, V, KI, VI])Vals() iter.Iter[V] {
    return iter.Map[basic.Pair[K,V],V](
        iter.MapVals[hash.Hash,basic.Pair[K,V]](m.internalHashMapImpl),
        func(index int, val basic.Pair[K, V]) (V, error) { return val.B,nil },
    )
}
// Description: Modifies the iterator chain returned by the unerlying 
// [HashMap.Vals] method such that a read lock will be placed on the underlying 
// hash map when the iterator is consumed. The hash map will have a read lock 
// the entire time the iteration is being performed. The lock will not be 
// applied until the iterator starts to be consumed.
//
// Lock Type: Read
//
// Time Complexity: O(n)
func (m *SyncedHashMap[K, V, KI, VI])Vals() iter.Iter[V] {
    return m.HashMap.Vals().SetupTeardown(
        func() error { m.RLock(); return nil },
        func() error { m.RUnlock(); return nil },
    )
}

// Panics, a hash set is not addressable.
func (m *HashMap[K, V, KI, VI])ValPntrs() iter.Iter[*V] {
    panic(getNonAddressablePanicText("hash map"))
}

// Description: Returns true if all the key value pairs in v are all contained 
// in other and the key value pairs in other are all contained in v. Returns 
// false otherwise. 
//
// Time Complexity: Dependent on the time complexity of the implementation of 
// the Get/GetPntr method on other. In big-O it might look something like this, 
// O(n*O(other.GetPntr))), where n is the number of elements in v and 
// O(other.ContainsPntr) represents the time complexity of the containsPntr 
// method on other.
func (m *HashMap[K, V, KI, VI])KeyedEq(
    other containerTypes.KeyedComparisonsOtherConstraint[K,V],
) bool {
    vw:=widgets.NewWidget[V,VI]()
    if len(m.internalHashMapImpl)!=other.Length() {
        return false
    }
    for _,v:=range(m.internalHashMapImpl) {
	if otherV,err:=addressableSafeGet[K,V](other,v.A); err==nil {
            if !vw.Eq(&v.B,otherV) {
                return false
            }
	} else {
	    return false
	}
    }
    return true
}
// Description: Places a read lock on the underlying hash map and then calls the 
// underlying hash map [HashMap.KeyedEq] method. Attempts to place a read lock 
// on other but whether or not that happens is implementation dependent.
//
// Lock Type: Read on this hash map, read on other
//
// Time Complexity: Dependent on the time complexity of the implementation of 
// the GetPntr method on other. In big-O it might look something like this, 
// O(n*O(other.GetPntr))), where n is the number of elements in v and 
// O(other.ContainsPntr) represents the time complexity of the containsPntr 
// method on other.
func (m *SyncedHashMap[K, V, KI, VI])KeyedEq(
    other containerTypes.KeyedComparisonsOtherConstraint[K,V],
) bool {
    m.RLock()
    other.RLock()
    defer m.RUnlock()
    defer other.RUnlock()
    return m.HashMap.KeyedEq(other)
}

// An equality function that implements the [algo.widget.WidgetInterface] 
// interface. Internally this is equivalent to [HashMap.KeyedEq]. Returns true
// if l==r, false otherwise.
func (m *HashMap[K, V, KI, VI])Eq(
    l *HashMap[K,V,KI,VI],
    r *HashMap[K,V,KI,VI],
) bool {
    return l.KeyedEq(r)
}
// An equality function that implements the [algo.widget.WidgetInterface] 
// interface. Internally this is equivalent to [SyncedHashMap.KeyedEq]. Returns
// true if l==r, false otherwise.
func (m *SyncedHashMap[K, V, KI, VI])Eq(
    l *SyncedHashMap[K,V,KI,VI], 
    r *SyncedHashMap[K,V,KI,VI],
) bool {
    return l.KeyedEq(r)
}

// Panics, hash maps cannot be compared for order.
func (m *HashMap[K, V, KI, VI])Lt(
    l *HashMap[K,V,KI,VI], 
    r *HashMap[K,V,KI,VI],
) bool {
    panic("Hash maps cannot be compared relative to each other.")
}
// Panics, hash maps cannot be compared for order.
func (m *SyncedHashMap[K, V, KI, VI])Lt(
    l *SyncedHashMap[K,V,KI,VI], 
    r *SyncedHashMap[K,V,KI,VI],
) bool {
    panic("Hash maps cannot be compared relative to each other.")
}

// A function that returns a hash of a hash map. To do this all of the 
// individual hashes that are produced from the elements of the hash map are 
// combined in a way that maintains identity, making it so the hash will 
// represent the same equality operation that [HashMap.KeyedEq] and 
// [HashMap.Eq] provide.
func (m *HashMap[K, V, KI, VI])Hash(other *HashMap[K,V,KI,VI]) hash.Hash {
    cntr:=0
    var rv hash.Hash
    kw:=widgets.NewWidget[K,KI]()
    vw:=widgets.NewWidget[V,VI]()
    for _,iterV:=range(m.internalHashMapImpl) {
        iterH:=kw.Hash(&iterV.A).Combine(vw.Hash(&iterV.B))
        if cntr==0 {
            rv=iterH
            cntr++
        } else {
            rv=rv.CombineUnordered(kw.Hash(&iterV.A).Combine(vw.Hash(&iterV.B)))
        }
    }
    return rv
}
// Places a read lock on the underlying hash map of other and then calls others
// underlying hash maps [HashMap.IsSubset] method.
func (m *SyncedHashMap[K, V, KI, VI])Hash(
    other *SyncedHashMap[K,V,KI,VI],
) hash.Hash {
    other.RLock()
    defer other.RUnlock()
    return m.HashMap.Hash(&other.HashMap)
}

// An zero function that implements the [algo.widget.WidgetInterface] interface.
// Internally this is equivalent to [HashMap.Clear].
func (m *HashMap[K, V, KI, VI])Zero(other *HashMap[K,V,KI,VI]) {
    other.Clear()
}
// An zero function that implements the [algo.widget.WidgetInterface] interface.
// Internally this is equivalent to [SyncedHashMap.Clear].
func (m *SyncedHashMap[K, V, KI, VI])Zero(other *SyncedHashMap[K,V,KI,VI]) {
    other.Clear()
}
