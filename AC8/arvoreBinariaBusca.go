package main

import "fmt"

type No struct {
	Chave   int
	Esquerda *No
	Direita *No
}

type Arvore struct {
    Raiz *No
}


func (t *Arvore) insereArvore(ch int) {
    novoNo := &No{Chave: ch}

    if t.Raiz == nil {
        t.Raiz = novoNo
    } else {
        insereNo(t.Raiz, novoNo)
    }
}

func insereNo(no, novoNo *No) {
    if novoNo.Chave < no.Chave {
        if no.Esquerda == nil {
            no.Esquerda = novoNo
        } else {
            insereNo(no.Esquerda, novoNo)
        }
    } else {
        if no.Direita == nil {
            no.Direita = novoNo
        } else {
            insereNo(no.Direita, novoNo)
        }
    }
}

func (t *Arvore) buscaArvore(ch int) bool {
	return buscaNo(t.Raiz, ch)
}


func buscaNo(no *No, ch int) bool {
	if no == nil {
		return false
	}

	if ch < no.Chave {
		return buscaNo(no.Esquerda, ch)
	} else if ch > no.Chave {
		return buscaNo(no.Direita, ch)
	}

	return true
}

func main() {
    arvore := &Arvore{}

    // Inserir elementos na árvore
    chaves := []int{10, 5, 15, 3, 7, 12, 18}
    for _, chave := range chaves {
        arvore.insereArvore(chave)
    }

    // Verificar se um valor existe na árvore
    buscaChave := 7
    if arvore.buscaArvore(buscaChave) {
        fmt.Printf("%d está na árvore.\n", buscaChave)
    } else {
        fmt.Printf("%d não está na árvore.\n", buscaChave)
    }

    buscaChave = 9
    if arvore.buscaArvore(buscaChave) {
        fmt.Printf("%d está na árvore.\n", buscaChave)
    } else {
        fmt.Printf("%d não está na árvore.\n", buscaChave)
    }
}