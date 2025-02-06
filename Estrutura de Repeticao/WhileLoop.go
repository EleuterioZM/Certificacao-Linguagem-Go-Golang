package main

import (
	"fmt"
	"time"
)

func main() {
	const delay = 100 * time.Millisecond // Definição do tempo de espera
	soma := 2
	contador := 0

	fmt.Println("Iniciando o loop de soma exponencial...")

	for soma < 600 {
		soma += soma
		contador++
		
		fmt.Printf("Iteração %d: soma = %d\n", contador, soma)
		time.Sleep(delay)
	}

	fmt.Println("Fim do loop. O valor ultrapassou 600.")
}
