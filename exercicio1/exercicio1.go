package main

import "fmt"

func main() {
	palavra := "Teste"
	fmt.Printf("A palavra possui %v caracteres\n", len(palavra))
	for i := 0; i < len(palavra); i++{
		fmt.Println(string(palavra[i])) 
	}
}
