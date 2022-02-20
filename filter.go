package ranges

func Filter[T any](rng Range[T], fn func(T) bool) Range[T] {
	return &filterRange[T]{rng: rng, fn: fn}
}

type filterRange[T any] struct {
	rng Range[T]
	fn  func(T) bool
}

func (this *filterRange[T]) Iterate() Iterator[T] {
	return &filterIterator[T]{it: this.rng.Iterate(), fn: this.fn}
}

type filterIterator[T any] struct {
	it Iterator[T]
	fn func(T) bool
}

func (this *filterIterator[T]) Value() T {
	return this.it.Value()
}

func (this *filterIterator[T]) Next() bool {
	for this.it.Next() {
		if this.fn(this.Value()) {
			return true
		}
	}
	return false
}
