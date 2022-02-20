package ranges

func ForEach[T any](rng Range[T], fn func(T)) {
	var it = rng.Iterate()
	for it.Next() {
		fn(it.Value())
	}
}
