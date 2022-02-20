package ranges

// Determines whether any element of a range satisfies a condition.
func Any[T any](rng Range[T], cond func(T) bool) bool {
	var it = rng.Iterate()
	for it.Next() {
		if cond(it.Value()) {
			return true
		}
	}
	return false
}
