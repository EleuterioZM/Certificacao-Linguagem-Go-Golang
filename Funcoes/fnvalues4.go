package main

import "fmt"

func adicionador() func(int) int {
	sum := 0

	//clousure
	return func(i int) int {
		sum = sum + i
		return sum
	}
}

func main() {
	ad1 := adicionador()
	fmt.Println(ad1(21))
	fmt.Println(ad1(13))
}