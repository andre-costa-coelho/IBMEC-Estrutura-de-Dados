package main

import (
	"fmt"
)

const M = 3

func main() {
  var fila [M]int 
  in, fim := -1,-1

  insereFila(&fila,&in,&fim,4)
  fmt.Println(fila)
  insereFila(&fila,&in,&fim,5)
  fmt.Println(fila)
  insereFila(&fila,&in,&fim,2)
  fmt.Println(fila)
  insereFila(&fila,&in,&fim,1 ) // Overflow
  fmt.Println(fila)

  removeFila(&fila,&in,&fim)
  fmt.Print(fila)
  removeFila(&fila,&in,&fim)
  fmt.Print(fila)
  removeFila(&fila,&in,&fim)
  fmt.Print(fila)
  removeFila(&fila,&in,&fim) // Underflow
  fmt.Print(fila)

}
func removeFila(f*[M]int, in * int, fim * int){
  if *in == *fim && *fim == -1{
    fmt.Println("Underflow")
    return
  }

  f[*in] = 0
  if *in == *fim {
    *in, *fim = -1,-1
  } else {
    *in = (*in + 1) % M
  }
}
func insereFila(f * [M]int, in *int, fim *int, valor int){
  if (*fim + 1) % M == *in {
    fmt.Println("Overflow")
    return
  }

  if *in == *fim && *fim == -1 {
    *in, *fim = 0,0
  } else {
    *fim = (*fim + 1) % M
  }
  f[*fim] = valor
  
}