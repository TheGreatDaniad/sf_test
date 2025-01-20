# Use the official Go 1.23 Alpine image as the builder
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install necessary dependencies
RUN apk add --no-cache git

# Copy Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

# Use the Alpine runtime image
FROM alpine:latest

WORKDIR /app

# Install PostgreSQL client tools
RUN apk add --no-cache postgresql-client bash

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Copy OpenAPI spec to the runtime image
COPY internal/api/docs /app/internal/api/docs

# Expose the application port
EXPOSE 8080

# Command to start the application
CMD ["./main"]
