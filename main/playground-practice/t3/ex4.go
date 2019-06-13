package main

import . "../../lib/go/g2d"

//Main
func main() {

    InitCanvas(arena.Size())
    MainLoop(tick)
}

//LOG CLASS -----------------------------------------------------------
type Log struct {

}

// Vehicle11 CLASS ------------------------------------------------------
type Vehicle struct {
    arena         *Arena
    x, y, w, h    int
    speed, dx, dy int
    category string
    direction string
}

func NewVehicle(arena *Arena, pos Point, width int, speed int, category string, direction string) *Vehicle {
    a := &Vehicle{arena, pos.X, pos.Y, width, 32, speed, speed, speed, category, direction}
    arena.Add(a)
    return a
}

func (a *Vehicle) Move() {
    as := a.arena.Size()

    //Controls whether the vehicle should be brought again to the other side of the arena
    if a.x > as.W + 100 {
        a.x = -100
    }
    if a.x < -100 {
        a.x = as.W + 100
    }

    //Controls in which direction the vehicle moves
    if	a.direction == "dx" {
        a.x += a.dx
    } else if a.direction == "sx" {
        a.x -= a.dx
    }
}

func (a *Vehicle) Position() Rect {
    return Rect{a.x, a.y, a.w, a.h}
}

//Controls the symbol of the vehicle based on its category and direction
func (a *Vehicle) Symbol() Rect {
    if	a.category == "fast_car" {
        if	a.direction == "dx" {
            return Rect{192, 0, a.h, a.w}
        } else {
            return Rect{192, 32, a.h, a.w}
        }
    }
    if a.category == "car" {
        if a.direction == "sx" {
            return Rect{256, 0, a.h, a.w}
        } else {
            return Rect{256, 32, a.h, a.w}
        }
    }
    if a.category == "truck" {
        if a.direction == "dx" {
            return Rect{256, 64, a.w, a.h}
        } else {
            return Rect{192, 64, a.w, a.h}
        }
    }
    if a.category == "tractor" {
        if	a.direction == "dx" {
            return Rect{224, 0, a.w, a.h}
        } else {
            return Rect{224, 32, a.w, a.h}
        }
    }
    return Rect{0, 0, 0, 0}
}

func (a *Vehicle) Speed() Point {
    return Point{a.dx, a.dy}
}

func (a *Vehicle) Collide(other Actor) {
}

//GHOST CLASS ----------------------------------------------------
type Ghost struct {
    arena      *Arena
    x, y, w, h int
    speed      int
    visible    bool
}

func NewGhost(arena *Arena, pos Point) *Ghost {
    a := &Ghost{arena, pos.X, pos.Y, 20, 20, 5, true}
    arena.Add(a)
    return a
}

func (a *Ghost) Move() {
    as := a.arena.Size()
    dx := RandInt(-1, 1) * a.speed
    dy := RandInt(-1, 1) * a.speed
    a.x = (a.x + dx + as.W) % as.W
    a.y = (a.y + dy + as.H) % as.H

    if RandInt(0, 99) == 0 {
        a.visible = !a.visible
    }
}

func (a *Ghost) Position() Rect {
    return Rect{a.x, a.y, a.w, a.h}
}

func (a *Ghost) Symbol() Rect {
    if a.visible {
        return Rect{20, 0, a.w, a.h}
    }
    return Rect{20, 20, a.w, a.h}
}

func (a *Ghost) Collide(other Actor) {
}

//Frog CLASS ------------------------------------------------
type Frog struct {
    arena         *Arena
    x, y, w, h    int
    speed, dx, dy int
    count int
}

func NewFrog(arena *Arena, pos Point) *Frog {
    a := &Frog{arena, pos.X, pos.Y, 32, 32, 8, 0, 0, 0}
    arena.Add(a)
    return a
}

func (a *Frog) Move() {
    if a.count > 0 {
        a.count--
        as := a.arena.Size()
        a.x += a.dx
        if a.x < 0 {
            a.x = 0
        } else if a.x > as.W-a.w {
            a.x = as.W - a.w
        }
        a.y += a.dy
        if a.y < 0 {
            a.y = 0
        } else if a.y > as.H-a.h {
            a.y = as.H - a.h
        }
    }

}
var moving bool = false

func (a *Frog) GoLeft() {
    if a.count == 0 {
        a.count = 4
    }
    a.dx, a.dy = -a.speed, 0
}

func (a *Frog) GoRight() {
    if a.count == 0 {
        a.count = 4
    }
    a.dx, a.dy = +a.speed, 0
}

func (a *Frog) GoUp() {
    if a.count == 0 {
        a.count = 4
    }
    a.dx, a.dy = 0, -a.speed
}


func (a *Frog) GoDown() {
    if a.count == 0 {
        a.count = 4
    }
    a.dx, a.dy = 0, +a.speed
}

func (a *Frog) Collide(other Actor) {
    _, isVehicle := other.(*Vehicle)
    if isVehicle {
        a.x = 300
        a.y = 416
    }
}

func (a *Frog) Position() Rect {
    return Rect{a.x, a.y, a.w, a.h}
}

func (a *Frog) Symbol() Rect {
    if a.count == 0 {
        return Rect{0, 0, a.w, a.h}
    }
    return Rect{32, 0, a.w, a.h}
}
//--------------------------------------------------
//Global Vars
var arena = NewArena(Size{640, 448})
var hero = NewFrog(arena, Point{300, 416})
var Vehicle1 = NewVehicle(arena, Point{150, 384}, 32, 10, "fast_car", "sx")
var Vehicle2 = NewVehicle(arena, Point{0, 320}, 32, 8, "car", "sx")
var Vehicle3 = NewVehicle(arena, Point{200, 288}, 32, 14, "fast_car", "dx")
var Vehicle4 = NewVehicle(arena, Point{200, 256}, 64, 6, "truck", "sx")
var Vehicle5 = NewVehicle(arena, Point{200, 352}, 64, 6, "truck", "dx")
var Vehicle6 = NewVehicle(arena, Point{300, 384}, 32, 10, "fast_car", "sx")
var Vehicle7 = NewVehicle(arena, Point{123, 320}, 32, 8, "car", "sx")
var Vehicle8 = NewVehicle(arena, Point{1000, 256}, 32, 6, "tractor", "sx")
var Vehicle9 = NewVehicle(arena, Point{400, 256}, 64, 6, "truck", "sx")
var Vehicle10 = NewVehicle(arena, Point{400, 352}, 64, 6, "truck", "dx")
var ghost = NewGhost(arena, Point{120, 80})

var sprites = LoadImage("https://tomamic.github.io/images/misc/frogger_32.png")
var bg = LoadImage("https://tomamic.github.io/images/misc/frogger_bg.png")

//Ticking function
func tick() {

    if (KeyPressed("ArrowUp")) {
        hero.GoUp()
    } else if KeyPressed("ArrowRight") {
        hero.GoRight()
    } else if KeyPressed("ArrowDown") {
        hero.GoDown()
    } else if KeyPressed("ArrowLeft") {
        hero.GoLeft()
    }


    arena.MoveAll()
    ClearCanvas()
    DrawImageClip(bg, Rect{0, 15, 640, 480}, Rect{0, 0, 640, 480})
    for _, a := range arena.Actors() {
        if a.Symbol().H != 0 {
            DrawImageClip(sprites, a.Symbol(), a.Position())
        } else {
            FillRect(a.Position())
        }
    }
}
