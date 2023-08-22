package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	
	http.HandleFunc("/", handler)
	port := os.Getenv("SERVER_PORT")
	if port == "" {
			port = "8080"
	}
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}