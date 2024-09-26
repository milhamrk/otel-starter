# OpenTelemetry Starter with Go

This project demonstrates how to set up tracing using OpenTelemetry with Go. It uses an OpenTelemetry collector and Jaeger for visualization.

## Prerequisites

- [Go](https://golang.org/doc/install) 1.21 or later
- [Docker](https://docs.docker.com/get-docker/)

## Features

- Tracing with OpenTelemetry
- Exports traces to Jaeger
- Visualize traces with Jaeger UI

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/otel-starter.git
   cd otel-starter
   ```

2. **Run Docker services:**

Start the OpenTelemetry collector and Jaeger:

    docker-compose up -d

3. **Run the Go application:**

After starting the collector, run the Go application:

    go run main.go

4. **View traces in Jaeger:**

Open your browser and go to `http://localhost:16686` to view the traces in Jaeger.