package api

import (
	"net/http"
)

// WebHealthcheck Healtcheck to some endpoint
func WebHealthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	url := r.Form.Get("url")
	if !Healthchecker(url) {
		w.Write([]byte(`{"status": "Unhealthy"}`))
		return
	}

	w.Write([]byte(`{"status": "Healthy"}`))
}
