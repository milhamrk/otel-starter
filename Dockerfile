FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o otel-starter

EXPOSE 8080

CMD ["./otel-starter"]
