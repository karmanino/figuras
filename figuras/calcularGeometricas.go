package figuras

type Figura interface {
	CalcularPerimetro() int
	CalcularArea() int
}

func CalcularArea(figura Figura) int {
	return figura.CalcularArea()
}
func CalcularPerimetro(figura Figura) int {
	return figura.CalcularPerimetro()
}
