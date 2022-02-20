package ranges

func Concat[T any](rng ...Range[T]) Range[T] {
	return &concatRange[T]{rng: rng}
}

type concatRange[T any] struct {
	rng []Range[T]
}

func (this *concatRange[T]) Iterate() Iterator[T] {
	return &concatIterator[T]{rng: this.rng}
}

type concatIterator[T any] struct {
	crt Iterator[T]
	rng []Range[T]
}

func (this *concatIterator[T]) Value() T {
	return this.crt.Value()
}

func (this *concatIterator[T]) Next() bool {
	if this.crt == nil {
		if len(this.rng) == 0 {
			return false
		}
		this.crt = this.rng[0].Iterate()
		this.rng = this.rng[1:]
	}
	for {
		if this.crt.Next() {
			return true
		}
		if len(this.rng) == 0 {
			return false
		}
		this.crt = this.rng[0].Iterate()
		this.rng = this.rng[1:]
	}
}
