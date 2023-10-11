package main

import (
	"fmt"
)

type No struct {
	chave int
	prox  *No
}

type ListaEncadeada struct {
	cabeca *No
}

// exibição da lista:
func (l *ListaEncadeada) exibe(){
	no := l.cabeca

	for no != nil {
		fmt.Println(no)
		no = no.prox
	}
}

// busca simples:
func (l *ListaEncadeada) buscaSimples(chave int) *No {
	no := l.cabeca

	for no != nil {
		if no.chave == chave {
			return no
		}
		no = no.prox
	}

	return nil
}
// inserção simples, ao final da lista:
func (l *ListaEncadeada) insereSimples(chave int) {
	novoNo := &No{chave: chave}
	no:= l.cabeca

	if no == nil {
		l.cabeca = novoNo
	} else {
		for no.prox != nil {
			no = no.prox
		}
		no.prox = novoNo
	}
}
// considerando uma lista é ordenada
// busca de uma lista ordenada:
func (l *ListaEncadeada) buscaOrd(chave int) (*No, *No) {
	var ant *No
	no := l.cabeca
	if no == nil {
		return ant, no
	}
	for no.chave < chave {
		ant = no
		no = no.prox
		if no == nil {
			return ant, no
		}
	}
	if no.chave == chave{
		return ant, no
	} 
return ant, nil		
}
// inserção de uma lista ordenada:
func (l *ListaEncadeada) insereOrd(chave int) {
	ant, no := l.buscaOrd(chave)

	if no != nil {
		return
	}
	novoNo := &No{chave: chave}
	if ant == nil {
		novoNo.prox = l.cabeca
		l.cabeca = novoNo
	} else {
		novoNo.prox = ant.prox
		ant.prox = novoNo
	}	
}	


// remoção de uma lista ordenada:
func (l *ListaEncadeada) remove(chave int) *No {
	ant, no := l.buscaOrd(chave)

	if no == nil {
		return nil
	}
	if ant == nil {
		l.cabeca = no.prox
	} else {
		ant.prox = no.prox
	}
	return no
}

func main(){
	var l ListaEncadeada

	fmt.Println(l)

	//l.insereSimples(2)
	//l.insereSimples(4)
	//l.insereSimples(6)

	l.insereOrd(2)
	l.insereOrd(4)
	l.insereOrd(6)
	l.insereOrd(3)
	l.insereOrd(1)

	//l.exibe()
	//fmt.Println(l.buscaOrd(2))
	//fmt.Println(l.buscaOrd(1))
	//fmt.Println(l.buscaOrd(8))
	//fmt.Println(l.buscaOrd(5))

	l.remove(5)
	l.remove(3)
	l.remove(1)
	l.remove(2)
	l.exibe()
	//fmt.Println(l.buscaSimples(4))
	//fmt.Println(l.buscaSimples(2))
	//fmt.Println(l.buscaSimples(6))
}