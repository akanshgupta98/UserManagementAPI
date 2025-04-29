package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct {
	webPort string
}

func main() {

	app := Config{
		webPort: "8081",
	}

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", app.webPort),
		Handler: app.routes(),
	}
	log.Printf("Server started at: %s", app.webPort)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("unable to start server at: %s", app.webPort)
		return
	}

}
