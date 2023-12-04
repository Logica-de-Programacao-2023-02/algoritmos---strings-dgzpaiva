package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string

	fmt.Print("Digite um número em ponto flutuante: ")
	fmt.Scanln(&input)

	valido := verificarNumeroFlutuante(input)

	if valido {
		fmt.Println("É um número em ponto flutuante válido.")
	} else {
		fmt.Println("Não é um número em ponto flutuante válido.")
	}
}

func verificarNumeroFlutuante(str string) bool {

	_, err := strconv.ParseFloat(str, 64)

	return err == nil
}
