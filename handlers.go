package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
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
	var taskResponse Task = req

	// Respond with success
	respondWithJSON(w, http.StatusCreated, taskResponse)
}
func PatchHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL path
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	// Fetch the existing task from the database
	var task Task
	result := DB.First(&task, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			respondWithError(w, http.StatusNotFound, "Task not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "Failed to fetch task")
		}
		return
	}

	// Decode the request body to get the updated fields
	var updates Task
	err = json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Update the task fields if they are provided in the request
	if updates.Task != "" {
		task.Task = updates.Task
	}
	if updates.IsDone != task.IsDone {
		task.IsDone = updates.IsDone
	}

	// Save the updated task to the database
	result = DB.Save(&task)
	if result.Error != nil {
		log.Printf("Database error: %v", result.Error)
		respondWithError(w, http.StatusInternalServerError, "Failed to update task")
		return
	}

	// Respond with the updated task
	respondWithJSON(w, http.StatusOK, task)
}
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL path
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid task ID")
		return
	}
	// Fetch the task from the database
	var task Task
	result := DB.First(&task, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			respondWithError(w, http.StatusNotFound, "Task not found")

		} else {
			respondWithError(w, http.StatusInternalServerError, "Failed to fetch task")
		}
		return
	}
	// Delete the task from the database
	result = DB.Delete(&task)
	if result.Error != nil {
		log.Printf("Database error: %v", result.Error)
		respondWithError(w, http.StatusInternalServerError, "Failed to delete task")
		return
	}

	// Return 204 No Content
	w.WriteHeader(http.StatusNoContent)
}
