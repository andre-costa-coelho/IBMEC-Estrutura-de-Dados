package main

import (
	"fmt"
	"AC3/contatos"
)

func main() {
	var listadeContatos [5]contatos.Contato

	x := 2

	for x != 0{
		fmt.Println("======================================")
		fmt.Println("          LISTA DE CONTATOS          ")
		fmt.Println("======================================")
		fmt.Println("Selecione a opção que deseja realizar:")
		fmt.Println("0 para sair, 1 para adicionar contato, 2 para excluir contato, 3 para imprimir a lista de contatos e 4 para editar o email de um contato:")
		fmt.Scanln(&x)

		switch x {
		case 0:
			break
		case 1:
			var nome string
			var email string
			fmt.Println("Insira o Nome: ")
			fmt.Scanln(&nome)
			fmt.Println("Insira o Email: ")
			fmt.Scanln(&email)
			contatos.AdicionaContato(nome, email ,&listadeContatos)
			break
		case 2:
			contatos.ExcluiContato(&listadeContatos)
			break
/* 4-Atualize a interface por linha de comando para incluir a opção de exibir todos os contatos salvos.
Cada contato deve ser exibido em uma linha, precedido pelo número do índice daquele elemento.*/
		case 3:
			for i := 0; i < len(&listadeContatos); i++{
				fmt.Println("Índice do Contato:",i)
				fmt.Println(listadeContatos[i])
			}
/*5- Atualize a interface por linha de comando para poder editar o e-mail de um contato previamente saldo.
A interface deve exibir os contatos armazenados e pedir para o usuário indicar o índice do contato que
quer ser alterado. Em seguida, o programa pede o novo e-mail e chama a função editaEmail, implementada
no exercício 3.*/
		case 4:
			var novoEmail string
			var indiceEmail int
			fmt.Println("Insira o novo email: ")
			fmt.Scanln(&novoEmail)
			fmt.Println("Insira o índice do contato que deseja alterar: ")
			fmt.Scanln(&indiceEmail)
			contatos.EditaEmail(indiceEmail, novoEmail, &listadeContatos)
		default:
			fmt.Println("Opção Inválida.")
			break
		}
		fmt.Println("------------------------------")
	}
}