package main

import (
	"log"
	"net/http"
)

func main() {
	initRoutes()

	log.Println("Serving :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
