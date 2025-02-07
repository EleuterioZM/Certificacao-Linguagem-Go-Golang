package main

import "fmt"



func main__() {


// var funcaoSomar func(float64, float64) float64 = func(a, b float64) float64 {
// 	return a + b
// }
// Incurtado

 funcaoSomar  := func(a, b float64) float64 {
	return a + b
}

funcaoMultiplicar := func (a, b int) int {
	return a * b
}

	fmt.Println("Funcao Somar",funcaoSomar(3, 4))
	fmt.Println("Funcao Multiplicar", funcaoMultiplicar(3, 4))
}