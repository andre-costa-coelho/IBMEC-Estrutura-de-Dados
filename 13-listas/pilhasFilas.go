package main

import (
	"fmt"
)

const M = 5

func main() {
	var pilha [M]int
  topo := 0

  insere(&pilha, &topo, 4)
  fmt.Println(pilha)
  insere(&pilha, &topo, 5)
  fmt.Println(pilha)
  insere(&pilha, &topo, 2)
  fmt.Println(pilha)
  insere(&pilha, &topo, 1) //Overflow
  fmt.Println(pilha)

  fmt.Println(remove(&pilha, &topo))
  fmt.Println(pilha)
  fmt.Println(remove(&pilha, &topo))
  fmt.Println(pilha)
  fmt.Println(remove(&pilha, &topo))
  fmt.Println(pilha)
  fmt.Println(remove(&pilha, &topo)) //Underflow
  fmt.Println(pilha)
  
}

func insere(p* [M]int, topo *int, valor int){
  if *topo == M{
    fmt.Println("overflow")
    return
  }

  p[*topo] = valor
  *topo++
}

func remove(p*[M]int, topo *int) int{
  if *topo == 0 {
    fmt.Println("underflow")
    return -1
  }

  *topo--
  valorRemovido := p[*topo]
  p[*topo] = 0
  return valorRemovido
}