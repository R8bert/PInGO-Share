#!/bin/bash

echo "🐳 Building and Pushing PInGO-Share to Docker Hub"
echo "================================================"

# Check if logged in
echo "📋 Checking Docker Hub login..."
if ! docker system info 2>/dev/null | grep -q "Username"; then
    echo "⚠️  You need to login to Docker Hub first!"
    echo "Run: docker login -u r8bert"
    echo "Or visit: https://login.docker.com/activate for web login"
    exit 1
fi

echo "✅ Docker Hub login verified"

# Build backend
echo "🔨 Building backend image..."
docker build -t r8bert/pingo:backend-latest -f ./backend/Dockerfile .

if [ $? -eq 0 ]; then
    echo "✅ Backend image built successfully"
else
    echo "❌ Backend build failed"
    exit 1
fi

# Build frontend
echo "🔨 Building frontend image..."
docker build -t r8bert/pingo:frontend-latest ./frontend/pingo/

if [ $? -eq 0 ]; then
    echo "✅ Frontend image built successfully"
else
    echo "❌ Frontend build failed"
    exit 1
fi

# Push backend
echo "📤 Pushing backend to Docker Hub..."
docker push r8bert/pingo:backend-latest

if [ $? -eq 0 ]; then
    echo "✅ Backend pushed successfully"
else
    echo "❌ Backend push failed"
    exit 1
fi

# Push frontend
echo "📤 Pushing frontend to Docker Hub..."
docker push r8bert/pingo:frontend-latest

if [ $? -eq 0 ]; then
    echo "✅ Frontend pushed successfully"
else
    echo "❌ Frontend push failed"
    exit 1
fi

echo ""
echo "🎉 SUCCESS! All images pushed to Docker Hub:"
echo "   📦 r8bert/pingo:backend-latest"
echo "   📦 r8bert/pingo:frontend-latest"
echo ""
echo "🌐 View your repository: https://hub.docker.com/r/r8bert/pingo"
echo ""
echo "🚀 Ready for Portainer deployment!"
echo "   Your portainer-stack.yml is already configured to use these images."
