package main

import . "g2d"

func main() {
    values := []int{}
	userIn := ToInt(Prompt("Inserire primo valore (0 per terminare l'inserimento)"))
    var max int = 0;
    
	for userIn != 0 {
		values = append(values, userIn)
		userIn = ToInt(Prompt("Inserire altro valore (0 per terminare l'inserimento)"))
        if max < userIn {
            max = userIn
        }
	}

	InitCanvas(Size{500, 40 * len(values)})

	for index, value := range values {
        FillRect(Rect{0, (40 * index)+1, (value*500)/max, 40})
	}
    
	MainLoop()
}
