package main

import (
	"fmt"
	"strings"
)

func calculaSalario(categoria string, horasTrabalhadas uint) float64 {
	switch strings.ToUpper(categoria) {
	case "A":
		salario := float64(horasTrabalhadas * 3000.00)
		if horasTrabalhadas > 160 {
			return salario * 1.5
		} else {
			return salario
		}
	case "B":
		salario := float64(horasTrabalhadas * 1500.00)
		if horasTrabalhadas > 160 {
			return salario * 1.2
		} else {
			return salario
		}
	case "C":
		return float64(horasTrabalhadas * 1000.00)
	default:
		return 0
	}
}

func main() {
	fmt.Printf("Funcionário A, Salário $ %.2f\n", calculaSalario("A", 172))
	fmt.Printf("Funcionário B, Salário $ %.2f\n", calculaSalario("B", 176))
	fmt.Printf("Funcionário C, Salário R$ %.2f\n", calculaSalario("C", 162))
}
