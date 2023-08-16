package main

import "fmt"

/* 1- Elabore um struct Contato, que deve conter os campos nome e e-mail. O struct deve conter um método para
alterar o e-mail.*/
type Contato struct {
	Nome	string
	Email	string
}
func(c *Contato) alteraEmail(email string){
	c.Email = email
}
/* 2- Elabore uma função adicionaContato, que recebe um contato e um array com até 5 elementos e inclui o
contato no primeiro elemento vazio do array.*/
func adicionaContato(nomeNovo string, emailNovo string, array[5]Contato) [5]Contato {
	for indice,valor := range array {
		if (valor == Contato{}) {
			array[indice] = Contato{Nome:nomeNovo, Email:emailNovo}
			break
		}
	}
	return array
}
/* 3- Elabore uma função excluiContato, que recebe um array de 5 elementos e elimina o último contato válido.
func excluiContato(array[5]Contato)*/
func excluiContato(array[5]Contato) [5]Contato{
	for indice,valor := range array {
		if (valor == Contato{}) {
			array[indice-1] = Contato{}
			break
		}
	}
	return array
}
/* 4- Elabore uma interface por linha de comando na função main, que cria um array de 5 elementos e permite
a inclusão ou exclusão de contatos. */
func main() {
	var listadeContatos [5]Contato

	x := 2

	for x != 0{
		fmt.Println("======================================")
		fmt.Println("          LISTA DE CONTATOS          ")
		fmt.Println("======================================")
		fmt.Println("Selecione a opção que deseja realizar:")
		fmt.Println("0 para sair, 1 para adicionar contato e 2 para excluir contato")
		fmt.Scanln(&x)

		if x == 0 {
			break
		} else if x == 1 {
			var nome string
			var email string
			fmt.Println("Insira o Nome: ")
			fmt.Scanln(&nome)
			fmt.Println("Insira o Email: ")
			fmt.Scanln(&email)
			listadeContatos = adicionaContato(nome, email , listadeContatos)
		} else if x == 2 {
			listadeContatos = excluiContato(listadeContatos)
		} else {
			fmt.Println("Opção Inválida.")
		}
		fmt.Println("------------------------------")
		fmt.Println(listadeContatos)
	}
}