package main

import "fmt"

type Produto interface {
	CalcularCusto()
}

type Ecommerce interface {
	Total()
	Adicionar()
}

type produto struct {
	Tipo  string
	Nome  string
	Preco float32
}

type loja struct {
	Produtos []produto
}

func novoProduto(tipo string, nome string, preco float32) produto {
	novoProduto := produto{
		tipo,
		nome,
		preco,
	}
	return novoProduto
}

func novaLoja(produtos ...produto) loja {

	novaLoja := loja{
		produtos,
	}
	return novaLoja
}

func (p produto) CalcularCusto() float32 {
	switch p.Tipo {
	case "p":
		return 0
	case "m":
		return p.Preco * 0.03
	case "g":
		return p.Preco*0.06 + 2500
	default:
		return 0
	}
}

func (l loja) Total() float32{
	var total float32 = 0.0
	for _, v := range l.Produtos{
		total = total + v.Preco + v.CalcularCusto()
	}
	return total
}

func (l *loja) Adicionar(p produto){
	l.Produtos = append(l.Produtos, p)
}

func main() {
	p1 := novoProduto("p", "mouse",10.00)
	p2 := novoProduto("m", "monitor",100.00)
	p3 := novoProduto("g", "CPU",100.00)
	l1 := novaLoja(p1, p2)

	fmt.Println(l1.Total())

	l1.Adicionar(p3)
	fmt.Println(l1.Total())
}
