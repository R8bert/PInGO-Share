# 🚀 Server Deployment Guide (No Docker Registry Required)

## Method 1: GitHub Actions Auto-Deploy (Recommended)

### Setup Steps:

1. **Add GitHub Secrets** (Go to your repo → Settings → Secrets and variables → Actions):
   ```
   SERVER_SSH_KEY: Your private SSH key content
   SERVER_HOST: your-server-ip-or-domain
   SERVER_USER: your-server-username (e.g., ubuntu, root)
   DEPLOY_PATH: /opt/pingo-share
   ```

2. **Generate SSH Key** (on your local machine):
   ```bash
   ssh-keygen -t ed25519 -C "github-deploy" -f ~/.ssh/github_deploy
   # Copy public key to server
   ssh-copy-id -i ~/.ssh/github_deploy.pub user@your-server
   # Copy private key content to GitHub Secrets (SERVER_SSH_KEY)
   cat ~/.ssh/github_deploy
   ```

3. **Push to GitHub** - Deployment will happen automatically!

### Usage:
- **Auto-deploy**: Push to `main` branch
- **Manual deploy**: Go to Actions tab → Deploy PInGO-Share → Run workflow

---

## Method 2: Manual Git Clone Deploy

### On Your Server:

1. **Download and run the deploy script**:
   ```bash
   curl -O https://raw.githubusercontent.com/R8bert/PInGO-Share/main/deploy-server.sh
   chmod +x deploy-server.sh
   ./deploy-server.sh
   ```

2. **Or manually**:
   ```bash
   # Clone repository
   git clone https://github.com/R8bert/PInGO-Share.git /opt/pingo-share
   cd /opt/pingo-share
   
   # Setup environment
   cp .env.production.example .env
   nano .env  # Edit with your values
   
   # Deploy
   docker-compose -f docker-compose.production.yml up -d --build
   ```

### To Update:
```bash
cd /opt/pingo-share
git pull origin main
docker-compose -f docker-compose.production.yml up -d --build
```

---

## Method 3: GitHub Container Registry (Free!)

### Setup:

1. **Enable GitHub Container Registry**:
   - Go to your repo → Settings → General
   - Scroll to "Danger Zone" → Change repository visibility → Make public (for free registry)
   - Or keep private (works with authentication)

2. **Push code** - Images will auto-build via GitHub Actions

3. **On your server**:
   ```bash
   # Login to GitHub Container Registry (if private repo)
   echo $GITHUB_TOKEN | docker login ghcr.io -u USERNAME --password-stdin
   
   # Deploy using pre-built images
   curl -O https://raw.githubusercontent.com/R8bert/PInGO-Share/main/docker-compose.github.yml
   curl -O https://raw.githubusercontent.com/R8bert/PInGO-Share/main/.env.production.example
   cp .env.production.example .env
   nano .env  # Edit your values
   
   docker-compose -f docker-compose.github.yml up -d
   ```

---

## 🔧 Environment Configuration

**Edit `.env` file on your server with these values:**

```bash
# Database (CHANGE THESE!)
DB_USER=user
DB_PASSWORD=your-super-secure-database-password
DB_NAME=filesharing

# Security (CHANGE THIS!)
JWT_SECRET=your-32-character-jwt-secret-here

# Domain/CORS (CHANGE THIS!)
ALLOWED_ORIGINS=https://yourdomain.com,https://www.yourdomain.com

# Port (optional)
FRONTEND_PORT=80
```

---

## 🌐 Domain Setup

### With Nginx Reverse Proxy:
```nginx
server {
    listen 80;
    server_name yourdomain.com;
    
    location / {
        proxy_pass http://localhost:3000;  # If using custom port
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

### Direct Port 80 (easier):
Set `FRONTEND_PORT=80` in `.env` file

---

## 📊 Management Commands

```bash
# Check status
docker-compose -f docker-compose.production.yml ps

# View logs
docker-compose -f docker-compose.production.yml logs -f

# Restart services
docker-compose -f docker-compose.production.yml restart

# Update and redeploy
cd /opt/pingo-share
git pull
docker-compose -f docker-compose.production.yml up -d --build

# Backup database
docker-compose -f docker-compose.production.yml exec postgres_pingo pg_dump -U user filesharing > backup.sql

# Stop everything
docker-compose -f docker-compose.production.yml down
```

---

## 🔍 Troubleshooting

### Check if services are running:
```bash
docker ps
```

### Check logs for errors:
```bash
docker-compose -f docker-compose.production.yml logs backend
docker-compose -f docker-compose.production.yml logs frontend
```

### Common Issues:
- **Port 80 in use**: Change `FRONTEND_PORT=8080` in `.env`
- **Database errors**: Check `DB_PASSWORD` in `.env`
- **CORS errors**: Verify `ALLOWED_ORIGINS` includes your domain
- **Build fails**: Ensure Docker has enough memory (2GB+)

---

## 🏆 Recommended Approach

**For beginners**: Use **Method 2** (Manual Git Clone)
**For automation**: Use **Method 1** (GitHub Actions)
**For advanced users**: Use **Method 3** (GitHub Container Registry)
