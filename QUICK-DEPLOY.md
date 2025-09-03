# 📦 GitHub Container Registry - Quick Reference

## 🚀 Super Quick Deploy (1-Command Setup)

### On Your Server:
```bash
curl -fsSL https://raw.githubusercontent.com/R8bert/PInGO-Share/main/deploy-github-registry.sh | bash
```

---

## 📋 Manual Steps Summary

### 1. Push Code to GitHub
```bash
git add .
git commit -m "Deploy setup"
git push origin main
```
*→ Wait for GitHub Actions to build images*

### 2. On Your Server
```bash
# Download files
curl -O https://raw.githubusercontent.com/R8bert/PInGO-Share/main/docker-compose.github.yml
curl -O https://raw.githubusercontent.com/R8bert/PInGO-Share/main/.env.production.example
cp .env.production.example .env

# Edit environment
nano .env  # Change DB_PASSWORD, JWT_SECRET, ALLOWED_ORIGINS

# Deploy
docker-compose -f docker-compose.github.yml up -d
```

---

## ⚙️ Required .env Changes

```bash
DB_PASSWORD=your_secure_password_here
JWT_SECRET=your-32-character-jwt-secret-here  
ALLOWED_ORIGINS=https://yourdomain.com
```

---

## 🔧 Management Commands

| Task | Command |
|------|---------|
| **Check Status** | `docker-compose -f docker-compose.github.yml ps` |
| **View Logs** | `docker-compose -f docker-compose.github.yml logs -f` |
| **Update App** | `docker-compose -f docker-compose.github.yml pull && docker-compose -f docker-compose.github.yml up -d` |
| **Restart** | `docker-compose -f docker-compose.github.yml restart` |
| **Stop All** | `docker-compose -f docker-compose.github.yml down` |

---

## 🌐 Access Your App

- **URL**: `http://your-server-ip`
- **Check**: GitHub → Packages (see built images)
- **Logs**: `docker-compose -f docker-compose.github.yml logs`

---

## 🔄 Update Process

1. **Push new code** to GitHub
2. **Wait** for GitHub Actions to build
3. **On server**: `docker-compose -f docker-compose.github.yml pull && docker-compose -f docker-compose.github.yml up -d`

---

## 🆘 Troubleshooting

| Issue | Solution |
|-------|----------|
| **Images not found** | Check GitHub Actions completed successfully |
| **Permission denied** | Run: `docker login ghcr.io -u USERNAME` |
| **Port conflict** | Change `FRONTEND_PORT=8080` in .env |
| **Database errors** | Check `DB_PASSWORD` in .env |

---

## 📱 GitHub Actions Status

Check: `https://github.com/R8bert/PInGO-Share/actions`

✅ **"Build and Push to GitHub Container Registry"** must be successful before deploying!
