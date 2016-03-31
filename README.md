######### go tutorial


functions in go

### Defining a Function:
```
func function_name( [parameter list] ) [return_types] {
	
	body of the function
}
```

example:
```
func max( num1, num2 int ) int {
	/* declear a local variable */

	result int

	if ( num1 > num2 ) {
		result = num1
	}
	else {
		result = num2
	}

	return result
}

### Calling a function

package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200
	var result int

	result = max(a,b)

	fmt.Printf("Max value = %d\n", result)	
}

```
############### Another Function Example 

```
package main

import "fmt"

func main () {
	
	listnumbers := []float64{1,2,3,4,5}

	fmt.Printf("Sum of listnumbers ", addListNumbers(listnumbers))
}

func addListNumbers ( number []float64 ) float64 {
	var sum := 0.0
	for _, val := range number {
		sum += val
	}
}

```
######### Function to return 2 values 

```
package main

import "fmt"

func main () {
	num1, num2 := next2Values(5)
	fmt.Println(num1, num2)
}

func next2Values(number int) (int, int) {
	return number+1, number+2
}

```
######## Another function call undefined number of args
```
package main
import "fmt"
func main () {
	fmt.Println (subtractNumbers(1,2,3,4,5))

}

func subtractNumbers(args ... int) int {
	finalValue := 0
	for _, val := range args {
		finalValue -= val
	}

	return finalValue
}

```
### catch errors and print out
```
package main
import "fmt"

func main() {
	
	fmt.Println(safeDiv(3,0))
	fmt.Println(safeDiv(4,2))
}

func saveDiv(num1,num2 int) int {

	defer func(){
		fmt.Println(recover())
	}()   //this will catch the Error when div by 0

	solution := num1/num2
	return solution
}
```

### Call to Panic and demo how to catch Error when Panic occures

```
package main
import "fmt"

func main() {
	demPanic()
}

func demoPanic(){
	defer func() {
		fmt.Println(recover())
	}()

	panic("OPs PANIC")
}

### Pointers in Go

package main
import "fmt"

func main () {
	num := 0

	/*
	or declare a pointer value with
	xPtr := new(int)
	changeXValue(xptr)
	fmt.Println("Value of xPtr is " *xPtr)
	*/

	changeXVal(&num)
	fmt.Println("Value of num :=", num)
	fmt.Println("Memory value of num :=", &num)
}

func changeXVal(x *int) {
	*x = 2
}

End





