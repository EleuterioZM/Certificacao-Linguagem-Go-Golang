package main

import "fmt"

func main_() {

	type posicao struct {
		X int
		Y int
	}
	pos := posicao{40, 67}

	//modificando valor
	pos.Y = 51

	fmt.Println(pos)
	fmt.Println("Valor do X", pos.X)
	fmt.Println("Valor do Y", pos.Y)

}
