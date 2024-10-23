package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	r := mux.NewRouter()

	r.HandleFunc("/user/{id:[0-9]+}", getUserHandler).Methods(http.MethodGet)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Bad Request: invalid user ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad Request: invalid user ID", http.StatusBadRequest)
		return
	}

	// Fetching the user data (simulated)
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
