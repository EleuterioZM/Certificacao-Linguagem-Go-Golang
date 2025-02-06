package main

import "fmt"

func main_som() {

	// Arrays
	 numeros := [] int {1, 2, 4, 5, 5, 1, 10}
	 soma := 0
for i := 0; i < len(numeros); i++ {
	fmt.Println("Na possicao ", i, "temos :", numeros[i])

	soma += numeros[i]

	
}

fmt.Println("Soma dos numeros", soma)

}
