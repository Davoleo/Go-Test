package main

import . "../../lib/go/g2d"

 

func main() {
    
    n := ToInt(Prompt("Inserisci il numero di lanci"))
    results := make([]int, n)
    
    for i := 0; i < n; i++ {
        results[i] = RandInt(2, 12)
    }
    
    getValueFrequency1(results, 11)
    
}

func getValueFrequency1(list []int, num int) {
    for index, value := range list {
        if num == value {
            Println(index, value)
        }
    }
}
