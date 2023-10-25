package main

import "fmt"

func main(){
	var expressoes []string = make([]string, 0)

	for {
		var expressao string
		fmt.Scanln(&expressao)

		if expressao == ""{
			break
		}

		expressoes = append(expressoes, expressao)

	}

	for _, expressao := range expressoes {
		fmt.Println(encontreOTelefone(expressao))
	}

}

func encontreOTelefone(expr string) string {
	telefone := ""

	for _, carac := range expr {
		switch carac {
		case '1','0','-':
			telefone += string(carac)
		case 'A', 'B', 'C':
			telefone += "1"
		case 'D', 'E', 'F':
			telefone += "2"
		case 'G', 'H', 'I':
			telefone += "3"
		case 'J', 'K', 'L':
			telefone += "4"
		case 'M', 'N', 'O':
			telefone += "5"
		case 'P', 'Q', 'R','S':
			telefone += "6"
		case 'T', 'U', 'V':
			telefone += "7"
		case 'W', 'X', 'Y','Z':
			telefone += "8"
		}
	}
	return telefone
}