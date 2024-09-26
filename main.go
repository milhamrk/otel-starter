package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"otel-starter/handlers"
	"otel-starter/tracing"
)

func main() {
	// Initialize OpenTelemetry Tracing
	tp := tracing.InitTracer()
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Fatal("Error shutting down tracer provider: ", err)
		}
	}()

	http.HandleFunc("/ping", handlers.PingHandler)
	http.HandleFunc("/hello", handlers.HelloHandler)

	// Start the server
	server := &http.Server{Addr: ":8080"}

	// Gracefully shutdown on interrupt
	go func() {
		log.Println("Server listening on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe failed: %v", err)
		}
	}()

	// Wait for termination signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	// Shutdown the server
	log.Println("Shutting down server...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
}
