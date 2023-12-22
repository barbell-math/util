package numeric

import (
    "fmt"
    "github.com/barbell-math/util/math"
)

type Vars[N math.Number] map[string]N;


func (v Vars[N])Access(_var string) (N,error) {
    if v,ok:=v[_var]; ok {
        return v,nil;
    }
    return N(0),math.MissingVariable(
        fmt.Sprintf("Requested: %s Have: %v",_var,v),
    );
}

func (v Vars[N])Copy() Vars[N] {
    rv:=Vars[N](make(map[string]N,len(v)));
    for k,v:=range(v) {
        rv[k]=v;
    }
    return rv;
}

func (v Vars[N])Apply(other Vars[N], op func(accum *N, iter N) error) (Vars[N],error) {
    for oKey,oVal:=range(other) {
        if val,ok:=v[oKey]; ok {
            if err:=op(&val,oVal); err==nil {
                v[oKey]=val;
            } else {
                return v,err;
            }
        }
    }
    return v,nil;
}

func (v Vars[N])ApplyConst(_const N, op func(accum *N, iter N) error) (Vars[N],error) {
    for k,val:=range(v) {
        if err:=op(&val,_const); err==nil {
            v[k]=val;
        } else {
            return v,err;
        }
    }
    return v,nil;
}
