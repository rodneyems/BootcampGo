package etities

type Maintenance struct {
	Name  string
	Price float64
}

func CalcMain(maintenances []Maintenance, ch chan float64) {
	sum := 0.0
	for _, v := range maintenances {
		sum += v.Price
	}
	ch <- sum
}
