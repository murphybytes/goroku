package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting goroku.....")
	log.Println("Getting port")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("No port found")
	}
	log.Printf("Binding to port %s", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request")
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
