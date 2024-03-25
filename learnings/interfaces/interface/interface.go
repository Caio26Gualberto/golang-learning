package main

import (
	"fmt"
	"math"
)

type Retangulo struct {
	altura  float64
	largura float64
}

type Circulo struct {
	raio float64
}

type Form interface {
	area() float64
}

func (r Retangulo) area() float64 {
	return r.altura * r.largura
}

func (c Circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func writeArea(f Form) {
	fmt.Printf("A area da forma é %0.2f\n", f.area())
}

func main() {
	r := Retangulo{10, 15}
	writeArea(r)

	c := Circulo{10}
	writeArea(c)
}
