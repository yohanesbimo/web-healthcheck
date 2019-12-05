package main

import (
	"net/http"

	"github.com/web-healthcheck/pkg/api"
)

func initRoutes() {
	http.HandleFunc("/web-healthcheck", api.WebHealthcheck)
}
