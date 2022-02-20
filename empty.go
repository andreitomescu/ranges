package ranges

func Empty[T any]() Range[T] {
	return Values[T]()
}
