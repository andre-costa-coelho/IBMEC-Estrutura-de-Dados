package main

import "fmt"

const M = 5

func main() {
	var lista [M]int
	var n = 0

	insereOrd(&lista, &n, 4)
	fmt.Println(lista)

	insereOrd(&lista, &n, 12)
	fmt.Println(lista)

	insereOrd(&lista, &n, 2)
	fmt.Println(lista)

	insereOrd(&lista, &n, 6)
	fmt.Println(lista)

	insereOrd(&lista, &n, 17)
	fmt.Println(lista)

	insereOrd(&lista, &n, 1)
	fmt.Println(lista)
}

func insereOrd(v *[M]int, n *int, novoValor int) {
  if *n == M{
    fmt.Println("Overflow")
    return
  }
  fmt.Println("Tentando inserir", novoValor)
  aux := novoValor
  for i := 0; i <*n; i++ {
    if v[i] == novoValor {
      fmt.Println("Valor já existe")
      return
    }
    if v[i] > novoValor{
      v[i], aux = aux, v[i]
    } 
  }
  fmt.Println(novoValor, "Não encontrado, inserindo na lista")
  v[*n] = aux
  *n++
}