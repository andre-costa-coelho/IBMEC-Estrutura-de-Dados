package main

import "fmt"

type No struct {
	Chave  int
	Esq, Dir *No
}

type Arvore struct {
	Raiz *No
}

func main (){
	arv := Arvore{}

	n1 := &No{Chave: 2}
	n2 := &No{Chave: 5}
	n3 := &No{Chave: 1}
	n4 := &No{Chave: 9}
	n5 := &No{Chave: 0}

	arv.Raiz = n1
	n1.Esq = n2
	n2.Esq = n3
	n2.Dir = n4
	n1.Dir = n5

	fmt.Println(arv)
}

func imprimeArvore(a Arvore){
	if a.Raiz != nil {
		//imprimePreOrdem(a.Raiz)
		//imprimeSimetrico(a.Raiz)
		imprimePosOrdem(a.Raiz)
	}
}

func imprimePreOrdem(n *No) {
	fmt.Println(n)
	if n.Esq != nil {
		imprimePreOrdem(n.Esq)
	}
	if n.Dir != nil {
		imprimePreOrdem(n.Dir)
	}
}

func imprimeSimetrico(n *No) {
	if n.Esq != nil {
		imprimeSimetrico(n.Esq)
	}
	fmt.Println(n)
	if n.Dir != nil {
		imprimeSimetrico(n.Dir)
	}
}

func imprimePosOrdem(n *No) {
	if n.Esq != nil {
		imprimePosOrdem(n.Esq)
	}
	if n.Dir != nil {
		imprimePosOrdem(n.Dir)
	}
	fmt.Println(n)
}

// impriveArvore(arv)

func buscaArvore(a Arvore, valor int) *No {
	if a.Raiz != nil {
		return nil
	}

	var noProcurado *No
	buscaSimetrico(a.Raiz, valor, &noProcurado)
	return noProcurado
}

func buscaSimetrico(n *No, valor int, nEcontrado **No) {
	if n.Esq != nil {
		buscaSimetrico(n.Esq, valor, nEcontrado)
	}
	if n.Chave == valor {
		*nEcontrado = n
	}
	if n.Dir != nil {
		buscaSimetrico(n.Dir, valor, nEcontrado)
	}
}