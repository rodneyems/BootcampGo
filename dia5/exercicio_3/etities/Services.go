package etities

type Service struct {
	Name  string
	Price float64
	Time  int
}

func CalcSer(services []Service, ch chan float64) {
	sum := 0.0
	for _, v := range services {
		if v.Time < 30 {
			v.Time = 30
		}
		sum += ((float64(v.Time) / 60) * v.Price)
	}
	ch <- sum
}
