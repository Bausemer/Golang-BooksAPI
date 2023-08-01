package main

//! PAGE 68

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	mux := http.NewServeMux()
	
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/snippet/view", snippetView)
	
	url := "localhost"+*addr

	log.Printf("Starting server on %s", url)
	err := http.ListenAndServe(url, mux)
	log.Fatal(err)
}