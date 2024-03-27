package kata

func Compose[T any](funcs ...func(T) T) func(T) T {
	return func(x T) T {
		for _, f := range funcs {
			x = f(x)
		}
		return x
	}
}
