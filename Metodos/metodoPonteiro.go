package main

import (
	"fmt"
	"strings"
)

type Pessoa2 struct {
	Nome, Sobrenome string
}

// Método NomeCompleto retorna o nome completo da pessoa
func (p Pessoa) NomeCompleto1() string {
	return p.Nome + " " + p.Sobrenome
}

func (p *Pessoa) CapitalizarNome() {
	
	p.Nome  = strings.ToUpper(p.Nome)
}

func FuncaoNomeCompleto1(p Pessoa) string {
	return p.Nome +" "+ p.Sobrenome;
}


func main1() {
	alguem := Pessoa{"Jose", "Zacarias"}

	// Imprime o nome completo
	alguem.CapitalizarNome()
	fmt.Println(alguem.NomeCompleto())
	// fmt.Println(FuncaoNomeCompleto(alguem))
}