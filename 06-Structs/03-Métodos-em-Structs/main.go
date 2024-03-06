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

func (client Client) disable() {
	client.active = false
	fmt.Printf("O client %s foi desativado", client.name)
}

func main() {
	david := Client{
		uid:    "zh3893-defh8e-ws545t",
		name:   "David",
		age:    18,
		active: true,
	}
	david.Address.city = "Aracaju"

	david.disable()
}
