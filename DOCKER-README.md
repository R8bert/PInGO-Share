# PInGO-Share Docker Setup

This document explains how to run PInGO-Share using Docker Compose.

## 🚀 Quick Start

1. **Clone the repository** (if you haven't already):
   ```bash
   git clone <your-repo-url>
   cd PInGO-Share
   ```

2. **Run the setup script**:
   ```bash
   ./docker-run.sh start
   ```

3. **Access the application**:
   - 🌐 Frontend: http://localhost:3000
   - 🔧 Backend API: http://localhost:8080
   - 🗄️ Database: localhost:5432

## 📋 Prerequisites

- Docker (version 20.10 or higher)
- Docker Compose (version 2.0 or higher)

## 🛠️ Available Commands

The `docker-run.sh` script provides easy management of your Docker services:

```bash
# Start all services
./docker-run.sh start

# Stop all services
./docker-run.sh stop

# Restart services
./docker-run.sh restart

# Build/rebuild images
./docker-run.sh build

# View logs
./docker-run.sh logs

# View logs for specific service
./docker-run.sh logs frontend
./docker-run.sh logs backend
./docker-run.sh logs postgres_pingo

# Check service status
./docker-run.sh status

# Clean up (removes containers, networks, volumes)
./docker-run.sh cleanup

# Run in production mode
./docker-run.sh prod

# Show help
./docker-run.sh help
```

## ⚙️ Configuration

### Environment Variables

The first time you run the script, it will create a `.env` file from `.env.example`. Update these values:

```env
# Database Configuration
DB_USER=user
DB_PASSWORD=your-secure-password

# JWT Secret (generate a strong secret)
JWT_SECRET=your-super-secret-jwt-key

# CORS Origins
ALLOWED_ORIGINS=http://localhost:3000
```

### Production Deployment

For production deployment:

1. **Update environment variables**:
   ```bash
   cp .env.example .env
   # Edit .env with production values
   ```

2. **Run in production mode**:
   ```bash
   ./docker-run.sh prod
   ```

## 🏗️ Architecture

The Docker Compose setup includes:

- **Frontend Service**: Vue.js application served by Nginx
- **Backend Service**: Go API server
- **Database Service**: PostgreSQL database
- **Network**: Isolated bridge network for service communication
- **Volumes**: Persistent storage for database and uploads

## 📁 Service Details

### Frontend Service
- **Port**: 3000 (external) → 80 (internal)
- **Technology**: Vue.js + Vite + Nginx
- **Features**: SPA routing, API proxy, static asset caching

### Backend Service
- **Port**: 8080 (internal only, accessed via frontend proxy)
- **Technology**: Go + Gin framework
- **Features**: JWT authentication, file uploads, CORS

### Database Service
- **Port**: 5432 (localhost only)
- **Technology**: PostgreSQL 15 Alpine
- **Features**: Health checks, secure configuration

## 🔧 Development

### Hot Reload Development

For development with hot reload, you can run services separately:

1. **Start only database**:
   ```bash
   docker-compose up postgres_pingo -d
   ```

2. **Run backend locally**:
   ```bash
   cd backend
   go run main.go
   ```

3. **Run frontend locally**:
   ```bash
   cd frontend/pingo
   npm run dev
   ```

### Building Individual Services

```bash
# Build only frontend
docker-compose build frontend

# Build only backend
docker-compose build backend
```

## 🐛 Troubleshooting

### Common Issues

1. **Port conflicts**:
   ```bash
   # Check what's using the ports
   sudo netstat -tulpn | grep :3000
   sudo netstat -tulpn | grep :8080
   ```

2. **Permission issues**:
   ```bash
   # Fix uploads directory permissions
   sudo chown -R $USER:$USER uploads/
   ```

3. **Database connection issues**:
   ```bash
   # Check database logs
   ./docker-run.sh logs postgres_pingo
   ```

4. **Clean restart**:
   ```bash
   ./docker-run.sh cleanup
   ./docker-run.sh build
   ./docker-run.sh start
   ```

### Viewing Logs

```bash
# All services
./docker-run.sh logs

# Specific service
./docker-run.sh logs frontend
./docker-run.sh logs backend
./docker-run.sh logs postgres_pingo

# Follow logs in real-time
docker-compose logs -f
```

### Health Checks

All services include health checks. Check status with:
```bash
./docker-run.sh status
```

## 🔒 Security Notes

- Database is bound to localhost only in development
- Non-root users in all containers
- Security headers in Nginx configuration
- Read-only containers where possible
- Minimal Alpine-based images

## 📦 File Structure

```
PInGO-Share/
├── docker-compose.yml          # Main compose file
├── docker-compose.prod.yml     # Production overrides
├── docker-run.sh              # Management script
├── .env.example               # Environment template
├── backend/
│   ├── Dockerfile            # Backend container
│   └── ...
├── frontend/pingo/
│   ├── Dockerfile           # Frontend container
│   ├── nginx.conf          # Nginx configuration
│   └── ...
└── uploads/                # Persistent uploads
```
