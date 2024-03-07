package main

import "fmt"

func main() {
	var myVar interface{} = "David Fontes"

	println(myVar.(string))
	res, ok := myVar.(int)
	fmt.Printf("The res value is %v and the result of the ok is %v", res, ok)
	// vai dar panic
	// res2 := myVar.(int)
	// fmt.Printf("The res2 value is %v", res2)
}
