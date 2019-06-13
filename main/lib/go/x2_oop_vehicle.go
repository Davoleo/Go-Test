package _go

import . "g2d"

var screen = Size{480, 360}
var margin = 100
var a = NewVehicle(Point{40, 40}, 5)

type Vehicle struct {
    x, y, w, h      int
    dx, left, right int
}

func NewVehicle(pos Point, dx int) *Vehicle {
    return &Vehicle{pos.X, pos.Y, 20, 20, dx, -margin, screen.W+margin}
}

func (a *Vehicle) Move() {
    if a.x + a.dx < a.left {
        a.x = a.right
    }
    if a.x + a.dx > a.right {
        a.x = a.left
    }
    a.x += a.dx
}

func (a *Vehicle) Position() Rect {
    return Rect{a.x, a.y, a.w, a.h}
}

func (a *Vehicle) Uturn() {
    a.dx *= -1
}

func tick() {
    if KeyPressed("Enter") {
        a.Uturn()
    }
    ClearCanvas()
    a.Move()
    FillRect(a.Position())
}

func main() {
    InitCanvas(screen)
    MainLoop(tick)
}

