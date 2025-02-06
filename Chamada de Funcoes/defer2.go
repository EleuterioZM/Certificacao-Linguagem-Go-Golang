package main

import "fmt"

func imprimir() string {
	fmt.Println("Imprimindo......")
	return "VALOR DE IMPRIMIR"
}

func main() {
	defer fmt.Println(imprimir())
	 fmt.Println(1)
	 fmt.Println(2)
	 fmt.Println(3)

}
