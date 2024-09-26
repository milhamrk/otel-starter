package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"otel-starter/tracing"

	"go.opentelemetry.io/otel"
)

func main() {
	// Initialize OpenTelemetry Tracing
	tp := tracing.InitTracer()
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Fatal("Error shutting down tracer provider: ", err)
		}
	}()

	tracer := otel.Tracer("otel-starter")
	ctx, span := tracer.Start(context.Background(), "main-operation")
	defer span.End()

	// Simulate work
	doWork(ctx)
	fmt.Println("Tracing completed.")
}

func doWork(ctx context.Context) {
	tracer := otel.Tracer("otel-starter")
	_, span := tracer.Start(ctx, "doWork")
	defer span.End()

	// Simulate some work
	time.Sleep(2 * time.Second)
	fmt.Println("Work is done.")
}
