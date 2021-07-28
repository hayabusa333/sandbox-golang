package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// 単純なハンドラ
	r.HandleFunc("/", EchoHandler)

	// サービスを起動する
	http.ListenAndServe(":8080", r)
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}
