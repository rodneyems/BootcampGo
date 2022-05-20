package main

import "fmt"

func calculaImposto(salario float64, imposto float64) float64{
	return salario * (imposto / 100)
}
func main(){
	fmt.Printf("Para o funcionário que ganha R$ 50.000,00, o valor de imposto cobrado será: R$ %.2f \n", calculaImposto(50000, 17))
	fmt.Printf("Para o funcionário que ganha R$ 150.000,00, o valor de imposto cobrado será: R$ %.2f\n", calculaImposto(150000, 10))
}