package main

import "fmt"

func main___() {
    i := 0

    for { // Loop infinito (while true)
        fmt.Println("Valor de i:", i)
        i++

        if i >= 5 { // Condição de saída
            break
        }
    }
}
