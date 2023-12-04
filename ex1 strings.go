package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	resultado := strings.ToUpper(input)

	fmt.Println("String em maiúsculas:", resultado)
}
