package main

import "fmt"

func main (){
 nota := 10
switch true {

case nota > 9:
	fmt.Println("Optimo")

case nota > 7:
	fmt.Println("Muito bem")

case nota > 6:
	fmt.Println("Bom")
default:
	fmt.Println("Precisa Melhorar, nota baixa")

}

	}
