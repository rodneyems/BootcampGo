package main

import "fmt"

func isEven(num int32){
	defer func(){
		err := recover() // funciona como um catch para o panic. Bom momento para fechar conexoes etc.

		fmt.Println("ERRO:", err)
	}()

	if num == 2 {
		panic("LASCOU")
	} else {
		fmt.Println("Nao lascou")
	}
}

func main(){
	isEven(2)
}