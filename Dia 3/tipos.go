package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome      string `json:"nome_json"` // aqui será um label para o json caso contrario sai o nome normal
	Genero    string
	Idade     int
	Profissao string
	Peso      int
	Gostos    Infos
}
type Infos struct {
	Comida  string
	Esporte string
	Lazer   string
}

func main() {
	rodney := Pessoa{
		Nome:      "Rodney",
		Genero:    "M",
		Idade:     28,
		Profissao: "Dev",
		Peso:      81,
	}
	fmt.Println(rodney)
	martinha := Pessoa{}
	martinha.Nome = "Martinha"
	fmt.Println(martinha)

	var tedi Pessoa
	tedi.Genero = "Cachorrao kkkk"
	tedi.Gostos.Comida = "Churrasco"
	fmt.Println(tedi)

	meuJSON, _ := json.Marshal(rodney)
	fmt.Println(string(meuJSON))
}
