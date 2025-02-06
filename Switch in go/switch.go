package main

import "fmt"

func main_a(){
nome := "Eleuterio"

	switch nome {
	case "Ana":  
		fmt.Println("E a Ana")

	case "Eleuterio":
		fmt.Println("E o Eleuterio")
	default:
		fmt.Println("Nao conheco")
	}
}