package main

import "fmt"

func main_SL() {

	// Arrays
	 nomes := [] string {
		"Eleuterio",
		"Zacarias",
		"Mabecuane",
	 }
//slice Literais

// println("Capacidade", cap(nomes))

nomes = nomes [:2]

fmt.Println(nomes)

fmt.Printf("len=%d, cap=%d\n", len(nomes),  cap(nomes))

fmt.Printf("%T\n", nomes)

}
