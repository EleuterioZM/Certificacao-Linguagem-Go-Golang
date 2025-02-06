package main

import (
	"fmt"
	"strings"
)

func main_() {

	tabuleiro := [][]string{
		{"X", "_", "_"},
		{"0", "X", "0"},
		{"_", "0", "X"},
	}
	for i := 0; i < len(tabuleiro); i++ {

		// fmt.Println(tabuleiro[i])
		fmt.Printf("%s\n", strings.Join(tabuleiro[i], "  "))
	}
}
