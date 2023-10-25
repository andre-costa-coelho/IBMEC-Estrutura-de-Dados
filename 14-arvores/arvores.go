package main

import "fmt"

type No struct {
	Chave  int
	Esq, Dir *No
	Altura int
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
		return  nil
	}

	return buscaPreOrdem(a.Raiz, valor)
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

func buscaPosOrdem(n *No, valor int, nEcontrado **No) {
	if n.Esq != nil {
		buscaPosOrdem(n.Esq, valor, nEcontrado)
	}
	if n.Dir != nil {
		buscaPosOrdem(n.Dir, valor, nEcontrado)
	}
	if n.Chave == valor {
		*nEcontrado = n
	}
}

func buscaPreOrdem(n *No, valor int) *No {
	if n.Chave == valor {
		return n
	}
	var noProcurado *No

	if n.Esq != nil {
		noProcurado = buscaPreOrdem(n.Esq, valor)
	}
	if n.Dir != nil {
		noProcurado = buscaPreOrdem(n.Dir, valor)
	}
	return noProcurado
}

func calculaAltura(n *No) {
	if n.Esq != nil {
		calculaAltura(n.Esq)
	}
	if n.Dir != nil {
		calculaAltura(n.Dir)
	}
	calculaAltura(n)
}

func calculaAlturaNo(n *No) {
	var altE, altD int
	if n.Esq != nil {
		altE = 0
	} else {
		altE = n.Esq.Altura
	}

	if n.Dir != nil {
		altD = 0
	} else {
		altD = n.Dir.Altura
	}

	if altE > altD {
		n.Altura = altE + 1
	} else {
		n.Altura = altD + 1
	}
}
