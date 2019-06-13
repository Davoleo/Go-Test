package main

import . "../../lib/go/g2d"

var ballObj *Ball2 = NewBall2(Point{20, 100})
var canvasSize2 Size = Size{500, 500}
var count int = 0

func tick2() {

	if count < 5 {
		ClearCanvas()
		ballObj.Move()
	}

	if KeyPressed("Enter") {
		count = 0
	}

	count++
}

func (b *Ball2) Move() {
	FillCircle(Point{b.x, b.y}, b.w/2)

	if b.x >= canvasSize2.W+100 {
		b.x = -100
	}

	//if b.x < -100 {
	//	b.x = canvSize.W + 100
	//}

	b.x += b.dx

}

func main() {

	InitCanvas(canvasSize2)
	MainLoop(tick2)
}

type Ball2 struct {
	x, y   int
	w, h   int
	dx, dy int
}

func NewBall2(pos Point) *Ball2 {
	b := &Ball2{pos.X, pos.Y, 50, 50, 5, 5}
	return b
}
