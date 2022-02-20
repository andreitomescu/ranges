package ranges

// Determines whether a range contains a specified element.
func Contains[T any](rng Range[T], val T, same func(T, T) bool) bool {
	var it = rng.Iterate()
	for it.Next() {
		if same(it.Value(), val) {
			return true
		}
	}
	return false
}
