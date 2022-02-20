package ranges

type Pair[K, V any] struct {
	Key   K
	Value V
}

func MakePair[K, V any](key K, val V) Pair[K, V] {
	return Pair[K, V]{Key: key, Value: val}
}
