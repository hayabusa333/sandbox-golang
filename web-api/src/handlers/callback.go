package handlers

import (
        "net/http"
)

func LoginCallback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
        w.Header().Set("location", "http://localhost:8080")
        w.WriteHeader(http.StatusFound)
}
