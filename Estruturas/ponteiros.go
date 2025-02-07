package main

import "fmt"

func main() {

	type posicao struct {
		X int
		Y int
	}
	var pos posicao = posicao{Y: 33, X: 23}

	var K *posicao = &pos
	
	(*K).X = 22

	fmt.Println(    (*K).X        )

}
