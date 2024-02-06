package basic;

type Pair[T any, U any] struct {
    A T;
    B U;
};

func (p *Pair[T,U])GetA() T {
    return p.A
}

func (p *Pair[T,U])SetA(v T) {
    p.A=v
}

func (p *Pair[T,U])GetB() U {
    return p.B
}

func (p *Pair[T,U])SetB(v U) {
    p.B=v
}
