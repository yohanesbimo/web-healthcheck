package main

import (
	"net/http"

	"github.com/web-healthcheck/pkg/web"
)

func initRoutes() {
	http.HandleFunc("/", web.HomeView)
}
