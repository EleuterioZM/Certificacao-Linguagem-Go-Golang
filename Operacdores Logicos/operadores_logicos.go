package main

import "fmt"

func main() {
	// Operadores lógicos: && (E), || (OU), ! (NÃO)
	var a int = 1
	var b int = 4

	// Verifica se 'a' é maior que 3 e 'b' é menor que 5
	if a > 3 && b < 5 {
		fmt.Println("Condição verdadeira: 'a' é maior que 3 E 'b' é menor que 5")
	} else {
		fmt.Println("Condição falsa: Pelo menos uma das comparações falhou")
	}

	// Verifica se 'a' é diferente de 'b'
	if a != b {
		fmt.Println("Os valores são diferentes: a =", a, "b =", b)
	} else {
		fmt.Println("Os valores são iguais: a =", a, "b =", b)
	}

	// Verifica se 'a' é maior que 3 ou 'b' é menor que 5
	if a > 3 || b < 5 {
		fmt.Println("Condição verdadeira: 'a' é maior que 3 OU 'b' é menor que 5")
	} else {
		fmt.Println("Condição falsa: Nenhuma das comparações é verdadeira")
	}

	if !(a > 3) {
		fmt.Println("Negação verdadeira: 'a' NÃO é maior que 3")
	} else {
		fmt.Println("Negação falsa: 'a' é maior que 3")
	}
}
