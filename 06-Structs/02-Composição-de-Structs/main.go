package main

import "fmt"

type Address struct {
	city  string
	state string
}

type Client struct {
	uid    string
	name   string
	age    int
	active bool
	Address
}

func main() {
	david := Client{
		uid:    "zh3893-defh8e-ws545t",
		name:   "David",
		age:    18,
		active: true,
	}
	david.Address.city = "Aracaju"

	david.active = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t, Cidade: %s", david.name, david.age, david.active, david.Address.city)

}
