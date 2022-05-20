package main

import "fmt"

type Circulo struct {
	raio float64
}

func (c *Circulo) SetRaio(raio float64){
	c.raio = raio
}

// contrutores 
func NewCirculo(r float64) *Circulo {
	return &Circulo{raio: r}
}
func main (){
	circ := Circulo{}
	circ.SetRaio(5.54)

	fmt.Println(circ.raio)
}