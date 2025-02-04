package main



func Adicao (a int ,b int )int{
	return a + b
}

func Sub (a int ,b int )int{
	return a - b
}

func Mul (a int ,b int )int{
	return a * b
}

func Div (a int ,b int )int{
	return a / b
}

func RestDiv (a int ,b int )int{
	return a % b
}
func main() {

//  + - * / %

println("A soma de a e b e :", Adicao(23, 2))
println("A subtracao de a e b e :", Sub(23, 2))   
println("A multiplicacao de a e b e :", Mul(23, 2))   
println("A divisao de a e b e :", Div(23, 2))   
println("O resto da divisao de a e b e :", RestDiv(23, 2))   


}
