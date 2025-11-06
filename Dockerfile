# Build stage
FROM golang:1.24-alpine AS builder

# Set environment variables
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the CLI binary
RUN go build -o weather .

FROM alpine:3.21

# Add SSL certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /bin

# Copy compiled binary from builder stage
COPY --from=builder /app/weather .

# Optional: copy .env (API_URL is public so its okay)
COPY .env .  

# Default command (runs the CLI)
ENTRYPOINT ["./weather"]
