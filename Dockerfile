# Dockerfile
# Stage 1: Build the application
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o todo-app ./cmd/main.go

# Stage 2: Run the application
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/todo-app .

EXPOSE 8080

CMD ["./todo-app"]