package _go

import . "g2d"

// way too much global stuff!
var x1, y1, dx1, dy1 = 40, 80, 5, 5
var x2, y2, dx2, dy2 = 80, 40, 5, 5
var screen = Size{480, 360}
var size = Size{20, 20}
var image = LoadImage("ball.png")

// encapsulates behaviour, but exposes data
func MoveBall(x, y, dx, dy int) (int, int, int, int) {
    if x + dx < 0 || x + dx + size.W > screen.W {
        dx = -dx
    }
    x += dx
    if y + dy < 0 || y + dy + size.H > screen.H {
        dy = -dy
    }
    y += dy
    return x, y, dx, dy
}

func tick() {
    ClearCanvas()                    // Draw background
    DrawImage(image, Point{x1, y1})  // Draw foreground
    DrawImage(image, Point{x2, y2})  // Draw foreground
    x1, y1, dx1, dy1 = MoveBall(x1, y1, dx1, dy1)
    x2, y2, dx2, dy2 = MoveBall(x2, y2, dx2, dy2)
}

func main() {
    InitCanvas(screen)
    MainLoop(tick)
}
