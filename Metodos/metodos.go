package main

import "fmt"

type Pessoa1 struct {
	Nome, Sobrenome string
}

// Método NomeCompleto retorna o nome completo da pessoa
func (p Pessoa) NomeCompleto2() string {
	return p.Nome + " " + p.Sobrenome
}

func main_() {
	alguem := Pessoa{"Jose", "Zacarias"}

	// Imprime o nome completo
	fmt.Println(alguem.NomeCompleto())
}