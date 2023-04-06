package main

import (
	"log"
	"net/http"
	"github.com/Michle99/goRestApi/routes"
)

func main() {

	router := routes.MovieRoutes()

	http.Handle("/api/v1", router)

	log.Println("Listening on")
	log.Fatal(http.ListenAndServe(":8000", router))
}