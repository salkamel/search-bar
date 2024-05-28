package main

import (
	"fmt"
	"groupie_tracker_search_bar/backend"
	"log"
	"net/http"
)

var HomeURL = "http://localhost:8080/"

func main() {
	err := backend.InitData()
	if err != nil {
		fmt.Println("Error:", err); return
	}
	http.HandleFunc("/", backend.HomePage)
	http.HandleFunc("/artistid/", backend.ArtistsPage)
	// Serve static files from the "frontend" directory
	fs := http.FileServer(http.Dir("frontend"))
	http.Handle("/frontend/", http.StripPrefix("/frontend/", fs))

	go func() { // To run server in background
		log.Fatal(http.ListenAndServe(":"+backend.GetURLPort(HomeURL), nil)) // Run server
	}()
	fmt.Println(HomeURL)
	select {} // Run server in background and print error message when server is not running in background and exit
}