package ranges

// Determines whether all elements of a range satisfy a condition.
func All[T any](rng Range[T], cond func(T) bool) bool {
	var it = rng.Iterate()
	for it.Next() {
		if cond(it.Value()) {
			continue
		}
		return false
	}
	return true
}
