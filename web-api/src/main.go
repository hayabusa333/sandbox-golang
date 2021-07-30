package main

import (
	"net/http"
	"github.com/web-api/handlers"
)

func main() {
	router := handlers.Router()

	// サービスを起動する
	http.ListenAndServe(":8080", router)
}
