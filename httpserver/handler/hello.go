package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"mime"
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
	// Parse the media type so we accept e.g. "application/json; charset=utf-8"
	// instead of only the bare "application/json".
	mediaType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil || mediaType != "application/json" {
		http.Error(w, "expected Content-Type: application/json", http.StatusBadRequest)
		return
	}

	var req request.Name

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(enc)
}
