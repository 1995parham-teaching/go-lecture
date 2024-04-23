package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/1995parham-teaching/go-lecture/httpserver/request"
)

type Hello struct {
	From string
}

func (h Hello) User(w http.ResponseWriter, r *http.Request) {
	value := r.PathValue("username")

	log.Println(value)

	w.WriteHeader(http.StatusNoContent)
}

func (h Hello) Get(w http.ResponseWriter, r *http.Request) {
	if value := r.FormValue("hello"); value != "" {
		log.Println(value)
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
	var req request.Name

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	if req.Count == nil {
		log.Println("There is no count")
	} else {
		log.Printf("There is a count %d", *req.Count)
	}

	enc, err := json.Marshal(fmt.Sprintf("Hello to %s from %s", req.Name, h.From))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(enc)
}
