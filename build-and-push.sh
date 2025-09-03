#!/bin/bash

# Build and push Docker images to registry
# Usage: ./build-and-push.sh [registry-url] [tag]

set -e

REGISTRY_URL=${1:-"your-registry.com/your-username"}
TAG=${2:-"latest"}

echo "Building and pushing PInGO-Share images..."
echo "Registry: $REGISTRY_URL"
echo "Tag: $TAG"

# Build backend image
echo "Building backend image..."
docker build -t "$REGISTRY_URL/pingo-share-backend:$TAG" -f ./backend/Dockerfile .

# Build frontend image
echo "Building frontend image..."
docker build -t "$REGISTRY_URL/pingo-share-frontend:$TAG" ./frontend/pingo/

# Push images to registry
echo "Pushing backend image..."
docker push "$REGISTRY_URL/pingo-share-backend:$TAG"

echo "Pushing frontend image..."
docker push "$REGISTRY_URL/pingo-share-frontend:$TAG"

echo "✅ Images built and pushed successfully!"
echo ""
echo "To deploy on server, use:"
echo "  REGISTRY_URL=$REGISTRY_URL TAG=$TAG docker-compose -f docker-compose.production.yml up -d"
