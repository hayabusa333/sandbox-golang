package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// 単純なハンドラ
	r.HandleFunc("/", EchoHandler)

	// 死活監視
	r.HandleFunc("/heartbeat", HearthbeatHandler)

	// サービスを起動する
	http.ListenAndServe(":8080", r)
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func HearthbeatHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}
