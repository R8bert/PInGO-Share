# 📦 GitHub Container Registry Deployment Guide

## 🎯 Complete Step-by-Step Guide

### Phase 1: Setup GitHub Container Registry

#### Step 1: Enable Container Registry for Your Repository

1. **Go to your GitHub repository**: `https://github.com/R8bert/PInGO-Share`

2. **Settings → General → Features**:
   - ✅ Enable "Packages" (should be enabled by default)

3. **Settings → Actions → General**:
   - ✅ Allow GitHub Actions to create and approve pull requests
   - ✅ Read and write permissions for GITHUB_TOKEN

#### Step 2: Push Your Code to GitHub

```bash
# In your local PInGO-Share directory
git add .
git commit -m "Add GitHub Container Registry deployment setup"
git push origin main
```

#### Step 3: Verify Auto-Build

1. **Go to Actions tab** in your GitHub repo
2. **Wait for "Build and Push to GitHub Container Registry" workflow** to complete
3. **Check Packages tab** - you should see:
   - `pingo-share-backend`
   - `pingo-share-frontend`

---

### Phase 2: Server Preparation

#### Step 4: Prepare Your Server

```bash
# SSH into your server
ssh your-username@your-server-ip

# Install Docker if not already installed
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo usermod -aG docker $USER
newgrp docker

# Install Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# Create deployment directory
sudo mkdir -p /opt/pingo-share
sudo chown $USER:$USER /opt/pingo-share
cd /opt/pingo-share
```

#### Step 5: Download Deployment Files

```bash
# Download the GitHub Container Registry docker-compose file
curl -O https://raw.githubusercontent.com/R8bert/PInGO-Share/main/docker-compose.github.yml

# Download environment template
curl -O https://raw.githubusercontent.com/R8bert/PInGO-Share/main/.env.production.example

# Copy template to .env
cp .env.production.example .env
```

#### Step 6: Configure Environment

```bash
# Edit the environment file
nano .env
```

**Required changes in `.env`:**

```bash
# Database Configuration (CHANGE THESE!)
DB_USER=pingo_user
DB_PASSWORD=your_super_secure_database_password_here
DB_NAME=filesharing

# Security (CHANGE THIS! - Must be 32+ characters)
JWT_SECRET=your-very-long-jwt-secret-key-here-32-chars-minimum

# Domain Configuration (CHANGE THIS!)
ALLOWED_ORIGINS=https://yourdomain.com,https://www.yourdomain.com,http://your-server-ip

# Port Configuration
FRONTEND_PORT=80

# GitHub Container Registry (keep as-is)
REGISTRY=ghcr.io
```

**Example `.env` file:**
```bash
DB_USER=pingo_user
DB_PASSWORD=MySecure123Database!Password
DB_NAME=filesharing
JWT_SECRET=this-is-my-super-secure-jwt-secret-key-with-32-characters
ALLOWED_ORIGINS=https://mydomain.com,https://www.mydomain.com
FRONTEND_PORT=80
```

---

### Phase 3: Authentication (For Private Repos)

#### Step 7: GitHub Authentication (Only if repo is private)

```bash
# Create GitHub Personal Access Token
# Go to: GitHub → Settings → Developer settings → Personal access tokens → Tokens (classic)
# Create token with 'read:packages' permission

# Login to GitHub Container Registry
echo "YOUR_GITHUB_TOKEN" | docker login ghcr.io -u YOUR_GITHUB_USERNAME --password-stdin
```

**Note**: If your repository is **public**, you can skip this step!

---

### Phase 4: Deployment

#### Step 8: Deploy the Application

```bash
# Pull the latest images
docker-compose -f docker-compose.github.yml pull

# Start the services
docker-compose -f docker-compose.github.yml up -d

# Check if everything is running
docker-compose -f docker-compose.github.yml ps
```

**Expected output:**
```
NAME                           IMAGE                                    STATUS
pingo-share-postgres_pingo-1   postgres:15-alpine                      Up (healthy)
pingo-share-backend-1          ghcr.io/r8bert/pingo-share-backend      Up (healthy)
pingo-share-frontend-1         ghcr.io/r8bert/pingo-share-frontend     Up
```

#### Step 9: Verify Deployment

```bash
# Check logs
docker-compose -f docker-compose.github.yml logs -f

# Test if frontend is accessible
curl -I http://localhost

# Check if API is responding (should return JSON)
curl http://localhost/api/health || curl http://localhost:8080/health
```

---

### Phase 5: Domain Setup (Optional)

#### Step 10: Configure Domain/Nginx (if using custom domain)

**Option A: Direct Port 80 (simpler)**
```bash
# Already configured if FRONTEND_PORT=80 in .env
# Your app will be available at: http://your-server-ip
```

**Option B: Nginx Reverse Proxy**
```bash
# Install nginx
sudo apt update && sudo apt install nginx

# Create nginx config
sudo nano /etc/nginx/sites-available/pingo-share
```

**Nginx config content:**
```nginx
server {
    listen 80;
    server_name yourdomain.com www.yourdomain.com;

    location / {
        proxy_pass http://localhost:3000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

```bash
# Enable the site
sudo ln -s /etc/nginx/sites-available/pingo-share /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
```

---

### Phase 6: Updates and Management

#### Step 11: Update Your Application

**When you push new code to GitHub:**

```bash
# SSH to your server
ssh your-username@your-server-ip
cd /opt/pingo-share

# Pull latest images (built automatically by GitHub Actions)
docker-compose -f docker-compose.github.yml pull

# Restart with new images
docker-compose -f docker-compose.github.yml up -d

# Clean up old images
docker image prune -f
```

#### Step 12: Management Commands

```bash
# View status
docker-compose -f docker-compose.github.yml ps

# View logs
docker-compose -f docker-compose.github.yml logs -f

# Restart specific service
docker-compose -f docker-compose.github.yml restart backend

# Stop everything
docker-compose -f docker-compose.github.yml down

# Start everything
docker-compose -f docker-compose.github.yml up -d

# Backup database
docker-compose -f docker-compose.github.yml exec postgres_pingo pg_dump -U pingo_user filesharing > backup_$(date +%Y%m%d).sql
```

---

## 🔧 Troubleshooting

### Issue 1: Images not found
```bash
# Check if images exist in GitHub Packages
curl -H "Authorization: token YOUR_GITHUB_TOKEN" \
  https://api.github.com/users/R8bert/packages/container/pingo-share-backend/versions

# Force rebuild images
# Go to GitHub → Actions → Manually run "Build and Push" workflow
```

### Issue 2: Permission denied
```bash
# Re-login to GitHub Container Registry
docker logout ghcr.io
echo "YOUR_GITHUB_TOKEN" | docker login ghcr.io -u R8bert --password-stdin
```

### Issue 3: Port conflicts
```bash
# Change port in .env file
echo "FRONTEND_PORT=8080" >> .env
docker-compose -f docker-compose.github.yml up -d
```

### Issue 4: Database connection failed
```bash
# Check database logs
docker-compose -f docker-compose.github.yml logs postgres_pingo

# Reset database
docker-compose -f docker-compose.github.yml down -v
docker-compose -f docker-compose.github.yml up -d
```

---

## 🎉 Success Checklist

- ✅ GitHub Actions built images successfully
- ✅ Images appear in GitHub Packages
- ✅ Server has Docker installed
- ✅ `.env` file configured with your values
- ✅ All containers are running (`docker-compose ps`)
- ✅ Frontend accessible at `http://your-server-ip`
- ✅ API responding (check logs)

---

## 🚀 Next Steps

1. **Set up HTTPS** with Let's Encrypt:
   ```bash
   sudo apt install certbot python3-certbot-nginx
   sudo certbot --nginx -d yourdomain.com
   ```

2. **Set up monitoring** with Portainer:
   ```bash
   docker run -d -p 9000:9000 --name portainer --restart always \
     -v /var/run/docker.sock:/var/run/docker.sock \
     -v portainer_data:/data portainer/portainer-ce
   ```

3. **Set up automated backups**:
   ```bash
   # Add to crontab
   0 2 * * * cd /opt/pingo-share && docker-compose -f docker-compose.github.yml exec -T postgres_pingo pg_dump -U pingo_user filesharing > /backup/pingo_$(date +\%Y\%m\%d).sql
   ```

Your PInGO-Share application is now deployed using GitHub Container Registry! 🎊
