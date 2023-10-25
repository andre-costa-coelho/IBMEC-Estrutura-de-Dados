package main

import	"fmt"

type Pessoa struct {
	Nome	string
}

func main(){
	var p Pessoa
	p = Pessoa{
		Nome:	"André",
	}

	var pessoas [3]Pessoa

	pessoas[0] = p
	for _, pessoa := range pessoas {
		if (pessoa != Pessoa{}) {
			fmt.Println(pessoa) // Quais elementos no array não estão vazios.
		}

	}
}
