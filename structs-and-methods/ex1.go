package main

import (
	"fmt"
	"math"	
)

type Ponto struct {
	x, y int
}

func calcDist(n1, n2 int) int{
	if n1 > n2 {
		return n1 - n2
	} 
	return n2 - n1
} 

func (p Ponto) distance(outroPonto Ponto) float64 {
	distX := calcDist(p.x,  outroPonto.x)
	distY := calcDist(p.y,  outroPonto.y)

	c1 := distX * distX
	c2 := distY * distY
	result := math.Sqrt(float64(c1) + float64(c2))
	return result
}

func main() {
	ponto := Ponto{2, 5}
	outroPonto := Ponto{-5, -2}
	fmt.Println(ponto.distance(outroPonto))
}