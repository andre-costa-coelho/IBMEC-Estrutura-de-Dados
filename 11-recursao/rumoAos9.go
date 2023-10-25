package main

import (
	"fmt"
	"strconv"
)

func main(){
	var numeros []string = make([]string, 0)

	for {
		var numero string
		fmt.Scanln(&numero)

		if numero == "0"{
			break
		}

		numeros = append(numeros, numero)
	}

	for _, numero := range numeros {
		fmt.Println(numero, rumoAos9(numero,0))
	}
}

func rumoAos9(n string, grau int)string {
	if n == "9"{
		if grau == 0{
			grau = 1
		}
		return "is a multiple of 9 and has 9-degree" + strconv.Itoa(grau) + "."
	}

	if len(n) == 1 {
		return "is not a multiple of 9"
	}

	soma := 0

	for _, digito := range n {
		soma += int(digito - '0')
	}

	return rumoAos9(strconv.Itoa(soma),grau + 1)
}