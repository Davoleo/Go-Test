package _go

import . "g2d"

func tick() {
    if KeyPressed("LeftButton") {
        mp := MousePosition()
        SetColor(Color{RandInt(0, 255), RandInt(0, 255), RandInt(0, 255)})
        if mp.X <= 25 && mp.Y <= 25 && Confirm("Exit?") {
            CloseCanvas()
        } else {
            FillCircle(mp, 25)
        }
    }
}

func main() {
    InitCanvas(Size{480, 360})
    MainLoop(tick)  // Call tick 30 times/second
}
