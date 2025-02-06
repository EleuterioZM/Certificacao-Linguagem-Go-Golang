package main

import (
	"fmt"

)

func main(){
	var nomes [] string;

var nomes2 = append(nomes, "Eleuterio")
 nomes2 = append(nomes2, "Maria")
 nomes2 = append(nomes2, "Ana", "Mateus", "Rogerio")

	fmt.Println(nomes)
	fmt.Println(nomes2)
	fmt.Printf("len = %d; cap = %d\n", len(nomes2), cap(nomes2))
}