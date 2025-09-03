# Multi-stage build for Vue.js frontend
FROM node:20-alpine AS builder

# Install security updates
RUN apk update && apk add --no-cache git

# Create app user for security
RUN adduser -D -g '' appuser

WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies (including dev dependencies for build)
RUN npm ci

# Copy source code
COPY . .

# Build the application
RUN npm run build

# Production stage - serve with nginx
FROM nginx:1.25-alpine

# Install security updates and curl for health checks
RUN apk update && apk upgrade && apk add --no-cache curl

# Copy built assets from builder stage
COPY --from=builder /app/dist /usr/share/nginx/html

# Copy nginx configuration
COPY nginx.conf /etc/nginx/nginx.conf

# Set permissions for static files and nginx cache directories
RUN chown -R nginx:nginx /usr/share/nginx/html && \
    chown -R nginx:nginx /var/cache/nginx && \
    chown -R nginx:nginx /var/log/nginx && \
    chown -R nginx:nginx /etc/nginx/conf.d
RUN touch /var/run/nginx.pid && \
    chown -R nginx:nginx /var/run/nginx.pid

# Expose port 80
EXPOSE 80

# Start nginx as nginx user
USER nginx
CMD ["nginx", "-g", "daemon off;"]
