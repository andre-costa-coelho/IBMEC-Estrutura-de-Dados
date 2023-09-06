/*Dado um array de números inteiros positivos e um valor alvo, encontre um par de números no array
cuja soma seja igual ao valor alvo. Se nenhum par for encontrado, retorne um valor (-1, -1)
indicando que nenhum par foi encontrado.*/

package main

import "fmt"

func main(){
	array := []int {1,2,3,4,5}
	numAlvo := 3

	fmt.Println(buscaPares(array,numAlvo))
}

func buscaPares(array[]int, numAlvo int) (int, int) {
	for _, i := range array {
		for _, j := range array {
			if (i+j == numAlvo) {
				return i,j
			}
		}
	}
	return -1,-1
}