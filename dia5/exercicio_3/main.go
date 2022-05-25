package main

import (
	"fmt"

	et "github.com/rodneyems/BootcampGo/dia5/exercicio_3/etities"
)

func main() {
	var maintenances = []et.Maintenance{
		{Name: "Montar", Price: 200.00},
		{Name: "Desmontar", Price: 100.00},
		{Name: "Quebrar", Price: 300.00},
	}
	var services = []et.Service{
		{Name: "Jardinagem", Price: 150.00, Time: 60},
		{Name: "Limpeza", Price: 100.00, Time: 60},
		{Name: "Desenvolvimento", Price: 2000.00, Time: 60},
	}
	var products = []et.Product{
		{Name: "Mouse", Price: 100.00, Quantity: 1},
		{Name: "Monitor", Price: 100.00, Quantity: 1},
		{Name: "Teclado", Price: 100.00, Quantity: 1},
	}
	c1 := make(chan float64)
	go et.CalcMain(maintenances, c1)
	c2 := make(chan float64)
	go et.CalcPro(products, c2)
	c3 := make(chan float64)
	go et.CalcSer(services, c3)

	fmt.Println("Total:", <-c1 + <-c2 + <-c3)
}
