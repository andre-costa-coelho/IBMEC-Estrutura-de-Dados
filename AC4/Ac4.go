package main

import "fmt"

func main()  {
	qtd := torreHanoi(3, "A", "C", "B", 1)
	fmt.Println("Movimentos mínimos necessarios:", qtd-1 )
}

func torreHanoi(peca int, orig string, dest string, trab string, mov int) int  {
	if peca>0 {
		mov = torreHanoi(peca-1, orig, trab, dest, mov)
		fmt.Println("Movimento: ", mov, "Peça: ",peca, "de:", orig,"para:", dest)
		mov++
		mov = torreHanoi(peca-1, trab, dest, orig, mov)
	}

	return mov
}