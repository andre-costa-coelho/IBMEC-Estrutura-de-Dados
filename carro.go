package main

import "fmt"

type Carro struct{
	Modelo	string
	Velocidade	int
	Cor	string
	Ligado	bool

}

func(c *Carro) Ligar(){
	c.Ligado = !c.Ligado
}

func(c *Carro) Acelerar(valor int){
	c.Velocidade += valor
}

func main(){
	carro1 := Carro{
		Modelo:	"Porsche 911 Carrera",
		Velocidade:	0,
		Cor:	"Branco",
		Ligado:	false,
	}

	fmt.Println(carro1)
	carro1.Ligar()
	fmt.Println(carro1)
	carro1.Acelerar(10)
	fmt.Println(carro1)

}