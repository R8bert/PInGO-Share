#!/bin/bash

# PInGO-Share Development Helper Script

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m'

print_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

case "${1:-help}" in
    "dev-frontend")
        print_info "Starting frontend development server..."
        cd frontend/pingo && npm run dev
        ;;
    "dev-backend")
        print_info "Starting backend development server..."
        cd backend && go run main.go
        ;;
    "dev-db")
        print_info "Starting only database for development..."
        docker-compose up postgres_pingo -d
        print_success "Database started on localhost:5432"
        ;;
    "dev-full")
        print_info "Starting full development environment..."
        echo "1. Starting database..."
        docker-compose up postgres_pingo -d
        sleep 3
        echo "2. Starting backend..."
        cd backend && go run main.go &
        BACKEND_PID=$!
        sleep 2
        echo "3. Starting frontend..."
        cd ../frontend/pingo && npm run dev &
        FRONTEND_PID=$!
        
        print_success "Development environment started!"
        echo "Frontend: http://localhost:5173"
        echo "Backend: http://localhost:8080"
        echo "Database: localhost:5432"
        echo ""
        echo "Press Ctrl+C to stop all services"
        
        trap "kill $BACKEND_PID $FRONTEND_PID; docker-compose stop postgres_pingo" EXIT
        wait
        ;;
    "install-frontend")
        print_info "Installing frontend dependencies..."
        cd frontend/pingo && npm install
        ;;
    "install-backend")
        print_info "Installing backend dependencies..."
        cd backend && go mod download
        ;;
    "build-frontend")
        print_info "Building frontend for production..."
        cd frontend/pingo && npm run build
        ;;
    "build-backend")
        print_info "Building backend for production..."
        cd backend && go build -o pingo-server main.go
        ;;
    *)
        echo "PInGO-Share Development Helper"
        echo ""
        echo "Usage: $0 [COMMAND]"
        echo ""
        echo "Development Commands:"
        echo "  dev-frontend      Start frontend dev server (port 5173)"
        echo "  dev-backend       Start backend dev server (port 8080)"
        echo "  dev-db           Start only database"
        echo "  dev-full         Start complete development environment"
        echo ""
        echo "Build Commands:"
        echo "  install-frontend  Install frontend dependencies"
        echo "  install-backend   Install backend dependencies"
        echo "  build-frontend    Build frontend for production"
        echo "  build-backend     Build backend for production"
        echo ""
        echo "For Docker deployment, use: ./docker-run.sh"
        ;;
esac
