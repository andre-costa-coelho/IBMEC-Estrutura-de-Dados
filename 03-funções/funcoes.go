package main

import "fmt"

func main(){
	fmt.Println(soma(2,6))
	fmt.Println(subtracao(2,6))

	var x, y = 4.5, -10.2
	x, y = trocaValores(x, y)
	fmt.Println(x, y)

	anonima()
}

// func nome (param1 tipo, param2 tipo,...) tipo do retorno {}
func soma(a int,b int) int{
	return a + b
}

// parametros do mesmo tipo, posso informar apenas uma vez
func subtracao(a,b float64) float64 {
	return a - b
}

// pode ter mais de um retorno
func trocaValores(a, b float64) (float64,float64) {
	return b, a
}

// anônimas (closure)
func anonima(){
	dobra := func(x int) int {
		return x*2
	}

	resultado := dobra(5)
	fmt.Println(resultado)

	calcular := func(f func(int) int, x int) int{
		return f(x)
	}

	resultado = calcular(dobra,8)
	fmt.Println(resultado)
}
