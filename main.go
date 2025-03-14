package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	InitDB()
	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()

	router.HandleFunc("/api", GetHandler).Methods(http.MethodGet)
	router.HandleFunc("/api", PostHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/{id}", PatchHandler).Methods(http.MethodPatch)
	router.HandleFunc("/api/{id}", DeleteHandler).Methods(http.MethodDelete)

	http.ListenAndServe(":8080", router)
}
