package ranges

func Group[T any, K comparable](rng Range[T], key func(T) K) map[K]Range[T] {
	var groups, it = make(map[K]Range[T]), rng.Iterate()
	for it.Next() {
		var group = key(it.Value())
		if slc, ok := groups[group]; ok {
			slc.(SliceRange[T]).Append(it.Value())
		} else {
			groups[group] = Values(it.Value())
		}
	}
	return groups
}
