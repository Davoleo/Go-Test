package playground_practice

/**
Generare all'inizio del programma un numero “segreto” a caso tra 1 e 90
Chiedere ripetutamente all'utente di immettere un numero, finché non indovina quello generato
Ad ogni tentativo, dire se il numero immesso è maggiore o minore del numero segreto
*/
func main() {

	secretNumber := RandInt(1, 90)
	userGuess := ToInt(Prompt("Indovina un numero da 1 a 90: "))

	for userGuess != secretNumber {
		if secretNumber > userGuess {
			Alert("Il numero segreto è maggiore di quello che hai inserito")
		} else if secretNumber < userGuess {
			Alert("Il numero segreto è minore di quello che hai inserito")
		}

		userGuess = ToInt(Prompt("Riprova:"))
	}

	Alert("Hai indovinato!")

}
