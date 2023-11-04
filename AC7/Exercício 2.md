## O que é uma lista duplamente encadeada?
É um tipo de lista encadeada que permite a navegação em ambas as direções da lista, devido que cada nó mantém uma referência tanto para o nó seguinte quanto para o nó anterior na lista.

### Exibição dos nós em uma lista duplamente encadeadas:
```sh
PROGRAMA exibeListaDup (lista){
 no := lista.primeiro
    ENQUANTO no != NULL 
    {
        ESCREVE (no)
        no = no.próximo 
}
 ENCERRA
}
```
### Busca de um nó:
```sh
PROGRAMA buscaListaDup (lista,ch){
DECLARA variável  ant *No
   no := lista.primeiro
   SE no == NULL 
   {
     RETORNA ant, no
}
ENQUANTO no.chave < ch 
{
   ant = no
   no = no.proximo
  SE no == NULL 
  {
     RETORNA ant, no 
  }
}
SE no.chave == ch 
{
   RETORNA ant, no
}
RETORNA ant, NULL
 ENCERRA 
}
```
### Inserção de um nó:
```sh
PROGRAMA insereListaDup (lista,ch){
ant, no := busca.no(ch)
        SE no != NULL
        {
            RETORNA
        }

        novo.no := &no {chave: ch}

        SE lista.primeiro == NULL
        {
            lista.primeiro = novo.no
            RETORNA
        }
        SE ant == NULL
        {
            novo.no.proximo = lista.primeiro
            lista.primeiro.anterior = novo.no
            lista.primeiro = novo.no
            RETORNA
        }
        novo.no.proximo = anterior.proximo
	    anterior.proximo = novo.no
	    novo.no.anterior = ant

	    SE novo.no.proximo != NULL
        {
		    novo.no.proximo.anterior = novo.no
	    } 
    ENCERRA
    }
```
### Remoção de um nó:
```sh
PROGRAMA removeListaDup (lista,ch){
    no := lista.primeiro
    ENQUANTO no != NULL 
    {
      SE no.chave == ch 
      {
         SE no.anterior == NULL E no.proximo == NULL 
         {
                   lista.primeiro = NULL
                   PARA
           }
SE no.anterior == NULL 
{
       lista.primeiro = no.proximo
       no.proximo.anterior = NULL
       PARA
}
SE no.proximo == NULL 
{
  no.amterior.proximo = NULL
  PARA
}
no.anterior.proximo = no.proximo
no.proximo.anterior = no.anterior

PARA
}
no = no.proximo
}
ENCERRA
}
```