package main

import (
	"errors"
	"fmt"
)

func Salary(hours int, price float64) (float64, error){
	if hours < 80 {
		err2 := errors.New("erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês")
		return 0.0, fmt.Errorf("erro 1: %v", err2)
	}
	salary := float64(hours) * price
	if salary > 20000{
		return salary * 0.9, nil
	} else {
		return salary, nil
	}
}

func main() {
	s, err := Salary(10, 100)
	if err != nil {
		fmt.Println(errors.Unwrap(err))
	} else {
		fmt.Println("Seu salario é:", s)
	}
}
