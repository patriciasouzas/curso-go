package main

import "fmt"

func main() {
	nome := []string{"Marcos", "Pedro", "Luiza", "Júlia"}
	for i := 0; i < len(nome); i++ {
		fmt.Println(nome[i])
	}

	//outra forma de fazer
	for k, nomes := range nome {
		fmt.Println(k, nomes)
	}

	//ignorando a key
	for _, nomes := range nome {
		fmt.Println(nomes)
	}

	//outra forma, como se fosse while
	var i int

	for i < len(nome) {
		fmt.Println(nome[i])
		i++
	}
}
