package main

import (
	"log"
	"net/http"
)

func main() {
	// Use the httpNewServeMux() to initialize a new servemux
	//register the home function as the handler  for the "/" URL pattern

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", home)
	//register new handlers
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Use the http.ListenAndServe() function to start a new web server
	// We pass in the TCP network address to listenn on and servemux
	//if error happens we log it and exit the process
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
