package main

import "fmt"

type aluno struct {
	Nome string
	Sobrenome string
	RG int
	DataDeAdmissao string
}

func (a aluno) Imprimir(){
	fmt.Println("Imprimindo...")
	fmt.Printf("Nome: %v Sobrenome: %v RG: %v Data de admissão: %v\n", a.Nome, a.Sobrenome, a.RG, a.DataDeAdmissao)
}

func main (){
	aluno1 := aluno{"Rodney", "Martins", 234562345, "16/05/2022"}

	aluno1.Imprimir()
}