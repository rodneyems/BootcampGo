package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Produto struct {
	id         int
	preco      float32
	quantidade int
}

func main() {
	file, _ := os.Open("listaDeProdutos.csv")
	csvRead := csv.NewReader(file)
	defer file.Close()
	data, _ := csvRead.ReadAll()
	soma := 0.0
	fmt.Println("ID\tPreco\tQuantidade")
	for _, linhas := range data {
		for _, linha := range linhas {
			l := strings.Split(linha, ";")
			fmt.Printf("%v\t%v\t%v\n", l[0], l[1], l[2])
			preco, _ := strconv.ParseFloat(l[1],64)
			soma += preco
		}
	}
	fmt.Printf("\t%v\n", soma)

}
