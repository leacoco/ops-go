//run using 
//leadel@opensuse-test01:~/cloudprojects/ops-go/src/github.com/leacoco/webserver> go install github.com/leacoco/webserver
//leadel@opensuse-test01:~/cloudprojects/ops-go/src/github.com/leacoco/webserver> $GOPATH/bin/webserver

package main 

import ("fmt"
"net/http")

func main() {
	
	http.HandleFunc("/", handler) // call the handler function anytime you request for / in the URL

	// Now we need to know what port to listen to

	http.ListenAndServe(":8080", nil)

}

// declear the handler function

func handler(w http.ResponseWriter, r *http.Request) {
	message := "welcome to my Go website. This page is Generate by go.\n"
	message += "Thanks for using this site\n"
	message += "Hope to see you again next time and have FUN :) "
	fmt.Fprintf(w, message) // This will be printed to the client screen	

}