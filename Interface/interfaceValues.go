package main

import "fmt"



type Potencia interface {
	Quad() int
}


type Inteiro int

func (i Inteiro) Quad() int {
	return int (i * i)
}


func descrever(Potencia Potencia)  {
	fmt.Printf("Valor : %v, Tipo : %T\n", Potencia, Potencia )
}

func main() {
	// var num Inteiro = 3

	
	var pot Potencia
	pot = Inteiro(3)
	fmt.Println(pot.Quad())

}