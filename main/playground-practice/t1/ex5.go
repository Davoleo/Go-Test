package main

import . "../../lib/go/g2d"

/**
Chiedere all'utente il numero di cerchi da disegnare
Disegnare i cerchi con raggio gradualmente decrescente, ma tutti con lo stesso centro
Far variare il colore dei cerchi
Dal rosso del livello più esterno
Fino al nero del livello più interno
*/
func main() {
	num := ToInt(Prompt("Numero di cicli?"))
	//step := 255/num
	//Println(step)

	InitCanvas(Size{510, 510})

	for i := num; i > 0; i -= 10 {
		SetColor(Color{i, 0, 0})
		FillCircle(Point{255, 255}, i)
	}

	MainLoop()
}
