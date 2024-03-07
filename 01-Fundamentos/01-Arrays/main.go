package main

import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[len(meuArray)-1] = 30

	for index, value := range meuArray {
		fmt.Printf("O valor do indice %d e o valor é %d\n", index, value)
	}
}
