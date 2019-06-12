package main

import . "g2d"

var ballObj *Ball = NewBall(Point{20, 100})
var canvasSize Size = Size{500, 500}
var count int = 0

func tick() {

	if count < 5 {
		ClearCanvas()
		ballObj.Move()
	}

	if KeyPressed("Enter") {
		count = 0
	}

	count++
}

func (b *Ball) Move() {
	FillCircle(Point{b.x, b.y}, b.w/2)

	if b.x >= canvasSize.W+100 {
		b.x = -100
	}

	//if b.x < -100 {
	//	b.x = canvasSize.W + 100
	//}

	b.x += b.dx

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
