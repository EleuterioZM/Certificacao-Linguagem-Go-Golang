package main

import "fmt"

func main() {

	a := 32

	// ponteiro armazena endereco na memoria
	p := &a
	*p = 18
	*p = *p + 2
	fmt.Println("Endereco na memoria", p)
	fmt.Println(*p)
}
