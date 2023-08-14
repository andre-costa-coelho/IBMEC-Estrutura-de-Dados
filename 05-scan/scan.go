package main

import (
	"fmt"
	"bufio"
	"os"
)


func main(){
	var x int
	var y float64

	fmt.Println("Informe um número: ")
	fmt.Scanln(&x)
	fmt.Println(x)

	fmt.Println("Informe um float: ")
	fmt.Scanln(&y)
	fmt.Println(y)

	//lendo frases inteiras
	leitor := bufio.NewReader(os.Stdin)
	fmt.Println("Escreva um texto: ")
	msg, _ := leitor.ReadString('\n')

	fmt.Println(msg)

}

