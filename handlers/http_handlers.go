package handlers

import (
	"fmt"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
)

// PingHandler responds with "pong"
func PingHandler(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("otel-starter")
	ctx, span := tracer.Start(r.Context(), "PingHandler")
	defer span.End()

	time.Sleep(100 * time.Millisecond) // Simulate work
	fmt.Fprintf(w, "pong")
}

// HelloHandler responds with "Hello, World!" and echoes the query
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("otel-starter")
	ctx, span := tracer.Start(r.Context(), "HelloHandler")
	defer span.End()

	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	time.Sleep(200 * time.Millisecond) // Simulate work
	fmt.Fprintf(w, "Hello, %s!", name)
}
