package kata

import "errors"

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

func ForEach[T any](collection []T, f func(T)) {
	for _, element := range collection {
		f(element)
	}
}

func Compose[T any](funcs ...func(T) T) func(T) T {
	return func(x T) T {
		for _, f := range funcs {
			x = f(x)
		}
		return x
	}
}

func Zip[T any](collection1, collection2 []T) [][]T {
	len := MinSize(collection1, collection2)
	tupleArray := [][]T{}

	for i := range len {
		tupleArray = append(tupleArray, []T{collection1[i], collection2[i]})
	}
	return tupleArray
}

func FindIndex[T any](collection []T, f func(T) bool) (int, error) {
	for i, element := range collection {
		if f(element) {
			return i, nil
		}
	}
	return -1, errors.New("kata: element not found within the slice")
}

func FlatMap[T any](collection [][]T, f func(T) T) []T {
	resutl := []T{}
	for _, sublist := range collection {
		for _, element := range sublist {
			resutl = append(resutl, f(element))
		}
	}
	return resutl
}

func GroupBy[T comparable](collection []T, f func(T) T) map[T][]T {
	result := map[T][]T{}
	for _, element := range collection {
		result[f(element)] = append(result[f(element)], element)
	}
	return result
}

func Take[T any](collection []T, n int) []T {
	if n > len(collection) {
		return collection
	}
	return collection[:n]
}

func TakeWhile[T any](collection []T, f func(T) bool) []T {
	index := 0
	for _, element := range collection {
		if f(element) {
			index++
		}
	}
	return collection[:index]
}

func Drop[T any](collection []T, n int) []T {
	if n < 0 {
		return []T{}
	}
	return collection[n:]
}

func DropWhile[T any](collection []T, f func(T) bool) []T {
	index := 0
	for _, element := range collection {
		if f(element) {
			index++
		} else {
			break
		}
	}
	return collection[index:]
}