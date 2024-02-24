package view

type NoRead struct {}
type Read[V any] struct {
}
func (r *Read[V])Vals() V { var tmp V; return tmp }
func (r *Read[V])Contains(v V) bool { return true }
func (r *Read[V])ContainsPntr(v *V) bool { return true }

type ReadAddressable[V any] struct {
	Read[V]
}
func (r *ReadAddressable[V])ValPntrs() *V { return nil }

type ReadKeyed[K any, V any] struct {
	Read[V]
}
func (r *ReadKeyed[K, V])Get(idx K) V { var tmp V; return tmp }

type ReadKeyedAddressable[K any, V any] struct {
	Read[V]
	ReadAddressable[V]
}
func (r *ReadKeyedAddressable[K, V])Get(idx K) *V { return nil }

// =============================================================================
type NoCreate struct {}
type Create[V any] struct {
}
func (w *Create[V])Append(vals ...V) error { return nil }

type CreateKeyed[K any, V any] struct {
	Create[V]
}
func (w *CreateKeyed[K, V])Add(k K, v V) error { return nil }

// =============================================================================
type NoUpdate struct {}
type KeyedUpdate[K any, V any] struct {
}
func (u *KeyedUpdate[K,V])Set(idx K, val V) error { return nil }

// =============================================================================
type NoDelete struct {}
type Delete[V any] struct {
}
func (d *Delete[V])Pop(v V, num int) int { return 0 }

// =============================================================================
type Container[K any, V any] interface {
	get(idx V)
}

type vecImpl[V any] []V
type Vector[V any] []V

type Readable[V any] interface { Read[V] | ReadAddressable[V] }
func (v *Vector[V])UnorderedEq(other Readable[V]) {}

// =============================================================================
type View[
	K any,
	V any,
	C Container[K,V],
	READ_ACCESS interface { 
		NoRead | 
		Read[V] | 
		ReadAddressable[V] | 
		ReadKeyed[K,V] | 
		ReadKeyedAddressable[K,V]
	},
	CREATE_ACCESS interface { NoCreate | Create[V] | CreateKeyed[K,V] },
	UPDATE_ACCESS interface { NoUpdate | KeyedUpdate[K,V] },
	DELETE_ACCESS interface { NoDelete | Delete[V] },
] struct {
	Read READ_ACCESS
	Create CREATE_ACCESS
	Update UPDATE_ACCESS
	Delete DELETE_ACCESS
	container C
}
