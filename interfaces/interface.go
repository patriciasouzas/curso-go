package main

import "fmt"

type Documento interface {
	Doc() string
}

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
	Pessoa
	RazaoSocial string
	cnpj        string
}

//uso do método String
func (p Pessoa) String() string {
	return fmt.Sprintf("olá, meu nome é %s e eu tenho %d anos.",
		p.Nome, p.Idade)
}

func (pf PessoaFisica) Doc() string {
	return fmt.Sprintf("meu documento é %s", pf.cpf)
}

func (pf PessoaJuridica) Doc() string {
	return fmt.Sprintf("meu documento é %s", pf.cnpj)
}

func show(d Documento) {
	fmt.Println(d)
	fmt.Println(d.Doc())
}

func main() {
	pf := PessoaFisica{
		Pessoa{Nome: "Carlos", Idade: 19, Status: true},
		"Silveira",
		"000.000.000-00",
	}

	show(pf)

	pj := PessoaJuridica{
		Pessoa{Nome: "Arco-iris Bar", Idade: 1, Status: true},
		"Arco-iris Bar LTDA.",
		"000.000.0000/000-00",
	}

	show(pj)
}
