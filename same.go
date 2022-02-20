package ranges

func Same[T any](first, second Range[T], same func(T, T) bool) bool {
	var fit, sit = first.Iterate(), second.Iterate()
	for {
		var fnext, snext = fit.Next(), sit.Next()
		if fnext && snext {
			if same(fit.Value(), sit.Value()) {
				continue
			}
			return false
		}
		if fnext || snext {
			return false
		}
		return true
	}
}
