package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
this program is a simple web server. it has three routs:
- /: this invokes index.html
- hello: invokes hello func
- /form: invokes form func which also invokes form.html
*/

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	//possible routs and how each one will get handled
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	//server connection and error handling
	fmt.Println("Server runs on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// response is what the server will send to the font-end and the request is what's coming from the front-end/user
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	//limit the users to only get from this rout, don't allow post
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "Hello..")

}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	//limit the users to only get from this rout, don't allow post
	if r.Method != "POST" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "POST request was successful:\n\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name = %s\n", name)
	fmt.Fprintf(w, "address = %s\n", address)

}
