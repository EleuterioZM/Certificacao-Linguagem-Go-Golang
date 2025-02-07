package main

import "fmt"

type Perfil_ struct {
	Altura float64
	Activo bool
	Profissao string
}
func main_li() {

	var Perfis map [string] Perfil = make(map [string] Perfil)
	Perfis ["Eleuterio"] = Perfil{

		1.74, true, "Medico",
	}
Perfis ["Maria"] = Perfil{
	1.66, true, "Eng",
}

	fmt.Println(Perfis)
}