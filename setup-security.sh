#!/bin/bash

# PInGO-Share Security Setup Script
# This script helps set up the backend securely

set -e

echo "🔒 PInGO-Share Security Setup"
echo "=============================="

# Check if .env exists
if [ ! -f .env ]; then
    echo "📝 Creating .env file from template..."
    cp .env.example .env
    echo "✅ .env file created"
else
    echo "⚠️  .env file already exists"
fi

# Generate JWT secret if not set
if grep -q "your-very-long-random-secret-key-here" .env; then
    echo "🔑 Generating secure JWT secret..."
    
    # Generate a secure JWT secret
    JWT_SECRET=$(openssl rand -base64 64 | tr -d '\n')
    
    # Replace in .env file
    sed -i "s/JWT_SECRET=your-very-long-random-secret-key-here-at-least-32-characters-long/JWT_SECRET=$JWT_SECRET/" .env
    
    echo "✅ JWT secret generated and set in .env"
else
    echo "✅ JWT secret already configured"
fi

# Generate database password if default
if grep -q "your-secure-database-password" .env; then
    echo "🔐 Generating secure database password..."
    
    DB_PASSWORD=$(openssl rand -base64 32 | tr -d '\n')
    
    sed -i "s/DB_PASSWORD=your-secure-database-password/DB_PASSWORD=$DB_PASSWORD/" .env
    
    echo "✅ Database password generated and set in .env"
else
    echo "✅ Database password already configured"
fi

# Create necessary directories
echo "📁 Creating necessary directories..."
mkdir -p uploads
mkdir -p logos
mkdir -p backgrounds
mkdir -p avatars

# Set proper permissions
chmod 755 uploads logos backgrounds avatars

echo "✅ Directories created with proper permissions"

# Check if Docker is available
if command -v docker >/dev/null 2>&1; then
    echo "🐳 Docker is available"
    
    if command -v docker-compose >/dev/null 2>&1; then
        echo "🐳 Docker Compose is available"
        
        echo ""
        echo "🚀 Ready to deploy! Run:"
        echo "   docker-compose up -d"
        echo ""
        echo "📋 To check status:"
        echo "   docker-compose ps"
        echo "   docker-compose logs -f backend"
        echo ""
    else
        echo "⚠️  Docker Compose not found. Please install docker-compose"
    fi
else
    echo "⚠️  Docker not found. Please install Docker"
fi

echo ""
echo "🔒 IMPORTANT SECURITY REMINDERS:"
echo "================================"
echo "1. Review and update ALLOWED_ORIGINS in .env for production"
echo "2. Use HTTPS with a reverse proxy (Nginx/Caddy) in production"
echo "3. Never expose ports 8080 or 5432 directly to the internet"
echo "4. Regularly update Docker images and dependencies"
echo "5. Monitor logs for suspicious activity"
echo ""
echo "📖 Read SECURITY.md for complete security guidelines"
echo ""
echo "✨ Setup complete! Your backend is ready for secure deployment."
