package ranges

func Map[T, R any](rng Range[T], fn func(T) R) Range[R] {
	return &mapRange[T, R]{rng: rng, fn: fn}
}

type mapRange[T, R any] struct {
	rng Range[T]
	fn  func(T) R
}

func (this *mapRange[T, R]) Iterate() Iterator[R] {
	return &mapIterator[T, R]{it: this.rng.Iterate(), fn: this.fn}
}

type mapIterator[T, R any] struct {
	it Iterator[T]
	fn func(T) R
}

func (this *mapIterator[T, R]) Value() R {
	return this.fn(this.it.Value())
}

func (this *mapIterator[T, R]) Next() bool {
	return this.it.Next()
}
