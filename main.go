package main

import (
	kata "kata/ArrayMethods"
	"log"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	c := kata.ToCollection(arr)

	log.Println(c)
	c.Map(func(v *int) {
		*v *= 3
	})

	log.Println(c)
	c.Filter(func(v int) bool {
		return v%2 == 0
	})

	log.Println(c)
}
