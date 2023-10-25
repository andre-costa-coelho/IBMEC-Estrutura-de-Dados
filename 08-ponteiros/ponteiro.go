package main

import "fmt"

func main(){
	// Ponteiros armazenam valores => Representam ENDEREÇOS
	// 3s ou 64 bits

	var x * string
	var y = "Olá"
	x = &y	//Aponta para x
	z := *x // pega o valor

	/* Var		End		Valor		|		End		Valor
		x		150		151(end)	|		150		151
		y		151		"olá"		|		151		"olá"
		z		152		"olá"		|		152				*/
}