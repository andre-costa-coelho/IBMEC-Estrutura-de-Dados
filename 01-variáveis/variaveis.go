package main

import "fmt"

var linguagem string
var soma bool

func main(){
	//declaração explícita do tipo
	var msg string
	var n1,n2 int
	// declaração implicita
	var n3,n4 = 15,25
	//declaração curta
	msg2:= "olá mundo"

	msg = "oi"
	n1 = 10
	n2 = 20

	fmt.Println(msg,n1,n2,n3,n4,msg2)
}

/*
inteiros:

int8	-128 a 127
uint8	0 a 255
int16	-32.768 a 32.767
uint16	0 a 65.535
int32
uint32
int64
uint64

byte	uint8
rune	int32	caractéres unicode

--------------------------------------------

flutuantes:

float32	precisão de 6 a 9 dígitos
float64	precisão de 15 a 17 dígitos

é necessário especificar o tamanho do float
caso a declaração seja implícita, é assumido float64

--------------------------------------------

complexos:

complex64
complex128

--------------------------------------------

string
bool	true / false

nil		ausência de dados

*/
