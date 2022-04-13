package figuras

type Cuadrado struct {
	Lado float64
}

func (cua *Cuadrado) Area() float64 {

	return cua.Lado * cua.Lado

}
func (cua *Cuadrado) Perimetro() float64 {
	return cua.Lado * 4
}
