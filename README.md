
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
### Another Function Example 

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
### Function to return 2 values 

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
### Another function call undefined number of args
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
```
### Pointers in Go
```
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
```
### Getting started with Struct in Go

```
package main
import "fmt"

func main() {
	rect1 := Rectangle{
		leftX: 0,
		topY: 50,
		height: 10,
		width: 30
	}
	## or rect1 := Rectangle{0,50,10,30}

	fmt.Println("Rectangle is ", rect1.height, " of height")
//	fmt.Println("Area of Rectangle is", area(&rect1.width, &rect1.height))
	fmt.Println("Area = ", rect1.area())
}

type Rectangle struct {
	leftX float64
	topY float64
	height float64
	width float64
}

/*
func area (w, h *float) float {
	return *w * *h
}
*/

func (rect *Rectangle) area () float64 { //define a function area with a reciever (rect *Rectangle)
	return rect.width * rect.height
}

```

### Getting started with interfaces in Go
```
package main
import "fmt"
import "math"

func main() {
	rect1 := Rectangle{20,30}
	cir1 := Circle {10.0}
	fmt.Println("Rectangle Area =", getArea (rect1))
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width float64
}

type Circle struct {
	radius float64
}

func ( rec Rectangle ) area() float64 {
	return rec.height * rec.width
}

func (cir Circle) area() float64 {
	return math.Pi * math.Pow(cir.radius, 2)
}

func getArea(sh Shape) float64 {
	return sh.area()
}

## Getting started with Strings

```
package main

import ("fmt"
"strings"
"sort"
"log"
"strconv")

func main() {
	sampString := "Hallo World"
	fmt.Println(strings.Contains(sampString, "lo"))	//sampString contains lo
	fmt.Println(strings.Index(sampString, "lo"))
	fmt.Println(strings.Count(sampString, "l"))
	fmt.Println(strings.Replace(sampString, "l", "x", 3)) // replace all first 3 'l's with 'x'
	fmt.Println(strings.Split(sampString, " "))

	listOfLetters := []string{"d","a","f"}
	sort.Strings(listOfLetters)

	fmt.Println("Sorted Letters = ", listOfLetters)


}
```

## File io
```
package main

import ( "fmt"
"os"
io/util)

func main() {
	file, err := os.Create(sample.txt)

	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("some text") // write to the file

	file.Close() // close the file

	stream, err := ioutil.ReadFile(sample.txt)
	if err != nil {
		log.Fatal(err)
	}

	readString := string(stream)

	fmt.Println(readString)
}

```
## Casting
```
package main

import ("fmt"
strconv)

func main() {
	myInt := 10
	myFloat := 10.5
	myString := "100"

	// convert int to float

	fmt.Println(float64(myInt))

	 // convert float to int

	fmt.Println(int(myFloat))

	// convert from string to int

	newInt, _ := strconv.ParseInt(myString, 0, 64)
	fmt.Println(newInt)

	// convert from string to float

	newFloat, _ := strconv.ParseFloat(myString, 64)
	fmt.Println(newFloat)

}

```

## Create HTTP request in Go, Webserver

```
package main

import("fmt"
"net/http")

func main() {
	
	http.HandleFunc("/", handler) // call the handler function anytime you request for / in the URL

	// Now we need to know what port to listen to

	http.ListenAndServe(":8080", nil)

}

// declear the handler function

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hi welcome\n") // This will be printed to the client screen	
} 


```
Now whan you execute the above code, 
accessing localhost:8080 show print out the result: Hi Welcome

## Another Example

```
package main

import ("fmt"
"time")

func main() {
	for i := 0; i < 10; i++ {
	go count(i) // call the go routine using the go call
	}

	time.Sleep(time.Millisecond * 11000)
}

func count(id int) {
	for i := 0; i < 10; i++ {
		fmt.Println(id, ":", i)
		time.Sleep(time.Millisecond * 1000) // sleep for one second
	}
}

```

## channels  to pass Data to go Routines

```
package main
import ("fmt"
"time"
"strconv")

var articleNum = 0
var articleName = ""

func bestellung (stringChan chan string) {
	articleNum++
	articleName = "Article " + strconv.Itoa(articleNum) // this will convert an int to a string

	fmt.Println("Make bestellung and more Articles")

	stringChan <- articleName

	time.Sleep(time.Millisecond * 10)
}

func moreBestellung (stringChan chan string) {
	articleName : <- stringChan
	
	fmt.Println("Add new article and send", articleName)

	stringChan <- articleName

	time.Sleep(time.Millisecond * 10)
}

func addArticles (stringChan chan string ) {
	articles : <- stringChan
	fmt.Println("Add article to", artcles, "and send")
	time.Sleep(time.Millisecond * 10)	
}


func main() {
	//create a channel to hold a string

	stringChan := make(chan string)

	for i := 0; i < 10; i++ {
		go bestellung(stringChan) // go routine to create 3 bestellung
		go moreBestellung(stringChan)
		go addArticles(stringChan)

		time.Sleep(time.Millisecond * 5000)
	}
}

```

