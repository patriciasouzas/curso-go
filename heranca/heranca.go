package main

import "fmt"

type Pessoa struct {
	Nome   string
	Idade  uint8
	Status bool
}

type PessoaFisica struct {
	Pessoa
	Sobrenome string
	cpf       string
}

type PessoaJuridica struct {
	RazaoSocial string
	cpf         string
}

//uso do método String
func (p Pessoa) String() string {
	return fmt.Sprintf("olá, meu nome é %s e eu tenho %d anos.",
		p.Nome, p.Idade)
}

func main() {
	p := PessoaFisica{
		Pessoa{Nome: "Carlos", Idade: 19, Status: true},
		"Silveira",
		"000.000.000-00",
	}

	fmt.Println(p)
}
