package figuras

type Circulo struct {
	Radio int
}

func (p *Circulo) CalcularPerimetro() int {
	return p.Radio * 2
}

func (p *Circulo) CalcularArea() int {
	return p.Radio * p.Radio * 2
}
