/* 1-Organize o projeto em pacotes: um pacote para o tipo Contato, um para a operação sobre os arrays
e um arquivo main.go.*/

package contatos

type Contato struct {
	Nome	string
	Email	string
}

func(c *Contato) alteraEmail(email string){
	c.Email = email
}
