package containers

import "github.com/barbell-math/util/container/containerTypes"

type (
    WidgetConstraint[T any, U any] interface {
        *U
        containerTypes.Widget[T]
    }
)
