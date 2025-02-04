package main

import "fmt"

func calcular(a int) (quadrado int, cubo int) {

	// var quadrado int = a * a
	// var cubo int = a * a * a

	quadrado = a * a
	cubo = a * a * a

	// return quadrado, cubo
	// returno null
	return
}

func main() {

	fmt.Println(calcular(2))

}
