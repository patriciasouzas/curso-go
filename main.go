package main

import "fmt"

var (
	nome string
	n1   int
	n2   int32
)

func hello(nome string) {
	fmt.Println("Hello,", nome, "!")
}

func main() {
	mensagem := "Aula 03 - Go 101" //declaração com atribuição
	fmt.Println(mensagem)

	var b, f, s = true, 2.3, "Olá"
	fmt.Println(b, f, s)

	var total int
	total++ //atribuição no valor default
	fmt.Println("total:", total)

	nome = "Patricia"
	fmt.Println("Hello,", nome, "!")

	var x, y = 10, 20
	fmt.Println(x, y)
	x, y = y, x //brincando com substituição de valores
	fmt.Println(x, y)

	hello("Carla Regina") //usando função
}
