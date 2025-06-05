# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

RUN apk update && apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o task-service ./cmd/main.go

# Stage 2: Run the binary in a minimal image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy binary, config, and schema
COPY --from=builder /app/task-service .
COPY --from=builder /app/config/config.yaml ./config.yaml
COPY --from=builder /app/pkg/db/schema.sql ./schema.sql

EXPOSE 8080

CMD ["./task-service"]
