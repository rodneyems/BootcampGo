package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	fmt.Println("Benjamin,",employees["Benjamin"], "anos")
	fmt.Println("\nLista com funcionários com mais de 21 anos:")
	for k, v := range employees{
		if v > 21 {
			fmt.Printf("%v, %v anos\n", k, v)
		}
	}
	employees["Federico"] = 21
	delete(employees, "Pedro")
	fmt.Println("\n\nLista completa atualizada:")
	for k, v := range employees{
		fmt.Printf("%v, %v anos\n", k, v)
	}
}
