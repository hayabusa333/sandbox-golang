package handlers

import (
	"encoding/json"
	"net/http"
)

type Item struct {
	Id	string	`json:"id"`
	Name	string	`json:"name,omitempty"`
}

var items []*Item

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// Mock用のデータ
func init() {
	items = []*Item{
		&Item{
			Id: "1",
			Name: "test_1",
		},
		&Item{
			Id: "2",
			Name: "test_2",
		},
		&Item{
			Id: "3",
			Name: "test_3",
		},
	}
}
