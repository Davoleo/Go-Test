package main

import . "../../lib/go/g2d"

var ball *Ball3 = NewBall3(Point{20, 100})
var size Size = Size{500, 500}

func tick3() {
	ClearCanvas()
	ball.Move()
}

func (b *Ball3) Move() {
	FillCircle(Point{b.x, b.y}, b.w/2)

	if b.x >= size.W+150 {
		b.x = -150
	}

	if b.x < -100 {
		b.x = size.W + 100
	}

	b.x += b.dx

	if KeyPressed("Enter") {
		b.dx = -b.dx
	}
}

func main() {

	InitCanvas(size)
	MainLoop(tick3)
}

type Ball3 struct {
	x, y   int
	w, h   int
	dx, dy int
}

func NewBall3(pos Point) *Ball3 {
	b := &Ball3{pos.X, pos.Y, 50, 50, 5, 5}
	return b
}
