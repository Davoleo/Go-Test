package main

import . "g2d"
import "math"

var xc, yc, k float64 = 250, 250, 200
var i int = 0
var canvasSize Size = Size{500, 500}

/**
Animazione periodica di una spirale
*/
func main() {

	InitCanvas(canvasSize)
	SetFrameRate(20)
	MainLoop(tick)

}

func tick() {
	ClearCanvas()
	Move()
}

func Move() {
	SetColor(Color{i * 2, 0, 255 - i*2})
	FillCircle(Point{int((xc * k) + 250), int((yc * k) + 250)}, 80-i)

	xc = math.Sin(float64(i))
	yc = math.Cos(float64(i))

	i++
	k -= 2.5
	if i >= 80 {
		i = 0
	}
	if k < 0 {
		k = 200
	}
}
