package main

import (
	"fmt"
	"strings"
)

type Pessoa struct{
	Nome	string
	Email	string
}

func (p*Pessoa) AlteraEmail(novoEmail string) {
	p.Email = novoEmail
}

func adicionaPessoa(p Pessoa, a * [5]Pessoa) {
	for ind, pessoa := range a{
		if (pessoa == Pessoa{}) {
			a[ind] = p
			break
		}
	}
}

func main(){
	var pessoas [5]Pessoa
	p1 := Pessoa{
		Nome:	"aaa",
		Email:	"bbb",
	}

	fmt.Println(pessoas)
	adicionaPessoa(p1,&pessoas)
	fmt.Println(pessoas)


	x := 5
	y := x
	fmt.Println(x, y)	//5, 5

	x = 6
	fmt.Println(x, y)	//6, 5

	z := &x // z é um ponteiro, que aponta para x
	fmt.Println(z)
	fmt.Println(x, *z) //6, 6	(*z -> dereferência)
	fmt.Println(&x) // (&x -> referência)

	var w * int
	fmt.Println(w)

	a := &z		//Ponteiro de ponteiro (Ñ é muito utilizado)
	fmt.Println(a)
	fmt.Println(*a)
	fmt.Println(**a)

	mensagem := "Olá, mundo"
	alteraMensagem(&mensagem)
	fmt.Println(mensagem)
}

func alteraMensagem(msg * string) {
	// strings.ReplaceAll(stringOriginal, textoProcurar string, textoSubstituir string)
	// "mundo" -> "turma"
	*msg = strings.ReplaceAll(*msg, "mundo", "turma")
}