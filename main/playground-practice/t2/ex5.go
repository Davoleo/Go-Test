package main

import . "../../lib/go/g2d"
import "math"

func main() {
	a := ToFloat(Prompt("Primo Semiasse?"))
	b := ToFloat(Prompt("Secondo Semiasse?"))

	ellipse := NewEllipse(Point{0, 0}, a, b)

	Println(ellipse.Area())
	Println(ellipse.FocalDistance())

}

type Ellipse struct {
	x, y int
	a, b float64
}

func NewEllipse(pos Point, a, b float64) *Ellipse {
	ellipse := &Ellipse{pos.X, pos.Y, a, b}
	return ellipse
}

func (e *Ellipse) Area() float64 {
	return math.Pi * e.a * e.b
}

func (e *Ellipse) FocalDistance() float64 {
	return 2 * (math.Sqrt(math.Abs((e.a * e.a) - (e.b * e.b))))
}
