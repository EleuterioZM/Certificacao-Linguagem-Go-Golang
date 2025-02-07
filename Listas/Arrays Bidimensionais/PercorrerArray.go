package main

import "fmt"

func main() {

	numeros := []int{10, 20, 30, 40}

	// for i := 0; i < len(numeros); i++ {
	// 	fmt.Println(numeros[i])
	// }

	for indice, Valor := range numeros {
		fmt.Println(indice, ":", Valor)
	}
}
