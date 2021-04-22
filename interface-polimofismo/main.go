package main

import "fmt"

type Pessoa struct {
	name  string
	idade int
}

type Dentista struct {
	Pessoa
	CRO string
}

type Arquiteto struct {
	Pessoa
	CRA string
}

func (a Dentista) hello() {
	fmt.Println(a.name, a.CRO)
}

func (a Arquiteto) hello() {
	fmt.Println(a.name, a.CRA)
}

type Profissao interface {
	hello()
}

func profissao(p Profissao) {
	p.hello()
}

func main() {
	dentista := Dentista{
		Pessoa: Pessoa{
			name:  "Elisson",
			idade: 26,
		},
		CRO: "123456",
	}
	arquiteto := Arquiteto{
		Pessoa: Pessoa{
			name:  "Eduarda",
			idade: 26,
		},
		CRA: "11223344",
	}

	profissao(dentista)
	profissao(arquiteto)
}
