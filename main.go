package main

import (
	"log"
	"net/http"
)

func main() {
	router := initializeRoutes()
	log.Fatal(http.ListenAndServe(":8080", router))
}
