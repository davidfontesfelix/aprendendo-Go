package main

import (
	"errors"
	"fmt"
)

func main() {
	value, err := sum(10, 4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}

func sum(number1, number2 int) (int, error) {
	if number1+number2 >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return number1 + number2, nil
}
