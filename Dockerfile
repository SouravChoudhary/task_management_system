FROM golang:1.22-alpine AS builder
# Stage 1: Build the Go binary

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

WORKDIR /root/

COPY --from=builder /app/task-service .

COPY config/config.yaml ./config.yaml

EXPOSE 8080

CMD ["./task-service"]
