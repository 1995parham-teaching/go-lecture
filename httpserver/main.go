package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/1995parham-teaching/go-lecture/httpserver/handler"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	h := handler.Hello{
		From:   "Golang",
		Logger: logger.With("handler", "hello"),
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello", h.Get)
	mux.HandleFunc("POST /hello", h.Post)
	mux.HandleFunc("GET /hello/{username}", h.User)

	if err := http.ListenAndServe("0.0.0.0:1373", mux); err != nil {
		logger.Error("http server failed", "error", err.Error())
	}
}
