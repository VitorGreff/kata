package main

import (
	kata "kata/ArrayMethods"
	"log"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	strings := []string{"oi", "alo"}
	log.Println(arr)
	log.Println(strings)
	// c := kata.ToCollection(arr)

	// log.Println(c)
	// c.Map(func(v *int) {
	// 	*v *= 3
	// })

	// log.Println(c)
	// c.Filter(func(v int) bool {
	// 	return v%2 == 0
	// })

	// log.Println(c)
	kata.Map(&arr, func(v *int) {
		*v *= 2
	})
	log.Println(arr)

	kata.Map(&strings, func(s *string) {
		*s += "world"
	})
	log.Println(strings)
}
