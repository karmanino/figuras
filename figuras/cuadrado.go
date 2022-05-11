package figuras

type Cuadrado struct {
	Lado int
}

func (p *Cuadrado) CalcularPerimetro() int {
	return p.Lado * 4
}

func (p *Cuadrado) CalcularArea() int {
	return p.Lado * 2
}
