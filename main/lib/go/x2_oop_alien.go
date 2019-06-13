package _go

import . "g2d"

var a = NewAlien(Point{40, 40})

type Alien struct {
    x, y, w, h int
    xmin, xmax int
    dx, dy     int
}

func NewAlien(pos Point) *Alien {
    return &Alien{pos.X, pos.Y, 20, 20, pos.X, pos.X+150, 5, 5}
}

func (a *Alien) Move() {
    if a.xmin <= a.x+a.dx && a.x+a.dx <= a.xmax {
        a.x += a.dx
    } else {
        a.dx = -a.dx
        a.y += a.dy
    }
}

func (a *Alien) Position() Rect {
    return Rect{a.x, a.y, a.w, a.h}
}

func tick() {
    ClearCanvas()
    a.Move()
    FillRect(a.Position())
}

func main() {
    InitCanvas(Size{480, 360})
    SetFrameRate(10)
    MainLoop(tick)
}
