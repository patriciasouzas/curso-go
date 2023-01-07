package main

import (
	"fmt"
	"strconv"
)

func Hello(nome string) {
	fmt.Println("Olá,", nome, "!")
}

func sum(a, b int) int {
	return a + b
}

func convertAndSum(a int, b string) (total int, err error) {
	// i, _ := strconv.Atoi(b) ignorando o erro
	i, err := strconv.Atoi(b)

	if err != nil {
		return
	}

	total = a + i

	return
}

func main() {
	Hello("Joana")

	fmt.Println("total:", sum(14, 6))
	total, err := convertAndSum(10, "2")
	fmt.Println("total:", total, err)
}
