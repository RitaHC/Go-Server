package main

import (
	"fmt"
	"log"
	"net/http"
)

type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Error handler for 404
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	// Error handler for method
	if r.Method != "GET" {
		http.Error(w, "Error is Not Supported", http.StatusNotFound)
		return
	}
	// Action to be performed when the page is found
	fmt.Fprintf(w, "Hello!")
}

type FormHandler struct{}

func (f FormHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse Form() err : %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name =%s", name)
	fmt.Fprintf(w, "Address =%s", address)
}
func main() {
	// Telling to go and check static directory
	fileServer := http.FileServer(http.Dir("./static"))

	////////// Register Routes //////////
	// handle function -> for route handling
	http.Handle("/", fileServer)
	http.Handle("/form", FormHandler{})
	http.Handle("/hello", HelloHandler{})

	fmt.Println("Server started at PORT 8080")
	// Error Handling
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
