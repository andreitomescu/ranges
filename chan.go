package ranges

func ToChan[T any](rng Range[T], cap int) chan T {
	var out = make(chan T, cap)
	go func() {
		var it = rng.Iterate()
		for it.Next() {
			out <- it.Value()
		}
		close(out)
	}()
	return out
}

func Chan[T any](ch chan T) Range[T] {
	return &chanRange[T]{ch: ch}
}

type chanRange[T any] struct {
	ch  chan T
	val T
}

func (this *chanRange[T]) Iterate() Iterator[T] {
	return this
}

func (this *chanRange[T]) Value() T {
	return this.val
}

func (this *chanRange[T]) Next() bool {
	if val, ok := <-this.ch; ok {
		this.val = val
		return true
	}
	return false
}
