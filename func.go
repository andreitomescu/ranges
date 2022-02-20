package ranges

type ordered interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 | ~string
}

func Identity[T any](val T) T {
	return val
}

func Equal[T comparable](x, y T) bool {
	return x == y
}

func Less[T ordered](x, y T) bool {
	return x < y
}

func Greater[T ordered](x, y T) bool {
	return x > y
}

func True[T any](T) bool {
	return true
}

func PairKey[K, V any](pair Pair[K, V]) K {
	return pair.Key
}

func PairValue[K, V any](pair Pair[K, V]) V {
	return pair.Value
}
