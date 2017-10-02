package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func main() {
	var (
		config   string
		instance int
	)
	flag.StringVar(&config, "config", "", "path to config file")
	flag.IntVar(&instance, "instance", -1, "each process has it's own instance")
	log.Println("Starting goroku.....")
	log.Println("Getting port")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("No port found using default %s", port)
	}
	log.Printf("Binding to port %s", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request")
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
