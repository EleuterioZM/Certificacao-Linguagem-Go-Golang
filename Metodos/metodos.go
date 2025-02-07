package main

import "fmt"

type Pessoa struct {
	Nome, Sobrenome string
}

// Método NomeCompleto retorna o nome completo da pessoa
func (p Pessoa) NomeCompleto() string {
	return p.Nome + " " + p.Sobrenome
}

func main() {
	alguem := Pessoa{"Jose", "Zacarias"}

	// Imprime o nome completo
	fmt.Println(alguem.NomeCompleto())
}