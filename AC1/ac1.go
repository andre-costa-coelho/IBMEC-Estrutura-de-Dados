package main

import "fmt"
func e_primo(n int) bool {
/*Elabore uma função e_primo() que retorna se um número é primo ou não.
Caso o número não seja primo, liste todos os números pelos quais ele é divisível.*/
for i := 2; i <= n-1; i++ {
	if n % i == 0{
		return true
	} else {
		fmt.Print(n)
		return false
	}
}
}
func main(){
	fmt.Println(e_primo(2))
}