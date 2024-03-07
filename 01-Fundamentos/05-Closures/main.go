package main

import "fmt"

func main() {
	// função anonyma 
	total := func() int {
		return sum(1, 3, 6, 9) * 2
	}()

	fmt.Println(total)
}

func sum(numbers ...int) int {
	total := 0
	for _, numbers := range numbers {
		total += numbers
	}
	return total
}
