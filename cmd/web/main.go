package main

import (
	"log"
	"net/http"
)

func main() {
	initRoutes()

	log.Println("Serving :80")

	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
