package main

import "net/http"

func (app *Config) routes() http.Handler {

	mux := http.DefaultServeMux
	// mux.HandleFunc("/", index)

	mux.Handle("/", logReq(http.HandlerFunc(index)))
	mux.Handle("/user", logReq(http.HandlerFunc(addUser)))
	return mux
}
