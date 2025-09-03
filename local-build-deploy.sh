#!/bin/bash

# Quick local build and deploy script
# Use this when GitHub Container Registry images aren't ready yet

set -e

echo "🏗️  Local Build and Deploy for PInGO-Share"
echo "=========================================="

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

print_step() {
    echo -e "${BLUE}📋 $1${NC}"
}

print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

# Step 1: Check if we're in the right directory
if [ ! -f "docker-compose.yml" ]; then
    echo "❌ This script must be run from the PInGO-Share project directory"
    echo "Current directory: $(pwd)"
    echo "Expected files: docker-compose.yml, backend/, frontend/"
    exit 1
fi

print_step "Building images locally..."

# Step 2: Build backend
print_warning "Building backend image..."
docker build -t ghcr.io/r8bert/pingo-share-backend:latest -f ./backend/Dockerfile .

# Step 3: Build frontend  
print_warning "Building frontend image..."
docker build -t ghcr.io/r8bert/pingo-share-frontend:latest ./frontend/pingo/

print_success "Images built successfully!"

# Step 4: Setup environment if needed
if [ ! -f ".env" ]; then
    if [ -f ".env.production.example" ]; then
        cp .env.production.example .env
        print_warning "Created .env file from template. Please edit it:"
        echo "  nano .env"
        echo ""
        echo "Required changes:"
        echo "- DB_PASSWORD: Set a secure password"  
        echo "- JWT_SECRET: Set a 32+ character secret"
        echo "- ALLOWED_ORIGINS: Set your domain"
        echo ""
        read -p "Press Enter after editing .env..."
    else
        echo "❌ No .env.production.example found. Creating basic .env..."
        cat > .env << EOF
DB_USER=user
DB_PASSWORD=secure_password_change_me
DB_NAME=filesharing
JWT_SECRET=your-32-character-jwt-secret-change-this
ALLOWED_ORIGINS=http://localhost:3000,http://127.0.0.1:3000
FRONTEND_PORT=80
EOF
        print_warning "Basic .env created. Edit it now:"
        echo "  nano .env"
        read -p "Press Enter after editing..."
    fi
fi

# Step 5: Deploy using github compose file
if [ -f "docker-compose.github.yml" ]; then
    print_step "Deploying with docker-compose.github.yml..."
    docker-compose -f docker-compose.github.yml down 2>/dev/null || true
    docker-compose -f docker-compose.github.yml up -d
    COMPOSE_FILE="docker-compose.github.yml"
elif [ -f "docker-compose.production.yml" ]; then
    print_step "Deploying with docker-compose.production.yml..."
    docker-compose -f docker-compose.production.yml down 2>/dev/null || true
    docker-compose -f docker-compose.production.yml up -d
    COMPOSE_FILE="docker-compose.production.yml"
else
    print_step "Deploying with docker-compose.yml..."
    docker-compose down 2>/dev/null || true
    docker-compose up -d
    COMPOSE_FILE="docker-compose.yml"
fi

# Step 6: Check status
print_step "Checking deployment status..."
sleep 5

docker-compose -f $COMPOSE_FILE ps

print_success "Local build and deployment completed!"

echo ""
echo "🌐 Your application should be available at:"
echo "   http://localhost"
echo "   http://127.0.0.1"

# Get external IP if possible
EXTERNAL_IP=$(curl -s ifconfig.me 2>/dev/null || curl -s icanhazip.com 2>/dev/null || echo "your-server-ip")
if [ "$EXTERNAL_IP" != "your-server-ip" ]; then
    echo "   http://$EXTERNAL_IP"
fi

echo ""
echo "🔧 Management commands:"
echo "  Status: docker-compose -f $COMPOSE_FILE ps"
echo "  Logs:   docker-compose -f $COMPOSE_FILE logs -f"
echo "  Stop:   docker-compose -f $COMPOSE_FILE down"
