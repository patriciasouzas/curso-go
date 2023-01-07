package main

import "fmt"

func main() {
	//slices
	nomes := []string{"Paulo", "João", "Dani", "Pedro", "Carla"}
	fmt.Println(nomes, len(nomes), cap(nomes))

	nomes = append(nomes, "Flávio", "Jussara")
	fmt.Println(nomes, len(nomes), cap(nomes))

	nomes = append(nomes, "Fátima")
	fmt.Println(nomes, len(nomes), cap(nomes))

	//definindo tipo, tamanho e capacidade do slice
	nums := make([]int, 4, 10)
	fmt.Println(nums)

	//maps
	idades := make(map[string]uint8)
	idades["Tiago"] = 15
	idades["Carlos"] = 35
	idades["Maria"] = 32
	idades["Amanda"] = 34

	fmt.Println(idades)
	fmt.Println(idades["Maria"])
	fmt.Println(idades["Xavier"]) //quando não tem o valor no map, ele retorna o default da chave

	val, ok := idades["Tiago"] //valor existe
	fmt.Println(val, ok)

	val, ok1 := idades["Lucas"] //valor não existe
	fmt.Println(val, ok1)
}
