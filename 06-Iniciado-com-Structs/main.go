package main

import "fmt"

type Client struct {
	uid    string
	name   string
	age    int
	active bool
}

func main() {
	david := Client{
		uid:    "zh3893-defh8e-ws545t",
		name:   "David",
		age:    18,
		active: true,
	}
	
	david.active = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", david.name, david.age, david.active)

}
