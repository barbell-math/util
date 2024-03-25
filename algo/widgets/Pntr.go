package widgets

import "github.com/barbell-math/util/algo/hash"

type (
    Pntr[T any, I WidgetInterface[T]] struct {
        iFace I
    }
)

func (p Pntr[T, I])Eq(l **T, r **T) bool {
    return p.iFace.Eq(*l,*r)
}

func (p Pntr[T, I])Lt(l **T, r **T) bool {
    return p.iFace.Lt(*l,*r)
}

func (p Pntr[T, I])Hash(other **T) hash.Hash {
    return p.iFace.Hash(*other)
}

func (p Pntr[T, I])Zero(other **T) {
    p.iFace.Zero(*other)
}
