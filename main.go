package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// Use the httpNewServeMux() to initialize a new servemux
	//register the home function as the handler  for the "/" URL pattern

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Use the http.ListenAndServe() function to start a new web server
	// We pass in the TCP network address to listenn on and servemux
	//if error happens we log it and exit the process
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
