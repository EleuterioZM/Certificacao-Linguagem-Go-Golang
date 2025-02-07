package main

import "fmt"

func computar(fn func(float64, float64) float64) float64 {
	return fn(5, 6)
}

func main___() {
	funcaoSomar := func(a, b float64) float64 {
		return a + b
	}

	funcaoMultiplicar := func(a, b int) int {
		return a * b
	}

	fmt.Println(computar(funcaoSomar)) // Isso funciona, pois funcaoSomar é compatível com o tipo esperado por computar
	fmt.Println("Funcao Multiplicar:", funcaoMultiplicar(3, 4)) // Isso também funciona, mas não está relacionado à função computar
}