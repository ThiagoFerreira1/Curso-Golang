package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(tipo(time.Now()))
}

func tipo(i interface{}) string {

	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "Não sei"
	}
}
