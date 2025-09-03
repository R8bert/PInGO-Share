# 🐳 How to Push PInGO-Share to Docker Hub

## Step-by-Step Guide

### Step 1: Login to Docker Hub

First, login to Docker Hub in your terminal:

```bash
docker login
```

Enter your Docker Hub credentials:
- **Username**: `r8bert`
- **Password**: Your Docker Hub password

You'll see: `Login Succeeded`

### Step 2: Build Your Images with Docker Hub Tags

Run these commands to build both images:

```bash
# Build backend image
docker build -t r8bert/pingo:backend-latest -f ./backend/Dockerfile .

# Build frontend image  
docker build -t r8bert/pingo:frontend-latest ./frontend/pingo/
```

### Step 3: Push Images to Docker Hub

Push both images:

```bash
# Push backend
docker push r8bert/pingo:backend-latest

# Push frontend
docker push r8bert/pingo:frontend-latest
```

You'll see progress bars showing the upload.

### Step 4: Verify on Docker Hub

1. Go to https://hub.docker.com/r/r8bert/pingo
2. You should see both tags:
   - `backend-latest`
   - `frontend-latest`

## 🚀 Complete Build and Push Script

I've created a script to do everything automatically:

```bash
#!/bin/bash

echo "🐳 Building and Pushing PInGO-Share to Docker Hub"
echo "================================================"

# Login check
echo "📋 Checking Docker Hub login..."
if ! docker info | grep -q "Username"; then
    echo "⚠️  Please login to Docker Hub first:"
    docker login
fi

# Build backend
echo "🔨 Building backend image..."
docker build -t r8bert/pingo:backend-latest -f ./backend/Dockerfile .

# Build frontend
echo "🔨 Building frontend image..."
docker build -t r8bert/pingo:frontend-latest ./frontend/pingo/

# Push backend
echo "📤 Pushing backend to Docker Hub..."
docker push r8bert/pingo:backend-latest

# Push frontend
echo "📤 Pushing frontend to Docker Hub..."
docker push r8bert/pingo:frontend-latest

echo "✅ Done! Images are now on Docker Hub:"
echo "   - r8bert/pingo:backend-latest"
echo "   - r8bert/pingo:frontend-latest"
echo ""
echo "🌐 View at: https://hub.docker.com/r/r8bert/pingo"
```

Save this as `push-to-dockerhub.sh` and run it!

## 🎯 Quick Commands Summary

```bash
# 1. Login
docker login

# 2. Build both images
docker build -t r8bert/pingo:backend-latest -f ./backend/Dockerfile .
docker build -t r8bert/pingo:frontend-latest ./frontend/pingo/

# 3. Push both images
docker push r8bert/pingo:backend-latest
docker push r8bert/pingo:frontend-latest
```

## 🔄 To Update Your Images Later

When you make changes to your code:

1. Rebuild the images (same commands as above)
2. Push again (Docker will only upload the changed layers)
3. In Portainer, update your stack to pull new images

## 📝 Notes

- **Image naming**: Using `r8bert/pingo:backend-latest` and `r8bert/pingo:frontend-latest`
- **Repository**: All images go to one repository: `r8bert/pingo`
- **Tags**: Different tags for backend and frontend
- **Public**: Docker Hub repos are public by default (free tier)

## ❓ Troubleshooting

**"Access Denied" error**: Make sure you're logged in with `docker login`

**"Repository not found"**: The repository is created automatically on first push

**Slow upload**: Docker uploads in layers, first push takes longest

**"Tag already exists"**: That's normal, it will overwrite with the new version
