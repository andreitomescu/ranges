package ranges

type Range[T any] interface {
	Iterate() Iterator[T]
}

type Iterator[T any] interface {
	Value() T
	Next() bool
}

type SliceRange[T any] interface {
	Slice() []T
	Append(T)
}
