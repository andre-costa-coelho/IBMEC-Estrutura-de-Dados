package main

import "fmt"

func e_primo(n int) bool {
	if n <= 1 {  // Verifica se o número é 1 ou menor que 1.
		return false
	}

	for i := 2; i*i <= n; i++ {  // Verifica se o número é primo
		if n%i == 0 {
			return false  // Se houver divisores que não seja 1 ou ele mesmo retorna falso.
		}
	}
	return true
}

func fibo(n int) int {
	
	if n <= 1 { // Verifica se o número é 1 ou menor que 1.
		return n
	} else {
			return fibo(n-1) + fibo(n-2) // Aplica a fórmula recursiva de Fibonacci.
		}
	}

func dia(n int) {

	if n == 1{ // Aplica os números aos seus respectivos dias
		fmt.Println("O dia é Domingo.")
	} else if n == 2 {
		fmt.Println("O dia é Segunda-Feira.") 
	} else if n == 3 {
		fmt.Println("O dia é Terça-Feira.")
	} else if n == 4 {
		fmt.Println("O dia é Quarta-Feira.")
	} else if n == 5 {
		fmt.Println("O dia é Quinta-Feira.")
	} else if n == 6 {
		fmt.Println("O dia é Sexta-Feira.")
	} else if n == 7 {
		fmt.Println("O dia é Sábado.")
	} else {
		fmt.Println("Valor inválido.")
	}
}

func e_bissexto(n int) bool {
	if (n % 4 == 0 && n % 100 != 0) || (n % 400 == 0) { // Verifica se o ano cumpre os requisitos para ser bissexto.
		return true
	}
	return false
}

func main(){
	fmt.Println("----------------------------------------------------------")
	//Aplicação da primeira função:
	fmt.Println("Função 'e_primo': ")
	var num1 int
	fmt.Print("Digite um número inteiro positivo: ") 
	fmt.Scan(&num1) //Lê um número inteiro positivo.

	if e_primo(num1) { // Utiliza a função para determinar se é primo ou não.
		fmt.Printf("%d é um número primo.\n", num1) // Imprime o resultado verdadeiro.
	} else {
		fmt.Printf("%d não é um número primo.\n", num1) // Imprime o resultado falso.
		fmt.Printf("Divisores de %d:", num1)
		for i := 1; i <= num1; i++ { //Imprime os divisores do número inteiro positivo (caso o resultado seja falso.)
			if num1%i == 0 {
				fmt.Printf(" %d", i)
			}
		}
		fmt.Println()
	}

	fmt.Println("----------------------------------------------------------")
	//Aplicação da segunda função:
	fmt.Println("Função 'fibo': ")
	var num2 int
	fmt.Print("Digite um número inteiro positivo: ") 
	fmt.Scan(&num2) //Lê um número inteiro positivo.

	fmt.Println("O",num2,"° elemento da série de Fibonacci é",fibo(num2)) // Imprime o “n-ésimo” elemento da série de Fibonacci

	fmt.Println("----------------------------------------------------------")
	//Aplicação da terceira função:
	fmt.Println("Função 'dia': ")
	var num3 int
	fmt.Println("Digite o número correspondente ao dia da semana: ") //Lê um numero inteiro
	fmt.Scan(&num3)
	dia(num3) //Aplica a função dia que corresponde ao dia do número lido.

	fmt.Println("----------------------------------------------------------")
	//Aplicação da quarta função:
	fmt.Println("Função 'e_bissexto': ")
	var num4 int
	fmt.Print("Digite um ano: ")
	fmt.Scan(&num4) //Lê o ano

	if e_bissexto(num4) {
		fmt.Printf("%d é um ano bissexto.\n", num4) // Se a função retornar verdadeiro imprime o resultado.
	} else {
		fmt.Printf("%d não é um ano bissexto.\n", num4) // Se a função retornar falso imprime o resultado.
	}
}