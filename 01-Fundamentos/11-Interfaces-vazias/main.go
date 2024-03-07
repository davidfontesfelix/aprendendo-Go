package main

import "fmt"

func main() {

	var valueX interface{} = 10
	var valueY interface{} = "Hello, World"

	showType(valueX)
	showType(valueY)
}

func showType(typeValue interface{}) {
	fmt.Printf("The variable type is %T and the value is %v \n", typeValue, typeValue)
}
