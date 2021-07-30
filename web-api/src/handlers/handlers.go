package handlers

import (
        "net/http"
        "fmt"
        "github.com/gorilla/mux"
)

func Router() *mux.Router {
        r := mux.NewRouter()

        // 死活監視
        r.HandleFunc("/heartbeat", HearthbeatHandler)

        return r
}

func HearthbeatHandler(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        fmt.Fprint(w, "OK")
}