package containers

//go:generate go run interfaceTest.go -type=HookedHashSet -category=dynamic -interface=Set -genericDecl=[int] -factory=generateHookedHashSet
//go:generate go run interfaceTest.go -type=SyncedHookedHashSet -category=dynamic -interface=Set -genericDecl=[int] -factory=generateSyncedHookedHashSet

type testSetHooks struct {
	addOpCntr int
	delOpCntr int
	clrOpCntr int
}

func (t *testSetHooks)addOp(hashLoc HashSetHash) {
	t.addOpCntr++
}
func (t *testSetHooks)deleteOp(updatedHashes map[OldHashSetHash]NewHashSetHash) {
	t.delOpCntr++
}
func (t *testSetHooks)clearOp() {
	t.clrOpCntr++
}

func generateHookedHashSet(capacity int) HookedHashSet[int, badBuiltinInt] {
	v, _:=NewHookedHashSet[int, badBuiltinInt](&testSetHooks{},capacity)
	return v
}

func generateSyncedHookedHashSet(
	capacity int,
) SyncedHookedHashSet[int, badBuiltinInt] {
	v, _:=NewSyncedHookedHashSet[int, badBuiltinInt](&testSetHooks{},capacity)
	return v
}

// TODO - test that the hooks are being properly called
