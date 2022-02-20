package ranges

func ToSlice[T any](rng Range[T]) []T {
	var out []T
	var it = rng.Iterate()
	for it.Next() {
		out = append(out, it.Value())
	}
	return out
}

func Slice[T any](slc []T) Range[T] {
	return &sliceRange[T]{slc: slc}
}

func Values[T any](slc ...T) Range[T] {
	return &sliceRange[T]{slc: slc}
}

type sliceRange[T any] struct {
	slc []T
}

func (this *sliceRange[T]) Iterate() Iterator[T] {
	return &sliceIterator[T]{slc: this.slc, idx: -1}
}

func (this *sliceRange[T]) Slice() []T {
	return this.slc
}

func (this *sliceRange[T]) Append(value T) {
	this.slc = append(this.slc, value)
}

type sliceIterator[T any] struct {
	slc []T
	idx int
}

func (this *sliceIterator[T]) Value() T {
	return this.slc[this.idx]
}

func (this *sliceIterator[T]) Next() bool {
	this.idx += 1
	return this.idx < len(this.slc)
}
