# 🎯 PInGO-Share

A file sharing application with Vue.js frontend and Go backend, deployed using Docker and GitHub Container Registry.

## 🚀 Quick Deploy

### Local Development
```bash
docker-compose up -d
```

### Production Deployment (GitHub Container Registry)
```bash
# Automated deployment script
./deploy-github-registry.sh

# Or manual local build
./local-build-deploy.sh
```

## 📁 Project Structure

```
├── backend/              # Go backend source
├── frontend/pingo/       # Vue.js frontend source  
├── .github/workflows/    # GitHub Actions (auto-builds images)
├── docker-compose.yml    # Local development
├── docker-compose.github.yml  # Production with GitHub Container Registry
└── .env                  # Environment configuration
```

## ⚙️ Configuration

Edit `.env` file:
```bash
DB_PASSWORD=your_secure_password
JWT_SECRET=your_32_character_secret
ALLOWED_ORIGINS=https://yourdomain.com
```

## 🔧 Management

```bash
# Check status
docker-compose -f docker-compose.github.yml ps

# View logs  
docker-compose -f docker-compose.github.yml logs -f

# Update (after pushing new code)
docker-compose -f docker-compose.github.yml pull && docker-compose -f docker-compose.github.yml up -d

# Stop
docker-compose -f docker-compose.github.yml down
```

## 🌐 Access

- **Local**: http://localhost:3000
- **Production**: http://your-server-ip

## 📦 Deployment Methods

1. **GitHub Container Registry** (Recommended)
   - Push code → GitHub Actions builds images → Deploy with `deploy-github-registry.sh`

2. **Local Build**  
   - Build and deploy locally with `local-build-deploy.sh`

## 🛠️ Tech Stack

- **Frontend**: Vue.js 3 + TypeScript + Vite
- **Backend**: Go + Gin Framework  
- **Database**: PostgreSQL
- **Deployment**: Docker + GitHub Container Registry
