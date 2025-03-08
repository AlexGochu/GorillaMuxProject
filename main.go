package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Task // Use a slice to store multiple tasks

	// Query all tasks from the database
	result := DB.Find(&tasks)
	if result.Error != nil {
		// Handle database error
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "Failed to retrieve tasks"}`)
		return
	}

	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode tasks into JSON and write to the response
	err := json.NewEncoder(w).Encode(tasks)
	if err != nil {
		// Handle JSON encoding error
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "Failed to encode tasks to JSON"}`)
		return
	}

}
func PostHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the request body into a Task struct
	var req Task
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		// Handle JSON decoding errors
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate the request data
	if req.Task == "" {
		respondWithError(w, http.StatusBadRequest, "Task field is required")
		return
	}

	// Create the task in the database
	result := DB.Create(&req)
	if result.Error != nil {
		// Log the error for debugging
		log.Printf("Database error: %v", result.Error)

		// Handle database errors
		respondWithError(w, http.StatusInternalServerError, "Failed to create task")
		return
	}

	// Respond with success
	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Task created successfully"})
}

func main() {
	InitDB()
	DB.AutoMigrate(&Task{})

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
