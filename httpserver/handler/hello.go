package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/1995parham-teaching/go-lecture/httpserver/request"
)

type Hello struct {
	From   string
	Logger *slog.Logger
}

func (h Hello) User(w http.ResponseWriter, r *http.Request) {
	value := r.PathValue("username")

	h.Logger.Info("read username from path parameter", "username", value)

	w.WriteHeader(http.StatusNoContent)
}

func (h Hello) Get(w http.ResponseWriter, r *http.Request) {
	if value := r.FormValue("hello"); value != "" {
		h.Logger.Info("read hello from query parameter", "hello", value)
	}

	enc, err := json.Marshal(fmt.Sprintf("Hello World from %s", h.From))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(enc)
}

func (h Hello) Post(w http.ResponseWriter, r *http.Request) {
	ct := r.Header.Get("Content-Type")

	if ct != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req request.Name

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	if req.Count == nil {
		h.Logger.Info("There is no count")
	} else {
		h.Logger.Info("There is a count", "count", *req.Count)
	}

	enc, err := json.Marshal(fmt.Sprintf("Hello to %s from %s", req.Name, h.From))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(enc)
}
