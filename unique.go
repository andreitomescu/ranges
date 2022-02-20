package ranges

func Unique[T any, K comparable](rng Range[T], key func(T) K) Range[T] {
	return &uniqueRange[T, K]{rng: rng, key: key}
}

type uniqueRange[T any, K comparable] struct {
	rng Range[T]
	key func(T) K
}

func (this *uniqueRange[T, K]) Iterate() Iterator[T] {
	return &uniqueIterator[T, K]{set: make(map[K]any), it: this.rng.Iterate(), key: this.key}
}

type uniqueIterator[T any, K comparable] struct {
	set map[K]any
	it  Iterator[T]
	key func(T) K
}

func (this *uniqueIterator[T, K]) Value() T {
	return this.it.Value()
}

func (this *uniqueIterator[T, K]) Next() bool {
	for this.it.Next() {
		var key = this.key(this.it.Value())
		if _, ok := this.set[key]; ok {
			continue
		}
		this.set[key] = nil
		return true
	}
	return false
}
