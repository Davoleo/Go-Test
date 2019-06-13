package main

import . "../../lib/go/g2d"

var canvSize Size = Size{500, 500}
var b *Ball = NewBall(Point{40, 40})

func main() {

	InitCanvas(canvSize)

	MainLoop(tick)

}

func tick() {

	ClearCanvas()
	b.Move()

}

func (b *Ball) Move() {

	FillCircle(Point{b.x, b.y}, b.w/2)

	if b.y >= 300 {
		b.x -= b.dx
	} else if b.x >= canvSize.W-b.w {
		b.y += b.dy
	} else {
		b.x += b.dx
	}

}

type Ball struct {
	x, y   int
	w, h   int
	dx, dy int
}

func NewBall(pos Point) *Ball {
	b := &Ball{pos.X, pos.Y, 80, 80, 5, 5}
	return b
}
