package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     uint8
	Status    bool
	cpf       string
}

//uso do método String
func (p Pessoa) String() string {
	return fmt.Sprintf("olá, meu nome é %s e eu tenho %d anos.",
		p.Nome, p.Idade)
}

type Categoria struct {
	Nome string
	Pai  *Categoria
}

//método sem referenciar ponteiro
func (c Categoria) HasParent() bool {
	return c.Pai != nil
}

//método referenciando ponteiro
func (c *Categoria) SetPai(pai *Categoria) {
	c.Pai = pai
}

func main() {

	p := Pessoa{"Carlos", "Silva", 38, true, "000.000.000-00"}
	p.cpf = "1"

	fmt.Println(p)

	//exemplo1 sem pai
	cat := Categoria{Nome: "Categoria 1"}
	if !cat.HasParent() {
		fmt.Println("não tem pai")
	}

	//exemplo2 com pai setado na mão
	cat = Categoria{Nome: "Categoria 1", Pai: &Categoria{Nome: "Pai1"}}
	if !cat.HasParent() {
		fmt.Println("não tem pai")
	} else {
		fmt.Println("o pai é", cat.Pai.Nome)
	}

	//exemplo3 com pai setado pelo método
	cat = Categoria{Nome: "Categoria 1"}
	cat.SetPai(&Categoria{Nome: "Pai2"})
	if !cat.HasParent() {
		fmt.Println("não tem pai")
	} else {
		fmt.Println("o pai é", cat.Pai.Nome)
	}
}
