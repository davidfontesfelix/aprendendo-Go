package main

func sum(value1 *int, value2 int) int {
	*value1 = 50
	value2 = 40
	return *value1 + value2
}

func main() {
	number1 := 10
	number2 := 20
	println(sum(&number1, number2))

	println(number1) //number1 vai passar a ser 50
	println(number2) //number2 vai continuar sendo 20
}
