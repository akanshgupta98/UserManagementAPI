package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func readJSON(r *http.Request, data any) {

	// Read from Request body.

	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		log.Printf("Error in reading from JSON request: %s", err)
		return
	}

}

func writeJSON(w http.ResponseWriter, payload any, status ...int) {
	// Create JSON data.
	jsonBytes, _ := json.MarshalIndent(payload, "", "\t")
	// Write the data onto the http Writer.
	w.Header().Add("Content-Type", "application/json")
	if len(status) == 0 {
		w.WriteHeader(http.StatusAccepted)
	} else {
		w.WriteHeader(status[0])
	}

	w.Write(jsonBytes)
}

func errorJSON(w http.ResponseWriter, err error, status ...int) {

	var errCode int
	var errMsg string
	if len(status) == 0 {
		errCode = http.StatusInternalServerError

	} else {
		errCode = status[0]
	}

	if err != nil {
		errMsg = err.Error()
	}
	http.Error(w, errMsg, errCode)

}
