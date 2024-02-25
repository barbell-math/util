package containers

import (
	"github.com/barbell-math/util/algo/widgets"
)

//go:generate go run interfaceTest.go -type=HashMap -category=dynamic -interface=Map -genericDecl=[int,int] -factory=generateHashMap
//go:generate go run interfaceTest.go -type=SyncedHashMap -category=dynamic -interface=Map -genericDecl=[int,int] -factory=generateSyncedHashMap

func generateHashMap() HashMap[int,int,widgets.BuiltinInt,widgets.BuiltinInt] {
    m,_:=NewHashMap[int,int,widgets.BuiltinInt,widgets.BuiltinInt](0)
    return m
}

func generateSyncedHashMap() SyncedHashMap[
    int,
    int,
    widgets.BuiltinInt,
    widgets.BuiltinInt,
] {
    m,_:=NewSyncedHashMap[int,int,widgets.BuiltinInt,widgets.BuiltinInt](0)
    return m
}
