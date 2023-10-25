package main

import (
	"fmt"
	"strconv"
)

const M = 7

func main() {
	lista, tam := [M]int{2, 1, 3, 9, 12}, 5
	lista2 := [M]int{1, 3, 5, 7, 10}
	lista3, tam3 := [M]int{1, 4}, 2

	fmt.Println(busca1(lista, 5, 17))
	fmt.Println(busca1(lista, 5, 9))

	fmt.Println(busca2(lista, 5, 17))
	fmt.Println(busca2(lista, 5, 9))

	fmt.Println(buscaOrd(lista2, 5, 7))
	fmt.Println(buscaOrd(lista2, 5, 6))

	fmt.Println(buscaBin(lista2, 5, 7))
	fmt.Println(buscaBin(lista2, 5, 6))

	insere(&lista, &tam, 9) // Elemento já existe
	fmt.Println(lista, tam)
	insere(&lista, &tam, 5) // Adiciona o elemento
	fmt.Println(lista, tam)
	insere(&lista, &tam, 16) // Adiciona o elemento
	fmt.Println(lista, tam)
	insere(&lista, &tam, 4) // Overflow
	fmt.Println(lista, tam)

	remove(&lista3, &tam3, 5) // Elemento não existe
	fmt.Println(lista3, tam3)
	remove(&lista3, &tam3, 4) // Remover o elemento
	fmt.Println(lista3, tam3)
	remove(&lista3, &tam3, 1) // Remover o elemento
	fmt.Println(lista3, tam3)
	remove(&lista3, &tam3, 1) // Underflow
	fmt.Println(lista3, tam3)
}

func removeMelhor(v *[M]int, n *int, valor int) string {
	if *n == 0 { return "Underflow" }

	ind := busca1(*v, *n, valor)
	if ind == -1 { return "Elemento não existe" }

	for i := ind; i < *n - 1; i++ {
		v[i] = v[i + 1]
	}
	v[*n - 1] = 0
	*n--
	return strconv.Itoa(valor)
}

func remove(v *[M]int, n *int, valor int) {
	if *n == 0 {
		fmt.Println("Underflow")
	} else {
		ind := busca1(*v, *n, valor)
		if ind == -1 {
			fmt.Println("Elemento não existe")
		} else {
			for i := ind; i < *n - 1; i++ {
				v[i] = v[i + 1]
			}
			v[*n - 1] = 0
			*n--
		}
	}
}

func insere(v *[M]int, n *int, valor int) {
	if *n == M {
		fmt.Println("Overflow")
	} else {
		if busca1(*v, *n, valor) == -1 {
			v[*n] = valor
			*n++
		} else {
			fmt.Println("Elemento já existe na tabela")
		}
	}
}

func buscaBin(v [M]int, n, x int) int {
	var meio int
	inf, sup := 0, n - 1
	for inf <= sup {
		meio = int((inf + sup) / 2)
		if v[meio] == x {
			return meio
		}
		if v[meio] < x {
			inf = meio + 1
		} else {
			sup = meio - 1
		}
	}
	return -1
}

func buscaOrd(v [M]int, n, x int) int {
	i := 0
	v[n] = x
	for v[i] < x {
		i++
	}
	if i == n + 1 || v[i] != x {
		return -1
	}
	return i
}

// Assumir que eu tenho pelo menos um espaço vazio na minha lista v
func busca2(v [M]int, n, x int) int {
	i := 0
	v[n] = x
	for v[i] != x {
		i++
	}
	if i != n {
		return i
	}
	return -1
}

func busca1(v [M]int, n, x int) int {
	i := 0
	for i < n {
		if v[i] == x {
			return i
		}
		i++
	}
	return -1
}