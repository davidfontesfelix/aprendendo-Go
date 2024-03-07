package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 3, 6, 9, 18))
}

func sum(numbers ...int) int {
	total := 0
	for _, numbers := range numbers {
		total += numbers
	}
	return total
}
