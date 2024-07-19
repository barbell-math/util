package containers

import (
	"github.com/barbell-math/util/algo/widgets"
)


type (
	HashSetHooks interface {
		addOp(hashLoc HashSetHash)
		deleteOp(updatedHashes map[OldHashSetHash]NewHashSetHash)
		clearOp()
	}

	HookedHashSet[T any, U widgets.WidgetInterface[T]] struct {
		HashSet[T,U]
		hooks HashSetHooks
	}

	SyncedHookedHashSet[T any, U widgets.WidgetInterface[T]] struct {
		SyncedHashSet[T,U]
		hooks HashSetHooks
	}
)

func NewHookedHashSet[T any, U widgets.WidgetInterface[T]](
	hooks HashSetHooks,
	size int,
) (HookedHashSet[T,U],error) {
	hs,err:=NewHashSet[T,U](size)
	if err!=nil {
		return HookedHashSet[T, U]{}, err
	}
	return HookedHashSet[T, U]{
		HashSet: hs,
		hooks: hooks,
	}, nil
}

func (h *HookedHashSet[T, U])GetFromHash(internalHash HashSetHash) (T,error) {
	if v,ok:=h.HashSet.internalHashSetImpl[internalHash]; ok {
		return v,nil
	}
	var tmp T
	return tmp, getKeyError[HashSetHash](&internalHash)
}

func (h *HookedHashSet[T, U])AppendUnique(vals ...T) error {
	for _,v:=range(vals) {
		h.AppendUnique(v)
		if vHash,ok:=h.getHashPosition(&v); ok {
			h.hooks.addOp(vHash)
		}
	}
	// The equivalent method on a normal hash set will never return an error.
	return nil
}

func (h *SyncedHookedHashSet[T, U])AppendUnique(vals ...T) error {
	h.Lock()
	defer h.Unlock()
	for _,v:=range(vals) {
		h.SyncedHashSet.HashSet.AppendUnique(v)
		if vHash,ok:=h.getHashPosition(&v); ok {
			h.hooks.addOp(vHash)
		}
	}
	// The equivalent method on a normal hash set will never return an error.
	return nil
}

func (h *HookedHashSet[T, U])Pop(v T) int {
	if res,ok:=h.getHashesAffectedByPop(&v); ok {
		h.hooks.deleteOp(res)
	}
	return h.HashSet.Pop(v)
}

func (h *SyncedHookedHashSet[T, U])Pop(v T) int {
	h.Lock()
	defer h.Unlock()
	if res,ok:=h.getHashesAffectedByPop(&v); ok {
		h.hooks.deleteOp(res)
	}
	return h.SyncedHashSet.HashSet.Pop(v)
}

func (h *HookedHashSet[T, U])PopPntr(v *T) int {
	if res,ok:=h.getHashesAffectedByPop(v); ok {
		h.hooks.deleteOp(res)
	}
	return h.HashSet.PopPntr(v)
}

func (h *SyncedHookedHashSet[T, U])PopPntr(v *T) int {
	h.Lock()
	defer h.Unlock()
	if res,ok:=h.getHashesAffectedByPop(v); ok {
		h.hooks.deleteOp(res)
	}
	return h.SyncedHashSet.HashSet.PopPntr(v)
}

func (h *HookedHashSet[T, U])Clear() {
	h.hooks.clearOp()
	h.HashSet.Clear()
}

func (h *SyncedHookedHashSet[T, U])Clear() {
	h.Lock()
	defer h.Unlock()
	h.hooks.clearOp()
	h.SyncedHashSet.HashSet.Clear()
}
