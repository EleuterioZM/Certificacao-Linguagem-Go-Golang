package main

import (
	"fmt"
)

func main () {
var animes [3]string;

animes = [3]string{"Dragon Ball", "Sailor Moon", "Yuyu Hakusho"}

var doisPrimeiros = animes[:2]
fmt.Println("Todo slice :",animes)
fmt.Println("Dois Primeiros do Slice :", doisPrimeiros)


fmt.Printf("len = %d, cap = %d\n", len(doisPrimeiros), cap(doisPrimeiros))


var doisUltimo = animes [1:]
fmt.Println("Dois Ultimos do Slice :", doisUltimo)
fmt.Printf("len = %d, cap = %d\n", len(doisUltimo), cap(doisUltimo))
}