package main

import (
	"fmt"
	"math"
)

type Geometria1 interface {
	area() float64
}

type Quadrado1 struct {
	lado float64
}

func (q Quadrado) area1() float64 {
	return q.lado * q.lado
}

type Circulo1 struct {
	raio float64
}

func (c Circulo) area1() float64 {
	return math.Pi * c.raio * c.raio
}

func main1() {
	var g Geometria
	g = Quadrado{3}
	fmt.Printf("Area do quadrado e %v\n", g.area())
	
	g = Circulo{5}
	fmt.Printf("A area do circulo e %v\n", g.area())
}