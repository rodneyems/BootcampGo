package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func findMin(notas ...int32) int32 {
	min := notas[0]
	for _, nota := range notas {
		if nota < min {
			min = nota
		}
	}
	return min
}
func findMax(notas ...int32) int32 {
	max := notas[0]
	for _, nota := range notas {
		if nota > max {
			max = nota
		}
	}
	return max
}
func findAvg(notas ...int32) int32 {
	var avg int32 = 0
	for _, nota := range notas {
		avg += nota
	}
	avg = avg / int32(len(notas))
	return avg
}
func operation(operation string) (func(...int32) int32, error) {
	switch operation {
	case minimum:
		return findMin, nil
	case average:
		return findMax, nil
	case maximum:
		return findAvg, nil
	default:
		return nil, errors.New("operação inválida")
	}
}
func main() {
	minFunc, err := operation(minimum)
	if err != nil {
		fmt.Println("erro:", err)
	} else {
		minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
		fmt.Println("Valor mínimo:", minValue)
	}
	averageFunc, err := operation(average)
	if err != nil {
		fmt.Println("erro:", err)
	} else {
		averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("Valor máximo:", averageValue)
	}
	maxFunc, err := operation(maximum)
	if err != nil {
		fmt.Println("erro:", err)
	} else {
		maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("Valor médio:", maxValue)
	}
}
