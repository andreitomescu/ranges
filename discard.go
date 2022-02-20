package ranges

func Discard[T any](rng Range[T]) {
	var it = rng.Iterate()
	for it.Next() {
	}
}
