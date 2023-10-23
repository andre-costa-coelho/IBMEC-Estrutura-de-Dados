## O que é uma Lista Encadeada Circular?
É uma lista encadeada na qual observamos uma estrutura de um nó apontando para o próximo, só que o último nó apontando para o primeiro, sendo assim formando um ciclo.

### Exibição dos nós em uma lista circular:
```sh
PROGRAMA exibirNos (lista) {
nó_atual := lista.primeiro
    ENQUANTO nó_atual != *NULL* 
    {
        ESCREVE(nó_atual.dado)
        nó_atual := nó_atual.proximo
        SE nó_atual == lista.primeiro 
        {
            PARA
        }
    }
ENCERRA
}
```
### Inserção de um nó no início da lista:
```sh
PROGRAMA insereInicio (lista,dado) {
novoNo := No(dado)
    no := lista.primeiro
    SE no == *NULL* 
    {
       lista,primeiro = novoNo
       novoNo.proxima = novoNo
    } SENÃO {
        PARA no.proximo != lista.primeiro 
        {
          no = no.proximo
       }
       no.proxima = novoNo
       novoNo.proxima = lista.primeiro
       lista.primeiro = novoNo
    }
 ENCERRA
}
```
### Exclusão do primeiro nó da lista:
```sh
PROGRAMA excluiPrimeiroNo (lista) {
no := lista.primeiro
	SE no == nil 
    {
		RETORNA
	}

	SE no.proxima == lista.primeiro 
    {
		lista.primeiro = nil
		RETORNA
	}

	ENQUANTO no.proxima != listam.primeiro 
    {
		no = no.proxima
	}
	no.proxima = lista.primeiro.proxima
	lista.primeiro = lista.primeiro.proxima
 ENCERRA
}
```