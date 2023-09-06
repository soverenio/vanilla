package defcon

func NewTuple2[T1 any, T2 any](first T1, second T2) Tuple2[T1, T2] {
	return Tuple2[T1, T2]{first: first, second: second}
}

type Tuple2[T1 any, T2 any] struct {
	first  T1
	second T2
}

func (t *Tuple2[T1, T2]) First() T1 {
	return t.first
}

func (t *Tuple2[T1, T2]) Second() T2 {
	return t.second
}

func (t *Tuple2[T1, T2]) Values() (T1, T2) {
	return t.first, t.second
}

func NewTuple3[T1 any, T2 any, T3 any](first T1, second T2, third T3) Tuple3[T1, T2, T3] {
	return Tuple3[T1, T2, T3]{first: first, second: second, third: third}
}

type Tuple3[T1 any, T2 any, T3 any] struct {
	first  T1
	second T2
	third  T3
}

func (t *Tuple3[T1, T2, T3]) First() T1 {
	return t.first
}

func (t *Tuple3[T1, T2, T3]) Second() T2 {
	return t.second
}

func (t *Tuple3[T1, T2, T3]) Third() T3 {
	return t.third
}

func (t *Tuple3[T1, T2, T3]) Values() (T1, T2, T3) {
	return t.first, t.second, t.third
}
