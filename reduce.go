package ranges

func Reduce[T, R any](rng Range[T], seed R, fn func(R, T) R) R {
	var it = rng.Iterate()
	for it.Next() {
		seed = fn(seed, it.Value())
	}
	return seed
}
