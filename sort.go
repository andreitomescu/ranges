package ranges

import (
	"sort"
)

func Sort[T any](rng Range[T], less func(T, T) bool) Range[T] {
	var slc = ToSlice(rng)
	sort.Slice(slc, func(i, j int) bool { return less(slc[i], slc[j]) })
	return Slice(slc)
}

func StableSort[T any](rng Range[T], less func(T, T) bool) Range[T] {
	var slc = ToSlice(rng)
	sort.SliceStable(slc, func(i, j int) bool { return less(slc[i], slc[j]) })
	return Slice(slc)
}
