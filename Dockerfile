# Use an official Golang image as the build environment
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to cache dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copying the rest of the application code
COPY . .

# Build the Go binary
RUN go build -o go-stocker

# Use a minimal base image for the final container
FROM alpine:latest

# Set the working directory inside the final container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/go-stocker /app/go-stocker

# Expose the application port (optional, if your app runs on a port)
EXPOSE 8080

# Command to run the application
CMD ["/app/go-stocker"]
