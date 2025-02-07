package main

import "fmt"

func main_() {

	var notas map[string]int
	notas = make(map[string]int)


	//Ana ->9
	//Maria ->10
	notas["Ana"] = 9
	notas["Maria"] = 10

	fmt.Println(notas)
	fmt.Println("Nota Ana", notas["Ana"])
	fmt.Println("Nota Maria", notas["Maria"])
	
valor, existe := notas["Maria"]
if existe {
	fmt.Println(valor)
}
fmt.Println(existe)

}