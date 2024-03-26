package kata

// type Collection[T any] struct {
// 	items []T
// }

// func ToCollection[T any](items []T) Collection[T] {
// 	return Collection[T]{items: items}
// }

// func (list *Collection[T]) Map(f func(*T)) {
// 	for i := range list.items {
// 		f(&(list.items[i]))
// 	}
// }

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

// All methods must iterate through indexes,
// as if we do something like:
// "for _, element := range"
// is the same thing as creating a copy of each value
// and not altering the value contained within the address
