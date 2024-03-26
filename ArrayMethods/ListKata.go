package kata

type Collection[T any] struct {
	items []T
}

func ToCollection[T any](items []T) Collection[T] {
	return Collection[T]{items: items}
}

func (list *Collection[T]) Map(f func(*T)) {
	for i := range list.items {
		f(&(list.items[i]))
	}
}

func Map[T any](list *[]T, f func(*T)) {
	for i := range *list {
		f(&(*list)[i])
	}
}

func (list *Collection[T]) Filter(f func(T) bool) {
	filteredCollection := Collection[T]{}
	for _, element := range list.items {
		if f(element) {
			filteredCollection.items = append(filteredCollection.items, element)
		}
	}
	list.items = filteredCollection.items
}

// func (list *Collection[T]) Foldr(f func(T), acc T) T {

// }

// All methods must iterate through indexes,
// as if we do something like:
// "for _, element := range"
// is the same thing as creating a copy of each value
// and not altering the value contained within the address
