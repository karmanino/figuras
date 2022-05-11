package main

import (
	"figuras/figuras"
	"fmt"

	"github.com/donvito/hellomod"
)

func main() {

	hellomod.SayHello()

	cuadrado := figuras.Cuadrado{Lado: 20}
	circulo := figuras.Circulo{Radio: 20}

	fmt.Println(figuras.CalcularArea(&cuadrado))
	fmt.Println(figuras.CalcularArea(&circulo))
	fmt.Println(figuras.CalcularPerimetro(&cuadrado))
	fmt.Println(figuras.CalcularPerimetro(&circulo))

}
