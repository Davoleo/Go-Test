package main

import . "../../lib/go/g2d"

func main() {
    
    n := ToInt(Prompt("Inserisci il numero di lanci"))
    results := make([]int, n)
    
    for i := 0; i < n; i++ {
        results[i] = RandInt(2, 12)
    }
    
    getValueFrequency(results)
    
}

//Questa funzione serve a contare quante volte escono i vari numeri
func getValueFrequency(list []int) {
    
    for i := 2; i <= 12; i = i + 1 {
    	count := 0
        for _, value := range list {
            if i == value {
                //Println(index, value)
                count++
            }
        }
        Println(i, " è uscito ", count, " volte")
    }
}
