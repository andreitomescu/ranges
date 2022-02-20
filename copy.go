package ranges

func Copy[T any](rng Range[T]) Range[T] {
	return Slice(ToSlice(rng))
}
