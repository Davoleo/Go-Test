package main

import . "../../lib/go/g2d"
import "math"

/**
Calcola l'area di un ellisse dai 2 semiassi inseriti dall'utente
*/
func main() {
	a := ToFloat(Prompt("Primo semiasse?"))
	b := ToFloat(Prompt("Secondo semiasse?"))

	Alert(EllipseArea(a, b))
}

func EllipseArea(a, b float64) float64 {
	return a * b * math.Pi
}
