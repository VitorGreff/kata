package kata

func FlatMap[T any](collection [][]T, f func(T) T) []T {
	result := []T{}
	for _, sublist := range collection {
		for _, element := range sublist {
			result = append(result, f(element))
		}
	}
	return result
}

func GroupBy[T comparable](collection []T, f func(T) T) map[T][]T {
	result := map[T][]T{}
	for _, element := range collection {
		result[f(element)] = append(result[f(element)], element)
	}
	return result
}

func Zip[T any](collection1, collection2 []T) [][]T {
	len := minSize(collection1, collection2)
	tupleArray := [][]T{}

	for i := range len {
		tupleArray = append(tupleArray, []T{collection1[i], collection2[i]})
	}
	return tupleArray
}

func minSize[T any](x, y []T) int {
	if len(x) < len(y) {
		return len(x)
	}
	return len(y)
}
