package ranges

func First[T any](rng Range[T]) T {
	var it = rng.Iterate()
	if it.Next() {
		return it.Value()
	}
	panic("no elements")
}

func FirstOr[T any](rng Range[T], value T) T {
	var it = rng.Iterate()
	if it.Next() {
		return it.Value()
	}
	return value
}
