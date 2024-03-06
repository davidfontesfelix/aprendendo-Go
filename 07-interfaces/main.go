package main

import "fmt"

type Address struct {
	city  string
	state string
}

type Person interface {
	Disable()
}

type Company struct {
	name string
}

func (company Company) Disable() {}

type Client struct {
	uid    string
	name   string
	age    int
	active bool
	Address
}

func (client Client) Disable() {
	client.active = false
	fmt.Printf("O cliente %s foi desativado", client.name)
}

func deactivation(person Person) {
	person.Disable()
}

func main() {
	david := Client{
		uid:    "zh3893-defh8e-ws545t",
		name:   "David",
		age:    18,
		active: true,
	}
	myCompany := Company{
		name: "zuks",
	}

	david.Address.city = "Aracaju"

	deactivation(david)
	deactivation(myCompany)
}
