package kata

import "errors"

func ForEach[T any](collection []T, f func(T)) {
	for _, element := range collection {
		f(element)
	}
}

func FindIndex[T any](collection []T, f func(T) bool) (int, error) {
	for i, element := range collection {
		if f(element) {
			return i, nil
		}
	}
	return -1, errors.New("kata: element not found within the slice")
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
