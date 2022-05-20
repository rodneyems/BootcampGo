package main

import "fmt"
// Só podemos utilizar um argumento Ellipsis e ele deve ser o último argumento da funcao. 
func soma(valores ...int) int32 {
	resultado := 0
	for _, valor := range valores {
		resultado += valor
	}
	return int32(resultado)
}

func main (){
	fmt.Println(soma(10,10,5,2))
}
