package kata

func Map[T any](collection []T, f func(T) T) []T {
	newCollection := []T{}
	for _, element := range collection {
		newCollection = append(newCollection, f(element))
	}
	return newCollection
}

func Filter[T any](collection []T, f func(T) bool) []T {
	newCollection := []T{}
	for _, element := range collection {
		if f(element) {
			newCollection = append(newCollection, element)
		}
	}
	return newCollection
}

func Foldr[T any](collection []T, f func(T, T) T, acc T) T {
	for _, element := range collection {
		acc = f(acc, element)
	}
	return acc
}
