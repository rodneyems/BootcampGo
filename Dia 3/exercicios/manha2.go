package main
import(
	"fmt"
	"errors"
)
func calculaMedia(notas ...int) (float64, error) {
	somaDasNotas := 0
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("a nota nao pode ser menor que zero")
		}
		somaDasNotas += nota
	}
	media := float64(somaDasNotas) / float64(len(notas))
	return media, nil
}

func main() {
	media, err := calculaMedia(10,10,5)
	if err != nil {
		fmt.Println("erro:", err)
	}else{
		fmt.Printf("A média calculada foi de: %.1f\n", media)
	}
}
