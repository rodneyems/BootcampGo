package main

import (
	"fmt"
)

type SalaryError struct {
	Level int
	Msg   string
}

func CheckSalary(salary *int) error {
	if *salary < 15000 {
		return fmt.Errorf("erro: o mínimo tributável é 15.000 e o salário informado é: %v", *salary)
	}
	return nil
}

func main() {
	salary := 1500

	err := CheckSalary(&salary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ótimo salario!")
	}
}
