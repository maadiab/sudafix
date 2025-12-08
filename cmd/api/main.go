package main

import (
	"log"
	"net/http"
)

func main() {

	// mux := routes.Route

	s := &http.Server{
		Addr: ":8080",
		// Handler: mux,
	}

	log.Printf("Server is running, and it listens to port: %s", s.Addr)

	s.ListenAndServe()
}
