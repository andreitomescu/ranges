package ranges

func Take[T any](rng Range[T], cnt int) Range[T] {
	return &takeRange[T]{rng: rng, cnt: cnt}
}

type takeRange[T any] struct {
	rng Range[T]
	cnt int
}

func (this *takeRange[T]) Iterate() Iterator[T] {
	return &takeIterator[T]{it: this.rng.Iterate(), cnt: this.cnt}
}

type takeIterator[T any] struct {
	it  Iterator[T]
	cnt int
}

func (this *takeIterator[T]) Value() T {
	return this.it.Value()
}

func (this *takeIterator[T]) Next() bool {
	if this.cnt > 0 {
		this.cnt -= 1
		return this.it.Next()
	}
	return false
}
