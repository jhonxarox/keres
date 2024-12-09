# Start with a base image for Go
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum for dependency management
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire application source code
COPY . .

# Build the Go application
RUN go build -o main ./cmd/main.go

# Use a smaller image for production
FROM alpine:3.18

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the application will run on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
