package main

import (
	"log"
	"net/http"

	"github.com/1995parham-teaching/go-lecture/httpserver/handler"
)

func main() {
	h := handler.Hello{
		From: "Golang",
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello", h.Get)
	mux.HandleFunc("POST /hello", h.Post)
	mux.HandleFunc("GET /hello/{username}", h.User)

	if err := http.ListenAndServe("0.0.0.0:1373", mux); err != nil {
		log.Fatal(err)
	}
}
