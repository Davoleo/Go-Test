package main

import . "../../lib/go/g2d"
//Allerta il più piccolo di 3 numeri randomici
func main() {
	var a, b, c int = RandInt(1, 6), RandInt(1, 6), RandInt(1, 6)

	Println("a = ", a)
	Println("b = ", b)
	Println("c = ", c)

	if a < b && a < c {
		Alert("Il numero minore è: ", a)
	} else if b < c {
		Alert("Il numero minore è: ", b)
	} else {
		Alert("Il numero minore è: ", c)
	}
}
