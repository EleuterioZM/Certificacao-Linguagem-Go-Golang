package main

import "fmt"

func main_mod() {

	// Arrays
	 nomes := [3] string {
		"Eleuterio",
		"Zacarias",
		"Mabecuane",
	 }
//slice

p1 := nomes[0:2]

fmt.Println(p1)

p1 [0] = "Ghost"

fmt.Println(p1)

fmt.Println("Array Original :", nomes)
}
