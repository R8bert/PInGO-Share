#!/bin/bash

# PInGO-Share Docker Compose Management Script

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if Docker and Docker Compose are installed
check_dependencies() {
    print_status "Checking dependencies..."
    
    if ! command -v docker &> /dev/null; then
        print_error "Docker is not installed. Please install Docker first."
        exit 1
    fi
    
    if ! command -v docker-compose &> /dev/null && ! docker compose version &> /dev/null; then
        print_error "Docker Compose is not installed. Please install Docker Compose first."
        exit 1
    fi
    
    print_success "Dependencies check passed!"
}

# Create .env file if it doesn't exist
setup_env() {
    if [ ! -f .env ]; then
        print_warning ".env file not found. Creating from .env.example..."
        cp .env.example .env
        print_success ".env file created. Please review and update the values before running."
        echo
        print_warning "IMPORTANT: Update the following in .env file:"
        echo "  - JWT_SECRET: Generate a strong secret key"
        echo "  - DB_PASSWORD: Set a secure database password"
        echo "  - ALLOWED_ORIGINS: Set your domain for CORS"
        echo
    fi
}

# Function to build all services
build_services() {
    print_status "Building Docker images..."
    docker-compose build --no-cache
    print_success "Build completed!"
}

# Function to start services
start_services() {
    print_status "Starting PInGO-Share services..."
    docker-compose up -d
    print_success "Services started!"
    
    print_status "Service URLs:"
    echo "  🌐 Frontend: http://localhost:3000"
    echo "  🔧 Backend API: http://localhost:8080"
    echo "  🗄️  Database: localhost:5432"
    echo
    print_status "To view logs: ./docker-run.sh logs"
    print_status "To stop services: ./docker-run.sh stop"
}

# Function to stop services
stop_services() {
    print_status "Stopping services..."
    docker-compose down
    print_success "Services stopped!"
}

# Function to restart services
restart_services() {
    print_status "Restarting services..."
    docker-compose restart
    print_success "Services restarted!"
}

# Function to show logs
show_logs() {
    if [ -n "$2" ]; then
        print_status "Showing logs for $2..."
        docker-compose logs -f "$2"
    else
        print_status "Showing logs for all services..."
        docker-compose logs -f
    fi
}

# Function to show service status
show_status() {
    print_status "Service status:"
    docker-compose ps
}

# Function to clean up
cleanup() {
    print_warning "This will remove all containers, networks, and volumes. Are you sure? (y/N)"
    read -r response
    if [[ "$response" =~ ^([yY][eE][sS]|[yY])$ ]]; then
        print_status "Cleaning up..."
        docker-compose down -v --remove-orphans
        docker system prune -f
        print_success "Cleanup completed!"
    else
        print_status "Cleanup cancelled."
    fi
}

# Function to run in production mode
production_mode() {
    print_status "Starting in production mode..."
    docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d
    print_success "Production services started!"
}

# Function to show help
show_help() {
    echo "PInGO-Share Docker Management Script"
    echo
    echo "Usage: $0 [COMMAND]"
    echo
    echo "Commands:"
    echo "  start      Start all services (frontend, backend, database)"
    echo "  stop       Stop all services"
    echo "  restart    Restart all services"
    echo "  build      Build all Docker images"
    echo "  logs       Show logs for all services"
    echo "  logs [service]  Show logs for specific service (frontend, backend, postgres_pingo)"
    echo "  status     Show service status"
    echo "  cleanup    Remove all containers, networks, and volumes"
    echo "  prod       Start in production mode"
    echo "  help       Show this help message"
    echo
    echo "Examples:"
    echo "  $0 start                 # Start all services"
    echo "  $0 logs backend          # Show backend logs"
    echo "  $0 build                 # Rebuild all images"
}

# Main script logic
main() {
    check_dependencies
    setup_env
    
    case "${1:-start}" in
        "start")
            start_services
            ;;
        "stop")
            stop_services
            ;;
        "restart")
            restart_services
            ;;
        "build")
            build_services
            ;;
        "logs")
            show_logs "$@"
            ;;
        "status")
            show_status
            ;;
        "cleanup")
            cleanup
            ;;
        "prod")
            production_mode
            ;;
        "help"|"-h"|"--help")
            show_help
            ;;
        *)
            print_error "Unknown command: $1"
            echo
            show_help
            exit 1
            ;;
    esac
}

# Run main function with all arguments
main "$@"
