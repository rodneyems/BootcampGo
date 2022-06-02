package main

import "fmt"

type SalaryError struct {
	Level int
	Msg   string
}

func (e *SalaryError) Error() string {
	return fmt.Sprintf("%v - %v", e.Level, e.Msg)
}

func CheckSalary(salary *int) error {
	if *salary < 15000 {
		return &SalaryError{
			Level: 2,
			Msg: "salario menor que 15000. favor trabalhar mais",
		}
	}
	return nil
}

func main() {
	salary := 15000

	err := CheckSalary(&salary)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Ótimo salario!")
	}
}
