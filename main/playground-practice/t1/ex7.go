package main

import . "../../lib/go/g2d"

func main() {

	resInput := ToFloat(Prompt("Inserisci una resistenza (0 per terminare l'inserimento)"))
	totalSerie := 0.0
	totalParallelo := 0.0

	for resInput != 0 {
		totalSerie += resInput
		totalParallelo += (1.0 / resInput)

		resInput = ToFloat(Prompt("Inserisci una resistenza (0 per terminare l'inserimento)"))
	}

	Alert("Resistenza totale in Serie: ", totalSerie)
	Alert("Resistenza totale in Parallelo: ", totalParallelo)

}
