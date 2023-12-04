package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	eSequenciaDecrescente := verificarSequenciaDecrescente(input)

	if eSequenciaDecrescente {
		fmt.Println("A string é uma sequência numérica decrescente.")
	} else {
		fmt.Println("A string NÃO é uma sequência numérica decrescente.")
	}
}

func verificarSequenciaDecrescente(str string) bool {
	if len(str) <= 1 {
		return false
	}

	for i := 0; i < len(str)-1; i++ {
		curr, _ := strconv.Atoi(string(str[i]))
		next, _ := strconv.Atoi(string(str[i+1]))

		if next != curr-1 {
			return false
		}
	}

	return true
}
