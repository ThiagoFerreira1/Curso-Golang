package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notafinal := int(nota)
	fmt.Println(notafinal)

	// warning
	fmt.Println("Teste" + string(97))

	// int para string
	fmt.Println("Teste" + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("True")
	if b {
		fmt.Println("Verdadeiro")
	}
}
