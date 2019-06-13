package main

import . "../../lib/go/g2d"

/**
Disegna un numero di quadrati in matrice con un gradiente di colore diverso a seconda della loro posizione x o y
*/
func main() {

	rows := ToInt(Prompt("Righe?"))
	cols := ToInt(Prompt("Colonne?"))

	InitCanvas(Size{800, 800})

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			SetColor(Color{0, j * 10, i * 10})
			FillRect(Rect{i * 21, j * 21, 20, 20})
		}
	}

	MainLoop()
}
