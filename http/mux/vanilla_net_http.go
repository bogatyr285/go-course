package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = map[int]User{
	1: {ID: 1, Name: "John Doe"},
	2: {ID: 2, Name: "Jane Roe"},
}

func main() {
	http.HandleFunc("/user/", userHandler)
	http.ListenAndServe(":8080", nil)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUserHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Crude routing by parsing URL manually
	idStr := strings.TrimPrefix(r.URL.Path, "/user/")
	id, err := strconv.Atoi(idStr)
	if err != nil || idStr == "" {
		http.Error(w, "Bad Request: invalid user ID", http.StatusBadRequest)
		return
	}

	// Fetching the user data
	user, found := users[id]
	if !found {
		http.Error(w, "Not Found: user does not exist", http.StatusNotFound)
		return
	}

	// Send the user data as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// 1. **Crude Routing**:

// 2. **Limited Middleware Support**:

// 3. **Poor Param Handling**:

// 4. **No Context**:

// 5. **Error Handling**:
