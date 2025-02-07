package main

import "fmt"

func main__() {

	type posicao struct {
		X int
		Y int
	}
	var pos posicao = posicao{Y: 33, X: 23}

	//modificando valor
	pos.Y = 51

	fmt.Println(pos)
	fmt.Println("Valor do X", pos.X)
	fmt.Println("Valor do Y", pos.Y)

}
