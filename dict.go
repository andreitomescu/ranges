package ranges

func ToDict[T, V any, K comparable](rng Range[T], key func(T) K, val func(T) V) map[K]V {
	var out = make(map[K]V)
	var it = rng.Iterate()
	for it.Next() {
		var crt = it.Value()
		out[key(crt)] = val(crt)
	}
	return out
}

func ToSet[T, K comparable](rng Range[T], key func(T) K) map[K]any {
	return ToDict(rng, key, func(T) any { return nil })
}

func Dict[K comparable, V any](dict map[K]V) Range[Pair[K, V]] {
	var slc []Pair[K, V]
	for key, val := range dict {
		slc = append(slc, MakePair(key, val))
	}
	return Slice(slc)
}
