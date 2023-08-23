/* 1-Organize o projeto em pacotes: um pacote para o tipo Contato, um para a operação sobre os arrays
e um arquivo main.go.*/

package contatos

/* 2-Refatore as funções que retornam arrays para que elas utilizem ponteiros.*/
func AdicionaContato(nomeNovo string, emailNovo string, array*[5]Contato) {
	for indice,valor := range array {
		if (valor == Contato{}) {
			array[indice] = Contato{Nome:nomeNovo, Email:emailNovo}
			break
		}
	}
}

func ExcluiContato(array*[5]Contato) {
	for indice,valor := range array {
		if (valor == Contato{}) {
			array[indice-1] = Contato{}
			break
		}
	}
}

/* 3-Crie uma nova função editaEmail, que recebe o índice do elemento no array e altera o e-mail
 do elemento indicado.*/
 func EditaEmail(indice int, novoEmail string, array*[5]Contato){
	if indice >= 0 && indice < len(array) {
		array[indice].alteraEmail(novoEmail)
	}
}