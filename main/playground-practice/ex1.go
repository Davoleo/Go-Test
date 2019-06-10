package playground_practice

//Draw a circle and print out the area
func main() {

	radius := ToInt(Prompt("radius?"))

	if radius > 0 && radius <= 200 {
		//Inizializzo il canvas
		InitCanvas(Size{500, 500})

		//Seleziono un colore randomico
		SetColor(Color{RandInt(1, 255), RandInt(1, 255), RandInt(1, 255)})

		//Disegno un cerchio pieno che ha come centro il centro del canvas e come grandezza il raggio inserito dall'utente
		FillCircle(Point{250, 250}, radius)
		MainLoop()

		//Stampo l'area in console
		Println((float64)(radius*radius) * math.Pi)
	} else {
		Alert("Error: Radius out of range!")
	}
}
