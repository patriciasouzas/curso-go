package main

import "fmt"

func main() {
	nomes := []string{"Paulo", "João", "Dani", "Pedro", "Carla"}
	fmt.Println(nomes, len(nomes), cap(nomes))

	nomes = append(nomes, "Flávio", "Jussara")
	fmt.Println(nomes, len(nomes), cap(nomes))

	nomes = append(nomes, "Fátima")
	fmt.Println(nomes, len(nomes), cap(nomes))

	//definindo tipo, tamanho e capacidade do slice
	idades := make([]int, 4, 10)
	fmt.Println(idades)

}
