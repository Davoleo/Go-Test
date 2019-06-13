package main

import . "../../lib/go/g2d"

var screen = Size{480, 360}
var margin = 100
var a = NewVehicle2(Point{40, 40}, 5)

type actor interface {
    Move()
    symbol()
    collide()
}

type Vehicle11 struct {
    x, y, w, h      int
    dx, left, right int
}

func NewVehicle2(pos Point, dx int) *Vehicle11 {
    return &Vehicle11{pos.X, pos.Y, 20, 20, dx, -margin, screen.W+margin}
}

func (a *Vehicle11) symbol() {
	
}

func (a *Vehicle11) collide() {
	
}

func (a *Vehicle11) Move() {
    if a.x + a.dx < a.left {
        a.x = a.right
    }
    if a.x + a.dx > a.right {
        a.x = a.left
    }
    a.x += a.dx
}

func (a *Vehicle11) Position() Rect {
    return Rect{a.x, a.y, a.w, a.h}
}

func (a *Vehicle11) Uturn() {
    a.dx *= -1
}

func tick2() {
    if KeyPressed("Enter") {
        a.Uturn()
    }
    ClearCanvas()
    a.Move()
    FillRect(a.Position())
}

func main() {
    InitCanvas(screen)
    MainLoop(tick2)
}

