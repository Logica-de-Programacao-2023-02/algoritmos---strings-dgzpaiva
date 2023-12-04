package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	eSequenciaCrescente := verificarSequenciaCrescente(input)

	if eSequenciaCrescente {
		fmt.Println("A string é uma sequência numérica crescente.")
	} else {
		fmt.Println("A string NÃO é uma sequência numérica crescente.")
	}
}

func verificarSequenciaCrescente(str string) bool {
	if len(str) <= 1 {
		return false
	}

	for i := 0; i < len(str)-1; i++ {
		curr, _ := strconv.Atoi(string(str[i]))
		next, _ := strconv.Atoi(string(str[i+1]))

		if next != curr+1 {
			return false
		}
	}

	return true
}
