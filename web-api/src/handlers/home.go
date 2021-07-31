package handlers

import (
        "net/http"
        "fmt"
)

func Home(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        fmt.Fprint(w, "Home page")
}
