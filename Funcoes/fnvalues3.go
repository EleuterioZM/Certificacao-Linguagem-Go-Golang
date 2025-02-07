package main

import "fmt"

func saudar(nome string) func() string {
	// return "Ola, "+ nome+ "!"
	// return fmt.Sprintf("Ola, %s!", nome)

	return func() string {
		return fmt.Sprintf("Ola, %s!", nome)
	}
}

func main_sf() {

	cumprimentarJoao := saudar("Eleuterio")
fmt.Println(cumprimentarJoao())
}