package main

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/1995parham-teaching/go-lecture/httpserver/handler"
)

func main() {
	// Demonstrate slog.NewMultiHandler (new in Go 1.26): fan-out logs to
	// stderr (text) and an in-memory JSON sink at the same time. In a real
	// service the second handler would write to a file or shipper.
	textHandler := slog.NewTextHandler(os.Stderr, nil)
	jsonHandler := slog.NewJSONHandler(io.Discard, &slog.HandlerOptions{Level: slog.LevelInfo})
	logger := slog.New(slog.NewMultiHandler(textHandler, jsonHandler))

	h := handler.Hello{
		From:   "Golang",
		Logger: logger.With("handler", "hello"),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", h.Get)
	mux.HandleFunc("POST /hello", h.Post)
	mux.HandleFunc("GET /hello/{username}", h.User)

	srv := &http.Server{
		Addr:              "0.0.0.0:1373",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	// signal.NotifyContext in Go 1.26 returns a CancelCauseFunc and the
	// context's cancel cause carries the received signal — useful when you
	// want to log *why* you're shutting down.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		logger.Info("http server listening", "addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("http server failed", "error", err.Error())
		}
	}()

	<-ctx.Done()
	logger.Info("shutdown requested", "cause", context.Cause(ctx))

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("graceful shutdown failed", "error", err.Error())
	}
}
