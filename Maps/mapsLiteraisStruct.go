package main

import "fmt"

type Perfil struct {
	Altura float64
	Activo bool
	Profissao string
}
func main() {

	var Perfis map [string] Perfil = map[string]Perfil{
		"Joao": {
			1.74, true, "Medico",
		},
		"Maria" : {
			1.66, true, "Eng",
		},
	}


	fmt.Println(Perfis)
}