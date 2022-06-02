package main

import (
	"errors"
	"fmt"
)

type SalaryError struct {
	Level int
	Msg   string
}

func CheckSalary(salary *int) error {
	if *salary < 15000 {
		return errors.New("salario menor que 15000. favor trabalhar mais")
	}
	return nil
}

func main() {
	salary := 1500

	err := CheckSalary(&salary)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Ótimo salario!")
	}
}
