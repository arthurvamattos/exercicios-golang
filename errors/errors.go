package main

import (
	"fmt"
	"math"
)

type sqrtError struct {
    arg  float64
    prob string
}

func (e *sqrtError) Error() string {
    return fmt.Sprintf("%s: %f", e.prob, e.arg)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return -1, &sqrtError{x, "Impossível obter a raiz de número negativo"}
	}
		return math.Sqrt(x), nil
	}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
