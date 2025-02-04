package main


// Em Go, o while tradicional de outras linguagens (como C, Java ou Python) não existe. 
//No entanto, podemos simular um while utilizando um for sem inicialização e sem incremento.
func main__() {
	
	soma := 2

	for soma < 600 { // while
		soma  = soma + soma
		
	
		println(soma)
	}
	println(soma)
}
