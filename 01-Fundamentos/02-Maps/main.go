package main

import "fmt"

func main() {
	salarios := map[string]int{"Wesley": 1000, "Maria": 1600, "Lucas": 2000}

	fmt.Println(salarios["Wesley"])
	for nome, salarios := range salarios {
		fmt.Printf("O slario de %s é %d\n", nome, salarios)
	}

	for _, salarios := range salarios {
		fmt.Printf("Salario %d\n", salarios)
	}

}
