package main

import "fmt"

func main() {
nome := "Rodney"
idade := 23
salario := 100001
tempoDeCasa := 2

if idade > 22 && tempoDeCasa > 1{
	if salario > 100000{
		fmt.Printf("Parabéns %v, você esta apto à receber o empréstimo sem nenhuma taxa de juros!\n", nome)
	} else {
		fmt.Printf("Parabéns %v, você esta apto à receber o empréstimo com uma pequena taxa de juros!\n", nome)
	}
}else {
	fmt.Printf("Desculpe %v, infelizmente você é inapto à receber o empréstimo\n", nome)
}
}
