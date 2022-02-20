package ranges

func Repeat[T any](val T, cnt int) Range[T] {
	return &repeatRange[T]{val: val, cnt: cnt}
}

type repeatRange[T any] struct {
	val T
	cnt int
}

func (this *repeatRange[T]) Iterate() Iterator[T] {
	return &repeatIterator[T]{val: this.val, cnt: this.cnt}
}

type repeatIterator[T any] struct {
	val T
	cnt int
}

func (this *repeatIterator[T]) Value() T {
	return this.val
}

func (this *repeatIterator[T]) Next() bool {
	if this.cnt > 0 {
		this.cnt -= 1
		return true
	}
	return false
}
