# Use Go 1.22 image
FROM golang:1.22 as builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the source code
COPY . ./

# Build the Go application
RUN go build -o otel-starter

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./otel-starter"]
