package ranges

// Returns the number of elements in a range.
func Count[T any](rng Range[T]) int {
	var it = rng.Iterate()
	var cnt = 0
	for it.Next() {
		cnt += 1
	}
	return cnt
}
