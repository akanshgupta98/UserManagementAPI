package main

import "net/http"

func (app *Config) routes() http.Handler {

	mux := http.NewServeMux()

	mux.Handle("/", logReq(http.HandlerFunc(index)))
	mux.Handle("/user", logReq(http.HandlerFunc(addUser)))
	mux.Handle("/echo", logReq(http.HandlerFunc(echo)))
	return mux
}
