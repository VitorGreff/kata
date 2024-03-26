package main

import (
	"fmt"
	kata "kata/ArrayMethods"
	"log"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	arr2 := []int{1, 2, 3, 4}
	strings := []string{"oi", "alo"}
	log.Println(arr)
	log.Println(strings)

	log.Println(kata.Map(arr, func(v int) int {
		return v * 2
	}))
	log.Println(kata.Filter(arr, func(v int) bool {
		return v%2 == 0
	}))
	log.Println(kata.Foldr(arr, func(acc, curr int) int {
		return acc + curr
	}, 0))
	log.Println(kata.Foldr(strings, func(acc, curr string) string {
		return acc + curr
	}, "test"))

	f := func(v int) bool { return v%2 == 0 }
	log.Println(kata.Filter(arr, f))

	log.Println(kata.Zip(arr, arr2))
	kata.ForEach(arr, func(v int) {
		log.Println(v)
	})

	addTwo := func(v int) int { return v + 2 }
	multiplyThree := func(v int) int { return v * 3 }
	multiplyTen := func(v int) int { return v * 10 }

	test := kata.Compose(addTwo, multiplyThree, multiplyTen)
	log.Println(test(4))
	log.Println(kata.Compose(multiplyTen, multiplyThree, multiplyTen)(221))

	nestedList := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	log.Println(kata.FlatMap(nestedList, multiplyTen))

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Função de agrupamento: retorna o resto da divisão por 2
	grouper := func(x int) int {
		return x % 2
	}

	// Aplicando groupBy para agrupar números pares e ímpares
	result := kata.GroupBy(numbers, grouper)

	fmt.Println(result) //
}
