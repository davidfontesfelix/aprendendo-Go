package main

func main() {
	// Memória -> Endereço -> Valor

	// variavel -> aponta para um ponteiro que tem um endereço na memoria -> valor

	value := 10
	var ponteiro *int = &value
	*ponteiro = 20

	value2 := &value
	*value2 = 30

	println(ponteiro)
	println(value)

	training()
}

func training() {
	var ponteiro *string

	nameInitial := "david"
	ponteiro = &nameInitial

	name2 := "julio"
	*ponteiro = name2

	name3 := "lucas"
	ponteiro = &name3


	println(*ponteiro, nameInitial, name2)

}
