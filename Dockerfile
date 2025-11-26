# Step 1: Build the Go binary

# Source image
FROM golang:1.25 AS builder

WORKDIR /app

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the app
COPY . . 

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tdl-go

# Step 2: Create final lightweight image
FROM alpine:latest

# Add CA certificates (needed if app calls HTTPS APIs)
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/tdl-go .

# Expose the app's port (optional)
EXPOSE 8080

# Run the app
CMD ["./tdl-go"]
