package main

import "fmt"

func main () {

	listnumbers := []float64{1,2,3,4,5}

	fmt.Printf("Sum of listnumbers = %f\n ", addListNumbers(listnumbers))
}

func addListNumbers ( number []float64 ) float64 {
        sum := 0.0
	for _, val := range number {
	sum += val
	}
	return sum
}
