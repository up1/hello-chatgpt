package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type Response struct {
	Message string `json:"message"`
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	// Check that the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body as a JSON object
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the request parameters
	if req.User == "" || req.Password == "" {
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	// Perform the registration
	err = doRegister(req.User, req.Password)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Return success response
	resp := Response{Message: "Register success"}
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/v1/register", registerHandler)
	fmt.Println("Listening on :8080...")
	http.ListenAndServe(":8080", nil)
}