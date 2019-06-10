package playground_practice

func main() {
	birthDay := ToInt(Prompt("Giorno di Nascita?"))
	birthMonth := ToInt(Prompt("Mese di Nascita?"))
	birthYear := ToInt(Prompt("Anno di Nascita?"))

	currentDay := ToInt(Prompt("Giorno Corrente?"))
	currentMonth := ToInt(Prompt("Mese Corrente?"))
	currentYear := ToInt(Prompt("Anno Corrente?"))

	age := currentYear - birthYear

	if birthMonth > currentMonth {
		age--
	} else if birthDay > currentDay {
		age--
	}

	Alert("Hai ", age, " anni")
}
