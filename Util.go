package kata

func MinSize[T any](x, y []T) int {
	if len(x) < len(y) {
		return len(x)
	}
	return len(y)
}
