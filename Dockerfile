# Stage 1: Builder stage - Builds the Go binary
# Uses the official Go image with version 1.21 for building the application
FROM golang:1.21-alpine AS builder

# Set working directory inside the container
WORKDIR /app

# Install git and ca-certificates (needed for fetching dependencies)
# Git is required if any dependencies come from git repositories
RUN apk add --no-cache git ca-certificates

# Copy go mod and sum files first for better Docker layer caching
# This allows Docker to cache the dependency download layer separately
# Copy go.mod and go.sum (if it exists) - go.sum might not exist if no dependencies
COPY go.mod ./
COPY go.sum* ./

# Download all dependencies
# This step is cached unless go.mod or go.sum changes
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application binary
# CGO_ENABLED=0 creates a static binary that doesn't require C libraries
# -a rebuilds all packages, -installsuffix cgo removes C dependencies
# -o specifies the output binary name and location
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mbongo-chain ./cmd/mbongo-chain

# Stage 2: Runner stage - Creates a minimal image to run the binary
# Uses a minimal Alpine Linux image (only 5MB) for security and size
FROM alpine:latest

# Install ca-certificates for HTTPS requests (if needed in the future)
RUN apk --no-cache add ca-certificates

# Create a non-root user for security best practices
RUN addgroup -g 1000 mbongo && \
    adduser -D -u 1000 -G mbongo mbongo

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/mbongo-chain .

# Change ownership to the non-root user
RUN chown -R mbongo:mbongo /app

# Switch to non-root user
USER mbongo

# Expose port 8080 (for future API functionality)
# This is informational and allows Docker to document which ports the container uses
EXPOSE 8080

# Run the binary
# The application will execute and then exit (for demo purposes)
CMD ["./mbongo-chain"]

