package main

import "fmt"

func main() {
	/*declear some local variables */
	var num1 int = 100
	var num2 int = 200
	var result int

	result = maximum(num1,num2)
	fmt.Printf("Max value = %d\n", result)

}

func maximum(a, b int) int {
	if ( a > b ) {
		return a
	}else {
		return b
	}
}
