package main

import "fmt"

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(0))
	fmt.Println(notaParaConceito(-1))
}

func notaParaConceito(n float64) string {
	var nota = int(n)

	switch {
	case nota >= 9 && nota <= 10:
		return "A"
	case nota >= 8 && nota < 9:
		return "B"
	case nota >= 5:
		return "C"
	case nota >= 3:
		return "D"
	case nota < 4:
		return "F"
	default:
		return "Nota inválida!"
	}
}
