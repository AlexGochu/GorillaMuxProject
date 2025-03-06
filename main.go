package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var task string

type requestBody struct {
	Task string `json:"task"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello, %s", task)
}
func PostHandler(w http.ResponseWriter, r *http.Request) {
	var req requestBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "JSON's decoder error", http.StatusBadRequest)
		return
	}

	task = req.Task
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			PostHandler(w, r)
		case http.MethodGet:
			GetHandler(w, r)
		default:
			http.Error(w, "Untreated method:", http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":8080", router)
}
