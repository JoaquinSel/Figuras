package figuras

import "math"

type Circulo struct {
	Radio float64
}

func (cir *Circulo) Area() float64 {
	return math.Pi * (cir.Radio * cir.Radio)
}

func (cir *Circulo) Perimetro() float64 {
	return math.Pi * 2 * cir.Radio
}
