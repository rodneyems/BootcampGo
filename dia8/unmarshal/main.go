package main

import (
	"encoding/json"
	"fmt"
)

type ser struct {
	Name  string
	Idade int
}

func main() {
	jsonData := []byte(`[
		{"Name": "Rodney", "Idade": 20},
		{"Name": "Martins", "Idade": 25},
		{"Name": "Tedi", "Idade": 15}
	]`)
	seres := []ser{}

	err := json.Unmarshal(jsonData, &seres)
	fmt.Println(err)
	for _, v := range seres {
		fmt.Println("\nNOVO")
		fmt.Println(v)
	}
}
