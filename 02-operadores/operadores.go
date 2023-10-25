package main

import "fmt"

func main(){
	var x,y,z = 4,10,-2

	// Aritiméticos:
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)

	// Atribuição:
	z += 2 //z=z+2
	z -= x
	z *= 3
	z /= 2
	z %= 3

	// Comparação:
	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x < y)
	fmt.Println(x <= y)
	fmt.Println(x > y)
	fmt.Println(x >= y)

	// Lógicos:
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	//Não vamos ver agora - Bitwise, Memória e Canal
}