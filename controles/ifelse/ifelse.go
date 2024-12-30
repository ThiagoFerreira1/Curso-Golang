package main

import "fmt"

func main() {
	imprimirResultado()
}

func imprimirResultado() {

	media := calculaMedia(8, 5, 2, 1)

	if media <= 6 {
		fmt.Println("Reprovado")
	}

	fmt.Println("Aprovado")

}

func calculaMedia(nota1 float32, nota2 float32, nota3 float32, nota4 float32) float32 {

	resultado := nota1 + nota2 + nota3 + nota4

	media := resultado / 4

	fmt.Println("Resultado: ", resultado)
	fmt.Println("Média: ", media)
	return media

}
