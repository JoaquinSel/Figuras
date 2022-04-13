package figuras

import "fmt"

type Geometria interface {
	Area() float64
	Perimetro() float64
}

func Medidas(g Geometria) {
	fmt.Println("Medidas: ", g)
	fmt.Println("Area: ", g.Area())
	fmt.Println("Perimetro: ", g.Perimetro())
}
