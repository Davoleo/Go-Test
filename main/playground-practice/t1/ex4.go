package main

import . "../../lib/go/g2d"

/**
Chiedere all'utente un numero n
Disegnare n quadrati:
Tutti con lato di 100 pixel
Ciascuno in posizione casuale
Ciascuno con un colore casuale
*/
func main() {
	num := ToInt(Prompt("Numero di cicli?"))

	InitCanvas(Size{500, 500})

	for i := 0; i < num; i++ {
		SetColor(Color{RandInt(1, 255), RandInt(1, 255), RandInt(1, 255)})
		FillRect(Rect{RandInt(1, 450), RandInt(1, 450), 100, 100})
	}

	MainLoop()
}
