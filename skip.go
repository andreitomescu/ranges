package ranges

func Skip[T any](rng Range[T], cnt int) Range[T] {
	return &skipRange[T]{rng: rng, cnt: cnt}
}

type skipRange[T any] struct {
	rng Range[T]
	cnt int
}

func (this *skipRange[T]) Iterate() Iterator[T] {
	return &skipIterator[T]{it: this.rng.Iterate(), cnt: this.cnt}
}

type skipIterator[T any] struct {
	it  Iterator[T]
	cnt int
}

func (this *skipIterator[T]) Value() T {
	return this.it.Value()
}

func (this *skipIterator[T]) Next() bool {
	for this.cnt > 0 {
		this.it.Next()
		this.cnt -= 1
	}
	return this.it.Next()
}
