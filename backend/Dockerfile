
# Multi-stage build for security and smaller image size
FROM golang:1.23-alpine AS builder

# Set GOTOOLCHAIN to auto to allow newer versions
ENV GOTOOLCHAIN=auto

# Install security updates
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

# Create app user for security with specific UID/GID that matches host
RUN addgroup -g 1000 appgroup && adduser -D -u 1000 -G appgroup appuser

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY backend/ .

# Build the binary with security flags
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

# Final stage - use alpine instead of scratch for file operations
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Create app user with same UID/GID as builder
RUN addgroup -g 1000 appgroup && adduser -D -u 1000 -G appgroup appuser

WORKDIR /app

# Create directories with proper ownership
RUN mkdir -p /app/uploads /app/avatars /app/backgrounds /app/logos && \
    chown -R appuser:appgroup /app

# Copy binary
COPY --from=builder /app/main /app/main

# Use non-root user
USER appuser

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD ["/app/main", "--health-check"]

# Run the binary
ENTRYPOINT ["/app/main"]