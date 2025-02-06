package main

import "fmt"

func main_Slice() {

	// Arrays
	 numeros := [] int {1, 2, 4, 5, 5, 1, 10}
//slice

 fmt.Println(numeros[2:5]) 

fmt.Println("Omitindo o indice Inicial",numeros[:5])

fmt.Println("Omitindo o indice Final", numeros[2:]) 

fmt.Println("Omitindo o indice Inicial e Final", numeros[:]) 
}
