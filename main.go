package main

import (
	"fmt"
	"log"
	"net/http"
)

const URL = "127.0.0.1:8080"

var Artists []Artist

func main() {

	err := InitTemplates()
	if err != nil {
		log.Fatal(err)
	}

	Artists = FetchArtists()

	http.HandleFunc("/indexTemplate", indexTemplate)
	http.HandleFunc("/artist", artist)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index)

	fmt.Printf("Listening on %s\n", URL)
	err = http.ListenAndServe(URL, nil)

	if err != nil {
		log.Fatal(err)
	}

}
