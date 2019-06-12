package main

import . "g2d"

var ballObj *Ball = NewBall(Point{20, 100})
var canvasSize Size = Size{500, 500}

func tick() {
	ClearCanvas()
	ballObj.Move()
}

func (b *Ball) Move() {
	FillCircle(Point{b.x, b.y}, b.w/2)

	if b.x >= canvasSize.W+150 {
		b.x = -150
	}

	if b.x < -100 {
		b.x = canvasSize.W + 100
	}

	b.x += b.dx

	if KeyPressed("Enter") {
		b.dx = -b.dx
	}
}

func main() {

	InitCanvas(canvasSize)
	MainLoop(tick)
}

type Ball struct {
	x, y   int
	w, h   int
	dx, dy int
}

func NewBall(pos Point) *Ball {
	b := &Ball{pos.X, pos.Y, 50, 50, 5, 5}
	return b
}
