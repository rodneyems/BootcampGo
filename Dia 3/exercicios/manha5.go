package main

import (
	"fmt"
)

const (
	dog = "dog"
	cat = "cat"
	ham = "ham"
	spi = "spi"
)

func Animal(animal string) (func(int) float64, func() string) {
	switch animal {
	case dog:
		return calcDog, nil
	case cat:
		return calcCat, nil
	case ham:
		return calcHam, nil
	case spi:
		return calcSpi, nil
	default:
		return nil, displayMsg
	}
}
func displayMsg() string {
	return "O animal não existe"
}
func calcDog(amount int) float64 {
	return float64(amount) * 10000
}
func calcCat(amount int) float64 {
	return float64(amount) * 5000
}
func calcHam(amount int) float64 {
	return float64(amount) * 250
}
func calcSpi(amount int) float64 {
	return float64(amount) * 150
}

func main() {
	var amount float64
	animalDog, msg := Animal(dog)
	if msg != nil {
		fmt.Println(msg())
	}
	animalCat, msg := Animal(cat)
	if msg != nil {
		fmt.Println(msg())
	}
	animalSpi, msg := Animal(spi)
	if msg != nil {
		fmt.Println(msg())
	}
	amount += animalDog(5)
	amount += animalCat(8)
	amount += animalSpi(2)
	fmt.Println("Será necessário comprar", amount/1000, "Kg de ração")
}
