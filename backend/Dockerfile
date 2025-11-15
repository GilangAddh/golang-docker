# =========================================================
# Stage 1: Build the Go binary
# =========================================================
FROM golang:1.24 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (for caching module download)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the app
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go


FROM alpine:3.19

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/app .

# Expose application port
EXPOSE 8080

# Run application
CMD ["./app"]
