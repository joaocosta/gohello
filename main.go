package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":49001", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	var input inputData
	var output outputData
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	output.Success = 1
	println(input.Message)
	json.NewEncoder(w).Encode(output)
}

type inputData struct {
	Message string `json:"message"`
}

type outputData struct {
	Success int `json:"success"`
}
