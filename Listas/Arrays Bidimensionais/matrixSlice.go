package main

import (
	"fmt"
	
)

func main() {

	tabuleiro := [][]string{
		{"X", "_", "_"},
		{"0", "X", "0"},
		{"_", "0", "X"},
	}
	
	for i := 0; i < len(tabuleiro); i++ {
		//tabuleiro[i]
		for j := 0; j < len(tabuleiro[i]); j++{
			//tabuleiro[i][j]
			fmt.Println(tabuleiro[i][j])

		}
	}
}
