package main

import (
	"chat-app/routes"
	"log"
	"net/http"
)

func main() {
	routes := routes.Mux()

	log.Println("Starting web server on port:8080")

	_ = http.ListenAndServe(":8080", routes)
}
