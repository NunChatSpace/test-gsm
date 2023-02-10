package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Body struct {
	Data string `json:"data"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body := Body{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Do something with the body
		fmt.Printf("Received body: %+v\n", body)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(body)
	}).Methods("POST")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Reached")
	}).Methods("GET")

	port := ":" + os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	http.ListenAndServe(":8080", r)
}
