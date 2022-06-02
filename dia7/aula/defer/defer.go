package main

import "fmt"

type Dog struct {
	Name string
}

func (d *Dog) late() {
	fmt.Println(d.Name, "AUAU")
}

func main() {
	dog := &Dog{"Tedi"}
	dog2 := Dog{"Tedi2"}
	if true {
		defer dog.late()
	}
	defer dog2.late()
}
