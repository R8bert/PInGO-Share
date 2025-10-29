#!/bin/bash

# PInGO Share Rust Backend - Quick Start Script

set -e

echo "🦀 PInGO Share Rust Backend Setup"
echo "=================================="
echo ""

# Check if Rust is installed
if ! command -v cargo &> /dev/null; then
    echo "❌ Rust is not installed!"
    echo "📦 Please install Rust from: https://rustup.rs/"
    echo ""
    echo "Run this command:"
    echo "curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh"
    exit 1
fi

echo "✅ Rust is installed"
rustc --version
cargo --version
echo ""

# Check for .env file
if [ ! -f ".env" ]; then
    echo "⚠️  No .env file found!"
    echo "📝 Creating .env file from template..."
    cat > .env << 'EOF'
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=user
DB_PASSWORD=pass
DB_NAME=filesharing

# JWT Secret (MUST be at least 32 characters)
JWT_SECRET=please-change-this-to-a-secure-secret-at-least-32-chars-long

# Server Configuration
SERVER_PORT=8080
ALLOWED_ORIGINS=*

# Logging
RUST_LOG=info
EOF
    echo "✅ Created .env file"
    echo "⚠️  Please update JWT_SECRET and database credentials in .env"
    echo ""
fi

# Check PostgreSQL connection
echo "🔍 Checking PostgreSQL connection..."
source .env
if command -v psql &> /dev/null; then
    if psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -c '\q' 2>/dev/null; then
        echo "✅ PostgreSQL connection successful"
    else
        echo "⚠️  Cannot connect to PostgreSQL"
        echo "   Make sure PostgreSQL is running and credentials are correct"
    fi
else
    echo "⚠️  psql command not found, skipping database check"
fi
echo ""

# Build the project
echo "🔨 Building Rust backend..."
echo "   (First build may take a few minutes)"
cargo build --release

echo ""
echo "✅ Build complete!"
echo ""
echo "🚀 Starting server..."
echo "   Server will be available at: http://localhost:$SERVER_PORT"
echo "   Press Ctrl+C to stop"
echo ""

# Run the server
cargo run --release
