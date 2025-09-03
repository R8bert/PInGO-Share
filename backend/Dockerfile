
# Multi-stage build for security and smaller image size
FROM golang:1.23-alpine AS builder

# Set GOTOOLCHAIN to auto to allow newer versions
ENV GOTOOLCHAIN=auto

# Install security updates
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

# Create app user for security
RUN adduser -D -g '' appuser

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY backend/ .

# Build the binary with security flags
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

# Final stage - minimal image
FROM scratch

# Copy certificates and user from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

# Copy binary
COPY --from=builder /app/main /main

# Create uploads directory
COPY --from=builder --chown=appuser:appuser /tmp /uploads

# Use non-root user
USER appuser

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD ["/main", "--health-check"]

# Run the binary
ENTRYPOINT ["/main"]