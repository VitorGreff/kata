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

func (list *Collection[T]) Filter(f func(T) bool) {
	filteredCollection := Collection[T]{}
	for i := range list.items {
		if f(list.items[i]) {
			filteredCollection.items = append(filteredCollection.items, list.items[i])
		}
	}
	list.items = filteredCollection.items
}

// All methods must iterate through indexes,
// as if we do something like:
// "for _, element := range"
// is the same thing as creating a copy of each value
// and not altering the value contained within the address
