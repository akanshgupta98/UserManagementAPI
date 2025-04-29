package main

import (
	"fmt"
	"log"
	"net/http"
)

func logReq(hdlr http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[LOG] - API request recieved: %s Method: %s", r.URL, r.Method)
		hdlr.ServeHTTP(w, r)
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	payload := "server is healthy"
	writeJSON(w, payload)
}

type addUserPayload struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type addUserJsonResponse struct {
	Message string         `json:"message"`
	Data    addUserPayload `json:"data"`
}

type User struct {
	FirstName string
	LastName  string
	Email     string
}

var users map[string]User = make(map[string]User)

func createUser(w http.ResponseWriter, r *http.Request) {
	var reqData addUserPayload
	readJSON(r, &reqData)
	switch {
	case reqData.FirstName == "":
		errorJSON(w, fmt.Errorf("first name is required"), http.StatusBadRequest)
		return
	case reqData.LastName == "":
		errorJSON(w, fmt.Errorf("last name is required"), http.StatusBadRequest)
		return
	case reqData.Email == "":
		errorJSON(w, fmt.Errorf("email name is required"), http.StatusBadRequest)
		return

	}
	if _, ok := users[reqData.Email]; ok {
		errorJSON(w, fmt.Errorf("email already exists"))
		return
	} else {
		user := User(reqData)
		users[reqData.Email] = user
	}

	response := addUserJsonResponse{
		Message: "User created successfully",
		Data:    reqData,
	}
	log.Printf("Final data to be sent is: %+v\n", response)
	writeJSON(w, response, http.StatusAccepted)

}

type GetUserPayload struct {
	Data User `json:"data"`
}

func fetchUserByEmail(w http.ResponseWriter, r *http.Request) {

	email := r.URL.Query().Get("email")
	if email == "" {
		errorJSON(w, nil, http.StatusBadRequest)
		return
	}

	if val, ok := users[email]; ok {
		userData := GetUserPayload{
			Data: val,
		}

		writeJSON(w, userData, http.StatusAccepted)
		return

	} else {
		errorJSON(w, nil, http.StatusNotFound)

	}
}
func addUser(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		createUser(w, r)
	case "GET":
		fetchUserByEmail(w, r)
	}
	// if r.Method != "POST" {
	// 	errorJSON(w, nil, http.StatusNotFound)
	// 	return
	// }

}
