package main

import (
	"fmt"
	"os"
)

type Produto struct {
	id         int
	preco      float32
	quantidade int
}

func main() {
	produtos := []Produto{}
	p1 := Produto{1, 10.90, 100}
	p2 := Produto{2, 20.00, 200}
	p3 := Produto{3, 15.00, 300}

	produtos = append(produtos, p1, p2, p3)

	file, err := os.OpenFile("listaDeProdutos.csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	defer file.Close()
	if err != nil {
		fmt.Println("Erro:", err)
	}
	for _, v := range produtos {
		s := fmt.Sprintf("%v;%v;%v\n", v.id, v.preco, v.quantidade)
		file.WriteString(s)
	}
}
