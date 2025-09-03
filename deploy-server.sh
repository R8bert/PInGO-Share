#!/bin/bash

# Manual deployment script for servers without Docker registry
# Run this script on your server

set -e

# Configuration - Edit these values
REPO_URL="https://github.com/R8bert/PInGO-Share.git"
DEPLOY_PATH="/opt/pingo-share"
BRANCH="main"

echo "🚀 Deploying PInGO-Share from GitHub..."

# Create deployment directory if it doesn't exist
sudo mkdir -p $DEPLOY_PATH
sudo chown $USER:$USER $DEPLOY_PATH

# Navigate to deployment directory
cd $DEPLOY_PATH

# Clone or update repository
if [ -d ".git" ]; then
    echo "📥 Updating existing repository..."
    git fetch origin
    git reset --hard origin/$BRANCH
    git clean -fd
else
    echo "📥 Cloning repository..."
    git clone $REPO_URL .
    git checkout $BRANCH
fi

# Setup environment file if it doesn't exist
if [ ! -f ".env" ]; then
    echo "⚙️ Setting up environment file..."
    cp .env.production.example .env
    echo ""
    echo "⚠️  IMPORTANT: Edit the .env file with your configuration:"
    echo "   nano .env"
    echo ""
    echo "Required changes:"
    echo "- DB_PASSWORD: Set a secure database password"
    echo "- JWT_SECRET: Set a secure JWT secret (32+ characters)"
    echo "- ALLOWED_ORIGINS: Set your domain (e.g., https://yourdomain.com)"
    echo ""
    read -p "Press Enter after editing .env file..."
fi

# Stop existing containers
echo "🛑 Stopping existing containers..."
docker-compose -f docker-compose.production.yml down || true

# Build new images
echo "🔨 Building Docker images..."
docker-compose -f docker-compose.production.yml build

# Start services
echo "🎯 Starting services..."
docker-compose -f docker-compose.production.yml up -d

# Show status
echo "📊 Service status:"
docker-compose -f docker-compose.production.yml ps

# Clean up old images
echo "🧹 Cleaning up old images..."
docker image prune -f

echo ""
echo "✅ Deployment complete!"
echo "🌐 Your application should be available at: http://your-server-ip"
echo ""
echo "To check logs: docker-compose -f docker-compose.production.yml logs -f"
echo "To restart: docker-compose -f docker-compose.production.yml restart"
