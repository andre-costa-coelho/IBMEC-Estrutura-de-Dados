package main

import "fmt"

func main(){
	var x = 10

	if x > 18 {
		fmt.Println("Você é maior de idade.")
	} else if x < 0 {
		fmt.Println("Valor inválido.")
	} else {
		fmt.Println("Você é menor de idade.")
	}


	var dia = "segunda"

	switch dia {
	case "segunda", "terça", "quarta", "quinta", "sexta":
		fmt.Println("Dia de semana.")
	case "sábado", "domingo":
		fmt.Println("Final de semana.")
	default:
		fmt.Println("Erro, dia inválido.")
	}

	//Estrutura de repetição:
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//for funcionando como while
	fmt.Println("--------------")
	x = 5

	for x > 0 {
		fmt.Println(x)
		x--
	}

	fmt.Println("--------------")
	for x < 10 {
		x++
		if x == 3 {
			continue // interrompe a interação que está em execução
		}
		fmt.Println(x)
		if x == 5{
			break // interrompe por completo do loop
		}
	}

}