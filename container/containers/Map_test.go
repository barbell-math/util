package containers

import (
	"github.com/barbell-math/util/algo/widgets"
	"github.com/barbell-math/util/container/dynamicContainers"
)

//go:generate go run interfaceTest.go -type=Map -category=dynamic -interface=Map -genericDecl=[int,int] -factory=generateVector
//go:generate go run interfaceTest.go -type=SyncedMap -category=dynamic -interface=Map -genericDecl=[int,int] -factory=generateSyncedMap

func generateMap() dynamicContainers.Map[int,int] {
    m,_:=NewMap[int,int,widgets.BuiltinInt,widgets.BuiltinInt](0)
    return &m
}

func generateSyncedMap() dynamicContainers.Map[int,int] {
    m,_:=NewSyncedMap[int,int,widgets.BuiltinInt,widgets.BuiltinInt](0)
    return &m
}
