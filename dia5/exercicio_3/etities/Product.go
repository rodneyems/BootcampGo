package etities

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func CalcPro(products []Product, ch chan float64) {
	sum := 0.0
	for _, v := range products {
		sum += (float64(v.Quantity) * v.Price)
	}
	ch <- sum
}
